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

	"H4sIAAAAAAAC/+xYXW/bNhf+KwTf91K2nK0DBt1tSbsZ64pgSa6KXjDkscxGIlmScmIY+u8DPyRbEu04",
	"DbDlYhcFFPJ8POc5HzzuDlNZKylAWIOLHVZEkxosaP/XqpKPM87cJwNDNVeWS4EL/KGSj4gzEJavOGic",
	"Ye6OFbFrnGFBasBFr51hDd8aroHhwuoGMmzoGmrizNqtcqLGai5K3LYZbpqUw7u75RWSK2TXgDQY2WgK",
	"aa9e/yUuWydslBQGfNDvtZbafVApLAjrPolSFafEgcm/Godod2BRaalAWx70odMfRuDNIu88G2PIcG3K",
	"Yyo1GEPKhFZ7GOXn6PdLLybvvwK1IcCEXef1k7QfZCPYK6JNJeuvmCC0vErF2uVvFm6Oaqe5GkU9NJY5",
	"POdQ0DsR0qKV58AJhSh9YJdVYyzoW6lkJcttIs+sjBRYqP3H/zWscIH/l+97Ko8W8/esBBd8REa0Jlv3",
	"t5DsBVY+SZawMqIkmMwiwCkbGfZgJgFV5B6qaT4+umO0csW7BuSMzlNZjT05Ub9dw8Gk6Hq4Z59Bd3TU",
	"siW6BHuu5SB9juURbf1Qif5SxLm5NyXuYExOwPu7RlfDHE9jPJXQ/SDdWzuG7kYBHbgatWtNSphVkhIb",
	"hhQ8kVpVzs49oQ8g2IwUFbFgbDLJoDecwizM2pR2J0Ge5XsIZWQ7Fd64df4kXMQuvZRixctpuJ1N6u/P",
	"77OboBfNpvIzQedb86zxeCf4t2ZQul1zuaJNtsBZrXlUWxEdJ/tQ/dqfe71k3yXH8+1W9Y0VNUE0tcto",
	"SSw8ku0+lwdZ3YA2zoCrAcYP63fvMMqYqVNHLuqvs+9tJN9DXiJVX8OkT1LJQFVyW0cmn2alnHXOlNpc",
	"zK/299n+esZrJbVXiQuKl8ZZWFsK/PCzmXOZE8VzopTJNxfhJYoMjlxRqWFzMb/p+T3hKMgmPbmr4Gk8",
	"BHvDB+GmH1QuVtKngFs/AS4/LvM/pLBaVuiX6yXu84kLfDFfzBeOY6lAEMVxgX/0RwGc5zdfA6kc0Mmy",
	"KTUKd4iugT4gGrzgDMd3waXJrytLhgv8G9jfg6nRZvfDYvGiTSexKQ6R3TSUgjGrpkKdIyfWZji3IIiw",
	"+c6to20eyPQ1JU0C8rU09tZr3DWchUry3Oy38c/pibUXyf3m234JCQVjf5Vs+6J4T03E6bBN8BGAI4KU",
	"lgxJUW0RDTqTjbx9ZW5OYfWvdBLeBrkXFFEN3i4yltjGuMJ8t3h3zGyPM+/35TbDPwXApxXipp0qCQck",
	"90Dg3LpwcV0GjTdVG/3ekeA84EUEsUj+f5WQrIRdXPLaMP8qCGUxLIcrfz4sCPdvyb6zILJn5brdM9TO",
	"YZqenqajuufVR/AGePVQj70SQyLNa5rqFfV71kIaCnmy3rztwrYHv52fz0H/S/vfScMp9sf/F5AgvrtD",
	"bi3StffyT5Detn8HAAD///evY0/DEwAA",
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
