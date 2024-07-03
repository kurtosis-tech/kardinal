package deployment

import (
	"bytes"
	"context"
	"errors"
	"github.com/kurtosis-tech/stacktrace"
	"gopkg.in/yaml.v3"
	"io"
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

type kubernetesClient struct {
	config          *rest.Config
	clientSet       *kubernetes.Clientset
	dynamicClient   *dynamic.DynamicClient
	discoveryMapper *restmapper.DeferredDiscoveryRESTMapper
}

func newKubernetesClient(config *rest.Config, clientSet *kubernetes.Clientset, dynamicClient *dynamic.DynamicClient, discoveryMapper *restmapper.DeferredDiscoveryRESTMapper) *kubernetesClient {
	return &kubernetesClient{config: config, clientSet: clientSet, dynamicClient: dynamicClient, discoveryMapper: discoveryMapper}
}

func (client *kubernetesClient) ApplyYamlFileContentInNamespace(ctx context.Context, namespace string, yamlFileContent []byte) error {
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

func (client *kubernetesClient) RemoveNamespaceResourcesByLabels(ctx context.Context, namespace string, labels map[string]string) error {

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
