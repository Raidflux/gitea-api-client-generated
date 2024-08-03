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
)

// NewRepoListBranchProtectionParams creates a new RepoListBranchProtectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepoListBranchProtectionParams() *RepoListBranchProtectionParams {
	return &RepoListBranchProtectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepoListBranchProtectionParamsWithTimeout creates a new RepoListBranchProtectionParams object
// with the ability to set a timeout on a request.
func NewRepoListBranchProtectionParamsWithTimeout(timeout time.Duration) *RepoListBranchProtectionParams {
	return &RepoListBranchProtectionParams{
		timeout: timeout,
	}
}

// NewRepoListBranchProtectionParamsWithContext creates a new RepoListBranchProtectionParams object
// with the ability to set a context for a request.
func NewRepoListBranchProtectionParamsWithContext(ctx context.Context) *RepoListBranchProtectionParams {
	return &RepoListBranchProtectionParams{
		Context: ctx,
	}
}

// NewRepoListBranchProtectionParamsWithHTTPClient creates a new RepoListBranchProtectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepoListBranchProtectionParamsWithHTTPClient(client *http.Client) *RepoListBranchProtectionParams {
	return &RepoListBranchProtectionParams{
		HTTPClient: client,
	}
}

/*
RepoListBranchProtectionParams contains all the parameters to send to the API endpoint

	for the repo list branch protection operation.

	Typically these are written to a http.Request.
*/
type RepoListBranchProtectionParams struct {

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

// WithDefaults hydrates default values in the repo list branch protection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoListBranchProtectionParams) WithDefaults() *RepoListBranchProtectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repo list branch protection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoListBranchProtectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repo list branch protection params
func (o *RepoListBranchProtectionParams) WithTimeout(timeout time.Duration) *RepoListBranchProtectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo list branch protection params
func (o *RepoListBranchProtectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo list branch protection params
func (o *RepoListBranchProtectionParams) WithContext(ctx context.Context) *RepoListBranchProtectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo list branch protection params
func (o *RepoListBranchProtectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo list branch protection params
func (o *RepoListBranchProtectionParams) WithHTTPClient(client *http.Client) *RepoListBranchProtectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo list branch protection params
func (o *RepoListBranchProtectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the repo list branch protection params
func (o *RepoListBranchProtectionParams) WithOwner(owner string) *RepoListBranchProtectionParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo list branch protection params
func (o *RepoListBranchProtectionParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo list branch protection params
func (o *RepoListBranchProtectionParams) WithRepo(repo string) *RepoListBranchProtectionParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo list branch protection params
func (o *RepoListBranchProtectionParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoListBranchProtectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
