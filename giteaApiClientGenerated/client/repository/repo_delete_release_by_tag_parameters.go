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

// NewRepoDeleteReleaseByTagParams creates a new RepoDeleteReleaseByTagParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepoDeleteReleaseByTagParams() *RepoDeleteReleaseByTagParams {
	return &RepoDeleteReleaseByTagParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepoDeleteReleaseByTagParamsWithTimeout creates a new RepoDeleteReleaseByTagParams object
// with the ability to set a timeout on a request.
func NewRepoDeleteReleaseByTagParamsWithTimeout(timeout time.Duration) *RepoDeleteReleaseByTagParams {
	return &RepoDeleteReleaseByTagParams{
		timeout: timeout,
	}
}

// NewRepoDeleteReleaseByTagParamsWithContext creates a new RepoDeleteReleaseByTagParams object
// with the ability to set a context for a request.
func NewRepoDeleteReleaseByTagParamsWithContext(ctx context.Context) *RepoDeleteReleaseByTagParams {
	return &RepoDeleteReleaseByTagParams{
		Context: ctx,
	}
}

// NewRepoDeleteReleaseByTagParamsWithHTTPClient creates a new RepoDeleteReleaseByTagParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepoDeleteReleaseByTagParamsWithHTTPClient(client *http.Client) *RepoDeleteReleaseByTagParams {
	return &RepoDeleteReleaseByTagParams{
		HTTPClient: client,
	}
}

/*
RepoDeleteReleaseByTagParams contains all the parameters to send to the API endpoint

	for the repo delete release by tag operation.

	Typically these are written to a http.Request.
*/
type RepoDeleteReleaseByTagParams struct {

	/* Owner.

	   owner of the repo
	*/
	Owner string

	/* Repo.

	   name of the repo
	*/
	Repo string

	/* Tag.

	   tag name of the release to delete
	*/
	Tag string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the repo delete release by tag params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoDeleteReleaseByTagParams) WithDefaults() *RepoDeleteReleaseByTagParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repo delete release by tag params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoDeleteReleaseByTagParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repo delete release by tag params
func (o *RepoDeleteReleaseByTagParams) WithTimeout(timeout time.Duration) *RepoDeleteReleaseByTagParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo delete release by tag params
func (o *RepoDeleteReleaseByTagParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo delete release by tag params
func (o *RepoDeleteReleaseByTagParams) WithContext(ctx context.Context) *RepoDeleteReleaseByTagParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo delete release by tag params
func (o *RepoDeleteReleaseByTagParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo delete release by tag params
func (o *RepoDeleteReleaseByTagParams) WithHTTPClient(client *http.Client) *RepoDeleteReleaseByTagParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo delete release by tag params
func (o *RepoDeleteReleaseByTagParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the repo delete release by tag params
func (o *RepoDeleteReleaseByTagParams) WithOwner(owner string) *RepoDeleteReleaseByTagParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo delete release by tag params
func (o *RepoDeleteReleaseByTagParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo delete release by tag params
func (o *RepoDeleteReleaseByTagParams) WithRepo(repo string) *RepoDeleteReleaseByTagParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo delete release by tag params
func (o *RepoDeleteReleaseByTagParams) SetRepo(repo string) {
	o.Repo = repo
}

// WithTag adds the tag to the repo delete release by tag params
func (o *RepoDeleteReleaseByTagParams) WithTag(tag string) *RepoDeleteReleaseByTagParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the repo delete release by tag params
func (o *RepoDeleteReleaseByTagParams) SetTag(tag string) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *RepoDeleteReleaseByTagParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param tag
	if err := r.SetPathParam("tag", o.Tag); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
