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

// NewOrgListCurrentUserOrgsParams creates a new OrgListCurrentUserOrgsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOrgListCurrentUserOrgsParams() *OrgListCurrentUserOrgsParams {
	return &OrgListCurrentUserOrgsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOrgListCurrentUserOrgsParamsWithTimeout creates a new OrgListCurrentUserOrgsParams object
// with the ability to set a timeout on a request.
func NewOrgListCurrentUserOrgsParamsWithTimeout(timeout time.Duration) *OrgListCurrentUserOrgsParams {
	return &OrgListCurrentUserOrgsParams{
		timeout: timeout,
	}
}

// NewOrgListCurrentUserOrgsParamsWithContext creates a new OrgListCurrentUserOrgsParams object
// with the ability to set a context for a request.
func NewOrgListCurrentUserOrgsParamsWithContext(ctx context.Context) *OrgListCurrentUserOrgsParams {
	return &OrgListCurrentUserOrgsParams{
		Context: ctx,
	}
}

// NewOrgListCurrentUserOrgsParamsWithHTTPClient creates a new OrgListCurrentUserOrgsParams object
// with the ability to set a custom HTTPClient for a request.
func NewOrgListCurrentUserOrgsParamsWithHTTPClient(client *http.Client) *OrgListCurrentUserOrgsParams {
	return &OrgListCurrentUserOrgsParams{
		HTTPClient: client,
	}
}

/*
OrgListCurrentUserOrgsParams contains all the parameters to send to the API endpoint

	for the org list current user orgs operation.

	Typically these are written to a http.Request.
*/
type OrgListCurrentUserOrgsParams struct {

	/* Limit.

	   page size of results
	*/
	Limit *int64

	/* Page.

	   page number of results to return (1-based)
	*/
	Page *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the org list current user orgs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrgListCurrentUserOrgsParams) WithDefaults() *OrgListCurrentUserOrgsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the org list current user orgs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrgListCurrentUserOrgsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the org list current user orgs params
func (o *OrgListCurrentUserOrgsParams) WithTimeout(timeout time.Duration) *OrgListCurrentUserOrgsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the org list current user orgs params
func (o *OrgListCurrentUserOrgsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the org list current user orgs params
func (o *OrgListCurrentUserOrgsParams) WithContext(ctx context.Context) *OrgListCurrentUserOrgsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the org list current user orgs params
func (o *OrgListCurrentUserOrgsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the org list current user orgs params
func (o *OrgListCurrentUserOrgsParams) WithHTTPClient(client *http.Client) *OrgListCurrentUserOrgsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the org list current user orgs params
func (o *OrgListCurrentUserOrgsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the org list current user orgs params
func (o *OrgListCurrentUserOrgsParams) WithLimit(limit *int64) *OrgListCurrentUserOrgsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the org list current user orgs params
func (o *OrgListCurrentUserOrgsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPage adds the page to the org list current user orgs params
func (o *OrgListCurrentUserOrgsParams) WithPage(page *int64) *OrgListCurrentUserOrgsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the org list current user orgs params
func (o *OrgListCurrentUserOrgsParams) SetPage(page *int64) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *OrgListCurrentUserOrgsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
