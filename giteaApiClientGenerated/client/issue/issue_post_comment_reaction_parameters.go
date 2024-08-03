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

	"github.com/Raidflux/gitea-api-client-generated/giteaApiClientGenerated/models"
)

// NewIssuePostCommentReactionParams creates a new IssuePostCommentReactionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIssuePostCommentReactionParams() *IssuePostCommentReactionParams {
	return &IssuePostCommentReactionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIssuePostCommentReactionParamsWithTimeout creates a new IssuePostCommentReactionParams object
// with the ability to set a timeout on a request.
func NewIssuePostCommentReactionParamsWithTimeout(timeout time.Duration) *IssuePostCommentReactionParams {
	return &IssuePostCommentReactionParams{
		timeout: timeout,
	}
}

// NewIssuePostCommentReactionParamsWithContext creates a new IssuePostCommentReactionParams object
// with the ability to set a context for a request.
func NewIssuePostCommentReactionParamsWithContext(ctx context.Context) *IssuePostCommentReactionParams {
	return &IssuePostCommentReactionParams{
		Context: ctx,
	}
}

// NewIssuePostCommentReactionParamsWithHTTPClient creates a new IssuePostCommentReactionParams object
// with the ability to set a custom HTTPClient for a request.
func NewIssuePostCommentReactionParamsWithHTTPClient(client *http.Client) *IssuePostCommentReactionParams {
	return &IssuePostCommentReactionParams{
		HTTPClient: client,
	}
}

/*
IssuePostCommentReactionParams contains all the parameters to send to the API endpoint

	for the issue post comment reaction operation.

	Typically these are written to a http.Request.
*/
type IssuePostCommentReactionParams struct {

	// Content.
	Content *models.EditReactionOption

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

// WithDefaults hydrates default values in the issue post comment reaction params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssuePostCommentReactionParams) WithDefaults() *IssuePostCommentReactionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the issue post comment reaction params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssuePostCommentReactionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the issue post comment reaction params
func (o *IssuePostCommentReactionParams) WithTimeout(timeout time.Duration) *IssuePostCommentReactionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the issue post comment reaction params
func (o *IssuePostCommentReactionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the issue post comment reaction params
func (o *IssuePostCommentReactionParams) WithContext(ctx context.Context) *IssuePostCommentReactionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the issue post comment reaction params
func (o *IssuePostCommentReactionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the issue post comment reaction params
func (o *IssuePostCommentReactionParams) WithHTTPClient(client *http.Client) *IssuePostCommentReactionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the issue post comment reaction params
func (o *IssuePostCommentReactionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContent adds the content to the issue post comment reaction params
func (o *IssuePostCommentReactionParams) WithContent(content *models.EditReactionOption) *IssuePostCommentReactionParams {
	o.SetContent(content)
	return o
}

// SetContent adds the content to the issue post comment reaction params
func (o *IssuePostCommentReactionParams) SetContent(content *models.EditReactionOption) {
	o.Content = content
}

// WithID adds the id to the issue post comment reaction params
func (o *IssuePostCommentReactionParams) WithID(id int64) *IssuePostCommentReactionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the issue post comment reaction params
func (o *IssuePostCommentReactionParams) SetID(id int64) {
	o.ID = id
}

// WithOwner adds the owner to the issue post comment reaction params
func (o *IssuePostCommentReactionParams) WithOwner(owner string) *IssuePostCommentReactionParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the issue post comment reaction params
func (o *IssuePostCommentReactionParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the issue post comment reaction params
func (o *IssuePostCommentReactionParams) WithRepo(repo string) *IssuePostCommentReactionParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the issue post comment reaction params
func (o *IssuePostCommentReactionParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *IssuePostCommentReactionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Content != nil {
		if err := r.SetBodyParam(o.Content); err != nil {
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
