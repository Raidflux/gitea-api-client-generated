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

// NewRepoGetPullRequestCommitsParams creates a new RepoGetPullRequestCommitsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepoGetPullRequestCommitsParams() *RepoGetPullRequestCommitsParams {
	return &RepoGetPullRequestCommitsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepoGetPullRequestCommitsParamsWithTimeout creates a new RepoGetPullRequestCommitsParams object
// with the ability to set a timeout on a request.
func NewRepoGetPullRequestCommitsParamsWithTimeout(timeout time.Duration) *RepoGetPullRequestCommitsParams {
	return &RepoGetPullRequestCommitsParams{
		timeout: timeout,
	}
}

// NewRepoGetPullRequestCommitsParamsWithContext creates a new RepoGetPullRequestCommitsParams object
// with the ability to set a context for a request.
func NewRepoGetPullRequestCommitsParamsWithContext(ctx context.Context) *RepoGetPullRequestCommitsParams {
	return &RepoGetPullRequestCommitsParams{
		Context: ctx,
	}
}

// NewRepoGetPullRequestCommitsParamsWithHTTPClient creates a new RepoGetPullRequestCommitsParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepoGetPullRequestCommitsParamsWithHTTPClient(client *http.Client) *RepoGetPullRequestCommitsParams {
	return &RepoGetPullRequestCommitsParams{
		HTTPClient: client,
	}
}

/*
RepoGetPullRequestCommitsParams contains all the parameters to send to the API endpoint

	for the repo get pull request commits operation.

	Typically these are written to a http.Request.
*/
type RepoGetPullRequestCommitsParams struct {

	/* Files.

	   include a list of affected files for every commit (disable for speedup, default 'true')
	*/
	Files *bool

	/* Index.

	   index of the pull request to get

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

	/* Verification.

	   include verification for every commit (disable for speedup, default 'true')
	*/
	Verification *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the repo get pull request commits params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoGetPullRequestCommitsParams) WithDefaults() *RepoGetPullRequestCommitsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repo get pull request commits params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoGetPullRequestCommitsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repo get pull request commits params
func (o *RepoGetPullRequestCommitsParams) WithTimeout(timeout time.Duration) *RepoGetPullRequestCommitsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo get pull request commits params
func (o *RepoGetPullRequestCommitsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo get pull request commits params
func (o *RepoGetPullRequestCommitsParams) WithContext(ctx context.Context) *RepoGetPullRequestCommitsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo get pull request commits params
func (o *RepoGetPullRequestCommitsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo get pull request commits params
func (o *RepoGetPullRequestCommitsParams) WithHTTPClient(client *http.Client) *RepoGetPullRequestCommitsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo get pull request commits params
func (o *RepoGetPullRequestCommitsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFiles adds the files to the repo get pull request commits params
func (o *RepoGetPullRequestCommitsParams) WithFiles(files *bool) *RepoGetPullRequestCommitsParams {
	o.SetFiles(files)
	return o
}

// SetFiles adds the files to the repo get pull request commits params
func (o *RepoGetPullRequestCommitsParams) SetFiles(files *bool) {
	o.Files = files
}

// WithIndex adds the index to the repo get pull request commits params
func (o *RepoGetPullRequestCommitsParams) WithIndex(index int64) *RepoGetPullRequestCommitsParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the repo get pull request commits params
func (o *RepoGetPullRequestCommitsParams) SetIndex(index int64) {
	o.Index = index
}

// WithLimit adds the limit to the repo get pull request commits params
func (o *RepoGetPullRequestCommitsParams) WithLimit(limit *int64) *RepoGetPullRequestCommitsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the repo get pull request commits params
func (o *RepoGetPullRequestCommitsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOwner adds the owner to the repo get pull request commits params
func (o *RepoGetPullRequestCommitsParams) WithOwner(owner string) *RepoGetPullRequestCommitsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo get pull request commits params
func (o *RepoGetPullRequestCommitsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithPage adds the page to the repo get pull request commits params
func (o *RepoGetPullRequestCommitsParams) WithPage(page *int64) *RepoGetPullRequestCommitsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the repo get pull request commits params
func (o *RepoGetPullRequestCommitsParams) SetPage(page *int64) {
	o.Page = page
}

// WithRepo adds the repo to the repo get pull request commits params
func (o *RepoGetPullRequestCommitsParams) WithRepo(repo string) *RepoGetPullRequestCommitsParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo get pull request commits params
func (o *RepoGetPullRequestCommitsParams) SetRepo(repo string) {
	o.Repo = repo
}

// WithVerification adds the verification to the repo get pull request commits params
func (o *RepoGetPullRequestCommitsParams) WithVerification(verification *bool) *RepoGetPullRequestCommitsParams {
	o.SetVerification(verification)
	return o
}

// SetVerification adds the verification to the repo get pull request commits params
func (o *RepoGetPullRequestCommitsParams) SetVerification(verification *bool) {
	o.Verification = verification
}

// WriteToRequest writes these params to a swagger request
func (o *RepoGetPullRequestCommitsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Files != nil {

		// query param files
		var qrFiles bool

		if o.Files != nil {
			qrFiles = *o.Files
		}
		qFiles := swag.FormatBool(qrFiles)
		if qFiles != "" {

			if err := r.SetQueryParam("files", qFiles); err != nil {
				return err
			}
		}
	}

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

	if o.Verification != nil {

		// query param verification
		var qrVerification bool

		if o.Verification != nil {
			qrVerification = *o.Verification
		}
		qVerification := swag.FormatBool(qrVerification)
		if qVerification != "" {

			if err := r.SetQueryParam("verification", qVerification); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
