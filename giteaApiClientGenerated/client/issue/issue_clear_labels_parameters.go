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

// NewIssueClearLabelsParams creates a new IssueClearLabelsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIssueClearLabelsParams() *IssueClearLabelsParams {
	return &IssueClearLabelsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIssueClearLabelsParamsWithTimeout creates a new IssueClearLabelsParams object
// with the ability to set a timeout on a request.
func NewIssueClearLabelsParamsWithTimeout(timeout time.Duration) *IssueClearLabelsParams {
	return &IssueClearLabelsParams{
		timeout: timeout,
	}
}

// NewIssueClearLabelsParamsWithContext creates a new IssueClearLabelsParams object
// with the ability to set a context for a request.
func NewIssueClearLabelsParamsWithContext(ctx context.Context) *IssueClearLabelsParams {
	return &IssueClearLabelsParams{
		Context: ctx,
	}
}

// NewIssueClearLabelsParamsWithHTTPClient creates a new IssueClearLabelsParams object
// with the ability to set a custom HTTPClient for a request.
func NewIssueClearLabelsParamsWithHTTPClient(client *http.Client) *IssueClearLabelsParams {
	return &IssueClearLabelsParams{
		HTTPClient: client,
	}
}

/*
IssueClearLabelsParams contains all the parameters to send to the API endpoint

	for the issue clear labels operation.

	Typically these are written to a http.Request.
*/
type IssueClearLabelsParams struct {

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

// WithDefaults hydrates default values in the issue clear labels params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueClearLabelsParams) WithDefaults() *IssueClearLabelsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the issue clear labels params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueClearLabelsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the issue clear labels params
func (o *IssueClearLabelsParams) WithTimeout(timeout time.Duration) *IssueClearLabelsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the issue clear labels params
func (o *IssueClearLabelsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the issue clear labels params
func (o *IssueClearLabelsParams) WithContext(ctx context.Context) *IssueClearLabelsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the issue clear labels params
func (o *IssueClearLabelsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the issue clear labels params
func (o *IssueClearLabelsParams) WithHTTPClient(client *http.Client) *IssueClearLabelsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the issue clear labels params
func (o *IssueClearLabelsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIndex adds the index to the issue clear labels params
func (o *IssueClearLabelsParams) WithIndex(index int64) *IssueClearLabelsParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the issue clear labels params
func (o *IssueClearLabelsParams) SetIndex(index int64) {
	o.Index = index
}

// WithOwner adds the owner to the issue clear labels params
func (o *IssueClearLabelsParams) WithOwner(owner string) *IssueClearLabelsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the issue clear labels params
func (o *IssueClearLabelsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the issue clear labels params
func (o *IssueClearLabelsParams) WithRepo(repo string) *IssueClearLabelsParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the issue clear labels params
func (o *IssueClearLabelsParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *IssueClearLabelsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
