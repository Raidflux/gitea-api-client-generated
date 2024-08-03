// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// OrganizationBlockUserReader is a Reader for the OrganizationBlockUser structure.
type OrganizationBlockUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationBlockUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewOrganizationBlockUserNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewOrganizationBlockUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewOrganizationBlockUserUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /orgs/{org}/blocks/{username}] organizationBlockUser", response, response.Code())
	}
}

// NewOrganizationBlockUserNoContent creates a OrganizationBlockUserNoContent with default headers values
func NewOrganizationBlockUserNoContent() *OrganizationBlockUserNoContent {
	return &OrganizationBlockUserNoContent{}
}

/*
OrganizationBlockUserNoContent describes a response with status code 204, with default header values.

APIEmpty is an empty response
*/
type OrganizationBlockUserNoContent struct {
}

// IsSuccess returns true when this organization block user no content response has a 2xx status code
func (o *OrganizationBlockUserNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this organization block user no content response has a 3xx status code
func (o *OrganizationBlockUserNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organization block user no content response has a 4xx status code
func (o *OrganizationBlockUserNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this organization block user no content response has a 5xx status code
func (o *OrganizationBlockUserNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this organization block user no content response a status code equal to that given
func (o *OrganizationBlockUserNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the organization block user no content response
func (o *OrganizationBlockUserNoContent) Code() int {
	return 204
}

func (o *OrganizationBlockUserNoContent) Error() string {
	return fmt.Sprintf("[PUT /orgs/{org}/blocks/{username}][%d] organizationBlockUserNoContent", 204)
}

func (o *OrganizationBlockUserNoContent) String() string {
	return fmt.Sprintf("[PUT /orgs/{org}/blocks/{username}][%d] organizationBlockUserNoContent", 204)
}

func (o *OrganizationBlockUserNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOrganizationBlockUserNotFound creates a OrganizationBlockUserNotFound with default headers values
func NewOrganizationBlockUserNotFound() *OrganizationBlockUserNotFound {
	return &OrganizationBlockUserNotFound{}
}

/*
OrganizationBlockUserNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type OrganizationBlockUserNotFound struct {
}

// IsSuccess returns true when this organization block user not found response has a 2xx status code
func (o *OrganizationBlockUserNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organization block user not found response has a 3xx status code
func (o *OrganizationBlockUserNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organization block user not found response has a 4xx status code
func (o *OrganizationBlockUserNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this organization block user not found response has a 5xx status code
func (o *OrganizationBlockUserNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this organization block user not found response a status code equal to that given
func (o *OrganizationBlockUserNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the organization block user not found response
func (o *OrganizationBlockUserNotFound) Code() int {
	return 404
}

func (o *OrganizationBlockUserNotFound) Error() string {
	return fmt.Sprintf("[PUT /orgs/{org}/blocks/{username}][%d] organizationBlockUserNotFound", 404)
}

func (o *OrganizationBlockUserNotFound) String() string {
	return fmt.Sprintf("[PUT /orgs/{org}/blocks/{username}][%d] organizationBlockUserNotFound", 404)
}

func (o *OrganizationBlockUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOrganizationBlockUserUnprocessableEntity creates a OrganizationBlockUserUnprocessableEntity with default headers values
func NewOrganizationBlockUserUnprocessableEntity() *OrganizationBlockUserUnprocessableEntity {
	return &OrganizationBlockUserUnprocessableEntity{}
}

/*
OrganizationBlockUserUnprocessableEntity describes a response with status code 422, with default header values.

APIValidationError is error format response related to input validation
*/
type OrganizationBlockUserUnprocessableEntity struct {
	Message string
	URL     string
}

// IsSuccess returns true when this organization block user unprocessable entity response has a 2xx status code
func (o *OrganizationBlockUserUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organization block user unprocessable entity response has a 3xx status code
func (o *OrganizationBlockUserUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organization block user unprocessable entity response has a 4xx status code
func (o *OrganizationBlockUserUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this organization block user unprocessable entity response has a 5xx status code
func (o *OrganizationBlockUserUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this organization block user unprocessable entity response a status code equal to that given
func (o *OrganizationBlockUserUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the organization block user unprocessable entity response
func (o *OrganizationBlockUserUnprocessableEntity) Code() int {
	return 422
}

func (o *OrganizationBlockUserUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /orgs/{org}/blocks/{username}][%d] organizationBlockUserUnprocessableEntity", 422)
}

func (o *OrganizationBlockUserUnprocessableEntity) String() string {
	return fmt.Sprintf("[PUT /orgs/{org}/blocks/{username}][%d] organizationBlockUserUnprocessableEntity", 422)
}

func (o *OrganizationBlockUserUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
