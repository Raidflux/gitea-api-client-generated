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

	"giteaApiClientGenerated/models"
)

// NewIssueEditCommentParams creates a new IssueEditCommentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIssueEditCommentParams() *IssueEditCommentParams {
	return &IssueEditCommentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIssueEditCommentParamsWithTimeout creates a new IssueEditCommentParams object
// with the ability to set a timeout on a request.
func NewIssueEditCommentParamsWithTimeout(timeout time.Duration) *IssueEditCommentParams {
	return &IssueEditCommentParams{
		timeout: timeout,
	}
}

// NewIssueEditCommentParamsWithContext creates a new IssueEditCommentParams object
// with the ability to set a context for a request.
func NewIssueEditCommentParamsWithContext(ctx context.Context) *IssueEditCommentParams {
	return &IssueEditCommentParams{
		Context: ctx,
	}
}

// NewIssueEditCommentParamsWithHTTPClient creates a new IssueEditCommentParams object
// with the ability to set a custom HTTPClient for a request.
func NewIssueEditCommentParamsWithHTTPClient(client *http.Client) *IssueEditCommentParams {
	return &IssueEditCommentParams{
		HTTPClient: client,
	}
}

/*
IssueEditCommentParams contains all the parameters to send to the API endpoint

	for the issue edit comment operation.

	Typically these are written to a http.Request.
*/
type IssueEditCommentParams struct {

	// Body.
	Body *models.EditIssueCommentOption

	/* ID.

	   id of the comment to edit

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

// WithDefaults hydrates default values in the issue edit comment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueEditCommentParams) WithDefaults() *IssueEditCommentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the issue edit comment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueEditCommentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the issue edit comment params
func (o *IssueEditCommentParams) WithTimeout(timeout time.Duration) *IssueEditCommentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the issue edit comment params
func (o *IssueEditCommentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the issue edit comment params
func (o *IssueEditCommentParams) WithContext(ctx context.Context) *IssueEditCommentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the issue edit comment params
func (o *IssueEditCommentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the issue edit comment params
func (o *IssueEditCommentParams) WithHTTPClient(client *http.Client) *IssueEditCommentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the issue edit comment params
func (o *IssueEditCommentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the issue edit comment params
func (o *IssueEditCommentParams) WithBody(body *models.EditIssueCommentOption) *IssueEditCommentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the issue edit comment params
func (o *IssueEditCommentParams) SetBody(body *models.EditIssueCommentOption) {
	o.Body = body
}

// WithID adds the id to the issue edit comment params
func (o *IssueEditCommentParams) WithID(id int64) *IssueEditCommentParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the issue edit comment params
func (o *IssueEditCommentParams) SetID(id int64) {
	o.ID = id
}

// WithOwner adds the owner to the issue edit comment params
func (o *IssueEditCommentParams) WithOwner(owner string) *IssueEditCommentParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the issue edit comment params
func (o *IssueEditCommentParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the issue edit comment params
func (o *IssueEditCommentParams) WithRepo(repo string) *IssueEditCommentParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the issue edit comment params
func (o *IssueEditCommentParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *IssueEditCommentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
