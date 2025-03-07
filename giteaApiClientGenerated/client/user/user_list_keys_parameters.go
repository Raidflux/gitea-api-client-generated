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

// NewUserListKeysParams creates a new UserListKeysParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserListKeysParams() *UserListKeysParams {
	return &UserListKeysParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserListKeysParamsWithTimeout creates a new UserListKeysParams object
// with the ability to set a timeout on a request.
func NewUserListKeysParamsWithTimeout(timeout time.Duration) *UserListKeysParams {
	return &UserListKeysParams{
		timeout: timeout,
	}
}

// NewUserListKeysParamsWithContext creates a new UserListKeysParams object
// with the ability to set a context for a request.
func NewUserListKeysParamsWithContext(ctx context.Context) *UserListKeysParams {
	return &UserListKeysParams{
		Context: ctx,
	}
}

// NewUserListKeysParamsWithHTTPClient creates a new UserListKeysParams object
// with the ability to set a custom HTTPClient for a request.
func NewUserListKeysParamsWithHTTPClient(client *http.Client) *UserListKeysParams {
	return &UserListKeysParams{
		HTTPClient: client,
	}
}

/*
UserListKeysParams contains all the parameters to send to the API endpoint

	for the user list keys operation.

	Typically these are written to a http.Request.
*/
type UserListKeysParams struct {

	/* Fingerprint.

	   fingerprint of the key
	*/
	Fingerprint *string

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

// WithDefaults hydrates default values in the user list keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserListKeysParams) WithDefaults() *UserListKeysParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user list keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserListKeysParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the user list keys params
func (o *UserListKeysParams) WithTimeout(timeout time.Duration) *UserListKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user list keys params
func (o *UserListKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user list keys params
func (o *UserListKeysParams) WithContext(ctx context.Context) *UserListKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user list keys params
func (o *UserListKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user list keys params
func (o *UserListKeysParams) WithHTTPClient(client *http.Client) *UserListKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user list keys params
func (o *UserListKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFingerprint adds the fingerprint to the user list keys params
func (o *UserListKeysParams) WithFingerprint(fingerprint *string) *UserListKeysParams {
	o.SetFingerprint(fingerprint)
	return o
}

// SetFingerprint adds the fingerprint to the user list keys params
func (o *UserListKeysParams) SetFingerprint(fingerprint *string) {
	o.Fingerprint = fingerprint
}

// WithLimit adds the limit to the user list keys params
func (o *UserListKeysParams) WithLimit(limit *int64) *UserListKeysParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the user list keys params
func (o *UserListKeysParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPage adds the page to the user list keys params
func (o *UserListKeysParams) WithPage(page *int64) *UserListKeysParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the user list keys params
func (o *UserListKeysParams) SetPage(page *int64) {
	o.Page = page
}

// WithUsername adds the username to the user list keys params
func (o *UserListKeysParams) WithUsername(username string) *UserListKeysParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the user list keys params
func (o *UserListKeysParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *UserListKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fingerprint != nil {

		// query param fingerprint
		var qrFingerprint string

		if o.Fingerprint != nil {
			qrFingerprint = *o.Fingerprint
		}
		qFingerprint := qrFingerprint
		if qFingerprint != "" {

			if err := r.SetQueryParam("fingerprint", qFingerprint); err != nil {
				return err
			}
		}
	}

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
