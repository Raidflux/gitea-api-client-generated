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

// NewOrgListTeamMemberParams creates a new OrgListTeamMemberParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOrgListTeamMemberParams() *OrgListTeamMemberParams {
	return &OrgListTeamMemberParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOrgListTeamMemberParamsWithTimeout creates a new OrgListTeamMemberParams object
// with the ability to set a timeout on a request.
func NewOrgListTeamMemberParamsWithTimeout(timeout time.Duration) *OrgListTeamMemberParams {
	return &OrgListTeamMemberParams{
		timeout: timeout,
	}
}

// NewOrgListTeamMemberParamsWithContext creates a new OrgListTeamMemberParams object
// with the ability to set a context for a request.
func NewOrgListTeamMemberParamsWithContext(ctx context.Context) *OrgListTeamMemberParams {
	return &OrgListTeamMemberParams{
		Context: ctx,
	}
}

// NewOrgListTeamMemberParamsWithHTTPClient creates a new OrgListTeamMemberParams object
// with the ability to set a custom HTTPClient for a request.
func NewOrgListTeamMemberParamsWithHTTPClient(client *http.Client) *OrgListTeamMemberParams {
	return &OrgListTeamMemberParams{
		HTTPClient: client,
	}
}

/*
OrgListTeamMemberParams contains all the parameters to send to the API endpoint

	for the org list team member operation.

	Typically these are written to a http.Request.
*/
type OrgListTeamMemberParams struct {

	/* ID.

	   id of the team

	   Format: int64
	*/
	ID int64

	/* Username.

	   username of the member to list
	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the org list team member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrgListTeamMemberParams) WithDefaults() *OrgListTeamMemberParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the org list team member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrgListTeamMemberParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the org list team member params
func (o *OrgListTeamMemberParams) WithTimeout(timeout time.Duration) *OrgListTeamMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the org list team member params
func (o *OrgListTeamMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the org list team member params
func (o *OrgListTeamMemberParams) WithContext(ctx context.Context) *OrgListTeamMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the org list team member params
func (o *OrgListTeamMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the org list team member params
func (o *OrgListTeamMemberParams) WithHTTPClient(client *http.Client) *OrgListTeamMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the org list team member params
func (o *OrgListTeamMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the org list team member params
func (o *OrgListTeamMemberParams) WithID(id int64) *OrgListTeamMemberParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the org list team member params
func (o *OrgListTeamMemberParams) SetID(id int64) {
	o.ID = id
}

// WithUsername adds the username to the org list team member params
func (o *OrgListTeamMemberParams) WithUsername(username string) *OrgListTeamMemberParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the org list team member params
func (o *OrgListTeamMemberParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *OrgListTeamMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
