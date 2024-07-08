package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/compose-spec/compose-go/cli"
	"github.com/compose-spec/compose-go/types"
	"github.com/kurtosis-tech/stacktrace"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"kardinal.cli/deployment"
	"kardinal.cli/kontrol"
	"kardinal.cli/tenant"
	"log"
	"net/http"

	api "github.com/kurtosis-tech/kardinal/libs/cli-kontrol-api/api/golang/client"
	api_types "github.com/kurtosis-tech/kardinal/libs/cli-kontrol-api/api/golang/types"
)

const (
	projectName          = "kardinal"
	devMode              = false
	kontrolServiceApiUrl = "ad718d90d54d54dd084dea50a9f011af-1140086995.us-east-1.elb.amazonaws.com"
	kontrolServicePort   = 8080

	kontrolBaseURLTmpl                  = "%s://%s"
	kontrolClusterResourcesEndpointTmpl = "%s/tenant/%s/cluster-resources"

	kontrolTrafficConfigurationURLTmpl = "%s/%s/traffic-configuration"

	localMinikubeKontrolAPIHost = "host.minikube.internal:8080"
	kloudKontrolHost            = "app.kardinal.dev"
	kloudKontrolAPIHost         = kloudKontrolHost + "/api"

	httpSchme   = "http"
	httpsScheme = httpSchme + "s"
)

var composeFile string

var rootCmd = &cobra.Command{
	Use:   "kardinal",
	Short: "Kardinal CLI to manage deployment flows",
}

var flowCmd = &cobra.Command{
	Use:   "flow",
	Short: "Manage deployment flows",
}

var managerCmd = &cobra.Command{
	Use:   "manager",
	Short: "Manage Kardinal manager",
}

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy services",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		services, err := parseComposeFile(composeFile)
		if err != nil {
			log.Fatalf("Error loading compose file: %v", err)
		}
		tenantUuid, err := tenant.GetOrCreateUserTenantUUID()
		if err != nil {
			log.Fatal("Error getting or creating user tenant UUID", err)
		}

		deploy(tenantUuid.String(), services)
	},
}

var createCmd = &cobra.Command{
	Use:   "create [service name] [image name]",
	Short: "Create a new service in development mode",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		serviceName, imageName := args[0], args[1]
		services, err := parseComposeFile(composeFile)
		if err != nil {
			log.Fatalf("Error loading compose file: %v", err)
		}

		tenantUuid, err := tenant.GetOrCreateUserTenantUUID()
		if err != nil {
			log.Fatal("Error getting or creating user tenant UUID", err)
		}

		fmt.Printf("Creating service %s with image %s in development mode...\n", serviceName, imageName)
		createDevFlow(tenantUuid.String(), services, imageName, serviceName)
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete services",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		services, err := parseComposeFile(composeFile)
		if err != nil {
			log.Fatalf("Error loading compose file: %v", err)
		}

		tenantUuid, err := tenant.GetOrCreateUserTenantUUID()
		if err != nil {
			log.Fatal("Error getting or creating user tenant UUID", err)
		}
		deleteFlow(tenantUuid.String(), services)

		fmt.Print("Deleting dev flow")
	},
}

var deployManagerCmd = &cobra.Command{
	Use:       fmt.Sprintf("deploy [kontrol location] accepted values: %s and %s ", kontrol.KontrolLocationLocalMinikube, kontrol.KontrolLocationKloudKontrol),
	Short:     "Deploy Kardinal manager into the cluster",
	ValidArgs: []string{kontrol.KontrolLocationLocalMinikube, kontrol.KontrolLocationKloudKontrol},
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {

		kontroLocation := args[0]

		if err := kontrol.SaveKontrolLocation(kontroLocation); err != nil {
			log.Fatal("Error saving the Kontrol location", err)
		}

		tenantUuid, err := tenant.GetOrCreateUserTenantUUID()
		if err != nil {
			log.Fatal("Error getting or creating user tenant UUID", err)
		}

		if err := deployManager(tenantUuid.String()); err != nil {
			log.Fatal("Error deploying Kardinal manager", err)
		}

		fmt.Printf("Kardinal manager deployed using '%s' Kontrol", kontroLocation)
	},
}

var removeManagerCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove Kardinal manager from the cluster",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if err := removeManager(); err != nil {
			log.Fatal("Error removing Kardinal manager", err)
		}

		fmt.Print("Kardinal manager removed from cluster")
	},
}

func init() {
	rootCmd.AddCommand(flowCmd)
	rootCmd.AddCommand(managerCmd)
	rootCmd.AddCommand(deployCmd)
	flowCmd.AddCommand(createCmd, deleteCmd)
	managerCmd.AddCommand(deployManagerCmd, removeManagerCmd)

	flowCmd.PersistentFlags().StringVarP(&composeFile, "docker-compose", "d", "", "Path to the Docker Compose file")
	flowCmd.MarkPersistentFlagRequired("docker-compose")
	deployCmd.PersistentFlags().StringVarP(&composeFile, "docker-compose", "d", "", "Path to the Docker Compose file")
	deployCmd.MarkPersistentFlagRequired("docker-compose")
}

func Execute() error {
	return rootCmd.Execute()
}

