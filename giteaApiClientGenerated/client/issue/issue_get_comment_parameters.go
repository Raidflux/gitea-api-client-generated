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

// NewIssueGetCommentParams creates a new IssueGetCommentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIssueGetCommentParams() *IssueGetCommentParams {
	return &IssueGetCommentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIssueGetCommentParamsWithTimeout creates a new IssueGetCommentParams object
// with the ability to set a timeout on a request.
func NewIssueGetCommentParamsWithTimeout(timeout time.Duration) *IssueGetCommentParams {
	return &IssueGetCommentParams{
		timeout: timeout,
	}
}

// NewIssueGetCommentParamsWithContext creates a new IssueGetCommentParams object
// with the ability to set a context for a request.
func NewIssueGetCommentParamsWithContext(ctx context.Context) *IssueGetCommentParams {
	return &IssueGetCommentParams{
		Context: ctx,
	}
}

// NewIssueGetCommentParamsWithHTTPClient creates a new IssueGetCommentParams object
// with the ability to set a custom HTTPClient for a request.
func NewIssueGetCommentParamsWithHTTPClient(client *http.Client) *IssueGetCommentParams {
	return &IssueGetCommentParams{
		HTTPClient: client,
	}
}

/*
IssueGetCommentParams contains all the parameters to send to the API endpoint

	for the issue get comment operation.

	Typically these are written to a http.Request.
*/
type IssueGetCommentParams struct {

	/* ID.

	   id of the comment

	   Format: int64
	*/
	ID int64

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

// WithDefaults hydrates default values in the issue get comment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueGetCommentParams) WithDefaults() *IssueGetCommentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the issue get comment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueGetCommentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the issue get comment params
func (o *IssueGetCommentParams) WithTimeout(timeout time.Duration) *IssueGetCommentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the issue get comment params
func (o *IssueGetCommentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the issue get comment params
func (o *IssueGetCommentParams) WithContext(ctx context.Context) *IssueGetCommentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the issue get comment params
func (o *IssueGetCommentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the issue get comment params
func (o *IssueGetCommentParams) WithHTTPClient(client *http.Client) *IssueGetCommentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the issue get comment params
func (o *IssueGetCommentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the issue get comment params
func (o *IssueGetCommentParams) WithID(id int64) *IssueGetCommentParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the issue get comment params
func (o *IssueGetCommentParams) SetID(id int64) {
	o.ID = id
}

// WithOwner adds the owner to the issue get comment params
func (o *IssueGetCommentParams) WithOwner(owner string) *IssueGetCommentParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the issue get comment params
func (o *IssueGetCommentParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the issue get comment params
func (o *IssueGetCommentParams) WithRepo(repo string) *IssueGetCommentParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the issue get comment params
func (o *IssueGetCommentParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *IssueGetCommentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
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
