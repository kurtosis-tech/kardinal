// Package types provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version 2.1.0 DO NOT EDIT.
package types

import (
	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
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
	ImageLocator string `json:"image-locator"`
	ServiceName  string `json:"service-name"`
}

// IngressConfig defines model for IngressConfig.
type IngressConfig struct {
	Ingress networkingv1.Ingress `json:"ingress"`
}

// MainClusterConfig defines model for MainClusterConfig.
type MainClusterConfig struct {
	IngressConfigs *[]IngressConfig `json:"ingress-configs,omitempty"`
	Namespace      *string          `json:"namespace,omitempty"`
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

// FlowId defines model for flow-id.
type FlowId = string

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
