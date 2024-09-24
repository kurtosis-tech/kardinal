package deployment

import (
	"context"
	"fmt"
	"hash/fnv"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"slices"
	"strings"
	"syscall"
	"time"

	api_types "github.com/kurtosis-tech/kardinal/libs/cli-kontrol-api/api/golang/types"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/portforward"
	"k8s.io/client-go/transport/spdy"
	gatewayclientset "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned"

	cli_k8s "kardinal.cli/kubernetes"
)

const (
	localPortStartRange = 61000
	proxyPortRangeStart = 59000
	proxyPortRangeEnd   = 60000
	maxRetries          = 10
	retryInterval       = 10 * time.Second
)

func hashStringToRange(s string, maxRange int) int {
	h := fnv.New32a()
	h.Write([]byte(s))
	hashedValue := h.Sum32()
	return int(hashedValue % uint32(maxRange))
}

func findAvailablePortInRange(host string, portsInUse *[]int) (int, error) {
	portRange := proxyPortRangeEnd - proxyPortRangeStart
	if portRange <= 0 {
		logrus.Fatalf("Invalid port range: %d-%d", proxyPortRangeStart, proxyPortRangeEnd)
	}

	portShift := hashStringToRange(host, portRange)
	port := proxyPortRangeStart + portShift
	if slices.Contains(*portsInUse, port) {
		logrus.Warnf("Attention! Port %d is already in use and cannot be uniquely mapped to the host %s. Changing the order of the arguments may result in different port-to-flow associations.\n", port, host)
		for i := 0; i < portRange; i++ {
			port = proxyPortRangeStart + i
			if !slices.Contains(*portsInUse, port) {
				return port, nil
			}
			logrus.Infof("Port %d is already in use. Trying next port...\n", port)
		}
		return -1, fmt.Errorf("no available ports in the range %d-%d", proxyPortRangeStart, proxyPortRangeEnd)
	}
	return port, nil
}

func StartGateway(hostFlowIdMap []api_types.IngressAccessEntry) error {
	k8sConfig, err := cli_k8s.GetConfig()
	if err != nil {
		return fmt.Errorf("an error occurred while creating a kubernetes client:\n %v", err)
	}
	client, err := cli_k8s.CreateKubernetesClient(k8sConfig)
	if err != nil {
		return fmt.Errorf("an error occurred while creating a kubernetes client:\n %v", err)
	}
	gatewayClient, err := cli_k8s.CreateGatewayApiClient(k8sConfig)
	if err != nil {
		return fmt.Errorf("an error occurred while creating a kubernetes client:\n %v", err)
	}

	servers := make([]*http.Server, 0)
	ports := make([]int, 0)
	stopChan := make(chan struct{}, 1)

	for entryIx, entry := range hostFlowIdMap {

		localPort := int32(localPortStartRange + entryIx)
		host := entry.Hostname

		logrus.Printf("Starting gateway for host: %s", host)

		err = assertBaselineNamespaceReady(client.GetClientSet(), entry.FlowId, entry.Namespace)
		if err != nil {
			return fmt.Errorf("failed to assert that baseline namespace is ready: %v", err)
		}

		// Find a pod for the service
		pod, port, namespace, err := findPodForService(client.GetClientSet(), gatewayClient, entry)
		if err != nil {
			// return fmt.Errorf("failed to find pod for service: %v", err)
			logrus.Errorf("failed to find pod for service: %v", err)
			continue
		}

		// Start port forwarding
		readyChan := make(chan struct{})
		go func() {
			for {
				err := portForwardPod(client.GetConfig(), namespace, pod, port, stopChan, readyChan, localPort)
				if err != nil {
					logrus.Printf("Port forwarding failed: %v. Retrying in 5 seconds...", err)
					time.Sleep(5 * time.Second)
					continue
				}
				break
			}
		}()

		// Wait for port forwarding to be ready
		<-readyChan

		availablePort, err := findAvailablePortInRange(host, &ports)
		if err != nil {
			return fmt.Errorf("failed to find an available port: %v", err)
		}
		// Start proxy server on the available port
		proxy := createProxy(host, localPort)
		server := &http.Server{
			Addr:    fmt.Sprintf(":%d", availablePort),
			Handler: proxy,
		}
		servers = append(servers, server)

		go func() {
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				logrus.Fatalf("Failed to start proxy server: %v", err)
			}
		}()

		fmt.Printf("\n🔗 Proxy server for host %s started on: http://localhost:%d\n", host, availablePort)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	logrus.Println("Shutting down...")
	close(stopChan)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	for _, server := range servers {
		if err := server.Shutdown(ctx); err != nil {
			logrus.Printf("Server shutdown error: %v", err)
		}
	}

	return nil
}

