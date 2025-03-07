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

	"github.com/Raidflux/gitea-api-client-generated/giteaApiClientGenerated/models"
)

// NewUserEditHookParams creates a new UserEditHookParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserEditHookParams() *UserEditHookParams {
	return &UserEditHookParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserEditHookParamsWithTimeout creates a new UserEditHookParams object
// with the ability to set a timeout on a request.
func NewUserEditHookParamsWithTimeout(timeout time.Duration) *UserEditHookParams {
	return &UserEditHookParams{
		timeout: timeout,
	}
}

// NewUserEditHookParamsWithContext creates a new UserEditHookParams object
// with the ability to set a context for a request.
func NewUserEditHookParamsWithContext(ctx context.Context) *UserEditHookParams {
	return &UserEditHookParams{
		Context: ctx,
	}
}

// NewUserEditHookParamsWithHTTPClient creates a new UserEditHookParams object
// with the ability to set a custom HTTPClient for a request.
func NewUserEditHookParamsWithHTTPClient(client *http.Client) *UserEditHookParams {
	return &UserEditHookParams{
		HTTPClient: client,
	}
}

/*
UserEditHookParams contains all the parameters to send to the API endpoint

	for the user edit hook operation.

	Typically these are written to a http.Request.
*/
type UserEditHookParams struct {

	// Body.
	Body *models.EditHookOption

	/* ID.

	   id of the hook to update

	   Format: int64
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the user edit hook params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserEditHookParams) WithDefaults() *UserEditHookParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user edit hook params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserEditHookParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the user edit hook params
func (o *UserEditHookParams) WithTimeout(timeout time.Duration) *UserEditHookParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user edit hook params
func (o *UserEditHookParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user edit hook params
func (o *UserEditHookParams) WithContext(ctx context.Context) *UserEditHookParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user edit hook params
func (o *UserEditHookParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user edit hook params
func (o *UserEditHookParams) WithHTTPClient(client *http.Client) *UserEditHookParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user edit hook params
func (o *UserEditHookParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the user edit hook params
func (o *UserEditHookParams) WithBody(body *models.EditHookOption) *UserEditHookParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the user edit hook params
func (o *UserEditHookParams) SetBody(body *models.EditHookOption) {
	o.Body = body
}

// WithID adds the id to the user edit hook params
func (o *UserEditHookParams) WithID(id int64) *UserEditHookParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the user edit hook params
func (o *UserEditHookParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UserEditHookParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
