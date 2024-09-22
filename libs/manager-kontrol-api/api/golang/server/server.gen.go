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
	// (GET /tenant/{uuid}/cluster-resources)
	GetTenantUuidClusterResources(ctx echo.Context, uuid Uuid) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetTenantUuidClusterResources converts echo context to params.
func (w *ServerInterfaceWrapper) GetTenantUuidClusterResources(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "uuid" -------------
	var uuid Uuid

	err = runtime.BindStyledParameterWithOptions("simple", "uuid", ctx.Param("uuid"), &uuid, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter uuid: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTenantUuidClusterResources(ctx, uuid)
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

	router.GET(baseURL+"/tenant/:uuid/cluster-resources", wrapper.GetTenantUuidClusterResources)

}

type NotOkJSONResponse ResponseInfo

type GetTenantUuidClusterResourcesRequestObject struct {
	Uuid Uuid `json:"uuid"`
}

type GetTenantUuidClusterResourcesResponseObject interface {
	VisitGetTenantUuidClusterResourcesResponse(w http.ResponseWriter) error
}

type GetTenantUuidClusterResources200JSONResponse ClusterResources

func (response GetTenantUuidClusterResources200JSONResponse) VisitGetTenantUuidClusterResourcesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetTenantUuidClusterResourcesdefaultJSONResponse struct {
	Body       ResponseInfo
	StatusCode int
}

func (response GetTenantUuidClusterResourcesdefaultJSONResponse) VisitGetTenantUuidClusterResourcesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Cluster resource definition
	// (GET /tenant/{uuid}/cluster-resources)
	GetTenantUuidClusterResources(ctx context.Context, request GetTenantUuidClusterResourcesRequestObject) (GetTenantUuidClusterResourcesResponseObject, error)
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

// GetTenantUuidClusterResources operation middleware
func (sh *strictHandler) GetTenantUuidClusterResources(ctx echo.Context, uuid Uuid) error {
	var request GetTenantUuidClusterResourcesRequestObject

	request.Uuid = uuid

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetTenantUuidClusterResources(ctx.Request().Context(), request.(GetTenantUuidClusterResourcesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetTenantUuidClusterResources")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetTenantUuidClusterResourcesResponseObject); ok {
		return validResponse.VisitGetTenantUuidClusterResourcesResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8SWS3PjNgzHv4oH7VEW482lo1um2U09nSYZr9MedjIZrgzL2EgkC4LedTP67h2K8iMb",
	"W8mhj5Me+AP8AaQEPEFpG2cNGvFQPIHTrBsU5O4pBFrE6wJ9yeSErIEC7u6mlyO7HMkKR4zeBi4RMqBo",
	"c1pWkIHRDUKR/DNg/DMQ4wIK4YAZ+HKFjY6BZeOizguTqaBt2yj2zhqPHcC1lZvHeFNaI2gk3mrnaip1",
	"hFFffCR6Ooj4I+MSCvhB7fNSyerVrA89NUubFvsuMYPfHJaCixEyW4Yo6Z1j7J/r4AV51uecCsbWIQul",
	"Jx1kZZn+6ugenK2p7C0k2HQ338aVHfd5ryefUfQkvzh0u41eG8j2yjE1znKXfF/Y3hGyVPACyAvZnKwq",
	"a0Ij48oq91gp7cgrj2Vgko3aesW0egLNrDfQVcLVdtNsz8FRXO2cX0/yy510GDLJ94yPP/lIqB2paFLr",
	"UyReyKQKcqgHy6drt9Ln+eXeZRZqfK12yettxTMoXy0/kqnUzvEYNZq13Twsqd5+O8PE76P8Q6f+P2gr",
	"LfhVb06D9oL8Kl2HGaudqEf0VPm83+7eOE7bTie3fSXiHtgGwdepfpnPb2dR+h9wkakYvR+gMij5NKmG",
	"eQzK0a/h2a4dhfDIayoHGErLuJ7kH5NuGCNpj5JE0ymGNbEEXT+8yrI75L8njzdB/RvnfP/Cfv6CpcQs",
	"nvWAF//v0i4wXpeWGy2xg5GR83ewC0RGsEKOkRr0Xld4pI1t1W/rRvOoTZ1v2yY/pQD7NbJEdj+Q0Lxf",
	"Ek1oYoT3s9nNDDKYXn+4gQz+uJhdT6+vDkIcNl3qqyEkdbT9po2ukNWv1gjbenRxO4UM1sg+dcpJfpaf",
	"xdWtQ6MdQQHn3au0e10tlaDRRtRTnAJaVabmOebD7llhdwbiFnR/7+kCCrhCmXeud4EWL3pu9mxG+XS8",
	"xnuJ6maQ9v67weLd2dk/Nla8QDwyWnwMZYneL0M92nKkXrfUoZZTK+yQVRqEunkkNI3mDRTbcWQ3g40W",
	"uCRD3YoZiK5ifeBl3e/btm3/DgAA//9xh5wN+QkAAA==",
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
