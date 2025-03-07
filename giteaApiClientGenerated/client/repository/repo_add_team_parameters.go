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

// NewRepoAddTeamParams creates a new RepoAddTeamParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepoAddTeamParams() *RepoAddTeamParams {
	return &RepoAddTeamParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepoAddTeamParamsWithTimeout creates a new RepoAddTeamParams object
// with the ability to set a timeout on a request.
func NewRepoAddTeamParamsWithTimeout(timeout time.Duration) *RepoAddTeamParams {
	return &RepoAddTeamParams{
		timeout: timeout,
	}
}

// NewRepoAddTeamParamsWithContext creates a new RepoAddTeamParams object
// with the ability to set a context for a request.
func NewRepoAddTeamParamsWithContext(ctx context.Context) *RepoAddTeamParams {
	return &RepoAddTeamParams{
		Context: ctx,
	}
}

// NewRepoAddTeamParamsWithHTTPClient creates a new RepoAddTeamParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepoAddTeamParamsWithHTTPClient(client *http.Client) *RepoAddTeamParams {
	return &RepoAddTeamParams{
		HTTPClient: client,
	}
}

/*
RepoAddTeamParams contains all the parameters to send to the API endpoint

	for the repo add team operation.

	Typically these are written to a http.Request.
*/
type RepoAddTeamParams struct {

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

// WithDefaults hydrates default values in the repo add team params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoAddTeamParams) WithDefaults() *RepoAddTeamParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repo add team params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoAddTeamParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repo add team params
func (o *RepoAddTeamParams) WithTimeout(timeout time.Duration) *RepoAddTeamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo add team params
func (o *RepoAddTeamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo add team params
func (o *RepoAddTeamParams) WithContext(ctx context.Context) *RepoAddTeamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo add team params
func (o *RepoAddTeamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo add team params
func (o *RepoAddTeamParams) WithHTTPClient(client *http.Client) *RepoAddTeamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo add team params
func (o *RepoAddTeamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the repo add team params
func (o *RepoAddTeamParams) WithOwner(owner string) *RepoAddTeamParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo add team params
func (o *RepoAddTeamParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo add team params
func (o *RepoAddTeamParams) WithRepo(repo string) *RepoAddTeamParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo add team params
func (o *RepoAddTeamParams) SetRepo(repo string) {
	o.Repo = repo
}

// WithTeam adds the team to the repo add team params
func (o *RepoAddTeamParams) WithTeam(team string) *RepoAddTeamParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the repo add team params
func (o *RepoAddTeamParams) SetTeam(team string) {
	o.Team = team
}

// WriteToRequest writes these params to a swagger request
func (o *RepoAddTeamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
