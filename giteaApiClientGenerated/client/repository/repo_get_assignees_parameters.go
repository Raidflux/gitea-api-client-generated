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

// NewRepoGetAssigneesParams creates a new RepoGetAssigneesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepoGetAssigneesParams() *RepoGetAssigneesParams {
	return &RepoGetAssigneesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepoGetAssigneesParamsWithTimeout creates a new RepoGetAssigneesParams object
// with the ability to set a timeout on a request.
func NewRepoGetAssigneesParamsWithTimeout(timeout time.Duration) *RepoGetAssigneesParams {
	return &RepoGetAssigneesParams{
		timeout: timeout,
	}
}

// NewRepoGetAssigneesParamsWithContext creates a new RepoGetAssigneesParams object
// with the ability to set a context for a request.
func NewRepoGetAssigneesParamsWithContext(ctx context.Context) *RepoGetAssigneesParams {
	return &RepoGetAssigneesParams{
		Context: ctx,
	}
}

// NewRepoGetAssigneesParamsWithHTTPClient creates a new RepoGetAssigneesParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepoGetAssigneesParamsWithHTTPClient(client *http.Client) *RepoGetAssigneesParams {
	return &RepoGetAssigneesParams{
		HTTPClient: client,
	}
}

/*
RepoGetAssigneesParams contains all the parameters to send to the API endpoint

	for the repo get assignees operation.

	Typically these are written to a http.Request.
*/
type RepoGetAssigneesParams struct {

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

// WithDefaults hydrates default values in the repo get assignees params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoGetAssigneesParams) WithDefaults() *RepoGetAssigneesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repo get assignees params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoGetAssigneesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repo get assignees params
func (o *RepoGetAssigneesParams) WithTimeout(timeout time.Duration) *RepoGetAssigneesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo get assignees params
func (o *RepoGetAssigneesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo get assignees params
func (o *RepoGetAssigneesParams) WithContext(ctx context.Context) *RepoGetAssigneesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo get assignees params
func (o *RepoGetAssigneesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo get assignees params
func (o *RepoGetAssigneesParams) WithHTTPClient(client *http.Client) *RepoGetAssigneesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo get assignees params
func (o *RepoGetAssigneesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the repo get assignees params
func (o *RepoGetAssigneesParams) WithOwner(owner string) *RepoGetAssigneesParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo get assignees params
func (o *RepoGetAssigneesParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo get assignees params
func (o *RepoGetAssigneesParams) WithRepo(repo string) *RepoGetAssigneesParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo get assignees params
func (o *RepoGetAssigneesParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoGetAssigneesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
