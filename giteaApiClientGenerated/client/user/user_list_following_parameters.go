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

// NewUserListFollowingParams creates a new UserListFollowingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserListFollowingParams() *UserListFollowingParams {
	return &UserListFollowingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserListFollowingParamsWithTimeout creates a new UserListFollowingParams object
// with the ability to set a timeout on a request.
func NewUserListFollowingParamsWithTimeout(timeout time.Duration) *UserListFollowingParams {
	return &UserListFollowingParams{
		timeout: timeout,
	}
}

// NewUserListFollowingParamsWithContext creates a new UserListFollowingParams object
// with the ability to set a context for a request.
func NewUserListFollowingParamsWithContext(ctx context.Context) *UserListFollowingParams {
	return &UserListFollowingParams{
		Context: ctx,
	}
}

// NewUserListFollowingParamsWithHTTPClient creates a new UserListFollowingParams object
// with the ability to set a custom HTTPClient for a request.
func NewUserListFollowingParamsWithHTTPClient(client *http.Client) *UserListFollowingParams {
	return &UserListFollowingParams{
		HTTPClient: client,
	}
}

/*
UserListFollowingParams contains all the parameters to send to the API endpoint

	for the user list following operation.

	Typically these are written to a http.Request.
*/
type UserListFollowingParams struct {

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

// WithDefaults hydrates default values in the user list following params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserListFollowingParams) WithDefaults() *UserListFollowingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user list following params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserListFollowingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the user list following params
func (o *UserListFollowingParams) WithTimeout(timeout time.Duration) *UserListFollowingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user list following params
func (o *UserListFollowingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user list following params
func (o *UserListFollowingParams) WithContext(ctx context.Context) *UserListFollowingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user list following params
func (o *UserListFollowingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user list following params
func (o *UserListFollowingParams) WithHTTPClient(client *http.Client) *UserListFollowingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user list following params
func (o *UserListFollowingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the user list following params
func (o *UserListFollowingParams) WithLimit(limit *int64) *UserListFollowingParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the user list following params
func (o *UserListFollowingParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPage adds the page to the user list following params
func (o *UserListFollowingParams) WithPage(page *int64) *UserListFollowingParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the user list following params
func (o *UserListFollowingParams) SetPage(page *int64) {
	o.Page = page
}

// WithUsername adds the username to the user list following params
func (o *UserListFollowingParams) WithUsername(username string) *UserListFollowingParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the user list following params
func (o *UserListFollowingParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *UserListFollowingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
