// Code generated by go-swagger; DO NOT EDIT.

package miscellaneous

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

// NewGetSigningKeyParams creates a new GetSigningKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSigningKeyParams() *GetSigningKeyParams {
	return &GetSigningKeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSigningKeyParamsWithTimeout creates a new GetSigningKeyParams object
// with the ability to set a timeout on a request.
func NewGetSigningKeyParamsWithTimeout(timeout time.Duration) *GetSigningKeyParams {
	return &GetSigningKeyParams{
		timeout: timeout,
	}
}

// NewGetSigningKeyParamsWithContext creates a new GetSigningKeyParams object
// with the ability to set a context for a request.
func NewGetSigningKeyParamsWithContext(ctx context.Context) *GetSigningKeyParams {
	return &GetSigningKeyParams{
		Context: ctx,
	}
}

// NewGetSigningKeyParamsWithHTTPClient creates a new GetSigningKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSigningKeyParamsWithHTTPClient(client *http.Client) *GetSigningKeyParams {
	return &GetSigningKeyParams{
		HTTPClient: client,
	}
}

/*
GetSigningKeyParams contains all the parameters to send to the API endpoint

	for the get signing key operation.

	Typically these are written to a http.Request.
*/
type GetSigningKeyParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get signing key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSigningKeyParams) WithDefaults() *GetSigningKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get signing key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSigningKeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get signing key params
func (o *GetSigningKeyParams) WithTimeout(timeout time.Duration) *GetSigningKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get signing key params
func (o *GetSigningKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get signing key params
func (o *GetSigningKeyParams) WithContext(ctx context.Context) *GetSigningKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get signing key params
func (o *GetSigningKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get signing key params
func (o *GetSigningKeyParams) WithHTTPClient(client *http.Client) *GetSigningKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get signing key params
func (o *GetSigningKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetSigningKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
