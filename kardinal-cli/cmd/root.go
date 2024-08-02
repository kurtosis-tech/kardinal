package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"kardinal.cli/consts"
	"kardinal.cli/multi_os_cmd_executor"

	"github.com/kurtosis-tech/stacktrace"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"kardinal.cli/deployment"
	"kardinal.cli/kontrol"
	"kardinal.cli/tenant"

	api "github.com/kurtosis-tech/kardinal/libs/cli-kontrol-api/api/golang/client"
	api_types "github.com/kurtosis-tech/kardinal/libs/cli-kontrol-api/api/golang/types"
	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

const (
	projectName = "kardinal"

	kontrolBaseURLTmpl                  = "%s://%s"
	kontrolClusterResourcesEndpointTmpl = "%s/tenant/%s/cluster-resources"

	kontrolTrafficConfigurationURLTmpl = "%s/%s/traffic-configuration"

	localMinikubeKontrolAPIHost = "host.minikube.internal:8080"
	localKontrolAPIHost         = "localhost:8080"
	localFrontendHost           = "localhost:5173"
	kloudKontrolHost            = "app.kardinal.dev"
	kloudKontrolAPIHost         = kloudKontrolHost + "/api"

	httpSchme   = "http"
	httpsScheme = httpSchme + "s"
)

var (
	kubernetesManifestFile string
	devMode                bool
)

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

var listCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all active flows",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		tenantUuid, err := tenant.GetOrCreateUserTenantUUID()
		if err != nil {
			log.Fatal("Error getting or creating user tenant UUID", err)
		}

		listDevFlow(tenantUuid.String())
	},
}

var createCmd = &cobra.Command{
	Use:   "create [service name] [image name]",
	Short: "Create a new service in development mode",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		serviceName, imageName := args[0], args[1]

		tenantUuid, err := tenant.GetOrCreateUserTenantUUID()
		if err != nil {
			log.Fatal("Error getting or creating user tenant UUID", err)
		}

		logrus.Infof("Creating service %s with image %s in development mode...\n", serviceName, imageName)
		createDevFlow(tenantUuid.String(), imageName, serviceName)
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete services",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		flowId := args[0]

		tenantUuid, err := tenant.GetOrCreateUserTenantUUID()
		if err != nil {
			log.Fatal("Error getting or creating user tenant UUID", err)
		}
		deleteFlow(tenantUuid.String(), flowId)

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
			log.Fatal("Error occurred opening the Kardinal dashboard", err)
		}
	},
}

func init() {
	devMode = false
	if os.Getenv("KARDINAL_CLI_DEV_MODE") == "TRUE" {
		devMode = true
	}

	rootCmd.AddCommand(flowCmd)
	rootCmd.AddCommand(managerCmd)
	rootCmd.AddCommand(deployCmd)
	rootCmd.AddCommand(dashboardCmd)
	flowCmd.AddCommand(listCmd, createCmd, deleteCmd)
	managerCmd.AddCommand(deployManagerCmd, removeManagerCmd)

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
	serviceConfigs := map[string]*api_types.ServiceConfig{}
	decode := scheme.Codecs.UniversalDeserializer().Decode
	for _, spec := range strings.Split(manifest, "---") {
		if len(spec) == 0 {
			continue
		}
		obj, _, err := decode([]byte(spec), nil, nil)
		if err != nil {
			return nil, stacktrace.Propagate(err, "An error occurred parsing the spec: %s", spec)
		}
		switch obj := obj.(type) {
		case *corev1.Service:
			service := obj
			serviceName := getObjectName(service.GetObjectMeta().(*metav1.ObjectMeta))
			_, ok := serviceConfigs[serviceName]
			if !ok {
				serviceConfigs[serviceName] = &api_types.ServiceConfig{
					Service: *service,
				}
			} else {
				serviceConfigs[serviceName].Service = *service
			}
		case *appv1.Deployment:
			deployment := obj
			deploymentName := getObjectName(deployment.GetObjectMeta().(*metav1.ObjectMeta))
			_, ok := serviceConfigs[deploymentName]
			if !ok {
				serviceConfigs[deploymentName] = &api_types.ServiceConfig{
					Deployment: *deployment,
				}
			} else {
				serviceConfigs[deploymentName].Deployment = *deployment
			}
		default:
			return nil, stacktrace.NewError("An error occurred parsing the manifest because of an unsupported kubernetes type")
		}
	}

	finalServiceConfigs := []api_types.ServiceConfig{}
	for _, serviceConfig := range serviceConfigs {
		finalServiceConfigs = append(finalServiceConfigs, *serviceConfig)
	}

	return finalServiceConfigs, nil
}

