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

// NewGetBlobParams creates a new GetBlobParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBlobParams() *GetBlobParams {
	return &GetBlobParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBlobParamsWithTimeout creates a new GetBlobParams object
// with the ability to set a timeout on a request.
func NewGetBlobParamsWithTimeout(timeout time.Duration) *GetBlobParams {
	return &GetBlobParams{
		timeout: timeout,
	}
}

// NewGetBlobParamsWithContext creates a new GetBlobParams object
// with the ability to set a context for a request.
func NewGetBlobParamsWithContext(ctx context.Context) *GetBlobParams {
	return &GetBlobParams{
		Context: ctx,
	}
}

// NewGetBlobParamsWithHTTPClient creates a new GetBlobParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBlobParamsWithHTTPClient(client *http.Client) *GetBlobParams {
	return &GetBlobParams{
		HTTPClient: client,
	}
}

/*
GetBlobParams contains all the parameters to send to the API endpoint

	for the get blob operation.

	Typically these are written to a http.Request.
*/
type GetBlobParams struct {

	/* Owner.

	   owner of the repo
	*/
	Owner string

	/* Repo.

	   name of the repo
	*/
	Repo string

	/* Sha.

	   sha of the commit
	*/
	Sha string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get blob params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBlobParams) WithDefaults() *GetBlobParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get blob params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBlobParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get blob params
func (o *GetBlobParams) WithTimeout(timeout time.Duration) *GetBlobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get blob params
func (o *GetBlobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get blob params
func (o *GetBlobParams) WithContext(ctx context.Context) *GetBlobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get blob params
func (o *GetBlobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get blob params
func (o *GetBlobParams) WithHTTPClient(client *http.Client) *GetBlobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get blob params
func (o *GetBlobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the get blob params
func (o *GetBlobParams) WithOwner(owner string) *GetBlobParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the get blob params
func (o *GetBlobParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the get blob params
func (o *GetBlobParams) WithRepo(repo string) *GetBlobParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the get blob params
func (o *GetBlobParams) SetRepo(repo string) {
	o.Repo = repo
}

// WithSha adds the sha to the get blob params
func (o *GetBlobParams) WithSha(sha string) *GetBlobParams {
	o.SetSha(sha)
	return o
}

// SetSha adds the sha to the get blob params
func (o *GetBlobParams) SetSha(sha string) {
	o.Sha = sha
}

// WriteToRequest writes these params to a swagger request
func (o *GetBlobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
