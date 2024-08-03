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

// NewRepoGetNoteParams creates a new RepoGetNoteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepoGetNoteParams() *RepoGetNoteParams {
	return &RepoGetNoteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepoGetNoteParamsWithTimeout creates a new RepoGetNoteParams object
// with the ability to set a timeout on a request.
func NewRepoGetNoteParamsWithTimeout(timeout time.Duration) *RepoGetNoteParams {
	return &RepoGetNoteParams{
		timeout: timeout,
	}
}

// NewRepoGetNoteParamsWithContext creates a new RepoGetNoteParams object
// with the ability to set a context for a request.
func NewRepoGetNoteParamsWithContext(ctx context.Context) *RepoGetNoteParams {
	return &RepoGetNoteParams{
		Context: ctx,
	}
}

// NewRepoGetNoteParamsWithHTTPClient creates a new RepoGetNoteParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepoGetNoteParamsWithHTTPClient(client *http.Client) *RepoGetNoteParams {
	return &RepoGetNoteParams{
		HTTPClient: client,
	}
}

/*
RepoGetNoteParams contains all the parameters to send to the API endpoint

	for the repo get note operation.

	Typically these are written to a http.Request.
*/
type RepoGetNoteParams struct {

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

	/* Verification.

	   include verification for every commit (disable for speedup, default 'true')
	*/
	Verification *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the repo get note params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoGetNoteParams) WithDefaults() *RepoGetNoteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repo get note params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoGetNoteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repo get note params
func (o *RepoGetNoteParams) WithTimeout(timeout time.Duration) *RepoGetNoteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo get note params
func (o *RepoGetNoteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo get note params
func (o *RepoGetNoteParams) WithContext(ctx context.Context) *RepoGetNoteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo get note params
func (o *RepoGetNoteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo get note params
func (o *RepoGetNoteParams) WithHTTPClient(client *http.Client) *RepoGetNoteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo get note params
func (o *RepoGetNoteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFiles adds the files to the repo get note params
func (o *RepoGetNoteParams) WithFiles(files *bool) *RepoGetNoteParams {
	o.SetFiles(files)
	return o
}

// SetFiles adds the files to the repo get note params
func (o *RepoGetNoteParams) SetFiles(files *bool) {
	o.Files = files
}

// WithOwner adds the owner to the repo get note params
func (o *RepoGetNoteParams) WithOwner(owner string) *RepoGetNoteParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo get note params
func (o *RepoGetNoteParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo get note params
func (o *RepoGetNoteParams) WithRepo(repo string) *RepoGetNoteParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo get note params
func (o *RepoGetNoteParams) SetRepo(repo string) {
	o.Repo = repo
}

// WithSha adds the sha to the repo get note params
func (o *RepoGetNoteParams) WithSha(sha string) *RepoGetNoteParams {
	o.SetSha(sha)
	return o
}

// SetSha adds the sha to the repo get note params
func (o *RepoGetNoteParams) SetSha(sha string) {
	o.Sha = sha
}

// WithVerification adds the verification to the repo get note params
func (o *RepoGetNoteParams) WithVerification(verification *bool) *RepoGetNoteParams {
	o.SetVerification(verification)
	return o
}

// SetVerification adds the verification to the repo get note params
func (o *RepoGetNoteParams) SetVerification(verification *bool) {
	o.Verification = verification
}

// WriteToRequest writes these params to a swagger request
func (o *RepoGetNoteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
