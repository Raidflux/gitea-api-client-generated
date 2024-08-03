// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// OrgDeleteAvatarReader is a Reader for the OrgDeleteAvatar structure.
type OrgDeleteAvatarReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrgDeleteAvatarReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewOrgDeleteAvatarNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewOrgDeleteAvatarNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /orgs/{org}/avatar] orgDeleteAvatar", response, response.Code())
	}
}

// NewOrgDeleteAvatarNoContent creates a OrgDeleteAvatarNoContent with default headers values
func NewOrgDeleteAvatarNoContent() *OrgDeleteAvatarNoContent {
	return &OrgDeleteAvatarNoContent{}
}

/*
OrgDeleteAvatarNoContent describes a response with status code 204, with default header values.

APIEmpty is an empty response
*/
type OrgDeleteAvatarNoContent struct {
}

// IsSuccess returns true when this org delete avatar no content response has a 2xx status code
func (o *OrgDeleteAvatarNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this org delete avatar no content response has a 3xx status code
func (o *OrgDeleteAvatarNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this org delete avatar no content response has a 4xx status code
func (o *OrgDeleteAvatarNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this org delete avatar no content response has a 5xx status code
func (o *OrgDeleteAvatarNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this org delete avatar no content response a status code equal to that given
func (o *OrgDeleteAvatarNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the org delete avatar no content response
func (o *OrgDeleteAvatarNoContent) Code() int {
	return 204
}

func (o *OrgDeleteAvatarNoContent) Error() string {
	return fmt.Sprintf("[DELETE /orgs/{org}/avatar][%d] orgDeleteAvatarNoContent", 204)
}

func (o *OrgDeleteAvatarNoContent) String() string {
	return fmt.Sprintf("[DELETE /orgs/{org}/avatar][%d] orgDeleteAvatarNoContent", 204)
}

func (o *OrgDeleteAvatarNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOrgDeleteAvatarNotFound creates a OrgDeleteAvatarNotFound with default headers values
func NewOrgDeleteAvatarNotFound() *OrgDeleteAvatarNotFound {
	return &OrgDeleteAvatarNotFound{}
}

/*
OrgDeleteAvatarNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type OrgDeleteAvatarNotFound struct {
}

// IsSuccess returns true when this org delete avatar not found response has a 2xx status code
func (o *OrgDeleteAvatarNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this org delete avatar not found response has a 3xx status code
func (o *OrgDeleteAvatarNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this org delete avatar not found response has a 4xx status code
func (o *OrgDeleteAvatarNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this org delete avatar not found response has a 5xx status code
func (o *OrgDeleteAvatarNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this org delete avatar not found response a status code equal to that given
func (o *OrgDeleteAvatarNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the org delete avatar not found response
func (o *OrgDeleteAvatarNotFound) Code() int {
	return 404
}

func (o *OrgDeleteAvatarNotFound) Error() string {
	return fmt.Sprintf("[DELETE /orgs/{org}/avatar][%d] orgDeleteAvatarNotFound", 404)
}

func (o *OrgDeleteAvatarNotFound) String() string {
	return fmt.Sprintf("[DELETE /orgs/{org}/avatar][%d] orgDeleteAvatarNotFound", 404)
}

func (o *OrgDeleteAvatarNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
