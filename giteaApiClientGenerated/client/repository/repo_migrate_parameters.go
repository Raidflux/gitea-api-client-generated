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

// NewRepoMigrateParams creates a new RepoMigrateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepoMigrateParams() *RepoMigrateParams {
	return &RepoMigrateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepoMigrateParamsWithTimeout creates a new RepoMigrateParams object
// with the ability to set a timeout on a request.
func NewRepoMigrateParamsWithTimeout(timeout time.Duration) *RepoMigrateParams {
	return &RepoMigrateParams{
		timeout: timeout,
	}
}

// NewRepoMigrateParamsWithContext creates a new RepoMigrateParams object
// with the ability to set a context for a request.
func NewRepoMigrateParamsWithContext(ctx context.Context) *RepoMigrateParams {
	return &RepoMigrateParams{
		Context: ctx,
	}
}

// NewRepoMigrateParamsWithHTTPClient creates a new RepoMigrateParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepoMigrateParamsWithHTTPClient(client *http.Client) *RepoMigrateParams {
	return &RepoMigrateParams{
		HTTPClient: client,
	}
}

/*
RepoMigrateParams contains all the parameters to send to the API endpoint

	for the repo migrate operation.

	Typically these are written to a http.Request.
*/
type RepoMigrateParams struct {

	// Body.
	Body *models.MigrateRepoOptions

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the repo migrate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoMigrateParams) WithDefaults() *RepoMigrateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repo migrate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoMigrateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repo migrate params
func (o *RepoMigrateParams) WithTimeout(timeout time.Duration) *RepoMigrateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo migrate params
func (o *RepoMigrateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo migrate params
func (o *RepoMigrateParams) WithContext(ctx context.Context) *RepoMigrateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo migrate params
func (o *RepoMigrateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo migrate params
func (o *RepoMigrateParams) WithHTTPClient(client *http.Client) *RepoMigrateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo migrate params
func (o *RepoMigrateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the repo migrate params
func (o *RepoMigrateParams) WithBody(body *models.MigrateRepoOptions) *RepoMigrateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the repo migrate params
func (o *RepoMigrateParams) SetBody(body *models.MigrateRepoOptions) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RepoMigrateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
