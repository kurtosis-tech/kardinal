package cmd

import (
	"context"
	"fmt"
	"kardinal.cli/consts"
	"kardinal.cli/multi_os_cmd_executor"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/kurtosis-tech/stacktrace"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"kardinal.cli/deployment"
	"kardinal.cli/kontrol"
	"kardinal.cli/tenant"

	api "github.com/kurtosis-tech/kardinal/libs/cli-kontrol-api/api/golang/client"
	api_types "github.com/kurtosis-tech/kardinal/libs/cli-kontrol-api/api/golang/types"
	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

const (
	projectName = "kardinal"
	devMode     = false

	kontrolBaseURLTmpl                  = "%s://%s"
	kontrolClusterResourcesEndpointTmpl = "%s/tenant/%s/cluster-resources"

	kontrolTrafficConfigurationURLTmpl = "%s/%s/traffic-configuration"

	localMinikubeKontrolAPIHost = "host.minikube.internal:8080"
	kloudKontrolHost            = "app.kardinal.dev"
	kloudKontrolAPIHost         = kloudKontrolHost + "/api"

	httpSchme   = "http"
	httpsScheme = httpSchme + "s"
)

var kubernetesManifestFile string

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
		serviceConfigs, err := parseKubernetesManifestFile(kubernetesManifestFile)
		if err != nil {
			log.Fatalf("Error loading k8s manifest file: %v", err)
		}
		tenantUuid, err := tenant.GetOrCreateUserTenantUUID()
		if err != nil {
			log.Fatal("Error getting or creating user tenant UUID", err)
		}

		deploy(tenantUuid.String(), serviceConfigs)
	},
}

var createCmd = &cobra.Command{
	Use:   "create [service name] [image name]",
	Short: "Create a new service in development mode",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		serviceName, imageName := args[0], args[1]
		serviceConfigs, err := parseKubernetesManifestFile(kubernetesManifestFile)
		if err != nil {
			log.Fatalf("Error loading k8s manifest file: %v", err)
		}

		tenantUuid, err := tenant.GetOrCreateUserTenantUUID()
		if err != nil {
			log.Fatal("Error getting or creating user tenant UUID", err)
		}

		fmt.Printf("Creating service %s with image %s in development mode...\n", serviceName, imageName)
		createDevFlow(tenantUuid.String(), serviceConfigs, imageName, serviceName)
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete services",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		serviceConfigs, err := parseKubernetesManifestFile(kubernetesManifestFile)
		if err != nil {
			log.Fatalf("Error loading k8s manifest file: %v", err)
		}

		tenantUuid, err := tenant.GetOrCreateUserTenantUUID()
		if err != nil {
			log.Fatal("Error getting or creating user tenant UUID", err)
		}
		deleteFlow(tenantUuid.String(), serviceConfigs)

		fmt.Print("Deleting dev flow")
	},
}

