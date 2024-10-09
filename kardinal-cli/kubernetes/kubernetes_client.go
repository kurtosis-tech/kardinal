package kubernetes

import (
	"bytes"
	"context"
	"errors"
	"io"

	"github.com/kurtosis-tech/stacktrace"
	"gopkg.in/yaml.v3"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
)

const (
	fieldManager = "kardinal-cli"

	listOptionsTimeoutSeconds       int64 = 10
	deleteOptionsGracePeriodSeconds int64 = 0
)

type KubernetesClient struct {
	config          *rest.Config
	clientSet       *kubernetes.Clientset
	dynamicClient   *dynamic.DynamicClient
	discoveryMapper *restmapper.DeferredDiscoveryRESTMapper
}

func newKubernetesClient(config *rest.Config, clientSet *kubernetes.Clientset, dynamicClient *dynamic.DynamicClient, discoveryMapper *restmapper.DeferredDiscoveryRESTMapper) *KubernetesClient {
	return &KubernetesClient{config: config, clientSet: clientSet, dynamicClient: dynamicClient, discoveryMapper: discoveryMapper}
}

func (client *KubernetesClient) GetClientSet() *kubernetes.Clientset {
	return client.clientSet
}

func (client *KubernetesClient) GetConfig() *rest.Config {
	return client.config
}

func (client *KubernetesClient) GetService(ctx context.Context, namespaceName string, name string) (*corev1.Service, error) {
	serviceClient := client.clientSet.CoreV1().Services(namespaceName)
	serviceObj, err := serviceClient.Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred getting service '%s' from namespace '%s'", name, namespaceName)
	}
	return serviceObj, nil
}

func (client *KubernetesClient) GetDeploymentsByLabels(ctx context.Context, namespace string, labels map[string]string) (*appsv1.DeploymentList, error) {
	deploymentClient := client.clientSet.AppsV1().Deployments(namespace)

	opts := buildListOptionsFromLabels(labels)
	deploymentResult, err := deploymentClient.List(ctx, opts)
	if err != nil {
		return nil, stacktrace.Propagate(err, "Failed to list deployments with labels '%+v' in namespace '%s'", labels, namespace)
	}

	// Only return objects not tombstoned by Kubernetes
	var deploymentsNotMarkedForDeletionList []appsv1.Deployment
	for _, deployment := range deploymentResult.Items {
		deletionTimestamp := deployment.GetObjectMeta().GetDeletionTimestamp()
		if deletionTimestamp == nil {
			deploymentsNotMarkedForDeletionList = append(deploymentsNotMarkedForDeletionList, deployment)
		}
	}
	deploymentList := appsv1.DeploymentList{
		Items:    deploymentsNotMarkedForDeletionList,
		TypeMeta: deploymentResult.TypeMeta,
		ListMeta: deploymentResult.ListMeta,
	}

	return &deploymentList, nil
}

func (client *KubernetesClient) ApplyYamlFileContentInNamespace(ctx context.Context, namespace string, yamlFileContent []byte) error {
	yamlReader := bytes.NewReader(yamlFileContent)

	dec := yaml.NewDecoder(yamlReader)

	for {
		unstructuredObject := &unstructured.Unstructured{Object: map[string]interface{}{}}
		err := dec.Decode(unstructuredObject.Object)
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			return stacktrace.Propagate(err, "An error occurred decoding the unstructured object")
		}
		if unstructuredObject.Object == nil {
			return stacktrace.NewError("Expected to find the object value after decoding the unstructured object but it was not found")
		}

		groupVersionKind := unstructuredObject.GroupVersionKind()
		restMapping, err := client.discoveryMapper.RESTMapping(groupVersionKind.GroupKind(), groupVersionKind.Version)
		if err != nil {
			return stacktrace.Propagate(err, "An error occurred getting the rest mapping for GVK")
		}

		groupVersionResource := restMapping.Resource

		if unstructuredObject.GetNamespace() != "" && namespace != unstructuredObject.GetNamespace() {
			return stacktrace.NewError(
				"The namespace '%s' in resource '%s' kind '%s' is different from the main namespace '%s'",
				unstructuredObject.GetNamespace(),
				unstructuredObject.GetName(),
				unstructuredObject.GetKind(),
				namespace,
			)
		}

		applyOpts := metav1.ApplyOptions{FieldManager: fieldManager}

		var resource dynamic.ResourceInterface

		resource = client.dynamicClient.Resource(groupVersionResource)
		if unstructuredObject.GetNamespace() != "" {
			resource = client.dynamicClient.Resource(groupVersionResource).Namespace(namespace)
		}

		_, err = resource.Apply(ctx, unstructuredObject.GetName(), unstructuredObject, applyOpts)
		if err != nil {
			return stacktrace.Propagate(err, "An error occurred applying the k8s resource with name '%s' in namespace '%s'", unstructuredObject.GetName(), unstructuredObject.GetNamespace())
		}

	}
}

