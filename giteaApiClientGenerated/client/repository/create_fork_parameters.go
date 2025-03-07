// Code generated by go-swagger; DO NOT EDIT.

package repository

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

// NewCreateForkParams creates a new CreateForkParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateForkParams() *CreateForkParams {
	return &CreateForkParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateForkParamsWithTimeout creates a new CreateForkParams object
// with the ability to set a timeout on a request.
func NewCreateForkParamsWithTimeout(timeout time.Duration) *CreateForkParams {
	return &CreateForkParams{
		timeout: timeout,
	}
}

// NewCreateForkParamsWithContext creates a new CreateForkParams object
// with the ability to set a context for a request.
func NewCreateForkParamsWithContext(ctx context.Context) *CreateForkParams {
	return &CreateForkParams{
		Context: ctx,
	}
}

// NewCreateForkParamsWithHTTPClient creates a new CreateForkParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateForkParamsWithHTTPClient(client *http.Client) *CreateForkParams {
	return &CreateForkParams{
		HTTPClient: client,
	}
}

/*
CreateForkParams contains all the parameters to send to the API endpoint

	for the create fork operation.

	Typically these are written to a http.Request.
*/
type CreateForkParams struct {

	// Body.
	Body *models.CreateForkOption

	/* Owner.

	   owner of the repo to fork
	*/
	Owner string

	/* Repo.

	   name of the repo to fork
	*/
	Repo string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create fork params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateForkParams) WithDefaults() *CreateForkParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create fork params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateForkParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create fork params
func (o *CreateForkParams) WithTimeout(timeout time.Duration) *CreateForkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create fork params
func (o *CreateForkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create fork params
func (o *CreateForkParams) WithContext(ctx context.Context) *CreateForkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create fork params
func (o *CreateForkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create fork params
func (o *CreateForkParams) WithHTTPClient(client *http.Client) *CreateForkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create fork params
func (o *CreateForkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create fork params
func (o *CreateForkParams) WithBody(body *models.CreateForkOption) *CreateForkParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create fork params
func (o *CreateForkParams) SetBody(body *models.CreateForkOption) {
	o.Body = body
}

// WithOwner adds the owner to the create fork params
func (o *CreateForkParams) WithOwner(owner string) *CreateForkParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the create fork params
func (o *CreateForkParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the create fork params
func (o *CreateForkParams) WithRepo(repo string) *CreateForkParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the create fork params
func (o *CreateForkParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *CreateForkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param repo
	if err := r.SetPathParam("repo", o.Repo); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
