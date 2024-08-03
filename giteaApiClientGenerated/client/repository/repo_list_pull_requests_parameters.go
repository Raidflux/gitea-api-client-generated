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

// NewRepoListPullRequestsParams creates a new RepoListPullRequestsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepoListPullRequestsParams() *RepoListPullRequestsParams {
	return &RepoListPullRequestsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepoListPullRequestsParamsWithTimeout creates a new RepoListPullRequestsParams object
// with the ability to set a timeout on a request.
func NewRepoListPullRequestsParamsWithTimeout(timeout time.Duration) *RepoListPullRequestsParams {
	return &RepoListPullRequestsParams{
		timeout: timeout,
	}
}

// NewRepoListPullRequestsParamsWithContext creates a new RepoListPullRequestsParams object
// with the ability to set a context for a request.
func NewRepoListPullRequestsParamsWithContext(ctx context.Context) *RepoListPullRequestsParams {
	return &RepoListPullRequestsParams{
		Context: ctx,
	}
}

// NewRepoListPullRequestsParamsWithHTTPClient creates a new RepoListPullRequestsParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepoListPullRequestsParamsWithHTTPClient(client *http.Client) *RepoListPullRequestsParams {
	return &RepoListPullRequestsParams{
		HTTPClient: client,
	}
}

/*
RepoListPullRequestsParams contains all the parameters to send to the API endpoint

	for the repo list pull requests operation.

	Typically these are written to a http.Request.
*/
type RepoListPullRequestsParams struct {

	/* Labels.

	   Label IDs
	*/
	Labels []int64

	/* Limit.

	   page size of results
	*/
	Limit *int64

	/* Milestone.

	   ID of the milestone

	   Format: int64
	*/
	Milestone *int64

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

	/* Sort.

	   Type of sort
	*/
	Sort *string

	/* State.

	   State of pull request: open or closed (optional)
	*/
	State *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the repo list pull requests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoListPullRequestsParams) WithDefaults() *RepoListPullRequestsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repo list pull requests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoListPullRequestsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repo list pull requests params
func (o *RepoListPullRequestsParams) WithTimeout(timeout time.Duration) *RepoListPullRequestsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo list pull requests params
func (o *RepoListPullRequestsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo list pull requests params
func (o *RepoListPullRequestsParams) WithContext(ctx context.Context) *RepoListPullRequestsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo list pull requests params
func (o *RepoListPullRequestsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo list pull requests params
func (o *RepoListPullRequestsParams) WithHTTPClient(client *http.Client) *RepoListPullRequestsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo list pull requests params
func (o *RepoListPullRequestsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLabels adds the labels to the repo list pull requests params
func (o *RepoListPullRequestsParams) WithLabels(labels []int64) *RepoListPullRequestsParams {
	o.SetLabels(labels)
	return o
}

// SetLabels adds the labels to the repo list pull requests params
func (o *RepoListPullRequestsParams) SetLabels(labels []int64) {
	o.Labels = labels
}

// WithLimit adds the limit to the repo list pull requests params
func (o *RepoListPullRequestsParams) WithLimit(limit *int64) *RepoListPullRequestsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the repo list pull requests params
func (o *RepoListPullRequestsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMilestone adds the milestone to the repo list pull requests params
func (o *RepoListPullRequestsParams) WithMilestone(milestone *int64) *RepoListPullRequestsParams {
	o.SetMilestone(milestone)
	return o
}

// SetMilestone adds the milestone to the repo list pull requests params
func (o *RepoListPullRequestsParams) SetMilestone(milestone *int64) {
	o.Milestone = milestone
}

// WithOwner adds the owner to the repo list pull requests params
func (o *RepoListPullRequestsParams) WithOwner(owner string) *RepoListPullRequestsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo list pull requests params
func (o *RepoListPullRequestsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithPage adds the page to the repo list pull requests params
func (o *RepoListPullRequestsParams) WithPage(page *int64) *RepoListPullRequestsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the repo list pull requests params
func (o *RepoListPullRequestsParams) SetPage(page *int64) {
	o.Page = page
}

// WithRepo adds the repo to the repo list pull requests params
func (o *RepoListPullRequestsParams) WithRepo(repo string) *RepoListPullRequestsParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo list pull requests params
func (o *RepoListPullRequestsParams) SetRepo(repo string) {
	o.Repo = repo
}

// WithSort adds the sort to the repo list pull requests params
func (o *RepoListPullRequestsParams) WithSort(sort *string) *RepoListPullRequestsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the repo list pull requests params
func (o *RepoListPullRequestsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithState adds the state to the repo list pull requests params
func (o *RepoListPullRequestsParams) WithState(state *string) *RepoListPullRequestsParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the repo list pull requests params
func (o *RepoListPullRequestsParams) SetState(state *string) {
	o.State = state
}

// WriteToRequest writes these params to a swagger request
func (o *RepoListPullRequestsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Labels != nil {

		// binding items for labels
		joinedLabels := o.bindParamLabels(reg)

		// query array param labels
		if err := r.SetQueryParam("labels", joinedLabels...); err != nil {
			return err
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

	if o.Milestone != nil {

		// query param milestone
		var qrMilestone int64

		if o.Milestone != nil {
			qrMilestone = *o.Milestone
		}
		qMilestone := swag.FormatInt64(qrMilestone)
		if qMilestone != "" {

			if err := r.SetQueryParam("milestone", qMilestone); err != nil {
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

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if o.State != nil {

		// query param state
		var qrState string

		if o.State != nil {
			qrState = *o.State
		}
		qState := qrState
		if qState != "" {

			if err := r.SetQueryParam("state", qState); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamRepoListPullRequests binds the parameter labels
func (o *RepoListPullRequestsParams) bindParamLabels(formats strfmt.Registry) []string {
	labelsIR := o.Labels

	var labelsIC []string
	for _, labelsIIR := range labelsIR { // explode []int64

		labelsIIV := swag.FormatInt64(labelsIIR) // int64 as string
		labelsIC = append(labelsIC, labelsIIV)
	}

	// items.CollectionFormat: "multi"
	labelsIS := swag.JoinByFormat(labelsIC, "multi")

	return labelsIS
}