var deployManagerCmd = &cobra.Command{
	Use:       fmt.Sprintf("deploy [kontrol location] accepted values: %s and %s ", kontrol.KontrolLocationLocalMinikube, kontrol.KontrolLocationKloudKontrol),
	Short:     "Deploy Kardinal manager into the cluster",
	ValidArgs: []string{kontrol.KontrolLocationLocalMinikube, kontrol.KontrolLocationKloudKontrol},
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {

		kontrolLocation := args[0]

		if err := kontrol.SaveKontrolLocation(kontrolLocation); err != nil {
			log.Fatal("Error saving the Kontrol location", err)
		}

		tenantUuid, err := tenant.GetOrCreateUserTenantUUID()
		if err != nil {
			log.Fatal("Error getting or creating user tenant UUID", err)
		}

		if err := deployManager(tenantUuid.String(), kontrolLocation); err != nil {
			log.Fatal("Error deploying Kardinal manager", err)
		}

		logrus.Infof("Kardinal manager deployed using '%s' Kontrol", kontrolLocation)
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

var dashboardCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "Open your Kardinal Dashboard",
	Args:  cobra.ExactArgs(0),
	Run: func(cmr *cobra.Command, args []string) {
		tenantUuid, err := tenant.GetOrCreateUserTenantUUID()
		if err != nil {
			log.Fatal("Error getting or creating user tenant UUID", err)
		}
		tenantUuidStr := tenantUuid.String()
		if err := multi_os_cmd_executor.OpenFile(path.Join(consts.KardinalDevURL, tenantUuidStr)); err != nil {
			log.Fatal("Error occurred opening the dashboard", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(flowCmd)
	rootCmd.AddCommand(managerCmd)
	rootCmd.AddCommand(deployCmd)
	rootCmd.AddCommand(dashboardCmd)
	flowCmd.AddCommand(createCmd, deleteCmd)
	managerCmd.AddCommand(deployManagerCmd, removeManagerCmd)

	flowCmd.PersistentFlags().StringVarP(&kubernetesManifestFile, "k8s-manifest", "k", "", "Path to the K8S manifest file")
	flowCmd.MarkPersistentFlagRequired("k8s-manifest")
	deployCmd.PersistentFlags().StringVarP(&kubernetesManifestFile, "k8s-manifest", "k", "", "Path to the K8S manifest file")
	deployCmd.MarkPersistentFlagRequired("k8s-manifest")
}

func Execute() error {
	return rootCmd.Execute()
}

func loadKubernetesManifestFile(filename string) ([]byte, error) {

	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return fileBytes, stacktrace.Propagate(err, "attempted to read kubernetes manifest file with path '%s' but failed", filename)
	}

	return fileBytes, nil
}

func parseKubernetesManifestFile(kubernetesManifestFile string) ([]api_types.ServiceConfig, error) {
	fileBytes, err := loadKubernetesManifestFile(kubernetesManifestFile)
	if err != nil {
		log.Fatalf("Error loading kubernetest manifest file: %v", err)
		return nil, err
	}

	manifest := string(fileBytes)
	// TODO: Check format of manifest file
	blocks := strings.Split(manifest, "---")
	if len(blocks)%2 != 0 {
		return nil, stacktrace.NewError("The manifest should contain pairs of service / deployment specifications")
	}
	serviceConfigs := make([]api_types.ServiceConfig, len(blocks)/2)
	decode := scheme.Codecs.UniversalDeserializer().Decode
	for index, spec := range strings.Split(manifest, "---") {
		if len(spec) == 0 {
			continue
		}
		obj, _, err := decode([]byte(spec), nil, nil)
		if err != nil {
			continue
		}
		switch obj := obj.(type) {
		case *corev1.Service:
			service := obj
			serviceConfigs[index/2].Service = *service
		case *appv1.Deployment:
			deployment := obj
			serviceConfigs[index/2].Deployment = *deployment
		default:
			return nil, stacktrace.NewError("An error occurred parsing the manifest because of an unsupported kubernetes type")
		}
	}

	return serviceConfigs, nil
}

func createDevFlow(tenantUuid api_types.Uuid, serviceConfigs []api_types.ServiceConfig, imageLocator, serviceName string) {
	ctx := context.Background()

	body := api_types.PostTenantUuidFlowCreateJSONRequestBody{
		ServiceConfigs: &serviceConfigs,
		ServiceName:    &serviceName,
		ImageLocator:   &imageLocator,
	}
	client := getKontrolServiceClient()

	resp, err := client.PostTenantUuidFlowCreateWithResponse(ctx, tenantUuid, body)
	if err != nil {
		log.Fatalf("Failed to create dev flow: %v", err)
	}

	fmt.Printf("Response: %s\n", string(resp.Body))
}

func deploy(tenantUuid api_types.Uuid, serviceConfigs []api_types.ServiceConfig) {
	ctx := context.Background()

	body := api_types.PostTenantUuidDeployJSONRequestBody{
		ServiceConfigs: &serviceConfigs,
	}
	client := getKontrolServiceClient()

	resp, err := client.PostTenantUuidDeployWithResponse(ctx, tenantUuid, body)
	if err != nil {
		log.Fatalf("Failed to deploy: %v", err)
	}

	fmt.Printf("Response: %s\n", string(resp.Body))

	trafficConfigurationURL, err := getTrafficConfigurationURL(tenantUuid)
	if err != nil {
		logrus.Warningf("The command run successfully but it was impossible to print the traffic configuration URL because and error ocurred, please make sure to run the 'kardinal manager deploy' command first")
		return
	}

	logrus.Infof("Visit: %s", trafficConfigurationURL)
}

func deleteFlow(tenantUuid api_types.Uuid, serviceConfigs []api_types.ServiceConfig) {
	ctx := context.Background()

	body := api_types.PostTenantUuidFlowDeleteJSONRequestBody{
		ServiceConfigs: &serviceConfigs,
	}
	client := getKontrolServiceClient()

	resp, err := client.PostTenantUuidFlowDeleteWithResponse(ctx, tenantUuid, body)
	if err != nil {
		log.Fatalf("Failed to delete flow: %v", err)
	}

	fmt.Printf("Response: %s\n", string(resp.Body))
}

func deployManager(tenantUuid api_types.Uuid, kontrolLocation string) error {

	ctx := context.Background()

	clusterResourcesURL, err := getClusterResourcesURL(tenantUuid)
	if err != nil {
		return stacktrace.Propagate(err, "Error getting cluster resources URL")
	}

	if err := deployment.DeployKardinalManagerInCluster(ctx, clusterResourcesURL, kontrolLocation); err != nil {
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
