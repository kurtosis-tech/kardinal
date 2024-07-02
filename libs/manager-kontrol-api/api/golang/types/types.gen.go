// Package types provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version 2.1.0 DO NOT EDIT.
package types

import (
	v1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

// Defines values for ResponseType.
const (
	ERROR   ResponseType = "ERROR"
	INFO    ResponseType = "INFO"
	WARNING ResponseType = "WARNING"
)

// ClusterResources defines model for ClusterResources.
type ClusterResources struct {
	Deployments      *[]appsv1.Deployment        `json:"deployments,omitempty"`
	DestinationRules *[]v1alpha3.DestinationRule `json:"destination_rules,omitempty"`
	Gateway          *v1alpha3.Gateway           `json:"gateway,omitempty"`
	Services         *[]corev1.Service           `json:"services,omitempty"`
	VirtualServices  *[]v1alpha3.VirtualService  `json:"virtual_services,omitempty"`
}

// ResponseInfo defines model for ResponseInfo.
type ResponseInfo struct {
	Code    uint32       `json:"code"`
	Message string       `json:"message"`
	Type    ResponseType `json:"type"`
}

// ResponseType defines model for ResponseType.
type ResponseType string

// Uuid defines model for uuid.
type Uuid = string

// NotOk defines model for NotOk.
type NotOk = ResponseInfo

// GetTenantUuidClusterResourcesParams defines parameters for GetTenantUuidClusterResources.
type GetTenantUuidClusterResourcesParams struct {
	// Namespace The namespace for which to retrieve the cluster resources
	Namespace *string `form:"namespace,omitempty" json:"namespace,omitempty"`
}
