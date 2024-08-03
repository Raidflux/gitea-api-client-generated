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

// NewUserListFollowersParams creates a new UserListFollowersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserListFollowersParams() *UserListFollowersParams {
	return &UserListFollowersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserListFollowersParamsWithTimeout creates a new UserListFollowersParams object
// with the ability to set a timeout on a request.
func NewUserListFollowersParamsWithTimeout(timeout time.Duration) *UserListFollowersParams {
	return &UserListFollowersParams{
		timeout: timeout,
	}
}

// NewUserListFollowersParamsWithContext creates a new UserListFollowersParams object
// with the ability to set a context for a request.
func NewUserListFollowersParamsWithContext(ctx context.Context) *UserListFollowersParams {
	return &UserListFollowersParams{
		Context: ctx,
	}
}

// NewUserListFollowersParamsWithHTTPClient creates a new UserListFollowersParams object
// with the ability to set a custom HTTPClient for a request.
func NewUserListFollowersParamsWithHTTPClient(client *http.Client) *UserListFollowersParams {
	return &UserListFollowersParams{
		HTTPClient: client,
	}
}

/*
UserListFollowersParams contains all the parameters to send to the API endpoint

	for the user list followers operation.

	Typically these are written to a http.Request.
*/
type UserListFollowersParams struct {

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

// WithDefaults hydrates default values in the user list followers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserListFollowersParams) WithDefaults() *UserListFollowersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user list followers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserListFollowersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the user list followers params
func (o *UserListFollowersParams) WithTimeout(timeout time.Duration) *UserListFollowersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user list followers params
func (o *UserListFollowersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user list followers params
func (o *UserListFollowersParams) WithContext(ctx context.Context) *UserListFollowersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user list followers params
func (o *UserListFollowersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user list followers params
func (o *UserListFollowersParams) WithHTTPClient(client *http.Client) *UserListFollowersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user list followers params
func (o *UserListFollowersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the user list followers params
func (o *UserListFollowersParams) WithLimit(limit *int64) *UserListFollowersParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the user list followers params
func (o *UserListFollowersParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPage adds the page to the user list followers params
func (o *UserListFollowersParams) WithPage(page *int64) *UserListFollowersParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the user list followers params
func (o *UserListFollowersParams) SetPage(page *int64) {
	o.Page = page
}

// WithUsername adds the username to the user list followers params
func (o *UserListFollowersParams) WithUsername(username string) *UserListFollowersParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the user list followers params
func (o *UserListFollowersParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *UserListFollowersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
