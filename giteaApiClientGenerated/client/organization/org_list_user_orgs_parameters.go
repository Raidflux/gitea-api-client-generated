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

// NewOrgListUserOrgsParams creates a new OrgListUserOrgsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOrgListUserOrgsParams() *OrgListUserOrgsParams {
	return &OrgListUserOrgsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOrgListUserOrgsParamsWithTimeout creates a new OrgListUserOrgsParams object
// with the ability to set a timeout on a request.
func NewOrgListUserOrgsParamsWithTimeout(timeout time.Duration) *OrgListUserOrgsParams {
	return &OrgListUserOrgsParams{
		timeout: timeout,
	}
}

// NewOrgListUserOrgsParamsWithContext creates a new OrgListUserOrgsParams object
// with the ability to set a context for a request.
func NewOrgListUserOrgsParamsWithContext(ctx context.Context) *OrgListUserOrgsParams {
	return &OrgListUserOrgsParams{
		Context: ctx,
	}
}

// NewOrgListUserOrgsParamsWithHTTPClient creates a new OrgListUserOrgsParams object
// with the ability to set a custom HTTPClient for a request.
func NewOrgListUserOrgsParamsWithHTTPClient(client *http.Client) *OrgListUserOrgsParams {
	return &OrgListUserOrgsParams{
		HTTPClient: client,
	}
}

/*
OrgListUserOrgsParams contains all the parameters to send to the API endpoint

	for the org list user orgs operation.

	Typically these are written to a http.Request.
*/
type OrgListUserOrgsParams struct {

	/* Limit.

	   page size of results
	*/
	Limit *int64

	/* Page.

	   page number of results to return (1-based)
	*/
	Page *int64

	/* Username.

	   username of user
	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the org list user orgs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrgListUserOrgsParams) WithDefaults() *OrgListUserOrgsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the org list user orgs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrgListUserOrgsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the org list user orgs params
func (o *OrgListUserOrgsParams) WithTimeout(timeout time.Duration) *OrgListUserOrgsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the org list user orgs params
func (o *OrgListUserOrgsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the org list user orgs params
func (o *OrgListUserOrgsParams) WithContext(ctx context.Context) *OrgListUserOrgsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the org list user orgs params
func (o *OrgListUserOrgsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the org list user orgs params
func (o *OrgListUserOrgsParams) WithHTTPClient(client *http.Client) *OrgListUserOrgsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the org list user orgs params
func (o *OrgListUserOrgsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the org list user orgs params
func (o *OrgListUserOrgsParams) WithLimit(limit *int64) *OrgListUserOrgsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the org list user orgs params
func (o *OrgListUserOrgsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPage adds the page to the org list user orgs params
func (o *OrgListUserOrgsParams) WithPage(page *int64) *OrgListUserOrgsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the org list user orgs params
func (o *OrgListUserOrgsParams) SetPage(page *int64) {
	o.Page = page
}

// WithUsername adds the username to the org list user orgs params
func (o *OrgListUserOrgsParams) WithUsername(username string) *OrgListUserOrgsParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the org list user orgs params
func (o *OrgListUserOrgsParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *OrgListUserOrgsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
