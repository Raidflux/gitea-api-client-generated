// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateOrgSecretReader is a Reader for the UpdateOrgSecret structure.
type UpdateOrgSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOrgSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateOrgSecretCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewUpdateOrgSecretNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateOrgSecretBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateOrgSecretNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /orgs/{org}/actions/secrets/{secretname}] updateOrgSecret", response, response.Code())
	}
}

// NewUpdateOrgSecretCreated creates a UpdateOrgSecretCreated with default headers values
func NewUpdateOrgSecretCreated() *UpdateOrgSecretCreated {
	return &UpdateOrgSecretCreated{}
}

/*
UpdateOrgSecretCreated describes a response with status code 201, with default header values.

response when creating a secret
*/
type UpdateOrgSecretCreated struct {
}

// IsSuccess returns true when this update org secret created response has a 2xx status code
func (o *UpdateOrgSecretCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update org secret created response has a 3xx status code
func (o *UpdateOrgSecretCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update org secret created response has a 4xx status code
func (o *UpdateOrgSecretCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update org secret created response has a 5xx status code
func (o *UpdateOrgSecretCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update org secret created response a status code equal to that given
func (o *UpdateOrgSecretCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update org secret created response
func (o *UpdateOrgSecretCreated) Code() int {
	return 201
}

func (o *UpdateOrgSecretCreated) Error() string {
	return fmt.Sprintf("[PUT /orgs/{org}/actions/secrets/{secretname}][%d] updateOrgSecretCreated", 201)
}

func (o *UpdateOrgSecretCreated) String() string {
	return fmt.Sprintf("[PUT /orgs/{org}/actions/secrets/{secretname}][%d] updateOrgSecretCreated", 201)
}

func (o *UpdateOrgSecretCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateOrgSecretNoContent creates a UpdateOrgSecretNoContent with default headers values
func NewUpdateOrgSecretNoContent() *UpdateOrgSecretNoContent {
	return &UpdateOrgSecretNoContent{}
}

/*
UpdateOrgSecretNoContent describes a response with status code 204, with default header values.

response when updating a secret
*/
type UpdateOrgSecretNoContent struct {
}

// IsSuccess returns true when this update org secret no content response has a 2xx status code
func (o *UpdateOrgSecretNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update org secret no content response has a 3xx status code
func (o *UpdateOrgSecretNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update org secret no content response has a 4xx status code
func (o *UpdateOrgSecretNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update org secret no content response has a 5xx status code
func (o *UpdateOrgSecretNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update org secret no content response a status code equal to that given
func (o *UpdateOrgSecretNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the update org secret no content response
func (o *UpdateOrgSecretNoContent) Code() int {
	return 204
}

func (o *UpdateOrgSecretNoContent) Error() string {
	return fmt.Sprintf("[PUT /orgs/{org}/actions/secrets/{secretname}][%d] updateOrgSecretNoContent", 204)
}

func (o *UpdateOrgSecretNoContent) String() string {
	return fmt.Sprintf("[PUT /orgs/{org}/actions/secrets/{secretname}][%d] updateOrgSecretNoContent", 204)
}

func (o *UpdateOrgSecretNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateOrgSecretBadRequest creates a UpdateOrgSecretBadRequest with default headers values
func NewUpdateOrgSecretBadRequest() *UpdateOrgSecretBadRequest {
	return &UpdateOrgSecretBadRequest{}
}

/*
UpdateOrgSecretBadRequest describes a response with status code 400, with default header values.

APIError is error format response
*/
type UpdateOrgSecretBadRequest struct {
	Message string
	URL     string
}

// IsSuccess returns true when this update org secret bad request response has a 2xx status code
func (o *UpdateOrgSecretBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update org secret bad request response has a 3xx status code
func (o *UpdateOrgSecretBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update org secret bad request response has a 4xx status code
func (o *UpdateOrgSecretBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update org secret bad request response has a 5xx status code
func (o *UpdateOrgSecretBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update org secret bad request response a status code equal to that given
func (o *UpdateOrgSecretBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update org secret bad request response
func (o *UpdateOrgSecretBadRequest) Code() int {
	return 400
}

func (o *UpdateOrgSecretBadRequest) Error() string {
	return fmt.Sprintf("[PUT /orgs/{org}/actions/secrets/{secretname}][%d] updateOrgSecretBadRequest", 400)
}

func (o *UpdateOrgSecretBadRequest) String() string {
	return fmt.Sprintf("[PUT /orgs/{org}/actions/secrets/{secretname}][%d] updateOrgSecretBadRequest", 400)
}

func (o *UpdateOrgSecretBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header message
	hdrMessage := response.GetHeader("message")

	if hdrMessage != "" {
		o.Message = hdrMessage
	}

	// hydrates response header url
	hdrURL := response.GetHeader("url")

	if hdrURL != "" {
		o.URL = hdrURL
	}

	return nil
}

// NewUpdateOrgSecretNotFound creates a UpdateOrgSecretNotFound with default headers values
func NewUpdateOrgSecretNotFound() *UpdateOrgSecretNotFound {
	return &UpdateOrgSecretNotFound{}
}

/*
UpdateOrgSecretNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type UpdateOrgSecretNotFound struct {
}

// IsSuccess returns true when this update org secret not found response has a 2xx status code
func (o *UpdateOrgSecretNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update org secret not found response has a 3xx status code
func (o *UpdateOrgSecretNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update org secret not found response has a 4xx status code
func (o *UpdateOrgSecretNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update org secret not found response has a 5xx status code
func (o *UpdateOrgSecretNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update org secret not found response a status code equal to that given
func (o *UpdateOrgSecretNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update org secret not found response
func (o *UpdateOrgSecretNotFound) Code() int {
	return 404
}

func (o *UpdateOrgSecretNotFound) Error() string {
	return fmt.Sprintf("[PUT /orgs/{org}/actions/secrets/{secretname}][%d] updateOrgSecretNotFound", 404)
}

func (o *UpdateOrgSecretNotFound) String() string {
	return fmt.Sprintf("[PUT /orgs/{org}/actions/secrets/{secretname}][%d] updateOrgSecretNotFound", 404)
}

func (o *UpdateOrgSecretNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
