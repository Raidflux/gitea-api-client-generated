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

// NewRepoCreateBranchProtectionParams creates a new RepoCreateBranchProtectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepoCreateBranchProtectionParams() *RepoCreateBranchProtectionParams {
	return &RepoCreateBranchProtectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepoCreateBranchProtectionParamsWithTimeout creates a new RepoCreateBranchProtectionParams object
// with the ability to set a timeout on a request.
func NewRepoCreateBranchProtectionParamsWithTimeout(timeout time.Duration) *RepoCreateBranchProtectionParams {
	return &RepoCreateBranchProtectionParams{
		timeout: timeout,
	}
}

// NewRepoCreateBranchProtectionParamsWithContext creates a new RepoCreateBranchProtectionParams object
// with the ability to set a context for a request.
func NewRepoCreateBranchProtectionParamsWithContext(ctx context.Context) *RepoCreateBranchProtectionParams {
	return &RepoCreateBranchProtectionParams{
		Context: ctx,
	}
}

// NewRepoCreateBranchProtectionParamsWithHTTPClient creates a new RepoCreateBranchProtectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepoCreateBranchProtectionParamsWithHTTPClient(client *http.Client) *RepoCreateBranchProtectionParams {
	return &RepoCreateBranchProtectionParams{
		HTTPClient: client,
	}
}

/*
RepoCreateBranchProtectionParams contains all the parameters to send to the API endpoint

	for the repo create branch protection operation.

	Typically these are written to a http.Request.
*/
type RepoCreateBranchProtectionParams struct {

	// Body.
	Body *models.CreateBranchProtectionOption

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

// WithDefaults hydrates default values in the repo create branch protection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoCreateBranchProtectionParams) WithDefaults() *RepoCreateBranchProtectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repo create branch protection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoCreateBranchProtectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repo create branch protection params
func (o *RepoCreateBranchProtectionParams) WithTimeout(timeout time.Duration) *RepoCreateBranchProtectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo create branch protection params
func (o *RepoCreateBranchProtectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo create branch protection params
func (o *RepoCreateBranchProtectionParams) WithContext(ctx context.Context) *RepoCreateBranchProtectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo create branch protection params
func (o *RepoCreateBranchProtectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo create branch protection params
func (o *RepoCreateBranchProtectionParams) WithHTTPClient(client *http.Client) *RepoCreateBranchProtectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo create branch protection params
func (o *RepoCreateBranchProtectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the repo create branch protection params
func (o *RepoCreateBranchProtectionParams) WithBody(body *models.CreateBranchProtectionOption) *RepoCreateBranchProtectionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the repo create branch protection params
func (o *RepoCreateBranchProtectionParams) SetBody(body *models.CreateBranchProtectionOption) {
	o.Body = body
}

// WithOwner adds the owner to the repo create branch protection params
func (o *RepoCreateBranchProtectionParams) WithOwner(owner string) *RepoCreateBranchProtectionParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo create branch protection params
func (o *RepoCreateBranchProtectionParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo create branch protection params
func (o *RepoCreateBranchProtectionParams) WithRepo(repo string) *RepoCreateBranchProtectionParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo create branch protection params
func (o *RepoCreateBranchProtectionParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoCreateBranchProtectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
