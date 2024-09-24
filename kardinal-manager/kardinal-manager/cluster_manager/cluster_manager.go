package cluster_manager

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/kurtosis-tech/kardinal/libs/manager-kontrol-api/api/golang/types"
	"github.com/kurtosis-tech/stacktrace"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
	istio "istio.io/api/networking/v1alpha3"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
	securityv1beta1 "istio.io/client-go/pkg/apis/security/v1beta1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	net "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kardinal.kontrol/kardinal-manager/topology"
	gateway "sigs.k8s.io/gateway-api/apis/v1"
	gatewayclientset "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned"
)

const (
	listOptionsTimeoutSeconds         int64 = 10
	fieldManager                            = "kardinal-manager"
	deleteOptionsGracePeriodSeconds   int64 = 0
	istioLabel                              = "istio-injection"
	enabledIstioValue                       = "enabled"
	telepresenceRestartedAtAnnotation       = "telepresence.getambassador.io/restartedAt"
	istioSystemNamespace                    = "istio-system"

	// TODO move these values to a shared library between Kardinal Manager, Kontrol and Kardinal CLI
	kardinalLabelKey = "kardinal.dev"
	enabledKardinal  = "enabled"
)

var (
	globalListOptions = metav1.ListOptions{
		TypeMeta: metav1.TypeMeta{
			Kind:       "",
			APIVersion: "",
		},
		LabelSelector:        "",
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

	globalGetOptions = metav1.GetOptions{
		TypeMeta: metav1.TypeMeta{
			Kind:       "",
			APIVersion: "",
		},
		ResourceVersion: "",
	}

	globalCreateOptions = metav1.CreateOptions{
		TypeMeta: metav1.TypeMeta{
			Kind:       "",
			APIVersion: "",
		},
		DryRun: nil,
		// We need every object to have this field manager so that the Kurtosis objects can all seamlessly modify Kubernetes resources
		FieldManager:    fieldManager,
		FieldValidation: "",
	}

	globalUpdateOptions = metav1.UpdateOptions{
		TypeMeta: metav1.TypeMeta{
			Kind:       "",
			APIVersion: "",
		},
		DryRun: nil,
		// We need every object to have this field manager so that the Kurtosis objects can all seamlessly modify Kubernetes resources
		FieldManager:    fieldManager,
		FieldValidation: "",
	}

	globalDeletePolicy = metav1.DeletePropagationForeground

	globalDeleteOptions = metav1.DeleteOptions{
		TypeMeta: metav1.TypeMeta{
			Kind:       "",
			APIVersion: "",
		},
		GracePeriodSeconds: int64Ptr(deleteOptionsGracePeriodSeconds),
		Preconditions:      nil,
		OrphanDependents:   nil,
		PropagationPolicy:  &globalDeletePolicy,
		DryRun:             nil,
	}
)

type ClusterManager struct {
	kubernetesClient *kubernetesClient
	istioClient      *istioClient
	gatewayClient    *gatewayclientset.Clientset
}

func NewClusterManager(kubernetesClient *kubernetesClient, istioClient *istioClient, gatewayClient *gatewayclientset.Clientset) *ClusterManager {
	return &ClusterManager{kubernetesClient: kubernetesClient, istioClient: istioClient, gatewayClient: gatewayClient}
}

func (manager *ClusterManager) GetVirtualServices(ctx context.Context, namespace string) ([]*v1alpha3.VirtualService, error) {
	virtServiceClient := manager.istioClient.clientSet.NetworkingV1alpha3().VirtualServices(namespace)

	virtualServiceList, err := virtServiceClient.List(ctx, globalListOptions)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred retrieving virtual services from IstIo client.")
	}
	return virtualServiceList.Items, nil
}

func (manager *ClusterManager) GetVirtualService(ctx context.Context, namespace string, name string) (*v1alpha3.VirtualService, error) {
	virtServiceClient := manager.istioClient.clientSet.NetworkingV1alpha3().VirtualServices(namespace)

	virtualService, err := virtServiceClient.Get(ctx, name, globalGetOptions)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred retrieving virtual service '%s' from IstIo client", name)
	}
	return virtualService, nil
}

func (manager *ClusterManager) GetDestinationRules(ctx context.Context, namespace string) ([]*v1alpha3.DestinationRule, error) {
	destRuleClient := manager.istioClient.clientSet.NetworkingV1alpha3().DestinationRules(namespace)

	destinationRules, err := destRuleClient.List(ctx, globalListOptions)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred retrieving destination rules.")
	}
	return destinationRules.Items, nil
}

