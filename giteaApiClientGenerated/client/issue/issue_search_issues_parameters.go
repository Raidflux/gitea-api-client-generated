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

// NewIssueSearchIssuesParams creates a new IssueSearchIssuesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIssueSearchIssuesParams() *IssueSearchIssuesParams {
	return &IssueSearchIssuesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIssueSearchIssuesParamsWithTimeout creates a new IssueSearchIssuesParams object
// with the ability to set a timeout on a request.
func NewIssueSearchIssuesParamsWithTimeout(timeout time.Duration) *IssueSearchIssuesParams {
	return &IssueSearchIssuesParams{
		timeout: timeout,
	}
}

// NewIssueSearchIssuesParamsWithContext creates a new IssueSearchIssuesParams object
// with the ability to set a context for a request.
func NewIssueSearchIssuesParamsWithContext(ctx context.Context) *IssueSearchIssuesParams {
	return &IssueSearchIssuesParams{
		Context: ctx,
	}
}

// NewIssueSearchIssuesParamsWithHTTPClient creates a new IssueSearchIssuesParams object
// with the ability to set a custom HTTPClient for a request.
func NewIssueSearchIssuesParamsWithHTTPClient(client *http.Client) *IssueSearchIssuesParams {
	return &IssueSearchIssuesParams{
		HTTPClient: client,
	}
}

