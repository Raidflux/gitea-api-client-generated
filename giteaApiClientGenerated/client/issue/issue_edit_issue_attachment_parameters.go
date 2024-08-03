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

// NewIssueEditIssueAttachmentParams creates a new IssueEditIssueAttachmentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIssueEditIssueAttachmentParams() *IssueEditIssueAttachmentParams {
	return &IssueEditIssueAttachmentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIssueEditIssueAttachmentParamsWithTimeout creates a new IssueEditIssueAttachmentParams object
// with the ability to set a timeout on a request.
func NewIssueEditIssueAttachmentParamsWithTimeout(timeout time.Duration) *IssueEditIssueAttachmentParams {
	return &IssueEditIssueAttachmentParams{
		timeout: timeout,
	}
}

// NewIssueEditIssueAttachmentParamsWithContext creates a new IssueEditIssueAttachmentParams object
// with the ability to set a context for a request.
func NewIssueEditIssueAttachmentParamsWithContext(ctx context.Context) *IssueEditIssueAttachmentParams {
	return &IssueEditIssueAttachmentParams{
		Context: ctx,
	}
}

// NewIssueEditIssueAttachmentParamsWithHTTPClient creates a new IssueEditIssueAttachmentParams object
// with the ability to set a custom HTTPClient for a request.
func NewIssueEditIssueAttachmentParamsWithHTTPClient(client *http.Client) *IssueEditIssueAttachmentParams {
	return &IssueEditIssueAttachmentParams{
		HTTPClient: client,
	}
}

/*
IssueEditIssueAttachmentParams contains all the parameters to send to the API endpoint

	for the issue edit issue attachment operation.

	Typically these are written to a http.Request.
*/
type IssueEditIssueAttachmentParams struct {

	/* AttachmentID.

	   id of the attachment to edit

	   Format: int64
	*/
	AttachmentID int64

	// Body.
	Body *models.EditAttachmentOptions

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

// WithDefaults hydrates default values in the issue edit issue attachment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueEditIssueAttachmentParams) WithDefaults() *IssueEditIssueAttachmentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the issue edit issue attachment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueEditIssueAttachmentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the issue edit issue attachment params
func (o *IssueEditIssueAttachmentParams) WithTimeout(timeout time.Duration) *IssueEditIssueAttachmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the issue edit issue attachment params
func (o *IssueEditIssueAttachmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the issue edit issue attachment params
func (o *IssueEditIssueAttachmentParams) WithContext(ctx context.Context) *IssueEditIssueAttachmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the issue edit issue attachment params
func (o *IssueEditIssueAttachmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the issue edit issue attachment params
func (o *IssueEditIssueAttachmentParams) WithHTTPClient(client *http.Client) *IssueEditIssueAttachmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the issue edit issue attachment params
func (o *IssueEditIssueAttachmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAttachmentID adds the attachmentID to the issue edit issue attachment params
func (o *IssueEditIssueAttachmentParams) WithAttachmentID(attachmentID int64) *IssueEditIssueAttachmentParams {
	o.SetAttachmentID(attachmentID)
	return o
}

// SetAttachmentID adds the attachmentId to the issue edit issue attachment params
func (o *IssueEditIssueAttachmentParams) SetAttachmentID(attachmentID int64) {
	o.AttachmentID = attachmentID
}

// WithBody adds the body to the issue edit issue attachment params
func (o *IssueEditIssueAttachmentParams) WithBody(body *models.EditAttachmentOptions) *IssueEditIssueAttachmentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the issue edit issue attachment params
func (o *IssueEditIssueAttachmentParams) SetBody(body *models.EditAttachmentOptions) {
	o.Body = body
}

// WithIndex adds the index to the issue edit issue attachment params
func (o *IssueEditIssueAttachmentParams) WithIndex(index int64) *IssueEditIssueAttachmentParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the issue edit issue attachment params
func (o *IssueEditIssueAttachmentParams) SetIndex(index int64) {
	o.Index = index
}

// WithOwner adds the owner to the issue edit issue attachment params
func (o *IssueEditIssueAttachmentParams) WithOwner(owner string) *IssueEditIssueAttachmentParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the issue edit issue attachment params
func (o *IssueEditIssueAttachmentParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the issue edit issue attachment params
func (o *IssueEditIssueAttachmentParams) WithRepo(repo string) *IssueEditIssueAttachmentParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the issue edit issue attachment params
func (o *IssueEditIssueAttachmentParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *IssueEditIssueAttachmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
