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
	"github.com/go-openapi/swag"
)

// NewRepoGetPullRequestParams creates a new RepoGetPullRequestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepoGetPullRequestParams() *RepoGetPullRequestParams {
	return &RepoGetPullRequestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepoGetPullRequestParamsWithTimeout creates a new RepoGetPullRequestParams object
// with the ability to set a timeout on a request.
func NewRepoGetPullRequestParamsWithTimeout(timeout time.Duration) *RepoGetPullRequestParams {
	return &RepoGetPullRequestParams{
		timeout: timeout,
	}
}

// NewRepoGetPullRequestParamsWithContext creates a new RepoGetPullRequestParams object
// with the ability to set a context for a request.
func NewRepoGetPullRequestParamsWithContext(ctx context.Context) *RepoGetPullRequestParams {
	return &RepoGetPullRequestParams{
		Context: ctx,
	}
}

// NewRepoGetPullRequestParamsWithHTTPClient creates a new RepoGetPullRequestParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepoGetPullRequestParamsWithHTTPClient(client *http.Client) *RepoGetPullRequestParams {
	return &RepoGetPullRequestParams{
		HTTPClient: client,
	}
}

/*
RepoGetPullRequestParams contains all the parameters to send to the API endpoint

	for the repo get pull request operation.

	Typically these are written to a http.Request.
*/
type RepoGetPullRequestParams struct {

	/* Index.

	   index of the pull request to get

	   Format: int64
	*/
	Index int64

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

// WithDefaults hydrates default values in the repo get pull request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoGetPullRequestParams) WithDefaults() *RepoGetPullRequestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repo get pull request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoGetPullRequestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repo get pull request params
func (o *RepoGetPullRequestParams) WithTimeout(timeout time.Duration) *RepoGetPullRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo get pull request params
func (o *RepoGetPullRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo get pull request params
func (o *RepoGetPullRequestParams) WithContext(ctx context.Context) *RepoGetPullRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo get pull request params
func (o *RepoGetPullRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo get pull request params
func (o *RepoGetPullRequestParams) WithHTTPClient(client *http.Client) *RepoGetPullRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo get pull request params
func (o *RepoGetPullRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIndex adds the index to the repo get pull request params
func (o *RepoGetPullRequestParams) WithIndex(index int64) *RepoGetPullRequestParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the repo get pull request params
func (o *RepoGetPullRequestParams) SetIndex(index int64) {
	o.Index = index
}

// WithOwner adds the owner to the repo get pull request params
func (o *RepoGetPullRequestParams) WithOwner(owner string) *RepoGetPullRequestParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo get pull request params
func (o *RepoGetPullRequestParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo get pull request params
func (o *RepoGetPullRequestParams) WithRepo(repo string) *RepoGetPullRequestParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo get pull request params
func (o *RepoGetPullRequestParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoGetPullRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param index
	if err := r.SetPathParam("index", swag.FormatInt64(o.Index)); err != nil {
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
