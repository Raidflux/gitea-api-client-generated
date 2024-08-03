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

// NewUpdateUserSecretParams creates a new UpdateUserSecretParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateUserSecretParams() *UpdateUserSecretParams {
	return &UpdateUserSecretParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateUserSecretParamsWithTimeout creates a new UpdateUserSecretParams object
// with the ability to set a timeout on a request.
func NewUpdateUserSecretParamsWithTimeout(timeout time.Duration) *UpdateUserSecretParams {
	return &UpdateUserSecretParams{
		timeout: timeout,
	}
}

// NewUpdateUserSecretParamsWithContext creates a new UpdateUserSecretParams object
// with the ability to set a context for a request.
func NewUpdateUserSecretParamsWithContext(ctx context.Context) *UpdateUserSecretParams {
	return &UpdateUserSecretParams{
		Context: ctx,
	}
}

// NewUpdateUserSecretParamsWithHTTPClient creates a new UpdateUserSecretParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateUserSecretParamsWithHTTPClient(client *http.Client) *UpdateUserSecretParams {
	return &UpdateUserSecretParams{
		HTTPClient: client,
	}
}

/*
UpdateUserSecretParams contains all the parameters to send to the API endpoint

	for the update user secret operation.

	Typically these are written to a http.Request.
*/
type UpdateUserSecretParams struct {

	// Body.
	Body *models.CreateOrUpdateSecretOption

	/* Secretname.

	   name of the secret
	*/
	Secretname string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update user secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUserSecretParams) WithDefaults() *UpdateUserSecretParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update user secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUserSecretParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update user secret params
func (o *UpdateUserSecretParams) WithTimeout(timeout time.Duration) *UpdateUserSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update user secret params
func (o *UpdateUserSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update user secret params
func (o *UpdateUserSecretParams) WithContext(ctx context.Context) *UpdateUserSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update user secret params
func (o *UpdateUserSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update user secret params
func (o *UpdateUserSecretParams) WithHTTPClient(client *http.Client) *UpdateUserSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update user secret params
func (o *UpdateUserSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update user secret params
func (o *UpdateUserSecretParams) WithBody(body *models.CreateOrUpdateSecretOption) *UpdateUserSecretParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update user secret params
func (o *UpdateUserSecretParams) SetBody(body *models.CreateOrUpdateSecretOption) {
	o.Body = body
}

// WithSecretname adds the secretname to the update user secret params
func (o *UpdateUserSecretParams) WithSecretname(secretname string) *UpdateUserSecretParams {
	o.SetSecretname(secretname)
	return o
}

// SetSecretname adds the secretname to the update user secret params
func (o *UpdateUserSecretParams) SetSecretname(secretname string) {
	o.Secretname = secretname
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateUserSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param secretname
	if err := r.SetPathParam("secretname", o.Secretname); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
