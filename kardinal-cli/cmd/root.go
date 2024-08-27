package cmd

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"

	"gopkg.in/yaml.v3"
	"kardinal.cli/consts"
	"kardinal.cli/multi_os_cmd_executor"

	"github.com/kurtosis-tech/stacktrace"
	"github.com/segmentio/analytics-go/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"kardinal.cli/deployment"
	"kardinal.cli/kontrol"
	"kardinal.cli/tenant"

	api "github.com/kurtosis-tech/kardinal/libs/cli-kontrol-api/api/golang/client"
	api_types "github.com/kurtosis-tech/kardinal/libs/cli-kontrol-api/api/golang/types"
	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	net "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

const (
	kontrolBaseURLTmpl                          = "%s://%s"
	kontrolClusterResourcesEndpointTmpl         = "%s/tenant/%s/cluster-resources"
	kontrolClusterResourcesManifestEndpointTmpl = "%s/tenant/%s/cluster-resources/manifest"

	kontrolTrafficConfigurationURLTmpl = "%s/%s/traffic-configuration"

	localMinikubeKontrolAPIHost = "host.minikube.internal:8080"
	localKontrolAPIHost         = "localhost:8080"
	localFrontendHost           = "localhost:5173"
	kloudKontrolHost            = "app.kardinal.dev"
	kloudKontrolAPIHost         = kloudKontrolHost + "/api"

	httpSchme   = "http"
	httpsScheme = httpSchme + "s"

	deleteAllDevFlowsFlagName = "all"

	addTraceRouterFlagName = "add-trace-router"
	yamlSeparator          = "---"
)

var (
	kubernetesManifestFile string
	devMode                bool
	serviceImagePairs      []string
	templateName           string
	templateYamlFile       string
	templateDescription    string
	templateArgsFile       string
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

var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "Manage template creation",
}

var topologyCmd = &cobra.Command{
	Use:   "topology",
	Short: "Manage Kardinal topologies",
}

var tenantCmd = &cobra.Command{
	Use:   "tenant",
	Short: "Manage tenant",
}

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy services",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		serviceConfigs, ingressConfigs, namespace, err := parseKubernetesManifestFile(kubernetesManifestFile)
		if err != nil {
			log.Fatalf("Error loading k8s manifest file: %v", err)
		}
		tenantUuid, err := tenant.GetOrCreateUserTenantUUID()
		if err != nil {
			log.Fatal("Error getting or creating user tenant UUID", err)
		}

		deploy(tenantUuid.String(), serviceConfigs, ingressConfigs, namespace)
	},
}

var templateCreateCmd = &cobra.Command{
	Use:   "create [template-name]",
	Short: "Create a new template",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		templateName := args[0]

		// TODO - add a special parser that throws an error if the template is not valid
		// A valid template only modifies the kardinal.service.dev annotations
		// A valid template only modifies services
		// A valid template has metadata.name
		// A valid template modifies at least one service
		serviceConfigs, _, _, err := parseKubernetesManifestFile(templateYamlFile)
		if err != nil {
			log.Fatalf("Error loading template file: %v", err)
		}

		var services []corev1.Service

		for _, config := range serviceConfigs {
			services = append(services, config.Service)
		}

		tenantUuid, err := tenant.GetOrCreateUserTenantUUID()
		if err != nil {
			log.Fatal("Error getting or creating user tenant UUID", err)
		}

		createTemplate(tenantUuid.String(), templateName, services, templateDescription)
	},
}

var templateDeleteCmd = &cobra.Command{
	Use:   "delete [template-name]",
	Short: "Delete a template",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		templateName := args[0]
		tenantUuid, err := tenant.GetOrCreateUserTenantUUID()
		if err != nil {
			log.Fatal("Error getting or creating user tenant UUID", err)
		}
		deleteTemplate(tenantUuid.String(), templateName)
	},
}

var templateListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all templates",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		tenantUuid, err := tenant.GetOrCreateUserTenantUUID()
		if err != nil {
			log.Fatal("Error getting or creating user tenant UUID", err)
		}
		listTemplates(tenantUuid.String())
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

		pairsMap := parsePairs(serviceImagePairs)
		pairsMap[serviceName] = imageName

		for key, value := range pairsMap {
			fmt.Printf("%s: %s\n", key, value)
		}

		tenantUuid, err := tenant.GetOrCreateUserTenantUUID()
		if err != nil {
			log.Fatal("Error getting or creating user tenant UUID", err)
		}

		logrus.Infof("Creating service %s with image %s in development mode...\n", serviceName, imageName)
		if templateName != "" {
			logrus.Infof("Using template: %s\n", templateName)
		}

		templateArgs, err := parseTemplateArgs(templateArgsFile)
		if err != nil {
			log.Fatalf("Error parsing template arguments: %v", err)
		}

		createDevFlow(tenantUuid.String(), pairsMap, templateName, templateArgs)
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete [flow-id]",
	Short: "Delete services",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var flowId string

		if len(args) > 0 {
			flowId = args[0]
		}

		shouldDeleteAllDevFlows, err := cmd.Flags().GetBool(deleteAllDevFlowsFlagName)
		if err != nil {
			log.Fatalf("Error getting %s flag: %v", deleteAllDevFlowsFlagName, err)
		}

		if !shouldDeleteAllDevFlows && flowId == "" {
			log.Fatal("Either 'flow-id' argument or 'all' flag must be set")
		}

		tenantUuid, err := tenant.GetOrCreateUserTenantUUID()
		if err != nil {
			log.Fatal("Error getting or creating user tenant UUID", err)
		}

		allFlowIds := []string{}
		if flowId != "" {
			allFlowIds = append(allFlowIds, flowId)
		}

		if shouldDeleteAllDevFlows {
			currentFlows, err := getTenantUuidFlows(tenantUuid.String())
			if err != nil {
				log.Fatalf("Failed to get the current dev flows: %v", err)
			}
			for _, currentDevFlow := range currentFlows {
				allFlowIds = append(allFlowIds, currentDevFlow.FlowId)
			}
		}

		for _, fID := range allFlowIds {
			deleteFlow(tenantUuid.String(), fID)
		}
	},
}

var topologyManifestCmd = &cobra.Command{
	Use:   "print-manifest",
	Short: "print the current cluster topology manifest deployed in Kontrol",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		addTraceRouter, err := cmd.Flags().GetBool(addTraceRouterFlagName)
		if err != nil {
			log.Fatalf("Error getting add-trace-router flag: %v", err)
		}

		tenantUuid, err := tenant.GetOrCreateUserTenantUUID()
		if err != nil {
			log.Fatal("Error getting or creating user tenant UUID", err)
		}

		ctx := context.Background()
		client := getKontrolServiceClient()

		resp, err := client.GetTenantUuidManifest(ctx, tenantUuid.String())
		if err != nil {
			log.Fatalf("Failed to get topology manifest: %v", err)
		}

		if resp == nil || resp.StatusCode != 200 {
			log.Fatalf("Not Topology manifest successfull response, response status code: %d", resp.StatusCode)
		}

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Error reading response body: %v\n", err)
			return
		}

		topologyManifest := string(bodyBytes)

		manifestToPrint := topologyManifest

		if addTraceRouter {
			traceRouterManifest, err := deployment.GetKardinalTraceRouterManifest()
			if err != nil {
				log.Fatalf("Error getting kardinal-trace router manifest: %v", err)
			}
			manifestToPrint = manifestToPrint + yamlSeparator + traceRouterManifest
		}

		fmt.Println(manifestToPrint)
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

		fmt.Print("Kardinal manager removed from cluster\n")
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
		// TODO support local-minikube deployments
		if err := multi_os_cmd_executor.OpenFile(path.Join(consts.KardinalDevURL, tenantUuidStr, consts.KardinalTrafficConfigurationSuffix)); err != nil {
			log.Fatal("Error occurred opening the Kardinal dashboard", err)
		}
	},
}