func (manager *ClusterManager) GetDestinationRule(ctx context.Context, namespace string, rule string) (*v1alpha3.DestinationRule, error) {
	destRuleClient := manager.istioClient.clientSet.NetworkingV1alpha3().DestinationRules(namespace)

	destinationRule, err := destRuleClient.Get(ctx, rule, globalGetOptions)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred retrieving destination rule '%s' from IstIo client", rule)
	}
	return destinationRule, nil
}

// how to expose API to configure ordering of routing rule? https://istio.io/latest/docs/concepts/traffic-management/#routing-rule-precedence
func (manager *ClusterManager) AddRoutingRule(ctx context.Context, namespace string, vsName string, routingRule *istio.HTTPRoute) error {
	virtServiceClient := manager.istioClient.clientSet.NetworkingV1alpha3().VirtualServices(namespace)

	vs, err := virtServiceClient.Get(ctx, vsName, globalGetOptions)
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred retrieving virtual service '%s'", vsName)
	}
	// always prepend routing rules due to routing rule precedence
	vs.Spec.Http = append([]*istio.HTTPRoute{routingRule}, vs.Spec.Http...)
	_, err = virtServiceClient.Update(ctx, vs, metav1.UpdateOptions{})
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred updating virtual service '%s' with routing rule: %v", vsName, routingRule)
	}
	return nil
}

func (manager *ClusterManager) AddSubset(ctx context.Context, namespace string, drName string, subset *istio.Subset) error {
	destRuleClient := manager.istioClient.clientSet.NetworkingV1alpha3().DestinationRules(namespace)

	dr, err := destRuleClient.Get(ctx, drName, globalGetOptions)
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred retrieving destination rule '%s'", drName)
	}
	// if there already exists a subset for the same, just update it
	shouldAddNewSubset := true
	for _, s := range dr.Spec.Subsets {
		if s.Name == subset.Name {
			s = subset
			shouldAddNewSubset = false
		}
	}
	if shouldAddNewSubset {
		dr.Spec.Subsets = append(dr.Spec.Subsets, subset)
	}
	_, err = destRuleClient.Update(ctx, dr, metav1.UpdateOptions{})
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred updating destination rule '%s' with subset: %v", drName, subset)
	}
	return nil
}

func (manager *ClusterManager) GetTopologyForNameSpace(namespace string) (map[string]*topology.Node, error) {
	return manager.istioClient.topologyManager.FetchTopology(namespace)
}

