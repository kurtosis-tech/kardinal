// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version 2.1.0 DO NOT EDIT.
package server

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	. "github.com/kurtosis-tech/kardinal/libs/cli-kontrol-api/api/golang/types"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /health)
	GetHealth(ctx echo.Context) error

	// (POST /tenant/{uuid}/deploy)
	PostTenantUuidDeploy(ctx echo.Context, uuid Uuid) error

	// (POST /tenant/{uuid}/flow/create)
	PostTenantUuidFlowCreate(ctx echo.Context, uuid Uuid) error

	// (DELETE /tenant/{uuid}/flow/{flow-id})
	DeleteTenantUuidFlowFlowId(ctx echo.Context, uuid Uuid, flowId FlowId) error

	// (GET /tenant/{uuid}/flows)
	GetTenantUuidFlows(ctx echo.Context, uuid Uuid) error

	// (GET /tenant/{uuid}/templates)
	GetTenantUuidTemplates(ctx echo.Context, uuid Uuid) error

	// (POST /tenant/{uuid}/templates/create)
	PostTenantUuidTemplatesCreate(ctx echo.Context, uuid Uuid) error

	// (DELETE /tenant/{uuid}/templates/{template-name})
	DeleteTenantUuidTemplatesTemplateName(ctx echo.Context, uuid Uuid, templateName TemplateName) error

	// (GET /tenant/{uuid}/topology)
	GetTenantUuidTopology(ctx echo.Context, uuid Uuid) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetHealth converts echo context to params.
func (w *ServerInterfaceWrapper) GetHealth(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetHealth(ctx)
	return err
}

