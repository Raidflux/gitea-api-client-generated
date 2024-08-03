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

// NewRepoGetEditorConfigParams creates a new RepoGetEditorConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepoGetEditorConfigParams() *RepoGetEditorConfigParams {
	return &RepoGetEditorConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepoGetEditorConfigParamsWithTimeout creates a new RepoGetEditorConfigParams object
// with the ability to set a timeout on a request.
func NewRepoGetEditorConfigParamsWithTimeout(timeout time.Duration) *RepoGetEditorConfigParams {
	return &RepoGetEditorConfigParams{
		timeout: timeout,
	}
}

// NewRepoGetEditorConfigParamsWithContext creates a new RepoGetEditorConfigParams object
// with the ability to set a context for a request.
func NewRepoGetEditorConfigParamsWithContext(ctx context.Context) *RepoGetEditorConfigParams {
	return &RepoGetEditorConfigParams{
		Context: ctx,
	}
}

// NewRepoGetEditorConfigParamsWithHTTPClient creates a new RepoGetEditorConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepoGetEditorConfigParamsWithHTTPClient(client *http.Client) *RepoGetEditorConfigParams {
	return &RepoGetEditorConfigParams{
		HTTPClient: client,
	}
}

/*
RepoGetEditorConfigParams contains all the parameters to send to the API endpoint

	for the repo get editor config operation.

	Typically these are written to a http.Request.
*/
type RepoGetEditorConfigParams struct {

	/* Filepath.

	   filepath of file to get
	*/
	Filepath string

	/* Owner.

	   owner of the repo
	*/
	Owner string

	/* Ref.

	   The name of the commit/branch/tag. Default the repository’s default branch (usually master)
	*/
	Ref *string

	/* Repo.

	   name of the repo
	*/
	Repo string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the repo get editor config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoGetEditorConfigParams) WithDefaults() *RepoGetEditorConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repo get editor config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoGetEditorConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repo get editor config params
func (o *RepoGetEditorConfigParams) WithTimeout(timeout time.Duration) *RepoGetEditorConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo get editor config params
func (o *RepoGetEditorConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo get editor config params
func (o *RepoGetEditorConfigParams) WithContext(ctx context.Context) *RepoGetEditorConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo get editor config params
func (o *RepoGetEditorConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo get editor config params
func (o *RepoGetEditorConfigParams) WithHTTPClient(client *http.Client) *RepoGetEditorConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo get editor config params
func (o *RepoGetEditorConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilepath adds the filepath to the repo get editor config params
func (o *RepoGetEditorConfigParams) WithFilepath(filepath string) *RepoGetEditorConfigParams {
	o.SetFilepath(filepath)
	return o
}

// SetFilepath adds the filepath to the repo get editor config params
func (o *RepoGetEditorConfigParams) SetFilepath(filepath string) {
	o.Filepath = filepath
}

// WithOwner adds the owner to the repo get editor config params
func (o *RepoGetEditorConfigParams) WithOwner(owner string) *RepoGetEditorConfigParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo get editor config params
func (o *RepoGetEditorConfigParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRef adds the ref to the repo get editor config params
func (o *RepoGetEditorConfigParams) WithRef(ref *string) *RepoGetEditorConfigParams {
	o.SetRef(ref)
	return o
}

// SetRef adds the ref to the repo get editor config params
func (o *RepoGetEditorConfigParams) SetRef(ref *string) {
	o.Ref = ref
}

// WithRepo adds the repo to the repo get editor config params
func (o *RepoGetEditorConfigParams) WithRepo(repo string) *RepoGetEditorConfigParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo get editor config params
func (o *RepoGetEditorConfigParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoGetEditorConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param filepath
	if err := r.SetPathParam("filepath", o.Filepath); err != nil {
		return err
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	if o.Ref != nil {

		// query param ref
		var qrRef string

		if o.Ref != nil {
			qrRef = *o.Ref
		}
		qRef := qrRef
		if qRef != "" {

			if err := r.SetQueryParam("ref", qRef); err != nil {
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
