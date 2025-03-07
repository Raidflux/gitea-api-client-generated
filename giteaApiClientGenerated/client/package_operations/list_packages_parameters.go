// Code generated by go-swagger; DO NOT EDIT.

package package_operations

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

// NewListPackagesParams creates a new ListPackagesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListPackagesParams() *ListPackagesParams {
	return &ListPackagesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListPackagesParamsWithTimeout creates a new ListPackagesParams object
// with the ability to set a timeout on a request.
func NewListPackagesParamsWithTimeout(timeout time.Duration) *ListPackagesParams {
	return &ListPackagesParams{
		timeout: timeout,
	}
}

// NewListPackagesParamsWithContext creates a new ListPackagesParams object
// with the ability to set a context for a request.
func NewListPackagesParamsWithContext(ctx context.Context) *ListPackagesParams {
	return &ListPackagesParams{
		Context: ctx,
	}
}

// NewListPackagesParamsWithHTTPClient creates a new ListPackagesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListPackagesParamsWithHTTPClient(client *http.Client) *ListPackagesParams {
	return &ListPackagesParams{
		HTTPClient: client,
	}
}

/*
ListPackagesParams contains all the parameters to send to the API endpoint

	for the list packages operation.

	Typically these are written to a http.Request.
*/
type ListPackagesParams struct {

	/* Limit.

	   page size of results
	*/
	Limit *int64

	/* Owner.

	   owner of the packages
	*/
	Owner string

	/* Page.

	   page number of results to return (1-based)
	*/
	Page *int64

	/* Q.

	   name filter
	*/
	Q *string

	/* Type.

	   package type filter
	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list packages params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListPackagesParams) WithDefaults() *ListPackagesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list packages params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListPackagesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list packages params
func (o *ListPackagesParams) WithTimeout(timeout time.Duration) *ListPackagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list packages params
func (o *ListPackagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list packages params
func (o *ListPackagesParams) WithContext(ctx context.Context) *ListPackagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list packages params
func (o *ListPackagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list packages params
func (o *ListPackagesParams) WithHTTPClient(client *http.Client) *ListPackagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list packages params
func (o *ListPackagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the list packages params
func (o *ListPackagesParams) WithLimit(limit *int64) *ListPackagesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list packages params
func (o *ListPackagesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOwner adds the owner to the list packages params
func (o *ListPackagesParams) WithOwner(owner string) *ListPackagesParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the list packages params
func (o *ListPackagesParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithPage adds the page to the list packages params
func (o *ListPackagesParams) WithPage(page *int64) *ListPackagesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the list packages params
func (o *ListPackagesParams) SetPage(page *int64) {
	o.Page = page
}

// WithQ adds the q to the list packages params
func (o *ListPackagesParams) WithQ(q *string) *ListPackagesParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the list packages params
func (o *ListPackagesParams) SetQ(q *string) {
	o.Q = q
}

// WithType adds the typeVar to the list packages params
func (o *ListPackagesParams) WithType(typeVar *string) *ListPackagesParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the list packages params
func (o *ListPackagesParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *ListPackagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
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

	if o.Type != nil {

		// query param type
		var qrType string

		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
