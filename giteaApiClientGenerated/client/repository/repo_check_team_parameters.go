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

// NewRepoCheckTeamParams creates a new RepoCheckTeamParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepoCheckTeamParams() *RepoCheckTeamParams {
	return &RepoCheckTeamParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepoCheckTeamParamsWithTimeout creates a new RepoCheckTeamParams object
// with the ability to set a timeout on a request.
func NewRepoCheckTeamParamsWithTimeout(timeout time.Duration) *RepoCheckTeamParams {
	return &RepoCheckTeamParams{
		timeout: timeout,
	}
}

// NewRepoCheckTeamParamsWithContext creates a new RepoCheckTeamParams object
// with the ability to set a context for a request.
func NewRepoCheckTeamParamsWithContext(ctx context.Context) *RepoCheckTeamParams {
	return &RepoCheckTeamParams{
		Context: ctx,
	}
}

// NewRepoCheckTeamParamsWithHTTPClient creates a new RepoCheckTeamParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepoCheckTeamParamsWithHTTPClient(client *http.Client) *RepoCheckTeamParams {
	return &RepoCheckTeamParams{
		HTTPClient: client,
	}
}

/*
RepoCheckTeamParams contains all the parameters to send to the API endpoint

	for the repo check team operation.

	Typically these are written to a http.Request.
*/
type RepoCheckTeamParams struct {

	/* Owner.

	   owner of the repo
	*/
	Owner string

	/* Repo.

	   name of the repo
	*/
	Repo string

	/* Team.

	   team name
	*/
	Team string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the repo check team params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoCheckTeamParams) WithDefaults() *RepoCheckTeamParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repo check team params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoCheckTeamParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repo check team params
func (o *RepoCheckTeamParams) WithTimeout(timeout time.Duration) *RepoCheckTeamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo check team params
func (o *RepoCheckTeamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo check team params
func (o *RepoCheckTeamParams) WithContext(ctx context.Context) *RepoCheckTeamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo check team params
func (o *RepoCheckTeamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo check team params
func (o *RepoCheckTeamParams) WithHTTPClient(client *http.Client) *RepoCheckTeamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo check team params
func (o *RepoCheckTeamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the repo check team params
func (o *RepoCheckTeamParams) WithOwner(owner string) *RepoCheckTeamParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo check team params
func (o *RepoCheckTeamParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo check team params
func (o *RepoCheckTeamParams) WithRepo(repo string) *RepoCheckTeamParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo check team params
func (o *RepoCheckTeamParams) SetRepo(repo string) {
	o.Repo = repo
}

// WithTeam adds the team to the repo check team params
func (o *RepoCheckTeamParams) WithTeam(team string) *RepoCheckTeamParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the repo check team params
func (o *RepoCheckTeamParams) SetTeam(team string) {
	o.Team = team
}

// WriteToRequest writes these params to a swagger request
func (o *RepoCheckTeamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param team
	if err := r.SetPathParam("team", o.Team); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
