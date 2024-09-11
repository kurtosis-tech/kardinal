package kubernetes

import (
	"github.com/kurtosis-tech/stacktrace"
	"k8s.io/client-go/discovery/cached/memory"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func CreateKubernetesClient() (*kubernetesClient, error) {
	var config *rest.Config

	// Load in-cluster configuration
	config, err := rest.InClusterConfig()
	if err != nil {
		// Fallback to out-of-cluster configuration (for local development)
		home := homedir.HomeDir()
		kubeConfig := filepath.Join(home, ".kube", "config")
		config, err = clientcmd.BuildConfigFromFlags("", kubeConfig)
		if err != nil {
			return nil, stacktrace.Propagate(err, "impossible to get kubernetes client config either inside or outside the cluster")
		}
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred while creating kubernetes client using config '%+v'", config)
	}

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred while creating kubernetes dynamic client using config '%+v'", config)
	}

	discoveryClient := memory.NewMemCacheClient(clientSet.Discovery())
	discoveryMapper := restmapper.NewDeferredDiscoveryRESTMapper(discoveryClient)

	kubernetesClientObj := newKubernetesClient(config, clientSet, dynamicClient, discoveryMapper)

	return kubernetesClientObj, nil
}