func (manager *ClusterManager) ApplyClusterResources(ctx context.Context, clusterResources *types.ClusterResources) error {
	if !isValid(clusterResources) {
		logrus.Debugf("the received cluster resources is not valid, nothing to apply.")
		return nil
	}

	allNSs := [][]string{
		lo.Uniq(lo.Map(*clusterResources.Services, func(item corev1.Service, _ int) string { return item.Namespace })),
		lo.Uniq(lo.Map(*clusterResources.Deployments, func(item appsv1.Deployment, _ int) string { return item.Namespace })),
		lo.Uniq(lo.Map(*clusterResources.VirtualServices, func(item v1alpha3.VirtualService, _ int) string { return item.Namespace })),
		lo.Uniq(lo.Map(*clusterResources.DestinationRules, func(item v1alpha3.DestinationRule, _ int) string { return item.Namespace })),
		lo.Uniq(lo.Map(*clusterResources.Gateways, func(item gateway.Gateway, _ int) string { return item.Namespace })),
		lo.Uniq(lo.Map(*clusterResources.HttpRoutes, func(item gateway.HTTPRoute, _ int) string { return item.Namespace })),
	}

	if clusterResources.EnvoyFilters != nil {
		envoyFiltersNS := lo.Uniq(lo.Map(*clusterResources.EnvoyFilters, func(item v1alpha3.EnvoyFilter, _ int) string { return item.Namespace }))
		allNSs = append(allNSs, [][]string{envoyFiltersNS}...)
	}

	if clusterResources.AuthorizationPolicies != nil {
		authPoliciesNS := lo.Uniq(lo.Map(*clusterResources.AuthorizationPolicies, func(item securityv1beta1.AuthorizationPolicy, _ int) string { return item.Namespace }))
		allNSs = append(allNSs, [][]string{authPoliciesNS}...)
	}

	uniqueNamespaces := lo.ReplaceAll(lo.Uniq(lo.Flatten(allNSs)), "", "default")

	for _, namespace := range uniqueNamespaces {
		if err := manager.ensureNamespace(ctx, namespace); err != nil {
			return stacktrace.Propagate(err, "An error occurred while creating or updating cluster namespace '%s'", namespace)
		}
	}

	for _, service := range *clusterResources.Services {
		if err := manager.createOrUpdateService(ctx, &service); err != nil {
			return stacktrace.Propagate(err, "An error occurred while creating or updating service '%s'", service.GetName())
		}
	}

	for _, deployment := range *clusterResources.Deployments {
		if err := manager.createOrUpdateDeployment(ctx, &deployment); err != nil {
			return stacktrace.Propagate(err, "An error occurred while creating or updating deployment '%s'", deployment.GetName())
		}
	}

	for _, virtualService := range *clusterResources.VirtualServices {
		if err := manager.createOrUpdateVirtualService(ctx, &virtualService); err != nil {
			return stacktrace.Propagate(err, "An error occurred while creating or updating virtual service '%s'", virtualService.GetName())
		}
	}

	for _, destinationRule := range *clusterResources.DestinationRules {
		if err := manager.createOrUpdateDestinationRule(ctx, &destinationRule); err != nil {
			return stacktrace.Propagate(err, "An error occurred while creating or updating destination rule '%s'", destinationRule.GetName())
		}
	}

	envoyFiltersLen := 0
	if clusterResources.EnvoyFilters != nil {
		envoyFiltersLen = len(*clusterResources.EnvoyFilters)
	}

	authPoliciesLen := 0
	if clusterResources.AuthorizationPolicies != nil {
		authPoliciesLen = len(*clusterResources.AuthorizationPolicies)
	}

	logrus.Infof("Have %d envoy filters and %d policies to apply", envoyFiltersLen, authPoliciesLen)
	if clusterResources.EnvoyFilters != nil {
		for _, envoyFilter := range *clusterResources.EnvoyFilters {
			if err := manager.createOrUpdateEnvoyFilter(ctx, &envoyFilter); err != nil {
				return stacktrace.Propagate(err, "An error occurred while creating or updating envoy filter '%s'", envoyFilter.GetName())
			}
		}
	}

	if clusterResources.AuthorizationPolicies != nil {
		for _, policy := range *clusterResources.AuthorizationPolicies {
			if err := manager.createOrUpdateAuthorizationPolicies(ctx, &policy); err != nil {
				return stacktrace.Propagate(err, "An error occurred while creating or updating envoy policies '%s'", policy.GetName())
			}
		}
	}

	for _, gateway := range *clusterResources.Gateways {
		if err := manager.createOrUpdateGateway(ctx, &gateway); err != nil {
			return stacktrace.Propagate(err, "An error occurred while creating or updating the cluster gateway")
		}
	}

	for _, httpRoute := range *clusterResources.HttpRoutes {
		if err := manager.createOrUpdateHttpRoute(ctx, &httpRoute); err != nil {
			return stacktrace.Propagate(err, "An error occurred while creating or updating the http route")
		}
	}

	for _, ingress := range *clusterResources.Ingresses {
		if err := manager.createOrUpdateIngress(ctx, &ingress); err != nil {
			return stacktrace.Propagate(err, "An error occurred while creating or updating the ingress")
		}
	}

	return nil
}

