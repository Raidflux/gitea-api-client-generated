// Code generated by go-swagger; DO NOT EDIT.

package miscellaneous

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

// NewGetGitignoreTemplateInfoParams creates a new GetGitignoreTemplateInfoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetGitignoreTemplateInfoParams() *GetGitignoreTemplateInfoParams {
	return &GetGitignoreTemplateInfoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetGitignoreTemplateInfoParamsWithTimeout creates a new GetGitignoreTemplateInfoParams object
// with the ability to set a timeout on a request.
func NewGetGitignoreTemplateInfoParamsWithTimeout(timeout time.Duration) *GetGitignoreTemplateInfoParams {
	return &GetGitignoreTemplateInfoParams{
		timeout: timeout,
	}
}

// NewGetGitignoreTemplateInfoParamsWithContext creates a new GetGitignoreTemplateInfoParams object
// with the ability to set a context for a request.
func NewGetGitignoreTemplateInfoParamsWithContext(ctx context.Context) *GetGitignoreTemplateInfoParams {
	return &GetGitignoreTemplateInfoParams{
		Context: ctx,
	}
}

// NewGetGitignoreTemplateInfoParamsWithHTTPClient creates a new GetGitignoreTemplateInfoParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetGitignoreTemplateInfoParamsWithHTTPClient(client *http.Client) *GetGitignoreTemplateInfoParams {
	return &GetGitignoreTemplateInfoParams{
		HTTPClient: client,
	}
}

/*
GetGitignoreTemplateInfoParams contains all the parameters to send to the API endpoint

	for the get gitignore template info operation.

	Typically these are written to a http.Request.
*/
type GetGitignoreTemplateInfoParams struct {

	/* Name.

	   name of the template
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get gitignore template info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGitignoreTemplateInfoParams) WithDefaults() *GetGitignoreTemplateInfoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get gitignore template info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGitignoreTemplateInfoParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get gitignore template info params
func (o *GetGitignoreTemplateInfoParams) WithTimeout(timeout time.Duration) *GetGitignoreTemplateInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get gitignore template info params
func (o *GetGitignoreTemplateInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get gitignore template info params
func (o *GetGitignoreTemplateInfoParams) WithContext(ctx context.Context) *GetGitignoreTemplateInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get gitignore template info params
func (o *GetGitignoreTemplateInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get gitignore template info params
func (o *GetGitignoreTemplateInfoParams) WithHTTPClient(client *http.Client) *GetGitignoreTemplateInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get gitignore template info params
func (o *GetGitignoreTemplateInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get gitignore template info params
func (o *GetGitignoreTemplateInfoParams) WithName(name string) *GetGitignoreTemplateInfoParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get gitignore template info params
func (o *GetGitignoreTemplateInfoParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetGitignoreTemplateInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
