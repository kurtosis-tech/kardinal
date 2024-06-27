package topology

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/portforward"
	"k8s.io/client-go/transport/spdy"
	"net/http"
	"os"
)

const (
	kialiServiceName = "kiali"
	namespaceName    = "istio-system"
)

type Manager struct {
	k8sConfig *rest.Config
}

func NewTopologyManager(k8sConfig *rest.Config) *Manager {
	return &Manager{k8sConfig: k8sConfig}
}

func (tf *Manager) FetchTopology(namespace string) (map[string]*Node, error) {
	stopChan, readyChan := make(chan struct{}, 1), make(chan struct{}, 1)
	go func() {
		err := setupPortForwarding(tf.k8sConfig, stopChan, readyChan)
		if err != nil {
			logrus.Fatalf("Error setting up port forwarding: %v", err)
		}
	}()
	<-readyChan // Wait for port forwarding to be ready

	defer close(stopChan)

	var graph map[string]*Node

	// Fetch the graph data
	graph, err := fetchGraphData(namespace)
	if err != nil {
		return nil, fmt.Errorf("Error fetching graph data: %v", err)
	}
	return graph, nil
}

func fetchGraphData(namespace string) (map[string]*Node, error) {
	fmt.Printf("Fetching graph data for namespace %s...\n", namespace)

	url := fmt.Sprintf("http://localhost:20001/kiali/api/namespaces/graph?duration=60s&graphType=versionedApp&includeIdleEdges=false&injectServiceNodes=true&boxBy=cluster,namespace&appenders=deadNode,istio,serviceEntry,meshCheck,workloadEntry,health&rateGrpc=requests&rateHttp=requests&rateTcp=sent&namespaces=%s", namespace)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch graph data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch graph data: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var graph RawKialiGraph

	if err := json.Unmarshal(body, &graph); err != nil {
		return nil, fmt.Errorf("failed to convert response body into inner representation")
	}

	return graphToNodesMap(&graph), nil
}

// setupPortForwarding this sets up a port forward for Kiali;
// TODO make this less fragile
func setupPortForwarding(config *rest.Config, stopChan, readyChan chan struct{}) error {
	roundTripper, upgrader, err := spdy.RoundTripperFor(config)
	if err != nil {
		return fmt.Errorf("error creating round tripper: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("error creating Kubernetes client: %v", err)
	}

	podName, err := getPodsForSvc(kialiServiceName, namespaceName, clientset)
	if err != nil {
		return err
	}

	// Define the port forwarding request
	req := clientset.CoreV1().RESTClient().Post().
		Resource("pods").
		Namespace(namespaceName).
		Name(podName).
		SubResource("portforward")

	// Set up the port forwarding options
	ports := []string{"20001:20001"}
	dialer := spdy.NewDialer(upgrader, &http.Client{Transport: roundTripper}, "POST", req.URL())

	// Create the port forwarder
	fw, err := portforward.New(dialer, ports, stopChan, readyChan, os.Stdout, os.Stderr)
	if err != nil {
		return fmt.Errorf("error creating port forwarder: %v", err)
	}

	// Start port forwarding
	err = fw.ForwardPorts()
	if err != nil {
		return fmt.Errorf("error forwarding ports: %v", err)
	}

	return nil
}

func getPodsForSvc(serviceName string, namespace string, clientset *kubernetes.Clientset) (string, error) {
	svc, err := clientset.CoreV1().Services(namespace).Get(context.TODO(), serviceName, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}

	// Use the service's selectors to find the pods
	selector := labels.SelectorFromSet(svc.Spec.Selector)
	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: selector.String(),
	})
	if err != nil {
		panic(err.Error())
	}

	if len(pods.Items) == 0 {
		return "", fmt.Errorf("Couldn't find a pod for service '%v' in name space '%v", serviceName, namespace)
	}

	return pods.Items[0].Name, nil
}
