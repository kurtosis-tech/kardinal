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

	// (POST /tenant/{uuid}/flow/{flow-id}/delete)
	PostTenantUuidFlowFlowIdDelete(ctx echo.Context, uuid Uuid, flowId FlowId) error

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

// PostTenantUuidFlowFlowIdDelete converts echo context to params.
func (w *ServerInterfaceWrapper) PostTenantUuidFlowFlowIdDelete(ctx echo.Context) error {
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
	err = w.Handler.PostTenantUuidFlowFlowIdDelete(ctx, uuid, flowId)
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
	router.POST(baseURL+"/tenant/:uuid/flow/:flow-id/delete", wrapper.PostTenantUuidFlowFlowIdDelete)
	router.GET(baseURL+"/tenant/:uuid/flows", wrapper.GetTenantUuidFlows)
	router.GET(baseURL+"/tenant/:uuid/topology", wrapper.GetTenantUuidTopology)

}

type NotFoundJSONResponse struct {
	// Id Resource ID
	Id *string `json:"id,omitempty"`

	// ResourceType Resource type
	ResourceType *string `json:"resource-type,omitempty"`
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

type PostTenantUuidDeploy200JSONResponse DevFlow

func (response PostTenantUuidDeploy200JSONResponse) VisitPostTenantUuidDeployResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostTenantUuidFlowCreateRequestObject struct {
	Uuid Uuid `json:"uuid"`
	Body *PostTenantUuidFlowCreateJSONRequestBody
}

type PostTenantUuidFlowCreateResponseObject interface {
	VisitPostTenantUuidFlowCreateResponse(w http.ResponseWriter) error
}

type PostTenantUuidFlowCreate200JSONResponse DevFlow

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

type PostTenantUuidFlowFlowIdDeleteRequestObject struct {
	Uuid   Uuid   `json:"uuid"`
	FlowId FlowId `json:"flow-id"`
	Body   *PostTenantUuidFlowFlowIdDeleteJSONRequestBody
}

type PostTenantUuidFlowFlowIdDeleteResponseObject interface {
	VisitPostTenantUuidFlowFlowIdDeleteResponse(w http.ResponseWriter) error
}

type PostTenantUuidFlowFlowIdDelete2xxResponse struct {
	StatusCode int
}

func (response PostTenantUuidFlowFlowIdDelete2xxResponse) VisitPostTenantUuidFlowFlowIdDeleteResponse(w http.ResponseWriter) error {
	w.WriteHeader(response.StatusCode)
	return nil
}

type PostTenantUuidFlowFlowIdDelete404JSONResponse struct{ NotFoundJSONResponse }

func (response PostTenantUuidFlowFlowIdDelete404JSONResponse) VisitPostTenantUuidFlowFlowIdDeleteResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type GetTenantUuidFlowsRequestObject struct {
	Uuid Uuid `json:"uuid"`
}

type GetTenantUuidFlowsResponseObject interface {
	VisitGetTenantUuidFlowsResponse(w http.ResponseWriter) error
}

type GetTenantUuidFlows200JSONResponse []DevFlow

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

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (GET /health)
	GetHealth(ctx context.Context, request GetHealthRequestObject) (GetHealthResponseObject, error)

	// (POST /tenant/{uuid}/deploy)
	PostTenantUuidDeploy(ctx context.Context, request PostTenantUuidDeployRequestObject) (PostTenantUuidDeployResponseObject, error)

	// (POST /tenant/{uuid}/flow/create)
	PostTenantUuidFlowCreate(ctx context.Context, request PostTenantUuidFlowCreateRequestObject) (PostTenantUuidFlowCreateResponseObject, error)

	// (POST /tenant/{uuid}/flow/{flow-id}/delete)
	PostTenantUuidFlowFlowIdDelete(ctx context.Context, request PostTenantUuidFlowFlowIdDeleteRequestObject) (PostTenantUuidFlowFlowIdDeleteResponseObject, error)

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

// PostTenantUuidFlowFlowIdDelete operation middleware
func (sh *strictHandler) PostTenantUuidFlowFlowIdDelete(ctx echo.Context, uuid Uuid, flowId FlowId) error {
	var request PostTenantUuidFlowFlowIdDeleteRequestObject

	request.Uuid = uuid
	request.FlowId = flowId

	var body PostTenantUuidFlowFlowIdDeleteJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostTenantUuidFlowFlowIdDelete(ctx.Request().Context(), request.(PostTenantUuidFlowFlowIdDeleteRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostTenantUuidFlowFlowIdDelete")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostTenantUuidFlowFlowIdDeleteResponseObject); ok {
		return validResponse.VisitPostTenantUuidFlowFlowIdDeleteResponse(ctx.Response())
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

	"H4sIAAAAAAAC/8xXW2/bNhT+KwS3hw2QrGTrw6C3LV43Y0URLMlT0QeGOrbZSDwsSTkxDP33gaQutkUl",
	"doOufTAgi+fyne9ceLSjHCuFEqQ1NN9RxTSrwIL2/5YlPqaicI8FGK6FsgIlzenbEh+JKEBasRSgaUKF",
	"e62YXdOESlYBzXvthGr4XAsNBc2triGhhq+hYs6s3SonaqwWckWbJqF1HXN4d7eYE1wSuwaiwWCtOcS9",
	"ev1zXDZO2CiUBnzQ79G+xVp6EBylBWndI1OqFJw5PNkn40Dt9owqjQq0FcFELIJ/W9RkMafJMQgPwR+n",
	"4WRS2x8nEd7aN3j/CbgNYU2YkGjJ0kfohEIMHvZVWRsL+hYVlrjajuOCYtUGaKHyDz9qWNKc/pANZZS1",
	"FrM/ixXQARnTmm3df4nFGVbeYxGx0uxn+ENrMmkBfhyxkdA5bFzRjmMqYJPu1fmLxPambhTwSOortoK0",
	"RM4sas/ZE6tU6SzcM/4AskhZXjILxsbKwIDeCA4pR7kUq9NZugl6V14tRnpnOHRJDFYnwU4pr4T67I7C",
	"L9k9lOPyfedekyVq38AuS7No9KGvR+q3a9ibNt0c6Mu5gO7VpGXL9ArsqZaD9CmWj+qwH0ytv1gl+oI+",
	"aWTcSfG5PsDXMeiQReM8if9JbcV0O+0O1a/9e68XJTc6sm63qmev1QRZV46mFbPwyLa0L8vhKd2ANs6A",
	"I7YQ+808OGxlzNipI5f0x8nQPBOgJ+aJv0K8RCyD1xqL6QnwlVo41oOHKpHRpkrcVm1Kn9IVpp1NpTaX",
	"s/lwngzHqagUaq/SXqlemibhos3pw29mJjBjSmRMKZNtLsNF0qbyyBVHDZvL2U2f6GccBdmoJ3cUPB23",
	"XG94L9yP0ftQyCX6WhDWT76rd4vsH5RWY0l+v17QvrBoTi9nF7MLxzEqkEwJmtNf/asAzvObrYGVDuho",
	"PUJNwhnha+APhAcvNKHtFHJp8rvEoqA5/Qvs38HU0S7yy8XFWWtIZLc5RHZTcw7GLOuSdI6cWJPQzIJk",
	"0mY7t0A1WSDT1xSaCORrNPbWa9zVogiV5LkZ9scP8XofRDK/qzUfQ0LB2D+w2J4V73P9dNClESoCZsKI",
	"0lgQlOWW8LAAjdbH5pVpeQ5mt5hEEW6IW00I1+BNE2OZrU08Y04y85Jwatqc46ug8V2lbn/DivASIBNG",
	"ipag7z9hCX1z8WbKcg816789pjO8a5dV16IlnJdr91sU86D3ZRlPXpTrlulv19cuvL42yE8aNqAtcXsu",
	"sTi0+88vl83T03i490n2/H+VJHvnUzfFYUrNazr3FR1y0krTt8pomfm/W8fufdO+TGz/BfxtuH2O0uNv",
	"9AiV3Rlx+46uvJcvo7Fp/gsAAP//MYhfRCASAAA=",
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
