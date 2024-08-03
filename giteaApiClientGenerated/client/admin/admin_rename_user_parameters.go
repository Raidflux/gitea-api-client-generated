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

	"giteaApiClientGenerated/models"
)

// NewAdminRenameUserParams creates a new AdminRenameUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAdminRenameUserParams() *AdminRenameUserParams {
	return &AdminRenameUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAdminRenameUserParamsWithTimeout creates a new AdminRenameUserParams object
// with the ability to set a timeout on a request.
func NewAdminRenameUserParamsWithTimeout(timeout time.Duration) *AdminRenameUserParams {
	return &AdminRenameUserParams{
		timeout: timeout,
	}
}

// NewAdminRenameUserParamsWithContext creates a new AdminRenameUserParams object
// with the ability to set a context for a request.
func NewAdminRenameUserParamsWithContext(ctx context.Context) *AdminRenameUserParams {
	return &AdminRenameUserParams{
		Context: ctx,
	}
}

// NewAdminRenameUserParamsWithHTTPClient creates a new AdminRenameUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewAdminRenameUserParamsWithHTTPClient(client *http.Client) *AdminRenameUserParams {
	return &AdminRenameUserParams{
		HTTPClient: client,
	}
}

/*
AdminRenameUserParams contains all the parameters to send to the API endpoint

	for the admin rename user operation.

	Typically these are written to a http.Request.
*/
type AdminRenameUserParams struct {

	// Body.
	Body *models.RenameUserOption

	/* Username.

	   existing username of user
	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the admin rename user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminRenameUserParams) WithDefaults() *AdminRenameUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the admin rename user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminRenameUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the admin rename user params
func (o *AdminRenameUserParams) WithTimeout(timeout time.Duration) *AdminRenameUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin rename user params
func (o *AdminRenameUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin rename user params
func (o *AdminRenameUserParams) WithContext(ctx context.Context) *AdminRenameUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin rename user params
func (o *AdminRenameUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin rename user params
func (o *AdminRenameUserParams) WithHTTPClient(client *http.Client) *AdminRenameUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin rename user params
func (o *AdminRenameUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the admin rename user params
func (o *AdminRenameUserParams) WithBody(body *models.RenameUserOption) *AdminRenameUserParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the admin rename user params
func (o *AdminRenameUserParams) SetBody(body *models.RenameUserOption) {
	o.Body = body
}

// WithUsername adds the username to the admin rename user params
func (o *AdminRenameUserParams) WithUsername(username string) *AdminRenameUserParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the admin rename user params
func (o *AdminRenameUserParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *AdminRenameUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
