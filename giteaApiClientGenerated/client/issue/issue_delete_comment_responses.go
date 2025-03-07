// Code generated by go-swagger; DO NOT EDIT.

package issue

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IssueDeleteCommentReader is a Reader for the IssueDeleteComment structure.
type IssueDeleteCommentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IssueDeleteCommentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIssueDeleteCommentNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewIssueDeleteCommentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewIssueDeleteCommentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /repos/{owner}/{repo}/issues/comments/{id}] issueDeleteComment", response, response.Code())
	}
}

// NewIssueDeleteCommentNoContent creates a IssueDeleteCommentNoContent with default headers values
func NewIssueDeleteCommentNoContent() *IssueDeleteCommentNoContent {
	return &IssueDeleteCommentNoContent{}
}

/*
IssueDeleteCommentNoContent describes a response with status code 204, with default header values.

APIEmpty is an empty response
*/
type IssueDeleteCommentNoContent struct {
}

// IsSuccess returns true when this issue delete comment no content response has a 2xx status code
func (o *IssueDeleteCommentNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this issue delete comment no content response has a 3xx status code
func (o *IssueDeleteCommentNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue delete comment no content response has a 4xx status code
func (o *IssueDeleteCommentNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this issue delete comment no content response has a 5xx status code
func (o *IssueDeleteCommentNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this issue delete comment no content response a status code equal to that given
func (o *IssueDeleteCommentNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the issue delete comment no content response
func (o *IssueDeleteCommentNoContent) Code() int {
	return 204
}

func (o *IssueDeleteCommentNoContent) Error() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/issues/comments/{id}][%d] issueDeleteCommentNoContent", 204)
}

func (o *IssueDeleteCommentNoContent) String() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/issues/comments/{id}][%d] issueDeleteCommentNoContent", 204)
}

func (o *IssueDeleteCommentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIssueDeleteCommentForbidden creates a IssueDeleteCommentForbidden with default headers values
func NewIssueDeleteCommentForbidden() *IssueDeleteCommentForbidden {
	return &IssueDeleteCommentForbidden{}
}

/*
IssueDeleteCommentForbidden describes a response with status code 403, with default header values.

APIForbiddenError is a forbidden error response
*/
type IssueDeleteCommentForbidden struct {
	Message string
	URL     string
}

// IsSuccess returns true when this issue delete comment forbidden response has a 2xx status code
func (o *IssueDeleteCommentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this issue delete comment forbidden response has a 3xx status code
func (o *IssueDeleteCommentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue delete comment forbidden response has a 4xx status code
func (o *IssueDeleteCommentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this issue delete comment forbidden response has a 5xx status code
func (o *IssueDeleteCommentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this issue delete comment forbidden response a status code equal to that given
func (o *IssueDeleteCommentForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the issue delete comment forbidden response
func (o *IssueDeleteCommentForbidden) Code() int {
	return 403
}

func (o *IssueDeleteCommentForbidden) Error() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/issues/comments/{id}][%d] issueDeleteCommentForbidden", 403)
}

func (o *IssueDeleteCommentForbidden) String() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/issues/comments/{id}][%d] issueDeleteCommentForbidden", 403)
}

func (o *IssueDeleteCommentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewIssueDeleteCommentNotFound creates a IssueDeleteCommentNotFound with default headers values
func NewIssueDeleteCommentNotFound() *IssueDeleteCommentNotFound {
	return &IssueDeleteCommentNotFound{}
}

/*
IssueDeleteCommentNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type IssueDeleteCommentNotFound struct {
}

// IsSuccess returns true when this issue delete comment not found response has a 2xx status code
func (o *IssueDeleteCommentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this issue delete comment not found response has a 3xx status code
func (o *IssueDeleteCommentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue delete comment not found response has a 4xx status code
func (o *IssueDeleteCommentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this issue delete comment not found response has a 5xx status code
func (o *IssueDeleteCommentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this issue delete comment not found response a status code equal to that given
func (o *IssueDeleteCommentNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the issue delete comment not found response
func (o *IssueDeleteCommentNotFound) Code() int {
	return 404
}

func (o *IssueDeleteCommentNotFound) Error() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/issues/comments/{id}][%d] issueDeleteCommentNotFound", 404)
}

func (o *IssueDeleteCommentNotFound) String() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/issues/comments/{id}][%d] issueDeleteCommentNotFound", 404)
}

func (o *IssueDeleteCommentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
