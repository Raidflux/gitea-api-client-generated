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

	"giteaApiClientGenerated/models"
)

// NewIssueCreateIssueBlockingParams creates a new IssueCreateIssueBlockingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIssueCreateIssueBlockingParams() *IssueCreateIssueBlockingParams {
	return &IssueCreateIssueBlockingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIssueCreateIssueBlockingParamsWithTimeout creates a new IssueCreateIssueBlockingParams object
// with the ability to set a timeout on a request.
func NewIssueCreateIssueBlockingParamsWithTimeout(timeout time.Duration) *IssueCreateIssueBlockingParams {
	return &IssueCreateIssueBlockingParams{
		timeout: timeout,
	}
}

// NewIssueCreateIssueBlockingParamsWithContext creates a new IssueCreateIssueBlockingParams object
// with the ability to set a context for a request.
func NewIssueCreateIssueBlockingParamsWithContext(ctx context.Context) *IssueCreateIssueBlockingParams {
	return &IssueCreateIssueBlockingParams{
		Context: ctx,
	}
}

// NewIssueCreateIssueBlockingParamsWithHTTPClient creates a new IssueCreateIssueBlockingParams object
// with the ability to set a custom HTTPClient for a request.
func NewIssueCreateIssueBlockingParamsWithHTTPClient(client *http.Client) *IssueCreateIssueBlockingParams {
	return &IssueCreateIssueBlockingParams{
		HTTPClient: client,
	}
}

/*
IssueCreateIssueBlockingParams contains all the parameters to send to the API endpoint

	for the issue create issue blocking operation.

	Typically these are written to a http.Request.
*/
type IssueCreateIssueBlockingParams struct {

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

// WithDefaults hydrates default values in the issue create issue blocking params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueCreateIssueBlockingParams) WithDefaults() *IssueCreateIssueBlockingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the issue create issue blocking params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueCreateIssueBlockingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the issue create issue blocking params
func (o *IssueCreateIssueBlockingParams) WithTimeout(timeout time.Duration) *IssueCreateIssueBlockingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the issue create issue blocking params
func (o *IssueCreateIssueBlockingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the issue create issue blocking params
func (o *IssueCreateIssueBlockingParams) WithContext(ctx context.Context) *IssueCreateIssueBlockingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the issue create issue blocking params
func (o *IssueCreateIssueBlockingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the issue create issue blocking params
func (o *IssueCreateIssueBlockingParams) WithHTTPClient(client *http.Client) *IssueCreateIssueBlockingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the issue create issue blocking params
func (o *IssueCreateIssueBlockingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the issue create issue blocking params
func (o *IssueCreateIssueBlockingParams) WithBody(body *models.IssueMeta) *IssueCreateIssueBlockingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the issue create issue blocking params
func (o *IssueCreateIssueBlockingParams) SetBody(body *models.IssueMeta) {
	o.Body = body
}

// WithIndex adds the index to the issue create issue blocking params
func (o *IssueCreateIssueBlockingParams) WithIndex(index string) *IssueCreateIssueBlockingParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the issue create issue blocking params
func (o *IssueCreateIssueBlockingParams) SetIndex(index string) {
	o.Index = index
}

// WithOwner adds the owner to the issue create issue blocking params
func (o *IssueCreateIssueBlockingParams) WithOwner(owner string) *IssueCreateIssueBlockingParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the issue create issue blocking params
func (o *IssueCreateIssueBlockingParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the issue create issue blocking params
func (o *IssueCreateIssueBlockingParams) WithRepo(repo string) *IssueCreateIssueBlockingParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the issue create issue blocking params
func (o *IssueCreateIssueBlockingParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *IssueCreateIssueBlockingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
