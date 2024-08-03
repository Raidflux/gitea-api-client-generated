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

// NewRepoChangeFilesParams creates a new RepoChangeFilesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepoChangeFilesParams() *RepoChangeFilesParams {
	return &RepoChangeFilesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepoChangeFilesParamsWithTimeout creates a new RepoChangeFilesParams object
// with the ability to set a timeout on a request.
func NewRepoChangeFilesParamsWithTimeout(timeout time.Duration) *RepoChangeFilesParams {
	return &RepoChangeFilesParams{
		timeout: timeout,
	}
}

// NewRepoChangeFilesParamsWithContext creates a new RepoChangeFilesParams object
// with the ability to set a context for a request.
func NewRepoChangeFilesParamsWithContext(ctx context.Context) *RepoChangeFilesParams {
	return &RepoChangeFilesParams{
		Context: ctx,
	}
}

// NewRepoChangeFilesParamsWithHTTPClient creates a new RepoChangeFilesParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepoChangeFilesParamsWithHTTPClient(client *http.Client) *RepoChangeFilesParams {
	return &RepoChangeFilesParams{
		HTTPClient: client,
	}
}

/*
RepoChangeFilesParams contains all the parameters to send to the API endpoint

	for the repo change files operation.

	Typically these are written to a http.Request.
*/
type RepoChangeFilesParams struct {

	// Body.
	Body *models.ChangeFilesOptions

	/* Owner.

	   owner of the repo
	*/
	Owner string

	/* Repo.

	   name of the repo
	*/
	Repo string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the repo change files params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoChangeFilesParams) WithDefaults() *RepoChangeFilesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repo change files params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoChangeFilesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repo change files params
func (o *RepoChangeFilesParams) WithTimeout(timeout time.Duration) *RepoChangeFilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo change files params
func (o *RepoChangeFilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo change files params
func (o *RepoChangeFilesParams) WithContext(ctx context.Context) *RepoChangeFilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo change files params
func (o *RepoChangeFilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo change files params
func (o *RepoChangeFilesParams) WithHTTPClient(client *http.Client) *RepoChangeFilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo change files params
func (o *RepoChangeFilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the repo change files params
func (o *RepoChangeFilesParams) WithBody(body *models.ChangeFilesOptions) *RepoChangeFilesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the repo change files params
func (o *RepoChangeFilesParams) SetBody(body *models.ChangeFilesOptions) {
	o.Body = body
}

// WithOwner adds the owner to the repo change files params
func (o *RepoChangeFilesParams) WithOwner(owner string) *RepoChangeFilesParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo change files params
func (o *RepoChangeFilesParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo change files params
func (o *RepoChangeFilesParams) WithRepo(repo string) *RepoChangeFilesParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo change files params
func (o *RepoChangeFilesParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoChangeFilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
