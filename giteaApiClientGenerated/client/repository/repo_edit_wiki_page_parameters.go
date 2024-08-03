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

	"giteaApiClientGenerated/models"
)

// NewRepoEditWikiPageParams creates a new RepoEditWikiPageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepoEditWikiPageParams() *RepoEditWikiPageParams {
	return &RepoEditWikiPageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepoEditWikiPageParamsWithTimeout creates a new RepoEditWikiPageParams object
// with the ability to set a timeout on a request.
func NewRepoEditWikiPageParamsWithTimeout(timeout time.Duration) *RepoEditWikiPageParams {
	return &RepoEditWikiPageParams{
		timeout: timeout,
	}
}

// NewRepoEditWikiPageParamsWithContext creates a new RepoEditWikiPageParams object
// with the ability to set a context for a request.
func NewRepoEditWikiPageParamsWithContext(ctx context.Context) *RepoEditWikiPageParams {
	return &RepoEditWikiPageParams{
		Context: ctx,
	}
}

// NewRepoEditWikiPageParamsWithHTTPClient creates a new RepoEditWikiPageParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepoEditWikiPageParamsWithHTTPClient(client *http.Client) *RepoEditWikiPageParams {
	return &RepoEditWikiPageParams{
		HTTPClient: client,
	}
}

/*
RepoEditWikiPageParams contains all the parameters to send to the API endpoint

	for the repo edit wiki page operation.

	Typically these are written to a http.Request.
*/
type RepoEditWikiPageParams struct {

	// Body.
	Body *models.CreateWikiPageOptions

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

// WithDefaults hydrates default values in the repo edit wiki page params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoEditWikiPageParams) WithDefaults() *RepoEditWikiPageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repo edit wiki page params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoEditWikiPageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repo edit wiki page params
func (o *RepoEditWikiPageParams) WithTimeout(timeout time.Duration) *RepoEditWikiPageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo edit wiki page params
func (o *RepoEditWikiPageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo edit wiki page params
func (o *RepoEditWikiPageParams) WithContext(ctx context.Context) *RepoEditWikiPageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo edit wiki page params
func (o *RepoEditWikiPageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo edit wiki page params
func (o *RepoEditWikiPageParams) WithHTTPClient(client *http.Client) *RepoEditWikiPageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo edit wiki page params
func (o *RepoEditWikiPageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the repo edit wiki page params
func (o *RepoEditWikiPageParams) WithBody(body *models.CreateWikiPageOptions) *RepoEditWikiPageParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the repo edit wiki page params
func (o *RepoEditWikiPageParams) SetBody(body *models.CreateWikiPageOptions) {
	o.Body = body
}

// WithOwner adds the owner to the repo edit wiki page params
func (o *RepoEditWikiPageParams) WithOwner(owner string) *RepoEditWikiPageParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo edit wiki page params
func (o *RepoEditWikiPageParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithPageName adds the pageName to the repo edit wiki page params
func (o *RepoEditWikiPageParams) WithPageName(pageName string) *RepoEditWikiPageParams {
	o.SetPageName(pageName)
	return o
}

// SetPageName adds the pageName to the repo edit wiki page params
func (o *RepoEditWikiPageParams) SetPageName(pageName string) {
	o.PageName = pageName
}

// WithRepo adds the repo to the repo edit wiki page params
func (o *RepoEditWikiPageParams) WithRepo(repo string) *RepoEditWikiPageParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo edit wiki page params
func (o *RepoEditWikiPageParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoEditWikiPageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
