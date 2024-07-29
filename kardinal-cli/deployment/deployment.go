package deployment

import (
	"bytes"
	"context"
	"kardinal.cli/kontrol"
	"text/template"

	"github.com/kurtosis-tech/stacktrace"
	"kardinal.cli/consts"
)

const (
	kardinalNamespace                 = "default"
	kardinalManagerDeploymentTmplName = "kardinal-manager-deployment"

	kardinalManagerDeploymentTmpl = `
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
    resources: ["namespaces", "pods", "services", "deployments", "virtualservices", "workloadgroups", "workloadentries", "sidecars", "serviceentries", "gateways", "envoyfilters", "destinationrules"]
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
---
apiVersion: v1
kind: Service
metadata:
  name: trace-router
  labels:
    {{.KardinalAppIDLabelKey}}: {{.KardinalManagerAppIDLabelValue}}
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
      containers:
        - name: trace-router
          image: kurtosistech/kardinal-router:latest
          imagePullPolicy: {{.KardinalManagerContainerImagePullPolicy}}
          ports:
            - containerPort: 8080
          env:
            - name: REDIS_HOST
              value: trace-router-redis
            - name: REDIS_PORT
              value: 6379
---
apiVersion: v1
kind: Service
metadata:
  name: trace-router-redis
  namespace: {{.Namespace}}
  labels:
    {{.KardinalAppIDLabelKey}}: {{.KardinalManagerAppIDLabelValue}}
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
      containers:
        - name: redis
          image: bitnami/redis:6.0.8
          ports:
            - containerPort: 6379
          env:
            - name: ALLOW_EMPTY_PASSWORD
              value: "yes"
`
)

type templateData struct {
	Namespace                               string
	ClusterResourcesURL                     string
	KardinalAppIDLabelKey                   string
	KardinalManagerAppIDLabelValue          string
	KardinalManagerContainerImagePullPolicy string
}

func DeployKardinalManagerInCluster(ctx context.Context, clusterResourcesURL string, kontrolLocation string) error {
	kubernetesClientObj, err := createKubernetesClient()
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred while creating the Kubernetes client")
	}

	kardinalManagerDeploymentTemplate, err := template.New(kardinalManagerDeploymentTmplName).Parse(kardinalManagerDeploymentTmpl)
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred while parsing the kardinal-manager deployment template")
	}

	var imagePullPolicy string

	switch kontrolLocation {
	case kontrol.KontrolLocationLocalMinikube:
		imagePullPolicy = "Never"
	case kontrol.KontrolLocationKloudKontrol:
		imagePullPolicy = "Always"
	default:
		stacktrace.NewError("invalid Kontrol location: %s", kontrolLocation)
	}

	templateDataObj := templateData{
		Namespace:                               kardinalNamespace,
		ClusterResourcesURL:                     clusterResourcesURL,
		KardinalAppIDLabelKey:                   consts.KardinalAppIDLabelKey,
		KardinalManagerAppIDLabelValue:          consts.KardinalManagerAppIDLabelValue,
		KardinalManagerContainerImagePullPolicy: imagePullPolicy,
	}

	yamlFileContentsBuffer := &bytes.Buffer{}

	if err = kardinalManagerDeploymentTemplate.Execute(yamlFileContentsBuffer, templateDataObj); err != nil {
		return stacktrace.Propagate(err, "An error occurred while executing the template '%s' with data objects '%+v'", kardinalManagerDeploymentTmplName, templateDataObj)
	}

	if err = kubernetesClientObj.ApplyYamlFileContentInNamespace(ctx, kardinalNamespace, yamlFileContentsBuffer.Bytes()); err != nil {
		return stacktrace.Propagate(err, "An error occurred while applying the kardinal-manager deployment")
	}

	return nil
}

func RemoveKardinalManagerFromCluster(ctx context.Context) error {
	kubernetesClientObj, err := createKubernetesClient()
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred while creating the Kubernetes client")
	}

	labels := map[string]string{
		consts.KardinalAppIDLabelKey: consts.KardinalManagerAppIDLabelValue,
	}

	if err = kubernetesClientObj.RemoveNamespaceResourcesByLabels(ctx, kardinalNamespace, labels); err != nil {
		return stacktrace.Propagate(err, "An error occurred while removing the kardinal-manager from the cluster using labels '%+v'", labels)
	}

	return nil
}