// Use in priority the label app value
func getObjectName(obj *metav1.ObjectMeta) string {
	labelApp, ok := obj.GetLabels()["app"]
	if ok {
		return labelApp
	}

	return obj.GetName()
}

func listDevFlow(tenantUuid api_types.Uuid) {
	ctx := context.Background()
	client := getKontrolServiceClient()

	resp, err := client.GetTenantUuidFlowsWithResponse(ctx, tenantUuid)
	if err != nil {
		log.Fatalf("Failed to create dev flow: %v", err)
	}

	if resp.StatusCode() == 200 {
		printTable(*resp.JSON200)
		return
	}

	if resp.StatusCode() == 404 {
		fmt.Printf("Could not create flow, missing %s: %s\n", resp.JSON404.ResourceType, resp.JSON404.Id)
	} else if resp.StatusCode() == 500 {
		fmt.Printf("Could not create flow, error %s: %v\n", resp.JSON500.Error, resp.JSON500.Msg)
	} else {
		fmt.Printf("Failed to create dev flow: %s\n", string(resp.Body))
	}
	os.Exit(1)
}

func createDevFlow(tenantUuid api_types.Uuid, imageLocator, serviceName string) {
	ctx := context.Background()

	devSpec := api_types.FlowSpec{
		{
			ServiceName:  serviceName,
			ImageLocator: imageLocator,
		},
	}
	client := getKontrolServiceClient()

	resp, err := client.PostTenantUuidFlowCreateWithResponse(ctx, tenantUuid, devSpec)
	if err != nil {
		log.Fatalf("Failed to create dev flow: %v", err)
	}

	if resp.StatusCode() == 200 {
		fmt.Printf("Flow \"%s\" created. Access it on:\n", resp.JSON200.FlowId)
		for _, url := range resp.JSON200.FlowUrls {
			fmt.Printf("🌐 http://%s\n", url)
		}
		return
	}

	if resp.StatusCode() == 404 {
		fmt.Printf("Could not create flow, missing %s: %s\n", resp.JSON404.ResourceType, resp.JSON404.Id)
	} else if resp.StatusCode() == 500 {
		fmt.Printf("Could not create flow, error %s: %v\n", resp.JSON500.Error, resp.JSON500.Msg)
	} else {
		fmt.Printf("Failed to create dev flow: %s\n", string(resp.Body))
	}
	os.Exit(1)
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

	trafficConfigurationURL, err := getTrafficConfigurationURL(tenantUuid)
	if err != nil {
		logrus.Warningf("The command run successfully but it was impossible to print the traffic configuration URL because and error ocurred, please make sure to run the 'kardinal manager deploy' command first")
		return
	}

	if resp.StatusCode() == 200 {
		fmt.Printf("Flow \"%s\" created. Access it on:\n", resp.JSON200.FlowId)
		for _, url := range resp.JSON200.FlowUrls {
			fmt.Printf("🌐 http://%s\n", url)
		}
		fmt.Printf("View and manage flows:\n⚙️  %s", trafficConfigurationURL)
		return
	}

	if resp.StatusCode() == 404 {
		fmt.Printf("Could not create flow, missing %s: %s\n", resp.JSON404.ResourceType, resp.JSON404.Id)
	} else if resp.StatusCode() == 500 {
		fmt.Printf("Could not create flow, error %s: %v\n", resp.JSON500.Error, resp.JSON500.Msg)
	} else {
		fmt.Printf("Failed to create dev flow: %s\n", string(resp.Body))
	}
	os.Exit(1)
}

