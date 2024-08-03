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

	"giteaApiClientGenerated/models"
)

// NewUserUpdateOAuth2ApplicationParams creates a new UserUpdateOAuth2ApplicationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserUpdateOAuth2ApplicationParams() *UserUpdateOAuth2ApplicationParams {
	return &UserUpdateOAuth2ApplicationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserUpdateOAuth2ApplicationParamsWithTimeout creates a new UserUpdateOAuth2ApplicationParams object
// with the ability to set a timeout on a request.
func NewUserUpdateOAuth2ApplicationParamsWithTimeout(timeout time.Duration) *UserUpdateOAuth2ApplicationParams {
	return &UserUpdateOAuth2ApplicationParams{
		timeout: timeout,
	}
}

// NewUserUpdateOAuth2ApplicationParamsWithContext creates a new UserUpdateOAuth2ApplicationParams object
// with the ability to set a context for a request.
func NewUserUpdateOAuth2ApplicationParamsWithContext(ctx context.Context) *UserUpdateOAuth2ApplicationParams {
	return &UserUpdateOAuth2ApplicationParams{
		Context: ctx,
	}
}

// NewUserUpdateOAuth2ApplicationParamsWithHTTPClient creates a new UserUpdateOAuth2ApplicationParams object
// with the ability to set a custom HTTPClient for a request.
func NewUserUpdateOAuth2ApplicationParamsWithHTTPClient(client *http.Client) *UserUpdateOAuth2ApplicationParams {
	return &UserUpdateOAuth2ApplicationParams{
		HTTPClient: client,
	}
}

/*
UserUpdateOAuth2ApplicationParams contains all the parameters to send to the API endpoint

	for the user update o auth2 application operation.

	Typically these are written to a http.Request.
*/
type UserUpdateOAuth2ApplicationParams struct {

	// Body.
	Body *models.CreateOAuth2ApplicationOptions

	/* ID.

	   application to be updated

	   Format: int64
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the user update o auth2 application params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserUpdateOAuth2ApplicationParams) WithDefaults() *UserUpdateOAuth2ApplicationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user update o auth2 application params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserUpdateOAuth2ApplicationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the user update o auth2 application params
func (o *UserUpdateOAuth2ApplicationParams) WithTimeout(timeout time.Duration) *UserUpdateOAuth2ApplicationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user update o auth2 application params
func (o *UserUpdateOAuth2ApplicationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user update o auth2 application params
func (o *UserUpdateOAuth2ApplicationParams) WithContext(ctx context.Context) *UserUpdateOAuth2ApplicationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user update o auth2 application params
func (o *UserUpdateOAuth2ApplicationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user update o auth2 application params
func (o *UserUpdateOAuth2ApplicationParams) WithHTTPClient(client *http.Client) *UserUpdateOAuth2ApplicationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user update o auth2 application params
func (o *UserUpdateOAuth2ApplicationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the user update o auth2 application params
func (o *UserUpdateOAuth2ApplicationParams) WithBody(body *models.CreateOAuth2ApplicationOptions) *UserUpdateOAuth2ApplicationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the user update o auth2 application params
func (o *UserUpdateOAuth2ApplicationParams) SetBody(body *models.CreateOAuth2ApplicationOptions) {
	o.Body = body
}

// WithID adds the id to the user update o auth2 application params
func (o *UserUpdateOAuth2ApplicationParams) WithID(id int64) *UserUpdateOAuth2ApplicationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the user update o auth2 application params
func (o *UserUpdateOAuth2ApplicationParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UserUpdateOAuth2ApplicationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
