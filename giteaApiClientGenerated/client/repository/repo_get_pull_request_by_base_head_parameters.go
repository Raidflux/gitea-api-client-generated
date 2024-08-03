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

// NewRepoGetPullRequestByBaseHeadParams creates a new RepoGetPullRequestByBaseHeadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepoGetPullRequestByBaseHeadParams() *RepoGetPullRequestByBaseHeadParams {
	return &RepoGetPullRequestByBaseHeadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepoGetPullRequestByBaseHeadParamsWithTimeout creates a new RepoGetPullRequestByBaseHeadParams object
// with the ability to set a timeout on a request.
func NewRepoGetPullRequestByBaseHeadParamsWithTimeout(timeout time.Duration) *RepoGetPullRequestByBaseHeadParams {
	return &RepoGetPullRequestByBaseHeadParams{
		timeout: timeout,
	}
}

// NewRepoGetPullRequestByBaseHeadParamsWithContext creates a new RepoGetPullRequestByBaseHeadParams object
// with the ability to set a context for a request.
func NewRepoGetPullRequestByBaseHeadParamsWithContext(ctx context.Context) *RepoGetPullRequestByBaseHeadParams {
	return &RepoGetPullRequestByBaseHeadParams{
		Context: ctx,
	}
}

// NewRepoGetPullRequestByBaseHeadParamsWithHTTPClient creates a new RepoGetPullRequestByBaseHeadParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepoGetPullRequestByBaseHeadParamsWithHTTPClient(client *http.Client) *RepoGetPullRequestByBaseHeadParams {
	return &RepoGetPullRequestByBaseHeadParams{
		HTTPClient: client,
	}
}

/*
RepoGetPullRequestByBaseHeadParams contains all the parameters to send to the API endpoint

	for the repo get pull request by base head operation.

	Typically these are written to a http.Request.
*/
type RepoGetPullRequestByBaseHeadParams struct {

	/* Base.

	   base of the pull request to get
	*/
	Base string

	/* Head.

	   head of the pull request to get
	*/
	Head string

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

// WithDefaults hydrates default values in the repo get pull request by base head params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoGetPullRequestByBaseHeadParams) WithDefaults() *RepoGetPullRequestByBaseHeadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repo get pull request by base head params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoGetPullRequestByBaseHeadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repo get pull request by base head params
func (o *RepoGetPullRequestByBaseHeadParams) WithTimeout(timeout time.Duration) *RepoGetPullRequestByBaseHeadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo get pull request by base head params
func (o *RepoGetPullRequestByBaseHeadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo get pull request by base head params
func (o *RepoGetPullRequestByBaseHeadParams) WithContext(ctx context.Context) *RepoGetPullRequestByBaseHeadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo get pull request by base head params
func (o *RepoGetPullRequestByBaseHeadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo get pull request by base head params
func (o *RepoGetPullRequestByBaseHeadParams) WithHTTPClient(client *http.Client) *RepoGetPullRequestByBaseHeadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo get pull request by base head params
func (o *RepoGetPullRequestByBaseHeadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBase adds the base to the repo get pull request by base head params
func (o *RepoGetPullRequestByBaseHeadParams) WithBase(base string) *RepoGetPullRequestByBaseHeadParams {
	o.SetBase(base)
	return o
}

// SetBase adds the base to the repo get pull request by base head params
func (o *RepoGetPullRequestByBaseHeadParams) SetBase(base string) {
	o.Base = base
}

// WithHead adds the head to the repo get pull request by base head params
func (o *RepoGetPullRequestByBaseHeadParams) WithHead(head string) *RepoGetPullRequestByBaseHeadParams {
	o.SetHead(head)
	return o
}

// SetHead adds the head to the repo get pull request by base head params
func (o *RepoGetPullRequestByBaseHeadParams) SetHead(head string) {
	o.Head = head
}

// WithOwner adds the owner to the repo get pull request by base head params
func (o *RepoGetPullRequestByBaseHeadParams) WithOwner(owner string) *RepoGetPullRequestByBaseHeadParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo get pull request by base head params
func (o *RepoGetPullRequestByBaseHeadParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo get pull request by base head params
func (o *RepoGetPullRequestByBaseHeadParams) WithRepo(repo string) *RepoGetPullRequestByBaseHeadParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo get pull request by base head params
func (o *RepoGetPullRequestByBaseHeadParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoGetPullRequestByBaseHeadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param base
	if err := r.SetPathParam("base", o.Base); err != nil {
		return err
	}

	// path param head
	if err := r.SetPathParam("head", o.Head); err != nil {
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
