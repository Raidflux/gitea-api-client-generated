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

// NewOrgListTeamsParams creates a new OrgListTeamsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOrgListTeamsParams() *OrgListTeamsParams {
	return &OrgListTeamsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOrgListTeamsParamsWithTimeout creates a new OrgListTeamsParams object
// with the ability to set a timeout on a request.
func NewOrgListTeamsParamsWithTimeout(timeout time.Duration) *OrgListTeamsParams {
	return &OrgListTeamsParams{
		timeout: timeout,
	}
}

// NewOrgListTeamsParamsWithContext creates a new OrgListTeamsParams object
// with the ability to set a context for a request.
func NewOrgListTeamsParamsWithContext(ctx context.Context) *OrgListTeamsParams {
	return &OrgListTeamsParams{
		Context: ctx,
	}
}

// NewOrgListTeamsParamsWithHTTPClient creates a new OrgListTeamsParams object
// with the ability to set a custom HTTPClient for a request.
func NewOrgListTeamsParamsWithHTTPClient(client *http.Client) *OrgListTeamsParams {
	return &OrgListTeamsParams{
		HTTPClient: client,
	}
}

/*
OrgListTeamsParams contains all the parameters to send to the API endpoint

	for the org list teams operation.

	Typically these are written to a http.Request.
*/
type OrgListTeamsParams struct {

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

// WithDefaults hydrates default values in the org list teams params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrgListTeamsParams) WithDefaults() *OrgListTeamsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the org list teams params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrgListTeamsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the org list teams params
func (o *OrgListTeamsParams) WithTimeout(timeout time.Duration) *OrgListTeamsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the org list teams params
func (o *OrgListTeamsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the org list teams params
func (o *OrgListTeamsParams) WithContext(ctx context.Context) *OrgListTeamsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the org list teams params
func (o *OrgListTeamsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the org list teams params
func (o *OrgListTeamsParams) WithHTTPClient(client *http.Client) *OrgListTeamsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the org list teams params
func (o *OrgListTeamsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the org list teams params
func (o *OrgListTeamsParams) WithLimit(limit *int64) *OrgListTeamsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the org list teams params
func (o *OrgListTeamsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOrg adds the org to the org list teams params
func (o *OrgListTeamsParams) WithOrg(org string) *OrgListTeamsParams {
	o.SetOrg(org)
	return o
}

// SetOrg adds the org to the org list teams params
func (o *OrgListTeamsParams) SetOrg(org string) {
	o.Org = org
}

// WithPage adds the page to the org list teams params
func (o *OrgListTeamsParams) WithPage(page *int64) *OrgListTeamsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the org list teams params
func (o *OrgListTeamsParams) SetPage(page *int64) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *OrgListTeamsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
