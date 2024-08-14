// Package types provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version 2.1.0 DO NOT EDIT.
package types

import (
	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

// Defines values for NodeType.
const (
	Gateway        NodeType = "gateway"
	Redis          NodeType = "redis"
	Service        NodeType = "service"
	ServiceVersion NodeType = "service-version"
)

// ClusterTopology defines model for ClusterTopology.
type ClusterTopology struct {
	Edges []Edge `json:"edges"`
	Nodes []Node `json:"nodes"`
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
	FlowId   string   `json:"flow-id"`
	FlowUrls []string `json:"flow-urls"`
}

// FlowSpec defines model for FlowSpec.
type FlowSpec = []struct {
	ImageLocator string        `json:"image-locator"`
	ServiceName  string        `json:"service-name"`
	TemplateSpec *TemplateSpec `json:"template-spec,omitempty"`
}

// MainClusterConfig defines model for MainClusterConfig.
type MainClusterConfig struct {
	ServiceConfigs *[]ServiceConfig `json:"service-configs,omitempty"`
}

// Node defines model for Node.
type Node struct {
	// Id Unique identifier for the node.
	Id string `json:"id"`

	// Label Label for the node.
	Label *string `json:"label,omitempty"`

	// Parent Parent node
	Parent *string `json:"parent,omitempty"`

	// Type Type of the node
	Type NodeType `json:"type"`

	// Versions Node versions
	Versions *[]string `json:"versions,omitempty"`
}

// NodeType Type of the node
type NodeType string

// ServiceConfig defines model for ServiceConfig.
type ServiceConfig struct {
	Deployment appv1.Deployment `json:"deployment"`
	Service    corev1.Service   `json:"service"`
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

// PostTenantUuidDeployJSONRequestBody defines body for PostTenantUuidDeploy for application/json ContentType.
type PostTenantUuidDeployJSONRequestBody = MainClusterConfig

// PostTenantUuidFlowCreateJSONRequestBody defines body for PostTenantUuidFlowCreate for application/json ContentType.
type PostTenantUuidFlowCreateJSONRequestBody = FlowSpec

// PostTenantUuidTemplatesCreateJSONRequestBody defines body for PostTenantUuidTemplatesCreate for application/json ContentType.
type PostTenantUuidTemplatesCreateJSONRequestBody = TemplateConfig
