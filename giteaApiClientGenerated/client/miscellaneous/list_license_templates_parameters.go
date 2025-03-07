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

// NewListLicenseTemplatesParams creates a new ListLicenseTemplatesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListLicenseTemplatesParams() *ListLicenseTemplatesParams {
	return &ListLicenseTemplatesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListLicenseTemplatesParamsWithTimeout creates a new ListLicenseTemplatesParams object
// with the ability to set a timeout on a request.
func NewListLicenseTemplatesParamsWithTimeout(timeout time.Duration) *ListLicenseTemplatesParams {
	return &ListLicenseTemplatesParams{
		timeout: timeout,
	}
}

// NewListLicenseTemplatesParamsWithContext creates a new ListLicenseTemplatesParams object
// with the ability to set a context for a request.
func NewListLicenseTemplatesParamsWithContext(ctx context.Context) *ListLicenseTemplatesParams {
	return &ListLicenseTemplatesParams{
		Context: ctx,
	}
}

// NewListLicenseTemplatesParamsWithHTTPClient creates a new ListLicenseTemplatesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListLicenseTemplatesParamsWithHTTPClient(client *http.Client) *ListLicenseTemplatesParams {
	return &ListLicenseTemplatesParams{
		HTTPClient: client,
	}
}

/*
ListLicenseTemplatesParams contains all the parameters to send to the API endpoint

	for the list license templates operation.

	Typically these are written to a http.Request.
*/
type ListLicenseTemplatesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list license templates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListLicenseTemplatesParams) WithDefaults() *ListLicenseTemplatesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list license templates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListLicenseTemplatesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list license templates params
func (o *ListLicenseTemplatesParams) WithTimeout(timeout time.Duration) *ListLicenseTemplatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list license templates params
func (o *ListLicenseTemplatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list license templates params
func (o *ListLicenseTemplatesParams) WithContext(ctx context.Context) *ListLicenseTemplatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list license templates params
func (o *ListLicenseTemplatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list license templates params
func (o *ListLicenseTemplatesParams) WithHTTPClient(client *http.Client) *ListLicenseTemplatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list license templates params
func (o *ListLicenseTemplatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListLicenseTemplatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
