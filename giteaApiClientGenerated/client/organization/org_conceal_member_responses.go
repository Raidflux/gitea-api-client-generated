// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// OrgConcealMemberReader is a Reader for the OrgConcealMember structure.
type OrgConcealMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrgConcealMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewOrgConcealMemberNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewOrgConcealMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOrgConcealMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /orgs/{org}/public_members/{username}] orgConcealMember", response, response.Code())
	}
}

// NewOrgConcealMemberNoContent creates a OrgConcealMemberNoContent with default headers values
func NewOrgConcealMemberNoContent() *OrgConcealMemberNoContent {
	return &OrgConcealMemberNoContent{}
}

/*
OrgConcealMemberNoContent describes a response with status code 204, with default header values.

APIEmpty is an empty response
*/
type OrgConcealMemberNoContent struct {
}

// IsSuccess returns true when this org conceal member no content response has a 2xx status code
func (o *OrgConcealMemberNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this org conceal member no content response has a 3xx status code
func (o *OrgConcealMemberNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this org conceal member no content response has a 4xx status code
func (o *OrgConcealMemberNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this org conceal member no content response has a 5xx status code
func (o *OrgConcealMemberNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this org conceal member no content response a status code equal to that given
func (o *OrgConcealMemberNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the org conceal member no content response
func (o *OrgConcealMemberNoContent) Code() int {
	return 204
}

func (o *OrgConcealMemberNoContent) Error() string {
	return fmt.Sprintf("[DELETE /orgs/{org}/public_members/{username}][%d] orgConcealMemberNoContent", 204)
}

func (o *OrgConcealMemberNoContent) String() string {
	return fmt.Sprintf("[DELETE /orgs/{org}/public_members/{username}][%d] orgConcealMemberNoContent", 204)
}

func (o *OrgConcealMemberNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOrgConcealMemberForbidden creates a OrgConcealMemberForbidden with default headers values
func NewOrgConcealMemberForbidden() *OrgConcealMemberForbidden {
	return &OrgConcealMemberForbidden{}
}

/*
OrgConcealMemberForbidden describes a response with status code 403, with default header values.

APIForbiddenError is a forbidden error response
*/
type OrgConcealMemberForbidden struct {
	Message string
	URL     string
}

// IsSuccess returns true when this org conceal member forbidden response has a 2xx status code
func (o *OrgConcealMemberForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this org conceal member forbidden response has a 3xx status code
func (o *OrgConcealMemberForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this org conceal member forbidden response has a 4xx status code
func (o *OrgConcealMemberForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this org conceal member forbidden response has a 5xx status code
func (o *OrgConcealMemberForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this org conceal member forbidden response a status code equal to that given
func (o *OrgConcealMemberForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the org conceal member forbidden response
func (o *OrgConcealMemberForbidden) Code() int {
	return 403
}

func (o *OrgConcealMemberForbidden) Error() string {
	return fmt.Sprintf("[DELETE /orgs/{org}/public_members/{username}][%d] orgConcealMemberForbidden", 403)
}

func (o *OrgConcealMemberForbidden) String() string {
	return fmt.Sprintf("[DELETE /orgs/{org}/public_members/{username}][%d] orgConcealMemberForbidden", 403)
}

func (o *OrgConcealMemberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewOrgConcealMemberNotFound creates a OrgConcealMemberNotFound with default headers values
func NewOrgConcealMemberNotFound() *OrgConcealMemberNotFound {
	return &OrgConcealMemberNotFound{}
}

/*
OrgConcealMemberNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type OrgConcealMemberNotFound struct {
}

// IsSuccess returns true when this org conceal member not found response has a 2xx status code
func (o *OrgConcealMemberNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this org conceal member not found response has a 3xx status code
func (o *OrgConcealMemberNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this org conceal member not found response has a 4xx status code
func (o *OrgConcealMemberNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this org conceal member not found response has a 5xx status code
func (o *OrgConcealMemberNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this org conceal member not found response a status code equal to that given
func (o *OrgConcealMemberNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the org conceal member not found response
func (o *OrgConcealMemberNotFound) Code() int {
	return 404
}

func (o *OrgConcealMemberNotFound) Error() string {
	return fmt.Sprintf("[DELETE /orgs/{org}/public_members/{username}][%d] orgConcealMemberNotFound", 404)
}

func (o *OrgConcealMemberNotFound) String() string {
	return fmt.Sprintf("[DELETE /orgs/{org}/public_members/{username}][%d] orgConcealMemberNotFound", 404)
}

func (o *OrgConcealMemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
