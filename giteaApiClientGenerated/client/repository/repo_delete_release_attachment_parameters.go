// Code generated by go-swagger; DO NOT EDIT.

package repository

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

// NewRepoDeleteReleaseAttachmentParams creates a new RepoDeleteReleaseAttachmentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepoDeleteReleaseAttachmentParams() *RepoDeleteReleaseAttachmentParams {
	return &RepoDeleteReleaseAttachmentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepoDeleteReleaseAttachmentParamsWithTimeout creates a new RepoDeleteReleaseAttachmentParams object
// with the ability to set a timeout on a request.
func NewRepoDeleteReleaseAttachmentParamsWithTimeout(timeout time.Duration) *RepoDeleteReleaseAttachmentParams {
	return &RepoDeleteReleaseAttachmentParams{
		timeout: timeout,
	}
}

// NewRepoDeleteReleaseAttachmentParamsWithContext creates a new RepoDeleteReleaseAttachmentParams object
// with the ability to set a context for a request.
func NewRepoDeleteReleaseAttachmentParamsWithContext(ctx context.Context) *RepoDeleteReleaseAttachmentParams {
	return &RepoDeleteReleaseAttachmentParams{
		Context: ctx,
	}
}

// NewRepoDeleteReleaseAttachmentParamsWithHTTPClient creates a new RepoDeleteReleaseAttachmentParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepoDeleteReleaseAttachmentParamsWithHTTPClient(client *http.Client) *RepoDeleteReleaseAttachmentParams {
	return &RepoDeleteReleaseAttachmentParams{
		HTTPClient: client,
	}
}

/*
RepoDeleteReleaseAttachmentParams contains all the parameters to send to the API endpoint

	for the repo delete release attachment operation.

	Typically these are written to a http.Request.
*/
type RepoDeleteReleaseAttachmentParams struct {

	/* AttachmentID.

	   id of the attachment to delete

	   Format: int64
	*/
	AttachmentID int64

	/* ID.

	   id of the release

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

// WithDefaults hydrates default values in the repo delete release attachment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoDeleteReleaseAttachmentParams) WithDefaults() *RepoDeleteReleaseAttachmentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repo delete release attachment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoDeleteReleaseAttachmentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repo delete release attachment params
func (o *RepoDeleteReleaseAttachmentParams) WithTimeout(timeout time.Duration) *RepoDeleteReleaseAttachmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo delete release attachment params
func (o *RepoDeleteReleaseAttachmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo delete release attachment params
func (o *RepoDeleteReleaseAttachmentParams) WithContext(ctx context.Context) *RepoDeleteReleaseAttachmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo delete release attachment params
func (o *RepoDeleteReleaseAttachmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo delete release attachment params
func (o *RepoDeleteReleaseAttachmentParams) WithHTTPClient(client *http.Client) *RepoDeleteReleaseAttachmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo delete release attachment params
func (o *RepoDeleteReleaseAttachmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAttachmentID adds the attachmentID to the repo delete release attachment params
func (o *RepoDeleteReleaseAttachmentParams) WithAttachmentID(attachmentID int64) *RepoDeleteReleaseAttachmentParams {
	o.SetAttachmentID(attachmentID)
	return o
}

// SetAttachmentID adds the attachmentId to the repo delete release attachment params
func (o *RepoDeleteReleaseAttachmentParams) SetAttachmentID(attachmentID int64) {
	o.AttachmentID = attachmentID
}

// WithID adds the id to the repo delete release attachment params
func (o *RepoDeleteReleaseAttachmentParams) WithID(id int64) *RepoDeleteReleaseAttachmentParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the repo delete release attachment params
func (o *RepoDeleteReleaseAttachmentParams) SetID(id int64) {
	o.ID = id
}

// WithOwner adds the owner to the repo delete release attachment params
func (o *RepoDeleteReleaseAttachmentParams) WithOwner(owner string) *RepoDeleteReleaseAttachmentParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo delete release attachment params
func (o *RepoDeleteReleaseAttachmentParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo delete release attachment params
func (o *RepoDeleteReleaseAttachmentParams) WithRepo(repo string) *RepoDeleteReleaseAttachmentParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo delete release attachment params
func (o *RepoDeleteReleaseAttachmentParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoDeleteReleaseAttachmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