var gatewayCmd = &cobra.Command{
	Use:   "gateway [flow-id]",
	Short: "Opens a gateway to the given flow",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmr *cobra.Command, args []string) {
		flowId := args[0]

		tenantUuid, err := tenant.GetOrCreateUserTenantUUID()
		if err != nil {
			log.Fatal("Error getting or creating user tenant UUID", err)
		}

		ctx := context.Background()
		client := getKontrolServiceClient()

		resp, err := client.GetTenantUuidFlowsWithResponse(ctx, tenantUuid.String())
		if err != nil {
			log.Fatalf("Failed to list flows: %v", err)
		}

		if resp == nil || resp.JSON200 == nil {
			log.Fatalf("List flow response is empty")
		}

		var host string

		for _, flow := range *resp.JSON200 {
			if flow.FlowId == flowId {
				if len(flow.FlowUrls) > 0 {
					host = flow.FlowUrls[0]
				} else {
					log.Fatalf("Flow '%s' has no hosts", flowId)
				}
			}
		}

		if host == "" {
			log.Fatalf("Couldn't find flow with id '%s'", flowId)
		}

		if err := deployment.StartGateway(host, flowId); err != nil {
			log.Fatal("An error occurred while creating a gateway", err)
		}
	},
}

var reportInstall = &cobra.Command{
	Use:   "report-install",
	Short: "Help us improve and grow Kardinal by anonymously reporting your install",
	Args:  cobra.ExactArgs(0),
	Run: func(cmr *cobra.Command, args []string) {
		tenantUuid, err := tenant.GetOrCreateUserTenantUUID()
		if err != nil {
			log.Fatal("Error getting or creating user tenant UUID", err)
		}
		// This write key is not sensitive. It is equivalent to a public key.
		analyticsClient := analytics.New("IMYNcUACcPpcIJuS6ChHpMd4z4ZpvVFq")
		defer analyticsClient.Close()

		props := analytics.NewProperties()
		if username, exists := os.LookupEnv("KARDINAL_PLAYGROUND_USERNAME"); exists {
			props.Set("playground_username", username)
		}

		analyticsClient.Enqueue(analytics.Track{
			Event:      "install_cli",
			UserId:     tenantUuid.String(),
			Properties: props,
		})
		log.Println("Thank you for helping us improve Kardinal! 🧡")
	},
}

var tenantShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show tenant UUID",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		tenantUuid, err := tenant.GetOrCreateUserTenantUUID()
		if err != nil {
			log.Fatal("Error getting or creating user tenant UUID", err)
		}
		fmt.Printf("%s\n", tenantUuid)
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
	rootCmd.AddCommand(templateCmd)
	rootCmd.AddCommand(dashboardCmd)
	rootCmd.AddCommand(gatewayCmd)
	rootCmd.AddCommand(reportInstall)
	rootCmd.AddCommand(topologyCmd)
	rootCmd.AddCommand(tenantCmd)

	flowCmd.AddCommand(listCmd, createCmd, deleteCmd)
	managerCmd.AddCommand(deployManagerCmd, removeManagerCmd)
	templateCmd.AddCommand(templateCreateCmd, templateDeleteCmd, templateListCmd)
	topologyCmd.AddCommand(topologyManifestCmd)
	topologyManifestCmd.Flags().BoolP(addTraceRouterFlagName, "", false, "Include the trace router in the printed manifest")
	tenantCmd.AddCommand(tenantShowCmd)

	createCmd.Flags().StringSliceVarP(&serviceImagePairs, "service-image", "s", []string{}, "Extra service and respective image to include in the same flow (can be used multiple times)")
	createCmd.Flags().StringVarP(&templateName, "template", "t", "", "Template name to use for the flow creation")
	createCmd.Flags().StringVarP(&templateArgsFile, "template-args", "a", "", "Path to YAML file containing template arguments")

	deployCmd.PersistentFlags().StringVarP(&kubernetesManifestFile, "k8s-manifest", "k", "", "Path to the K8S manifest file")
	deployCmd.MarkPersistentFlagRequired("k8s-manifest")

	templateCreateCmd.Flags().StringVarP(&templateYamlFile, "template-yaml", "t", "", "Path to the YAML file containing the template")
	templateCreateCmd.Flags().StringVarP(&templateDescription, "description", "d", "", "Description of the template")
	templateCreateCmd.MarkFlagRequired("template-yaml")
	deleteCmd.Flags().BoolP(deleteAllDevFlowsFlagName, "", false, "Delete all the current dev flows")
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

