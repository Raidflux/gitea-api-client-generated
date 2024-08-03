// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewGetVerificationTokenParams creates a new GetVerificationTokenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVerificationTokenParams() *GetVerificationTokenParams {
	return &GetVerificationTokenParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVerificationTokenParamsWithTimeout creates a new GetVerificationTokenParams object
// with the ability to set a timeout on a request.
func NewGetVerificationTokenParamsWithTimeout(timeout time.Duration) *GetVerificationTokenParams {
	return &GetVerificationTokenParams{
		timeout: timeout,
	}
}

// NewGetVerificationTokenParamsWithContext creates a new GetVerificationTokenParams object
// with the ability to set a context for a request.
func NewGetVerificationTokenParamsWithContext(ctx context.Context) *GetVerificationTokenParams {
	return &GetVerificationTokenParams{
		Context: ctx,
	}
}

// NewGetVerificationTokenParamsWithHTTPClient creates a new GetVerificationTokenParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVerificationTokenParamsWithHTTPClient(client *http.Client) *GetVerificationTokenParams {
	return &GetVerificationTokenParams{
		HTTPClient: client,
	}
}

/*
GetVerificationTokenParams contains all the parameters to send to the API endpoint

	for the get verification token operation.

	Typically these are written to a http.Request.
*/
type GetVerificationTokenParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get verification token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVerificationTokenParams) WithDefaults() *GetVerificationTokenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get verification token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVerificationTokenParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get verification token params
func (o *GetVerificationTokenParams) WithTimeout(timeout time.Duration) *GetVerificationTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get verification token params
func (o *GetVerificationTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get verification token params
func (o *GetVerificationTokenParams) WithContext(ctx context.Context) *GetVerificationTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get verification token params
func (o *GetVerificationTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get verification token params
func (o *GetVerificationTokenParams) WithHTTPClient(client *http.Client) *GetVerificationTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get verification token params
func (o *GetVerificationTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetVerificationTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
