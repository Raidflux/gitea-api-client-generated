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

// NewUserCurrentListFollowingParams creates a new UserCurrentListFollowingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserCurrentListFollowingParams() *UserCurrentListFollowingParams {
	return &UserCurrentListFollowingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserCurrentListFollowingParamsWithTimeout creates a new UserCurrentListFollowingParams object
// with the ability to set a timeout on a request.
func NewUserCurrentListFollowingParamsWithTimeout(timeout time.Duration) *UserCurrentListFollowingParams {
	return &UserCurrentListFollowingParams{
		timeout: timeout,
	}
}

// NewUserCurrentListFollowingParamsWithContext creates a new UserCurrentListFollowingParams object
// with the ability to set a context for a request.
func NewUserCurrentListFollowingParamsWithContext(ctx context.Context) *UserCurrentListFollowingParams {
	return &UserCurrentListFollowingParams{
		Context: ctx,
	}
}

// NewUserCurrentListFollowingParamsWithHTTPClient creates a new UserCurrentListFollowingParams object
// with the ability to set a custom HTTPClient for a request.
func NewUserCurrentListFollowingParamsWithHTTPClient(client *http.Client) *UserCurrentListFollowingParams {
	return &UserCurrentListFollowingParams{
		HTTPClient: client,
	}
}

/*
UserCurrentListFollowingParams contains all the parameters to send to the API endpoint

	for the user current list following operation.

	Typically these are written to a http.Request.
*/
type UserCurrentListFollowingParams struct {

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

// WithDefaults hydrates default values in the user current list following params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserCurrentListFollowingParams) WithDefaults() *UserCurrentListFollowingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user current list following params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserCurrentListFollowingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the user current list following params
func (o *UserCurrentListFollowingParams) WithTimeout(timeout time.Duration) *UserCurrentListFollowingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user current list following params
func (o *UserCurrentListFollowingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user current list following params
func (o *UserCurrentListFollowingParams) WithContext(ctx context.Context) *UserCurrentListFollowingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user current list following params
func (o *UserCurrentListFollowingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user current list following params
func (o *UserCurrentListFollowingParams) WithHTTPClient(client *http.Client) *UserCurrentListFollowingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user current list following params
func (o *UserCurrentListFollowingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the user current list following params
func (o *UserCurrentListFollowingParams) WithLimit(limit *int64) *UserCurrentListFollowingParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the user current list following params
func (o *UserCurrentListFollowingParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPage adds the page to the user current list following params
func (o *UserCurrentListFollowingParams) WithPage(page *int64) *UserCurrentListFollowingParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the user current list following params
func (o *UserCurrentListFollowingParams) SetPage(page *int64) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *UserCurrentListFollowingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