func loadComposeFile(filename string) (*types.Project, error) {
	opts, err := cli.NewProjectOptions([]string{filename},
		cli.WithOsEnv,
		cli.WithDotEnv,
		cli.WithName(projectName),
	)
	if err != nil {
		return nil, err
	}

	project, err := cli.ProjectFromOptions(opts)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func parseComposeFile(composeFile string) ([]types.ServiceConfig, error) {
	project, err := loadComposeFile(composeFile)
	if err != nil {
		log.Fatalf("Error loading compose file: %v", err)
		return nil, err
	}

	fmt.Println("Services in the Docker Compose file:")
	for _, service := range project.Services {
		fmt.Println(service.Name)
	}

	projectYAML, err := project.MarshalJSON()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var dockerCompose map[string]interface{}
	err = json.Unmarshal(projectYAML, &dockerCompose)
	if err != nil {
		log.Fatalf("error: %v", err)
		return nil, err
	}

	return project.Services, nil
}

func createDevFlow(tenantUuid api_types.Uuid, services []types.ServiceConfig, imageLocator, serviceName string) {
	ctx := context.Background()

	body := api_types.PostTenantUuidFlowCreateJSONRequestBody{
		DockerCompose: &services,
		ServiceName:   &serviceName,
		ImageLocator:  &imageLocator,
	}
	client := getKontrolServiceClient()

	resp, err := client.PostTenantUuidFlowCreateWithResponse(ctx, tenantUuid, body)
	if err != nil {
		log.Fatalf("Failed to create dev flow: %v", err)
	}

	fmt.Printf("Response: %s\n", string(resp.Body))
}

func deploy(tenantUuid api_types.Uuid, services []types.ServiceConfig) {
	ctx := context.Background()

	body := api_types.PostTenantUuidDeployJSONRequestBody{
		DockerCompose: &services,
	}
	client := getKontrolServiceClient()

	resp, err := client.PostTenantUuidDeployWithResponse(ctx, tenantUuid, body)
	if err != nil {
		log.Fatalf("Failed to deploy: %v", err)
	}

	fmt.Printf("Response: %s\n", string(resp.Body))

	trafficConfigurationURL, err := getTrafficConfigurationURL(tenantUuid)
	if err != nil {
		log.Fatalf("Failed to get the traffic configuration URL for tenant UUID: %s. Error:\n%v", tenantUuid, err)
	}

	logrus.Infof("Visit: %s", trafficConfigurationURL)
}

func deleteFlow(tenantUuid api_types.Uuid, services []types.ServiceConfig) {
	ctx := context.Background()

	body := api_types.PostTenantUuidFlowDeleteJSONRequestBody{
		DockerCompose: &services,
	}
	client := getKontrolServiceClient()

	resp, err := client.PostTenantUuidFlowDeleteWithResponse(ctx, tenantUuid, body)
	if err != nil {
		log.Fatalf("Failed to delete flow: %v", err)
	}

	fmt.Printf("Response: %s\n", string(resp.Body))
}

func deployManager(tenantUuid api_types.Uuid) error {

	ctx := context.Background()

	clusterResourcesURL, err := getClusterResourcesURL(tenantUuid)
	if err != nil {
		return stacktrace.Propagate(err, "Error getting cluster resources URL")
	}

	if err := deployment.DeployKardinalManagerInCluster(ctx, clusterResourcesURL); err != nil {
		return stacktrace.Propagate(err, "An error occurred deploying Kardinal manager into the cluster with cluster resources URL '%s'", clusterResourcesURL)
	}

	return nil
}

func removeManager() error {
	ctx := context.Background()

	if err := deployment.RemoveKardinalManagerFromCluster(ctx); err != nil {
		return stacktrace.Propagate(err, "An error occurred removing Kardinal manager from the cluster")
	}

	return nil
}

func getKontrolServiceClient() *api.ClientWithResponses {
	if devMode {
		client, err := api.NewClientWithResponses("http://localhost:8080", api.WithHTTPClient(http.DefaultClient))
		if err != nil {
			log.Fatalf("Failed to create client: %v", err)
		}
		return client
	} else {
		client, err := api.NewClientWithResponses(fmt.Sprintf("https://%s", kloudKontrolAPIHost))
		if err != nil {
			log.Fatalf("Failed to create client: %v", err)
		}
		return client
	}
}

func getKontrolBaseURL(useApiHost bool) (string, error) {
	kontrolLocation, err := kontrol.GetKontrolLocation()
	if err != nil {
		return "", stacktrace.Propagate(err, "An error occurred getting the Kontrol location")
	}

	var (
		scheme string
		host   string
	)

	switch kontrolLocation {
	case kontrol.KontrolLocationLocalMinikube:
		scheme = httpSchme
		host = localMinikubeKontrolAPIHost
	case kontrol.KontrolLocationKloudKontrol:
		scheme = httpsScheme
		if useApiHost {
			host = kloudKontrolAPIHost
		} else {
			host = kloudKontrolHost
		}
	default:
		return "", stacktrace.NewError("invalid Kontrol location: %s", kontrolLocation)
	}

	baseURL := fmt.Sprintf(kontrolBaseURLTmpl, scheme, host)

	return baseURL, nil
}

func getTrafficConfigurationURL(tenantUuid api_types.Uuid) (string, error) {

	kontrolBaseURL, err := getKontrolBaseURL(false)
	if err != nil {
		return "", stacktrace.Propagate(err, "An error occurred getting the Kontrol base URL")
	}

	trafficConfigurationURL := fmt.Sprintf(kontrolTrafficConfigurationURLTmpl, kontrolBaseURL, tenantUuid)

	return trafficConfigurationURL, nil
}

func getClusterResourcesURL(tenantUuid api_types.Uuid) (string, error) {

	kontrolBaseURL, err := getKontrolBaseURL(true)
	if err != nil {
		return "", stacktrace.Propagate(err, "An error occurred getting the Kontrol base URL")
	}

	clusterResourcesURL := fmt.Sprintf(kontrolClusterResourcesEndpointTmpl, kontrolBaseURL, tenantUuid)

	return clusterResourcesURL, nil
}
