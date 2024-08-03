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

// IssueCreateIssueBlockingReader is a Reader for the IssueCreateIssueBlocking structure.
type IssueCreateIssueBlockingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IssueCreateIssueBlockingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIssueCreateIssueBlockingCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewIssueCreateIssueBlockingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /repos/{owner}/{repo}/issues/{index}/blocks] issueCreateIssueBlocking", response, response.Code())
	}
}

// NewIssueCreateIssueBlockingCreated creates a IssueCreateIssueBlockingCreated with default headers values
func NewIssueCreateIssueBlockingCreated() *IssueCreateIssueBlockingCreated {
	return &IssueCreateIssueBlockingCreated{}
}

/*
IssueCreateIssueBlockingCreated describes a response with status code 201, with default header values.

Issue
*/
type IssueCreateIssueBlockingCreated struct {
	Payload *models.Issue
}

// IsSuccess returns true when this issue create issue blocking created response has a 2xx status code
func (o *IssueCreateIssueBlockingCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this issue create issue blocking created response has a 3xx status code
func (o *IssueCreateIssueBlockingCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue create issue blocking created response has a 4xx status code
func (o *IssueCreateIssueBlockingCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this issue create issue blocking created response has a 5xx status code
func (o *IssueCreateIssueBlockingCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this issue create issue blocking created response a status code equal to that given
func (o *IssueCreateIssueBlockingCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the issue create issue blocking created response
func (o *IssueCreateIssueBlockingCreated) Code() int {
	return 201
}

func (o *IssueCreateIssueBlockingCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/issues/{index}/blocks][%d] issueCreateIssueBlockingCreated %s", 201, payload)
}

func (o *IssueCreateIssueBlockingCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/issues/{index}/blocks][%d] issueCreateIssueBlockingCreated %s", 201, payload)
}

func (o *IssueCreateIssueBlockingCreated) GetPayload() *models.Issue {
	return o.Payload
}

func (o *IssueCreateIssueBlockingCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Issue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIssueCreateIssueBlockingNotFound creates a IssueCreateIssueBlockingNotFound with default headers values
func NewIssueCreateIssueBlockingNotFound() *IssueCreateIssueBlockingNotFound {
	return &IssueCreateIssueBlockingNotFound{}
}

/*
IssueCreateIssueBlockingNotFound describes a response with status code 404, with default header values.

the issue does not exist
*/
type IssueCreateIssueBlockingNotFound struct {
}

// IsSuccess returns true when this issue create issue blocking not found response has a 2xx status code
func (o *IssueCreateIssueBlockingNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this issue create issue blocking not found response has a 3xx status code
func (o *IssueCreateIssueBlockingNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue create issue blocking not found response has a 4xx status code
func (o *IssueCreateIssueBlockingNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this issue create issue blocking not found response has a 5xx status code
func (o *IssueCreateIssueBlockingNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this issue create issue blocking not found response a status code equal to that given
func (o *IssueCreateIssueBlockingNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the issue create issue blocking not found response
func (o *IssueCreateIssueBlockingNotFound) Code() int {
	return 404
}

func (o *IssueCreateIssueBlockingNotFound) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/issues/{index}/blocks][%d] issueCreateIssueBlockingNotFound", 404)
}

func (o *IssueCreateIssueBlockingNotFound) String() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/issues/{index}/blocks][%d] issueCreateIssueBlockingNotFound", 404)
}

func (o *IssueCreateIssueBlockingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
