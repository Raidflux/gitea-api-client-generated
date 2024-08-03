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

// NewRepoListPullReviewsParams creates a new RepoListPullReviewsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepoListPullReviewsParams() *RepoListPullReviewsParams {
	return &RepoListPullReviewsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepoListPullReviewsParamsWithTimeout creates a new RepoListPullReviewsParams object
// with the ability to set a timeout on a request.
func NewRepoListPullReviewsParamsWithTimeout(timeout time.Duration) *RepoListPullReviewsParams {
	return &RepoListPullReviewsParams{
		timeout: timeout,
	}
}

// NewRepoListPullReviewsParamsWithContext creates a new RepoListPullReviewsParams object
// with the ability to set a context for a request.
func NewRepoListPullReviewsParamsWithContext(ctx context.Context) *RepoListPullReviewsParams {
	return &RepoListPullReviewsParams{
		Context: ctx,
	}
}

// NewRepoListPullReviewsParamsWithHTTPClient creates a new RepoListPullReviewsParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepoListPullReviewsParamsWithHTTPClient(client *http.Client) *RepoListPullReviewsParams {
	return &RepoListPullReviewsParams{
		HTTPClient: client,
	}
}

/*
RepoListPullReviewsParams contains all the parameters to send to the API endpoint

	for the repo list pull reviews operation.

	Typically these are written to a http.Request.
*/
type RepoListPullReviewsParams struct {

	/* Index.

	   index of the pull request

	   Format: int64
	*/
	Index int64

	/* Limit.

	   page size of results
	*/
	Limit *int64

	/* Owner.

	   owner of the repo
	*/
	Owner string

	/* Page.

	   page number of results to return (1-based)
	*/
	Page *int64

	/* Repo.

	   name of the repo
	*/
	Repo string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the repo list pull reviews params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoListPullReviewsParams) WithDefaults() *RepoListPullReviewsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repo list pull reviews params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoListPullReviewsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repo list pull reviews params
func (o *RepoListPullReviewsParams) WithTimeout(timeout time.Duration) *RepoListPullReviewsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo list pull reviews params
func (o *RepoListPullReviewsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo list pull reviews params
func (o *RepoListPullReviewsParams) WithContext(ctx context.Context) *RepoListPullReviewsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo list pull reviews params
func (o *RepoListPullReviewsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo list pull reviews params
func (o *RepoListPullReviewsParams) WithHTTPClient(client *http.Client) *RepoListPullReviewsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo list pull reviews params
func (o *RepoListPullReviewsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIndex adds the index to the repo list pull reviews params
func (o *RepoListPullReviewsParams) WithIndex(index int64) *RepoListPullReviewsParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the repo list pull reviews params
func (o *RepoListPullReviewsParams) SetIndex(index int64) {
	o.Index = index
}

// WithLimit adds the limit to the repo list pull reviews params
func (o *RepoListPullReviewsParams) WithLimit(limit *int64) *RepoListPullReviewsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the repo list pull reviews params
func (o *RepoListPullReviewsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOwner adds the owner to the repo list pull reviews params
func (o *RepoListPullReviewsParams) WithOwner(owner string) *RepoListPullReviewsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo list pull reviews params
func (o *RepoListPullReviewsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithPage adds the page to the repo list pull reviews params
func (o *RepoListPullReviewsParams) WithPage(page *int64) *RepoListPullReviewsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the repo list pull reviews params
func (o *RepoListPullReviewsParams) SetPage(page *int64) {
	o.Page = page
}

// WithRepo adds the repo to the repo list pull reviews params
func (o *RepoListPullReviewsParams) WithRepo(repo string) *RepoListPullReviewsParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo list pull reviews params
func (o *RepoListPullReviewsParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoListPullReviewsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param index
	if err := r.SetPathParam("index", swag.FormatInt64(o.Index)); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
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