func (manager *ClusterManager) CleanUpClusterResources(ctx context.Context, clusterResources *types.ClusterResources) error {
	if !isValid(clusterResources) {
		logrus.Debugf("the received cluster resources is not valid, nothing to clean up.")
		return nil
	}

	// Clean up virtual services
	virtualServicesByNS := lo.GroupBy(*clusterResources.VirtualServices, func(item v1alpha3.VirtualService) string { return item.Namespace })
	for namespace, virtualServices := range virtualServicesByNS {
		if err := manager.cleanUpVirtualServicesInNamespace(ctx, namespace, virtualServices); err != nil {
			return stacktrace.Propagate(err, "An error occurred cleaning up virtual services '%+v' in namespace '%s'", virtualServices, namespace)
		}
	}

	// Clean up destination rules
	destinationRulesByNS := lo.GroupBy(*clusterResources.DestinationRules, func(item v1alpha3.DestinationRule) string {
		return item.Namespace
	})
	for namespace, destinationRules := range destinationRulesByNS {
		if err := manager.cleanUpDestinationRulesInNamespace(ctx, namespace, destinationRules); err != nil {
			return stacktrace.Propagate(err, "An error occurred cleaning up destination rules '%+v' in namespace '%s'", destinationRules, namespace)
		}
	}

	// Clean up Ingresses
	ingressesByNs := lo.GroupBy(*clusterResources.Ingresses, func(item net.Ingress) string {
		return item.Namespace
	})
	for namespace, ingresses := range ingressesByNs {
		if err := manager.cleanUpIngressesInNamespace(ctx, namespace, ingresses); err != nil {
			return stacktrace.Propagate(err, "An error occurred cleaning up ingresses '%+v' in namespace '%s'", ingresses, namespace)
		}
	}

	// Clean up http routes
	routesByNs := lo.GroupBy(*clusterResources.HttpRoutes, func(item gateway.HTTPRoute) string {
		return item.Namespace
	})
	for namespace, routes := range routesByNs {
		if err := manager.cleanUpHttpRoutesInNamespace(ctx, namespace, routes); err != nil {
			return stacktrace.Propagate(err, "An error occurred cleaning up http routes '%+v' in namespace '%s'", routes, namespace)
		}
	}

	// Clean up gateway
	gatewaysByNs := lo.GroupBy(*clusterResources.Gateways, func(item gateway.Gateway) string {
		return item.Namespace
	})
	for namespace, gateways := range gatewaysByNs {
		if err := manager.cleanUpGatewaysInNamespace(ctx, namespace, gateways); err != nil {
			return stacktrace.Propagate(err, "An error occurred cleaning up gateways '%+v' in namespace '%s'", gateways, namespace)
		}
	}

	// Cleanup envoy filters
	if clusterResources.EnvoyFilters != nil {
		envoyFiltersByNS := lo.GroupBy(*clusterResources.EnvoyFilters, func(item v1alpha3.EnvoyFilter) string {
			return item.Namespace
		})
		for namespace, envoyFilters := range envoyFiltersByNS {
			if err := manager.cleanupEnvoyFiltersInNamespace(ctx, namespace, envoyFilters); err != nil {
				return stacktrace.Propagate(err, "An error occurred cleaning up envoy filters '%+v' in namespace '%s'", envoyFilters, namespace)
			}
		}
	}

	// Clean up services
	servicesByNS := lo.GroupBy(*clusterResources.Services, func(item corev1.Service) string {
		return item.Namespace
	})
	for namespace, services := range servicesByNS {
		gateways := gatewaysByNs[namespace]
		if err := manager.cleanUpServicesInNamespace(ctx, namespace, services, gateways); err != nil {
			return stacktrace.Propagate(err, "An error occurred cleaning up services '%+v' in namespace '%s'", services, namespace)
		}
	}

	// Clean up deployments
	deploymentsByNS := lo.GroupBy(*clusterResources.Deployments, func(item appsv1.Deployment) string { return item.Namespace })
	for namespace, deployments := range deploymentsByNS {
		if err := manager.cleanUpDeploymentsInNamespace(ctx, namespace, deployments); err != nil {
			return stacktrace.Propagate(err, "An error occurred cleaning up deployments '%+v' in namespace '%s'", deployments, namespace)
		}
	}

	// Cleanup authorization policies
	if clusterResources.AuthorizationPolicies != nil {
		authorizationPoliciesByNS := lo.GroupBy(*clusterResources.AuthorizationPolicies, func(item securityv1beta1.AuthorizationPolicy) string {
			return item.Namespace
		})
		for namespace, authorizationPolicies := range authorizationPoliciesByNS {
			if err := manager.cleanupAuthorizationPoliciesInNamespace(ctx, namespace, authorizationPolicies); err != nil {
				return stacktrace.Propagate(err, "An error occurred cleaning up authorization policies '%+v' in namespace '%s'", authorizationPolicies, namespace)
			}
		}
	}

	return nil
}

func (manager *ClusterManager) ensureNamespace(ctx context.Context, name string) error {

	if name == istioSystemNamespace {
		// Some resources might be under the istio system namespace but we don't want to alter
		// this namespace because it is managed by Istio
		return nil
	}

	existingNamespace, err := manager.kubernetesClient.clientSet.CoreV1().Namespaces().Get(ctx, name, metav1.GetOptions{})
	if err == nil && existingNamespace != nil {
		value, found := existingNamespace.Labels[istioLabel]
		if !found || value != enabledIstioValue {
			existingNamespace.Labels[istioLabel] = enabledIstioValue
		}
		value, found = existingNamespace.Labels[kardinalLabelKey]
		if !found || value != enabledKardinal {
			existingNamespace.Labels[kardinalLabelKey] = enabledKardinal
		}
		_, err = manager.kubernetesClient.clientSet.CoreV1().Namespaces().Update(ctx, existingNamespace, globalUpdateOptions)
		if err != nil {
			return stacktrace.Propagate(err, "Failed to update Namespace: %s", name)
		}
	} else {
		newNamespace := corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: name,
				Labels: map[string]string{
					istioLabel:       enabledIstioValue,
					kardinalLabelKey: enabledKardinal,
				},
			},
		}
		_, err = manager.kubernetesClient.clientSet.CoreV1().Namespaces().Create(ctx, &newNamespace, globalCreateOptions)
		if err != nil {
			return stacktrace.Propagate(err, "Failed to create Namespace: %s", name)
		}
	}

	return nil
}

