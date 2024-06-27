package cluster_manager

import (
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
)

type kubernetesClient struct {
	config          *rest.Config
	clientSet       *kubernetes.Clientset
	dynamicClient   *dynamic.DynamicClient
	discoveryMapper *restmapper.DeferredDiscoveryRESTMapper
}

func newKubernetesClient(config *rest.Config, clientSet *kubernetes.Clientset, dynamicClient *dynamic.DynamicClient, discoveryMapper *restmapper.DeferredDiscoveryRESTMapper) *kubernetesClient {
	return &kubernetesClient{config: config, clientSet: clientSet, dynamicClient: dynamicClient, discoveryMapper: discoveryMapper}
}
