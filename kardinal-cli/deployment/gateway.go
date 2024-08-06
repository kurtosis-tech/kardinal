package deployment

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/portforward"
	"k8s.io/client-go/transport/spdy"
)

const (
	namespace           = "istio-system"
	service             = "istio-ingressgateway"
	localPortForIstio   = 9080
	istioGatewayPodPort = 8080
	proxyServerPort     = 9060
	maxRetries          = 10
	retryInterval       = 10 * time.Second
	prodNamespace       = "prod"
)

func StartGateway(host string) error {
	log.Printf("Starting gateway for host: %s", host)

	client, err := createKubernetesClient()
	if err != nil {
		return fmt.Errorf("an error occurred while creating a kubernetes client:\n %v", err)
	}

	// Check for pods in the prod namespace
	err = assertProdNamespaceReady(client.clientSet)
	if err != nil {
		return fmt.Errorf("failed to assert that prod namespace is ready: %v", err)
	}

	// Find a pod for the service
	pod, err := findPodForService(client.clientSet)
	if err != nil {
		return fmt.Errorf("failed to find pod for service: %v", err)
	}

	// Start port forwarding
	stopChan := make(chan struct{}, 1)
	readyChan := make(chan struct{})
	go func() {
		for {
			err := portForwardPod(client.config, pod, stopChan, readyChan)
			if err != nil {
				log.Printf("Port forwarding failed: %v. Retrying in 5 seconds...", err)
				time.Sleep(5 * time.Second)
				continue
			}
			break
		}
	}()

	// Wait for port forwarding to be ready
	<-readyChan

	// Start proxy server
	proxy := createProxy(host)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", proxyServerPort),
		Handler: proxy,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start proxy server: %v", err)
		}
	}()

	log.Printf("Proxy server for host %s started on http://localhost:%d", host, proxyServerPort)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Println("Shutting down...")
	close(stopChan)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %v", err)
	}

	return nil
}

func assertProdNamespaceReady(client *kubernetes.Clientset) error {
	for retry := 0; retry < maxRetries; retry++ {
		pods, err := client.CoreV1().Pods(prodNamespace).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Printf("Error listing pods in prod namespace (attempt %d/%d): %v", retry+1, maxRetries, err)
			time.Sleep(retryInterval)
			continue
		}

		if len(pods.Items) == 0 {
			log.Printf("No pods found in namespace %s (attempt %d/%d)", prodNamespace, retry+1, maxRetries)
			time.Sleep(retryInterval)
			continue
		}

		allReady := true
		for _, pod := range pods.Items {
			if !isPodReady(&pod) {
				allReady = false
				log.Printf("Pod %s is not ready", pod.Name)
				break
			}
		}

		if allReady {
			log.Printf("All pods in namespace %s are ready", prodNamespace)
			return nil
		}

		log.Printf("Waiting for all pods to be ready (attempt %d/%d)", retry+1, maxRetries)
		time.Sleep(retryInterval)
	}

	return fmt.Errorf("failed to assert all pods are ready in namespace %s after %d attempts", prodNamespace, maxRetries)
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

func findPodForService(client *kubernetes.Clientset) (string, error) {
	svc, err := client.CoreV1().Services(namespace).Get(context.Background(), service, metav1.GetOptions{})
	if err != nil {
		return "", fmt.Errorf("error getting service: %v", err)
	}

	var labelSelectors []string
	for key, value := range svc.Spec.Selector {
		labelSelectors = append(labelSelectors, fmt.Sprintf("%s=%s", key, value))
	}
	selector := strings.Join(labelSelectors, ",")

	pods, err := client.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{LabelSelector: selector})
	if err != nil {
		return "", fmt.Errorf("error listing pods: %v", err)
	}

	if len(pods.Items) == 0 {
		return "", fmt.Errorf("no pods found for service %s", service)
	}

	podName := pods.Items[0].Name
	return podName, nil
}

func portForwardPod(config *rest.Config, podName string, stopChan <-chan struct{}, readyChan chan struct{}) error {
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

	ports := []string{fmt.Sprintf("%d:%d", localPortForIstio, istioGatewayPodPort)}
	forwarder, err := portforward.New(dialer, ports, stopChan, readyChan, io.Discard, os.Stderr)
	if err != nil {
		return fmt.Errorf("failed to create port forwarder: %v", err)
	}

	return forwarder.ForwardPorts()
}

func createProxy(host string) *httputil.ReverseProxy {
	target, _ := url.Parse(fmt.Sprintf("http://localhost:%d", localPortForIstio))
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
