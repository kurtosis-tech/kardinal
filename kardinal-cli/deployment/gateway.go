package deployment

import (
	"context"
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/portforward"
	"k8s.io/client-go/transport/spdy"
)

const (
	namespace           = "istio-system"
	service             = "istio-ingressgateway"
	localPort           = 9080
	istioGatewayPodPort = 80
)

func StartGateway(host string) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("Starting gateway for host: %s", host)

	client, err := createKubernetesClient()
	if err != nil {
		return fmt.Errorf("an error occurred while creating a kubernetes client:\n %v", err)
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
		Addr:    ":8080",
		Handler: proxy,
	}

	go func() {
		log.Println("Starting proxy server on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start proxy server: %v", err)
		}
	}()

	log.Println("Proxy server started on :8080")

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

func findPodForService(client *kubernetes.Clientset) (string, error) {
	log.Printf("Finding pod for service %s in namespace %s", service, namespace)
	svc, err := client.CoreV1().Services(namespace).Get(context.Background(), service, metav1.GetOptions{})
	if err != nil {
		return "", fmt.Errorf("error getting service: %v", err)
	}

	// Convert the service selector to a label selector string
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

	// Return the name of the first pod found
	podName := pods.Items[0].Name
	log.Printf("Found pod %s for service %s", podName, service)
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

	log.Printf("Attempting to connect to pod %s at: %s", podName, serverURL.String())

	dialer := spdy.NewDialer(upgrader, &http.Client{Transport: roundTripper}, http.MethodPost, serverURL)

	ports := []string{fmt.Sprintf("%d:%d", localPort, istioGatewayPodPort)}
	forwarder, err := portforward.New(dialer, ports, stopChan, readyChan, os.Stdout, os.Stderr)
	if err != nil {
		return fmt.Errorf("failed to create port forwarder: %v", err)
	}

	log.Printf("Starting port forwarding for pod %s...", podName)
	return forwarder.ForwardPorts()
}

func createProxy(host string) *httputil.ReverseProxy {
	target, _ := url.Parse(fmt.Sprintf("http://localhost:%d", localPort))
	proxy := httputil.NewSingleHostReverseProxy(target)

	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		req.Host = host // Set the Host header to the provided host
		req.Header.Set("X-Forwarded-Host", host)
		log.Printf("Proxying request to: %s%s", req.Host, req.URL.Path)
	}

	proxy.ModifyResponse = func(resp *http.Response) error {
		// Set cache-control headers
		resp.Header.Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate, max-age=0")
		resp.Header.Set("Pragma", "no-cache")
		resp.Header.Set("Expires", "0")
		log.Printf("Received response with status: %d", resp.StatusCode)
		return nil
	}

	return proxy
}