func (manager *ClusterManager) createOrUpdateService(ctx context.Context, service *corev1.Service) error {
	serviceClient := manager.kubernetesClient.clientSet.CoreV1().Services(service.Namespace)
	existingService, err := serviceClient.Get(ctx, service.Name, metav1.GetOptions{})
	if err != nil {
		_, err = serviceClient.Create(ctx, service, globalCreateOptions)
		if err != nil {
			return stacktrace.Propagate(err, "Failed to create service: %s", service.GetName())
		}
	} else {
		if !deepCheckEqual(existingService.Spec, service.Spec) {
			service.ResourceVersion = existingService.ResourceVersion
			_, err = serviceClient.Update(ctx, service, globalUpdateOptions)
			if err != nil {
				return stacktrace.Propagate(err, "Failed to update service: %s", service.GetName())
			}
		}
	}

	return nil
}

func (manager *ClusterManager) createOrUpdateDeployment(ctx context.Context, deployment *appsv1.Deployment) error {
	deploymentClient := manager.kubernetesClient.clientSet.AppsV1().Deployments(deployment.Namespace)
	existingDeployment, err := deploymentClient.Get(ctx, deployment.Name, metav1.GetOptions{})
	if err != nil {
		_, err = deploymentClient.Create(ctx, deployment, globalCreateOptions)
		if err != nil {
			return stacktrace.Propagate(err, "Failed to create deployment: %s", deployment.GetName())
		}
	} else {
		if !deepCheckEqual(existingDeployment.Spec, deployment.Spec) {
			updateDeploymentWithRelevantValuesFromCurrentDeployment(deployment, existingDeployment)
			_, err = deploymentClient.Update(ctx, deployment, globalUpdateOptions)
			if err != nil {
				return stacktrace.Propagate(err, "Failed to update deployment: %s", deployment.GetName())
			}
		}
	}

	return nil
}

func updateDeploymentWithRelevantValuesFromCurrentDeployment(newDeployment *appsv1.Deployment, currentDeployment *appsv1.Deployment) {
	newDeployment.ResourceVersion = currentDeployment.ResourceVersion
	// merge annotations
	newAnnotations := newDeployment.Spec.Template.GetAnnotations()
	currentAnnotations := currentDeployment.Spec.Template.GetAnnotations()

	for annotationKey, annotationValue := range currentAnnotations {
		if annotationKey == telepresenceRestartedAtAnnotation {
			// This key is necessary for Kardinal/Telepresence (https://www.telepresence.io/) integration
			// keeping this annotation because otherwise the telepresence traffic-agent container will be removed from the pod
			newAnnotations[annotationKey] = annotationValue
		}
	}
	newDeployment.Spec.Template.Annotations = newAnnotations
}

func (manager *ClusterManager) createOrUpdateVirtualService(ctx context.Context, virtualService *v1alpha3.VirtualService) error {
	virtServiceClient := manager.istioClient.clientSet.NetworkingV1alpha3().VirtualServices(virtualService.GetNamespace())

	existingVirtService, err := virtServiceClient.Get(ctx, virtualService.Name, metav1.GetOptions{})
	if err != nil {
		_, err = virtServiceClient.Create(ctx, virtualService, globalCreateOptions)
		if err != nil {
			return stacktrace.Propagate(err, "Failed to create virtual service: %s", virtualService.GetName())
		}
	} else {
		if !deepCheckEqual(existingVirtService.Spec, virtualService.Spec) {
			virtualService.ResourceVersion = existingVirtService.ResourceVersion
			_, err = virtServiceClient.Update(ctx, virtualService, globalUpdateOptions)
			if err != nil {
				return stacktrace.Propagate(err, "Failed to update virtual service: %s", virtualService.GetName())
			}
		}
	}

	return nil
}

func (manager *ClusterManager) createOrUpdateDestinationRule(ctx context.Context, destinationRule *v1alpha3.DestinationRule) error {
	destRuleClient := manager.istioClient.clientSet.NetworkingV1alpha3().DestinationRules(destinationRule.GetNamespace())

	existingDestRule, err := destRuleClient.Get(ctx, destinationRule.Name, metav1.GetOptions{})
	if err != nil {
		_, err = destRuleClient.Create(ctx, destinationRule, globalCreateOptions)
		if err != nil {
			return stacktrace.Propagate(err, "Failed to create destination rule: %s", destinationRule.GetName())
		}
	} else {
		if !deepCheckEqual(existingDestRule.Spec, destinationRule.Spec) {
			destinationRule.ResourceVersion = existingDestRule.ResourceVersion
			_, err = destRuleClient.Update(ctx, destinationRule, globalUpdateOptions)
			if err != nil {
				return stacktrace.Propagate(err, "Failed to update destination rule: %s", destinationRule.GetName())
			}
		}
	}

	return nil
}

