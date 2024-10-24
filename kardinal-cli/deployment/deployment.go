package deployment

import (
	"bytes"
	"context"
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"net/http"
	"text/template"

	"kardinal.cli/kubernetes"

	"kardinal.cli/kontrol"

	"github.com/kurtosis-tech/stacktrace"
	"github.com/sirupsen/logrus"
	"kardinal.cli/consts"
)

const (
	kardinalNamespace                   = "default"
	kardinalManagerDeploymentTmplName   = "kardinal-manager-deployment"
	kardinalTraceRouterManifestTmplName = "kardinal-trace-router-manifest"
	gatewayAPIInstallYamlURL            = "https://github.com/kubernetes-sigs/gateway-api/releases/download/v1.1.0/standard-install.yaml"

	kardinalManagerAuthTmpl = `
apiVersion: v1
kind: ServiceAccount
metadata:
  name: kardinal-manager
  namespace: {{.Namespace}}
  labels:
    {{.KardinalAppIDLabelKey}}: {{.KardinalManagerAppIDLabelValue}}

---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: kardinal-manager-role
  labels:
    {{.KardinalAppIDLabelKey}}: {{.KardinalManagerAppIDLabelValue}}
rules:
  - apiGroups: ["*"]
    resources: ["namespaces", "pods", "services", "deployments", "statefulsets", "virtualservices", "workloadgroups", "workloadentries", "sidecars", "serviceentries", "gateways", "envoyfilters", "destinationrules", "authorizationpolicies", "ingresses", "httproutes"]
    verbs: ["get", "list", "watch", "create", "update", "patch", "delete"]

---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: kardinal-manager-binding
  labels:
    {{.KardinalAppIDLabelKey}}: {{.KardinalManagerAppIDLabelValue}}
subjects:
  - kind: ServiceAccount
    name: kardinal-manager
    namespace: default
roleRef:
  kind: ClusterRole
  name: kardinal-manager-role
  apiGroup: rbac.authorization.k8s.io
---
`

	kardinalManagerDeploymentTmpl = `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: kardinal-manager
  namespace: {{.Namespace}}
  labels:
    {{.KardinalAppIDLabelKey}}: {{.KardinalManagerAppIDLabelValue}}
spec:
  replicas: 1
  selector:
    matchLabels:
      {{.KardinalAppIDLabelKey}}: {{.KardinalManagerAppIDLabelValue}}
  template:
    metadata:
      labels:
        {{.KardinalAppIDLabelKey}}: {{.KardinalManagerAppIDLabelValue}}
        app: kardinal-manager
    spec:
      serviceAccountName: kardinal-manager
      containers:
        - name: kardinal-manager
          image: kurtosistech/kardinal-manager:latest
          imagePullPolicy: {{.KardinalManagerContainerImagePullPolicy}}
          env:
            - name: KUBERNETES_SERVICE_HOST
              value: "kubernetes.default.svc"
            - name: KUBERNETES_SERVICE_PORT
              value: "443"
            - name: KARDINAL_MANAGER_CLUSTER_CONFIG_ENDPOINT
              value: "{{.ClusterResourcesURL}}"
            - name: KARDINAL_MANAGER_FETCHER_JOB_DURATION_SECONDS
              value: "10"
---`

	kardinalTraceRouterDeploymentTmpl = `
apiVersion: v1
kind: Service
metadata:
  name: trace-router
  labels:
    {{.KardinalAppIDLabelKey}}: {{.KardinalManagerAppIDLabelValue}}
  namespace: {{.Namespace}}
spec:
  ports:
    - port: 8080
      targetPort: 8080
  selector:
    app: trace-router
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: trace-router
  namespace: {{.Namespace}}
  labels:
    {{.KardinalAppIDLabelKey}}: {{.KardinalManagerAppIDLabelValue}}
    app: trace-router
spec:
  replicas: 1
  selector:
    matchLabels:
      app: trace-router
  template:
    metadata:
      labels:
        app: trace-router
    spec:
      serviceAccountName: kardinal-manager
      containers:
        - name: trace-router
          image: kurtosistech/kardinal-router:latest
          imagePullPolicy: {{.KardinalTraceRouterContainerImagePullPolicy}}
          ports:
            - containerPort: 8080
          env:
            - name: REDIS_HOST
              value: trace-router-redis
            - name: REDIS_PORT
              value: "6379"
---
apiVersion: v1
kind: Service
metadata:
  name: trace-router-redis
  namespace: {{.Namespace}}
  labels:
    {{.KardinalAppIDLabelKey}}: {{.KardinalManagerAppIDLabelValue}}
    app: trace-router-redis
spec:
  ports:
    - port: 6379
      targetPort: 6379
  selector:
    app: trace-router-redis
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: trace-router-redis
  labels:
    {{.KardinalAppIDLabelKey}}: {{.KardinalManagerAppIDLabelValue}}
  namespace: {{.Namespace}}
spec:
  replicas: 1
  selector:
    matchLabels:
      app: trace-router-redis
  template:
    metadata:
      labels:
        app: trace-router-redis
    spec:
      serviceAccountName: kardinal-manager
      containers:
        - name: redis
          image: bitnami/redis:6.2
          ports:
            - containerPort: 6379
          env:
            - name: ALLOW_EMPTY_PASSWORD
              value: "yes"
`

	kardinalTraceRouterManifestTmpl = kardinalManagerAuthTmpl + kardinalTraceRouterDeploymentTmpl

	allKardinalTmpls = kardinalManagerAuthTmpl + kardinalManagerDeploymentTmpl + kardinalTraceRouterDeploymentTmpl
)

