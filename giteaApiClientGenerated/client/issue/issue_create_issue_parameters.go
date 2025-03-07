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

// NewIssueCreateIssueParams creates a new IssueCreateIssueParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIssueCreateIssueParams() *IssueCreateIssueParams {
	return &IssueCreateIssueParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIssueCreateIssueParamsWithTimeout creates a new IssueCreateIssueParams object
// with the ability to set a timeout on a request.
func NewIssueCreateIssueParamsWithTimeout(timeout time.Duration) *IssueCreateIssueParams {
	return &IssueCreateIssueParams{
		timeout: timeout,
	}
}

// NewIssueCreateIssueParamsWithContext creates a new IssueCreateIssueParams object
// with the ability to set a context for a request.
func NewIssueCreateIssueParamsWithContext(ctx context.Context) *IssueCreateIssueParams {
	return &IssueCreateIssueParams{
		Context: ctx,
	}
}

// NewIssueCreateIssueParamsWithHTTPClient creates a new IssueCreateIssueParams object
// with the ability to set a custom HTTPClient for a request.
func NewIssueCreateIssueParamsWithHTTPClient(client *http.Client) *IssueCreateIssueParams {
	return &IssueCreateIssueParams{
		HTTPClient: client,
	}
}

/*
IssueCreateIssueParams contains all the parameters to send to the API endpoint

	for the issue create issue operation.

	Typically these are written to a http.Request.
*/
type IssueCreateIssueParams struct {

	// Body.
	Body *models.CreateIssueOption

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

// WithDefaults hydrates default values in the issue create issue params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueCreateIssueParams) WithDefaults() *IssueCreateIssueParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the issue create issue params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueCreateIssueParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the issue create issue params
func (o *IssueCreateIssueParams) WithTimeout(timeout time.Duration) *IssueCreateIssueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the issue create issue params
func (o *IssueCreateIssueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the issue create issue params
func (o *IssueCreateIssueParams) WithContext(ctx context.Context) *IssueCreateIssueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the issue create issue params
func (o *IssueCreateIssueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the issue create issue params
func (o *IssueCreateIssueParams) WithHTTPClient(client *http.Client) *IssueCreateIssueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the issue create issue params
func (o *IssueCreateIssueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the issue create issue params
func (o *IssueCreateIssueParams) WithBody(body *models.CreateIssueOption) *IssueCreateIssueParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the issue create issue params
func (o *IssueCreateIssueParams) SetBody(body *models.CreateIssueOption) {
	o.Body = body
}

// WithOwner adds the owner to the issue create issue params
func (o *IssueCreateIssueParams) WithOwner(owner string) *IssueCreateIssueParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the issue create issue params
func (o *IssueCreateIssueParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the issue create issue params
func (o *IssueCreateIssueParams) WithRepo(repo string) *IssueCreateIssueParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the issue create issue params
func (o *IssueCreateIssueParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *IssueCreateIssueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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
