package cluster_manager

import (
	"istio.io/client-go/pkg/clientset/versioned"
	"kardinal.kontrol/kardinal-manager/topology"
)

// IstIO ontology:
// - virtual services
// 	 - host
//   - routing rules
//   - destination rules
//
// - destination rules
// 		- host
//		- traffic policy
//		- subsets
//
// TODO: implement this ontology later
// - gateways
// - service entries

// use cases IstIo manager needs to support:
// - ability to configure traffic routing rules for services in a cluster
//   - change the distribution of traffic to a service
//   - redirect which service traffic is going to
//   - duplicate traffic to services
//
// - ability to add new versions of a service
//   - updating destination rules

type istioClient struct {
	clientSet *versioned.Clientset

	topologyManager *topology.Manager
}

func newIstioClient(clientSet *versioned.Clientset, topologyManager *topology.Manager) *istioClient {
	return &istioClient{clientSet: clientSet, topologyManager: topologyManager}
}
