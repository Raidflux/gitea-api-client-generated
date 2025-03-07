// Code generated by go-swagger; DO NOT EDIT.

package notification

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

// NewNotifyReadThreadParams creates a new NotifyReadThreadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNotifyReadThreadParams() *NotifyReadThreadParams {
	return &NotifyReadThreadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNotifyReadThreadParamsWithTimeout creates a new NotifyReadThreadParams object
// with the ability to set a timeout on a request.
func NewNotifyReadThreadParamsWithTimeout(timeout time.Duration) *NotifyReadThreadParams {
	return &NotifyReadThreadParams{
		timeout: timeout,
	}
}

// NewNotifyReadThreadParamsWithContext creates a new NotifyReadThreadParams object
// with the ability to set a context for a request.
func NewNotifyReadThreadParamsWithContext(ctx context.Context) *NotifyReadThreadParams {
	return &NotifyReadThreadParams{
		Context: ctx,
	}
}

// NewNotifyReadThreadParamsWithHTTPClient creates a new NotifyReadThreadParams object
// with the ability to set a custom HTTPClient for a request.
func NewNotifyReadThreadParamsWithHTTPClient(client *http.Client) *NotifyReadThreadParams {
	return &NotifyReadThreadParams{
		HTTPClient: client,
	}
}

/*
NotifyReadThreadParams contains all the parameters to send to the API endpoint

	for the notify read thread operation.

	Typically these are written to a http.Request.
*/
type NotifyReadThreadParams struct {

	/* ID.

	   id of notification thread
	*/
	ID string

	/* ToStatus.

	   Status to mark notifications as

	   Default: "read"
	*/
	ToStatus *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the notify read thread params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NotifyReadThreadParams) WithDefaults() *NotifyReadThreadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the notify read thread params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NotifyReadThreadParams) SetDefaults() {
	var (
		toStatusDefault = string("read")
	)

	val := NotifyReadThreadParams{
		ToStatus: &toStatusDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the notify read thread params
func (o *NotifyReadThreadParams) WithTimeout(timeout time.Duration) *NotifyReadThreadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the notify read thread params
func (o *NotifyReadThreadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the notify read thread params
func (o *NotifyReadThreadParams) WithContext(ctx context.Context) *NotifyReadThreadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the notify read thread params
func (o *NotifyReadThreadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the notify read thread params
func (o *NotifyReadThreadParams) WithHTTPClient(client *http.Client) *NotifyReadThreadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the notify read thread params
func (o *NotifyReadThreadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the notify read thread params
func (o *NotifyReadThreadParams) WithID(id string) *NotifyReadThreadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the notify read thread params
func (o *NotifyReadThreadParams) SetID(id string) {
	o.ID = id
}

// WithToStatus adds the toStatus to the notify read thread params
func (o *NotifyReadThreadParams) WithToStatus(toStatus *string) *NotifyReadThreadParams {
	o.SetToStatus(toStatus)
	return o
}

// SetToStatus adds the toStatus to the notify read thread params
func (o *NotifyReadThreadParams) SetToStatus(toStatus *string) {
	o.ToStatus = toStatus
}

// WriteToRequest writes these params to a swagger request
func (o *NotifyReadThreadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.ToStatus != nil {

		// query param to-status
		var qrToStatus string

		if o.ToStatus != nil {
			qrToStatus = *o.ToStatus
		}
		qToStatus := qrToStatus
		if qToStatus != "" {

			if err := r.SetQueryParam("to-status", qToStatus); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
