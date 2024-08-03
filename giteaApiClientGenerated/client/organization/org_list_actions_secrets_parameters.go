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

// NewOrgListActionsSecretsParams creates a new OrgListActionsSecretsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOrgListActionsSecretsParams() *OrgListActionsSecretsParams {
	return &OrgListActionsSecretsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOrgListActionsSecretsParamsWithTimeout creates a new OrgListActionsSecretsParams object
// with the ability to set a timeout on a request.
func NewOrgListActionsSecretsParamsWithTimeout(timeout time.Duration) *OrgListActionsSecretsParams {
	return &OrgListActionsSecretsParams{
		timeout: timeout,
	}
}

// NewOrgListActionsSecretsParamsWithContext creates a new OrgListActionsSecretsParams object
// with the ability to set a context for a request.
func NewOrgListActionsSecretsParamsWithContext(ctx context.Context) *OrgListActionsSecretsParams {
	return &OrgListActionsSecretsParams{
		Context: ctx,
	}
}

// NewOrgListActionsSecretsParamsWithHTTPClient creates a new OrgListActionsSecretsParams object
// with the ability to set a custom HTTPClient for a request.
func NewOrgListActionsSecretsParamsWithHTTPClient(client *http.Client) *OrgListActionsSecretsParams {
	return &OrgListActionsSecretsParams{
		HTTPClient: client,
	}
}

/*
OrgListActionsSecretsParams contains all the parameters to send to the API endpoint

	for the org list actions secrets operation.

	Typically these are written to a http.Request.
*/
type OrgListActionsSecretsParams struct {

	/* Limit.

	   page size of results
	*/
	Limit *int64

	/* Org.

	   name of the organization
	*/
	Org string

	/* Page.

	   page number of results to return (1-based)
	*/
	Page *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the org list actions secrets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrgListActionsSecretsParams) WithDefaults() *OrgListActionsSecretsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the org list actions secrets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrgListActionsSecretsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the org list actions secrets params
func (o *OrgListActionsSecretsParams) WithTimeout(timeout time.Duration) *OrgListActionsSecretsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the org list actions secrets params
func (o *OrgListActionsSecretsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the org list actions secrets params
func (o *OrgListActionsSecretsParams) WithContext(ctx context.Context) *OrgListActionsSecretsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the org list actions secrets params
func (o *OrgListActionsSecretsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the org list actions secrets params
func (o *OrgListActionsSecretsParams) WithHTTPClient(client *http.Client) *OrgListActionsSecretsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the org list actions secrets params
func (o *OrgListActionsSecretsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the org list actions secrets params
func (o *OrgListActionsSecretsParams) WithLimit(limit *int64) *OrgListActionsSecretsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the org list actions secrets params
func (o *OrgListActionsSecretsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOrg adds the org to the org list actions secrets params
func (o *OrgListActionsSecretsParams) WithOrg(org string) *OrgListActionsSecretsParams {
	o.SetOrg(org)
	return o
}

// SetOrg adds the org to the org list actions secrets params
func (o *OrgListActionsSecretsParams) SetOrg(org string) {
	o.Org = org
}

// WithPage adds the page to the org list actions secrets params
func (o *OrgListActionsSecretsParams) WithPage(page *int64) *OrgListActionsSecretsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the org list actions secrets params
func (o *OrgListActionsSecretsParams) SetPage(page *int64) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *OrgListActionsSecretsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	// path param org
	if err := r.SetPathParam("org", o.Org); err != nil {
		return err
	}

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
