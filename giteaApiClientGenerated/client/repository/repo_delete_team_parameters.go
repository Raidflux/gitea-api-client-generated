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

// NewRepoDeleteTeamParams creates a new RepoDeleteTeamParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepoDeleteTeamParams() *RepoDeleteTeamParams {
	return &RepoDeleteTeamParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepoDeleteTeamParamsWithTimeout creates a new RepoDeleteTeamParams object
// with the ability to set a timeout on a request.
func NewRepoDeleteTeamParamsWithTimeout(timeout time.Duration) *RepoDeleteTeamParams {
	return &RepoDeleteTeamParams{
		timeout: timeout,
	}
}

// NewRepoDeleteTeamParamsWithContext creates a new RepoDeleteTeamParams object
// with the ability to set a context for a request.
func NewRepoDeleteTeamParamsWithContext(ctx context.Context) *RepoDeleteTeamParams {
	return &RepoDeleteTeamParams{
		Context: ctx,
	}
}

// NewRepoDeleteTeamParamsWithHTTPClient creates a new RepoDeleteTeamParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepoDeleteTeamParamsWithHTTPClient(client *http.Client) *RepoDeleteTeamParams {
	return &RepoDeleteTeamParams{
		HTTPClient: client,
	}
}

/*
RepoDeleteTeamParams contains all the parameters to send to the API endpoint

	for the repo delete team operation.

	Typically these are written to a http.Request.
*/
type RepoDeleteTeamParams struct {

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

// WithDefaults hydrates default values in the repo delete team params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoDeleteTeamParams) WithDefaults() *RepoDeleteTeamParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repo delete team params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoDeleteTeamParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repo delete team params
func (o *RepoDeleteTeamParams) WithTimeout(timeout time.Duration) *RepoDeleteTeamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo delete team params
func (o *RepoDeleteTeamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo delete team params
func (o *RepoDeleteTeamParams) WithContext(ctx context.Context) *RepoDeleteTeamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo delete team params
func (o *RepoDeleteTeamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo delete team params
func (o *RepoDeleteTeamParams) WithHTTPClient(client *http.Client) *RepoDeleteTeamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo delete team params
func (o *RepoDeleteTeamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the repo delete team params
func (o *RepoDeleteTeamParams) WithOwner(owner string) *RepoDeleteTeamParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo delete team params
func (o *RepoDeleteTeamParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo delete team params
func (o *RepoDeleteTeamParams) WithRepo(repo string) *RepoDeleteTeamParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo delete team params
func (o *RepoDeleteTeamParams) SetRepo(repo string) {
	o.Repo = repo
}

// WithTeam adds the team to the repo delete team params
func (o *RepoDeleteTeamParams) WithTeam(team string) *RepoDeleteTeamParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the repo delete team params
func (o *RepoDeleteTeamParams) SetTeam(team string) {
	o.Team = team
}

// WriteToRequest writes these params to a swagger request
func (o *RepoDeleteTeamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
