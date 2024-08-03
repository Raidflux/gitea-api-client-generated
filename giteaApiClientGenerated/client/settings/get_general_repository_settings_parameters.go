// Code generated by go-swagger; DO NOT EDIT.

package settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetGeneralRepositorySettingsParams creates a new GetGeneralRepositorySettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetGeneralRepositorySettingsParams() *GetGeneralRepositorySettingsParams {
	return &GetGeneralRepositorySettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetGeneralRepositorySettingsParamsWithTimeout creates a new GetGeneralRepositorySettingsParams object
// with the ability to set a timeout on a request.
func NewGetGeneralRepositorySettingsParamsWithTimeout(timeout time.Duration) *GetGeneralRepositorySettingsParams {
	return &GetGeneralRepositorySettingsParams{
		timeout: timeout,
	}
}

// NewGetGeneralRepositorySettingsParamsWithContext creates a new GetGeneralRepositorySettingsParams object
// with the ability to set a context for a request.
func NewGetGeneralRepositorySettingsParamsWithContext(ctx context.Context) *GetGeneralRepositorySettingsParams {
	return &GetGeneralRepositorySettingsParams{
		Context: ctx,
	}
}

// NewGetGeneralRepositorySettingsParamsWithHTTPClient creates a new GetGeneralRepositorySettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetGeneralRepositorySettingsParamsWithHTTPClient(client *http.Client) *GetGeneralRepositorySettingsParams {
	return &GetGeneralRepositorySettingsParams{
		HTTPClient: client,
	}
}

/*
GetGeneralRepositorySettingsParams contains all the parameters to send to the API endpoint

	for the get general repository settings operation.

	Typically these are written to a http.Request.
*/
type GetGeneralRepositorySettingsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get general repository settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGeneralRepositorySettingsParams) WithDefaults() *GetGeneralRepositorySettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get general repository settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGeneralRepositorySettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get general repository settings params
func (o *GetGeneralRepositorySettingsParams) WithTimeout(timeout time.Duration) *GetGeneralRepositorySettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get general repository settings params
func (o *GetGeneralRepositorySettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get general repository settings params
func (o *GetGeneralRepositorySettingsParams) WithContext(ctx context.Context) *GetGeneralRepositorySettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get general repository settings params
func (o *GetGeneralRepositorySettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get general repository settings params
func (o *GetGeneralRepositorySettingsParams) WithHTTPClient(client *http.Client) *GetGeneralRepositorySettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get general repository settings params
func (o *GetGeneralRepositorySettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetGeneralRepositorySettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
