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

// NewUnpinIssueParams creates a new UnpinIssueParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUnpinIssueParams() *UnpinIssueParams {
	return &UnpinIssueParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUnpinIssueParamsWithTimeout creates a new UnpinIssueParams object
// with the ability to set a timeout on a request.
func NewUnpinIssueParamsWithTimeout(timeout time.Duration) *UnpinIssueParams {
	return &UnpinIssueParams{
		timeout: timeout,
	}
}

// NewUnpinIssueParamsWithContext creates a new UnpinIssueParams object
// with the ability to set a context for a request.
func NewUnpinIssueParamsWithContext(ctx context.Context) *UnpinIssueParams {
	return &UnpinIssueParams{
		Context: ctx,
	}
}

// NewUnpinIssueParamsWithHTTPClient creates a new UnpinIssueParams object
// with the ability to set a custom HTTPClient for a request.
func NewUnpinIssueParamsWithHTTPClient(client *http.Client) *UnpinIssueParams {
	return &UnpinIssueParams{
		HTTPClient: client,
	}
}

/*
UnpinIssueParams contains all the parameters to send to the API endpoint

	for the unpin issue operation.

	Typically these are written to a http.Request.
*/
type UnpinIssueParams struct {

	/* Index.

	   index of issue to unpin

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

// WithDefaults hydrates default values in the unpin issue params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnpinIssueParams) WithDefaults() *UnpinIssueParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the unpin issue params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnpinIssueParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the unpin issue params
func (o *UnpinIssueParams) WithTimeout(timeout time.Duration) *UnpinIssueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unpin issue params
func (o *UnpinIssueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unpin issue params
func (o *UnpinIssueParams) WithContext(ctx context.Context) *UnpinIssueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unpin issue params
func (o *UnpinIssueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unpin issue params
func (o *UnpinIssueParams) WithHTTPClient(client *http.Client) *UnpinIssueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unpin issue params
func (o *UnpinIssueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIndex adds the index to the unpin issue params
func (o *UnpinIssueParams) WithIndex(index int64) *UnpinIssueParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the unpin issue params
func (o *UnpinIssueParams) SetIndex(index int64) {
	o.Index = index
}

// WithOwner adds the owner to the unpin issue params
func (o *UnpinIssueParams) WithOwner(owner string) *UnpinIssueParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the unpin issue params
func (o *UnpinIssueParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the unpin issue params
func (o *UnpinIssueParams) WithRepo(repo string) *UnpinIssueParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the unpin issue params
func (o *UnpinIssueParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *UnpinIssueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
