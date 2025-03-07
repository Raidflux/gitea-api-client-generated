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

// NewUserCurrentDeleteSubscriptionParams creates a new UserCurrentDeleteSubscriptionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserCurrentDeleteSubscriptionParams() *UserCurrentDeleteSubscriptionParams {
	return &UserCurrentDeleteSubscriptionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserCurrentDeleteSubscriptionParamsWithTimeout creates a new UserCurrentDeleteSubscriptionParams object
// with the ability to set a timeout on a request.
func NewUserCurrentDeleteSubscriptionParamsWithTimeout(timeout time.Duration) *UserCurrentDeleteSubscriptionParams {
	return &UserCurrentDeleteSubscriptionParams{
		timeout: timeout,
	}
}

// NewUserCurrentDeleteSubscriptionParamsWithContext creates a new UserCurrentDeleteSubscriptionParams object
// with the ability to set a context for a request.
func NewUserCurrentDeleteSubscriptionParamsWithContext(ctx context.Context) *UserCurrentDeleteSubscriptionParams {
	return &UserCurrentDeleteSubscriptionParams{
		Context: ctx,
	}
}

// NewUserCurrentDeleteSubscriptionParamsWithHTTPClient creates a new UserCurrentDeleteSubscriptionParams object
// with the ability to set a custom HTTPClient for a request.
func NewUserCurrentDeleteSubscriptionParamsWithHTTPClient(client *http.Client) *UserCurrentDeleteSubscriptionParams {
	return &UserCurrentDeleteSubscriptionParams{
		HTTPClient: client,
	}
}

/*
UserCurrentDeleteSubscriptionParams contains all the parameters to send to the API endpoint

	for the user current delete subscription operation.

	Typically these are written to a http.Request.
*/
type UserCurrentDeleteSubscriptionParams struct {

	/* Owner.

	   owner of the repo
	*/
	Owner string

	/* Repo.

	   name of the repo
	*/
	Repo string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the user current delete subscription params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserCurrentDeleteSubscriptionParams) WithDefaults() *UserCurrentDeleteSubscriptionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user current delete subscription params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserCurrentDeleteSubscriptionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the user current delete subscription params
func (o *UserCurrentDeleteSubscriptionParams) WithTimeout(timeout time.Duration) *UserCurrentDeleteSubscriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user current delete subscription params
func (o *UserCurrentDeleteSubscriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user current delete subscription params
func (o *UserCurrentDeleteSubscriptionParams) WithContext(ctx context.Context) *UserCurrentDeleteSubscriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user current delete subscription params
func (o *UserCurrentDeleteSubscriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user current delete subscription params
func (o *UserCurrentDeleteSubscriptionParams) WithHTTPClient(client *http.Client) *UserCurrentDeleteSubscriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user current delete subscription params
func (o *UserCurrentDeleteSubscriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the user current delete subscription params
func (o *UserCurrentDeleteSubscriptionParams) WithOwner(owner string) *UserCurrentDeleteSubscriptionParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the user current delete subscription params
func (o *UserCurrentDeleteSubscriptionParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the user current delete subscription params
func (o *UserCurrentDeleteSubscriptionParams) WithRepo(repo string) *UserCurrentDeleteSubscriptionParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the user current delete subscription params
func (o *UserCurrentDeleteSubscriptionParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *UserCurrentDeleteSubscriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