/*
IssueSearchIssuesParams contains all the parameters to send to the API endpoint

	for the issue search issues operation.

	Typically these are written to a http.Request.
*/
type IssueSearchIssuesParams struct {

	/* Assigned.

	   filter (issues / pulls) assigned to you, default is false
	*/
	Assigned *bool

	/* Before.

	   Only show notifications updated before the given time. This is a timestamp in RFC 3339 format

	   Format: date-time
	*/
	Before *strfmt.DateTime

	/* Created.

	   filter (issues / pulls) created by you, default is false
	*/
	Created *bool

	/* Labels.

	   comma separated list of labels. Fetch only issues that have any of this labels. Non existent labels are discarded
	*/
	Labels *string

	/* Limit.

	   page size of results
	*/
	Limit *int64

	/* Mentioned.

	   filter (issues / pulls) mentioning you, default is false
	*/
	Mentioned *bool

	/* Milestones.

	   comma separated list of milestone names. Fetch only issues that have any of this milestones. Non existent are discarded
	*/
	Milestones *string

	/* Owner.

	   filter by owner
	*/
	Owner *string

	/* Page.

	   page number of results to return (1-based)
	*/
	Page *int64

	/* PriorityRepoID.

	   repository to prioritize in the results

	   Format: int64
	*/
	PriorityRepoID *int64

	/* Q.

	   search string
	*/
	Q *string

	/* ReviewRequested.

	   filter pulls requesting your review, default is false
	*/
	ReviewRequested *bool

	/* Reviewed.

	   filter pulls reviewed by you, default is false
	*/
	Reviewed *bool

	/* Since.

	   Only show notifications updated after the given time. This is a timestamp in RFC 3339 format

	   Format: date-time
	*/
	Since *strfmt.DateTime

	/* State.

	   whether issue is open or closed
	*/
	State *string

	/* Team.

	   filter by team (requires organization owner parameter to be provided)
	*/
	Team *string

	/* Type.

	   filter by type (issues / pulls) if set
	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the issue search issues params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueSearchIssuesParams) WithDefaults() *IssueSearchIssuesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the issue search issues params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueSearchIssuesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the issue search issues params
func (o *IssueSearchIssuesParams) WithTimeout(timeout time.Duration) *IssueSearchIssuesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the issue search issues params
func (o *IssueSearchIssuesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the issue search issues params
func (o *IssueSearchIssuesParams) WithContext(ctx context.Context) *IssueSearchIssuesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the issue search issues params
func (o *IssueSearchIssuesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the issue search issues params
func (o *IssueSearchIssuesParams) WithHTTPClient(client *http.Client) *IssueSearchIssuesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the issue search issues params
func (o *IssueSearchIssuesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssigned adds the assigned to the issue search issues params
func (o *IssueSearchIssuesParams) WithAssigned(assigned *bool) *IssueSearchIssuesParams {
	o.SetAssigned(assigned)
	return o
}

// SetAssigned adds the assigned to the issue search issues params
func (o *IssueSearchIssuesParams) SetAssigned(assigned *bool) {
	o.Assigned = assigned
}

// WithBefore adds the before to the issue search issues params
func (o *IssueSearchIssuesParams) WithBefore(before *strfmt.DateTime) *IssueSearchIssuesParams {
	o.SetBefore(before)
	return o
}

// SetBefore adds the before to the issue search issues params
func (o *IssueSearchIssuesParams) SetBefore(before *strfmt.DateTime) {
	o.Before = before
}

// WithCreated adds the created to the issue search issues params
func (o *IssueSearchIssuesParams) WithCreated(created *bool) *IssueSearchIssuesParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the issue search issues params
func (o *IssueSearchIssuesParams) SetCreated(created *bool) {
	o.Created = created
}

// WithLabels adds the labels to the issue search issues params
func (o *IssueSearchIssuesParams) WithLabels(labels *string) *IssueSearchIssuesParams {
	o.SetLabels(labels)
	return o
}

// SetLabels adds the labels to the issue search issues params
func (o *IssueSearchIssuesParams) SetLabels(labels *string) {
	o.Labels = labels
}

// WithLimit adds the limit to the issue search issues params
func (o *IssueSearchIssuesParams) WithLimit(limit *int64) *IssueSearchIssuesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the issue search issues params
func (o *IssueSearchIssuesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMentioned adds the mentioned to the issue search issues params
func (o *IssueSearchIssuesParams) WithMentioned(mentioned *bool) *IssueSearchIssuesParams {
	o.SetMentioned(mentioned)
	return o
}

// SetMentioned adds the mentioned to the issue search issues params
func (o *IssueSearchIssuesParams) SetMentioned(mentioned *bool) {
	o.Mentioned = mentioned
}

// WithMilestones adds the milestones to the issue search issues params
func (o *IssueSearchIssuesParams) WithMilestones(milestones *string) *IssueSearchIssuesParams {
	o.SetMilestones(milestones)
	return o
}

// SetMilestones adds the milestones to the issue search issues params
func (o *IssueSearchIssuesParams) SetMilestones(milestones *string) {
	o.Milestones = milestones
}

// WithOwner adds the owner to the issue search issues params
func (o *IssueSearchIssuesParams) WithOwner(owner *string) *IssueSearchIssuesParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the issue search issues params
func (o *IssueSearchIssuesParams) SetOwner(owner *string) {
	o.Owner = owner
}

// WithPage adds the page to the issue search issues params
func (o *IssueSearchIssuesParams) WithPage(page *int64) *IssueSearchIssuesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the issue search issues params
func (o *IssueSearchIssuesParams) SetPage(page *int64) {
	o.Page = page
}

// WithPriorityRepoID adds the priorityRepoID to the issue search issues params
func (o *IssueSearchIssuesParams) WithPriorityRepoID(priorityRepoID *int64) *IssueSearchIssuesParams {
	o.SetPriorityRepoID(priorityRepoID)
	return o
}

// SetPriorityRepoID adds the priorityRepoId to the issue search issues params
func (o *IssueSearchIssuesParams) SetPriorityRepoID(priorityRepoID *int64) {
	o.PriorityRepoID = priorityRepoID
}

// WithQ adds the q to the issue search issues params
func (o *IssueSearchIssuesParams) WithQ(q *string) *IssueSearchIssuesParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the issue search issues params
func (o *IssueSearchIssuesParams) SetQ(q *string) {
	o.Q = q
}

// WithReviewRequested adds the reviewRequested to the issue search issues params
func (o *IssueSearchIssuesParams) WithReviewRequested(reviewRequested *bool) *IssueSearchIssuesParams {
	o.SetReviewRequested(reviewRequested)
	return o
}

// SetReviewRequested adds the reviewRequested to the issue search issues params
func (o *IssueSearchIssuesParams) SetReviewRequested(reviewRequested *bool) {
	o.ReviewRequested = reviewRequested
}

// WithReviewed adds the reviewed to the issue search issues params
func (o *IssueSearchIssuesParams) WithReviewed(reviewed *bool) *IssueSearchIssuesParams {
	o.SetReviewed(reviewed)
	return o
}

// SetReviewed adds the reviewed to the issue search issues params
func (o *IssueSearchIssuesParams) SetReviewed(reviewed *bool) {
	o.Reviewed = reviewed
}

// WithSince adds the since to the issue search issues params
func (o *IssueSearchIssuesParams) WithSince(since *strfmt.DateTime) *IssueSearchIssuesParams {
	o.SetSince(since)
	return o
}

// SetSince adds the since to the issue search issues params
func (o *IssueSearchIssuesParams) SetSince(since *strfmt.DateTime) {
	o.Since = since
}

// WithState adds the state to the issue search issues params
func (o *IssueSearchIssuesParams) WithState(state *string) *IssueSearchIssuesParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the issue search issues params
func (o *IssueSearchIssuesParams) SetState(state *string) {
	o.State = state
}

// WithTeam adds the team to the issue search issues params
func (o *IssueSearchIssuesParams) WithTeam(team *string) *IssueSearchIssuesParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the issue search issues params
func (o *IssueSearchIssuesParams) SetTeam(team *string) {
	o.Team = team
}

// WithType adds the typeVar to the issue search issues params
func (o *IssueSearchIssuesParams) WithType(typeVar *string) *IssueSearchIssuesParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the issue search issues params
func (o *IssueSearchIssuesParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *IssueSearchIssuesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Assigned != nil {

		// query param assigned
		var qrAssigned bool

		if o.Assigned != nil {
			qrAssigned = *o.Assigned
		}
		qAssigned := swag.FormatBool(qrAssigned)
		if qAssigned != "" {

			if err := r.SetQueryParam("assigned", qAssigned); err != nil {
				return err
			}
		}
	}

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

	if o.Created != nil {

		// query param created
		var qrCreated bool

		if o.Created != nil {
			qrCreated = *o.Created
		}
		qCreated := swag.FormatBool(qrCreated)
		if qCreated != "" {

			if err := r.SetQueryParam("created", qCreated); err != nil {
				return err
			}
		}
	}

	if o.Labels != nil {

		// query param labels
		var qrLabels string

		if o.Labels != nil {
			qrLabels = *o.Labels
		}
		qLabels := qrLabels
		if qLabels != "" {

			if err := r.SetQueryParam("labels", qLabels); err != nil {
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

	if o.Mentioned != nil {

		// query param mentioned
		var qrMentioned bool

		if o.Mentioned != nil {
			qrMentioned = *o.Mentioned
		}
		qMentioned := swag.FormatBool(qrMentioned)
		if qMentioned != "" {

			if err := r.SetQueryParam("mentioned", qMentioned); err != nil {
				return err
			}
		}
	}

	if o.Milestones != nil {

		// query param milestones
		var qrMilestones string

		if o.Milestones != nil {
			qrMilestones = *o.Milestones
		}
		qMilestones := qrMilestones
		if qMilestones != "" {

			if err := r.SetQueryParam("milestones", qMilestones); err != nil {
				return err
			}
		}
	}

	if o.Owner != nil {

		// query param owner
		var qrOwner string

		if o.Owner != nil {
			qrOwner = *o.Owner
		}
		qOwner := qrOwner
		if qOwner != "" {

			if err := r.SetQueryParam("owner", qOwner); err != nil {
				return err
			}
		}
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

	if o.PriorityRepoID != nil {

		// query param priority_repo_id
		var qrPriorityRepoID int64

		if o.PriorityRepoID != nil {
			qrPriorityRepoID = *o.PriorityRepoID
		}
		qPriorityRepoID := swag.FormatInt64(qrPriorityRepoID)
		if qPriorityRepoID != "" {

			if err := r.SetQueryParam("priority_repo_id", qPriorityRepoID); err != nil {
				return err
			}
		}
	}

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}
	}

	if o.ReviewRequested != nil {

		// query param review_requested
		var qrReviewRequested bool

		if o.ReviewRequested != nil {
			qrReviewRequested = *o.ReviewRequested
		}
		qReviewRequested := swag.FormatBool(qrReviewRequested)
		if qReviewRequested != "" {

			if err := r.SetQueryParam("review_requested", qReviewRequested); err != nil {
				return err
			}
		}
	}

	if o.Reviewed != nil {

		// query param reviewed
		var qrReviewed bool

		if o.Reviewed != nil {
			qrReviewed = *o.Reviewed
		}
		qReviewed := swag.FormatBool(qrReviewed)
		if qReviewed != "" {

			if err := r.SetQueryParam("reviewed", qReviewed); err != nil {
				return err
			}
		}
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

	if o.Team != nil {

		// query param team
		var qrTeam string

		if o.Team != nil {
			qrTeam = *o.Team
		}
		qTeam := qrTeam
		if qTeam != "" {

			if err := r.SetQueryParam("team", qTeam); err != nil {
				return err
			}
		}
	}

	if o.Type != nil {

		// query param type
		var qrType string

		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
