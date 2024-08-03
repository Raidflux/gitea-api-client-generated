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

// NewIssueEditIssueCommentAttachmentParams creates a new IssueEditIssueCommentAttachmentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIssueEditIssueCommentAttachmentParams() *IssueEditIssueCommentAttachmentParams {
	return &IssueEditIssueCommentAttachmentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIssueEditIssueCommentAttachmentParamsWithTimeout creates a new IssueEditIssueCommentAttachmentParams object
// with the ability to set a timeout on a request.
func NewIssueEditIssueCommentAttachmentParamsWithTimeout(timeout time.Duration) *IssueEditIssueCommentAttachmentParams {
	return &IssueEditIssueCommentAttachmentParams{
		timeout: timeout,
	}
}

// NewIssueEditIssueCommentAttachmentParamsWithContext creates a new IssueEditIssueCommentAttachmentParams object
// with the ability to set a context for a request.
func NewIssueEditIssueCommentAttachmentParamsWithContext(ctx context.Context) *IssueEditIssueCommentAttachmentParams {
	return &IssueEditIssueCommentAttachmentParams{
		Context: ctx,
	}
}

// NewIssueEditIssueCommentAttachmentParamsWithHTTPClient creates a new IssueEditIssueCommentAttachmentParams object
// with the ability to set a custom HTTPClient for a request.
func NewIssueEditIssueCommentAttachmentParamsWithHTTPClient(client *http.Client) *IssueEditIssueCommentAttachmentParams {
	return &IssueEditIssueCommentAttachmentParams{
		HTTPClient: client,
	}
}

/*
IssueEditIssueCommentAttachmentParams contains all the parameters to send to the API endpoint

	for the issue edit issue comment attachment operation.

	Typically these are written to a http.Request.
*/
type IssueEditIssueCommentAttachmentParams struct {

	/* AttachmentID.

	   id of the attachment to edit

	   Format: int64
	*/
	AttachmentID int64

	// Body.
	Body *models.EditAttachmentOptions

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

// WithDefaults hydrates default values in the issue edit issue comment attachment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueEditIssueCommentAttachmentParams) WithDefaults() *IssueEditIssueCommentAttachmentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the issue edit issue comment attachment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueEditIssueCommentAttachmentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the issue edit issue comment attachment params
func (o *IssueEditIssueCommentAttachmentParams) WithTimeout(timeout time.Duration) *IssueEditIssueCommentAttachmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the issue edit issue comment attachment params
func (o *IssueEditIssueCommentAttachmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the issue edit issue comment attachment params
func (o *IssueEditIssueCommentAttachmentParams) WithContext(ctx context.Context) *IssueEditIssueCommentAttachmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the issue edit issue comment attachment params
func (o *IssueEditIssueCommentAttachmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the issue edit issue comment attachment params
func (o *IssueEditIssueCommentAttachmentParams) WithHTTPClient(client *http.Client) *IssueEditIssueCommentAttachmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the issue edit issue comment attachment params
func (o *IssueEditIssueCommentAttachmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAttachmentID adds the attachmentID to the issue edit issue comment attachment params
func (o *IssueEditIssueCommentAttachmentParams) WithAttachmentID(attachmentID int64) *IssueEditIssueCommentAttachmentParams {
	o.SetAttachmentID(attachmentID)
	return o
}

// SetAttachmentID adds the attachmentId to the issue edit issue comment attachment params
func (o *IssueEditIssueCommentAttachmentParams) SetAttachmentID(attachmentID int64) {
	o.AttachmentID = attachmentID
}

// WithBody adds the body to the issue edit issue comment attachment params
func (o *IssueEditIssueCommentAttachmentParams) WithBody(body *models.EditAttachmentOptions) *IssueEditIssueCommentAttachmentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the issue edit issue comment attachment params
func (o *IssueEditIssueCommentAttachmentParams) SetBody(body *models.EditAttachmentOptions) {
	o.Body = body
}

// WithID adds the id to the issue edit issue comment attachment params
func (o *IssueEditIssueCommentAttachmentParams) WithID(id int64) *IssueEditIssueCommentAttachmentParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the issue edit issue comment attachment params
func (o *IssueEditIssueCommentAttachmentParams) SetID(id int64) {
	o.ID = id
}

// WithOwner adds the owner to the issue edit issue comment attachment params
func (o *IssueEditIssueCommentAttachmentParams) WithOwner(owner string) *IssueEditIssueCommentAttachmentParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the issue edit issue comment attachment params
func (o *IssueEditIssueCommentAttachmentParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the issue edit issue comment attachment params
func (o *IssueEditIssueCommentAttachmentParams) WithRepo(repo string) *IssueEditIssueCommentAttachmentParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the issue edit issue comment attachment params
func (o *IssueEditIssueCommentAttachmentParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *IssueEditIssueCommentAttachmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param attachment_id
	if err := r.SetPathParam("attachment_id", swag.FormatInt64(o.AttachmentID)); err != nil {
		return err
	}
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
