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

// NewIssueEditCommentDeprecatedParams creates a new IssueEditCommentDeprecatedParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIssueEditCommentDeprecatedParams() *IssueEditCommentDeprecatedParams {
	return &IssueEditCommentDeprecatedParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIssueEditCommentDeprecatedParamsWithTimeout creates a new IssueEditCommentDeprecatedParams object
// with the ability to set a timeout on a request.
func NewIssueEditCommentDeprecatedParamsWithTimeout(timeout time.Duration) *IssueEditCommentDeprecatedParams {
	return &IssueEditCommentDeprecatedParams{
		timeout: timeout,
	}
}

// NewIssueEditCommentDeprecatedParamsWithContext creates a new IssueEditCommentDeprecatedParams object
// with the ability to set a context for a request.
func NewIssueEditCommentDeprecatedParamsWithContext(ctx context.Context) *IssueEditCommentDeprecatedParams {
	return &IssueEditCommentDeprecatedParams{
		Context: ctx,
	}
}

// NewIssueEditCommentDeprecatedParamsWithHTTPClient creates a new IssueEditCommentDeprecatedParams object
// with the ability to set a custom HTTPClient for a request.
func NewIssueEditCommentDeprecatedParamsWithHTTPClient(client *http.Client) *IssueEditCommentDeprecatedParams {
	return &IssueEditCommentDeprecatedParams{
		HTTPClient: client,
	}
}

/*
IssueEditCommentDeprecatedParams contains all the parameters to send to the API endpoint

	for the issue edit comment deprecated operation.

	Typically these are written to a http.Request.
*/
type IssueEditCommentDeprecatedParams struct {

	// Body.
	Body *models.EditIssueCommentOption

	/* ID.

	   id of the comment to edit

	   Format: int64
	*/
	ID int64

	/* Index.

	   this parameter is ignored
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

// WithDefaults hydrates default values in the issue edit comment deprecated params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueEditCommentDeprecatedParams) WithDefaults() *IssueEditCommentDeprecatedParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the issue edit comment deprecated params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueEditCommentDeprecatedParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the issue edit comment deprecated params
func (o *IssueEditCommentDeprecatedParams) WithTimeout(timeout time.Duration) *IssueEditCommentDeprecatedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the issue edit comment deprecated params
func (o *IssueEditCommentDeprecatedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the issue edit comment deprecated params
func (o *IssueEditCommentDeprecatedParams) WithContext(ctx context.Context) *IssueEditCommentDeprecatedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the issue edit comment deprecated params
func (o *IssueEditCommentDeprecatedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the issue edit comment deprecated params
func (o *IssueEditCommentDeprecatedParams) WithHTTPClient(client *http.Client) *IssueEditCommentDeprecatedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the issue edit comment deprecated params
func (o *IssueEditCommentDeprecatedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the issue edit comment deprecated params
func (o *IssueEditCommentDeprecatedParams) WithBody(body *models.EditIssueCommentOption) *IssueEditCommentDeprecatedParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the issue edit comment deprecated params
func (o *IssueEditCommentDeprecatedParams) SetBody(body *models.EditIssueCommentOption) {
	o.Body = body
}

// WithID adds the id to the issue edit comment deprecated params
func (o *IssueEditCommentDeprecatedParams) WithID(id int64) *IssueEditCommentDeprecatedParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the issue edit comment deprecated params
func (o *IssueEditCommentDeprecatedParams) SetID(id int64) {
	o.ID = id
}

// WithIndex adds the index to the issue edit comment deprecated params
func (o *IssueEditCommentDeprecatedParams) WithIndex(index int64) *IssueEditCommentDeprecatedParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the issue edit comment deprecated params
func (o *IssueEditCommentDeprecatedParams) SetIndex(index int64) {
	o.Index = index
}

// WithOwner adds the owner to the issue edit comment deprecated params
func (o *IssueEditCommentDeprecatedParams) WithOwner(owner string) *IssueEditCommentDeprecatedParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the issue edit comment deprecated params
func (o *IssueEditCommentDeprecatedParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the issue edit comment deprecated params
func (o *IssueEditCommentDeprecatedParams) WithRepo(repo string) *IssueEditCommentDeprecatedParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the issue edit comment deprecated params
func (o *IssueEditCommentDeprecatedParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *IssueEditCommentDeprecatedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