type templateData struct {
	Namespace                                   string
	ClusterResourcesURL                         string
	KardinalAppIDLabelKey                       string
	KardinalManagerAppIDLabelValue              string
	KardinalManagerContainerImagePullPolicy     string
	KardinalTraceRouterContainerImagePullPolicy string
}

func DeployKardinalManagerInCluster(ctx context.Context, clusterResourcesURL string, kontrolLocation string) error {
	k8sConfig, err := kubernetes.GetConfig()
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred while creating the Kubernetes client")
	}
	kubernetesClientObj, err := kubernetes.CreateKubernetesClient(k8sConfig)
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred while creating the Kubernetes client")
	}

	kardinalManagerDeploymentTemplate, err := template.New(kardinalManagerDeploymentTmplName).Parse(allKardinalTmpls)
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred while parsing the kardinal-manager deployment template")
	}

	var imagePullPolicy string

	switch kontrolLocation {
	case kontrol.KontrolLocationLocal:
		imagePullPolicy = "IfNotPresent"
	case kontrol.KontrolLocationKloud:
		imagePullPolicy = "Always"
	default:
		stacktrace.NewError("invalid Kontrol location: %s", kontrolLocation)
	}

	templateDataObj := templateData{
		Namespace:                                   kardinalNamespace,
		ClusterResourcesURL:                         clusterResourcesURL,
		KardinalAppIDLabelKey:                       consts.KardinalAppIDLabelKey,
		KardinalManagerAppIDLabelValue:              consts.KardinalManagerAppIDLabelValue,
		KardinalManagerContainerImagePullPolicy:     imagePullPolicy,
		KardinalTraceRouterContainerImagePullPolicy: imagePullPolicy,
	}

	yamlFileContentsBuffer := &bytes.Buffer{}

	if err = kardinalManagerDeploymentTemplate.Execute(yamlFileContentsBuffer, templateDataObj); err != nil {
		return stacktrace.Propagate(err, "An error occurred while executing the template '%s' with data objects '%+v'", kardinalManagerDeploymentTmplName, templateDataObj)
	}

	if err = kubernetesClientObj.ApplyYamlFileContentInNamespace(ctx, kardinalNamespace, yamlFileContentsBuffer.Bytes()); err != nil {
		return stacktrace.Propagate(err, "An error occurred while applying the kardinal-manager deployment")
	}

	if err := installGatewayAPI(ctx, kubernetesClientObj); err != nil {
		return stacktrace.Propagate(err, "An error occurred while installing the Gateway API")
	}

	return nil
}

type k8sResourceKindAndMetadata struct {
	Kind     string `yaml:"kind"`
	Metadata struct {
		Name      string `yaml:"name"`
		Namespace string `yaml:"namespace"`
	}
}

