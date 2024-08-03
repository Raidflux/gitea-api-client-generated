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

// NewIssueResetTimeParams creates a new IssueResetTimeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIssueResetTimeParams() *IssueResetTimeParams {
	return &IssueResetTimeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIssueResetTimeParamsWithTimeout creates a new IssueResetTimeParams object
// with the ability to set a timeout on a request.
func NewIssueResetTimeParamsWithTimeout(timeout time.Duration) *IssueResetTimeParams {
	return &IssueResetTimeParams{
		timeout: timeout,
	}
}

// NewIssueResetTimeParamsWithContext creates a new IssueResetTimeParams object
// with the ability to set a context for a request.
func NewIssueResetTimeParamsWithContext(ctx context.Context) *IssueResetTimeParams {
	return &IssueResetTimeParams{
		Context: ctx,
	}
}

// NewIssueResetTimeParamsWithHTTPClient creates a new IssueResetTimeParams object
// with the ability to set a custom HTTPClient for a request.
func NewIssueResetTimeParamsWithHTTPClient(client *http.Client) *IssueResetTimeParams {
	return &IssueResetTimeParams{
		HTTPClient: client,
	}
}

/*
IssueResetTimeParams contains all the parameters to send to the API endpoint

	for the issue reset time operation.

	Typically these are written to a http.Request.
*/
type IssueResetTimeParams struct {

	/* Index.

	   index of the issue to add tracked time to

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

// WithDefaults hydrates default values in the issue reset time params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueResetTimeParams) WithDefaults() *IssueResetTimeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the issue reset time params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueResetTimeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the issue reset time params
func (o *IssueResetTimeParams) WithTimeout(timeout time.Duration) *IssueResetTimeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the issue reset time params
func (o *IssueResetTimeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the issue reset time params
func (o *IssueResetTimeParams) WithContext(ctx context.Context) *IssueResetTimeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the issue reset time params
func (o *IssueResetTimeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the issue reset time params
func (o *IssueResetTimeParams) WithHTTPClient(client *http.Client) *IssueResetTimeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the issue reset time params
func (o *IssueResetTimeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIndex adds the index to the issue reset time params
func (o *IssueResetTimeParams) WithIndex(index int64) *IssueResetTimeParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the issue reset time params
func (o *IssueResetTimeParams) SetIndex(index int64) {
	o.Index = index
}

// WithOwner adds the owner to the issue reset time params
func (o *IssueResetTimeParams) WithOwner(owner string) *IssueResetTimeParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the issue reset time params
func (o *IssueResetTimeParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the issue reset time params
func (o *IssueResetTimeParams) WithRepo(repo string) *IssueResetTimeParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the issue reset time params
func (o *IssueResetTimeParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *IssueResetTimeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
