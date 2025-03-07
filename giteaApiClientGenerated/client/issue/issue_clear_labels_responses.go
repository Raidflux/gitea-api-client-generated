// Code generated by go-swagger; DO NOT EDIT.

package issue

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IssueClearLabelsReader is a Reader for the IssueClearLabels structure.
type IssueClearLabelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IssueClearLabelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIssueClearLabelsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewIssueClearLabelsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewIssueClearLabelsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /repos/{owner}/{repo}/issues/{index}/labels] issueClearLabels", response, response.Code())
	}
}

// NewIssueClearLabelsNoContent creates a IssueClearLabelsNoContent with default headers values
func NewIssueClearLabelsNoContent() *IssueClearLabelsNoContent {
	return &IssueClearLabelsNoContent{}
}

/*
IssueClearLabelsNoContent describes a response with status code 204, with default header values.

APIEmpty is an empty response
*/
type IssueClearLabelsNoContent struct {
}

// IsSuccess returns true when this issue clear labels no content response has a 2xx status code
func (o *IssueClearLabelsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this issue clear labels no content response has a 3xx status code
func (o *IssueClearLabelsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue clear labels no content response has a 4xx status code
func (o *IssueClearLabelsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this issue clear labels no content response has a 5xx status code
func (o *IssueClearLabelsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this issue clear labels no content response a status code equal to that given
func (o *IssueClearLabelsNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the issue clear labels no content response
func (o *IssueClearLabelsNoContent) Code() int {
	return 204
}

func (o *IssueClearLabelsNoContent) Error() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/issues/{index}/labels][%d] issueClearLabelsNoContent", 204)
}

func (o *IssueClearLabelsNoContent) String() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/issues/{index}/labels][%d] issueClearLabelsNoContent", 204)
}

func (o *IssueClearLabelsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIssueClearLabelsForbidden creates a IssueClearLabelsForbidden with default headers values
func NewIssueClearLabelsForbidden() *IssueClearLabelsForbidden {
	return &IssueClearLabelsForbidden{}
}

/*
IssueClearLabelsForbidden describes a response with status code 403, with default header values.

APIForbiddenError is a forbidden error response
*/
type IssueClearLabelsForbidden struct {
	Message string
	URL     string
}

// IsSuccess returns true when this issue clear labels forbidden response has a 2xx status code
func (o *IssueClearLabelsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this issue clear labels forbidden response has a 3xx status code
func (o *IssueClearLabelsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue clear labels forbidden response has a 4xx status code
func (o *IssueClearLabelsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this issue clear labels forbidden response has a 5xx status code
func (o *IssueClearLabelsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this issue clear labels forbidden response a status code equal to that given
func (o *IssueClearLabelsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the issue clear labels forbidden response
func (o *IssueClearLabelsForbidden) Code() int {
	return 403
}

func (o *IssueClearLabelsForbidden) Error() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/issues/{index}/labels][%d] issueClearLabelsForbidden", 403)
}

func (o *IssueClearLabelsForbidden) String() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/issues/{index}/labels][%d] issueClearLabelsForbidden", 403)
}

func (o *IssueClearLabelsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewIssueClearLabelsNotFound creates a IssueClearLabelsNotFound with default headers values
func NewIssueClearLabelsNotFound() *IssueClearLabelsNotFound {
	return &IssueClearLabelsNotFound{}
}

/*
IssueClearLabelsNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type IssueClearLabelsNotFound struct {
}

// IsSuccess returns true when this issue clear labels not found response has a 2xx status code
func (o *IssueClearLabelsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this issue clear labels not found response has a 3xx status code
func (o *IssueClearLabelsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue clear labels not found response has a 4xx status code
func (o *IssueClearLabelsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this issue clear labels not found response has a 5xx status code
func (o *IssueClearLabelsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this issue clear labels not found response a status code equal to that given
func (o *IssueClearLabelsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the issue clear labels not found response
func (o *IssueClearLabelsNotFound) Code() int {
	return 404
}

func (o *IssueClearLabelsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/issues/{index}/labels][%d] issueClearLabelsNotFound", 404)
}

func (o *IssueClearLabelsNotFound) String() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/issues/{index}/labels][%d] issueClearLabelsNotFound", 404)
}

func (o *IssueClearLabelsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
