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

// NewRepoListKeysParams creates a new RepoListKeysParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepoListKeysParams() *RepoListKeysParams {
	return &RepoListKeysParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepoListKeysParamsWithTimeout creates a new RepoListKeysParams object
// with the ability to set a timeout on a request.
func NewRepoListKeysParamsWithTimeout(timeout time.Duration) *RepoListKeysParams {
	return &RepoListKeysParams{
		timeout: timeout,
	}
}

// NewRepoListKeysParamsWithContext creates a new RepoListKeysParams object
// with the ability to set a context for a request.
func NewRepoListKeysParamsWithContext(ctx context.Context) *RepoListKeysParams {
	return &RepoListKeysParams{
		Context: ctx,
	}
}

// NewRepoListKeysParamsWithHTTPClient creates a new RepoListKeysParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepoListKeysParamsWithHTTPClient(client *http.Client) *RepoListKeysParams {
	return &RepoListKeysParams{
		HTTPClient: client,
	}
}

/*
RepoListKeysParams contains all the parameters to send to the API endpoint

	for the repo list keys operation.

	Typically these are written to a http.Request.
*/
type RepoListKeysParams struct {

	/* Fingerprint.

	   fingerprint of the key
	*/
	Fingerprint *string

	/* KeyID.

	   the key_id to search for
	*/
	KeyID *int64

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

// WithDefaults hydrates default values in the repo list keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoListKeysParams) WithDefaults() *RepoListKeysParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repo list keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoListKeysParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repo list keys params
func (o *RepoListKeysParams) WithTimeout(timeout time.Duration) *RepoListKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo list keys params
func (o *RepoListKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo list keys params
func (o *RepoListKeysParams) WithContext(ctx context.Context) *RepoListKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo list keys params
func (o *RepoListKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo list keys params
func (o *RepoListKeysParams) WithHTTPClient(client *http.Client) *RepoListKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo list keys params
func (o *RepoListKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFingerprint adds the fingerprint to the repo list keys params
func (o *RepoListKeysParams) WithFingerprint(fingerprint *string) *RepoListKeysParams {
	o.SetFingerprint(fingerprint)
	return o
}

// SetFingerprint adds the fingerprint to the repo list keys params
func (o *RepoListKeysParams) SetFingerprint(fingerprint *string) {
	o.Fingerprint = fingerprint
}

// WithKeyID adds the keyID to the repo list keys params
func (o *RepoListKeysParams) WithKeyID(keyID *int64) *RepoListKeysParams {
	o.SetKeyID(keyID)
	return o
}

// SetKeyID adds the keyId to the repo list keys params
func (o *RepoListKeysParams) SetKeyID(keyID *int64) {
	o.KeyID = keyID
}

// WithLimit adds the limit to the repo list keys params
func (o *RepoListKeysParams) WithLimit(limit *int64) *RepoListKeysParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the repo list keys params
func (o *RepoListKeysParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOwner adds the owner to the repo list keys params
func (o *RepoListKeysParams) WithOwner(owner string) *RepoListKeysParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo list keys params
func (o *RepoListKeysParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithPage adds the page to the repo list keys params
func (o *RepoListKeysParams) WithPage(page *int64) *RepoListKeysParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the repo list keys params
func (o *RepoListKeysParams) SetPage(page *int64) {
	o.Page = page
}

// WithRepo adds the repo to the repo list keys params
func (o *RepoListKeysParams) WithRepo(repo string) *RepoListKeysParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo list keys params
func (o *RepoListKeysParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoListKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fingerprint != nil {

		// query param fingerprint
		var qrFingerprint string

		if o.Fingerprint != nil {
			qrFingerprint = *o.Fingerprint
		}
		qFingerprint := qrFingerprint
		if qFingerprint != "" {

			if err := r.SetQueryParam("fingerprint", qFingerprint); err != nil {
				return err
			}
		}
	}

	if o.KeyID != nil {

		// query param key_id
		var qrKeyID int64

		if o.KeyID != nil {
			qrKeyID = *o.KeyID
		}
		qKeyID := swag.FormatInt64(qrKeyID)
		if qKeyID != "" {

			if err := r.SetQueryParam("key_id", qKeyID); err != nil {
				return err
			}
		}
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
