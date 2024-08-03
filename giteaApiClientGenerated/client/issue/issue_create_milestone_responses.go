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

// IssueCreateMilestoneReader is a Reader for the IssueCreateMilestone structure.
type IssueCreateMilestoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IssueCreateMilestoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIssueCreateMilestoneCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewIssueCreateMilestoneNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /repos/{owner}/{repo}/milestones] issueCreateMilestone", response, response.Code())
	}
}

// NewIssueCreateMilestoneCreated creates a IssueCreateMilestoneCreated with default headers values
func NewIssueCreateMilestoneCreated() *IssueCreateMilestoneCreated {
	return &IssueCreateMilestoneCreated{}
}

/*
IssueCreateMilestoneCreated describes a response with status code 201, with default header values.

Milestone
*/
type IssueCreateMilestoneCreated struct {
	Payload *models.Milestone
}

// IsSuccess returns true when this issue create milestone created response has a 2xx status code
func (o *IssueCreateMilestoneCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this issue create milestone created response has a 3xx status code
func (o *IssueCreateMilestoneCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue create milestone created response has a 4xx status code
func (o *IssueCreateMilestoneCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this issue create milestone created response has a 5xx status code
func (o *IssueCreateMilestoneCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this issue create milestone created response a status code equal to that given
func (o *IssueCreateMilestoneCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the issue create milestone created response
func (o *IssueCreateMilestoneCreated) Code() int {
	return 201
}

func (o *IssueCreateMilestoneCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/milestones][%d] issueCreateMilestoneCreated %s", 201, payload)
}

func (o *IssueCreateMilestoneCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/milestones][%d] issueCreateMilestoneCreated %s", 201, payload)
}

func (o *IssueCreateMilestoneCreated) GetPayload() *models.Milestone {
	return o.Payload
}

func (o *IssueCreateMilestoneCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Milestone)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIssueCreateMilestoneNotFound creates a IssueCreateMilestoneNotFound with default headers values
func NewIssueCreateMilestoneNotFound() *IssueCreateMilestoneNotFound {
	return &IssueCreateMilestoneNotFound{}
}

/*
IssueCreateMilestoneNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type IssueCreateMilestoneNotFound struct {
}

// IsSuccess returns true when this issue create milestone not found response has a 2xx status code
func (o *IssueCreateMilestoneNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this issue create milestone not found response has a 3xx status code
func (o *IssueCreateMilestoneNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue create milestone not found response has a 4xx status code
func (o *IssueCreateMilestoneNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this issue create milestone not found response has a 5xx status code
func (o *IssueCreateMilestoneNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this issue create milestone not found response a status code equal to that given
func (o *IssueCreateMilestoneNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the issue create milestone not found response
func (o *IssueCreateMilestoneNotFound) Code() int {
	return 404
}

func (o *IssueCreateMilestoneNotFound) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/milestones][%d] issueCreateMilestoneNotFound", 404)
}

func (o *IssueCreateMilestoneNotFound) String() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/milestones][%d] issueCreateMilestoneNotFound", 404)
}

func (o *IssueCreateMilestoneNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
