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

// NewRepoTrackedTimesParams creates a new RepoTrackedTimesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepoTrackedTimesParams() *RepoTrackedTimesParams {
	return &RepoTrackedTimesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepoTrackedTimesParamsWithTimeout creates a new RepoTrackedTimesParams object
// with the ability to set a timeout on a request.
func NewRepoTrackedTimesParamsWithTimeout(timeout time.Duration) *RepoTrackedTimesParams {
	return &RepoTrackedTimesParams{
		timeout: timeout,
	}
}

// NewRepoTrackedTimesParamsWithContext creates a new RepoTrackedTimesParams object
// with the ability to set a context for a request.
func NewRepoTrackedTimesParamsWithContext(ctx context.Context) *RepoTrackedTimesParams {
	return &RepoTrackedTimesParams{
		Context: ctx,
	}
}

// NewRepoTrackedTimesParamsWithHTTPClient creates a new RepoTrackedTimesParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepoTrackedTimesParamsWithHTTPClient(client *http.Client) *RepoTrackedTimesParams {
	return &RepoTrackedTimesParams{
		HTTPClient: client,
	}
}

/*
RepoTrackedTimesParams contains all the parameters to send to the API endpoint

	for the repo tracked times operation.

	Typically these are written to a http.Request.
*/
type RepoTrackedTimesParams struct {

	/* Before.

	   Only show times updated before the given time. This is a timestamp in RFC 3339 format

	   Format: date-time
	*/
	Before *strfmt.DateTime

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

	/* Since.

	   Only show times updated after the given time. This is a timestamp in RFC 3339 format

	   Format: date-time
	*/
	Since *strfmt.DateTime

	/* User.

	   optional filter by user (available for issue managers)
	*/
	User *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the repo tracked times params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoTrackedTimesParams) WithDefaults() *RepoTrackedTimesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repo tracked times params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoTrackedTimesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repo tracked times params
func (o *RepoTrackedTimesParams) WithTimeout(timeout time.Duration) *RepoTrackedTimesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo tracked times params
func (o *RepoTrackedTimesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo tracked times params
func (o *RepoTrackedTimesParams) WithContext(ctx context.Context) *RepoTrackedTimesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo tracked times params
func (o *RepoTrackedTimesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo tracked times params
func (o *RepoTrackedTimesParams) WithHTTPClient(client *http.Client) *RepoTrackedTimesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo tracked times params
func (o *RepoTrackedTimesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBefore adds the before to the repo tracked times params
func (o *RepoTrackedTimesParams) WithBefore(before *strfmt.DateTime) *RepoTrackedTimesParams {
	o.SetBefore(before)
	return o
}

// SetBefore adds the before to the repo tracked times params
func (o *RepoTrackedTimesParams) SetBefore(before *strfmt.DateTime) {
	o.Before = before
}

// WithLimit adds the limit to the repo tracked times params
func (o *RepoTrackedTimesParams) WithLimit(limit *int64) *RepoTrackedTimesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the repo tracked times params
func (o *RepoTrackedTimesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOwner adds the owner to the repo tracked times params
func (o *RepoTrackedTimesParams) WithOwner(owner string) *RepoTrackedTimesParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo tracked times params
func (o *RepoTrackedTimesParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithPage adds the page to the repo tracked times params
func (o *RepoTrackedTimesParams) WithPage(page *int64) *RepoTrackedTimesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the repo tracked times params
func (o *RepoTrackedTimesParams) SetPage(page *int64) {
	o.Page = page
}

// WithRepo adds the repo to the repo tracked times params
func (o *RepoTrackedTimesParams) WithRepo(repo string) *RepoTrackedTimesParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo tracked times params
func (o *RepoTrackedTimesParams) SetRepo(repo string) {
	o.Repo = repo
}

// WithSince adds the since to the repo tracked times params
func (o *RepoTrackedTimesParams) WithSince(since *strfmt.DateTime) *RepoTrackedTimesParams {
	o.SetSince(since)
	return o
}

// SetSince adds the since to the repo tracked times params
func (o *RepoTrackedTimesParams) SetSince(since *strfmt.DateTime) {
	o.Since = since
}

// WithUser adds the user to the repo tracked times params
func (o *RepoTrackedTimesParams) WithUser(user *string) *RepoTrackedTimesParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the repo tracked times params
func (o *RepoTrackedTimesParams) SetUser(user *string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *RepoTrackedTimesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Before != nil {

		// query param before
		var qrBefore strfmt.DateTime

		if o.Before != nil {
			qrBefore = *o.Before
		}
		qBefore := qrBefore.String()
		if qBefore != "" {

			if err := r.SetQueryParam("before", qBefore); err != nil {
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

	if o.Since != nil {

		// query param since
		var qrSince strfmt.DateTime

		if o.Since != nil {
			qrSince = *o.Since
		}
		qSince := qrSince.String()
		if qSince != "" {

			if err := r.SetQueryParam("since", qSince); err != nil {
				return err
			}
		}
	}

	if o.User != nil {

		// query param user
		var qrUser string

		if o.User != nil {
			qrUser = *o.User
		}
		qUser := qrUser
		if qUser != "" {

			if err := r.SetQueryParam("user", qUser); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
