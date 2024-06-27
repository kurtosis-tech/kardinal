package cluster_manager

import (
	"context"
	"github.com/stretchr/testify/require"
	istio "istio.io/api/networking/v1alpha3"
	"testing"
)

const (
	defaultNamespace = "default"
)

func TestClusterManager_GetVirtualServices(t *testing.T) {
	ctx := context.Background()
	clusterManager, err := getClusterManagerForTesting(t)
	require.NoError(t, err)

	virtualServices, err := clusterManager.GetVirtualServices(ctx, defaultNamespace)
	require.NoError(t, err)
	require.NotEmpty(t, virtualServices)
}

func TestClusterManager_GetTopologyForNameSpace(t *testing.T) {
	clusterManager, err := getClusterManagerForTesting(t)
	require.NoError(t, err)

	graph, err := clusterManager.GetTopologyForNameSpace("ms-demo")
	require.Empty(t, err)
	require.NotNil(t, graph)
}

// This test is to demonstrate using the ClusterManager to accomplish certain workflows
// assumes
// - default k8s namespace contains the services from the sample bookinfo application: https://istio.io/latest/docs/examples/bookinfo/, run
// - a destination rule for reviews service has been preconfigured with one version of reviews
// - a virtual service for reviews service has been preconfigured with one routing rule
func TestClusterManagerWorkflows(t *testing.T) {
	ctx := context.Background()
	clusterManager, err := getClusterManagerForTesting(t)
	require.NoError(t, err)

	// verify that there exists a destination rule for the "reviews" service that only sends traffic to v1
	reviewsDestinationRule, err := clusterManager.GetDestinationRule(ctx, defaultNamespace, "reviews")
	require.NoError(t, err)
	require.NotEmpty(t, reviewsDestinationRule)
	require.NotEmpty(t, reviewsDestinationRule.Spec.Subsets)
	require.Equal(t, reviewsDestinationRule.Spec.Subsets[0].Name, "v1")

	// verify that there exists a virtual service for the "reviews" service
	reviewsVirtualService, err := clusterManager.GetVirtualService(ctx, defaultNamespace, "reviews")
	require.NoError(t, err)
	require.NotEmpty(t, reviewsVirtualService)
	require.NotEmpty(t, reviewsVirtualService.Spec.Http)
	// TODO: may want to implement types in house to manage some of this stuff but for now jus use objects directly
	require.Equal(t, reviewsVirtualService.Spec.Http[0].Route[0].Destination.Host, "reviews")
	require.Equal(t, reviewsVirtualService.Spec.Http[0].Route[0].Destination.Subset, "v1")

	// register a new version of the reviews service
	v2subset := &istio.Subset{
		Name: "v2",
		Labels: map[string]string{
			"version": "v2",
		},
		TrafficPolicy: nil,
	}
	err = clusterManager.AddSubset(ctx, defaultNamespace, "reviews", v2subset)
	require.NoError(t, err)

	// add a routing rule that splits traffic between v1 and v2
	splitTraffic5050Rule := &istio.HTTPRoute{
		Route: []*istio.HTTPRouteDestination{
			{
				Destination: &istio.Destination{
					Host:   "reviews",
					Subset: "v1",
					Port:   nil,
				},
				Weight: 50,
			},
			{
				Destination: &istio.Destination{
					Host:   "reviews",
					Subset: "v2",
					Port:   nil,
				},
				Weight: 50,
			},
		},
	}
	// can consider adjusting the AddRoutingRule api to only take in params we care about to make the api easier to use but again for now, KISS till we know more about use cases
	err = clusterManager.AddRoutingRule(ctx, defaultNamespace, "reviews", splitTraffic5050Rule)
}

// Note: test will only work if kubeconfig is available locally and a cluster is running
// these code is meant for local iteration for now and less for unit testing
func getClusterManagerForTesting(t *testing.T) (*ClusterManager, error) {
	clusterManager, err := CreateClusterManager()
	require.NoError(t, err)
	return clusterManager, nil
}
