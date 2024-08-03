// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// OrgPublicizeMemberReader is a Reader for the OrgPublicizeMember structure.
type OrgPublicizeMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrgPublicizeMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewOrgPublicizeMemberNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewOrgPublicizeMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOrgPublicizeMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /orgs/{org}/public_members/{username}] orgPublicizeMember", response, response.Code())
	}
}

// NewOrgPublicizeMemberNoContent creates a OrgPublicizeMemberNoContent with default headers values
func NewOrgPublicizeMemberNoContent() *OrgPublicizeMemberNoContent {
	return &OrgPublicizeMemberNoContent{}
}

/*
OrgPublicizeMemberNoContent describes a response with status code 204, with default header values.

membership publicized
*/
type OrgPublicizeMemberNoContent struct {
}

// IsSuccess returns true when this org publicize member no content response has a 2xx status code
func (o *OrgPublicizeMemberNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this org publicize member no content response has a 3xx status code
func (o *OrgPublicizeMemberNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this org publicize member no content response has a 4xx status code
func (o *OrgPublicizeMemberNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this org publicize member no content response has a 5xx status code
func (o *OrgPublicizeMemberNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this org publicize member no content response a status code equal to that given
func (o *OrgPublicizeMemberNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the org publicize member no content response
func (o *OrgPublicizeMemberNoContent) Code() int {
	return 204
}

func (o *OrgPublicizeMemberNoContent) Error() string {
	return fmt.Sprintf("[PUT /orgs/{org}/public_members/{username}][%d] orgPublicizeMemberNoContent", 204)
}

func (o *OrgPublicizeMemberNoContent) String() string {
	return fmt.Sprintf("[PUT /orgs/{org}/public_members/{username}][%d] orgPublicizeMemberNoContent", 204)
}

func (o *OrgPublicizeMemberNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOrgPublicizeMemberForbidden creates a OrgPublicizeMemberForbidden with default headers values
func NewOrgPublicizeMemberForbidden() *OrgPublicizeMemberForbidden {
	return &OrgPublicizeMemberForbidden{}
}

/*
OrgPublicizeMemberForbidden describes a response with status code 403, with default header values.

APIForbiddenError is a forbidden error response
*/
type OrgPublicizeMemberForbidden struct {
	Message string
	URL     string
}

// IsSuccess returns true when this org publicize member forbidden response has a 2xx status code
func (o *OrgPublicizeMemberForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this org publicize member forbidden response has a 3xx status code
func (o *OrgPublicizeMemberForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this org publicize member forbidden response has a 4xx status code
func (o *OrgPublicizeMemberForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this org publicize member forbidden response has a 5xx status code
func (o *OrgPublicizeMemberForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this org publicize member forbidden response a status code equal to that given
func (o *OrgPublicizeMemberForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the org publicize member forbidden response
func (o *OrgPublicizeMemberForbidden) Code() int {
	return 403
}

func (o *OrgPublicizeMemberForbidden) Error() string {
	return fmt.Sprintf("[PUT /orgs/{org}/public_members/{username}][%d] orgPublicizeMemberForbidden", 403)
}

func (o *OrgPublicizeMemberForbidden) String() string {
	return fmt.Sprintf("[PUT /orgs/{org}/public_members/{username}][%d] orgPublicizeMemberForbidden", 403)
}

func (o *OrgPublicizeMemberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewOrgPublicizeMemberNotFound creates a OrgPublicizeMemberNotFound with default headers values
func NewOrgPublicizeMemberNotFound() *OrgPublicizeMemberNotFound {
	return &OrgPublicizeMemberNotFound{}
}

/*
OrgPublicizeMemberNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type OrgPublicizeMemberNotFound struct {
}

// IsSuccess returns true when this org publicize member not found response has a 2xx status code
func (o *OrgPublicizeMemberNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this org publicize member not found response has a 3xx status code
func (o *OrgPublicizeMemberNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this org publicize member not found response has a 4xx status code
func (o *OrgPublicizeMemberNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this org publicize member not found response has a 5xx status code
func (o *OrgPublicizeMemberNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this org publicize member not found response a status code equal to that given
func (o *OrgPublicizeMemberNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the org publicize member not found response
func (o *OrgPublicizeMemberNotFound) Code() int {
	return 404
}

func (o *OrgPublicizeMemberNotFound) Error() string {
	return fmt.Sprintf("[PUT /orgs/{org}/public_members/{username}][%d] orgPublicizeMemberNotFound", 404)
}

func (o *OrgPublicizeMemberNotFound) String() string {
	return fmt.Sprintf("[PUT /orgs/{org}/public_members/{username}][%d] orgPublicizeMemberNotFound", 404)
}

func (o *OrgPublicizeMemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
