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

	"github.com/Raidflux/gitea-api-client-generated/giteaApiClientGenerated/models"
)

// NewOrgEditHookParams creates a new OrgEditHookParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOrgEditHookParams() *OrgEditHookParams {
	return &OrgEditHookParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOrgEditHookParamsWithTimeout creates a new OrgEditHookParams object
// with the ability to set a timeout on a request.
func NewOrgEditHookParamsWithTimeout(timeout time.Duration) *OrgEditHookParams {
	return &OrgEditHookParams{
		timeout: timeout,
	}
}

// NewOrgEditHookParamsWithContext creates a new OrgEditHookParams object
// with the ability to set a context for a request.
func NewOrgEditHookParamsWithContext(ctx context.Context) *OrgEditHookParams {
	return &OrgEditHookParams{
		Context: ctx,
	}
}

// NewOrgEditHookParamsWithHTTPClient creates a new OrgEditHookParams object
// with the ability to set a custom HTTPClient for a request.
func NewOrgEditHookParamsWithHTTPClient(client *http.Client) *OrgEditHookParams {
	return &OrgEditHookParams{
		HTTPClient: client,
	}
}

/*
OrgEditHookParams contains all the parameters to send to the API endpoint

	for the org edit hook operation.

	Typically these are written to a http.Request.
*/
type OrgEditHookParams struct {

	// Body.
	Body *models.EditHookOption

	/* ID.

	   id of the hook to update

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

// WithDefaults hydrates default values in the org edit hook params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrgEditHookParams) WithDefaults() *OrgEditHookParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the org edit hook params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrgEditHookParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the org edit hook params
func (o *OrgEditHookParams) WithTimeout(timeout time.Duration) *OrgEditHookParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the org edit hook params
func (o *OrgEditHookParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the org edit hook params
func (o *OrgEditHookParams) WithContext(ctx context.Context) *OrgEditHookParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the org edit hook params
func (o *OrgEditHookParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the org edit hook params
func (o *OrgEditHookParams) WithHTTPClient(client *http.Client) *OrgEditHookParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the org edit hook params
func (o *OrgEditHookParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the org edit hook params
func (o *OrgEditHookParams) WithBody(body *models.EditHookOption) *OrgEditHookParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the org edit hook params
func (o *OrgEditHookParams) SetBody(body *models.EditHookOption) {
	o.Body = body
}

// WithID adds the id to the org edit hook params
func (o *OrgEditHookParams) WithID(id int64) *OrgEditHookParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the org edit hook params
func (o *OrgEditHookParams) SetID(id int64) {
	o.ID = id
}

// WithOrg adds the org to the org edit hook params
func (o *OrgEditHookParams) WithOrg(org string) *OrgEditHookParams {
	o.SetOrg(org)
	return o
}

// SetOrg adds the org to the org edit hook params
func (o *OrgEditHookParams) SetOrg(org string) {
	o.Org = org
}

// WriteToRequest writes these params to a swagger request
func (o *OrgEditHookParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
