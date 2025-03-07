// Code generated by go-swagger; DO NOT EDIT.

package admin

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
	"github.com/go-openapi/swag"
)

// NewAdminGetHookParams creates a new AdminGetHookParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAdminGetHookParams() *AdminGetHookParams {
	return &AdminGetHookParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAdminGetHookParamsWithTimeout creates a new AdminGetHookParams object
// with the ability to set a timeout on a request.
func NewAdminGetHookParamsWithTimeout(timeout time.Duration) *AdminGetHookParams {
	return &AdminGetHookParams{
		timeout: timeout,
	}
}

// NewAdminGetHookParamsWithContext creates a new AdminGetHookParams object
// with the ability to set a context for a request.
func NewAdminGetHookParamsWithContext(ctx context.Context) *AdminGetHookParams {
	return &AdminGetHookParams{
		Context: ctx,
	}
}

// NewAdminGetHookParamsWithHTTPClient creates a new AdminGetHookParams object
// with the ability to set a custom HTTPClient for a request.
func NewAdminGetHookParamsWithHTTPClient(client *http.Client) *AdminGetHookParams {
	return &AdminGetHookParams{
		HTTPClient: client,
	}
}

/*
AdminGetHookParams contains all the parameters to send to the API endpoint

	for the admin get hook operation.

	Typically these are written to a http.Request.
*/
type AdminGetHookParams struct {

	/* ID.

	   id of the hook to get

	   Format: int64
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the admin get hook params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminGetHookParams) WithDefaults() *AdminGetHookParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the admin get hook params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminGetHookParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the admin get hook params
func (o *AdminGetHookParams) WithTimeout(timeout time.Duration) *AdminGetHookParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin get hook params
func (o *AdminGetHookParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin get hook params
func (o *AdminGetHookParams) WithContext(ctx context.Context) *AdminGetHookParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin get hook params
func (o *AdminGetHookParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin get hook params
func (o *AdminGetHookParams) WithHTTPClient(client *http.Client) *AdminGetHookParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin get hook params
func (o *AdminGetHookParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the admin get hook params
func (o *AdminGetHookParams) WithID(id int64) *AdminGetHookParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the admin get hook params
func (o *AdminGetHookParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AdminGetHookParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