func (manager *ClusterManager) createOrUpdateGateway(ctx context.Context, gateway *gateway.Gateway) error {
	gatewayClient := manager.gatewayClient.GatewayV1().Gateways(gateway.GetNamespace())
	existingGateway, err := gatewayClient.Get(ctx, gateway.Name, metav1.GetOptions{})
	if err != nil {
		_, err = gatewayClient.Create(ctx, gateway, globalCreateOptions)
		if err != nil {
			return stacktrace.Propagate(err, "Failed to create gateway: %s", gateway.GetName())
		}
	} else {
		if !deepCheckEqual(existingGateway.Spec, gateway.Spec) {
			gateway.ResourceVersion = existingGateway.ResourceVersion
			_, err = gatewayClient.Update(ctx, gateway, globalUpdateOptions)
			if err != nil {
				return stacktrace.Propagate(err, "Failed to update gateway: %s", gateway.GetName())
			}
		}
	}

	return nil
}

func (manager *ClusterManager) createOrUpdateHttpRoute(ctx context.Context, route *gateway.HTTPRoute) error {
	routeClient := manager.gatewayClient.GatewayV1().HTTPRoutes(route.GetNamespace())
	existingRoute, err := routeClient.Get(ctx, route.Name, metav1.GetOptions{})
	if err != nil {
		_, err = routeClient.Create(ctx, route, globalCreateOptions)
		if err != nil {
			return stacktrace.Propagate(err, "Failed to create route: %s", route.GetName())
		}
	} else {
		if !deepCheckEqual(existingRoute.Spec, route.Spec) {
			route.ResourceVersion = existingRoute.ResourceVersion
			_, err = routeClient.Update(ctx, route, globalUpdateOptions)
			if err != nil {
				return stacktrace.Propagate(err, "Failed to update route: %s", route.GetName())
			}
		}
	}

	return nil
}

func (manager *ClusterManager) createOrUpdateIngress(ctx context.Context, ingress *net.Ingress) error {
	routeClient := manager.kubernetesClient.clientSet.NetworkingV1().Ingresses(ingress.GetNamespace())
	existingRoute, err := routeClient.Get(ctx, ingress.Name, metav1.GetOptions{})
	if err != nil {
		_, err = routeClient.Create(ctx, ingress, globalCreateOptions)
		if err != nil {
			return stacktrace.Propagate(err, "Failed to create ingress: %s", ingress.GetName())
		}
	} else {
		if !deepCheckEqual(existingRoute.Spec, ingress.Spec) {
			ingress.ResourceVersion = existingRoute.ResourceVersion
			_, err = routeClient.Update(ctx, ingress, globalUpdateOptions)
			if err != nil {
				return stacktrace.Propagate(err, "Failed to update ingress: %s", ingress.GetName())
			}
		}
	}

	return nil
}

func (manager *ClusterManager) createOrUpdateEnvoyFilter(ctx context.Context, filter *v1alpha3.EnvoyFilter) error {
	envoyFilterClient := manager.istioClient.clientSet.NetworkingV1alpha3().EnvoyFilters(filter.GetNamespace())
	existingFilter, err := envoyFilterClient.Get(ctx, filter.Name, metav1.GetOptions{})
	if err != nil {
		_, err = envoyFilterClient.Create(ctx, filter, globalCreateOptions)
		if err != nil {
			return stacktrace.Propagate(err, "Failed to create envoy filter: %s", filter.GetName())
		}
	} else {
		if !deepCheckEqual(existingFilter.Spec, filter.Spec) {
			filter.ResourceVersion = existingFilter.ResourceVersion
			_, err = envoyFilterClient.Update(ctx, filter, globalUpdateOptions)
			if err != nil {
				return stacktrace.Propagate(err, "Failed to update filter: %s", filter.GetName())
			}
		}
	}
	return nil
}

func (manager *ClusterManager) createOrUpdateAuthorizationPolicies(ctx context.Context, policy *securityv1beta1.AuthorizationPolicy) error {
	authorizationPolicyClient := manager.istioClient.clientSet.SecurityV1beta1().AuthorizationPolicies(policy.GetNamespace())
	existingPolicy, err := authorizationPolicyClient.Get(ctx, policy.Name, metav1.GetOptions{})
	if err != nil {
		_, err = authorizationPolicyClient.Create(ctx, policy, globalCreateOptions)
		if err != nil {
			return stacktrace.Propagate(err, "Failed to create policy: %s", policy.GetName())
		}
	} else {
		if !deepCheckEqual(existingPolicy.Spec, policy.Spec) {
			policy.ResourceVersion = existingPolicy.ResourceVersion
			_, err = authorizationPolicyClient.Update(ctx, policy, globalUpdateOptions)
			if err != nil {
				return stacktrace.Propagate(err, "Failed to update policy: %s", policy.GetName())
			}
		}
	}
	return nil
}

