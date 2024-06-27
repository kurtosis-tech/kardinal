package fetcher

import (
	"context"
	"github.com/stretchr/testify/require"
	"kardinal.kontrol/kardinal-manager/cluster_manager"
	"testing"
)

// This test can be executed and use Minikube dashboard and Kiali Dashboard to see the changes between prod apply and devInProd apply
// these code is meant for local iteration for now and less for unit testing
func TestVotingAppDemoProdAndDevCase(t *testing.T) {
	clusterManager, err := cluster_manager.CreateClusterManager()
	require.NoError(t, err)

	prodOnlyDemoConfigEndpoint := "https://gist.githubusercontent.com/leoporoli/477b9b95238ffa994fb62849debb9abc/raw/b911cbe28df8cb65bf84834f666f94488937c364/cluster-resources-examples.json"

	prodFetcher := NewFetcher(clusterManager, prodOnlyDemoConfigEndpoint)

	ctx := context.Background()

	err = prodFetcher.fetchAndApply(ctx)
	require.NoError(t, err)

	// Sleep to check the Cluster topology in Minikube and Kiali, prod topology should be created in voting-app namespace
	//time.Sleep(2 * time.Minute)

	devInProdEndpoint := "https://gist.githubusercontent.com/leoporoli/d3e3afb29fa0dcc12738df558b263154/raw/7da19c18d34edf09bd2fe2939134b1d0424d1c2b/cluster-resources-for-dev.json"

	devInProdFetcher := NewFetcher(clusterManager, devInProdEndpoint)

	err = devInProdFetcher.fetchAndApply(ctx)
	require.NoError(t, err)

	// Sleep to check the Cluster topology in Minikube and Kiali, dev topology should be added in voting-app namespace
	//time.Sleep(2 * time.Minute)

	//Executing prodFetcher again to remove the Dev resources
	//err = prodFetcher.fetchAndApply(ctx)
	//require.NoError(t, err)

	// Now you can check that dev components has been removed from the cluster
}
