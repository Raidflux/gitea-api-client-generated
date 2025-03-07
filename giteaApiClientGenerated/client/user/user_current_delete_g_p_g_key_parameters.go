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
	"github.com/go-openapi/swag"
)

// NewUserCurrentDeleteGPGKeyParams creates a new UserCurrentDeleteGPGKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserCurrentDeleteGPGKeyParams() *UserCurrentDeleteGPGKeyParams {
	return &UserCurrentDeleteGPGKeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserCurrentDeleteGPGKeyParamsWithTimeout creates a new UserCurrentDeleteGPGKeyParams object
// with the ability to set a timeout on a request.
func NewUserCurrentDeleteGPGKeyParamsWithTimeout(timeout time.Duration) *UserCurrentDeleteGPGKeyParams {
	return &UserCurrentDeleteGPGKeyParams{
		timeout: timeout,
	}
}

// NewUserCurrentDeleteGPGKeyParamsWithContext creates a new UserCurrentDeleteGPGKeyParams object
// with the ability to set a context for a request.
func NewUserCurrentDeleteGPGKeyParamsWithContext(ctx context.Context) *UserCurrentDeleteGPGKeyParams {
	return &UserCurrentDeleteGPGKeyParams{
		Context: ctx,
	}
}

// NewUserCurrentDeleteGPGKeyParamsWithHTTPClient creates a new UserCurrentDeleteGPGKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewUserCurrentDeleteGPGKeyParamsWithHTTPClient(client *http.Client) *UserCurrentDeleteGPGKeyParams {
	return &UserCurrentDeleteGPGKeyParams{
		HTTPClient: client,
	}
}

/*
UserCurrentDeleteGPGKeyParams contains all the parameters to send to the API endpoint

	for the user current delete g p g key operation.

	Typically these are written to a http.Request.
*/
type UserCurrentDeleteGPGKeyParams struct {

	/* ID.

	   id of key to delete

	   Format: int64
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the user current delete g p g key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserCurrentDeleteGPGKeyParams) WithDefaults() *UserCurrentDeleteGPGKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user current delete g p g key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserCurrentDeleteGPGKeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the user current delete g p g key params
func (o *UserCurrentDeleteGPGKeyParams) WithTimeout(timeout time.Duration) *UserCurrentDeleteGPGKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user current delete g p g key params
func (o *UserCurrentDeleteGPGKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user current delete g p g key params
func (o *UserCurrentDeleteGPGKeyParams) WithContext(ctx context.Context) *UserCurrentDeleteGPGKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user current delete g p g key params
func (o *UserCurrentDeleteGPGKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user current delete g p g key params
func (o *UserCurrentDeleteGPGKeyParams) WithHTTPClient(client *http.Client) *UserCurrentDeleteGPGKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user current delete g p g key params
func (o *UserCurrentDeleteGPGKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the user current delete g p g key params
func (o *UserCurrentDeleteGPGKeyParams) WithID(id int64) *UserCurrentDeleteGPGKeyParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the user current delete g p g key params
func (o *UserCurrentDeleteGPGKeyParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UserCurrentDeleteGPGKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
