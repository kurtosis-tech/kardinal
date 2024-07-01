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
	. "github.com/kurtosis-tech/kardinal/libs/manager-kontrol-api/api/golang/types"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Cluster resource definition
	// (GET /cluster-resources)
	GetClusterResources(ctx echo.Context, params GetClusterResourcesParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetClusterResources converts echo context to params.
func (w *ServerInterfaceWrapper) GetClusterResources(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetClusterResourcesParams
	// ------------- Optional query parameter "namespace" -------------

	err = runtime.BindQueryParameter("form", true, false, "namespace", ctx.QueryParams(), &params.Namespace)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter namespace: %s", err))
	}

	// ------------- Required query parameter "tenant" -------------

	err = runtime.BindQueryParameter("form", true, true, "tenant", ctx.QueryParams(), &params.Tenant)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tenant: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetClusterResources(ctx, params)
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

	router.GET(baseURL+"/cluster-resources", wrapper.GetClusterResources)

}

type NotOkJSONResponse ResponseInfo

type GetClusterResourcesRequestObject struct {
	Params GetClusterResourcesParams
}

type GetClusterResourcesResponseObject interface {
	VisitGetClusterResourcesResponse(w http.ResponseWriter) error
}

type GetClusterResources200JSONResponse ClusterResources

func (response GetClusterResources200JSONResponse) VisitGetClusterResourcesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetClusterResourcesdefaultJSONResponse struct {
	Body       ResponseInfo
	StatusCode int
}

func (response GetClusterResourcesdefaultJSONResponse) VisitGetClusterResourcesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Cluster resource definition
	// (GET /cluster-resources)
	GetClusterResources(ctx context.Context, request GetClusterResourcesRequestObject) (GetClusterResourcesResponseObject, error)
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

// GetClusterResources operation middleware
func (sh *strictHandler) GetClusterResources(ctx echo.Context, params GetClusterResourcesParams) error {
	var request GetClusterResourcesRequestObject

	request.Params = params

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetClusterResources(ctx.Request().Context(), request.(GetClusterResourcesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetClusterResources")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetClusterResourcesResponseObject); ok {
		return validResponse.VisitGetClusterResourcesResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7yVTW/bPAzHv4rB5zk6cdJeBt+KdSuCYemQttuhKApNoR01tqRStNug8HcfZDlO2rys",
	"h22n2CZF/khGf76ANKU1GjU7SF/AChIlMlL7xqiFZv80RydJWVZGQwo3N5PzyGQRLzDqfGJQ3vJYIa0g",
	"Bi1KhBR6I+FjpQjnkDJVGIOTCyxFm2NlvadjUjqHpmm8s7NGO2wZpoYvl/5BGs0YaIS1hZLC0yQPziO9",
	"bEX8nzCDFP5LNqUlweqSWRd6ojMTkr2pTOOzRck4j5DIEHiX7rCP/bGoHCPN0JmKZCC0ZCwSq/A2R1uY",
	"VbluqGIs24fnQW4GXbHCWlePh+e9K8Qb+0CV1lBbZ9fF4A4xWMELSGH5wQ2VSYRViTcl9bjFXAcnEisI",
	"lbHSbZfuqSrwME89FoVdiNPh+ebIrCrwONb61AZMOVbGo8lCoeZBbhK7zD2oSzTyk6Gl0nnSH9xHnQvG",
	"J7E6iHjR2f8BmkOqlTzSN2kI6/HwKvgdRwq+e4foTYeGWCviShT3v2XpG/Q9nHgX1N8Y4eaD+fmAkn0V",
	"r67dzpWRZo7+NzNUCoYUKqX59AT6QEoz5kg+UonOiRz3KMfa+30CcO19g9islek2BNjkiAPZ3ZGCrruU",
	"qKvSR/g0m13OIIbJ9PMlxPDjbDadTC+2QmzrnOq6wYoLb/sqtMiRki9GM5kiOvs2gRhqJBfEaTwcDUc+",
	"u7GohVWQwmn7KUyv7WUig0INaFuicmyn7pveXu7JHFK4QN6Rs/jVBrh9K/zXC4z8X8dZITHKDEVPCyUX",
	"EZuIkElhje1S6CAi2gq8bz/0seDYSoj3j3RDmnR7prl7szxORqM/tjp2erVnfVxVUqJzWVVEa46gxJmo",
	"Cj6UoUdOwrJrtacqS0ErSNcrp+9lNMdMadVmjIFF7gcFu2O/a5qm+RUAAP//EDi0luAHAAA=",
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
