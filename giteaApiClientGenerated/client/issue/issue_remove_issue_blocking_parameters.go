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

	"github.com/Raidflux/gitea-api-client-generated/giteaApiClientGenerated/models"
)

// NewIssueRemoveIssueBlockingParams creates a new IssueRemoveIssueBlockingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIssueRemoveIssueBlockingParams() *IssueRemoveIssueBlockingParams {
	return &IssueRemoveIssueBlockingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIssueRemoveIssueBlockingParamsWithTimeout creates a new IssueRemoveIssueBlockingParams object
// with the ability to set a timeout on a request.
func NewIssueRemoveIssueBlockingParamsWithTimeout(timeout time.Duration) *IssueRemoveIssueBlockingParams {
	return &IssueRemoveIssueBlockingParams{
		timeout: timeout,
	}
}

// NewIssueRemoveIssueBlockingParamsWithContext creates a new IssueRemoveIssueBlockingParams object
// with the ability to set a context for a request.
func NewIssueRemoveIssueBlockingParamsWithContext(ctx context.Context) *IssueRemoveIssueBlockingParams {
	return &IssueRemoveIssueBlockingParams{
		Context: ctx,
	}
}

// NewIssueRemoveIssueBlockingParamsWithHTTPClient creates a new IssueRemoveIssueBlockingParams object
// with the ability to set a custom HTTPClient for a request.
func NewIssueRemoveIssueBlockingParamsWithHTTPClient(client *http.Client) *IssueRemoveIssueBlockingParams {
	return &IssueRemoveIssueBlockingParams{
		HTTPClient: client,
	}
}

/*
IssueRemoveIssueBlockingParams contains all the parameters to send to the API endpoint

	for the issue remove issue blocking operation.

	Typically these are written to a http.Request.
*/
type IssueRemoveIssueBlockingParams struct {

	// Body.
	Body *models.IssueMeta

	/* Index.

	   index of the issue
	*/
	Index string

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

// WithDefaults hydrates default values in the issue remove issue blocking params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueRemoveIssueBlockingParams) WithDefaults() *IssueRemoveIssueBlockingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the issue remove issue blocking params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueRemoveIssueBlockingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the issue remove issue blocking params
func (o *IssueRemoveIssueBlockingParams) WithTimeout(timeout time.Duration) *IssueRemoveIssueBlockingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the issue remove issue blocking params
func (o *IssueRemoveIssueBlockingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the issue remove issue blocking params
func (o *IssueRemoveIssueBlockingParams) WithContext(ctx context.Context) *IssueRemoveIssueBlockingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the issue remove issue blocking params
func (o *IssueRemoveIssueBlockingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the issue remove issue blocking params
func (o *IssueRemoveIssueBlockingParams) WithHTTPClient(client *http.Client) *IssueRemoveIssueBlockingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the issue remove issue blocking params
func (o *IssueRemoveIssueBlockingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the issue remove issue blocking params
func (o *IssueRemoveIssueBlockingParams) WithBody(body *models.IssueMeta) *IssueRemoveIssueBlockingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the issue remove issue blocking params
func (o *IssueRemoveIssueBlockingParams) SetBody(body *models.IssueMeta) {
	o.Body = body
}

// WithIndex adds the index to the issue remove issue blocking params
func (o *IssueRemoveIssueBlockingParams) WithIndex(index string) *IssueRemoveIssueBlockingParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the issue remove issue blocking params
func (o *IssueRemoveIssueBlockingParams) SetIndex(index string) {
	o.Index = index
}

// WithOwner adds the owner to the issue remove issue blocking params
func (o *IssueRemoveIssueBlockingParams) WithOwner(owner string) *IssueRemoveIssueBlockingParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the issue remove issue blocking params
func (o *IssueRemoveIssueBlockingParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the issue remove issue blocking params
func (o *IssueRemoveIssueBlockingParams) WithRepo(repo string) *IssueRemoveIssueBlockingParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the issue remove issue blocking params
func (o *IssueRemoveIssueBlockingParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *IssueRemoveIssueBlockingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param index
	if err := r.SetPathParam("index", o.Index); err != nil {
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
