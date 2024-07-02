package fetcher

import (
	"context"
	"encoding/json"
	"github.com/kurtosis-tech/kardinal/libs/manager-kontrol-api/api/golang/types"
	"github.com/kurtosis-tech/stacktrace"
	"github.com/sirupsen/logrus"
	"io"
	"kardinal.kontrol/kardinal-manager/cluster_manager"
	"kardinal.kontrol/kardinal-manager/utils"
	"net/http"
	"net/url"
	"time"
)

const (
	defaultTickerDuration              = time.Second * 5
	fetcherJobDurationSecondsEnvVarKey = "KARDINAL_MANAGER_FETCHER_JOB_DURATION_SECONDS"
)

type fetcher struct {
	clusterManager *cluster_manager.ClusterManager
	configEndpoint string
}

func NewFetcher(clusterManager *cluster_manager.ClusterManager, configEndpoint string) *fetcher {
	return &fetcher{clusterManager: clusterManager, configEndpoint: configEndpoint}
}

func (fetcher *fetcher) Run(ctx context.Context) error {

	fetcherTickerDuration := defaultTickerDuration

	fetcherJobDurationSecondsEnVarValue, err := utils.GetIntFromEnvVar(fetcherJobDurationSecondsEnvVarKey, "fetcher job duration seconds")
	if err != nil {
		logrus.Debugf("an error occurred while getting the fetcher job durations seconds from the env var, using default value '%s'. Error:\n%s", defaultTickerDuration, err)
	}

	if fetcherJobDurationSecondsEnVarValue != 0 {
		fetcherTickerDuration = time.Second * time.Duration(int64(fetcherJobDurationSecondsEnVarValue))
	}

	ticker := time.NewTicker(fetcherTickerDuration)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			logrus.Debugf("New fetcher execution at %s", time.Now())
			if err := fetcher.fetchAndApply(ctx); err != nil {
				return stacktrace.Propagate(err, "Failed to fetch and apply the cluster configuration")
			}
		}
	}
}

func (fetcher *fetcher) fetchAndApply(ctx context.Context) error {
	clusterResources, err := fetcher.getClusterResourcesFromCloud()
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred fetching cluster resources from cloud")
	}

	if err = fetcher.clusterManager.ApplyClusterResources(ctx, clusterResources); err != nil {
		return stacktrace.Propagate(err, "Failed to apply cluster resources '%+v'", clusterResources)
	}

	if err = fetcher.clusterManager.CleanUpClusterResources(ctx, clusterResources); err != nil {
		return stacktrace.Propagate(err, "Failed to clean up cluster resources '%+v'", clusterResources)
	}

	return nil
}

func (fetcher *fetcher) getClusterResourcesFromCloud() (*types.ClusterResources, error) {

	configEndpointURL, err := url.Parse(fetcher.configEndpoint)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred parsing the config endpoint '%s'", fetcher.configEndpoint)
	}

	resp, err := http.Get(configEndpointURL.String())
	if err != nil {
		return nil, stacktrace.Propagate(err, "Error fetching cluster resources from endpoint '%s'", fetcher.configEndpoint)
	}
	defer resp.Body.Close()

	responseBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, stacktrace.Propagate(err, "Error reading the response from '%v'", fetcher.configEndpoint)
	}

	if len(responseBodyBytes) == 0 {
		logrus.Debugf("The cluster resources endpoint '%s' returned an empty body", fetcher.configEndpoint)
		return nil, nil
	}

	var clusterResources *types.ClusterResources

	if err = json.Unmarshal(responseBodyBytes, &clusterResources); err != nil {
		return nil, stacktrace.Propagate(err, "And error occurred unmarshalling the response to a config response object")
	}

	return clusterResources, nil
}
