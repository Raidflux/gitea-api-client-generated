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

// NewIssueCreateMilestoneParams creates a new IssueCreateMilestoneParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIssueCreateMilestoneParams() *IssueCreateMilestoneParams {
	return &IssueCreateMilestoneParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIssueCreateMilestoneParamsWithTimeout creates a new IssueCreateMilestoneParams object
// with the ability to set a timeout on a request.
func NewIssueCreateMilestoneParamsWithTimeout(timeout time.Duration) *IssueCreateMilestoneParams {
	return &IssueCreateMilestoneParams{
		timeout: timeout,
	}
}

// NewIssueCreateMilestoneParamsWithContext creates a new IssueCreateMilestoneParams object
// with the ability to set a context for a request.
func NewIssueCreateMilestoneParamsWithContext(ctx context.Context) *IssueCreateMilestoneParams {
	return &IssueCreateMilestoneParams{
		Context: ctx,
	}
}

// NewIssueCreateMilestoneParamsWithHTTPClient creates a new IssueCreateMilestoneParams object
// with the ability to set a custom HTTPClient for a request.
func NewIssueCreateMilestoneParamsWithHTTPClient(client *http.Client) *IssueCreateMilestoneParams {
	return &IssueCreateMilestoneParams{
		HTTPClient: client,
	}
}

/*
IssueCreateMilestoneParams contains all the parameters to send to the API endpoint

	for the issue create milestone operation.

	Typically these are written to a http.Request.
*/
type IssueCreateMilestoneParams struct {

	// Body.
	Body *models.CreateMilestoneOption

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

// WithDefaults hydrates default values in the issue create milestone params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueCreateMilestoneParams) WithDefaults() *IssueCreateMilestoneParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the issue create milestone params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueCreateMilestoneParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the issue create milestone params
func (o *IssueCreateMilestoneParams) WithTimeout(timeout time.Duration) *IssueCreateMilestoneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the issue create milestone params
func (o *IssueCreateMilestoneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the issue create milestone params
func (o *IssueCreateMilestoneParams) WithContext(ctx context.Context) *IssueCreateMilestoneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the issue create milestone params
func (o *IssueCreateMilestoneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the issue create milestone params
func (o *IssueCreateMilestoneParams) WithHTTPClient(client *http.Client) *IssueCreateMilestoneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the issue create milestone params
func (o *IssueCreateMilestoneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the issue create milestone params
func (o *IssueCreateMilestoneParams) WithBody(body *models.CreateMilestoneOption) *IssueCreateMilestoneParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the issue create milestone params
func (o *IssueCreateMilestoneParams) SetBody(body *models.CreateMilestoneOption) {
	o.Body = body
}

// WithOwner adds the owner to the issue create milestone params
func (o *IssueCreateMilestoneParams) WithOwner(owner string) *IssueCreateMilestoneParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the issue create milestone params
func (o *IssueCreateMilestoneParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the issue create milestone params
func (o *IssueCreateMilestoneParams) WithRepo(repo string) *IssueCreateMilestoneParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the issue create milestone params
func (o *IssueCreateMilestoneParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *IssueCreateMilestoneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
