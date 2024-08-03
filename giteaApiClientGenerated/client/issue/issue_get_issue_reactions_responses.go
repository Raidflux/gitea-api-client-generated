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

	"github.com/Raidflux/gitea-api-client-generated/giteaApiClientGenerated/models"
)

// IssueGetIssueReactionsReader is a Reader for the IssueGetIssueReactions structure.
type IssueGetIssueReactionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IssueGetIssueReactionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIssueGetIssueReactionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewIssueGetIssueReactionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewIssueGetIssueReactionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /repos/{owner}/{repo}/issues/{index}/reactions] issueGetIssueReactions", response, response.Code())
	}
}

// NewIssueGetIssueReactionsOK creates a IssueGetIssueReactionsOK with default headers values
func NewIssueGetIssueReactionsOK() *IssueGetIssueReactionsOK {
	return &IssueGetIssueReactionsOK{}
}

/*
IssueGetIssueReactionsOK describes a response with status code 200, with default header values.

ReactionList
*/
type IssueGetIssueReactionsOK struct {
	Payload []*models.Reaction
}

// IsSuccess returns true when this issue get issue reactions o k response has a 2xx status code
func (o *IssueGetIssueReactionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this issue get issue reactions o k response has a 3xx status code
func (o *IssueGetIssueReactionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue get issue reactions o k response has a 4xx status code
func (o *IssueGetIssueReactionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this issue get issue reactions o k response has a 5xx status code
func (o *IssueGetIssueReactionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this issue get issue reactions o k response a status code equal to that given
func (o *IssueGetIssueReactionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the issue get issue reactions o k response
func (o *IssueGetIssueReactionsOK) Code() int {
	return 200
}

func (o *IssueGetIssueReactionsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/issues/{index}/reactions][%d] issueGetIssueReactionsOK %s", 200, payload)
}

func (o *IssueGetIssueReactionsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/issues/{index}/reactions][%d] issueGetIssueReactionsOK %s", 200, payload)
}

func (o *IssueGetIssueReactionsOK) GetPayload() []*models.Reaction {
	return o.Payload
}

func (o *IssueGetIssueReactionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIssueGetIssueReactionsForbidden creates a IssueGetIssueReactionsForbidden with default headers values
func NewIssueGetIssueReactionsForbidden() *IssueGetIssueReactionsForbidden {
	return &IssueGetIssueReactionsForbidden{}
}

/*
IssueGetIssueReactionsForbidden describes a response with status code 403, with default header values.

APIForbiddenError is a forbidden error response
*/
type IssueGetIssueReactionsForbidden struct {
	Message string
	URL     string
}

// IsSuccess returns true when this issue get issue reactions forbidden response has a 2xx status code
func (o *IssueGetIssueReactionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this issue get issue reactions forbidden response has a 3xx status code
func (o *IssueGetIssueReactionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue get issue reactions forbidden response has a 4xx status code
func (o *IssueGetIssueReactionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this issue get issue reactions forbidden response has a 5xx status code
func (o *IssueGetIssueReactionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this issue get issue reactions forbidden response a status code equal to that given
func (o *IssueGetIssueReactionsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the issue get issue reactions forbidden response
func (o *IssueGetIssueReactionsForbidden) Code() int {
	return 403
}

func (o *IssueGetIssueReactionsForbidden) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/issues/{index}/reactions][%d] issueGetIssueReactionsForbidden", 403)
}

func (o *IssueGetIssueReactionsForbidden) String() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/issues/{index}/reactions][%d] issueGetIssueReactionsForbidden", 403)
}

func (o *IssueGetIssueReactionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewIssueGetIssueReactionsNotFound creates a IssueGetIssueReactionsNotFound with default headers values
func NewIssueGetIssueReactionsNotFound() *IssueGetIssueReactionsNotFound {
	return &IssueGetIssueReactionsNotFound{}
}

/*
IssueGetIssueReactionsNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type IssueGetIssueReactionsNotFound struct {
}

// IsSuccess returns true when this issue get issue reactions not found response has a 2xx status code
func (o *IssueGetIssueReactionsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this issue get issue reactions not found response has a 3xx status code
func (o *IssueGetIssueReactionsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue get issue reactions not found response has a 4xx status code
func (o *IssueGetIssueReactionsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this issue get issue reactions not found response has a 5xx status code
func (o *IssueGetIssueReactionsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this issue get issue reactions not found response a status code equal to that given
func (o *IssueGetIssueReactionsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the issue get issue reactions not found response
func (o *IssueGetIssueReactionsNotFound) Code() int {
	return 404
}

func (o *IssueGetIssueReactionsNotFound) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/issues/{index}/reactions][%d] issueGetIssueReactionsNotFound", 404)
}

func (o *IssueGetIssueReactionsNotFound) String() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/issues/{index}/reactions][%d] issueGetIssueReactionsNotFound", 404)
}

func (o *IssueGetIssueReactionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
