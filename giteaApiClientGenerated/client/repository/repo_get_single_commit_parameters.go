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

// NewRepoGetSingleCommitParams creates a new RepoGetSingleCommitParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepoGetSingleCommitParams() *RepoGetSingleCommitParams {
	return &RepoGetSingleCommitParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepoGetSingleCommitParamsWithTimeout creates a new RepoGetSingleCommitParams object
// with the ability to set a timeout on a request.
func NewRepoGetSingleCommitParamsWithTimeout(timeout time.Duration) *RepoGetSingleCommitParams {
	return &RepoGetSingleCommitParams{
		timeout: timeout,
	}
}

// NewRepoGetSingleCommitParamsWithContext creates a new RepoGetSingleCommitParams object
// with the ability to set a context for a request.
func NewRepoGetSingleCommitParamsWithContext(ctx context.Context) *RepoGetSingleCommitParams {
	return &RepoGetSingleCommitParams{
		Context: ctx,
	}
}

// NewRepoGetSingleCommitParamsWithHTTPClient creates a new RepoGetSingleCommitParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepoGetSingleCommitParamsWithHTTPClient(client *http.Client) *RepoGetSingleCommitParams {
	return &RepoGetSingleCommitParams{
		HTTPClient: client,
	}
}

/*
RepoGetSingleCommitParams contains all the parameters to send to the API endpoint

	for the repo get single commit operation.

	Typically these are written to a http.Request.
*/
type RepoGetSingleCommitParams struct {

	/* Files.

	   include a list of affected files for every commit (disable for speedup, default 'true')
	*/
	Files *bool

	/* Owner.

	   owner of the repo
	*/
	Owner string

	/* Repo.

	   name of the repo
	*/
	Repo string

	/* Sha.

	   a git ref or commit sha
	*/
	Sha string

	/* Stat.

	   include diff stats for every commit (disable for speedup, default 'true')
	*/
	Stat *bool

	/* Verification.

	   include verification for every commit (disable for speedup, default 'true')
	*/
	Verification *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the repo get single commit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoGetSingleCommitParams) WithDefaults() *RepoGetSingleCommitParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repo get single commit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoGetSingleCommitParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repo get single commit params
func (o *RepoGetSingleCommitParams) WithTimeout(timeout time.Duration) *RepoGetSingleCommitParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo get single commit params
func (o *RepoGetSingleCommitParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo get single commit params
func (o *RepoGetSingleCommitParams) WithContext(ctx context.Context) *RepoGetSingleCommitParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo get single commit params
func (o *RepoGetSingleCommitParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo get single commit params
func (o *RepoGetSingleCommitParams) WithHTTPClient(client *http.Client) *RepoGetSingleCommitParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo get single commit params
func (o *RepoGetSingleCommitParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFiles adds the files to the repo get single commit params
func (o *RepoGetSingleCommitParams) WithFiles(files *bool) *RepoGetSingleCommitParams {
	o.SetFiles(files)
	return o
}

// SetFiles adds the files to the repo get single commit params
func (o *RepoGetSingleCommitParams) SetFiles(files *bool) {
	o.Files = files
}

// WithOwner adds the owner to the repo get single commit params
func (o *RepoGetSingleCommitParams) WithOwner(owner string) *RepoGetSingleCommitParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo get single commit params
func (o *RepoGetSingleCommitParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo get single commit params
func (o *RepoGetSingleCommitParams) WithRepo(repo string) *RepoGetSingleCommitParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo get single commit params
func (o *RepoGetSingleCommitParams) SetRepo(repo string) {
	o.Repo = repo
}

// WithSha adds the sha to the repo get single commit params
func (o *RepoGetSingleCommitParams) WithSha(sha string) *RepoGetSingleCommitParams {
	o.SetSha(sha)
	return o
}

// SetSha adds the sha to the repo get single commit params
func (o *RepoGetSingleCommitParams) SetSha(sha string) {
	o.Sha = sha
}

// WithStat adds the stat to the repo get single commit params
func (o *RepoGetSingleCommitParams) WithStat(stat *bool) *RepoGetSingleCommitParams {
	o.SetStat(stat)
	return o
}

// SetStat adds the stat to the repo get single commit params
func (o *RepoGetSingleCommitParams) SetStat(stat *bool) {
	o.Stat = stat
}

// WithVerification adds the verification to the repo get single commit params
func (o *RepoGetSingleCommitParams) WithVerification(verification *bool) *RepoGetSingleCommitParams {
	o.SetVerification(verification)
	return o
}

// SetVerification adds the verification to the repo get single commit params
func (o *RepoGetSingleCommitParams) SetVerification(verification *bool) {
	o.Verification = verification
}

// WriteToRequest writes these params to a swagger request
func (o *RepoGetSingleCommitParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Stat != nil {

		// query param stat
		var qrStat bool

		if o.Stat != nil {
			qrStat = *o.Stat
		}
		qStat := swag.FormatBool(qrStat)
		if qStat != "" {

			if err := r.SetQueryParam("stat", qStat); err != nil {
				return err
			}
		}
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
