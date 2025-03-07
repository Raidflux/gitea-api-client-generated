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

// NewIssueGetIssueCommentAttachmentParams creates a new IssueGetIssueCommentAttachmentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIssueGetIssueCommentAttachmentParams() *IssueGetIssueCommentAttachmentParams {
	return &IssueGetIssueCommentAttachmentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIssueGetIssueCommentAttachmentParamsWithTimeout creates a new IssueGetIssueCommentAttachmentParams object
// with the ability to set a timeout on a request.
func NewIssueGetIssueCommentAttachmentParamsWithTimeout(timeout time.Duration) *IssueGetIssueCommentAttachmentParams {
	return &IssueGetIssueCommentAttachmentParams{
		timeout: timeout,
	}
}

// NewIssueGetIssueCommentAttachmentParamsWithContext creates a new IssueGetIssueCommentAttachmentParams object
// with the ability to set a context for a request.
func NewIssueGetIssueCommentAttachmentParamsWithContext(ctx context.Context) *IssueGetIssueCommentAttachmentParams {
	return &IssueGetIssueCommentAttachmentParams{
		Context: ctx,
	}
}

// NewIssueGetIssueCommentAttachmentParamsWithHTTPClient creates a new IssueGetIssueCommentAttachmentParams object
// with the ability to set a custom HTTPClient for a request.
func NewIssueGetIssueCommentAttachmentParamsWithHTTPClient(client *http.Client) *IssueGetIssueCommentAttachmentParams {
	return &IssueGetIssueCommentAttachmentParams{
		HTTPClient: client,
	}
}

/*
IssueGetIssueCommentAttachmentParams contains all the parameters to send to the API endpoint

	for the issue get issue comment attachment operation.

	Typically these are written to a http.Request.
*/
type IssueGetIssueCommentAttachmentParams struct {

	/* AttachmentID.

	   id of the attachment to get

	   Format: int64
	*/
	AttachmentID int64

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

// WithDefaults hydrates default values in the issue get issue comment attachment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueGetIssueCommentAttachmentParams) WithDefaults() *IssueGetIssueCommentAttachmentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the issue get issue comment attachment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueGetIssueCommentAttachmentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the issue get issue comment attachment params
func (o *IssueGetIssueCommentAttachmentParams) WithTimeout(timeout time.Duration) *IssueGetIssueCommentAttachmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the issue get issue comment attachment params
func (o *IssueGetIssueCommentAttachmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the issue get issue comment attachment params
func (o *IssueGetIssueCommentAttachmentParams) WithContext(ctx context.Context) *IssueGetIssueCommentAttachmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the issue get issue comment attachment params
func (o *IssueGetIssueCommentAttachmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the issue get issue comment attachment params
func (o *IssueGetIssueCommentAttachmentParams) WithHTTPClient(client *http.Client) *IssueGetIssueCommentAttachmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the issue get issue comment attachment params
func (o *IssueGetIssueCommentAttachmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAttachmentID adds the attachmentID to the issue get issue comment attachment params
func (o *IssueGetIssueCommentAttachmentParams) WithAttachmentID(attachmentID int64) *IssueGetIssueCommentAttachmentParams {
	o.SetAttachmentID(attachmentID)
	return o
}

// SetAttachmentID adds the attachmentId to the issue get issue comment attachment params
func (o *IssueGetIssueCommentAttachmentParams) SetAttachmentID(attachmentID int64) {
	o.AttachmentID = attachmentID
}

// WithID adds the id to the issue get issue comment attachment params
func (o *IssueGetIssueCommentAttachmentParams) WithID(id int64) *IssueGetIssueCommentAttachmentParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the issue get issue comment attachment params
func (o *IssueGetIssueCommentAttachmentParams) SetID(id int64) {
	o.ID = id
}

// WithOwner adds the owner to the issue get issue comment attachment params
func (o *IssueGetIssueCommentAttachmentParams) WithOwner(owner string) *IssueGetIssueCommentAttachmentParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the issue get issue comment attachment params
func (o *IssueGetIssueCommentAttachmentParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the issue get issue comment attachment params
func (o *IssueGetIssueCommentAttachmentParams) WithRepo(repo string) *IssueGetIssueCommentAttachmentParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the issue get issue comment attachment params
func (o *IssueGetIssueCommentAttachmentParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *IssueGetIssueCommentAttachmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param attachment_id
	if err := r.SetPathParam("attachment_id", swag.FormatInt64(o.AttachmentID)); err != nil {
		return err
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
