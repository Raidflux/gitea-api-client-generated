// Code generated by go-swagger; DO NOT EDIT.

package settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new settings API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new settings API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new settings API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for settings API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// This client is generated with a few options you might find useful for your swagger spec.
//
// Feel free to add you own set of options.

// WithContentType allows the client to force the Content-Type header
// to negotiate a specific Consumer from the server.
//
// You may use this option to set arbitrary extensions to your MIME media type.
func WithContentType(mime string) ClientOption {
	return func(r *runtime.ClientOperation) {
		r.ConsumesMediaTypes = []string{mime}
	}
}

// WithContentTypeApplicationJSON sets the Content-Type header to "application/json".
func WithContentTypeApplicationJSON(r *runtime.ClientOperation) {
	r.ConsumesMediaTypes = []string{"application/json"}
}

// WithContentTypeTextPlain sets the Content-Type header to "text/plain".
func WithContentTypeTextPlain(r *runtime.ClientOperation) {
	r.ConsumesMediaTypes = []string{"text/plain"}
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetGeneralAPISettings(params *GetGeneralAPISettingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGeneralAPISettingsOK, error)

	GetGeneralAttachmentSettings(params *GetGeneralAttachmentSettingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGeneralAttachmentSettingsOK, error)

	GetGeneralRepositorySettings(params *GetGeneralRepositorySettingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGeneralRepositorySettingsOK, error)

	GetGeneralUISettings(params *GetGeneralUISettingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGeneralUISettingsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetGeneralAPISettings gets instance s global settings for api
*/
func (a *Client) GetGeneralAPISettings(params *GetGeneralAPISettingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGeneralAPISettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGeneralAPISettingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getGeneralAPISettings",
		Method:             "GET",
		PathPattern:        "/settings/api",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetGeneralAPISettingsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetGeneralAPISettingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getGeneralAPISettings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetGeneralAttachmentSettings gets instance s global settings for attachment
*/
func (a *Client) GetGeneralAttachmentSettings(params *GetGeneralAttachmentSettingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGeneralAttachmentSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGeneralAttachmentSettingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getGeneralAttachmentSettings",
		Method:             "GET",
		PathPattern:        "/settings/attachment",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetGeneralAttachmentSettingsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetGeneralAttachmentSettingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getGeneralAttachmentSettings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetGeneralRepositorySettings gets instance s global settings for repositories
*/
func (a *Client) GetGeneralRepositorySettings(params *GetGeneralRepositorySettingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGeneralRepositorySettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGeneralRepositorySettingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getGeneralRepositorySettings",
		Method:             "GET",
		PathPattern:        "/settings/repository",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetGeneralRepositorySettingsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetGeneralRepositorySettingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getGeneralRepositorySettings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetGeneralUISettings gets instance s global settings for ui
*/
func (a *Client) GetGeneralUISettings(params *GetGeneralUISettingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGeneralUISettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGeneralUISettingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getGeneralUISettings",
		Method:             "GET",
		PathPattern:        "/settings/ui",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetGeneralUISettingsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetGeneralUISettingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getGeneralUISettings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
