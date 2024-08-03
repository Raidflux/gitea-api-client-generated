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

// NewRepoUpdateAvatarParams creates a new RepoUpdateAvatarParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepoUpdateAvatarParams() *RepoUpdateAvatarParams {
	return &RepoUpdateAvatarParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepoUpdateAvatarParamsWithTimeout creates a new RepoUpdateAvatarParams object
// with the ability to set a timeout on a request.
func NewRepoUpdateAvatarParamsWithTimeout(timeout time.Duration) *RepoUpdateAvatarParams {
	return &RepoUpdateAvatarParams{
		timeout: timeout,
	}
}

// NewRepoUpdateAvatarParamsWithContext creates a new RepoUpdateAvatarParams object
// with the ability to set a context for a request.
func NewRepoUpdateAvatarParamsWithContext(ctx context.Context) *RepoUpdateAvatarParams {
	return &RepoUpdateAvatarParams{
		Context: ctx,
	}
}

// NewRepoUpdateAvatarParamsWithHTTPClient creates a new RepoUpdateAvatarParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepoUpdateAvatarParamsWithHTTPClient(client *http.Client) *RepoUpdateAvatarParams {
	return &RepoUpdateAvatarParams{
		HTTPClient: client,
	}
}

/*
RepoUpdateAvatarParams contains all the parameters to send to the API endpoint

	for the repo update avatar operation.

	Typically these are written to a http.Request.
*/
type RepoUpdateAvatarParams struct {

	// Body.
	Body *models.UpdateRepoAvatarOption

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

// WithDefaults hydrates default values in the repo update avatar params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoUpdateAvatarParams) WithDefaults() *RepoUpdateAvatarParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repo update avatar params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoUpdateAvatarParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repo update avatar params
func (o *RepoUpdateAvatarParams) WithTimeout(timeout time.Duration) *RepoUpdateAvatarParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo update avatar params
func (o *RepoUpdateAvatarParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo update avatar params
func (o *RepoUpdateAvatarParams) WithContext(ctx context.Context) *RepoUpdateAvatarParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo update avatar params
func (o *RepoUpdateAvatarParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo update avatar params
func (o *RepoUpdateAvatarParams) WithHTTPClient(client *http.Client) *RepoUpdateAvatarParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo update avatar params
func (o *RepoUpdateAvatarParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the repo update avatar params
func (o *RepoUpdateAvatarParams) WithBody(body *models.UpdateRepoAvatarOption) *RepoUpdateAvatarParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the repo update avatar params
func (o *RepoUpdateAvatarParams) SetBody(body *models.UpdateRepoAvatarOption) {
	o.Body = body
}

// WithOwner adds the owner to the repo update avatar params
func (o *RepoUpdateAvatarParams) WithOwner(owner string) *RepoUpdateAvatarParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo update avatar params
func (o *RepoUpdateAvatarParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo update avatar params
func (o *RepoUpdateAvatarParams) WithRepo(repo string) *RepoUpdateAvatarParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo update avatar params
func (o *RepoUpdateAvatarParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoUpdateAvatarParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
