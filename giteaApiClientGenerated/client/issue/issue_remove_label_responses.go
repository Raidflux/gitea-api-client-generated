// Code generated by go-swagger; DO NOT EDIT.

package issue

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IssueRemoveLabelReader is a Reader for the IssueRemoveLabel structure.
type IssueRemoveLabelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IssueRemoveLabelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIssueRemoveLabelNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewIssueRemoveLabelForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewIssueRemoveLabelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewIssueRemoveLabelUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /repos/{owner}/{repo}/issues/{index}/labels/{id}] issueRemoveLabel", response, response.Code())
	}
}

// NewIssueRemoveLabelNoContent creates a IssueRemoveLabelNoContent with default headers values
func NewIssueRemoveLabelNoContent() *IssueRemoveLabelNoContent {
	return &IssueRemoveLabelNoContent{}
}

/*
IssueRemoveLabelNoContent describes a response with status code 204, with default header values.

APIEmpty is an empty response
*/
type IssueRemoveLabelNoContent struct {
}

// IsSuccess returns true when this issue remove label no content response has a 2xx status code
func (o *IssueRemoveLabelNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this issue remove label no content response has a 3xx status code
func (o *IssueRemoveLabelNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue remove label no content response has a 4xx status code
func (o *IssueRemoveLabelNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this issue remove label no content response has a 5xx status code
func (o *IssueRemoveLabelNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this issue remove label no content response a status code equal to that given
func (o *IssueRemoveLabelNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the issue remove label no content response
func (o *IssueRemoveLabelNoContent) Code() int {
	return 204
}

func (o *IssueRemoveLabelNoContent) Error() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/issues/{index}/labels/{id}][%d] issueRemoveLabelNoContent", 204)
}

func (o *IssueRemoveLabelNoContent) String() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/issues/{index}/labels/{id}][%d] issueRemoveLabelNoContent", 204)
}

func (o *IssueRemoveLabelNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIssueRemoveLabelForbidden creates a IssueRemoveLabelForbidden with default headers values
func NewIssueRemoveLabelForbidden() *IssueRemoveLabelForbidden {
	return &IssueRemoveLabelForbidden{}
}

/*
IssueRemoveLabelForbidden describes a response with status code 403, with default header values.

APIForbiddenError is a forbidden error response
*/
type IssueRemoveLabelForbidden struct {
	Message string
	URL     string
}

// IsSuccess returns true when this issue remove label forbidden response has a 2xx status code
func (o *IssueRemoveLabelForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this issue remove label forbidden response has a 3xx status code
func (o *IssueRemoveLabelForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue remove label forbidden response has a 4xx status code
func (o *IssueRemoveLabelForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this issue remove label forbidden response has a 5xx status code
func (o *IssueRemoveLabelForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this issue remove label forbidden response a status code equal to that given
func (o *IssueRemoveLabelForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the issue remove label forbidden response
func (o *IssueRemoveLabelForbidden) Code() int {
	return 403
}

func (o *IssueRemoveLabelForbidden) Error() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/issues/{index}/labels/{id}][%d] issueRemoveLabelForbidden", 403)
}

func (o *IssueRemoveLabelForbidden) String() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/issues/{index}/labels/{id}][%d] issueRemoveLabelForbidden", 403)
}

func (o *IssueRemoveLabelForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewIssueRemoveLabelNotFound creates a IssueRemoveLabelNotFound with default headers values
func NewIssueRemoveLabelNotFound() *IssueRemoveLabelNotFound {
	return &IssueRemoveLabelNotFound{}
}

/*
IssueRemoveLabelNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type IssueRemoveLabelNotFound struct {
}

// IsSuccess returns true when this issue remove label not found response has a 2xx status code
func (o *IssueRemoveLabelNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this issue remove label not found response has a 3xx status code
func (o *IssueRemoveLabelNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue remove label not found response has a 4xx status code
func (o *IssueRemoveLabelNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this issue remove label not found response has a 5xx status code
func (o *IssueRemoveLabelNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this issue remove label not found response a status code equal to that given
func (o *IssueRemoveLabelNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the issue remove label not found response
func (o *IssueRemoveLabelNotFound) Code() int {
	return 404
}

func (o *IssueRemoveLabelNotFound) Error() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/issues/{index}/labels/{id}][%d] issueRemoveLabelNotFound", 404)
}

func (o *IssueRemoveLabelNotFound) String() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/issues/{index}/labels/{id}][%d] issueRemoveLabelNotFound", 404)
}

func (o *IssueRemoveLabelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIssueRemoveLabelUnprocessableEntity creates a IssueRemoveLabelUnprocessableEntity with default headers values
func NewIssueRemoveLabelUnprocessableEntity() *IssueRemoveLabelUnprocessableEntity {
	return &IssueRemoveLabelUnprocessableEntity{}
}

/*
IssueRemoveLabelUnprocessableEntity describes a response with status code 422, with default header values.

APIValidationError is error format response related to input validation
*/
type IssueRemoveLabelUnprocessableEntity struct {
	Message string
	URL     string
}

// IsSuccess returns true when this issue remove label unprocessable entity response has a 2xx status code
func (o *IssueRemoveLabelUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this issue remove label unprocessable entity response has a 3xx status code
func (o *IssueRemoveLabelUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue remove label unprocessable entity response has a 4xx status code
func (o *IssueRemoveLabelUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this issue remove label unprocessable entity response has a 5xx status code
func (o *IssueRemoveLabelUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this issue remove label unprocessable entity response a status code equal to that given
func (o *IssueRemoveLabelUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the issue remove label unprocessable entity response
func (o *IssueRemoveLabelUnprocessableEntity) Code() int {
	return 422
}

func (o *IssueRemoveLabelUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/issues/{index}/labels/{id}][%d] issueRemoveLabelUnprocessableEntity", 422)
}

func (o *IssueRemoveLabelUnprocessableEntity) String() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/issues/{index}/labels/{id}][%d] issueRemoveLabelUnprocessableEntity", 422)
}

func (o *IssueRemoveLabelUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
