// Code generated by go-swagger; DO NOT EDIT.

package organization

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

// NewCreateOrgRepoParams creates a new CreateOrgRepoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateOrgRepoParams() *CreateOrgRepoParams {
	return &CreateOrgRepoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateOrgRepoParamsWithTimeout creates a new CreateOrgRepoParams object
// with the ability to set a timeout on a request.
func NewCreateOrgRepoParamsWithTimeout(timeout time.Duration) *CreateOrgRepoParams {
	return &CreateOrgRepoParams{
		timeout: timeout,
	}
}

// NewCreateOrgRepoParamsWithContext creates a new CreateOrgRepoParams object
// with the ability to set a context for a request.
func NewCreateOrgRepoParamsWithContext(ctx context.Context) *CreateOrgRepoParams {
	return &CreateOrgRepoParams{
		Context: ctx,
	}
}

// NewCreateOrgRepoParamsWithHTTPClient creates a new CreateOrgRepoParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateOrgRepoParamsWithHTTPClient(client *http.Client) *CreateOrgRepoParams {
	return &CreateOrgRepoParams{
		HTTPClient: client,
	}
}

/*
CreateOrgRepoParams contains all the parameters to send to the API endpoint

	for the create org repo operation.

	Typically these are written to a http.Request.
*/
type CreateOrgRepoParams struct {

	// Body.
	Body *models.CreateRepoOption

	/* Org.

	   name of organization
	*/
	Org string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create org repo params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrgRepoParams) WithDefaults() *CreateOrgRepoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create org repo params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrgRepoParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create org repo params
func (o *CreateOrgRepoParams) WithTimeout(timeout time.Duration) *CreateOrgRepoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create org repo params
func (o *CreateOrgRepoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create org repo params
func (o *CreateOrgRepoParams) WithContext(ctx context.Context) *CreateOrgRepoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create org repo params
func (o *CreateOrgRepoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create org repo params
func (o *CreateOrgRepoParams) WithHTTPClient(client *http.Client) *CreateOrgRepoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create org repo params
func (o *CreateOrgRepoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create org repo params
func (o *CreateOrgRepoParams) WithBody(body *models.CreateRepoOption) *CreateOrgRepoParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create org repo params
func (o *CreateOrgRepoParams) SetBody(body *models.CreateRepoOption) {
	o.Body = body
}

// WithOrg adds the org to the create org repo params
func (o *CreateOrgRepoParams) WithOrg(org string) *CreateOrgRepoParams {
	o.SetOrg(org)
	return o
}

// SetOrg adds the org to the create org repo params
func (o *CreateOrgRepoParams) SetOrg(org string) {
	o.Org = org
}

// WriteToRequest writes these params to a swagger request
func (o *CreateOrgRepoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param org
	if err := r.SetPathParam("org", o.Org); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