func (client *KubernetesClient) RemoveNamespaceResourcesByLabels(ctx context.Context, namespace string, labels map[string]string) error {
	opts := buildListOptionsFromLabels(labels)

	deleteOptions := metav1.NewDeleteOptions(deleteOptionsGracePeriodSeconds)

	// Delete deployments
	if err := client.clientSet.AppsV1().Deployments(namespace).DeleteCollection(ctx, *deleteOptions, opts); err != nil {
		return stacktrace.Propagate(err, "An error occurred removing deployments in namespace '%s'", namespace)
	}

	// Delete services one by one because there is not DeleteCollection function for services
	servicesToRemove, err := client.clientSet.CoreV1().Services(namespace).List(ctx, opts)
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred listing services")
	}

	for _, service := range servicesToRemove.Items {
		if err := client.clientSet.CoreV1().Services(namespace).Delete(ctx, service.GetName(), *deleteOptions); err != nil {
			return stacktrace.Propagate(err, "An error occurred removing service '%s' from namespace '%s'", service.GetName(), namespace)
		}
	}

	// Delete cluster role bindings
	if err := client.clientSet.RbacV1().ClusterRoleBindings().DeleteCollection(ctx, *deleteOptions, opts); err != nil {
		return stacktrace.Propagate(err, "An error occurred removing cluster role bindings")
	}

	// Delete cluster roles
	if err := client.clientSet.RbacV1().ClusterRoles().DeleteCollection(ctx, *deleteOptions, opts); err != nil {
		return stacktrace.Propagate(err, "An error occurred removing cluster roles")
	}

	// Delete service accounts
	if err := client.clientSet.CoreV1().ServiceAccounts(namespace).DeleteCollection(ctx, *deleteOptions, opts); err != nil {
		return stacktrace.Propagate(err, "An error occurred removing service accounts from namespace '%s'", namespace)
	}

	return nil
}

func (client *KubernetesClient) GetNamespacesByLabels(ctx context.Context, namespaceLabels map[string]string) (*corev1.NamespaceList, error) {
	namespaceClient := client.clientSet.CoreV1().Namespaces()

	listOptions := buildListOptionsFromLabels(namespaceLabels)
	namespaces, err := namespaceClient.List(ctx, listOptions)
	if err != nil {
		return nil, stacktrace.Propagate(err, "Failed to list namespaces with labels '%+v'", namespaceLabels)
	}

	// Only return objects not tombstoned by Kubernetes
	var namespacesNotMarkedForDeletionList []corev1.Namespace
	for _, namespace := range namespaces.Items {
		deletionTimestamp := namespace.GetObjectMeta().GetDeletionTimestamp()
		if deletionTimestamp == nil {
			namespacesNotMarkedForDeletionList = append(namespacesNotMarkedForDeletionList, namespace)
		}
	}
	namespacesNotMarkedForDeletionnamespaceList := corev1.NamespaceList{
		Items:    namespacesNotMarkedForDeletionList,
		TypeMeta: namespaces.TypeMeta,
		ListMeta: namespaces.ListMeta,
	}
	return &namespacesNotMarkedForDeletionnamespaceList, nil
}

func buildListOptionsFromLabels(labelsMap map[string]string) metav1.ListOptions {
	return metav1.ListOptions{
		TypeMeta: metav1.TypeMeta{
			Kind:       "",
			APIVersion: "",
		},
		LabelSelector:        labels.SelectorFromSet(labelsMap).String(),
		FieldSelector:        "",
		Watch:                false,
		AllowWatchBookmarks:  false,
		ResourceVersion:      "",
		ResourceVersionMatch: "",
		TimeoutSeconds:       int64Ptr(listOptionsTimeoutSeconds),
		Limit:                0,
		Continue:             "",
		SendInitialEvents:    nil,
	}
}

func int64Ptr(i int64) *int64 { return &i }