func DeployResourceSpecs(ctx context.Context, namespace string, resourceSpecs []string) error {
	if len(resourceSpecs) == 0 {
		return nil
	}

	k8sConfig, err := kubernetes.GetConfig()
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred while creating the Kubernetes client")
	}
	kubernetesClientObj, err := kubernetes.CreateKubernetesClient(k8sConfig)
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred while creating the Kubernetes client")
	}

	if err := kubernetesClientObj.EnsureNamespace(ctx, namespace); err != nil {
		return stacktrace.Propagate(err, "An error occurred while ensuring the namespace '%s'", namespace)
	}

	for _, resourceSpec := range resourceSpecs {
		resourceSpecBytes := []byte(resourceSpec)

		var k8sResourceMetadataObj k8sResourceKindAndMetadata
		err := yaml.Unmarshal([]byte(resourceSpec), &k8sResourceMetadataObj)
		if err != nil {
			return stacktrace.Propagate(err, "An error occurred while unmarshalling a resource spec")
		}
		namespaceName := k8sResourceMetadataObj.Metadata.Namespace

		logrus.Debugf("Deploying resource '%s' kind '%s' in namespace '%s'", k8sResourceMetadataObj.Metadata.Name, k8sResourceMetadataObj.Kind, namespaceName)
		if err := kubernetesClientObj.ApplyYamlFileContentInNamespace(ctx, namespaceName, resourceSpecBytes); err != nil {
			return stacktrace.Propagate(err, "An error occurred while applying resource '%s' kind '%s' in namespace '%s'", k8sResourceMetadataObj.Metadata.Name, k8sResourceMetadataObj.Kind, namespaceName)
		}
	}

	return nil
}

func installGatewayAPI(ctx context.Context, kubernetesClientObj *kubernetes.KubernetesClient) error {
	resp, err := http.Get(gatewayAPIInstallYamlURL)
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred while downloading the gateway API YAML")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return stacktrace.NewError("Failed to download file, status code: %d", resp.StatusCode)
	}

	var buf bytes.Buffer

	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred while reading the response body")
	}

	logrus.Info("ℹ️  Installing the Gateway API (https://gateway-api.sigs.k8s.io/) in the cluster.")
	if err := kubernetesClientObj.ApplyYamlFileContentInNamespace(ctx, kardinalNamespace, buf.Bytes()); err != nil {
		return stacktrace.Propagate(err, "An error occurred while applying the gateway API YAML")
	}
	return nil
}

func RemoveKardinalManagerFromCluster(ctx context.Context) error {
	k8sConfig, err := kubernetes.GetConfig()
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred while creating the Kubernetes client")
	}
	kubernetesClientObj, err := kubernetes.CreateKubernetesClient(k8sConfig)
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred while creating the Kubernetes client")
	}

	labels := map[string]string{
		consts.KardinalAppIDLabelKey: consts.KardinalManagerAppIDLabelValue,
	}

	if err = kubernetesClientObj.RemoveNamespaceResourcesByLabels(ctx, kardinalNamespace, labels); err != nil {
		return stacktrace.Propagate(err, "An error occurred while removing the kardinal-manager from the cluster using labels '%+v'", labels)
	}
	fmt.Printf("⚠️  If you also want to unistall the Gateway API, run the following command:\nkubectl delete -f %s\n\n", gatewayAPIInstallYamlURL)

	return nil
}

func GetKardinalTraceRouterManifest() (string, error) {
	kardinalTraceRouterManifestTemplate, err := template.New(kardinalTraceRouterManifestTmplName).Parse(kardinalTraceRouterManifestTmpl)
	if err != nil {
		return "", stacktrace.Propagate(err, "An error occurred while parsing the Kardinal trace router manifest template")
	}

	imagePullPolicy := "Always"

	templateDataObj := templateData{
		Namespace:                                   kardinalNamespace,
		KardinalAppIDLabelKey:                       consts.KardinalAppIDLabelKey,
		KardinalManagerAppIDLabelValue:              consts.KardinalManagerAppIDLabelValue,
		KardinalTraceRouterContainerImagePullPolicy: imagePullPolicy,
	}

	yamlFileContentsBuffer := &bytes.Buffer{}

	if err = kardinalTraceRouterManifestTemplate.Execute(yamlFileContentsBuffer, templateDataObj); err != nil {
		return "", stacktrace.Propagate(err, "An error occurred while executing the template '%s' with data objects '%+v'", kardinalTraceRouterManifestTmplName, templateDataObj)
	}

	return yamlFileContentsBuffer.String(), nil
}
