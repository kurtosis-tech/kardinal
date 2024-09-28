package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path"
	"slices"
	"strings"
	"time"

	"kardinal.cli/kubernetes"

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
	k8snet "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	gateway "sigs.k8s.io/gateway-api/apis/v1"
	gatewayscheme "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned/scheme"
)

const (
	kontrolBaseURLTmpl                  = "%s://%s"
	kontrolClusterResourcesEndpointTmpl = "%s/tenant/%s/cluster-resources"

	kontrolTrafficConfigurationURLTmpl = "%s/%s/traffic-configuration"

	localMinikubeKontrolAPIHost = "host.minikube.internal:8080"
	localhost                   = "localhost"
	localKontrolAPIHost         = localhost + ":8080"
	localFrontendHost           = localhost + ":5173"
	kloudKontrolHost            = "app.kardinal.dev"
	kloudKontrolAPIHost         = kloudKontrolHost + "/api"

	tcpProtocol = "tcp"
	httpSchme   = "http"
	httpsScheme = httpSchme + "s"

	deleteAllDevFlowsFlagName = "all"

	addTraceRouterFlagName = "add-trace-router"
	yamlSeparator          = "---"

	telepresenceCmdName     = "telepresence"
	telepresenceInstallDocs = "https://www.telepresence.io/docs/latest/quick-start/"
	telepresenceAppLabel    = "traffic-manager"
	ambassadorNamespace     = "ambassador"

	appLabelKey      = "app"
	versionLabelKey  = "version"
	portCheckTimeout = 5 * time.Second
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
		serviceConfigs, ingressConfigs, gatewayConfigs, routeConfigs, namespace, err := parseKubernetesManifestFile(kubernetesManifestFile)
		if err != nil {
			log.Fatalf("Error loading k8s manifest file: %v", err)
		}
		tenantUuid, err := tenant.GetOrCreateUserTenantUUID()
		if err != nil {
			log.Fatal("Error getting or creating user tenant UUID", err)
		}

		deploy(tenantUuid.String(), serviceConfigs, ingressConfigs, gatewayConfigs, routeConfigs, namespace)
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
		serviceConfigs, _, _, _, _, err := parseKubernetesManifestFile(templateYamlFile)
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

var telepresenceInterceptCmd = &cobra.Command{
	Use:   "telepresence-intercept [flow-id] [service name] [local port]",
	Short: "Execute a Telepresence intercept for a service in a dev flow",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		flowId, serviceName, localPort := args[0], args[1], args[2]

		// GUARDRAILS
		// Is telepresence CLI in host
		_, err := exec.LookPath(telepresenceCmdName)
		if err != nil {
			log.Fatalf("The %s command was not foun in the host, you can install it following this docs: %s", telepresenceCmdName, telepresenceInstallDocs)
		}

		// Is the port open and HTTP
		if err := isPortOpenAndHTTP(localPort); err != nil {
			log.Fatalf("An error occurred checking HTTP server on port '%s': %s", localPort, err)
		}

		// is Traffic-manager installed
		k8sConfig, err := kubernetes.GetConfig()
		if err != nil {
			log.Fatalf("An error occurred while creating the Kubernetes client: %s", err)
			return
		}
		kubernetesClt, err := kubernetes.CreateKubernetesClient(k8sConfig)
		if err != nil {
			log.Fatalf("An error occurred while creating the Kubernetes client: %s", err)
			return
		}

		trafficManagerDeploymentLabels := map[string]string{
			appLabelKey: telepresenceAppLabel,
		}

		ambassadorNamespaceName := ambassadorNamespace
		trafficManagerDeployments, err := kubernetesClt.GetDeploymentsByLabels(ctx, ambassadorNamespaceName, trafficManagerDeploymentLabels)
		if err != nil {
			log.Fatalf("An error occurred getting deployments with labels '%+v' in namespace '%s': %s", trafficManagerDeploymentLabels, ambassadorNamespaceName, err)
		}
		if len(trafficManagerDeployments.Items) == 0 {
			log.Fatalf("The 'traffic-manager' deployment was not foun in '%s' namespace", ambassadorNamespaceName)
		}

		tenantUuid, err := tenant.GetOrCreateUserTenantUUID()
		if err != nil {
			log.Fatal("Error getting or creating user tenant UUID", err)
		}

		var namespaceName string
		currentFlows, err := getTenantUuidFlows(tenantUuid.String())
		if err != nil {
			log.Fatalf("Failed to get the current dev flows: %v", err)
		}
		for _, currentFlow := range currentFlows {
			if *currentFlow.IsBaseline {
				namespaceName = currentFlow.FlowId
			}
		}

		serviceObj, err := kubernetesClt.GetService(ctx, namespaceName, serviceName)
		if err != nil {
			log.Fatalf("An error occurred getting service '%s': %s", serviceName, err)
		}

		var appLabel string
		// getting the app label
		for labelKey, labelValue := range serviceObj.GetLabels() {
			if labelKey == appLabelKey {
				appLabel = labelValue
			}
		}
		if appLabel == "" {
			log.Fatalf("Won't be possible to create the intercept because service '%s' doesn't have the '%s' label which is necessary to get the deployment linked to this", serviceName, appLabelKey)
		}

		deploymentLabels := map[string]string{
			appLabelKey:     appLabel,
			versionLabelKey: flowId,
		}

		deployments, err := kubernetesClt.GetDeploymentsByLabels(ctx, namespaceName, deploymentLabels)
		if err != nil {
			log.Fatalf("An error occurred getting deployments with labels '%+v' in namespace '%s': %s", deploymentLabels, namespaceName, err)
		}
		if len(deployments.Items) > 1 {
			log.Fatalf("Found more than one deployment with labels '%+v' in namespace '%s'", deploymentLabels, namespaceName)
		}

		interceptBaseName := deployments.Items[0].GetName()

		logrus.Info("Executing Telepresence intercept...")

		telepresenceConnectCmdArgs := []string{"connect", "-n", namespaceName}
		logrus.Infof("Executing Telepresence connect to namespace '%s'...", namespaceName)
		telepresenceConnectCmd := exec.Command(telepresenceCmdName, telepresenceConnectCmdArgs...)
		telepresenceConnectOutput, err := telepresenceConnectCmd.CombinedOutput()
		logrus.Infof("Telepresence connect command output: %s", string(telepresenceConnectOutput))
		if err != nil {
			log.Fatalf("An error occurred connecting Telepresence: %v", err)
		}
		logrus.Infof("Telepresence has been successfully connected to namespace '%s'...", namespaceName)

		logrus.Infof("Executing Telepresence intercept to flow id '%s'...", flowId)

		telepresenceInterceptCmdArgs := []string{"intercept", "--port", fmt.Sprintf("%s:http", localPort), "--service", interceptBaseName, interceptBaseName}
		telepresenceInterceptCmd := exec.Command(telepresenceCmdName, telepresenceInterceptCmdArgs...)
		telepresenceInterceptOutput, err := telepresenceInterceptCmd.CombinedOutput()
		logrus.Infof("Telepresence intercept command output: %s", string(telepresenceInterceptOutput))
		if err != nil {
			log.Fatalf("An error occurred running Telepresence intercept: %v", err)
		}
		logrus.Infof("Telepresence intercept successfully created in flow ID '%s' and service '%s'", flowId, serviceName)
	},
}

