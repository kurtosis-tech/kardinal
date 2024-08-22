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
	"io"
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
	// Cluster resource definition in a manifest YAML response
	// (GET /tenant/{uuid}/cluster-resources/manifest)
	GetTenantUuidClusterResourcesManifest(ctx echo.Context, uuid Uuid) error
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

// GetTenantUuidClusterResourcesManifest converts echo context to params.
func (w *ServerInterfaceWrapper) GetTenantUuidClusterResourcesManifest(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "uuid" -------------
	var uuid Uuid

	err = runtime.BindStyledParameterWithOptions("simple", "uuid", ctx.Param("uuid"), &uuid, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter uuid: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTenantUuidClusterResourcesManifest(ctx, uuid)
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
	router.GET(baseURL+"/tenant/:uuid/cluster-resources/manifest", wrapper.GetTenantUuidClusterResourcesManifest)

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

type GetTenantUuidClusterResourcesManifestRequestObject struct {
	Uuid Uuid `json:"uuid"`
}

type GetTenantUuidClusterResourcesManifestResponseObject interface {
	VisitGetTenantUuidClusterResourcesManifestResponse(w http.ResponseWriter) error
}

type GetTenantUuidClusterResourcesManifest200ApplicationxYamlResponse struct {
	Body          io.Reader
	ContentLength int64
}

func (response GetTenantUuidClusterResourcesManifest200ApplicationxYamlResponse) VisitGetTenantUuidClusterResourcesManifestResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/x-yaml")
	if response.ContentLength != 0 {
		w.Header().Set("Content-Length", fmt.Sprint(response.ContentLength))
	}
	w.WriteHeader(200)

	if closer, ok := response.Body.(io.ReadCloser); ok {
		defer closer.Close()
	}
	_, err := io.Copy(w, response.Body)
	return err
}

type GetTenantUuidClusterResourcesManifestdefaultJSONResponse struct {
	Body       ResponseInfo
	StatusCode int
}

func (response GetTenantUuidClusterResourcesManifestdefaultJSONResponse) VisitGetTenantUuidClusterResourcesManifestResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Cluster resource definition
	// (GET /tenant/{uuid}/cluster-resources)
	GetTenantUuidClusterResources(ctx context.Context, request GetTenantUuidClusterResourcesRequestObject) (GetTenantUuidClusterResourcesResponseObject, error)
	// Cluster resource definition in a manifest YAML response
	// (GET /tenant/{uuid}/cluster-resources/manifest)
	GetTenantUuidClusterResourcesManifest(ctx context.Context, request GetTenantUuidClusterResourcesManifestRequestObject) (GetTenantUuidClusterResourcesManifestResponseObject, error)
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

// GetTenantUuidClusterResourcesManifest operation middleware
func (sh *strictHandler) GetTenantUuidClusterResourcesManifest(ctx echo.Context, uuid Uuid) error {
	var request GetTenantUuidClusterResourcesManifestRequestObject

	request.Uuid = uuid

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetTenantUuidClusterResourcesManifest(ctx.Request().Context(), request.(GetTenantUuidClusterResourcesManifestRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetTenantUuidClusterResourcesManifest")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetTenantUuidClusterResourcesManifestResponseObject); ok {
		return validResponse.VisitGetTenantUuidClusterResourcesManifestResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8RW32/bNhD+VwRuj7IZNy+D34KlDYwuSeEmG4YiCBj5JF9Dkdzx6EYL/L8PFCXLaWwn",
	"A7bmST/u13cfT/fpURS2dtaAYS+mj8IpUjUwUPsUAi7idQG+IHSM1oipuL6enWa2zHgJGYG3gQoQucBo",
	"c4qXIhdG1SCmKT4XBH8FJFiIKVOAXPhiCbWKiblx0c8zoanEer2Ozt5Z46EFcGH58j7eFNYwGI63yjmN",
	"hYpg5FcfET1uZfyZoBRT8ZMc+pLJ6uW8Sz0zpU3FvmvMwIODgmGRAZElEV264Jj7Vx08A827nhNhZB0Q",
	"Y3pSgZeW8O8W3a2zGovOggx1e/Mwquyo63s1uQNWk/HJdtinGNWIfPAcYe0stc13xHaBIk+ETwV6RjtG",
	"KwuNYHhUWenuK6kceumhCITcyD4qttUhUESqES0TTtum7udgJ1zlnF9Nxqcb18Mgk/uA8f4XHxEqhzKa",
	"5GofEs9oEoMU9EH6lHZLdTw+HULmQcNL3KWo15FngL9ZukdTyU3gLtRgVra5LVH3385hxO+j+4fW+y3Q",
	"Vorhm2r2wjvr7D8AmgdaYXHglAtLsJqMPye/w5CS786Ri6Z9I7dC4qD07YtYNgT9niJeBer/OMLhhb37",
	"CgXHLp6st2erqbALiNfSUq04Lmc0fPxObBKhYaiAYqYavFcV7NjQvffrFu1V9E1LvVeALynBUCNPyG4O",
	"NHTVlQQT6pjh/Xx+ORe5mF18uBS5+ONkfjG7ONtKsa0n2LHByDrazpVRFZD8aA2T1dnJp5nIxQrIJxGY",
	"jI/GR7G6dWCUQzEVx+2rdHotl5LBKMPyMQrcWhZJF0a0LQwVtDMQj6BdTLOFmIoz4Ks29Drg4pmc5E/k",
	"98tujgcX2crr+uY7zXx3dPSfKeYziDtU83MoCvC+DDrrcaQ1XqqgeV+FDWSZNL5dBaGuFTVi2ivt5vci",
	"W0CJBtuKuWBVRX7Ec95vYpqXTkfWymAJnreO6WlLV0v0GZiFs2g4I+BAxmdK6/af52O4AzLA4Ad8paXW",
	"1hXL2DqrbdVkaDJrIKuDZtxAyHoEA2P5v5mU876BHzExD6NG1frpzOz4zt50KCLNamD1z5Pz37ap3T8w",
	"udgMw816vV7/EwAA//8hLsi2EAsAAA==",
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
