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

// NewUserDeleteAvatarParams creates a new UserDeleteAvatarParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserDeleteAvatarParams() *UserDeleteAvatarParams {
	return &UserDeleteAvatarParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserDeleteAvatarParamsWithTimeout creates a new UserDeleteAvatarParams object
// with the ability to set a timeout on a request.
func NewUserDeleteAvatarParamsWithTimeout(timeout time.Duration) *UserDeleteAvatarParams {
	return &UserDeleteAvatarParams{
		timeout: timeout,
	}
}

// NewUserDeleteAvatarParamsWithContext creates a new UserDeleteAvatarParams object
// with the ability to set a context for a request.
func NewUserDeleteAvatarParamsWithContext(ctx context.Context) *UserDeleteAvatarParams {
	return &UserDeleteAvatarParams{
		Context: ctx,
	}
}

// NewUserDeleteAvatarParamsWithHTTPClient creates a new UserDeleteAvatarParams object
// with the ability to set a custom HTTPClient for a request.
func NewUserDeleteAvatarParamsWithHTTPClient(client *http.Client) *UserDeleteAvatarParams {
	return &UserDeleteAvatarParams{
		HTTPClient: client,
	}
}

/*
UserDeleteAvatarParams contains all the parameters to send to the API endpoint

	for the user delete avatar operation.

	Typically these are written to a http.Request.
*/
type UserDeleteAvatarParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the user delete avatar params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserDeleteAvatarParams) WithDefaults() *UserDeleteAvatarParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user delete avatar params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserDeleteAvatarParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the user delete avatar params
func (o *UserDeleteAvatarParams) WithTimeout(timeout time.Duration) *UserDeleteAvatarParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user delete avatar params
func (o *UserDeleteAvatarParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user delete avatar params
func (o *UserDeleteAvatarParams) WithContext(ctx context.Context) *UserDeleteAvatarParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user delete avatar params
func (o *UserDeleteAvatarParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user delete avatar params
func (o *UserDeleteAvatarParams) WithHTTPClient(client *http.Client) *UserDeleteAvatarParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user delete avatar params
func (o *UserDeleteAvatarParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *UserDeleteAvatarParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
