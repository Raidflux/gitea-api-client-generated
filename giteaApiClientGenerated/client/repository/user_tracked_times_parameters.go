// Code generated by go-swagger; DO NOT EDIT.

package repository

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
)

// NewUserTrackedTimesParams creates a new UserTrackedTimesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserTrackedTimesParams() *UserTrackedTimesParams {
	return &UserTrackedTimesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserTrackedTimesParamsWithTimeout creates a new UserTrackedTimesParams object
// with the ability to set a timeout on a request.
func NewUserTrackedTimesParamsWithTimeout(timeout time.Duration) *UserTrackedTimesParams {
	return &UserTrackedTimesParams{
		timeout: timeout,
	}
}

// NewUserTrackedTimesParamsWithContext creates a new UserTrackedTimesParams object
// with the ability to set a context for a request.
func NewUserTrackedTimesParamsWithContext(ctx context.Context) *UserTrackedTimesParams {
	return &UserTrackedTimesParams{
		Context: ctx,
	}
}

// NewUserTrackedTimesParamsWithHTTPClient creates a new UserTrackedTimesParams object
// with the ability to set a custom HTTPClient for a request.
func NewUserTrackedTimesParamsWithHTTPClient(client *http.Client) *UserTrackedTimesParams {
	return &UserTrackedTimesParams{
		HTTPClient: client,
	}
}

/*
UserTrackedTimesParams contains all the parameters to send to the API endpoint

	for the user tracked times operation.

	Typically these are written to a http.Request.
*/
type UserTrackedTimesParams struct {

	/* Owner.

	   owner of the repo
	*/
	Owner string

	/* Repo.

	   name of the repo
	*/
	Repo string

	/* User.

	   username of user
	*/
	User string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the user tracked times params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserTrackedTimesParams) WithDefaults() *UserTrackedTimesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user tracked times params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserTrackedTimesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the user tracked times params
func (o *UserTrackedTimesParams) WithTimeout(timeout time.Duration) *UserTrackedTimesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user tracked times params
func (o *UserTrackedTimesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user tracked times params
func (o *UserTrackedTimesParams) WithContext(ctx context.Context) *UserTrackedTimesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user tracked times params
func (o *UserTrackedTimesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user tracked times params
func (o *UserTrackedTimesParams) WithHTTPClient(client *http.Client) *UserTrackedTimesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user tracked times params
func (o *UserTrackedTimesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the user tracked times params
func (o *UserTrackedTimesParams) WithOwner(owner string) *UserTrackedTimesParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the user tracked times params
func (o *UserTrackedTimesParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the user tracked times params
func (o *UserTrackedTimesParams) WithRepo(repo string) *UserTrackedTimesParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the user tracked times params
func (o *UserTrackedTimesParams) SetRepo(repo string) {
	o.Repo = repo
}

// WithUser adds the user to the user tracked times params
func (o *UserTrackedTimesParams) WithUser(user string) *UserTrackedTimesParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the user tracked times params
func (o *UserTrackedTimesParams) SetUser(user string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *UserTrackedTimesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param repo
	if err := r.SetPathParam("repo", o.Repo); err != nil {
		return err
	}

	// path param user
	if err := r.SetPathParam("user", o.User); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
