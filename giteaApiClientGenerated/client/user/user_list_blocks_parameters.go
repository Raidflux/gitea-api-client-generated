// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewUserListBlocksParams creates a new UserListBlocksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserListBlocksParams() *UserListBlocksParams {
	return &UserListBlocksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserListBlocksParamsWithTimeout creates a new UserListBlocksParams object
// with the ability to set a timeout on a request.
func NewUserListBlocksParamsWithTimeout(timeout time.Duration) *UserListBlocksParams {
	return &UserListBlocksParams{
		timeout: timeout,
	}
}

// NewUserListBlocksParamsWithContext creates a new UserListBlocksParams object
// with the ability to set a context for a request.
func NewUserListBlocksParamsWithContext(ctx context.Context) *UserListBlocksParams {
	return &UserListBlocksParams{
		Context: ctx,
	}
}

// NewUserListBlocksParamsWithHTTPClient creates a new UserListBlocksParams object
// with the ability to set a custom HTTPClient for a request.
func NewUserListBlocksParamsWithHTTPClient(client *http.Client) *UserListBlocksParams {
	return &UserListBlocksParams{
		HTTPClient: client,
	}
}

/*
UserListBlocksParams contains all the parameters to send to the API endpoint

	for the user list blocks operation.

	Typically these are written to a http.Request.
*/
type UserListBlocksParams struct {

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

// WithDefaults hydrates default values in the user list blocks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserListBlocksParams) WithDefaults() *UserListBlocksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user list blocks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserListBlocksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the user list blocks params
func (o *UserListBlocksParams) WithTimeout(timeout time.Duration) *UserListBlocksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user list blocks params
func (o *UserListBlocksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user list blocks params
func (o *UserListBlocksParams) WithContext(ctx context.Context) *UserListBlocksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user list blocks params
func (o *UserListBlocksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user list blocks params
func (o *UserListBlocksParams) WithHTTPClient(client *http.Client) *UserListBlocksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user list blocks params
func (o *UserListBlocksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the user list blocks params
func (o *UserListBlocksParams) WithLimit(limit *int64) *UserListBlocksParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the user list blocks params
func (o *UserListBlocksParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPage adds the page to the user list blocks params
func (o *UserListBlocksParams) WithPage(page *int64) *UserListBlocksParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the user list blocks params
func (o *UserListBlocksParams) SetPage(page *int64) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *UserListBlocksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
