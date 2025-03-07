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

// NewIssueListIssueDependenciesParams creates a new IssueListIssueDependenciesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIssueListIssueDependenciesParams() *IssueListIssueDependenciesParams {
	return &IssueListIssueDependenciesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIssueListIssueDependenciesParamsWithTimeout creates a new IssueListIssueDependenciesParams object
// with the ability to set a timeout on a request.
func NewIssueListIssueDependenciesParamsWithTimeout(timeout time.Duration) *IssueListIssueDependenciesParams {
	return &IssueListIssueDependenciesParams{
		timeout: timeout,
	}
}

// NewIssueListIssueDependenciesParamsWithContext creates a new IssueListIssueDependenciesParams object
// with the ability to set a context for a request.
func NewIssueListIssueDependenciesParamsWithContext(ctx context.Context) *IssueListIssueDependenciesParams {
	return &IssueListIssueDependenciesParams{
		Context: ctx,
	}
}

// NewIssueListIssueDependenciesParamsWithHTTPClient creates a new IssueListIssueDependenciesParams object
// with the ability to set a custom HTTPClient for a request.
func NewIssueListIssueDependenciesParamsWithHTTPClient(client *http.Client) *IssueListIssueDependenciesParams {
	return &IssueListIssueDependenciesParams{
		HTTPClient: client,
	}
}

/*
IssueListIssueDependenciesParams contains all the parameters to send to the API endpoint

	for the issue list issue dependencies operation.

	Typically these are written to a http.Request.
*/
type IssueListIssueDependenciesParams struct {

	/* Index.

	   index of the issue
	*/
	Index string

	/* Limit.

	   page size of results
	*/
	Limit *int64

	/* Owner.

	   owner of the repo
	*/
	Owner string

	/* Page.

	   page number of results to return (1-based)
	*/
	Page *int64

	/* Repo.

	   name of the repo
	*/
	Repo string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the issue list issue dependencies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueListIssueDependenciesParams) WithDefaults() *IssueListIssueDependenciesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the issue list issue dependencies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueListIssueDependenciesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the issue list issue dependencies params
func (o *IssueListIssueDependenciesParams) WithTimeout(timeout time.Duration) *IssueListIssueDependenciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the issue list issue dependencies params
func (o *IssueListIssueDependenciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the issue list issue dependencies params
func (o *IssueListIssueDependenciesParams) WithContext(ctx context.Context) *IssueListIssueDependenciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the issue list issue dependencies params
func (o *IssueListIssueDependenciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the issue list issue dependencies params
func (o *IssueListIssueDependenciesParams) WithHTTPClient(client *http.Client) *IssueListIssueDependenciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the issue list issue dependencies params
func (o *IssueListIssueDependenciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIndex adds the index to the issue list issue dependencies params
func (o *IssueListIssueDependenciesParams) WithIndex(index string) *IssueListIssueDependenciesParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the issue list issue dependencies params
func (o *IssueListIssueDependenciesParams) SetIndex(index string) {
	o.Index = index
}

// WithLimit adds the limit to the issue list issue dependencies params
func (o *IssueListIssueDependenciesParams) WithLimit(limit *int64) *IssueListIssueDependenciesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the issue list issue dependencies params
func (o *IssueListIssueDependenciesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOwner adds the owner to the issue list issue dependencies params
func (o *IssueListIssueDependenciesParams) WithOwner(owner string) *IssueListIssueDependenciesParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the issue list issue dependencies params
func (o *IssueListIssueDependenciesParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithPage adds the page to the issue list issue dependencies params
func (o *IssueListIssueDependenciesParams) WithPage(page *int64) *IssueListIssueDependenciesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the issue list issue dependencies params
func (o *IssueListIssueDependenciesParams) SetPage(page *int64) {
	o.Page = page
}

// WithRepo adds the repo to the issue list issue dependencies params
func (o *IssueListIssueDependenciesParams) WithRepo(repo string) *IssueListIssueDependenciesParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the issue list issue dependencies params
func (o *IssueListIssueDependenciesParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *IssueListIssueDependenciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param index
	if err := r.SetPathParam("index", o.Index); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
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
