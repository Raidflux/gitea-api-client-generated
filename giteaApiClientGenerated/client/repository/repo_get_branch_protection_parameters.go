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

// NewRepoGetBranchProtectionParams creates a new RepoGetBranchProtectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepoGetBranchProtectionParams() *RepoGetBranchProtectionParams {
	return &RepoGetBranchProtectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepoGetBranchProtectionParamsWithTimeout creates a new RepoGetBranchProtectionParams object
// with the ability to set a timeout on a request.
func NewRepoGetBranchProtectionParamsWithTimeout(timeout time.Duration) *RepoGetBranchProtectionParams {
	return &RepoGetBranchProtectionParams{
		timeout: timeout,
	}
}

// NewRepoGetBranchProtectionParamsWithContext creates a new RepoGetBranchProtectionParams object
// with the ability to set a context for a request.
func NewRepoGetBranchProtectionParamsWithContext(ctx context.Context) *RepoGetBranchProtectionParams {
	return &RepoGetBranchProtectionParams{
		Context: ctx,
	}
}

// NewRepoGetBranchProtectionParamsWithHTTPClient creates a new RepoGetBranchProtectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepoGetBranchProtectionParamsWithHTTPClient(client *http.Client) *RepoGetBranchProtectionParams {
	return &RepoGetBranchProtectionParams{
		HTTPClient: client,
	}
}

/*
RepoGetBranchProtectionParams contains all the parameters to send to the API endpoint

	for the repo get branch protection operation.

	Typically these are written to a http.Request.
*/
type RepoGetBranchProtectionParams struct {

	/* Name.

	   name of protected branch
	*/
	Name string

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

// WithDefaults hydrates default values in the repo get branch protection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoGetBranchProtectionParams) WithDefaults() *RepoGetBranchProtectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repo get branch protection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoGetBranchProtectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repo get branch protection params
func (o *RepoGetBranchProtectionParams) WithTimeout(timeout time.Duration) *RepoGetBranchProtectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo get branch protection params
func (o *RepoGetBranchProtectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo get branch protection params
func (o *RepoGetBranchProtectionParams) WithContext(ctx context.Context) *RepoGetBranchProtectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo get branch protection params
func (o *RepoGetBranchProtectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo get branch protection params
func (o *RepoGetBranchProtectionParams) WithHTTPClient(client *http.Client) *RepoGetBranchProtectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo get branch protection params
func (o *RepoGetBranchProtectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the repo get branch protection params
func (o *RepoGetBranchProtectionParams) WithName(name string) *RepoGetBranchProtectionParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the repo get branch protection params
func (o *RepoGetBranchProtectionParams) SetName(name string) {
	o.Name = name
}

// WithOwner adds the owner to the repo get branch protection params
func (o *RepoGetBranchProtectionParams) WithOwner(owner string) *RepoGetBranchProtectionParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo get branch protection params
func (o *RepoGetBranchProtectionParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo get branch protection params
func (o *RepoGetBranchProtectionParams) WithRepo(repo string) *RepoGetBranchProtectionParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo get branch protection params
func (o *RepoGetBranchProtectionParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoGetBranchProtectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
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
