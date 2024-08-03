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

// IssueGetCommentsAndTimelineReader is a Reader for the IssueGetCommentsAndTimeline structure.
type IssueGetCommentsAndTimelineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IssueGetCommentsAndTimelineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIssueGetCommentsAndTimelineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewIssueGetCommentsAndTimelineNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /repos/{owner}/{repo}/issues/{index}/timeline] issueGetCommentsAndTimeline", response, response.Code())
	}
}

// NewIssueGetCommentsAndTimelineOK creates a IssueGetCommentsAndTimelineOK with default headers values
func NewIssueGetCommentsAndTimelineOK() *IssueGetCommentsAndTimelineOK {
	return &IssueGetCommentsAndTimelineOK{}
}

/*
IssueGetCommentsAndTimelineOK describes a response with status code 200, with default header values.

TimelineList
*/
type IssueGetCommentsAndTimelineOK struct {
	Payload []*models.TimelineComment
}

// IsSuccess returns true when this issue get comments and timeline o k response has a 2xx status code
func (o *IssueGetCommentsAndTimelineOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this issue get comments and timeline o k response has a 3xx status code
func (o *IssueGetCommentsAndTimelineOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue get comments and timeline o k response has a 4xx status code
func (o *IssueGetCommentsAndTimelineOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this issue get comments and timeline o k response has a 5xx status code
func (o *IssueGetCommentsAndTimelineOK) IsServerError() bool {
	return false
}

// IsCode returns true when this issue get comments and timeline o k response a status code equal to that given
func (o *IssueGetCommentsAndTimelineOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the issue get comments and timeline o k response
func (o *IssueGetCommentsAndTimelineOK) Code() int {
	return 200
}

func (o *IssueGetCommentsAndTimelineOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/issues/{index}/timeline][%d] issueGetCommentsAndTimelineOK %s", 200, payload)
}

func (o *IssueGetCommentsAndTimelineOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/issues/{index}/timeline][%d] issueGetCommentsAndTimelineOK %s", 200, payload)
}

func (o *IssueGetCommentsAndTimelineOK) GetPayload() []*models.TimelineComment {
	return o.Payload
}

func (o *IssueGetCommentsAndTimelineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIssueGetCommentsAndTimelineNotFound creates a IssueGetCommentsAndTimelineNotFound with default headers values
func NewIssueGetCommentsAndTimelineNotFound() *IssueGetCommentsAndTimelineNotFound {
	return &IssueGetCommentsAndTimelineNotFound{}
}

/*
IssueGetCommentsAndTimelineNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type IssueGetCommentsAndTimelineNotFound struct {
}

// IsSuccess returns true when this issue get comments and timeline not found response has a 2xx status code
func (o *IssueGetCommentsAndTimelineNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this issue get comments and timeline not found response has a 3xx status code
func (o *IssueGetCommentsAndTimelineNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue get comments and timeline not found response has a 4xx status code
func (o *IssueGetCommentsAndTimelineNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this issue get comments and timeline not found response has a 5xx status code
func (o *IssueGetCommentsAndTimelineNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this issue get comments and timeline not found response a status code equal to that given
func (o *IssueGetCommentsAndTimelineNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the issue get comments and timeline not found response
func (o *IssueGetCommentsAndTimelineNotFound) Code() int {
	return 404
}

func (o *IssueGetCommentsAndTimelineNotFound) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/issues/{index}/timeline][%d] issueGetCommentsAndTimelineNotFound", 404)
}

func (o *IssueGetCommentsAndTimelineNotFound) String() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/issues/{index}/timeline][%d] issueGetCommentsAndTimelineNotFound", 404)
}

func (o *IssueGetCommentsAndTimelineNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
