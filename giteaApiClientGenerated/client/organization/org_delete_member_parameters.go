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
)

// NewOrgDeleteMemberParams creates a new OrgDeleteMemberParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOrgDeleteMemberParams() *OrgDeleteMemberParams {
	return &OrgDeleteMemberParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOrgDeleteMemberParamsWithTimeout creates a new OrgDeleteMemberParams object
// with the ability to set a timeout on a request.
func NewOrgDeleteMemberParamsWithTimeout(timeout time.Duration) *OrgDeleteMemberParams {
	return &OrgDeleteMemberParams{
		timeout: timeout,
	}
}

// NewOrgDeleteMemberParamsWithContext creates a new OrgDeleteMemberParams object
// with the ability to set a context for a request.
func NewOrgDeleteMemberParamsWithContext(ctx context.Context) *OrgDeleteMemberParams {
	return &OrgDeleteMemberParams{
		Context: ctx,
	}
}

// NewOrgDeleteMemberParamsWithHTTPClient creates a new OrgDeleteMemberParams object
// with the ability to set a custom HTTPClient for a request.
func NewOrgDeleteMemberParamsWithHTTPClient(client *http.Client) *OrgDeleteMemberParams {
	return &OrgDeleteMemberParams{
		HTTPClient: client,
	}
}

/*
OrgDeleteMemberParams contains all the parameters to send to the API endpoint

	for the org delete member operation.

	Typically these are written to a http.Request.
*/
type OrgDeleteMemberParams struct {

	/* Org.

	   name of the organization
	*/
	Org string

	/* Username.

	   username of the user
	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the org delete member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrgDeleteMemberParams) WithDefaults() *OrgDeleteMemberParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the org delete member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrgDeleteMemberParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the org delete member params
func (o *OrgDeleteMemberParams) WithTimeout(timeout time.Duration) *OrgDeleteMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the org delete member params
func (o *OrgDeleteMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the org delete member params
func (o *OrgDeleteMemberParams) WithContext(ctx context.Context) *OrgDeleteMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the org delete member params
func (o *OrgDeleteMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the org delete member params
func (o *OrgDeleteMemberParams) WithHTTPClient(client *http.Client) *OrgDeleteMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the org delete member params
func (o *OrgDeleteMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrg adds the org to the org delete member params
func (o *OrgDeleteMemberParams) WithOrg(org string) *OrgDeleteMemberParams {
	o.SetOrg(org)
	return o
}

// SetOrg adds the org to the org delete member params
func (o *OrgDeleteMemberParams) SetOrg(org string) {
	o.Org = org
}

// WithUsername adds the username to the org delete member params
func (o *OrgDeleteMemberParams) WithUsername(username string) *OrgDeleteMemberParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the org delete member params
func (o *OrgDeleteMemberParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *OrgDeleteMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param org
	if err := r.SetPathParam("org", o.Org); err != nil {
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
