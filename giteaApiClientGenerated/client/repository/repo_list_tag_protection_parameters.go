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

// NewRepoListTagProtectionParams creates a new RepoListTagProtectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepoListTagProtectionParams() *RepoListTagProtectionParams {
	return &RepoListTagProtectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepoListTagProtectionParamsWithTimeout creates a new RepoListTagProtectionParams object
// with the ability to set a timeout on a request.
func NewRepoListTagProtectionParamsWithTimeout(timeout time.Duration) *RepoListTagProtectionParams {
	return &RepoListTagProtectionParams{
		timeout: timeout,
	}
}

// NewRepoListTagProtectionParamsWithContext creates a new RepoListTagProtectionParams object
// with the ability to set a context for a request.
func NewRepoListTagProtectionParamsWithContext(ctx context.Context) *RepoListTagProtectionParams {
	return &RepoListTagProtectionParams{
		Context: ctx,
	}
}

// NewRepoListTagProtectionParamsWithHTTPClient creates a new RepoListTagProtectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepoListTagProtectionParamsWithHTTPClient(client *http.Client) *RepoListTagProtectionParams {
	return &RepoListTagProtectionParams{
		HTTPClient: client,
	}
}

/*
RepoListTagProtectionParams contains all the parameters to send to the API endpoint

	for the repo list tag protection operation.

	Typically these are written to a http.Request.
*/
type RepoListTagProtectionParams struct {

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

// WithDefaults hydrates default values in the repo list tag protection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoListTagProtectionParams) WithDefaults() *RepoListTagProtectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repo list tag protection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoListTagProtectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repo list tag protection params
func (o *RepoListTagProtectionParams) WithTimeout(timeout time.Duration) *RepoListTagProtectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo list tag protection params
func (o *RepoListTagProtectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo list tag protection params
func (o *RepoListTagProtectionParams) WithContext(ctx context.Context) *RepoListTagProtectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo list tag protection params
func (o *RepoListTagProtectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo list tag protection params
func (o *RepoListTagProtectionParams) WithHTTPClient(client *http.Client) *RepoListTagProtectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo list tag protection params
func (o *RepoListTagProtectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the repo list tag protection params
func (o *RepoListTagProtectionParams) WithOwner(owner string) *RepoListTagProtectionParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo list tag protection params
func (o *RepoListTagProtectionParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo list tag protection params
func (o *RepoListTagProtectionParams) WithRepo(repo string) *RepoListTagProtectionParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo list tag protection params
func (o *RepoListTagProtectionParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoListTagProtectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
