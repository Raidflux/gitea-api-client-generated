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

// NewOrgGetLabelParams creates a new OrgGetLabelParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOrgGetLabelParams() *OrgGetLabelParams {
	return &OrgGetLabelParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOrgGetLabelParamsWithTimeout creates a new OrgGetLabelParams object
// with the ability to set a timeout on a request.
func NewOrgGetLabelParamsWithTimeout(timeout time.Duration) *OrgGetLabelParams {
	return &OrgGetLabelParams{
		timeout: timeout,
	}
}

// NewOrgGetLabelParamsWithContext creates a new OrgGetLabelParams object
// with the ability to set a context for a request.
func NewOrgGetLabelParamsWithContext(ctx context.Context) *OrgGetLabelParams {
	return &OrgGetLabelParams{
		Context: ctx,
	}
}

// NewOrgGetLabelParamsWithHTTPClient creates a new OrgGetLabelParams object
// with the ability to set a custom HTTPClient for a request.
func NewOrgGetLabelParamsWithHTTPClient(client *http.Client) *OrgGetLabelParams {
	return &OrgGetLabelParams{
		HTTPClient: client,
	}
}

/*
OrgGetLabelParams contains all the parameters to send to the API endpoint

	for the org get label operation.

	Typically these are written to a http.Request.
*/
type OrgGetLabelParams struct {

	/* ID.

	   id of the label to get

	   Format: int64
	*/
	ID int64

	/* Org.

	   name of the organization
	*/
	Org string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the org get label params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrgGetLabelParams) WithDefaults() *OrgGetLabelParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the org get label params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrgGetLabelParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the org get label params
func (o *OrgGetLabelParams) WithTimeout(timeout time.Duration) *OrgGetLabelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the org get label params
func (o *OrgGetLabelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the org get label params
func (o *OrgGetLabelParams) WithContext(ctx context.Context) *OrgGetLabelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the org get label params
func (o *OrgGetLabelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the org get label params
func (o *OrgGetLabelParams) WithHTTPClient(client *http.Client) *OrgGetLabelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the org get label params
func (o *OrgGetLabelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the org get label params
func (o *OrgGetLabelParams) WithID(id int64) *OrgGetLabelParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the org get label params
func (o *OrgGetLabelParams) SetID(id int64) {
	o.ID = id
}

// WithOrg adds the org to the org get label params
func (o *OrgGetLabelParams) WithOrg(org string) *OrgGetLabelParams {
	o.SetOrg(org)
	return o
}

// SetOrg adds the org to the org get label params
func (o *OrgGetLabelParams) SetOrg(org string) {
	o.Org = org
}

// WriteToRequest writes these params to a swagger request
func (o *OrgGetLabelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
