// Code generated by go-swagger; DO NOT EDIT.

package issue

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

// NewIssueCheckSubscriptionParams creates a new IssueCheckSubscriptionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIssueCheckSubscriptionParams() *IssueCheckSubscriptionParams {
	return &IssueCheckSubscriptionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIssueCheckSubscriptionParamsWithTimeout creates a new IssueCheckSubscriptionParams object
// with the ability to set a timeout on a request.
func NewIssueCheckSubscriptionParamsWithTimeout(timeout time.Duration) *IssueCheckSubscriptionParams {
	return &IssueCheckSubscriptionParams{
		timeout: timeout,
	}
}

// NewIssueCheckSubscriptionParamsWithContext creates a new IssueCheckSubscriptionParams object
// with the ability to set a context for a request.
func NewIssueCheckSubscriptionParamsWithContext(ctx context.Context) *IssueCheckSubscriptionParams {
	return &IssueCheckSubscriptionParams{
		Context: ctx,
	}
}

// NewIssueCheckSubscriptionParamsWithHTTPClient creates a new IssueCheckSubscriptionParams object
// with the ability to set a custom HTTPClient for a request.
func NewIssueCheckSubscriptionParamsWithHTTPClient(client *http.Client) *IssueCheckSubscriptionParams {
	return &IssueCheckSubscriptionParams{
		HTTPClient: client,
	}
}

/*
IssueCheckSubscriptionParams contains all the parameters to send to the API endpoint

	for the issue check subscription operation.

	Typically these are written to a http.Request.
*/
type IssueCheckSubscriptionParams struct {

	/* Index.

	   index of the issue

	   Format: int64
	*/
	Index int64

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

// WithDefaults hydrates default values in the issue check subscription params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueCheckSubscriptionParams) WithDefaults() *IssueCheckSubscriptionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the issue check subscription params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueCheckSubscriptionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the issue check subscription params
func (o *IssueCheckSubscriptionParams) WithTimeout(timeout time.Duration) *IssueCheckSubscriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the issue check subscription params
func (o *IssueCheckSubscriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the issue check subscription params
func (o *IssueCheckSubscriptionParams) WithContext(ctx context.Context) *IssueCheckSubscriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the issue check subscription params
func (o *IssueCheckSubscriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the issue check subscription params
func (o *IssueCheckSubscriptionParams) WithHTTPClient(client *http.Client) *IssueCheckSubscriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the issue check subscription params
func (o *IssueCheckSubscriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIndex adds the index to the issue check subscription params
func (o *IssueCheckSubscriptionParams) WithIndex(index int64) *IssueCheckSubscriptionParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the issue check subscription params
func (o *IssueCheckSubscriptionParams) SetIndex(index int64) {
	o.Index = index
}

// WithOwner adds the owner to the issue check subscription params
func (o *IssueCheckSubscriptionParams) WithOwner(owner string) *IssueCheckSubscriptionParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the issue check subscription params
func (o *IssueCheckSubscriptionParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the issue check subscription params
func (o *IssueCheckSubscriptionParams) WithRepo(repo string) *IssueCheckSubscriptionParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the issue check subscription params
func (o *IssueCheckSubscriptionParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *IssueCheckSubscriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param index
	if err := r.SetPathParam("index", swag.FormatInt64(o.Index)); err != nil {
		return err
	}

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
