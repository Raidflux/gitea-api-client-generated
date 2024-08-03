// Code generated by go-swagger; DO NOT EDIT.

package admin

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

// NewAdminSearchEmailsParams creates a new AdminSearchEmailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAdminSearchEmailsParams() *AdminSearchEmailsParams {
	return &AdminSearchEmailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAdminSearchEmailsParamsWithTimeout creates a new AdminSearchEmailsParams object
// with the ability to set a timeout on a request.
func NewAdminSearchEmailsParamsWithTimeout(timeout time.Duration) *AdminSearchEmailsParams {
	return &AdminSearchEmailsParams{
		timeout: timeout,
	}
}

// NewAdminSearchEmailsParamsWithContext creates a new AdminSearchEmailsParams object
// with the ability to set a context for a request.
func NewAdminSearchEmailsParamsWithContext(ctx context.Context) *AdminSearchEmailsParams {
	return &AdminSearchEmailsParams{
		Context: ctx,
	}
}

// NewAdminSearchEmailsParamsWithHTTPClient creates a new AdminSearchEmailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewAdminSearchEmailsParamsWithHTTPClient(client *http.Client) *AdminSearchEmailsParams {
	return &AdminSearchEmailsParams{
		HTTPClient: client,
	}
}

/*
AdminSearchEmailsParams contains all the parameters to send to the API endpoint

	for the admin search emails operation.

	Typically these are written to a http.Request.
*/
type AdminSearchEmailsParams struct {

	/* Limit.

	   page size of results
	*/
	Limit *int64

	/* Page.

	   page number of results to return (1-based)
	*/
	Page *int64

	/* Q.

	   keyword
	*/
	Q *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the admin search emails params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminSearchEmailsParams) WithDefaults() *AdminSearchEmailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the admin search emails params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminSearchEmailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the admin search emails params
func (o *AdminSearchEmailsParams) WithTimeout(timeout time.Duration) *AdminSearchEmailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin search emails params
func (o *AdminSearchEmailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin search emails params
func (o *AdminSearchEmailsParams) WithContext(ctx context.Context) *AdminSearchEmailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin search emails params
func (o *AdminSearchEmailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin search emails params
func (o *AdminSearchEmailsParams) WithHTTPClient(client *http.Client) *AdminSearchEmailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin search emails params
func (o *AdminSearchEmailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the admin search emails params
func (o *AdminSearchEmailsParams) WithLimit(limit *int64) *AdminSearchEmailsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the admin search emails params
func (o *AdminSearchEmailsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPage adds the page to the admin search emails params
func (o *AdminSearchEmailsParams) WithPage(page *int64) *AdminSearchEmailsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the admin search emails params
func (o *AdminSearchEmailsParams) SetPage(page *int64) {
	o.Page = page
}

// WithQ adds the q to the admin search emails params
func (o *AdminSearchEmailsParams) WithQ(q *string) *AdminSearchEmailsParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the admin search emails params
func (o *AdminSearchEmailsParams) SetQ(q *string) {
	o.Q = q
}

// WriteToRequest writes these params to a swagger request
func (o *AdminSearchEmailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