// PostTenantUuidDeploy converts echo context to params.
func (w *ServerInterfaceWrapper) PostTenantUuidDeploy(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "uuid" -------------
	var uuid Uuid

	err = runtime.BindStyledParameterWithOptions("simple", "uuid", ctx.Param("uuid"), &uuid, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter uuid: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostTenantUuidDeploy(ctx, uuid)
	return err
}

// PostTenantUuidFlowCreate converts echo context to params.
func (w *ServerInterfaceWrapper) PostTenantUuidFlowCreate(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "uuid" -------------
	var uuid Uuid

	err = runtime.BindStyledParameterWithOptions("simple", "uuid", ctx.Param("uuid"), &uuid, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter uuid: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostTenantUuidFlowCreate(ctx, uuid)
	return err
}

// DeleteTenantUuidFlowFlowId converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteTenantUuidFlowFlowId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "uuid" -------------
	var uuid Uuid

	err = runtime.BindStyledParameterWithOptions("simple", "uuid", ctx.Param("uuid"), &uuid, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter uuid: %s", err))
	}

	// ------------- Path parameter "flow-id" -------------
	var flowId FlowId

	err = runtime.BindStyledParameterWithOptions("simple", "flow-id", ctx.Param("flow-id"), &flowId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter flow-id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteTenantUuidFlowFlowId(ctx, uuid, flowId)
	return err
}

// GetTenantUuidFlows converts echo context to params.
func (w *ServerInterfaceWrapper) GetTenantUuidFlows(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "uuid" -------------
	var uuid Uuid

	err = runtime.BindStyledParameterWithOptions("simple", "uuid", ctx.Param("uuid"), &uuid, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter uuid: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTenantUuidFlows(ctx, uuid)
	return err
}

// GetTenantUuidTemplates converts echo context to params.
func (w *ServerInterfaceWrapper) GetTenantUuidTemplates(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "uuid" -------------
	var uuid Uuid

	err = runtime.BindStyledParameterWithOptions("simple", "uuid", ctx.Param("uuid"), &uuid, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter uuid: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTenantUuidTemplates(ctx, uuid)
	return err
}

// PostTenantUuidTemplatesCreate converts echo context to params.
func (w *ServerInterfaceWrapper) PostTenantUuidTemplatesCreate(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "uuid" -------------
	var uuid Uuid

	err = runtime.BindStyledParameterWithOptions("simple", "uuid", ctx.Param("uuid"), &uuid, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter uuid: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostTenantUuidTemplatesCreate(ctx, uuid)
	return err
}

// DeleteTenantUuidTemplatesTemplateName converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteTenantUuidTemplatesTemplateName(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "uuid" -------------
	var uuid Uuid

	err = runtime.BindStyledParameterWithOptions("simple", "uuid", ctx.Param("uuid"), &uuid, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter uuid: %s", err))
	}

	// ------------- Path parameter "template-name" -------------
	var templateName TemplateName

	err = runtime.BindStyledParameterWithOptions("simple", "template-name", ctx.Param("template-name"), &templateName, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter template-name: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteTenantUuidTemplatesTemplateName(ctx, uuid, templateName)
	return err
}

// GetTenantUuidTopology converts echo context to params.
func (w *ServerInterfaceWrapper) GetTenantUuidTopology(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "uuid" -------------
	var uuid Uuid

	err = runtime.BindStyledParameterWithOptions("simple", "uuid", ctx.Param("uuid"), &uuid, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter uuid: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTenantUuidTopology(ctx, uuid)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/health", wrapper.GetHealth)
	router.POST(baseURL+"/tenant/:uuid/deploy", wrapper.PostTenantUuidDeploy)
	router.POST(baseURL+"/tenant/:uuid/flow/create", wrapper.PostTenantUuidFlowCreate)
	router.DELETE(baseURL+"/tenant/:uuid/flow/:flow-id", wrapper.DeleteTenantUuidFlowFlowId)
	router.GET(baseURL+"/tenant/:uuid/flows", wrapper.GetTenantUuidFlows)
	router.GET(baseURL+"/tenant/:uuid/templates", wrapper.GetTenantUuidTemplates)
	router.POST(baseURL+"/tenant/:uuid/templates/create", wrapper.PostTenantUuidTemplatesCreate)
	router.DELETE(baseURL+"/tenant/:uuid/templates/:template-name", wrapper.DeleteTenantUuidTemplatesTemplateName)
	router.GET(baseURL+"/tenant/:uuid/topology", wrapper.GetTenantUuidTopology)

}

type ErrorJSONResponse struct {
	// Error Error type
	Error string `json:"error"`

	// Msg Error message
	Msg *string `json:"msg,omitempty"`
}

type NotFoundJSONResponse struct {
	// Id Resource ID
	Id string `json:"id"`

	// ResourceType Resource type
	ResourceType string `json:"resource-type"`
}

type GetHealthRequestObject struct {
}

type GetHealthResponseObject interface {
	VisitGetHealthResponse(w http.ResponseWriter) error
}

type GetHealth200JSONResponse string

func (response GetHealth200JSONResponse) VisitGetHealthResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostTenantUuidDeployRequestObject struct {
	Uuid Uuid `json:"uuid"`
	Body *PostTenantUuidDeployJSONRequestBody
}

type PostTenantUuidDeployResponseObject interface {
	VisitPostTenantUuidDeployResponse(w http.ResponseWriter) error
}

type PostTenantUuidDeploy200JSONResponse Flow

func (response PostTenantUuidDeploy200JSONResponse) VisitPostTenantUuidDeployResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostTenantUuidDeploy404JSONResponse struct{ NotFoundJSONResponse }

func (response PostTenantUuidDeploy404JSONResponse) VisitPostTenantUuidDeployResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type PostTenantUuidDeploy500JSONResponse struct{ ErrorJSONResponse }

func (response PostTenantUuidDeploy500JSONResponse) VisitPostTenantUuidDeployResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type PostTenantUuidFlowCreateRequestObject struct {
	Uuid Uuid `json:"uuid"`
	Body *PostTenantUuidFlowCreateJSONRequestBody
}

type PostTenantUuidFlowCreateResponseObject interface {
	VisitPostTenantUuidFlowCreateResponse(w http.ResponseWriter) error
}

type PostTenantUuidFlowCreate200JSONResponse Flow

func (response PostTenantUuidFlowCreate200JSONResponse) VisitPostTenantUuidFlowCreateResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostTenantUuidFlowCreate404JSONResponse struct{ NotFoundJSONResponse }

func (response PostTenantUuidFlowCreate404JSONResponse) VisitPostTenantUuidFlowCreateResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type PostTenantUuidFlowCreate500JSONResponse struct{ ErrorJSONResponse }

func (response PostTenantUuidFlowCreate500JSONResponse) VisitPostTenantUuidFlowCreateResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type DeleteTenantUuidFlowFlowIdRequestObject struct {
	Uuid   Uuid   `json:"uuid"`
	FlowId FlowId `json:"flow-id"`
}

type DeleteTenantUuidFlowFlowIdResponseObject interface {
	VisitDeleteTenantUuidFlowFlowIdResponse(w http.ResponseWriter) error
}

type DeleteTenantUuidFlowFlowId2xxResponse struct {
	StatusCode int
}

func (response DeleteTenantUuidFlowFlowId2xxResponse) VisitDeleteTenantUuidFlowFlowIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(response.StatusCode)
	return nil
}

type DeleteTenantUuidFlowFlowId404JSONResponse struct{ NotFoundJSONResponse }

func (response DeleteTenantUuidFlowFlowId404JSONResponse) VisitDeleteTenantUuidFlowFlowIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type DeleteTenantUuidFlowFlowId500JSONResponse struct{ ErrorJSONResponse }

func (response DeleteTenantUuidFlowFlowId500JSONResponse) VisitDeleteTenantUuidFlowFlowIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type GetTenantUuidFlowsRequestObject struct {
	Uuid Uuid `json:"uuid"`
}

type GetTenantUuidFlowsResponseObject interface {
	VisitGetTenantUuidFlowsResponse(w http.ResponseWriter) error
}

type GetTenantUuidFlows200JSONResponse []Flow

func (response GetTenantUuidFlows200JSONResponse) VisitGetTenantUuidFlowsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetTenantUuidFlows404JSONResponse struct{ NotFoundJSONResponse }

func (response GetTenantUuidFlows404JSONResponse) VisitGetTenantUuidFlowsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type GetTenantUuidFlows500JSONResponse struct{ ErrorJSONResponse }

func (response GetTenantUuidFlows500JSONResponse) VisitGetTenantUuidFlowsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type GetTenantUuidTemplatesRequestObject struct {
	Uuid Uuid `json:"uuid"`
}

type GetTenantUuidTemplatesResponseObject interface {
	VisitGetTenantUuidTemplatesResponse(w http.ResponseWriter) error
}

type GetTenantUuidTemplates200JSONResponse []Template

func (response GetTenantUuidTemplates200JSONResponse) VisitGetTenantUuidTemplatesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetTenantUuidTemplates404JSONResponse struct{ NotFoundJSONResponse }

func (response GetTenantUuidTemplates404JSONResponse) VisitGetTenantUuidTemplatesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type GetTenantUuidTemplates500JSONResponse struct{ ErrorJSONResponse }

func (response GetTenantUuidTemplates500JSONResponse) VisitGetTenantUuidTemplatesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type PostTenantUuidTemplatesCreateRequestObject struct {
	Uuid Uuid `json:"uuid"`
	Body *PostTenantUuidTemplatesCreateJSONRequestBody
}

type PostTenantUuidTemplatesCreateResponseObject interface {
	VisitPostTenantUuidTemplatesCreateResponse(w http.ResponseWriter) error
}

type PostTenantUuidTemplatesCreate200JSONResponse Template

func (response PostTenantUuidTemplatesCreate200JSONResponse) VisitPostTenantUuidTemplatesCreateResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostTenantUuidTemplatesCreate404JSONResponse struct{ NotFoundJSONResponse }

func (response PostTenantUuidTemplatesCreate404JSONResponse) VisitPostTenantUuidTemplatesCreateResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type PostTenantUuidTemplatesCreate500JSONResponse struct{ ErrorJSONResponse }

func (response PostTenantUuidTemplatesCreate500JSONResponse) VisitPostTenantUuidTemplatesCreateResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type DeleteTenantUuidTemplatesTemplateNameRequestObject struct {
	Uuid         Uuid         `json:"uuid"`
	TemplateName TemplateName `json:"template-name"`
}

type DeleteTenantUuidTemplatesTemplateNameResponseObject interface {
	VisitDeleteTenantUuidTemplatesTemplateNameResponse(w http.ResponseWriter) error
}

type DeleteTenantUuidTemplatesTemplateName2xxResponse struct {
	StatusCode int
}

func (response DeleteTenantUuidTemplatesTemplateName2xxResponse) VisitDeleteTenantUuidTemplatesTemplateNameResponse(w http.ResponseWriter) error {
	w.WriteHeader(response.StatusCode)
	return nil
}

type DeleteTenantUuidTemplatesTemplateName404JSONResponse struct{ NotFoundJSONResponse }

func (response DeleteTenantUuidTemplatesTemplateName404JSONResponse) VisitDeleteTenantUuidTemplatesTemplateNameResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type DeleteTenantUuidTemplatesTemplateName500JSONResponse struct{ ErrorJSONResponse }

func (response DeleteTenantUuidTemplatesTemplateName500JSONResponse) VisitDeleteTenantUuidTemplatesTemplateNameResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type GetTenantUuidTopologyRequestObject struct {
	Uuid Uuid `json:"uuid"`
}

type GetTenantUuidTopologyResponseObject interface {
	VisitGetTenantUuidTopologyResponse(w http.ResponseWriter) error
}

type GetTenantUuidTopology200JSONResponse ClusterTopology

func (response GetTenantUuidTopology200JSONResponse) VisitGetTenantUuidTopologyResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetTenantUuidTopology404JSONResponse struct{ NotFoundJSONResponse }

func (response GetTenantUuidTopology404JSONResponse) VisitGetTenantUuidTopologyResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type GetTenantUuidTopology500JSONResponse struct{ ErrorJSONResponse }

func (response GetTenantUuidTopology500JSONResponse) VisitGetTenantUuidTopologyResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (GET /health)
	GetHealth(ctx context.Context, request GetHealthRequestObject) (GetHealthResponseObject, error)

	// (POST /tenant/{uuid}/deploy)
	PostTenantUuidDeploy(ctx context.Context, request PostTenantUuidDeployRequestObject) (PostTenantUuidDeployResponseObject, error)

	// (POST /tenant/{uuid}/flow/create)
	PostTenantUuidFlowCreate(ctx context.Context, request PostTenantUuidFlowCreateRequestObject) (PostTenantUuidFlowCreateResponseObject, error)

	// (DELETE /tenant/{uuid}/flow/{flow-id})
	DeleteTenantUuidFlowFlowId(ctx context.Context, request DeleteTenantUuidFlowFlowIdRequestObject) (DeleteTenantUuidFlowFlowIdResponseObject, error)

	// (GET /tenant/{uuid}/flows)
	GetTenantUuidFlows(ctx context.Context, request GetTenantUuidFlowsRequestObject) (GetTenantUuidFlowsResponseObject, error)

	// (GET /tenant/{uuid}/templates)
	GetTenantUuidTemplates(ctx context.Context, request GetTenantUuidTemplatesRequestObject) (GetTenantUuidTemplatesResponseObject, error)

	// (POST /tenant/{uuid}/templates/create)
	PostTenantUuidTemplatesCreate(ctx context.Context, request PostTenantUuidTemplatesCreateRequestObject) (PostTenantUuidTemplatesCreateResponseObject, error)

	// (DELETE /tenant/{uuid}/templates/{template-name})
	DeleteTenantUuidTemplatesTemplateName(ctx context.Context, request DeleteTenantUuidTemplatesTemplateNameRequestObject) (DeleteTenantUuidTemplatesTemplateNameResponseObject, error)

	// (GET /tenant/{uuid}/topology)
	GetTenantUuidTopology(ctx context.Context, request GetTenantUuidTopologyRequestObject) (GetTenantUuidTopologyResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// GetHealth operation middleware
func (sh *strictHandler) GetHealth(ctx echo.Context) error {
	var request GetHealthRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetHealth(ctx.Request().Context(), request.(GetHealthRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetHealth")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetHealthResponseObject); ok {
		return validResponse.VisitGetHealthResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostTenantUuidDeploy operation middleware
func (sh *strictHandler) PostTenantUuidDeploy(ctx echo.Context, uuid Uuid) error {
	var request PostTenantUuidDeployRequestObject

	request.Uuid = uuid

	var body PostTenantUuidDeployJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostTenantUuidDeploy(ctx.Request().Context(), request.(PostTenantUuidDeployRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostTenantUuidDeploy")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostTenantUuidDeployResponseObject); ok {
		return validResponse.VisitPostTenantUuidDeployResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostTenantUuidFlowCreate operation middleware
func (sh *strictHandler) PostTenantUuidFlowCreate(ctx echo.Context, uuid Uuid) error {
	var request PostTenantUuidFlowCreateRequestObject

	request.Uuid = uuid

	var body PostTenantUuidFlowCreateJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostTenantUuidFlowCreate(ctx.Request().Context(), request.(PostTenantUuidFlowCreateRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostTenantUuidFlowCreate")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostTenantUuidFlowCreateResponseObject); ok {
		return validResponse.VisitPostTenantUuidFlowCreateResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// DeleteTenantUuidFlowFlowId operation middleware
func (sh *strictHandler) DeleteTenantUuidFlowFlowId(ctx echo.Context, uuid Uuid, flowId FlowId) error {
	var request DeleteTenantUuidFlowFlowIdRequestObject

	request.Uuid = uuid
	request.FlowId = flowId

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteTenantUuidFlowFlowId(ctx.Request().Context(), request.(DeleteTenantUuidFlowFlowIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteTenantUuidFlowFlowId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteTenantUuidFlowFlowIdResponseObject); ok {
		return validResponse.VisitDeleteTenantUuidFlowFlowIdResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetTenantUuidFlows operation middleware
func (sh *strictHandler) GetTenantUuidFlows(ctx echo.Context, uuid Uuid) error {
	var request GetTenantUuidFlowsRequestObject

	request.Uuid = uuid

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetTenantUuidFlows(ctx.Request().Context(), request.(GetTenantUuidFlowsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetTenantUuidFlows")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetTenantUuidFlowsResponseObject); ok {
		return validResponse.VisitGetTenantUuidFlowsResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetTenantUuidTemplates operation middleware
func (sh *strictHandler) GetTenantUuidTemplates(ctx echo.Context, uuid Uuid) error {
	var request GetTenantUuidTemplatesRequestObject

	request.Uuid = uuid

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetTenantUuidTemplates(ctx.Request().Context(), request.(GetTenantUuidTemplatesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetTenantUuidTemplates")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetTenantUuidTemplatesResponseObject); ok {
		return validResponse.VisitGetTenantUuidTemplatesResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostTenantUuidTemplatesCreate operation middleware
func (sh *strictHandler) PostTenantUuidTemplatesCreate(ctx echo.Context, uuid Uuid) error {
	var request PostTenantUuidTemplatesCreateRequestObject

	request.Uuid = uuid

	var body PostTenantUuidTemplatesCreateJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostTenantUuidTemplatesCreate(ctx.Request().Context(), request.(PostTenantUuidTemplatesCreateRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostTenantUuidTemplatesCreate")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostTenantUuidTemplatesCreateResponseObject); ok {
		return validResponse.VisitPostTenantUuidTemplatesCreateResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// DeleteTenantUuidTemplatesTemplateName operation middleware
func (sh *strictHandler) DeleteTenantUuidTemplatesTemplateName(ctx echo.Context, uuid Uuid, templateName TemplateName) error {
	var request DeleteTenantUuidTemplatesTemplateNameRequestObject

	request.Uuid = uuid
	request.TemplateName = templateName

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteTenantUuidTemplatesTemplateName(ctx.Request().Context(), request.(DeleteTenantUuidTemplatesTemplateNameRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteTenantUuidTemplatesTemplateName")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteTenantUuidTemplatesTemplateNameResponseObject); ok {
		return validResponse.VisitDeleteTenantUuidTemplatesTemplateNameResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetTenantUuidTopology operation middleware
func (sh *strictHandler) GetTenantUuidTopology(ctx echo.Context, uuid Uuid) error {
	var request GetTenantUuidTopologyRequestObject

	request.Uuid = uuid

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetTenantUuidTopology(ctx.Request().Context(), request.(GetTenantUuidTopologyRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetTenantUuidTopology")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetTenantUuidTopologyResponseObject); ok {
		return validResponse.VisitGetTenantUuidTopologyResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8xZW2/bNhT+KwS3R9lKtg4Y/LYl7WasK4LFeSqKgpWOZTYSyZKUE8PQfx9I6k7akRs0",
	"8YMBWTz37/BcoD1OeCE4A6YVXuyxIJIUoEHaf+ucP8xoah5TUImkQlPO8AK/y/kDoikwTdcUJI4wNa8F",
	"0RscYUYKwIuWO8ISvpVUQooXWpYQYZVsoCBGrN4JQ6q0pCzDVRVhDYXIiYaZkzLWbN4ivkZ6A6ghDasf",
	"CjrNiLIMeX13t7xudEtQvJTJAd2W/xSVlSFWgjMFNvJvpeTSPCScaWDaPBIhcpoQY0z8VRmL9j2JQnIB",
	"UlPHDw3/0AMrFlnl0diGCBcqO8RSgFIkC3BVfS8/1no/tWT8y1dItHMwINdo/cD1O16y9BnehsD6rwYI",
	"La9Dvjb4zdzJQe5wrEZeD4VFxp4pIWiVMK7R2sbAEDkvrWNXeak0yBUXPOfZLoBzmtUh0FDYh58lrPEC",
	"/xR3FzuuJcZv0wyM87VlREqyM/8ZT0+Q8oGnASmjkDiRUW2gH40IW2M8h3LyBXIfj/fmNVqb5N0AMkLn",
	"IVTrO+mxrzbQK1fNHW6jn7Yl5aBkTWQGeqpkRz1F8ihsbVGp9YUCZ4qvH7herfaMt2elzIcY+z4eA7Sr",
	"5p20Q9bdCkgGqkbXtSAZzHKeEO2KFDySQuRGzheS3ANLZ2RhSrfSQZBBbmnSNQifu6EgT8Z7aMpIdsi9",
	"8dX5l1BW39IrztY0891tZCb2fPo9u3V8tdgQPp519mpOKo93jH4rB6nbXC6TtMErMOlqHuQWRNaVfch+",
	"Y99bvuC9C5bn1U60F6vmBFYWBtGMaHgguw7LHqpbkMoIMDmQ0n7+dgprGuUrNcFF7XH0vRfJ3iFLEcqv",
	"IegelCmInO+KOpKPs4zPGmVCbC/n19151B3PaCG4tCz1gGKpceTGlgW+/13NKY+JoDERQsXbS9eJ6giO",
	"VCVcwvZyftvG94giRxvUZI6cpnERbAX33A0Fa9UMgIE49ZALINTUDh+6ZmoMltKRoX3iWuYxMw+DOjDW",
	"bzC9N4HZtw+N76EvzY7QmqOMbmGarF4etEn/QxPiiSvUJciTMW9a0TDiRGZl0ew8JE2pCQ/Jb3pEbmL3",
	"638t9/NJ68m0NPp8wBtDT9ma24Sk2ra6q/fL+B/OtOQ5+uNmidvChRf4cn4xvzDGcgGMCIoX+Ff7ygXd",
	"Oh1vgOQGAG+14xK5M5RsILlHidOCI1wPQCZGdi5fpniB/wL9txM1WmF+ubg4aaQPrERDy27LJAGl1mWO",
	"GkWGrIpwrIERpuO92buq2FUNizpXAZNvuNIry3FX0tSVTBubbvf9GG7NHUlsV7zqk0MSlP6Tp7uT/D3W",
	"+v2pIhAPZzgiSEieIs7yHUocj7d6Vs/E5pitdhwNmrdFZlREiQQrFylNdKlMYr65eHNIbGtn3C6GVYR/",
	"cwYfZ6hXylBKGENiawhMzQvj15XjePHc8Ef7z6ouZE9BYQtev1BNYRxUy9Dk76RMWWpdyBBBaYN/qSjL",
	"UGNbhLhw1TbfoQeqN4iggfpzyN13Z5W3+3r3qly1zsEl8TB5r+37Yfqa3zL9zvSNnqRrVkKX6X2IHh/9",
	"xtLWA+vBGcTVmnqopw0DqZ5TAp6Ru5P2RJfE3sh03gW5KU8TQVi15GcMRLuOTADjPVXaDoqNY+0GbcZy",
	"hly0XheaExtmi9Frdc0p2Bwep9rGxeChP7u/XDfq8se3rjk7r8sb7wdfWU5qUG26NA8f3HeaH9Orhl+D",
	"pnWsNubn0LF07xvAhHrZUL9OuTyW5eNvGqFkr8+Q2XplYbW8RNCr6v8AAAD//3+qJLUQHQAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
