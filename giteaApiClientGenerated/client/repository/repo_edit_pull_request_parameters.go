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

	"github.com/Raidflux/gitea-api-client-generated/giteaApiClientGenerated/models"
)

// NewRepoEditPullRequestParams creates a new RepoEditPullRequestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepoEditPullRequestParams() *RepoEditPullRequestParams {
	return &RepoEditPullRequestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepoEditPullRequestParamsWithTimeout creates a new RepoEditPullRequestParams object
// with the ability to set a timeout on a request.
func NewRepoEditPullRequestParamsWithTimeout(timeout time.Duration) *RepoEditPullRequestParams {
	return &RepoEditPullRequestParams{
		timeout: timeout,
	}
}

// NewRepoEditPullRequestParamsWithContext creates a new RepoEditPullRequestParams object
// with the ability to set a context for a request.
func NewRepoEditPullRequestParamsWithContext(ctx context.Context) *RepoEditPullRequestParams {
	return &RepoEditPullRequestParams{
		Context: ctx,
	}
}

// NewRepoEditPullRequestParamsWithHTTPClient creates a new RepoEditPullRequestParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepoEditPullRequestParamsWithHTTPClient(client *http.Client) *RepoEditPullRequestParams {
	return &RepoEditPullRequestParams{
		HTTPClient: client,
	}
}

/*
RepoEditPullRequestParams contains all the parameters to send to the API endpoint

	for the repo edit pull request operation.

	Typically these are written to a http.Request.
*/
type RepoEditPullRequestParams struct {

	// Body.
	Body *models.EditPullRequestOption

	/* Index.

	   index of the pull request to edit

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

// WithDefaults hydrates default values in the repo edit pull request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoEditPullRequestParams) WithDefaults() *RepoEditPullRequestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repo edit pull request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoEditPullRequestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repo edit pull request params
func (o *RepoEditPullRequestParams) WithTimeout(timeout time.Duration) *RepoEditPullRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo edit pull request params
func (o *RepoEditPullRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo edit pull request params
func (o *RepoEditPullRequestParams) WithContext(ctx context.Context) *RepoEditPullRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo edit pull request params
func (o *RepoEditPullRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo edit pull request params
func (o *RepoEditPullRequestParams) WithHTTPClient(client *http.Client) *RepoEditPullRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo edit pull request params
func (o *RepoEditPullRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the repo edit pull request params
func (o *RepoEditPullRequestParams) WithBody(body *models.EditPullRequestOption) *RepoEditPullRequestParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the repo edit pull request params
func (o *RepoEditPullRequestParams) SetBody(body *models.EditPullRequestOption) {
	o.Body = body
}

// WithIndex adds the index to the repo edit pull request params
func (o *RepoEditPullRequestParams) WithIndex(index int64) *RepoEditPullRequestParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the repo edit pull request params
func (o *RepoEditPullRequestParams) SetIndex(index int64) {
	o.Index = index
}

// WithOwner adds the owner to the repo edit pull request params
func (o *RepoEditPullRequestParams) WithOwner(owner string) *RepoEditPullRequestParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo edit pull request params
func (o *RepoEditPullRequestParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo edit pull request params
func (o *RepoEditPullRequestParams) WithRepo(repo string) *RepoEditPullRequestParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo edit pull request params
func (o *RepoEditPullRequestParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoEditPullRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