func (manager *ClusterManager) cleanUpServicesInNamespace(ctx context.Context, namespace string, servicesToKeep []corev1.Service, gateways []gateway.Gateway) error {
	serviceClient := manager.kubernetesClient.clientSet.CoreV1().Services(namespace)
	gatewayNames := lo.Map(gateways, func(item gateway.Gateway, _ int) string { return item.Name })
	allServices, err := serviceClient.List(ctx, globalListOptions)
	if err != nil {
		return stacktrace.Propagate(err, "Failed to list services in namespace %s", namespace)
	}

	allServicesSkippingGateways := lo.Filter(allServices.Items, func(item corev1.Service, _ int) bool {
		_, found := lo.Find(gatewayNames, func(gatewayName string) bool { return strings.HasPrefix(item.Name, gatewayName) })
		if found {
			logrus.Infof("Skipping deletion service %s because it seems a service from one of the gateways %v", item.Name, gatewayNames)
		}
		return !found
	})
	for _, service := range allServicesSkippingGateways {
		_, exists := lo.Find(servicesToKeep, func(item corev1.Service) bool { return item.Name == service.Name })
		if !exists {
			logrus.Infof("Deleting service %s", service.Name)
			err = serviceClient.Delete(ctx, service.Name, globalDeleteOptions)
			if err != nil {
				return stacktrace.Propagate(err, "Failed to delete service %s", service.GetName())
			}
		}
	}
	return nil
}

func (manager *ClusterManager) cleanUpDeploymentsInNamespace(ctx context.Context, namespace string, deploymentsToKeep []appsv1.Deployment) error {
	deploymentClient := manager.kubernetesClient.clientSet.AppsV1().Deployments(namespace)
	allDeployments, err := deploymentClient.List(ctx, globalListOptions)
	if err != nil {
		return stacktrace.Propagate(err, "Failed to list deployments in namespace %s", namespace)
	}
	for _, deployment := range allDeployments.Items {
		_, exists := lo.Find(deploymentsToKeep, func(item appsv1.Deployment) bool { return item.Name == deployment.Name })
		if !exists {
			err = deploymentClient.Delete(ctx, deployment.Name, globalDeleteOptions)
			if err != nil {
				return stacktrace.Propagate(err, "Failed to delete deployment %s", deployment.GetName())
			}
		}
	}
	return nil
}

func (manager *ClusterManager) cleanUpVirtualServicesInNamespace(ctx context.Context, namespace string, virtualServicesToKeep []v1alpha3.VirtualService) error {
	virtServiceClient := manager.istioClient.clientSet.NetworkingV1alpha3().VirtualServices(namespace)
	allVirtServices, err := virtServiceClient.List(ctx, globalListOptions)
	if err != nil {
		return stacktrace.Propagate(err, "Failed to list virtual services in namespace %s", namespace)
	}
	for _, virtService := range allVirtServices.Items {
		_, exists := lo.Find(virtualServicesToKeep, func(item v1alpha3.VirtualService) bool { return item.Name == virtService.Name })
		if !exists {
			err = virtServiceClient.Delete(ctx, virtService.Name, globalDeleteOptions)
			if err != nil {
				return stacktrace.Propagate(err, "Failed to delete virtual service %s", virtService.GetName())
			}
		}
	}

	return nil
}

func (manager *ClusterManager) cleanUpDestinationRulesInNamespace(ctx context.Context, namespace string, destinationRulesToKeep []v1alpha3.DestinationRule) error {
	destRuleClient := manager.istioClient.clientSet.NetworkingV1alpha3().DestinationRules(namespace)
	allDestRules, err := destRuleClient.List(ctx, globalListOptions)
	if err != nil {
		return stacktrace.Propagate(err, "Failed to list destination rules in namespace %s", namespace)
	}
	for _, destRule := range allDestRules.Items {
		_, exists := lo.Find(destinationRulesToKeep, func(item v1alpha3.DestinationRule) bool { return item.Name == destRule.Name })
		if !exists {
			err = destRuleClient.Delete(ctx, destRule.Name, globalDeleteOptions)
			if err != nil {
				return stacktrace.Propagate(err, "Failed to delete destination rule %s", destRule.GetName())
			}
		}
	}

	return nil
}

func (manager *ClusterManager) cleanUpGatewaysInNamespace(ctx context.Context, namespace string, gatewaysToKeep []gateway.Gateway) error {
	gatewayClient := manager.gatewayClient.GatewayV1().Gateways(namespace)
	allGateways, err := gatewayClient.List(ctx, globalListOptions)
	if err != nil {
		return stacktrace.Propagate(err, "Failed to list gateways in namespace %s", namespace)
	}
	for _, gatewayItem := range allGateways.Items {
		_, exists := lo.Find(gatewaysToKeep, func(item gateway.Gateway) bool { return item.Name == gatewayItem.Name })
		if !exists {
			err = gatewayClient.Delete(ctx, gatewayItem.Name, globalDeleteOptions)
			if err != nil {
				return stacktrace.Propagate(err, "Failed to delete gateway %s", gatewayItem.GetName())
			}
		}
	}

	return nil
}

