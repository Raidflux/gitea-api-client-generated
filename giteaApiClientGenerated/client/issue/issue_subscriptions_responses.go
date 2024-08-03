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

// IssueSubscriptionsReader is a Reader for the IssueSubscriptions structure.
type IssueSubscriptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IssueSubscriptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIssueSubscriptionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewIssueSubscriptionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /repos/{owner}/{repo}/issues/{index}/subscriptions] issueSubscriptions", response, response.Code())
	}
}

// NewIssueSubscriptionsOK creates a IssueSubscriptionsOK with default headers values
func NewIssueSubscriptionsOK() *IssueSubscriptionsOK {
	return &IssueSubscriptionsOK{}
}

/*
IssueSubscriptionsOK describes a response with status code 200, with default header values.

UserList
*/
type IssueSubscriptionsOK struct {
	Payload []*models.User
}

// IsSuccess returns true when this issue subscriptions o k response has a 2xx status code
func (o *IssueSubscriptionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this issue subscriptions o k response has a 3xx status code
func (o *IssueSubscriptionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue subscriptions o k response has a 4xx status code
func (o *IssueSubscriptionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this issue subscriptions o k response has a 5xx status code
func (o *IssueSubscriptionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this issue subscriptions o k response a status code equal to that given
func (o *IssueSubscriptionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the issue subscriptions o k response
func (o *IssueSubscriptionsOK) Code() int {
	return 200
}

func (o *IssueSubscriptionsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/issues/{index}/subscriptions][%d] issueSubscriptionsOK %s", 200, payload)
}

func (o *IssueSubscriptionsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/issues/{index}/subscriptions][%d] issueSubscriptionsOK %s", 200, payload)
}

func (o *IssueSubscriptionsOK) GetPayload() []*models.User {
	return o.Payload
}

func (o *IssueSubscriptionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIssueSubscriptionsNotFound creates a IssueSubscriptionsNotFound with default headers values
func NewIssueSubscriptionsNotFound() *IssueSubscriptionsNotFound {
	return &IssueSubscriptionsNotFound{}
}

/*
IssueSubscriptionsNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type IssueSubscriptionsNotFound struct {
}

// IsSuccess returns true when this issue subscriptions not found response has a 2xx status code
func (o *IssueSubscriptionsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this issue subscriptions not found response has a 3xx status code
func (o *IssueSubscriptionsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue subscriptions not found response has a 4xx status code
func (o *IssueSubscriptionsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this issue subscriptions not found response has a 5xx status code
func (o *IssueSubscriptionsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this issue subscriptions not found response a status code equal to that given
func (o *IssueSubscriptionsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the issue subscriptions not found response
func (o *IssueSubscriptionsNotFound) Code() int {
	return 404
}

func (o *IssueSubscriptionsNotFound) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/issues/{index}/subscriptions][%d] issueSubscriptionsNotFound", 404)
}

func (o *IssueSubscriptionsNotFound) String() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/issues/{index}/subscriptions][%d] issueSubscriptionsNotFound", 404)
}

func (o *IssueSubscriptionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
