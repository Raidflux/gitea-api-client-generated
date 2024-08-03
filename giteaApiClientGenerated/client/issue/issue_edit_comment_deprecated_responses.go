// Code generated by go-swagger; DO NOT EDIT.

package issue

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"giteaApiClientGenerated/models"
)

// IssueEditCommentDeprecatedReader is a Reader for the IssueEditCommentDeprecated structure.
type IssueEditCommentDeprecatedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IssueEditCommentDeprecatedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIssueEditCommentDeprecatedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewIssueEditCommentDeprecatedNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewIssueEditCommentDeprecatedForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewIssueEditCommentDeprecatedNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /repos/{owner}/{repo}/issues/{index}/comments/{id}] issueEditCommentDeprecated", response, response.Code())
	}
}

// NewIssueEditCommentDeprecatedOK creates a IssueEditCommentDeprecatedOK with default headers values
func NewIssueEditCommentDeprecatedOK() *IssueEditCommentDeprecatedOK {
	return &IssueEditCommentDeprecatedOK{}
}

/*
IssueEditCommentDeprecatedOK describes a response with status code 200, with default header values.

Comment
*/
type IssueEditCommentDeprecatedOK struct {
	Payload *models.Comment
}

// IsSuccess returns true when this issue edit comment deprecated o k response has a 2xx status code
func (o *IssueEditCommentDeprecatedOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this issue edit comment deprecated o k response has a 3xx status code
func (o *IssueEditCommentDeprecatedOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue edit comment deprecated o k response has a 4xx status code
func (o *IssueEditCommentDeprecatedOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this issue edit comment deprecated o k response has a 5xx status code
func (o *IssueEditCommentDeprecatedOK) IsServerError() bool {
	return false
}

// IsCode returns true when this issue edit comment deprecated o k response a status code equal to that given
func (o *IssueEditCommentDeprecatedOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the issue edit comment deprecated o k response
func (o *IssueEditCommentDeprecatedOK) Code() int {
	return 200
}

func (o *IssueEditCommentDeprecatedOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}/issues/{index}/comments/{id}][%d] issueEditCommentDeprecatedOK %s", 200, payload)
}

func (o *IssueEditCommentDeprecatedOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}/issues/{index}/comments/{id}][%d] issueEditCommentDeprecatedOK %s", 200, payload)
}

func (o *IssueEditCommentDeprecatedOK) GetPayload() *models.Comment {
	return o.Payload
}

func (o *IssueEditCommentDeprecatedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Comment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIssueEditCommentDeprecatedNoContent creates a IssueEditCommentDeprecatedNoContent with default headers values
func NewIssueEditCommentDeprecatedNoContent() *IssueEditCommentDeprecatedNoContent {
	return &IssueEditCommentDeprecatedNoContent{}
}

/*
IssueEditCommentDeprecatedNoContent describes a response with status code 204, with default header values.

APIEmpty is an empty response
*/
type IssueEditCommentDeprecatedNoContent struct {
}

// IsSuccess returns true when this issue edit comment deprecated no content response has a 2xx status code
func (o *IssueEditCommentDeprecatedNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this issue edit comment deprecated no content response has a 3xx status code
func (o *IssueEditCommentDeprecatedNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue edit comment deprecated no content response has a 4xx status code
func (o *IssueEditCommentDeprecatedNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this issue edit comment deprecated no content response has a 5xx status code
func (o *IssueEditCommentDeprecatedNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this issue edit comment deprecated no content response a status code equal to that given
func (o *IssueEditCommentDeprecatedNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the issue edit comment deprecated no content response
func (o *IssueEditCommentDeprecatedNoContent) Code() int {
	return 204
}

func (o *IssueEditCommentDeprecatedNoContent) Error() string {
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}/issues/{index}/comments/{id}][%d] issueEditCommentDeprecatedNoContent", 204)
}

func (o *IssueEditCommentDeprecatedNoContent) String() string {
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}/issues/{index}/comments/{id}][%d] issueEditCommentDeprecatedNoContent", 204)
}

func (o *IssueEditCommentDeprecatedNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIssueEditCommentDeprecatedForbidden creates a IssueEditCommentDeprecatedForbidden with default headers values
func NewIssueEditCommentDeprecatedForbidden() *IssueEditCommentDeprecatedForbidden {
	return &IssueEditCommentDeprecatedForbidden{}
}

/*
IssueEditCommentDeprecatedForbidden describes a response with status code 403, with default header values.

APIForbiddenError is a forbidden error response
*/
type IssueEditCommentDeprecatedForbidden struct {
	Message string
	URL     string
}

// IsSuccess returns true when this issue edit comment deprecated forbidden response has a 2xx status code
func (o *IssueEditCommentDeprecatedForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this issue edit comment deprecated forbidden response has a 3xx status code
func (o *IssueEditCommentDeprecatedForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue edit comment deprecated forbidden response has a 4xx status code
func (o *IssueEditCommentDeprecatedForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this issue edit comment deprecated forbidden response has a 5xx status code
func (o *IssueEditCommentDeprecatedForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this issue edit comment deprecated forbidden response a status code equal to that given
func (o *IssueEditCommentDeprecatedForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the issue edit comment deprecated forbidden response
func (o *IssueEditCommentDeprecatedForbidden) Code() int {
	return 403
}

func (o *IssueEditCommentDeprecatedForbidden) Error() string {
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}/issues/{index}/comments/{id}][%d] issueEditCommentDeprecatedForbidden", 403)
}

func (o *IssueEditCommentDeprecatedForbidden) String() string {
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}/issues/{index}/comments/{id}][%d] issueEditCommentDeprecatedForbidden", 403)
}

func (o *IssueEditCommentDeprecatedForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewIssueEditCommentDeprecatedNotFound creates a IssueEditCommentDeprecatedNotFound with default headers values
func NewIssueEditCommentDeprecatedNotFound() *IssueEditCommentDeprecatedNotFound {
	return &IssueEditCommentDeprecatedNotFound{}
}

/*
IssueEditCommentDeprecatedNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type IssueEditCommentDeprecatedNotFound struct {
}

// IsSuccess returns true when this issue edit comment deprecated not found response has a 2xx status code
func (o *IssueEditCommentDeprecatedNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this issue edit comment deprecated not found response has a 3xx status code
func (o *IssueEditCommentDeprecatedNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue edit comment deprecated not found response has a 4xx status code
func (o *IssueEditCommentDeprecatedNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this issue edit comment deprecated not found response has a 5xx status code
func (o *IssueEditCommentDeprecatedNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this issue edit comment deprecated not found response a status code equal to that given
func (o *IssueEditCommentDeprecatedNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the issue edit comment deprecated not found response
func (o *IssueEditCommentDeprecatedNotFound) Code() int {
	return 404
}

func (o *IssueEditCommentDeprecatedNotFound) Error() string {
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}/issues/{index}/comments/{id}][%d] issueEditCommentDeprecatedNotFound", 404)
}

func (o *IssueEditCommentDeprecatedNotFound) String() string {
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}/issues/{index}/comments/{id}][%d] issueEditCommentDeprecatedNotFound", 404)
}

func (o *IssueEditCommentDeprecatedNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
