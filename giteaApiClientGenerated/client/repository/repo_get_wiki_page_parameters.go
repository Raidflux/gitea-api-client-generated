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

// NewRepoGetWikiPageParams creates a new RepoGetWikiPageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepoGetWikiPageParams() *RepoGetWikiPageParams {
	return &RepoGetWikiPageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepoGetWikiPageParamsWithTimeout creates a new RepoGetWikiPageParams object
// with the ability to set a timeout on a request.
func NewRepoGetWikiPageParamsWithTimeout(timeout time.Duration) *RepoGetWikiPageParams {
	return &RepoGetWikiPageParams{
		timeout: timeout,
	}
}

// NewRepoGetWikiPageParamsWithContext creates a new RepoGetWikiPageParams object
// with the ability to set a context for a request.
func NewRepoGetWikiPageParamsWithContext(ctx context.Context) *RepoGetWikiPageParams {
	return &RepoGetWikiPageParams{
		Context: ctx,
	}
}

// NewRepoGetWikiPageParamsWithHTTPClient creates a new RepoGetWikiPageParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepoGetWikiPageParamsWithHTTPClient(client *http.Client) *RepoGetWikiPageParams {
	return &RepoGetWikiPageParams{
		HTTPClient: client,
	}
}

/*
RepoGetWikiPageParams contains all the parameters to send to the API endpoint

	for the repo get wiki page operation.

	Typically these are written to a http.Request.
*/
type RepoGetWikiPageParams struct {

	/* Owner.

	   owner of the repo
	*/
	Owner string

	/* PageName.

	   name of the page
	*/
	PageName string

	/* Repo.

	   name of the repo
	*/
	Repo string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the repo get wiki page params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoGetWikiPageParams) WithDefaults() *RepoGetWikiPageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repo get wiki page params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoGetWikiPageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repo get wiki page params
func (o *RepoGetWikiPageParams) WithTimeout(timeout time.Duration) *RepoGetWikiPageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo get wiki page params
func (o *RepoGetWikiPageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo get wiki page params
func (o *RepoGetWikiPageParams) WithContext(ctx context.Context) *RepoGetWikiPageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo get wiki page params
func (o *RepoGetWikiPageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo get wiki page params
func (o *RepoGetWikiPageParams) WithHTTPClient(client *http.Client) *RepoGetWikiPageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo get wiki page params
func (o *RepoGetWikiPageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the repo get wiki page params
func (o *RepoGetWikiPageParams) WithOwner(owner string) *RepoGetWikiPageParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo get wiki page params
func (o *RepoGetWikiPageParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithPageName adds the pageName to the repo get wiki page params
func (o *RepoGetWikiPageParams) WithPageName(pageName string) *RepoGetWikiPageParams {
	o.SetPageName(pageName)
	return o
}

// SetPageName adds the pageName to the repo get wiki page params
func (o *RepoGetWikiPageParams) SetPageName(pageName string) {
	o.PageName = pageName
}

// WithRepo adds the repo to the repo get wiki page params
func (o *RepoGetWikiPageParams) WithRepo(repo string) *RepoGetWikiPageParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo get wiki page params
func (o *RepoGetWikiPageParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoGetWikiPageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param pageName
	if err := r.SetPathParam("pageName", o.PageName); err != nil {
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
