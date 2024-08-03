// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// OrganizationCheckUserBlockReader is a Reader for the OrganizationCheckUserBlock structure.
type OrganizationCheckUserBlockReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationCheckUserBlockReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewOrganizationCheckUserBlockNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewOrganizationCheckUserBlockNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /orgs/{org}/blocks/{username}] organizationCheckUserBlock", response, response.Code())
	}
}

// NewOrganizationCheckUserBlockNoContent creates a OrganizationCheckUserBlockNoContent with default headers values
func NewOrganizationCheckUserBlockNoContent() *OrganizationCheckUserBlockNoContent {
	return &OrganizationCheckUserBlockNoContent{}
}

/*
OrganizationCheckUserBlockNoContent describes a response with status code 204, with default header values.

APIEmpty is an empty response
*/
type OrganizationCheckUserBlockNoContent struct {
}

// IsSuccess returns true when this organization check user block no content response has a 2xx status code
func (o *OrganizationCheckUserBlockNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this organization check user block no content response has a 3xx status code
func (o *OrganizationCheckUserBlockNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organization check user block no content response has a 4xx status code
func (o *OrganizationCheckUserBlockNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this organization check user block no content response has a 5xx status code
func (o *OrganizationCheckUserBlockNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this organization check user block no content response a status code equal to that given
func (o *OrganizationCheckUserBlockNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the organization check user block no content response
func (o *OrganizationCheckUserBlockNoContent) Code() int {
	return 204
}

func (o *OrganizationCheckUserBlockNoContent) Error() string {
	return fmt.Sprintf("[GET /orgs/{org}/blocks/{username}][%d] organizationCheckUserBlockNoContent", 204)
}

func (o *OrganizationCheckUserBlockNoContent) String() string {
	return fmt.Sprintf("[GET /orgs/{org}/blocks/{username}][%d] organizationCheckUserBlockNoContent", 204)
}

func (o *OrganizationCheckUserBlockNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOrganizationCheckUserBlockNotFound creates a OrganizationCheckUserBlockNotFound with default headers values
func NewOrganizationCheckUserBlockNotFound() *OrganizationCheckUserBlockNotFound {
	return &OrganizationCheckUserBlockNotFound{}
}

/*
OrganizationCheckUserBlockNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type OrganizationCheckUserBlockNotFound struct {
}

// IsSuccess returns true when this organization check user block not found response has a 2xx status code
func (o *OrganizationCheckUserBlockNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organization check user block not found response has a 3xx status code
func (o *OrganizationCheckUserBlockNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organization check user block not found response has a 4xx status code
func (o *OrganizationCheckUserBlockNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this organization check user block not found response has a 5xx status code
func (o *OrganizationCheckUserBlockNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this organization check user block not found response a status code equal to that given
func (o *OrganizationCheckUserBlockNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the organization check user block not found response
func (o *OrganizationCheckUserBlockNotFound) Code() int {
	return 404
}

func (o *OrganizationCheckUserBlockNotFound) Error() string {
	return fmt.Sprintf("[GET /orgs/{org}/blocks/{username}][%d] organizationCheckUserBlockNotFound", 404)
}

func (o *OrganizationCheckUserBlockNotFound) String() string {
	return fmt.Sprintf("[GET /orgs/{org}/blocks/{username}][%d] organizationCheckUserBlockNotFound", 404)
}

func (o *OrganizationCheckUserBlockNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
