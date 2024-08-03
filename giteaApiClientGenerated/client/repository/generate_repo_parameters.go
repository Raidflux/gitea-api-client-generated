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

	"giteaApiClientGenerated/models"
)

// NewGenerateRepoParams creates a new GenerateRepoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGenerateRepoParams() *GenerateRepoParams {
	return &GenerateRepoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGenerateRepoParamsWithTimeout creates a new GenerateRepoParams object
// with the ability to set a timeout on a request.
func NewGenerateRepoParamsWithTimeout(timeout time.Duration) *GenerateRepoParams {
	return &GenerateRepoParams{
		timeout: timeout,
	}
}

// NewGenerateRepoParamsWithContext creates a new GenerateRepoParams object
// with the ability to set a context for a request.
func NewGenerateRepoParamsWithContext(ctx context.Context) *GenerateRepoParams {
	return &GenerateRepoParams{
		Context: ctx,
	}
}

// NewGenerateRepoParamsWithHTTPClient creates a new GenerateRepoParams object
// with the ability to set a custom HTTPClient for a request.
func NewGenerateRepoParamsWithHTTPClient(client *http.Client) *GenerateRepoParams {
	return &GenerateRepoParams{
		HTTPClient: client,
	}
}

/*
GenerateRepoParams contains all the parameters to send to the API endpoint

	for the generate repo operation.

	Typically these are written to a http.Request.
*/
type GenerateRepoParams struct {

	// Body.
	Body *models.GenerateRepoOption

	/* TemplateOwner.

	   name of the template repository owner
	*/
	TemplateOwner string

	/* TemplateRepo.

	   name of the template repository
	*/
	TemplateRepo string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the generate repo params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GenerateRepoParams) WithDefaults() *GenerateRepoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the generate repo params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GenerateRepoParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the generate repo params
func (o *GenerateRepoParams) WithTimeout(timeout time.Duration) *GenerateRepoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the generate repo params
func (o *GenerateRepoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the generate repo params
func (o *GenerateRepoParams) WithContext(ctx context.Context) *GenerateRepoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the generate repo params
func (o *GenerateRepoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the generate repo params
func (o *GenerateRepoParams) WithHTTPClient(client *http.Client) *GenerateRepoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the generate repo params
func (o *GenerateRepoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the generate repo params
func (o *GenerateRepoParams) WithBody(body *models.GenerateRepoOption) *GenerateRepoParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the generate repo params
func (o *GenerateRepoParams) SetBody(body *models.GenerateRepoOption) {
	o.Body = body
}

// WithTemplateOwner adds the templateOwner to the generate repo params
func (o *GenerateRepoParams) WithTemplateOwner(templateOwner string) *GenerateRepoParams {
	o.SetTemplateOwner(templateOwner)
	return o
}

// SetTemplateOwner adds the templateOwner to the generate repo params
func (o *GenerateRepoParams) SetTemplateOwner(templateOwner string) {
	o.TemplateOwner = templateOwner
}

// WithTemplateRepo adds the templateRepo to the generate repo params
func (o *GenerateRepoParams) WithTemplateRepo(templateRepo string) *GenerateRepoParams {
	o.SetTemplateRepo(templateRepo)
	return o
}

// SetTemplateRepo adds the templateRepo to the generate repo params
func (o *GenerateRepoParams) SetTemplateRepo(templateRepo string) {
	o.TemplateRepo = templateRepo
}

// WriteToRequest writes these params to a swagger request
func (o *GenerateRepoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param template_owner
	if err := r.SetPathParam("template_owner", o.TemplateOwner); err != nil {
		return err
	}

	// path param template_repo
	if err := r.SetPathParam("template_repo", o.TemplateRepo); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
