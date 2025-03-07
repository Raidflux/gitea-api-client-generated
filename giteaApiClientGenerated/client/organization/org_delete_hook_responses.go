// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// OrgDeleteHookReader is a Reader for the OrgDeleteHook structure.
type OrgDeleteHookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrgDeleteHookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewOrgDeleteHookNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewOrgDeleteHookNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /orgs/{org}/hooks/{id}] orgDeleteHook", response, response.Code())
	}
}

// NewOrgDeleteHookNoContent creates a OrgDeleteHookNoContent with default headers values
func NewOrgDeleteHookNoContent() *OrgDeleteHookNoContent {
	return &OrgDeleteHookNoContent{}
}

/*
OrgDeleteHookNoContent describes a response with status code 204, with default header values.

APIEmpty is an empty response
*/
type OrgDeleteHookNoContent struct {
}

// IsSuccess returns true when this org delete hook no content response has a 2xx status code
func (o *OrgDeleteHookNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this org delete hook no content response has a 3xx status code
func (o *OrgDeleteHookNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this org delete hook no content response has a 4xx status code
func (o *OrgDeleteHookNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this org delete hook no content response has a 5xx status code
func (o *OrgDeleteHookNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this org delete hook no content response a status code equal to that given
func (o *OrgDeleteHookNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the org delete hook no content response
func (o *OrgDeleteHookNoContent) Code() int {
	return 204
}

func (o *OrgDeleteHookNoContent) Error() string {
	return fmt.Sprintf("[DELETE /orgs/{org}/hooks/{id}][%d] orgDeleteHookNoContent", 204)
}

func (o *OrgDeleteHookNoContent) String() string {
	return fmt.Sprintf("[DELETE /orgs/{org}/hooks/{id}][%d] orgDeleteHookNoContent", 204)
}

func (o *OrgDeleteHookNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOrgDeleteHookNotFound creates a OrgDeleteHookNotFound with default headers values
func NewOrgDeleteHookNotFound() *OrgDeleteHookNotFound {
	return &OrgDeleteHookNotFound{}
}

/*
OrgDeleteHookNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type OrgDeleteHookNotFound struct {
}

// IsSuccess returns true when this org delete hook not found response has a 2xx status code
func (o *OrgDeleteHookNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this org delete hook not found response has a 3xx status code
func (o *OrgDeleteHookNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this org delete hook not found response has a 4xx status code
func (o *OrgDeleteHookNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this org delete hook not found response has a 5xx status code
func (o *OrgDeleteHookNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this org delete hook not found response a status code equal to that given
func (o *OrgDeleteHookNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the org delete hook not found response
func (o *OrgDeleteHookNotFound) Code() int {
	return 404
}

func (o *OrgDeleteHookNotFound) Error() string {
	return fmt.Sprintf("[DELETE /orgs/{org}/hooks/{id}][%d] orgDeleteHookNotFound", 404)
}

func (o *OrgDeleteHookNotFound) String() string {
	return fmt.Sprintf("[DELETE /orgs/{org}/hooks/{id}][%d] orgDeleteHookNotFound", 404)
}

func (o *OrgDeleteHookNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
