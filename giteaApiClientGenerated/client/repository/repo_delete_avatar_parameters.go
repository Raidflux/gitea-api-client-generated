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

// NewRepoDeleteAvatarParams creates a new RepoDeleteAvatarParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepoDeleteAvatarParams() *RepoDeleteAvatarParams {
	return &RepoDeleteAvatarParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepoDeleteAvatarParamsWithTimeout creates a new RepoDeleteAvatarParams object
// with the ability to set a timeout on a request.
func NewRepoDeleteAvatarParamsWithTimeout(timeout time.Duration) *RepoDeleteAvatarParams {
	return &RepoDeleteAvatarParams{
		timeout: timeout,
	}
}

// NewRepoDeleteAvatarParamsWithContext creates a new RepoDeleteAvatarParams object
// with the ability to set a context for a request.
func NewRepoDeleteAvatarParamsWithContext(ctx context.Context) *RepoDeleteAvatarParams {
	return &RepoDeleteAvatarParams{
		Context: ctx,
	}
}

// NewRepoDeleteAvatarParamsWithHTTPClient creates a new RepoDeleteAvatarParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepoDeleteAvatarParamsWithHTTPClient(client *http.Client) *RepoDeleteAvatarParams {
	return &RepoDeleteAvatarParams{
		HTTPClient: client,
	}
}

/*
RepoDeleteAvatarParams contains all the parameters to send to the API endpoint

	for the repo delete avatar operation.

	Typically these are written to a http.Request.
*/
type RepoDeleteAvatarParams struct {

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

// WithDefaults hydrates default values in the repo delete avatar params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoDeleteAvatarParams) WithDefaults() *RepoDeleteAvatarParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repo delete avatar params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoDeleteAvatarParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repo delete avatar params
func (o *RepoDeleteAvatarParams) WithTimeout(timeout time.Duration) *RepoDeleteAvatarParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo delete avatar params
func (o *RepoDeleteAvatarParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo delete avatar params
func (o *RepoDeleteAvatarParams) WithContext(ctx context.Context) *RepoDeleteAvatarParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo delete avatar params
func (o *RepoDeleteAvatarParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo delete avatar params
func (o *RepoDeleteAvatarParams) WithHTTPClient(client *http.Client) *RepoDeleteAvatarParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo delete avatar params
func (o *RepoDeleteAvatarParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the repo delete avatar params
func (o *RepoDeleteAvatarParams) WithOwner(owner string) *RepoDeleteAvatarParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo delete avatar params
func (o *RepoDeleteAvatarParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo delete avatar params
func (o *RepoDeleteAvatarParams) WithRepo(repo string) *RepoDeleteAvatarParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo delete avatar params
func (o *RepoDeleteAvatarParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoDeleteAvatarParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