func deleteFlow(tenantUuid api_types.Uuid, flowId api_types.FlowId) {
	ctx := context.Background()

	client := getKontrolServiceClient()

	resp, err := client.DeleteTenantUuidFlowFlowId(ctx, tenantUuid, flowId)
	if err != nil {
		log.Fatalf("Failed to delete flow: %v", err)
	}

	respCode := resp.StatusCode
	if respCode == 200 || respCode == 204 {
		fmt.Printf("Dev flow %s has been deleted", flowId)
	} else {
		fmt.Printf("Failed to delete dev flow!\n")
		os.Exit(1)
	}
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
	kontrolHostApi, err := getKontrolBaseURLForCLI()
	if err != nil {
		logrus.Fatalf("An error occurred getting the Kontrol location:\n%v", err)
		os.Exit(1)
	}
	client, err := api.NewClientWithResponses(kontrolHostApi)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		os.Exit(1)
	}
	return client
}

func getKontrolBaseURLForUI() (string, error) {
	var (
		scheme string
		host   string
	)

	if devMode {
		scheme = httpSchme
		host = localFrontendHost
	} else {
		scheme = httpsScheme
		host = kloudKontrolHost
	}

	baseURL := fmt.Sprintf(kontrolBaseURLTmpl, scheme, host)

	return baseURL, nil
}

func getKontrolBaseURLForCLI() (string, error) {
	var (
		scheme string
		host   string
	)

	if devMode {
		scheme = httpSchme
		host = localKontrolAPIHost
	} else {
		scheme = httpsScheme
		host = kloudKontrolAPIHost
	}

	baseURL := fmt.Sprintf(kontrolBaseURLTmpl, scheme, host)

	return baseURL, nil
}

func getKontrolBaseURLForManager() (string, error) {
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
		host = kloudKontrolAPIHost
	default:
		return "", stacktrace.NewError("invalid Kontrol location: %s", kontrolLocation)
	}

	baseURL := fmt.Sprintf(kontrolBaseURLTmpl, scheme, host)
	return baseURL, nil
}

func getTrafficConfigurationURL(tenantUuid api_types.Uuid) (string, error) {
	kontrolBaseURL, err := getKontrolBaseURLForUI()
	if err != nil {
		return "", stacktrace.Propagate(err, "An error occurred getting the Kontrol base URL")
	}

	trafficConfigurationURL := fmt.Sprintf(kontrolTrafficConfigurationURLTmpl, kontrolBaseURL, tenantUuid)

	return trafficConfigurationURL, nil
}

func getClusterResourcesURL(tenantUuid api_types.Uuid) (string, error) {
	kontrolBaseURL, err := getKontrolBaseURLForManager()
	if err != nil {
		return "", stacktrace.Propagate(err, "An error occurred getting the Kontrol base URL")
	}

	clusterResourcesURL := fmt.Sprintf(kontrolClusterResourcesEndpointTmpl, kontrolBaseURL, tenantUuid)

	return clusterResourcesURL, nil
}

func printTable(flows []api_types.Flow) {
	// Find the maximum width of each column
	data := lo.Map(flows, func(flow api_types.Flow, _ int) []string {
		return []string{
			flow.FlowId,
			strings.Join(lo.Map(flow.FlowUrls, func(item string, _ int) string { return fmt.Sprintf("http://%s", item) }), ", "),
		}
	})
	header := [][]string{{"Flow ID", "Flow URL"}}
	data = append(header, data...)

	colWidths := make([]int, len(data[0]))
	for _, row := range data {
		for i, cell := range row {
			if len(cell) > colWidths[i] {
				colWidths[i] = len(cell)
			}
		}
	}

	for _, width := range colWidths {
		fmt.Print("|", strings.Repeat("-", width+2))
	}
	fmt.Println("|")

	// Print the table
	for rowNum, row := range data {
		for i, cell := range row {
			fmt.Printf("| %-*s ", colWidths[i], cell)
		}
		fmt.Println("|")

		// Print separator after header
		if rowNum == 0 {
			for _, width := range colWidths {
				fmt.Print("|", strings.Repeat("-", width+2))
			}
			fmt.Println("|")
		}
	}

	for _, width := range colWidths {
		fmt.Print("|", strings.Repeat("-", width+2))
	}
	fmt.Println("|")
}
