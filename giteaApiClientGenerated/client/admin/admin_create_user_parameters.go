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

	"github.com/Raidflux/gitea-api-client-generated/giteaApiClientGenerated/models"
)

// NewAdminCreateUserParams creates a new AdminCreateUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAdminCreateUserParams() *AdminCreateUserParams {
	return &AdminCreateUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAdminCreateUserParamsWithTimeout creates a new AdminCreateUserParams object
// with the ability to set a timeout on a request.
func NewAdminCreateUserParamsWithTimeout(timeout time.Duration) *AdminCreateUserParams {
	return &AdminCreateUserParams{
		timeout: timeout,
	}
}

// NewAdminCreateUserParamsWithContext creates a new AdminCreateUserParams object
// with the ability to set a context for a request.
func NewAdminCreateUserParamsWithContext(ctx context.Context) *AdminCreateUserParams {
	return &AdminCreateUserParams{
		Context: ctx,
	}
}

// NewAdminCreateUserParamsWithHTTPClient creates a new AdminCreateUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewAdminCreateUserParamsWithHTTPClient(client *http.Client) *AdminCreateUserParams {
	return &AdminCreateUserParams{
		HTTPClient: client,
	}
}

/*
AdminCreateUserParams contains all the parameters to send to the API endpoint

	for the admin create user operation.

	Typically these are written to a http.Request.
*/
type AdminCreateUserParams struct {

	// Body.
	Body *models.CreateUserOption

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the admin create user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminCreateUserParams) WithDefaults() *AdminCreateUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the admin create user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminCreateUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the admin create user params
func (o *AdminCreateUserParams) WithTimeout(timeout time.Duration) *AdminCreateUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin create user params
func (o *AdminCreateUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin create user params
func (o *AdminCreateUserParams) WithContext(ctx context.Context) *AdminCreateUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin create user params
func (o *AdminCreateUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin create user params
func (o *AdminCreateUserParams) WithHTTPClient(client *http.Client) *AdminCreateUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin create user params
func (o *AdminCreateUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the admin create user params
func (o *AdminCreateUserParams) WithBody(body *models.CreateUserOption) *AdminCreateUserParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the admin create user params
func (o *AdminCreateUserParams) SetBody(body *models.CreateUserOption) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AdminCreateUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
