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
	"github.com/go-openapi/swag"
)

// NewRepoGetHookParams creates a new RepoGetHookParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepoGetHookParams() *RepoGetHookParams {
	return &RepoGetHookParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepoGetHookParamsWithTimeout creates a new RepoGetHookParams object
// with the ability to set a timeout on a request.
func NewRepoGetHookParamsWithTimeout(timeout time.Duration) *RepoGetHookParams {
	return &RepoGetHookParams{
		timeout: timeout,
	}
}

// NewRepoGetHookParamsWithContext creates a new RepoGetHookParams object
// with the ability to set a context for a request.
func NewRepoGetHookParamsWithContext(ctx context.Context) *RepoGetHookParams {
	return &RepoGetHookParams{
		Context: ctx,
	}
}

// NewRepoGetHookParamsWithHTTPClient creates a new RepoGetHookParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepoGetHookParamsWithHTTPClient(client *http.Client) *RepoGetHookParams {
	return &RepoGetHookParams{
		HTTPClient: client,
	}
}

/*
RepoGetHookParams contains all the parameters to send to the API endpoint

	for the repo get hook operation.

	Typically these are written to a http.Request.
*/
type RepoGetHookParams struct {

	/* ID.

	   id of the hook to get

	   Format: int64
	*/
	ID int64

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

// WithDefaults hydrates default values in the repo get hook params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoGetHookParams) WithDefaults() *RepoGetHookParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repo get hook params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoGetHookParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repo get hook params
func (o *RepoGetHookParams) WithTimeout(timeout time.Duration) *RepoGetHookParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo get hook params
func (o *RepoGetHookParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo get hook params
func (o *RepoGetHookParams) WithContext(ctx context.Context) *RepoGetHookParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo get hook params
func (o *RepoGetHookParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo get hook params
func (o *RepoGetHookParams) WithHTTPClient(client *http.Client) *RepoGetHookParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo get hook params
func (o *RepoGetHookParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the repo get hook params
func (o *RepoGetHookParams) WithID(id int64) *RepoGetHookParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the repo get hook params
func (o *RepoGetHookParams) SetID(id int64) {
	o.ID = id
}

// WithOwner adds the owner to the repo get hook params
func (o *RepoGetHookParams) WithOwner(owner string) *RepoGetHookParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo get hook params
func (o *RepoGetHookParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo get hook params
func (o *RepoGetHookParams) WithRepo(repo string) *RepoGetHookParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo get hook params
func (o *RepoGetHookParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoGetHookParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
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
