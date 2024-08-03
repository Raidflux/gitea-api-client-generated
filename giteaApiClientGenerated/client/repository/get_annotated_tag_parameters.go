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
)

// NewGetAnnotatedTagParams creates a new GetAnnotatedTagParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAnnotatedTagParams() *GetAnnotatedTagParams {
	return &GetAnnotatedTagParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAnnotatedTagParamsWithTimeout creates a new GetAnnotatedTagParams object
// with the ability to set a timeout on a request.
func NewGetAnnotatedTagParamsWithTimeout(timeout time.Duration) *GetAnnotatedTagParams {
	return &GetAnnotatedTagParams{
		timeout: timeout,
	}
}

// NewGetAnnotatedTagParamsWithContext creates a new GetAnnotatedTagParams object
// with the ability to set a context for a request.
func NewGetAnnotatedTagParamsWithContext(ctx context.Context) *GetAnnotatedTagParams {
	return &GetAnnotatedTagParams{
		Context: ctx,
	}
}

// NewGetAnnotatedTagParamsWithHTTPClient creates a new GetAnnotatedTagParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAnnotatedTagParamsWithHTTPClient(client *http.Client) *GetAnnotatedTagParams {
	return &GetAnnotatedTagParams{
		HTTPClient: client,
	}
}

/*
GetAnnotatedTagParams contains all the parameters to send to the API endpoint

	for the get annotated tag operation.

	Typically these are written to a http.Request.
*/
type GetAnnotatedTagParams struct {

	/* Owner.

	   owner of the repo
	*/
	Owner string

	/* Repo.

	   name of the repo
	*/
	Repo string

	/* Sha.

	   sha of the tag. The Git tags API only supports annotated tag objects, not lightweight tags.
	*/
	Sha string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get annotated tag params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAnnotatedTagParams) WithDefaults() *GetAnnotatedTagParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get annotated tag params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAnnotatedTagParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get annotated tag params
func (o *GetAnnotatedTagParams) WithTimeout(timeout time.Duration) *GetAnnotatedTagParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get annotated tag params
func (o *GetAnnotatedTagParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get annotated tag params
func (o *GetAnnotatedTagParams) WithContext(ctx context.Context) *GetAnnotatedTagParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get annotated tag params
func (o *GetAnnotatedTagParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get annotated tag params
func (o *GetAnnotatedTagParams) WithHTTPClient(client *http.Client) *GetAnnotatedTagParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get annotated tag params
func (o *GetAnnotatedTagParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the get annotated tag params
func (o *GetAnnotatedTagParams) WithOwner(owner string) *GetAnnotatedTagParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the get annotated tag params
func (o *GetAnnotatedTagParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the get annotated tag params
func (o *GetAnnotatedTagParams) WithRepo(repo string) *GetAnnotatedTagParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the get annotated tag params
func (o *GetAnnotatedTagParams) SetRepo(repo string) {
	o.Repo = repo
}

// WithSha adds the sha to the get annotated tag params
func (o *GetAnnotatedTagParams) WithSha(sha string) *GetAnnotatedTagParams {
	o.SetSha(sha)
	return o
}

// SetSha adds the sha to the get annotated tag params
func (o *GetAnnotatedTagParams) SetSha(sha string) {
	o.Sha = sha
}

// WriteToRequest writes these params to a swagger request
func (o *GetAnnotatedTagParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param repo
	if err := r.SetPathParam("repo", o.Repo); err != nil {
		return err
	}

	// path param sha
	if err := r.SetPathParam("sha", o.Sha); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
