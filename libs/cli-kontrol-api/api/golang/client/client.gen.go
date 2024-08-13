// Package client provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version 2.1.0 DO NOT EDIT.
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	. "github.com/kurtosis-tech/kardinal/libs/cli-kontrol-api/api/golang/types"
	"github.com/oapi-codegen/runtime"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetHealth request
	GetHealth(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostTenantUuidDeployWithBody request with any body
	PostTenantUuidDeployWithBody(ctx context.Context, uuid Uuid, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostTenantUuidDeploy(ctx context.Context, uuid Uuid, body PostTenantUuidDeployJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostTenantUuidFlowCreateWithBody request with any body
	PostTenantUuidFlowCreateWithBody(ctx context.Context, uuid Uuid, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostTenantUuidFlowCreate(ctx context.Context, uuid Uuid, body PostTenantUuidFlowCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteTenantUuidFlowFlowId request
	DeleteTenantUuidFlowFlowId(ctx context.Context, uuid Uuid, flowId FlowId, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetTenantUuidFlows request
	GetTenantUuidFlows(ctx context.Context, uuid Uuid, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostTenantUuidTemplatesCreateWithBody request with any body
	PostTenantUuidTemplatesCreateWithBody(ctx context.Context, uuid Uuid, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostTenantUuidTemplatesCreate(ctx context.Context, uuid Uuid, body PostTenantUuidTemplatesCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetTenantUuidTopology request
	GetTenantUuidTopology(ctx context.Context, uuid Uuid, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetHealth(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetHealthRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostTenantUuidDeployWithBody(ctx context.Context, uuid Uuid, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostTenantUuidDeployRequestWithBody(c.Server, uuid, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostTenantUuidDeploy(ctx context.Context, uuid Uuid, body PostTenantUuidDeployJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostTenantUuidDeployRequest(c.Server, uuid, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostTenantUuidFlowCreateWithBody(ctx context.Context, uuid Uuid, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostTenantUuidFlowCreateRequestWithBody(c.Server, uuid, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostTenantUuidFlowCreate(ctx context.Context, uuid Uuid, body PostTenantUuidFlowCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostTenantUuidFlowCreateRequest(c.Server, uuid, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteTenantUuidFlowFlowId(ctx context.Context, uuid Uuid, flowId FlowId, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteTenantUuidFlowFlowIdRequest(c.Server, uuid, flowId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetTenantUuidFlows(ctx context.Context, uuid Uuid, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetTenantUuidFlowsRequest(c.Server, uuid)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostTenantUuidTemplatesCreateWithBody(ctx context.Context, uuid Uuid, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostTenantUuidTemplatesCreateRequestWithBody(c.Server, uuid, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostTenantUuidTemplatesCreate(ctx context.Context, uuid Uuid, body PostTenantUuidTemplatesCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostTenantUuidTemplatesCreateRequest(c.Server, uuid, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetTenantUuidTopology(ctx context.Context, uuid Uuid, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetTenantUuidTopologyRequest(c.Server, uuid)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetHealthRequest generates requests for GetHealth
func NewGetHealthRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/health")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPostTenantUuidDeployRequest calls the generic PostTenantUuidDeploy builder with application/json body
func NewPostTenantUuidDeployRequest(server string, uuid Uuid, body PostTenantUuidDeployJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostTenantUuidDeployRequestWithBody(server, uuid, "application/json", bodyReader)
}

// NewPostTenantUuidDeployRequestWithBody generates requests for PostTenantUuidDeploy with any type of body
func NewPostTenantUuidDeployRequestWithBody(server string, uuid Uuid, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "uuid", runtime.ParamLocationPath, uuid)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/tenant/%s/deploy", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewPostTenantUuidFlowCreateRequest calls the generic PostTenantUuidFlowCreate builder with application/json body
func NewPostTenantUuidFlowCreateRequest(server string, uuid Uuid, body PostTenantUuidFlowCreateJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostTenantUuidFlowCreateRequestWithBody(server, uuid, "application/json", bodyReader)
}

// NewPostTenantUuidFlowCreateRequestWithBody generates requests for PostTenantUuidFlowCreate with any type of body
func NewPostTenantUuidFlowCreateRequestWithBody(server string, uuid Uuid, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "uuid", runtime.ParamLocationPath, uuid)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/tenant/%s/flow/create", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewDeleteTenantUuidFlowFlowIdRequest generates requests for DeleteTenantUuidFlowFlowId
func NewDeleteTenantUuidFlowFlowIdRequest(server string, uuid Uuid, flowId FlowId) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "uuid", runtime.ParamLocationPath, uuid)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "flow-id", runtime.ParamLocationPath, flowId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/tenant/%s/flow/%s", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetTenantUuidFlowsRequest generates requests for GetTenantUuidFlows
func NewGetTenantUuidFlowsRequest(server string, uuid Uuid) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "uuid", runtime.ParamLocationPath, uuid)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/tenant/%s/flows", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPostTenantUuidTemplatesCreateRequest calls the generic PostTenantUuidTemplatesCreate builder with application/json body
func NewPostTenantUuidTemplatesCreateRequest(server string, uuid Uuid, body PostTenantUuidTemplatesCreateJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostTenantUuidTemplatesCreateRequestWithBody(server, uuid, "application/json", bodyReader)
}

// NewPostTenantUuidTemplatesCreateRequestWithBody generates requests for PostTenantUuidTemplatesCreate with any type of body
func NewPostTenantUuidTemplatesCreateRequestWithBody(server string, uuid Uuid, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "uuid", runtime.ParamLocationPath, uuid)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/tenant/%s/templates/create", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewGetTenantUuidTopologyRequest generates requests for GetTenantUuidTopology
func NewGetTenantUuidTopologyRequest(server string, uuid Uuid) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "uuid", runtime.ParamLocationPath, uuid)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/tenant/%s/topology", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetHealthWithResponse request
	GetHealthWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetHealthResponse, error)

	// PostTenantUuidDeployWithBodyWithResponse request with any body
	PostTenantUuidDeployWithBodyWithResponse(ctx context.Context, uuid Uuid, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostTenantUuidDeployResponse, error)

	PostTenantUuidDeployWithResponse(ctx context.Context, uuid Uuid, body PostTenantUuidDeployJSONRequestBody, reqEditors ...RequestEditorFn) (*PostTenantUuidDeployResponse, error)

	// PostTenantUuidFlowCreateWithBodyWithResponse request with any body
	PostTenantUuidFlowCreateWithBodyWithResponse(ctx context.Context, uuid Uuid, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostTenantUuidFlowCreateResponse, error)

	PostTenantUuidFlowCreateWithResponse(ctx context.Context, uuid Uuid, body PostTenantUuidFlowCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*PostTenantUuidFlowCreateResponse, error)

	// DeleteTenantUuidFlowFlowIdWithResponse request
	DeleteTenantUuidFlowFlowIdWithResponse(ctx context.Context, uuid Uuid, flowId FlowId, reqEditors ...RequestEditorFn) (*DeleteTenantUuidFlowFlowIdResponse, error)

	// GetTenantUuidFlowsWithResponse request
	GetTenantUuidFlowsWithResponse(ctx context.Context, uuid Uuid, reqEditors ...RequestEditorFn) (*GetTenantUuidFlowsResponse, error)

	// PostTenantUuidTemplatesCreateWithBodyWithResponse request with any body
	PostTenantUuidTemplatesCreateWithBodyWithResponse(ctx context.Context, uuid Uuid, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostTenantUuidTemplatesCreateResponse, error)

	PostTenantUuidTemplatesCreateWithResponse(ctx context.Context, uuid Uuid, body PostTenantUuidTemplatesCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*PostTenantUuidTemplatesCreateResponse, error)

	// GetTenantUuidTopologyWithResponse request
	GetTenantUuidTopologyWithResponse(ctx context.Context, uuid Uuid, reqEditors ...RequestEditorFn) (*GetTenantUuidTopologyResponse, error)
}

type GetHealthResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *string
}

// Status returns HTTPResponse.Status
func (r GetHealthResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetHealthResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostTenantUuidDeployResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Flow
	JSON404      *NotFound
	JSON500      *Error
}

// Status returns HTTPResponse.Status
func (r PostTenantUuidDeployResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostTenantUuidDeployResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostTenantUuidFlowCreateResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Flow
	JSON404      *NotFound
	JSON500      *Error
}

// Status returns HTTPResponse.Status
func (r PostTenantUuidFlowCreateResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostTenantUuidFlowCreateResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteTenantUuidFlowFlowIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON404      *NotFound
	JSON500      *Error
}

// Status returns HTTPResponse.Status
func (r DeleteTenantUuidFlowFlowIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteTenantUuidFlowFlowIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetTenantUuidFlowsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]Flow
	JSON404      *NotFound
	JSON500      *Error
}

// Status returns HTTPResponse.Status
func (r GetTenantUuidFlowsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetTenantUuidFlowsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostTenantUuidTemplatesCreateResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Template
	JSON404      *NotFound
	JSON500      *Error
}

// Status returns HTTPResponse.Status
func (r PostTenantUuidTemplatesCreateResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostTenantUuidTemplatesCreateResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetTenantUuidTopologyResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ClusterTopology
	JSON404      *NotFound
	JSON500      *Error
}

// Status returns HTTPResponse.Status
func (r GetTenantUuidTopologyResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetTenantUuidTopologyResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetHealthWithResponse request returning *GetHealthResponse
func (c *ClientWithResponses) GetHealthWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetHealthResponse, error) {
	rsp, err := c.GetHealth(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetHealthResponse(rsp)
}

// PostTenantUuidDeployWithBodyWithResponse request with arbitrary body returning *PostTenantUuidDeployResponse
func (c *ClientWithResponses) PostTenantUuidDeployWithBodyWithResponse(ctx context.Context, uuid Uuid, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostTenantUuidDeployResponse, error) {
	rsp, err := c.PostTenantUuidDeployWithBody(ctx, uuid, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostTenantUuidDeployResponse(rsp)
}

func (c *ClientWithResponses) PostTenantUuidDeployWithResponse(ctx context.Context, uuid Uuid, body PostTenantUuidDeployJSONRequestBody, reqEditors ...RequestEditorFn) (*PostTenantUuidDeployResponse, error) {
	rsp, err := c.PostTenantUuidDeploy(ctx, uuid, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostTenantUuidDeployResponse(rsp)
}

// PostTenantUuidFlowCreateWithBodyWithResponse request with arbitrary body returning *PostTenantUuidFlowCreateResponse
func (c *ClientWithResponses) PostTenantUuidFlowCreateWithBodyWithResponse(ctx context.Context, uuid Uuid, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostTenantUuidFlowCreateResponse, error) {
	rsp, err := c.PostTenantUuidFlowCreateWithBody(ctx, uuid, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostTenantUuidFlowCreateResponse(rsp)
}

func (c *ClientWithResponses) PostTenantUuidFlowCreateWithResponse(ctx context.Context, uuid Uuid, body PostTenantUuidFlowCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*PostTenantUuidFlowCreateResponse, error) {
	rsp, err := c.PostTenantUuidFlowCreate(ctx, uuid, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostTenantUuidFlowCreateResponse(rsp)
}

// DeleteTenantUuidFlowFlowIdWithResponse request returning *DeleteTenantUuidFlowFlowIdResponse
func (c *ClientWithResponses) DeleteTenantUuidFlowFlowIdWithResponse(ctx context.Context, uuid Uuid, flowId FlowId, reqEditors ...RequestEditorFn) (*DeleteTenantUuidFlowFlowIdResponse, error) {
	rsp, err := c.DeleteTenantUuidFlowFlowId(ctx, uuid, flowId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteTenantUuidFlowFlowIdResponse(rsp)
}

// GetTenantUuidFlowsWithResponse request returning *GetTenantUuidFlowsResponse
func (c *ClientWithResponses) GetTenantUuidFlowsWithResponse(ctx context.Context, uuid Uuid, reqEditors ...RequestEditorFn) (*GetTenantUuidFlowsResponse, error) {
	rsp, err := c.GetTenantUuidFlows(ctx, uuid, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetTenantUuidFlowsResponse(rsp)
}

// PostTenantUuidTemplatesCreateWithBodyWithResponse request with arbitrary body returning *PostTenantUuidTemplatesCreateResponse
func (c *ClientWithResponses) PostTenantUuidTemplatesCreateWithBodyWithResponse(ctx context.Context, uuid Uuid, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostTenantUuidTemplatesCreateResponse, error) {
	rsp, err := c.PostTenantUuidTemplatesCreateWithBody(ctx, uuid, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostTenantUuidTemplatesCreateResponse(rsp)
}

func (c *ClientWithResponses) PostTenantUuidTemplatesCreateWithResponse(ctx context.Context, uuid Uuid, body PostTenantUuidTemplatesCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*PostTenantUuidTemplatesCreateResponse, error) {
	rsp, err := c.PostTenantUuidTemplatesCreate(ctx, uuid, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostTenantUuidTemplatesCreateResponse(rsp)
}

// GetTenantUuidTopologyWithResponse request returning *GetTenantUuidTopologyResponse
func (c *ClientWithResponses) GetTenantUuidTopologyWithResponse(ctx context.Context, uuid Uuid, reqEditors ...RequestEditorFn) (*GetTenantUuidTopologyResponse, error) {
	rsp, err := c.GetTenantUuidTopology(ctx, uuid, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetTenantUuidTopologyResponse(rsp)
}

// ParseGetHealthResponse parses an HTTP response from a GetHealthWithResponse call
func ParseGetHealthResponse(rsp *http.Response) (*GetHealthResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetHealthResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest string
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParsePostTenantUuidDeployResponse parses an HTTP response from a PostTenantUuidDeployWithResponse call
func ParsePostTenantUuidDeployResponse(rsp *http.Response) (*PostTenantUuidDeployResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostTenantUuidDeployResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Flow
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParsePostTenantUuidFlowCreateResponse parses an HTTP response from a PostTenantUuidFlowCreateWithResponse call
func ParsePostTenantUuidFlowCreateResponse(rsp *http.Response) (*PostTenantUuidFlowCreateResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostTenantUuidFlowCreateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Flow
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseDeleteTenantUuidFlowFlowIdResponse parses an HTTP response from a DeleteTenantUuidFlowFlowIdWithResponse call
func ParseDeleteTenantUuidFlowFlowIdResponse(rsp *http.Response) (*DeleteTenantUuidFlowFlowIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteTenantUuidFlowFlowIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetTenantUuidFlowsResponse parses an HTTP response from a GetTenantUuidFlowsWithResponse call
func ParseGetTenantUuidFlowsResponse(rsp *http.Response) (*GetTenantUuidFlowsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetTenantUuidFlowsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []Flow
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParsePostTenantUuidTemplatesCreateResponse parses an HTTP response from a PostTenantUuidTemplatesCreateWithResponse call
func ParsePostTenantUuidTemplatesCreateResponse(rsp *http.Response) (*PostTenantUuidTemplatesCreateResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostTenantUuidTemplatesCreateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Template
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetTenantUuidTopologyResponse parses an HTTP response from a GetTenantUuidTopologyWithResponse call
func ParseGetTenantUuidTopologyResponse(rsp *http.Response) (*GetTenantUuidTopologyResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetTenantUuidTopologyResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ClusterTopology
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}
