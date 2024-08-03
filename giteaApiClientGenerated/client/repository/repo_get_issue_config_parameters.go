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

// NewRepoGetIssueConfigParams creates a new RepoGetIssueConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepoGetIssueConfigParams() *RepoGetIssueConfigParams {
	return &RepoGetIssueConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepoGetIssueConfigParamsWithTimeout creates a new RepoGetIssueConfigParams object
// with the ability to set a timeout on a request.
func NewRepoGetIssueConfigParamsWithTimeout(timeout time.Duration) *RepoGetIssueConfigParams {
	return &RepoGetIssueConfigParams{
		timeout: timeout,
	}
}

// NewRepoGetIssueConfigParamsWithContext creates a new RepoGetIssueConfigParams object
// with the ability to set a context for a request.
func NewRepoGetIssueConfigParamsWithContext(ctx context.Context) *RepoGetIssueConfigParams {
	return &RepoGetIssueConfigParams{
		Context: ctx,
	}
}

// NewRepoGetIssueConfigParamsWithHTTPClient creates a new RepoGetIssueConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepoGetIssueConfigParamsWithHTTPClient(client *http.Client) *RepoGetIssueConfigParams {
	return &RepoGetIssueConfigParams{
		HTTPClient: client,
	}
}

/*
RepoGetIssueConfigParams contains all the parameters to send to the API endpoint

	for the repo get issue config operation.

	Typically these are written to a http.Request.
*/
type RepoGetIssueConfigParams struct {

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

// WithDefaults hydrates default values in the repo get issue config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoGetIssueConfigParams) WithDefaults() *RepoGetIssueConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repo get issue config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoGetIssueConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repo get issue config params
func (o *RepoGetIssueConfigParams) WithTimeout(timeout time.Duration) *RepoGetIssueConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo get issue config params
func (o *RepoGetIssueConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo get issue config params
func (o *RepoGetIssueConfigParams) WithContext(ctx context.Context) *RepoGetIssueConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo get issue config params
func (o *RepoGetIssueConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo get issue config params
func (o *RepoGetIssueConfigParams) WithHTTPClient(client *http.Client) *RepoGetIssueConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo get issue config params
func (o *RepoGetIssueConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the repo get issue config params
func (o *RepoGetIssueConfigParams) WithOwner(owner string) *RepoGetIssueConfigParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo get issue config params
func (o *RepoGetIssueConfigParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo get issue config params
func (o *RepoGetIssueConfigParams) WithRepo(repo string) *RepoGetIssueConfigParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo get issue config params
func (o *RepoGetIssueConfigParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoGetIssueConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
