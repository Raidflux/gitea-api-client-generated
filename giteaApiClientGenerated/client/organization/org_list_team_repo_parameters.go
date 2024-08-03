// Code generated by go-swagger; DO NOT EDIT.

package organization

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
	"github.com/go-openapi/swag"
)

// NewOrgListTeamRepoParams creates a new OrgListTeamRepoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOrgListTeamRepoParams() *OrgListTeamRepoParams {
	return &OrgListTeamRepoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOrgListTeamRepoParamsWithTimeout creates a new OrgListTeamRepoParams object
// with the ability to set a timeout on a request.
func NewOrgListTeamRepoParamsWithTimeout(timeout time.Duration) *OrgListTeamRepoParams {
	return &OrgListTeamRepoParams{
		timeout: timeout,
	}
}

// NewOrgListTeamRepoParamsWithContext creates a new OrgListTeamRepoParams object
// with the ability to set a context for a request.
func NewOrgListTeamRepoParamsWithContext(ctx context.Context) *OrgListTeamRepoParams {
	return &OrgListTeamRepoParams{
		Context: ctx,
	}
}

// NewOrgListTeamRepoParamsWithHTTPClient creates a new OrgListTeamRepoParams object
// with the ability to set a custom HTTPClient for a request.
func NewOrgListTeamRepoParamsWithHTTPClient(client *http.Client) *OrgListTeamRepoParams {
	return &OrgListTeamRepoParams{
		HTTPClient: client,
	}
}

/*
OrgListTeamRepoParams contains all the parameters to send to the API endpoint

	for the org list team repo operation.

	Typically these are written to a http.Request.
*/
type OrgListTeamRepoParams struct {

	/* ID.

	   id of the team

	   Format: int64
	*/
	ID int64

	/* Org.

	   organization that owns the repo to list
	*/
	Org string

	/* Repo.

	   name of the repo to list
	*/
	Repo string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the org list team repo params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrgListTeamRepoParams) WithDefaults() *OrgListTeamRepoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the org list team repo params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrgListTeamRepoParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the org list team repo params
func (o *OrgListTeamRepoParams) WithTimeout(timeout time.Duration) *OrgListTeamRepoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the org list team repo params
func (o *OrgListTeamRepoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the org list team repo params
func (o *OrgListTeamRepoParams) WithContext(ctx context.Context) *OrgListTeamRepoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the org list team repo params
func (o *OrgListTeamRepoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the org list team repo params
func (o *OrgListTeamRepoParams) WithHTTPClient(client *http.Client) *OrgListTeamRepoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the org list team repo params
func (o *OrgListTeamRepoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the org list team repo params
func (o *OrgListTeamRepoParams) WithID(id int64) *OrgListTeamRepoParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the org list team repo params
func (o *OrgListTeamRepoParams) SetID(id int64) {
	o.ID = id
}

// WithOrg adds the org to the org list team repo params
func (o *OrgListTeamRepoParams) WithOrg(org string) *OrgListTeamRepoParams {
	o.SetOrg(org)
	return o
}

// SetOrg adds the org to the org list team repo params
func (o *OrgListTeamRepoParams) SetOrg(org string) {
	o.Org = org
}

// WithRepo adds the repo to the org list team repo params
func (o *OrgListTeamRepoParams) WithRepo(repo string) *OrgListTeamRepoParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the org list team repo params
func (o *OrgListTeamRepoParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *OrgListTeamRepoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// path param org
	if err := r.SetPathParam("org", o.Org); err != nil {
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
