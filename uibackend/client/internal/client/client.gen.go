// Package client provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.2.0 DO NOT EDIT.
package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
	. "github.com/openclarity/vmclarity/uibackend/types"
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
	// GetDashboardFindingsImpact request
	GetDashboardFindingsImpact(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetDashboardFindingsTrends request
	GetDashboardFindingsTrends(ctx context.Context, params *GetDashboardFindingsTrendsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetDashboardRiskiestAssets request
	GetDashboardRiskiestAssets(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetDashboardRiskiestRegions request
	GetDashboardRiskiestRegions(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetDashboardFindingsImpact(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetDashboardFindingsImpactRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetDashboardFindingsTrends(ctx context.Context, params *GetDashboardFindingsTrendsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetDashboardFindingsTrendsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetDashboardRiskiestAssets(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetDashboardRiskiestAssetsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetDashboardRiskiestRegions(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetDashboardRiskiestRegionsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetDashboardFindingsImpactRequest generates requests for GetDashboardFindingsImpact
func NewGetDashboardFindingsImpactRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/dashboard/findingsImpact")
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

// NewGetDashboardFindingsTrendsRequest generates requests for GetDashboardFindingsTrends
func NewGetDashboardFindingsTrendsRequest(server string, params *GetDashboardFindingsTrendsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/dashboard/findingsTrends")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "startTime", runtime.ParamLocationQuery, params.StartTime); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "endTime", runtime.ParamLocationQuery, params.EndTime); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetDashboardRiskiestAssetsRequest generates requests for GetDashboardRiskiestAssets
func NewGetDashboardRiskiestAssetsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/dashboard/riskiestAssets")
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

// NewGetDashboardRiskiestRegionsRequest generates requests for GetDashboardRiskiestRegions
func NewGetDashboardRiskiestRegionsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/dashboard/riskiestRegions")
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
	// GetDashboardFindingsImpactWithResponse request
	GetDashboardFindingsImpactWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetDashboardFindingsImpactResponse, error)

	// GetDashboardFindingsTrendsWithResponse request
	GetDashboardFindingsTrendsWithResponse(ctx context.Context, params *GetDashboardFindingsTrendsParams, reqEditors ...RequestEditorFn) (*GetDashboardFindingsTrendsResponse, error)

	// GetDashboardRiskiestAssetsWithResponse request
	GetDashboardRiskiestAssetsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetDashboardRiskiestAssetsResponse, error)

	// GetDashboardRiskiestRegionsWithResponse request
	GetDashboardRiskiestRegionsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetDashboardRiskiestRegionsResponse, error)
}

type GetDashboardFindingsImpactResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *FindingsImpact
	JSONDefault  *UnknownError
}

// Status returns HTTPResponse.Status
func (r GetDashboardFindingsImpactResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetDashboardFindingsImpactResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetDashboardFindingsTrendsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *FindingsTrends
	JSONDefault  *UnknownError
}

// Status returns HTTPResponse.Status
func (r GetDashboardFindingsTrendsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetDashboardFindingsTrendsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetDashboardRiskiestAssetsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *RiskiestAssets
	JSONDefault  *UnknownError
}

// Status returns HTTPResponse.Status
func (r GetDashboardRiskiestAssetsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetDashboardRiskiestAssetsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetDashboardRiskiestRegionsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *RiskiestRegions
	JSONDefault  *UnknownError
}

// Status returns HTTPResponse.Status
func (r GetDashboardRiskiestRegionsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetDashboardRiskiestRegionsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetDashboardFindingsImpactWithResponse request returning *GetDashboardFindingsImpactResponse
func (c *ClientWithResponses) GetDashboardFindingsImpactWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetDashboardFindingsImpactResponse, error) {
	rsp, err := c.GetDashboardFindingsImpact(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetDashboardFindingsImpactResponse(rsp)
}

// GetDashboardFindingsTrendsWithResponse request returning *GetDashboardFindingsTrendsResponse
func (c *ClientWithResponses) GetDashboardFindingsTrendsWithResponse(ctx context.Context, params *GetDashboardFindingsTrendsParams, reqEditors ...RequestEditorFn) (*GetDashboardFindingsTrendsResponse, error) {
	rsp, err := c.GetDashboardFindingsTrends(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetDashboardFindingsTrendsResponse(rsp)
}

// GetDashboardRiskiestAssetsWithResponse request returning *GetDashboardRiskiestAssetsResponse
func (c *ClientWithResponses) GetDashboardRiskiestAssetsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetDashboardRiskiestAssetsResponse, error) {
	rsp, err := c.GetDashboardRiskiestAssets(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetDashboardRiskiestAssetsResponse(rsp)
}

// GetDashboardRiskiestRegionsWithResponse request returning *GetDashboardRiskiestRegionsResponse
func (c *ClientWithResponses) GetDashboardRiskiestRegionsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetDashboardRiskiestRegionsResponse, error) {
	rsp, err := c.GetDashboardRiskiestRegions(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetDashboardRiskiestRegionsResponse(rsp)
}

// ParseGetDashboardFindingsImpactResponse parses an HTTP response from a GetDashboardFindingsImpactWithResponse call
func ParseGetDashboardFindingsImpactResponse(rsp *http.Response) (*GetDashboardFindingsImpactResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetDashboardFindingsImpactResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest FindingsImpact
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest UnknownError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseGetDashboardFindingsTrendsResponse parses an HTTP response from a GetDashboardFindingsTrendsWithResponse call
func ParseGetDashboardFindingsTrendsResponse(rsp *http.Response) (*GetDashboardFindingsTrendsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetDashboardFindingsTrendsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest FindingsTrends
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest UnknownError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseGetDashboardRiskiestAssetsResponse parses an HTTP response from a GetDashboardRiskiestAssetsWithResponse call
func ParseGetDashboardRiskiestAssetsResponse(rsp *http.Response) (*GetDashboardRiskiestAssetsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetDashboardRiskiestAssetsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest RiskiestAssets
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest UnknownError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseGetDashboardRiskiestRegionsResponse parses an HTTP response from a GetDashboardRiskiestRegionsWithResponse call
func ParseGetDashboardRiskiestRegionsResponse(rsp *http.Response) (*GetDashboardRiskiestRegionsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetDashboardRiskiestRegionsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest RiskiestRegions
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest UnknownError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}
