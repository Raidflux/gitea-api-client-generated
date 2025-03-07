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

// NewAdminCronListParams creates a new AdminCronListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAdminCronListParams() *AdminCronListParams {
	return &AdminCronListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAdminCronListParamsWithTimeout creates a new AdminCronListParams object
// with the ability to set a timeout on a request.
func NewAdminCronListParamsWithTimeout(timeout time.Duration) *AdminCronListParams {
	return &AdminCronListParams{
		timeout: timeout,
	}
}

// NewAdminCronListParamsWithContext creates a new AdminCronListParams object
// with the ability to set a context for a request.
func NewAdminCronListParamsWithContext(ctx context.Context) *AdminCronListParams {
	return &AdminCronListParams{
		Context: ctx,
	}
}

// NewAdminCronListParamsWithHTTPClient creates a new AdminCronListParams object
// with the ability to set a custom HTTPClient for a request.
func NewAdminCronListParamsWithHTTPClient(client *http.Client) *AdminCronListParams {
	return &AdminCronListParams{
		HTTPClient: client,
	}
}

/*
AdminCronListParams contains all the parameters to send to the API endpoint

	for the admin cron list operation.

	Typically these are written to a http.Request.
*/
type AdminCronListParams struct {

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

// WithDefaults hydrates default values in the admin cron list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminCronListParams) WithDefaults() *AdminCronListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the admin cron list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminCronListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the admin cron list params
func (o *AdminCronListParams) WithTimeout(timeout time.Duration) *AdminCronListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin cron list params
func (o *AdminCronListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin cron list params
func (o *AdminCronListParams) WithContext(ctx context.Context) *AdminCronListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin cron list params
func (o *AdminCronListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin cron list params
func (o *AdminCronListParams) WithHTTPClient(client *http.Client) *AdminCronListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin cron list params
func (o *AdminCronListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the admin cron list params
func (o *AdminCronListParams) WithLimit(limit *int64) *AdminCronListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the admin cron list params
func (o *AdminCronListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPage adds the page to the admin cron list params
func (o *AdminCronListParams) WithPage(page *int64) *AdminCronListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the admin cron list params
func (o *AdminCronListParams) SetPage(page *int64) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *AdminCronListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
