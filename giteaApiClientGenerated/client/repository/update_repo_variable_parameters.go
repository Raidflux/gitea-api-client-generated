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

	"github.com/Raidflux/gitea-api-client-generated/giteaApiClientGenerated/models"
)

// NewUpdateRepoVariableParams creates a new UpdateRepoVariableParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateRepoVariableParams() *UpdateRepoVariableParams {
	return &UpdateRepoVariableParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRepoVariableParamsWithTimeout creates a new UpdateRepoVariableParams object
// with the ability to set a timeout on a request.
func NewUpdateRepoVariableParamsWithTimeout(timeout time.Duration) *UpdateRepoVariableParams {
	return &UpdateRepoVariableParams{
		timeout: timeout,
	}
}

// NewUpdateRepoVariableParamsWithContext creates a new UpdateRepoVariableParams object
// with the ability to set a context for a request.
func NewUpdateRepoVariableParamsWithContext(ctx context.Context) *UpdateRepoVariableParams {
	return &UpdateRepoVariableParams{
		Context: ctx,
	}
}

// NewUpdateRepoVariableParamsWithHTTPClient creates a new UpdateRepoVariableParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateRepoVariableParamsWithHTTPClient(client *http.Client) *UpdateRepoVariableParams {
	return &UpdateRepoVariableParams{
		HTTPClient: client,
	}
}

/*
UpdateRepoVariableParams contains all the parameters to send to the API endpoint

	for the update repo variable operation.

	Typically these are written to a http.Request.
*/
type UpdateRepoVariableParams struct {

	// Body.
	Body *models.UpdateVariableOption

	/* Owner.

	   name of the owner
	*/
	Owner string

	/* Repo.

	   name of the repository
	*/
	Repo string

	/* Variablename.

	   name of the variable
	*/
	Variablename string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update repo variable params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRepoVariableParams) WithDefaults() *UpdateRepoVariableParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update repo variable params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRepoVariableParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update repo variable params
func (o *UpdateRepoVariableParams) WithTimeout(timeout time.Duration) *UpdateRepoVariableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update repo variable params
func (o *UpdateRepoVariableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update repo variable params
func (o *UpdateRepoVariableParams) WithContext(ctx context.Context) *UpdateRepoVariableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update repo variable params
func (o *UpdateRepoVariableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update repo variable params
func (o *UpdateRepoVariableParams) WithHTTPClient(client *http.Client) *UpdateRepoVariableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update repo variable params
func (o *UpdateRepoVariableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update repo variable params
func (o *UpdateRepoVariableParams) WithBody(body *models.UpdateVariableOption) *UpdateRepoVariableParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update repo variable params
func (o *UpdateRepoVariableParams) SetBody(body *models.UpdateVariableOption) {
	o.Body = body
}

// WithOwner adds the owner to the update repo variable params
func (o *UpdateRepoVariableParams) WithOwner(owner string) *UpdateRepoVariableParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the update repo variable params
func (o *UpdateRepoVariableParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the update repo variable params
func (o *UpdateRepoVariableParams) WithRepo(repo string) *UpdateRepoVariableParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the update repo variable params
func (o *UpdateRepoVariableParams) SetRepo(repo string) {
	o.Repo = repo
}

// WithVariablename adds the variablename to the update repo variable params
func (o *UpdateRepoVariableParams) WithVariablename(variablename string) *UpdateRepoVariableParams {
	o.SetVariablename(variablename)
	return o
}

// SetVariablename adds the variablename to the update repo variable params
func (o *UpdateRepoVariableParams) SetVariablename(variablename string) {
	o.Variablename = variablename
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRepoVariableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param repo
	if err := r.SetPathParam("repo", o.Repo); err != nil {
		return err
	}

	// path param variablename
	if err := r.SetPathParam("variablename", o.Variablename); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
