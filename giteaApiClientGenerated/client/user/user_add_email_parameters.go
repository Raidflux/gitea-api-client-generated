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

	"github.com/Raidflux/gitea-api-client-generated/giteaApiClientGenerated/models"
)

// NewUserAddEmailParams creates a new UserAddEmailParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserAddEmailParams() *UserAddEmailParams {
	return &UserAddEmailParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserAddEmailParamsWithTimeout creates a new UserAddEmailParams object
// with the ability to set a timeout on a request.
func NewUserAddEmailParamsWithTimeout(timeout time.Duration) *UserAddEmailParams {
	return &UserAddEmailParams{
		timeout: timeout,
	}
}

// NewUserAddEmailParamsWithContext creates a new UserAddEmailParams object
// with the ability to set a context for a request.
func NewUserAddEmailParamsWithContext(ctx context.Context) *UserAddEmailParams {
	return &UserAddEmailParams{
		Context: ctx,
	}
}

// NewUserAddEmailParamsWithHTTPClient creates a new UserAddEmailParams object
// with the ability to set a custom HTTPClient for a request.
func NewUserAddEmailParamsWithHTTPClient(client *http.Client) *UserAddEmailParams {
	return &UserAddEmailParams{
		HTTPClient: client,
	}
}

/*
UserAddEmailParams contains all the parameters to send to the API endpoint

	for the user add email operation.

	Typically these are written to a http.Request.
*/
type UserAddEmailParams struct {

	// Body.
	Body *models.CreateEmailOption

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the user add email params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserAddEmailParams) WithDefaults() *UserAddEmailParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user add email params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserAddEmailParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the user add email params
func (o *UserAddEmailParams) WithTimeout(timeout time.Duration) *UserAddEmailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user add email params
func (o *UserAddEmailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user add email params
func (o *UserAddEmailParams) WithContext(ctx context.Context) *UserAddEmailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user add email params
func (o *UserAddEmailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user add email params
func (o *UserAddEmailParams) WithHTTPClient(client *http.Client) *UserAddEmailParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user add email params
func (o *UserAddEmailParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the user add email params
func (o *UserAddEmailParams) WithBody(body *models.CreateEmailOption) *UserAddEmailParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the user add email params
func (o *UserAddEmailParams) SetBody(body *models.CreateEmailOption) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UserAddEmailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