func (manager *ClusterManager) cleanUpIngressesInNamespace(ctx context.Context, namespace string, ingressesToKeep []net.Ingress) error {
	ingressClient := manager.kubernetesClient.clientSet.NetworkingV1().Ingresses(namespace)
	allingresss, err := ingressClient.List(ctx, globalListOptions)
	if err != nil {
		return stacktrace.Propagate(err, "Failed to list Ingress in namespace %s", namespace)
	}
	for _, ingressItem := range allingresss.Items {
		_, exists := lo.Find(ingressesToKeep, func(item net.Ingress) bool { return item.Name == ingressItem.Name })
		if !exists {
			err = ingressClient.Delete(ctx, ingressItem.Name, globalDeleteOptions)
			if err != nil {
				return stacktrace.Propagate(err, "Failed to delete gateway %s", ingressItem.GetName())
			}
		}
	}

	return nil
}

func (manager *ClusterManager) cleanUpHttpRoutesInNamespace(ctx context.Context, namespace string, routesToKeep []gateway.HTTPRoute) error {
	routeClient := manager.gatewayClient.GatewayV1().HTTPRoutes(namespace)
	allRoutes, err := routeClient.List(ctx, globalListOptions)
	if err != nil {
		return stacktrace.Propagate(err, "Failed to list Routes in namespace %s", namespace)
	}
	for _, routeItem := range allRoutes.Items {
		_, exists := lo.Find(routesToKeep, func(item gateway.HTTPRoute) bool { return item.Name == routeItem.Name })
		if !exists {
			err = routeClient.Delete(ctx, routeItem.Name, globalDeleteOptions)
			if err != nil {
				return stacktrace.Propagate(err, "Failed to delete gateway %s", routeItem.GetName())
			}
		}
	}

	return nil
}

func (manager *ClusterManager) cleanupEnvoyFiltersInNamespace(ctx context.Context, namespace string, filtersToKeep []v1alpha3.EnvoyFilter) error {
	envoyFilterClient := manager.istioClient.clientSet.NetworkingV1alpha3().EnvoyFilters(namespace)
	allEnvoyFilters, err := envoyFilterClient.List(ctx, globalListOptions)
	if err != nil {
		return stacktrace.Propagate(err, "Failed to list envoy filter in namespace %s", namespace)
	}
	for _, filter := range allEnvoyFilters.Items {
		_, exists := lo.Find(filtersToKeep, func(item v1alpha3.EnvoyFilter) bool { return item.Name == filter.Name })
		if !exists {
			err = envoyFilterClient.Delete(ctx, filter.Name, globalDeleteOptions)
			if err != nil {
				return stacktrace.Propagate(err, "Failed to delete filter %s", filter.GetName())
			}
		}
	}

	return nil
}

func (manager *ClusterManager) cleanupAuthorizationPoliciesInNamespace(ctx context.Context, namespace string, policiesToKeep []securityv1beta1.AuthorizationPolicy) error {
	authorizationPolicyClient := manager.istioClient.clientSet.SecurityV1beta1().AuthorizationPolicies(namespace)
	allPolicies, err := authorizationPolicyClient.List(ctx, globalListOptions)
	if err != nil {
		return stacktrace.Propagate(err, "Failed to list policy in namespace %s", namespace)
	}
	for _, policy := range allPolicies.Items {
		_, exists := lo.Find(policiesToKeep, func(item securityv1beta1.AuthorizationPolicy) bool { return item.Name == policy.Name })
		if !exists {
			err = authorizationPolicyClient.Delete(ctx, policy.Name, globalDeleteOptions)
			if err != nil {
				return stacktrace.Propagate(err, "Failed to delete policy %s", policy.GetName())
			}
		}
	}

	return nil
}

func int64Ptr(i int64) *int64 { return &i }

func isValid(clusterResources *types.ClusterResources) bool {
	if clusterResources == nil {
		logrus.Debugf("cluster resources is nil.")
		return false
	}

	if clusterResources.Gateways == nil &&
		clusterResources.HttpRoutes == nil &&
		clusterResources.Deployments == nil &&
		clusterResources.DestinationRules == nil &&
		clusterResources.Services == nil &&
		clusterResources.VirtualServices == nil {
		logrus.Debugf("cluster resources is empty.")
		return false
	}

	return true
}

func deepCheckEqual(a, b interface{}) bool {
	aj, err := json.Marshal(a)
	if err != nil {
		return false
	}
	bj, err := json.Marshal(b)
	if err != nil {
		return false
	}
	return string(aj) == string(bj)
}
