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

// NewRepoGetCommitPullRequestParams creates a new RepoGetCommitPullRequestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepoGetCommitPullRequestParams() *RepoGetCommitPullRequestParams {
	return &RepoGetCommitPullRequestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepoGetCommitPullRequestParamsWithTimeout creates a new RepoGetCommitPullRequestParams object
// with the ability to set a timeout on a request.
func NewRepoGetCommitPullRequestParamsWithTimeout(timeout time.Duration) *RepoGetCommitPullRequestParams {
	return &RepoGetCommitPullRequestParams{
		timeout: timeout,
	}
}

// NewRepoGetCommitPullRequestParamsWithContext creates a new RepoGetCommitPullRequestParams object
// with the ability to set a context for a request.
func NewRepoGetCommitPullRequestParamsWithContext(ctx context.Context) *RepoGetCommitPullRequestParams {
	return &RepoGetCommitPullRequestParams{
		Context: ctx,
	}
}

// NewRepoGetCommitPullRequestParamsWithHTTPClient creates a new RepoGetCommitPullRequestParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepoGetCommitPullRequestParamsWithHTTPClient(client *http.Client) *RepoGetCommitPullRequestParams {
	return &RepoGetCommitPullRequestParams{
		HTTPClient: client,
	}
}

/*
RepoGetCommitPullRequestParams contains all the parameters to send to the API endpoint

	for the repo get commit pull request operation.

	Typically these are written to a http.Request.
*/
type RepoGetCommitPullRequestParams struct {

	/* Owner.

	   owner of the repo
	*/
	Owner string

	/* Repo.

	   name of the repo
	*/
	Repo string

	/* Sha.

	   SHA of the commit to get
	*/
	Sha string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the repo get commit pull request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoGetCommitPullRequestParams) WithDefaults() *RepoGetCommitPullRequestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repo get commit pull request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoGetCommitPullRequestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repo get commit pull request params
func (o *RepoGetCommitPullRequestParams) WithTimeout(timeout time.Duration) *RepoGetCommitPullRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo get commit pull request params
func (o *RepoGetCommitPullRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo get commit pull request params
func (o *RepoGetCommitPullRequestParams) WithContext(ctx context.Context) *RepoGetCommitPullRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo get commit pull request params
func (o *RepoGetCommitPullRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo get commit pull request params
func (o *RepoGetCommitPullRequestParams) WithHTTPClient(client *http.Client) *RepoGetCommitPullRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo get commit pull request params
func (o *RepoGetCommitPullRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the repo get commit pull request params
func (o *RepoGetCommitPullRequestParams) WithOwner(owner string) *RepoGetCommitPullRequestParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo get commit pull request params
func (o *RepoGetCommitPullRequestParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo get commit pull request params
func (o *RepoGetCommitPullRequestParams) WithRepo(repo string) *RepoGetCommitPullRequestParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo get commit pull request params
func (o *RepoGetCommitPullRequestParams) SetRepo(repo string) {
	o.Repo = repo
}

// WithSha adds the sha to the repo get commit pull request params
func (o *RepoGetCommitPullRequestParams) WithSha(sha string) *RepoGetCommitPullRequestParams {
	o.SetSha(sha)
	return o
}

// SetSha adds the sha to the repo get commit pull request params
func (o *RepoGetCommitPullRequestParams) SetSha(sha string) {
	o.Sha = sha
}

// WriteToRequest writes these params to a swagger request
func (o *RepoGetCommitPullRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param sha
	if err := r.SetPathParam("sha", o.Sha); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