// TODO move to the kubernetes package
func assertBaselineNamespaceReady(client *kubernetes.Clientset, flowId string, baselineNamespace string) error {
	for retry := 0; retry < maxRetries; retry++ {
		pods, err := client.CoreV1().Pods(baselineNamespace).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			logrus.Printf("Error listing pods in baseline namespace (attempt %d/%d): %v", retry+1, maxRetries, err)
			time.Sleep(retryInterval)
			continue
		}

		if len(pods.Items) == 0 {
			logrus.Printf("No pods found in namespace %s (attempt %d/%d)", baselineNamespace, retry+1, maxRetries)
			time.Sleep(retryInterval)
			continue
		}

		allReady := true
		flowIdFound := false
		for _, pod := range pods.Items {
			if strings.Contains(pod.Name, flowId) {
				flowIdFound = true
			}
			if !isPodReady(&pod) {
				allReady = false
				logrus.Printf("Pod %s is not ready", pod.Name)
				break
			}
		}

		if !flowIdFound {
			logrus.Printf("FlowId %s not found in any pod name (attempt %d/%d)", flowId, retry+1, maxRetries)
			time.Sleep(retryInterval)
			continue
		}

		if allReady && flowIdFound {
			logrus.Printf("All pods in namespace %s are ready and flowId %s found", baselineNamespace, flowId)
			return nil
		}

		logrus.Printf("Waiting for all pods to be ready and flowId to be found (attempt %d/%d)", retry+1, maxRetries)
		time.Sleep(retryInterval)
	}

	return fmt.Errorf("failed to assert all pods are ready and flowId %s found in namespace %s after %d attempts", flowId, baselineNamespace, maxRetries)
}

func isPodReady(pod *corev1.Pod) bool {
	if pod.Status.Phase != corev1.PodRunning {
		return false
	}

	if len(pod.Status.ContainerStatuses) != 2 {
		return false
	}

	for _, containerStatus := range pod.Status.ContainerStatuses {
		if !containerStatus.Ready {
			return false
		}
	}

	return true
}

func findPodForService(client *kubernetes.Clientset, gwClient *gatewayclientset.Clientset, accessEntry api_types.IngressAccessEntry) (string, int32, string, error) {
	var labelSelectors []string
	var port int32
	var namespace string
	if accessEntry.Type == "ingress" {
		ingress, err := client.NetworkingV1().Ingresses(accessEntry.Namespace).Get(context.Background(), accessEntry.Service, metav1.GetOptions{})
		if err != nil {
			return "", -1, "", fmt.Errorf("error getting ingress: %v", err)
		}

		ingressClassName := "nginx"
		if ingress.Spec.IngressClassName != nil {
			ingressClassName = *ingress.Spec.IngressClassName
		}
		ingressClass, err := client.NetworkingV1().IngressClasses().Get(context.Background(), ingressClassName, metav1.GetOptions{})
		if err != nil {
			return "", -1, "", fmt.Errorf("error getting ingress class: %v", err)
		}

		for key, value := range ingressClass.Labels {
			labelSelectors = append(labelSelectors, fmt.Sprintf("%s=%s", key, value))
		}
		namespace = ingressClass.Labels["app.kubernetes.io/instance"]
		port = 80

	} else if accessEntry.Type == "gateway" {
		gw, err := gwClient.GatewayV1().Gateways(accessEntry.Namespace).Get(context.Background(), accessEntry.Service, metav1.GetOptions{})
		if err != nil {
			return "", 0, "", fmt.Errorf("error getting gateway: %v", err)
		}
		port = int32(gw.Spec.Listeners[0].Port)
		labelSelectors = append(labelSelectors, fmt.Sprintf("gateway.networking.k8s.io/gateway-name=%s", accessEntry.Service))
		namespace = accessEntry.Namespace

	} else {
		return "", -1, "", fmt.Errorf("unkown access type: %s", accessEntry.Type)
	}

	selector := strings.Join(labelSelectors, ",")
	pods, err := client.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{LabelSelector: selector})
	if err != nil {
		return "", -1, "", fmt.Errorf("error listing pods: %v", err)
	}

	if len(pods.Items) == 0 {
		return "", -1, "", fmt.Errorf("no pods found for service %s", accessEntry.Service)
	}

	podName := pods.Items[0].Name
	return podName, port, namespace, nil
}

func portForwardPod(config *rest.Config, namespace string, podName string, port int32, stopChan <-chan struct{}, readyChan chan struct{}, localPort int32) error {
	roundTripper, upgrader, err := spdy.RoundTripperFor(config)
	if err != nil {
		return fmt.Errorf("failed to create round tripper: %v", err)
	}

	path := fmt.Sprintf("/api/v1/namespaces/%s/pods/%s/portforward", namespace, podName)
	hostIP := strings.TrimLeft(config.Host, "htps:/")

	serverURL, err := url.Parse(fmt.Sprintf("https://%s%s", hostIP, path))
	if err != nil {
		return fmt.Errorf("failed to parse URL: %v", err)
	}

	dialer := spdy.NewDialer(upgrader, &http.Client{Transport: roundTripper}, http.MethodPost, serverURL)

	ports := []string{fmt.Sprintf("%d:%d", localPort, port)}
	forwarder, err := portforward.New(dialer, ports, stopChan, readyChan, io.Discard, os.Stderr)
	if err != nil {
		return fmt.Errorf("failed to create port forwarder: %v", err)
	}

	return forwarder.ForwardPorts()
}

func createProxy(host string, localPort int32) *httputil.ReverseProxy {
	target, _ := url.Parse(fmt.Sprintf("http://localhost:%d", localPort))
	proxy := httputil.NewSingleHostReverseProxy(target)

	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		req.Host = host // Set the Host header to the provided host
		req.Header.Set("X-Forwarded-Host", host)
	}

	proxy.ModifyResponse = func(resp *http.Response) error {
		// Set cache-control headers
		resp.Header.Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate, max-age=0")
		resp.Header.Set("Pragma", "no-cache")
		resp.Header.Set("Expires", "0")
		return nil
	}

	return proxy
}
