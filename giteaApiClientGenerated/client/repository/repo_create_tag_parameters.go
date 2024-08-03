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

// NewRepoCreateTagParams creates a new RepoCreateTagParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepoCreateTagParams() *RepoCreateTagParams {
	return &RepoCreateTagParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepoCreateTagParamsWithTimeout creates a new RepoCreateTagParams object
// with the ability to set a timeout on a request.
func NewRepoCreateTagParamsWithTimeout(timeout time.Duration) *RepoCreateTagParams {
	return &RepoCreateTagParams{
		timeout: timeout,
	}
}

// NewRepoCreateTagParamsWithContext creates a new RepoCreateTagParams object
// with the ability to set a context for a request.
func NewRepoCreateTagParamsWithContext(ctx context.Context) *RepoCreateTagParams {
	return &RepoCreateTagParams{
		Context: ctx,
	}
}

// NewRepoCreateTagParamsWithHTTPClient creates a new RepoCreateTagParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepoCreateTagParamsWithHTTPClient(client *http.Client) *RepoCreateTagParams {
	return &RepoCreateTagParams{
		HTTPClient: client,
	}
}

/*
RepoCreateTagParams contains all the parameters to send to the API endpoint

	for the repo create tag operation.

	Typically these are written to a http.Request.
*/
type RepoCreateTagParams struct {

	// Body.
	Body *models.CreateTagOption

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

// WithDefaults hydrates default values in the repo create tag params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoCreateTagParams) WithDefaults() *RepoCreateTagParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repo create tag params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoCreateTagParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repo create tag params
func (o *RepoCreateTagParams) WithTimeout(timeout time.Duration) *RepoCreateTagParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo create tag params
func (o *RepoCreateTagParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo create tag params
func (o *RepoCreateTagParams) WithContext(ctx context.Context) *RepoCreateTagParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo create tag params
func (o *RepoCreateTagParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo create tag params
func (o *RepoCreateTagParams) WithHTTPClient(client *http.Client) *RepoCreateTagParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo create tag params
func (o *RepoCreateTagParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the repo create tag params
func (o *RepoCreateTagParams) WithBody(body *models.CreateTagOption) *RepoCreateTagParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the repo create tag params
func (o *RepoCreateTagParams) SetBody(body *models.CreateTagOption) {
	o.Body = body
}

// WithOwner adds the owner to the repo create tag params
func (o *RepoCreateTagParams) WithOwner(owner string) *RepoCreateTagParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo create tag params
func (o *RepoCreateTagParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo create tag params
func (o *RepoCreateTagParams) WithRepo(repo string) *RepoCreateTagParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo create tag params
func (o *RepoCreateTagParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoCreateTagParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param repo
	if err := r.SetPathParam("repo", o.Repo); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