func parsePairs(pairs []string) map[string]string {
	pairsMap := make(map[string]string)
	for _, pair := range pairs {
		kv := strings.SplitN(pair, "=", 2)
		if len(kv) == 2 {
			pairsMap[kv[0]] = kv[1]
		}
	}
	return pairsMap
}

func parseKubernetesManifestFile(kubernetesManifestFile string) ([]api_types.ServiceConfig, []api_types.IngressConfig, string, error) {
	fileBytes, err := loadKubernetesManifestFile(kubernetesManifestFile)
	if err != nil {
		log.Fatalf("Error loading kubernetest manifest file: %v", err)
		return nil, nil, "", err
	}

	manifest := string(fileBytes)
	var namespace string
	serviceConfigs := map[string]*api_types.ServiceConfig{}
	ingressConfigs := map[string]*api_types.IngressConfig{}
	decode := scheme.Codecs.UniversalDeserializer().Decode
	for _, spec := range strings.Split(manifest, "---") {
		if len(spec) == 0 {
			continue
		}
		obj, _, err := decode([]byte(spec), nil, nil)
		if err != nil {
			return nil, nil, "", stacktrace.Propagate(err, "An error occurred parsing the spec: %s", spec)
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
		case *net.Ingress:
			ingress := obj
			ingressName := getObjectName(ingress.GetObjectMeta().(*metav1.ObjectMeta))
			ingressConfigs[ingressName] = &api_types.IngressConfig{Ingress: *ingress}
		case *corev1.Namespace:
			namespaceObj := obj
			namespaceName := getObjectName(namespaceObj.GetObjectMeta().(*metav1.ObjectMeta))
			namespace = namespaceName
		default:
			return nil, nil, "", stacktrace.NewError("An error occurred parsing the manifest because of an unsupported kubernetes type")
		}
	}

	finalServiceConfigs := []api_types.ServiceConfig{}
	for _, serviceConfig := range serviceConfigs {
		finalServiceConfigs = append(finalServiceConfigs, *serviceConfig)
	}

	finalIngressConfigs := []api_types.IngressConfig{}
	for _, ingressConfig := range ingressConfigs {
		finalIngressConfigs = append(finalIngressConfigs, *ingressConfig)
	}

	return finalServiceConfigs, finalIngressConfigs, namespace, nil
}

