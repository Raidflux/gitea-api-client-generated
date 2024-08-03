// Code generated by go-swagger; DO NOT EDIT.

package admin

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

// NewAdminCronRunParams creates a new AdminCronRunParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAdminCronRunParams() *AdminCronRunParams {
	return &AdminCronRunParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAdminCronRunParamsWithTimeout creates a new AdminCronRunParams object
// with the ability to set a timeout on a request.
func NewAdminCronRunParamsWithTimeout(timeout time.Duration) *AdminCronRunParams {
	return &AdminCronRunParams{
		timeout: timeout,
	}
}

// NewAdminCronRunParamsWithContext creates a new AdminCronRunParams object
// with the ability to set a context for a request.
func NewAdminCronRunParamsWithContext(ctx context.Context) *AdminCronRunParams {
	return &AdminCronRunParams{
		Context: ctx,
	}
}

// NewAdminCronRunParamsWithHTTPClient creates a new AdminCronRunParams object
// with the ability to set a custom HTTPClient for a request.
func NewAdminCronRunParamsWithHTTPClient(client *http.Client) *AdminCronRunParams {
	return &AdminCronRunParams{
		HTTPClient: client,
	}
}

/*
AdminCronRunParams contains all the parameters to send to the API endpoint

	for the admin cron run operation.

	Typically these are written to a http.Request.
*/
type AdminCronRunParams struct {

	/* Task.

	   task to run
	*/
	Task string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the admin cron run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminCronRunParams) WithDefaults() *AdminCronRunParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the admin cron run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminCronRunParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the admin cron run params
func (o *AdminCronRunParams) WithTimeout(timeout time.Duration) *AdminCronRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin cron run params
func (o *AdminCronRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin cron run params
func (o *AdminCronRunParams) WithContext(ctx context.Context) *AdminCronRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin cron run params
func (o *AdminCronRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin cron run params
func (o *AdminCronRunParams) WithHTTPClient(client *http.Client) *AdminCronRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin cron run params
func (o *AdminCronRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTask adds the task to the admin cron run params
func (o *AdminCronRunParams) WithTask(task string) *AdminCronRunParams {
	o.SetTask(task)
	return o
}

// SetTask adds the task to the admin cron run params
func (o *AdminCronRunParams) SetTask(task string) {
	o.Task = task
}

// WriteToRequest writes these params to a swagger request
func (o *AdminCronRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param task
	if err := r.SetPathParam("task", o.Task); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
