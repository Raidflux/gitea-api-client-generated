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

	"github.com/Raidflux/gitea-api-client-generated/giteaApiClientGenerated/models"
)

// NewRepoAddCollaboratorParams creates a new RepoAddCollaboratorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepoAddCollaboratorParams() *RepoAddCollaboratorParams {
	return &RepoAddCollaboratorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepoAddCollaboratorParamsWithTimeout creates a new RepoAddCollaboratorParams object
// with the ability to set a timeout on a request.
func NewRepoAddCollaboratorParamsWithTimeout(timeout time.Duration) *RepoAddCollaboratorParams {
	return &RepoAddCollaboratorParams{
		timeout: timeout,
	}
}

// NewRepoAddCollaboratorParamsWithContext creates a new RepoAddCollaboratorParams object
// with the ability to set a context for a request.
func NewRepoAddCollaboratorParamsWithContext(ctx context.Context) *RepoAddCollaboratorParams {
	return &RepoAddCollaboratorParams{
		Context: ctx,
	}
}

// NewRepoAddCollaboratorParamsWithHTTPClient creates a new RepoAddCollaboratorParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepoAddCollaboratorParamsWithHTTPClient(client *http.Client) *RepoAddCollaboratorParams {
	return &RepoAddCollaboratorParams{
		HTTPClient: client,
	}
}

/*
RepoAddCollaboratorParams contains all the parameters to send to the API endpoint

	for the repo add collaborator operation.

	Typically these are written to a http.Request.
*/
type RepoAddCollaboratorParams struct {

	// Body.
	Body *models.AddCollaboratorOption

	/* Collaborator.

	   username of the collaborator to add
	*/
	Collaborator string

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

// WithDefaults hydrates default values in the repo add collaborator params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoAddCollaboratorParams) WithDefaults() *RepoAddCollaboratorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repo add collaborator params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoAddCollaboratorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repo add collaborator params
func (o *RepoAddCollaboratorParams) WithTimeout(timeout time.Duration) *RepoAddCollaboratorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo add collaborator params
func (o *RepoAddCollaboratorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo add collaborator params
func (o *RepoAddCollaboratorParams) WithContext(ctx context.Context) *RepoAddCollaboratorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo add collaborator params
func (o *RepoAddCollaboratorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo add collaborator params
func (o *RepoAddCollaboratorParams) WithHTTPClient(client *http.Client) *RepoAddCollaboratorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo add collaborator params
func (o *RepoAddCollaboratorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the repo add collaborator params
func (o *RepoAddCollaboratorParams) WithBody(body *models.AddCollaboratorOption) *RepoAddCollaboratorParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the repo add collaborator params
func (o *RepoAddCollaboratorParams) SetBody(body *models.AddCollaboratorOption) {
	o.Body = body
}

// WithCollaborator adds the collaborator to the repo add collaborator params
func (o *RepoAddCollaboratorParams) WithCollaborator(collaborator string) *RepoAddCollaboratorParams {
	o.SetCollaborator(collaborator)
	return o
}

// SetCollaborator adds the collaborator to the repo add collaborator params
func (o *RepoAddCollaboratorParams) SetCollaborator(collaborator string) {
	o.Collaborator = collaborator
}

// WithOwner adds the owner to the repo add collaborator params
func (o *RepoAddCollaboratorParams) WithOwner(owner string) *RepoAddCollaboratorParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo add collaborator params
func (o *RepoAddCollaboratorParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo add collaborator params
func (o *RepoAddCollaboratorParams) WithRepo(repo string) *RepoAddCollaboratorParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo add collaborator params
func (o *RepoAddCollaboratorParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoAddCollaboratorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param collaborator
	if err := r.SetPathParam("collaborator", o.Collaborator); err != nil {
		return err
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
