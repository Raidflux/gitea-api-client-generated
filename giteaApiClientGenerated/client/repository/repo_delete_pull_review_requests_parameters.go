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

	"giteaApiClientGenerated/models"
)

// NewRepoDeletePullReviewRequestsParams creates a new RepoDeletePullReviewRequestsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepoDeletePullReviewRequestsParams() *RepoDeletePullReviewRequestsParams {
	return &RepoDeletePullReviewRequestsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepoDeletePullReviewRequestsParamsWithTimeout creates a new RepoDeletePullReviewRequestsParams object
// with the ability to set a timeout on a request.
func NewRepoDeletePullReviewRequestsParamsWithTimeout(timeout time.Duration) *RepoDeletePullReviewRequestsParams {
	return &RepoDeletePullReviewRequestsParams{
		timeout: timeout,
	}
}

// NewRepoDeletePullReviewRequestsParamsWithContext creates a new RepoDeletePullReviewRequestsParams object
// with the ability to set a context for a request.
func NewRepoDeletePullReviewRequestsParamsWithContext(ctx context.Context) *RepoDeletePullReviewRequestsParams {
	return &RepoDeletePullReviewRequestsParams{
		Context: ctx,
	}
}

// NewRepoDeletePullReviewRequestsParamsWithHTTPClient creates a new RepoDeletePullReviewRequestsParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepoDeletePullReviewRequestsParamsWithHTTPClient(client *http.Client) *RepoDeletePullReviewRequestsParams {
	return &RepoDeletePullReviewRequestsParams{
		HTTPClient: client,
	}
}

/*
RepoDeletePullReviewRequestsParams contains all the parameters to send to the API endpoint

	for the repo delete pull review requests operation.

	Typically these are written to a http.Request.
*/
type RepoDeletePullReviewRequestsParams struct {

	// Body.
	Body *models.PullReviewRequestOptions

	/* Index.

	   index of the pull request

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

// WithDefaults hydrates default values in the repo delete pull review requests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoDeletePullReviewRequestsParams) WithDefaults() *RepoDeletePullReviewRequestsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repo delete pull review requests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoDeletePullReviewRequestsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repo delete pull review requests params
func (o *RepoDeletePullReviewRequestsParams) WithTimeout(timeout time.Duration) *RepoDeletePullReviewRequestsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo delete pull review requests params
func (o *RepoDeletePullReviewRequestsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo delete pull review requests params
func (o *RepoDeletePullReviewRequestsParams) WithContext(ctx context.Context) *RepoDeletePullReviewRequestsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo delete pull review requests params
func (o *RepoDeletePullReviewRequestsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo delete pull review requests params
func (o *RepoDeletePullReviewRequestsParams) WithHTTPClient(client *http.Client) *RepoDeletePullReviewRequestsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo delete pull review requests params
func (o *RepoDeletePullReviewRequestsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the repo delete pull review requests params
func (o *RepoDeletePullReviewRequestsParams) WithBody(body *models.PullReviewRequestOptions) *RepoDeletePullReviewRequestsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the repo delete pull review requests params
func (o *RepoDeletePullReviewRequestsParams) SetBody(body *models.PullReviewRequestOptions) {
	o.Body = body
}

// WithIndex adds the index to the repo delete pull review requests params
func (o *RepoDeletePullReviewRequestsParams) WithIndex(index int64) *RepoDeletePullReviewRequestsParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the repo delete pull review requests params
func (o *RepoDeletePullReviewRequestsParams) SetIndex(index int64) {
	o.Index = index
}

// WithOwner adds the owner to the repo delete pull review requests params
func (o *RepoDeletePullReviewRequestsParams) WithOwner(owner string) *RepoDeletePullReviewRequestsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo delete pull review requests params
func (o *RepoDeletePullReviewRequestsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo delete pull review requests params
func (o *RepoDeletePullReviewRequestsParams) WithRepo(repo string) *RepoDeletePullReviewRequestsParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo delete pull review requests params
func (o *RepoDeletePullReviewRequestsParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoDeletePullReviewRequestsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