func isPortOpenAndHTTP(localPortStr string) error {
	// Check if the port is open
	localServiceAddress := fmt.Sprintf("%s:%s", localhost, localPortStr)
	conn, err := net.DialTimeout(tcpProtocol, localServiceAddress, portCheckTimeout)
	if err != nil {
		return stacktrace.Propagate(err, "Port %s is not open or not reachable", localPortStr)
	}
	defer conn.Close()

	logrus.Debugf("Port %s is open", localPortStr)

	client := &http.Client{
		Timeout: portCheckTimeout,
	}

	// Check if there is an HTTP server running on the port
	httpServerAddr := fmt.Sprintf("%s://%s", httpSchme, localServiceAddress)
	resp, err := client.Get(httpServerAddr)
	if err != nil {
		return stacktrace.Propagate(err, "failing to call an HTTP server on '%s'", httpServerAddr)
	}
	defer resp.Body.Close()

	return nil
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
	Use:       fmt.Sprintf("deploy [kardinal-kontrol service location] accepted values: %s and %s ", kontrol.KontrolLocationLocal, kontrol.KontrolLocationKloud),
	Short:     "Deploy Kardinal manager into the cluster",
	Long:      "The Kardinal Manager retrieves the latest configuration from the Kardinal Kontrol service and applies changes to the user K8S topology.  The Kardinal Kontrol service can be the one running in the Kardinal Cloud or can be the one deployed locally.",
	ValidArgs: []string{kontrol.KontrolLocationLocal, kontrol.KontrolLocationKloud},
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
	Use:   "gateway [flow-id] ...",
	Short: "Opens a gateway to the given list of flows",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1)),
	Run: func(cmr *cobra.Command, args []string) {
		flowIds := args

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

		hostFlowIdMap := make([]api_types.IngressAccessEntry, 0)

		for _, flow := range *resp.JSON200 {
			if slices.Contains(flowIds, flow.FlowId) {
				flowId := flow.FlowId
				if len(flow.AccessEntry) > 0 {
					hostFlowIdMap = append(hostFlowIdMap, flow.AccessEntry...)
				} else {
					log.Fatalf("Flow '%s' has no hosts", flowId)
				}
			}
		}

		if err := deployment.StartGateway(hostFlowIdMap); err != nil {
			log.Fatalf("An error occurred while creating a gateway: %v", err)
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

	flowCmd.AddCommand(listCmd, createCmd, deleteCmd, telepresenceInterceptCmd)
	managerCmd.AddCommand(deployManagerCmd, removeManagerCmd)
	templateCmd.AddCommand(templateCreateCmd, templateDeleteCmd, templateListCmd)
	topologyCmd.AddCommand(topologyManifestCmd)
	topologyManifestCmd.Flags().BoolP(addTraceRouterFlagName, "", false, "Include the trace router in the printed manifest")
	tenantCmd.AddCommand(tenantShowCmd)

	createCmd.Flags().StringSliceVarP(&serviceImagePairs, "service-image", "s", []string{}, "Extra service and respective image to include in the same flow (can be used multiple times)")
	createCmd.Flags().StringVarP(&templateName, "template", "t", "", "Template name to use for the flow creation")
	createCmd.Flags().StringVarP(&templateArgsFile, "template-args", "a", "", "JSON with the template arguments or path to YAML file containing template arguments")

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

func parseKubernetesManifestFile(kubernetesManifestFile string) ([]api_types.ServiceConfig, []api_types.IngressConfig, []api_types.GatewayConfig, []api_types.RouteConfig, string, error) {
	fileBytes, err := loadKubernetesManifestFile(kubernetesManifestFile)
	if err != nil {
		log.Fatalf("Error loading kubernetest manifest file: %v", err)
		return nil, nil, nil, nil, "", err
	}

	manifest := string(fileBytes)
	var namespace string
	serviceConfigs := map[string]*api_types.ServiceConfig{}
	ingressConfigs := map[string]*api_types.IngressConfig{}
	gatewayConfigs := map[string]*api_types.GatewayConfig{}
	routeConfigs := map[string]*api_types.RouteConfig{}

	// Register the gateway scheme to parse the Gateway CRD
	gatewayscheme.AddToScheme(scheme.Scheme)
	decode := scheme.Codecs.UniversalDeserializer().Decode
	for _, spec := range strings.Split(manifest, "---") {
		if len(spec) == 0 {
			continue
		}
		obj, _, err := decode([]byte(spec), nil, nil)
		if err != nil {
			return nil, nil, nil, nil, "", stacktrace.Propagate(err, "An error occurred parsing the spec: %s", spec)
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
		case *k8snet.Ingress:
			ingress := obj
			ingressName := getObjectName(ingress.GetObjectMeta().(*metav1.ObjectMeta))
			ingressConfigs[ingressName] = &api_types.IngressConfig{Ingress: *ingress}
		case *corev1.Namespace:
			namespaceObj := obj
			namespaceName := getObjectName(namespaceObj.GetObjectMeta().(*metav1.ObjectMeta))
			namespace = namespaceName
		case *gateway.Gateway:
			gatewayObj := obj
			gatewayName := getObjectName(gatewayObj.GetObjectMeta().(*metav1.ObjectMeta))
			gatewayConfigs[gatewayName] = &api_types.GatewayConfig{Gateway: *gatewayObj}
		case *gateway.HTTPRoute:
			routeObj := obj
			routeName := getObjectName(routeObj.GetObjectMeta().(*metav1.ObjectMeta))
			routeConfigs[routeName] = &api_types.RouteConfig{HttpRoute: *routeObj}
		default:
			return nil, nil, nil, nil, "", stacktrace.NewError("An error occurred parsing the manifest because of an unsupported kubernetes type")
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

	finalGatewayConfigs := []api_types.GatewayConfig{}
	for _, gatewayConfig := range gatewayConfigs {
		finalGatewayConfigs = append(finalGatewayConfigs, *gatewayConfig)
	}

	finalRouteConfigs := []api_types.RouteConfig{}
	for _, routeConfig := range routeConfigs {
		finalRouteConfigs = append(finalRouteConfigs, *routeConfig)
	}

	return finalServiceConfigs, finalIngressConfigs, finalGatewayConfigs, finalRouteConfigs, namespace, nil
}

func parseTemplateArgs(filepathOrJson string) (map[string]interface{}, error) {
	if filepathOrJson == "" {
		return nil, nil
	}

	var args map[string]interface{}

	if _, err := os.Stat(filepathOrJson); errors.Is(err, os.ErrNotExist) {
		err = json.Unmarshal([]byte(filepathOrJson), &args)
		if err != nil {
			return nil, stacktrace.Propagate(err, "Template arguments must be a valid JSON object or a path to a JSON file")
		}

		return args, nil
	} else {
		data, err := os.ReadFile(filepathOrJson)
		if err != nil {
			return nil, err
		}

		err = yaml.Unmarshal(data, &args)
		if err != nil {
			return nil, err
		}

		return args, nil

	}
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
		for _, entry := range resp.JSON200.AccessEntry {
			fmt.Printf("🌐 http://%s\n", entry.Hostname)
		}
		return
	}

	if resp.StatusCode() == 404 {
		fmt.Printf("Could not create flow, missing %s: %s\n", resp.JSON404.ResourceType, resp.JSON404.Id)
	} else if resp.StatusCode() == 500 {
		fmt.Printf("Could not create flow, error %s: %s\n", resp.JSON500.Error, *resp.JSON500.Msg)
	} else {
		fmt.Printf("Failed to create dev flow: %s\n", string(resp.Body))
	}
	os.Exit(1)
}

func deploy(
	tenantUuid api_types.Uuid,
	serviceConfigs []api_types.ServiceConfig,
	ingressConfigs []api_types.IngressConfig,
	gatewayConfigs []api_types.GatewayConfig,
	routeConfigs []api_types.RouteConfig,
	namespace string,
) {
	ctx := context.Background()

	body := api_types.PostTenantUuidDeployJSONRequestBody{
		ServiceConfigs: &serviceConfigs,
		IngressConfigs: &ingressConfigs,
		GatewayConfigs: &gatewayConfigs,
		RouteConfigs:   &routeConfigs,
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
		for _, entry := range resp.JSON200.AccessEntry {
			fmt.Printf("🌐 http://%s\n", entry.Hostname)
		}
		fmt.Printf("View and manage flows:\n⚙️  %s\n", trafficConfigurationURL)
		return
	}

	if resp.StatusCode() == 404 {
		fmt.Printf("Could not create flow, missing %s: %s\n", resp.JSON404.ResourceType, resp.JSON404.Id)
	} else if resp.StatusCode() == 500 {
		fmt.Printf("Could not create flow, error %s: %s\n", resp.JSON500.Error, *resp.JSON500.Msg)
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
		fmt.Printf("Could not create template, error %s: %s\n", resp.JSON500.Error, *resp.JSON500.Msg)
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
		fmt.Printf("Could not list templates, error %s: %s\n", resp.JSON500.Error, *resp.JSON500.Msg)
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
	case kontrol.KontrolLocationLocal:
		scheme = httpSchme
		host = localMinikubeKontrolAPIHost
	case kontrol.KontrolLocationKloud:
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
