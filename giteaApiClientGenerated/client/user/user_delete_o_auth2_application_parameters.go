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

// NewUserDeleteOAuth2ApplicationParams creates a new UserDeleteOAuth2ApplicationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserDeleteOAuth2ApplicationParams() *UserDeleteOAuth2ApplicationParams {
	return &UserDeleteOAuth2ApplicationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserDeleteOAuth2ApplicationParamsWithTimeout creates a new UserDeleteOAuth2ApplicationParams object
// with the ability to set a timeout on a request.
func NewUserDeleteOAuth2ApplicationParamsWithTimeout(timeout time.Duration) *UserDeleteOAuth2ApplicationParams {
	return &UserDeleteOAuth2ApplicationParams{
		timeout: timeout,
	}
}

// NewUserDeleteOAuth2ApplicationParamsWithContext creates a new UserDeleteOAuth2ApplicationParams object
// with the ability to set a context for a request.
func NewUserDeleteOAuth2ApplicationParamsWithContext(ctx context.Context) *UserDeleteOAuth2ApplicationParams {
	return &UserDeleteOAuth2ApplicationParams{
		Context: ctx,
	}
}

// NewUserDeleteOAuth2ApplicationParamsWithHTTPClient creates a new UserDeleteOAuth2ApplicationParams object
// with the ability to set a custom HTTPClient for a request.
func NewUserDeleteOAuth2ApplicationParamsWithHTTPClient(client *http.Client) *UserDeleteOAuth2ApplicationParams {
	return &UserDeleteOAuth2ApplicationParams{
		HTTPClient: client,
	}
}

/*
UserDeleteOAuth2ApplicationParams contains all the parameters to send to the API endpoint

	for the user delete o auth2 application operation.

	Typically these are written to a http.Request.
*/
type UserDeleteOAuth2ApplicationParams struct {

	/* ID.

	   token to be deleted

	   Format: int64
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the user delete o auth2 application params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserDeleteOAuth2ApplicationParams) WithDefaults() *UserDeleteOAuth2ApplicationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user delete o auth2 application params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserDeleteOAuth2ApplicationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the user delete o auth2 application params
func (o *UserDeleteOAuth2ApplicationParams) WithTimeout(timeout time.Duration) *UserDeleteOAuth2ApplicationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user delete o auth2 application params
func (o *UserDeleteOAuth2ApplicationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user delete o auth2 application params
func (o *UserDeleteOAuth2ApplicationParams) WithContext(ctx context.Context) *UserDeleteOAuth2ApplicationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user delete o auth2 application params
func (o *UserDeleteOAuth2ApplicationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user delete o auth2 application params
func (o *UserDeleteOAuth2ApplicationParams) WithHTTPClient(client *http.Client) *UserDeleteOAuth2ApplicationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user delete o auth2 application params
func (o *UserDeleteOAuth2ApplicationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the user delete o auth2 application params
func (o *UserDeleteOAuth2ApplicationParams) WithID(id int64) *UserDeleteOAuth2ApplicationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the user delete o auth2 application params
func (o *UserDeleteOAuth2ApplicationParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UserDeleteOAuth2ApplicationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
