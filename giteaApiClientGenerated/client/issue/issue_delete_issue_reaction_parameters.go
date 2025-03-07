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

// NewIssueDeleteIssueReactionParams creates a new IssueDeleteIssueReactionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIssueDeleteIssueReactionParams() *IssueDeleteIssueReactionParams {
	return &IssueDeleteIssueReactionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIssueDeleteIssueReactionParamsWithTimeout creates a new IssueDeleteIssueReactionParams object
// with the ability to set a timeout on a request.
func NewIssueDeleteIssueReactionParamsWithTimeout(timeout time.Duration) *IssueDeleteIssueReactionParams {
	return &IssueDeleteIssueReactionParams{
		timeout: timeout,
	}
}

// NewIssueDeleteIssueReactionParamsWithContext creates a new IssueDeleteIssueReactionParams object
// with the ability to set a context for a request.
func NewIssueDeleteIssueReactionParamsWithContext(ctx context.Context) *IssueDeleteIssueReactionParams {
	return &IssueDeleteIssueReactionParams{
		Context: ctx,
	}
}

// NewIssueDeleteIssueReactionParamsWithHTTPClient creates a new IssueDeleteIssueReactionParams object
// with the ability to set a custom HTTPClient for a request.
func NewIssueDeleteIssueReactionParamsWithHTTPClient(client *http.Client) *IssueDeleteIssueReactionParams {
	return &IssueDeleteIssueReactionParams{
		HTTPClient: client,
	}
}

/*
IssueDeleteIssueReactionParams contains all the parameters to send to the API endpoint

	for the issue delete issue reaction operation.

	Typically these are written to a http.Request.
*/
type IssueDeleteIssueReactionParams struct {

	// Content.
	Content *models.EditReactionOption

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

// WithDefaults hydrates default values in the issue delete issue reaction params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueDeleteIssueReactionParams) WithDefaults() *IssueDeleteIssueReactionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the issue delete issue reaction params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueDeleteIssueReactionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the issue delete issue reaction params
func (o *IssueDeleteIssueReactionParams) WithTimeout(timeout time.Duration) *IssueDeleteIssueReactionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the issue delete issue reaction params
func (o *IssueDeleteIssueReactionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the issue delete issue reaction params
func (o *IssueDeleteIssueReactionParams) WithContext(ctx context.Context) *IssueDeleteIssueReactionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the issue delete issue reaction params
func (o *IssueDeleteIssueReactionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the issue delete issue reaction params
func (o *IssueDeleteIssueReactionParams) WithHTTPClient(client *http.Client) *IssueDeleteIssueReactionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the issue delete issue reaction params
func (o *IssueDeleteIssueReactionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContent adds the content to the issue delete issue reaction params
func (o *IssueDeleteIssueReactionParams) WithContent(content *models.EditReactionOption) *IssueDeleteIssueReactionParams {
	o.SetContent(content)
	return o
}

// SetContent adds the content to the issue delete issue reaction params
func (o *IssueDeleteIssueReactionParams) SetContent(content *models.EditReactionOption) {
	o.Content = content
}

// WithIndex adds the index to the issue delete issue reaction params
func (o *IssueDeleteIssueReactionParams) WithIndex(index int64) *IssueDeleteIssueReactionParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the issue delete issue reaction params
func (o *IssueDeleteIssueReactionParams) SetIndex(index int64) {
	o.Index = index
}

// WithOwner adds the owner to the issue delete issue reaction params
func (o *IssueDeleteIssueReactionParams) WithOwner(owner string) *IssueDeleteIssueReactionParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the issue delete issue reaction params
func (o *IssueDeleteIssueReactionParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the issue delete issue reaction params
func (o *IssueDeleteIssueReactionParams) WithRepo(repo string) *IssueDeleteIssueReactionParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the issue delete issue reaction params
func (o *IssueDeleteIssueReactionParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *IssueDeleteIssueReactionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Content != nil {
		if err := r.SetBodyParam(o.Content); err != nil {
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
