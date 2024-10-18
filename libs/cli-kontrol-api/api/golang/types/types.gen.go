// Package types provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version 2.1.0 DO NOT EDIT.
package types

import (
	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	gateway "sigs.k8s.io/gateway-api/apis/v1"
)

// Defines values for NodeType.
const (
	External NodeType = "external"
	Gateway  NodeType = "gateway"
	Service  NodeType = "service"
)

// ClusterTopology defines model for ClusterTopology.
type ClusterTopology struct {
	Edges []Edge `json:"edges"`
	Nodes []Node `json:"nodes"`
}

// DeploymentConfig defines model for DeploymentConfig.
type DeploymentConfig struct {
	Deployment appv1.Deployment `json:"deployment"`
}

// Edge defines model for Edge.
type Edge struct {
	// Label Label for the edge.
	Label *string `json:"label,omitempty"`

	// Source The identifier of the source node of the edge.
	Source string `json:"source"`

	// Target The identifier of the target node of the edge.
	Target string `json:"target"`
}

// Flow defines model for Flow.
type Flow struct {
	AccessEntry []IngressAccessEntry `json:"access-entry"`
	FlowId      string               `json:"flow-id"`
	IsBaseline  *bool                `json:"is-baseline,omitempty"`
}

// FlowSpec defines model for FlowSpec.
type FlowSpec = []struct {
	EnvVarOverrides       *map[string]string `json:"env-var-overrides,omitempty"`
	ImageLocator          string             `json:"image-locator"`
	SecretEnvVarOverrides *map[string]string `json:"secret-env-var-overrides,omitempty"`
	ServiceName           string             `json:"service-name"`
}

// GatewayConfig defines model for GatewayConfig.
type GatewayConfig struct {
	Gateway gateway.Gateway `json:"gateway"`
}

// IngressAccessEntry defines model for IngressAccessEntry.
type IngressAccessEntry struct {
	FlowId        string `json:"flow-id"`
	FlowNamespace string `json:"flow-namespace"`
	Hostname      string `json:"hostname"`
	Namespace     string `json:"namespace"`
	Service       string `json:"service"`
	Type          string `json:"type"`
}

// IngressConfig defines model for IngressConfig.
type IngressConfig struct {
	Ingress networkingv1.Ingress `json:"ingress"`
}

// MainClusterConfig defines model for MainClusterConfig.
type MainClusterConfig struct {
	DeploymentConfigs  *[]DeploymentConfig  `json:"deployment-configs,omitempty"`
	GatewayConfigs     *[]GatewayConfig     `json:"gateway-configs,omitempty"`
	IngressConfigs     *[]IngressConfig     `json:"ingress-configs,omitempty"`
	Namespace          *string              `json:"namespace,omitempty"`
	RouteConfigs       *[]RouteConfig       `json:"route-configs,omitempty"`
	ServiceConfigs     *[]ServiceConfig     `json:"service-configs,omitempty"`
	StatefulSetConfigs *[]StatefulSetConfig `json:"stateful-set-configs,omitempty"`
}

// Node defines model for Node.
type Node struct {
	// Id Unique identifier for the node.
	Id string `json:"id"`

	// Label Label for the node.
	Label string `json:"label"`

	// Type Type of the node
	Type NodeType `json:"type"`

	// Versions Node versions
	Versions *[]NodeVersion `json:"versions,omitempty"`
}

// NodeType Type of the node
type NodeType string

// NodeVersion defines model for NodeVersion.
type NodeVersion struct {
	FlowId     string  `json:"flowId"`
	ImageTag   *string `json:"imageTag,omitempty"`
	IsBaseline bool    `json:"isBaseline"`
}

// RouteConfig defines model for RouteConfig.
type RouteConfig struct {
	HttpRoute gateway.HTTPRoute `json:"httpRoute"`
}

// ServiceConfig defines model for ServiceConfig.
type ServiceConfig struct {
	Service corev1.Service `json:"service"`
}

// StatefulSetConfig defines model for StatefulSetConfig.
type StatefulSetConfig struct {
	StatefulSet appv1.StatefulSet `json:"stateful-set"`
}

// Template defines model for Template.
type Template struct {
	Description *string `json:"description,omitempty"`
	Name        string  `json:"name"`
	TemplateId  string  `json:"template-id"`
}

// TemplateConfig defines model for TemplateConfig.
type TemplateConfig struct {
	// Description The description of the template
	Description *string `json:"description,omitempty"`

	// Name The name to give the template
	Name    string           `json:"name"`
	Service []corev1.Service `json:"service"`
}

// TemplateSpec defines model for TemplateSpec.
type TemplateSpec struct {
	Arguments *map[string]interface{} `json:"arguments,omitempty"`

	// TemplateName name of the template
	TemplateName string `json:"template_name"`
}

// FlowId defines model for flow-id.
type FlowId = string

// TemplateName defines model for template-name.
type TemplateName = string

// Uuid defines model for uuid.
type Uuid = string

// Error defines model for Error.
type Error struct {
	// Error Error type
	Error string `json:"error"`

	// Msg Error message
	Msg *string `json:"msg,omitempty"`
}

// NotFound defines model for NotFound.
type NotFound struct {
	// Id Resource ID
	Id string `json:"id"`

	// ResourceType Resource type
	ResourceType string `json:"resource-type"`
}

// RequestError defines model for RequestError.
type RequestError struct {
	// Error Error type
	Error string `json:"error"`

	// Msg Error message
	Msg *string `json:"msg,omitempty"`
}

// PostTenantUuidFlowCreateJSONBody defines parameters for PostTenantUuidFlowCreate.
type PostTenantUuidFlowCreateJSONBody struct {
	FlowId       *string       `json:"flow-id,omitempty"`
	FlowSpec     FlowSpec      `json:"flow_spec"`
	TemplateSpec *TemplateSpec `json:"template_spec,omitempty"`
}

// PostTenantUuidDeployJSONRequestBody defines body for PostTenantUuidDeploy for application/json ContentType.
type PostTenantUuidDeployJSONRequestBody = MainClusterConfig

// PostTenantUuidFlowCreateJSONRequestBody defines body for PostTenantUuidFlowCreate for application/json ContentType.
type PostTenantUuidFlowCreateJSONRequestBody PostTenantUuidFlowCreateJSONBody

// PostTenantUuidTemplatesCreateJSONRequestBody defines body for PostTenantUuidTemplatesCreate for application/json ContentType.
type PostTenantUuidTemplatesCreateJSONRequestBody = TemplateConfig
