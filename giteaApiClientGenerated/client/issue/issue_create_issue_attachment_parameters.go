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

// NewIssueCreateIssueAttachmentParams creates a new IssueCreateIssueAttachmentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIssueCreateIssueAttachmentParams() *IssueCreateIssueAttachmentParams {
	return &IssueCreateIssueAttachmentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIssueCreateIssueAttachmentParamsWithTimeout creates a new IssueCreateIssueAttachmentParams object
// with the ability to set a timeout on a request.
func NewIssueCreateIssueAttachmentParamsWithTimeout(timeout time.Duration) *IssueCreateIssueAttachmentParams {
	return &IssueCreateIssueAttachmentParams{
		timeout: timeout,
	}
}

// NewIssueCreateIssueAttachmentParamsWithContext creates a new IssueCreateIssueAttachmentParams object
// with the ability to set a context for a request.
func NewIssueCreateIssueAttachmentParamsWithContext(ctx context.Context) *IssueCreateIssueAttachmentParams {
	return &IssueCreateIssueAttachmentParams{
		Context: ctx,
	}
}

// NewIssueCreateIssueAttachmentParamsWithHTTPClient creates a new IssueCreateIssueAttachmentParams object
// with the ability to set a custom HTTPClient for a request.
func NewIssueCreateIssueAttachmentParamsWithHTTPClient(client *http.Client) *IssueCreateIssueAttachmentParams {
	return &IssueCreateIssueAttachmentParams{
		HTTPClient: client,
	}
}

/*
IssueCreateIssueAttachmentParams contains all the parameters to send to the API endpoint

	for the issue create issue attachment operation.

	Typically these are written to a http.Request.
*/
type IssueCreateIssueAttachmentParams struct {

	/* Attachment.

	   attachment to upload
	*/
	Attachment runtime.NamedReadCloser

	/* Index.

	   index of the issue

	   Format: int64
	*/
	Index int64

	/* Name.

	   name of the attachment
	*/
	Name *string

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

// WithDefaults hydrates default values in the issue create issue attachment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueCreateIssueAttachmentParams) WithDefaults() *IssueCreateIssueAttachmentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the issue create issue attachment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueCreateIssueAttachmentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the issue create issue attachment params
func (o *IssueCreateIssueAttachmentParams) WithTimeout(timeout time.Duration) *IssueCreateIssueAttachmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the issue create issue attachment params
func (o *IssueCreateIssueAttachmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the issue create issue attachment params
func (o *IssueCreateIssueAttachmentParams) WithContext(ctx context.Context) *IssueCreateIssueAttachmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the issue create issue attachment params
func (o *IssueCreateIssueAttachmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the issue create issue attachment params
func (o *IssueCreateIssueAttachmentParams) WithHTTPClient(client *http.Client) *IssueCreateIssueAttachmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the issue create issue attachment params
func (o *IssueCreateIssueAttachmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAttachment adds the attachment to the issue create issue attachment params
func (o *IssueCreateIssueAttachmentParams) WithAttachment(attachment runtime.NamedReadCloser) *IssueCreateIssueAttachmentParams {
	o.SetAttachment(attachment)
	return o
}

// SetAttachment adds the attachment to the issue create issue attachment params
func (o *IssueCreateIssueAttachmentParams) SetAttachment(attachment runtime.NamedReadCloser) {
	o.Attachment = attachment
}

// WithIndex adds the index to the issue create issue attachment params
func (o *IssueCreateIssueAttachmentParams) WithIndex(index int64) *IssueCreateIssueAttachmentParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the issue create issue attachment params
func (o *IssueCreateIssueAttachmentParams) SetIndex(index int64) {
	o.Index = index
}

// WithName adds the name to the issue create issue attachment params
func (o *IssueCreateIssueAttachmentParams) WithName(name *string) *IssueCreateIssueAttachmentParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the issue create issue attachment params
func (o *IssueCreateIssueAttachmentParams) SetName(name *string) {
	o.Name = name
}

// WithOwner adds the owner to the issue create issue attachment params
func (o *IssueCreateIssueAttachmentParams) WithOwner(owner string) *IssueCreateIssueAttachmentParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the issue create issue attachment params
func (o *IssueCreateIssueAttachmentParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the issue create issue attachment params
func (o *IssueCreateIssueAttachmentParams) WithRepo(repo string) *IssueCreateIssueAttachmentParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the issue create issue attachment params
func (o *IssueCreateIssueAttachmentParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *IssueCreateIssueAttachmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	// form file param attachment
	if err := r.SetFileParam("attachment", o.Attachment); err != nil {
		return err
	}

	// path param index
	if err := r.SetPathParam("index", swag.FormatInt64(o.Index)); err != nil {
		return err
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
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