func parseTemplateArgs(filename string) (map[string]interface{}, error) {
	if filename == "" {
		return nil, nil
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var args map[string]interface{}
	err = yaml.Unmarshal(data, &args)
	if err != nil {
		return nil, err
	}

	return args, nil
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
	flows, err := getTenantUuidFlows(tenantUuid)
	if err != nil {
		log.Fatalf("Failed to get dev flow: %v", err)
	}

	printFlowTable(flows)
	return
}

func getTenantUuidFlows(tenantUuid api_types.Uuid) ([]api_types.Flow, error) {
	ctx := context.Background()
	client := getKontrolServiceClient()

	resp, err := client.GetTenantUuidFlowsWithResponse(ctx, tenantUuid)
	if err != nil {
		return nil, stacktrace.Propagate(err, "Failed to get tenant UUID %s dev flows", tenantUuid)
	}

	if resp.StatusCode() == 200 {
		return *resp.JSON200, nil
	}

	return nil, stacktrace.NewError("Failed to get tenant UUID '%s' dev flows, '%d' status code received", tenantUuid, resp.StatusCode())
}

func createDevFlow(tenantUuid api_types.Uuid, pairsMap map[string]string, templateName string, templateArgs map[string]interface{}) {
	ctx := context.Background()

	devSpec := api_types.FlowSpec{}
	for serviceName, imageLocator := range pairsMap {
		devSpec = append(devSpec, struct {
			ImageLocator string `json:"image-locator"`
			ServiceName  string `json:"service-name"`
		}{
			ImageLocator: imageLocator,
			ServiceName:  serviceName,
		})
	}

	client := getKontrolServiceClient()

	var templateSpec *api_types.TemplateSpec
	if templateName != "" {
		templateSpec = &api_types.TemplateSpec{
			TemplateName: templateName,
			Arguments:    &templateArgs,
		}
	}

	resp, err := client.PostTenantUuidFlowCreateWithResponse(ctx, tenantUuid, api_types.PostTenantUuidFlowCreateJSONRequestBody{
		FlowSpec:     devSpec,
		TemplateSpec: templateSpec,
	})
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

func deploy(tenantUuid api_types.Uuid, serviceConfigs []api_types.ServiceConfig, ingressConfigs []api_types.IngressConfig, namespace string) {
	ctx := context.Background()

	body := api_types.PostTenantUuidDeployJSONRequestBody{
		ServiceConfigs: &serviceConfigs,
		IngressConfigs: &ingressConfigs,
		Namespace:      &namespace,
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
		fmt.Printf("View and manage flows:\n⚙️  %s\n", trafficConfigurationURL)
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
		fmt.Printf("Dev flow %s has been deleted\n", flowId)
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

func createTemplate(tenantUuid api_types.Uuid, templateName string, services []corev1.Service, description string) {
	ctx := context.Background()

	client := getKontrolServiceClient()

	templateConfig := api_types.TemplateConfig{
		Name:    templateName,
		Service: services,
	}

	if description != "" {
		templateConfig.Description = &description
	}

	resp, err := client.PostTenantUuidTemplatesCreateWithResponse(ctx, tenantUuid, api_types.PostTenantUuidTemplatesCreateJSONRequestBody(templateConfig))
	if err != nil {
		log.Fatalf("Failed to create template: %v", err)
	}

	if resp.StatusCode() == 200 {
		fmt.Printf("Template '%s' created successfully. Template ID: %s\n", resp.JSON200.Name, resp.JSON200.TemplateId)
		if resp.JSON200.Description != nil {
			fmt.Printf("Template Description: %s\n", *resp.JSON200.Description)
		}
		return
	}

	if resp.StatusCode() == 404 {
		fmt.Printf("Could not create template, missing %s: %s\n", resp.JSON404.ResourceType, resp.JSON404.Id)
	} else if resp.StatusCode() == 500 {
		fmt.Printf("Could not create template, error %s: %v\n", resp.JSON500.Error, resp.JSON500.Msg)
	} else {
		fmt.Printf("Failed to create template: %s\n", string(resp.Body))
	}
	log.Fatal("Template creation failed")
}

func deleteTemplate(tenantUuid api_types.Uuid, templateName string) {
	ctx := context.Background()
	client := getKontrolServiceClient()

	resp, err := client.DeleteTenantUuidTemplatesTemplateName(ctx, tenantUuid, templateName)
	if err != nil {
		log.Fatalf("Failed to delete template: %v", err)
	}

	respCode := resp.StatusCode
	if respCode >= 200 && respCode < 300 {
		fmt.Printf("Template '%s' has been deleted\n", templateName)
	} else {
		fmt.Printf("Failed to delete template!\n")
		os.Exit(1)
	}
}

func listTemplates(tenantUuid api_types.Uuid) {
	ctx := context.Background()
	client := getKontrolServiceClient()

	resp, err := client.GetTenantUuidTemplatesWithResponse(ctx, tenantUuid)
	if err != nil {
		log.Fatalf("Failed to list templates: %v", err)
	}

	if resp.StatusCode() == 200 {
		printTemplateTable(*resp.JSON200)
		return
	}

	if resp.StatusCode() == 404 {
		fmt.Printf("Could not list templates, missing %s: %s\n", resp.JSON404.ResourceType, resp.JSON404.Id)
	} else if resp.StatusCode() == 500 {
		fmt.Printf("Could not list templates, error %s: %v\n", resp.JSON500.Error, resp.JSON500.Msg)
	} else {
		fmt.Printf("Failed to list templates: %s\n", string(resp.Body))
	}
	os.Exit(1)
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

func getClusterResourcesManifestURL(tenantUuid api_types.Uuid) (string, error) {
	kontrolBaseURL, err := getKontrolBaseURLForManager()
	if err != nil {
		return "", stacktrace.Propagate(err, "An error occurred getting the Kontrol base URL")
	}

	clusterResourcesURL := fmt.Sprintf(kontrolClusterResourcesManifestEndpointTmpl, kontrolBaseURL, tenantUuid)

	return clusterResourcesURL, nil
}
