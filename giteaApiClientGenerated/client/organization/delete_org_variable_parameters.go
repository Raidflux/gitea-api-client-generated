// Code generated by go-swagger; DO NOT EDIT.

package organization

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

// NewDeleteOrgVariableParams creates a new DeleteOrgVariableParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteOrgVariableParams() *DeleteOrgVariableParams {
	return &DeleteOrgVariableParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOrgVariableParamsWithTimeout creates a new DeleteOrgVariableParams object
// with the ability to set a timeout on a request.
func NewDeleteOrgVariableParamsWithTimeout(timeout time.Duration) *DeleteOrgVariableParams {
	return &DeleteOrgVariableParams{
		timeout: timeout,
	}
}

// NewDeleteOrgVariableParamsWithContext creates a new DeleteOrgVariableParams object
// with the ability to set a context for a request.
func NewDeleteOrgVariableParamsWithContext(ctx context.Context) *DeleteOrgVariableParams {
	return &DeleteOrgVariableParams{
		Context: ctx,
	}
}

// NewDeleteOrgVariableParamsWithHTTPClient creates a new DeleteOrgVariableParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteOrgVariableParamsWithHTTPClient(client *http.Client) *DeleteOrgVariableParams {
	return &DeleteOrgVariableParams{
		HTTPClient: client,
	}
}

/*
DeleteOrgVariableParams contains all the parameters to send to the API endpoint

	for the delete org variable operation.

	Typically these are written to a http.Request.
*/
type DeleteOrgVariableParams struct {

	/* Org.

	   name of the organization
	*/
	Org string

	/* Variablename.

	   name of the variable
	*/
	Variablename string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete org variable params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrgVariableParams) WithDefaults() *DeleteOrgVariableParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete org variable params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrgVariableParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete org variable params
func (o *DeleteOrgVariableParams) WithTimeout(timeout time.Duration) *DeleteOrgVariableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete org variable params
func (o *DeleteOrgVariableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete org variable params
func (o *DeleteOrgVariableParams) WithContext(ctx context.Context) *DeleteOrgVariableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete org variable params
func (o *DeleteOrgVariableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete org variable params
func (o *DeleteOrgVariableParams) WithHTTPClient(client *http.Client) *DeleteOrgVariableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete org variable params
func (o *DeleteOrgVariableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrg adds the org to the delete org variable params
func (o *DeleteOrgVariableParams) WithOrg(org string) *DeleteOrgVariableParams {
	o.SetOrg(org)
	return o
}

// SetOrg adds the org to the delete org variable params
func (o *DeleteOrgVariableParams) SetOrg(org string) {
	o.Org = org
}

// WithVariablename adds the variablename to the delete org variable params
func (o *DeleteOrgVariableParams) WithVariablename(variablename string) *DeleteOrgVariableParams {
	o.SetVariablename(variablename)
	return o
}

// SetVariablename adds the variablename to the delete org variable params
func (o *DeleteOrgVariableParams) SetVariablename(variablename string) {
	o.Variablename = variablename
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOrgVariableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param org
	if err := r.SetPathParam("org", o.Org); err != nil {
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
