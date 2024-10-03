package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"kardinal.kontrol/kardinal-manager/cluster_manager"
	"kardinal.kontrol/kardinal-manager/fetcher"
	"kardinal.kontrol/kardinal-manager/logger"
	"kardinal.kontrol/kardinal-manager/utils"
	"os"
)

const (
	successExitCode                = 0
	clusterConfigEndpointEnvVarKey = "KARDINAL_MANAGER_CLUSTER_CONFIG_ENDPOINT"
)

func main() {

	// Create context
	ctx := context.Background()

	if err := logger.ConfigureLogger(); err != nil {
		logrus.Fatal("An error occurred configuring the logger!\nError was: %s", err)
	}

	configEndpoint, err := utils.GetFromEnvVar(clusterConfigEndpointEnvVarKey, "the config endpoint")
	if err != nil {
		logrus.Fatal("An error occurred getting the config endpoint from the env vars!\nError was: %s", err)
	}

	clusterManager, err := cluster_manager.CreateClusterManager()
	if err != nil {
		logrus.Fatal("An error occurred while creating the cluster manager!\nError was: %s", err)
	}

	fetcher := fetcher.NewFetcher(clusterManager, configEndpoint)

	if err = fetcher.Run(ctx); err != nil {
		logrus.Fatalf("An error occurred while running the fetcher!\nError was: %s", err)
	}

	// No external clients connection so-far
	//if err := server.CreateAndStartRestAPIServer(); err != nil {
	//	logrus.Fatalf("The REST API server is down, exiting!\nError was: %s", err)
	//}

	os.Exit(successExitCode)
}
