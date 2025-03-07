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

// IssueListIssueCommentAttachmentsReader is a Reader for the IssueListIssueCommentAttachments structure.
type IssueListIssueCommentAttachmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IssueListIssueCommentAttachmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIssueListIssueCommentAttachmentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewIssueListIssueCommentAttachmentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /repos/{owner}/{repo}/issues/comments/{id}/assets] issueListIssueCommentAttachments", response, response.Code())
	}
}

// NewIssueListIssueCommentAttachmentsOK creates a IssueListIssueCommentAttachmentsOK with default headers values
func NewIssueListIssueCommentAttachmentsOK() *IssueListIssueCommentAttachmentsOK {
	return &IssueListIssueCommentAttachmentsOK{}
}

/*
IssueListIssueCommentAttachmentsOK describes a response with status code 200, with default header values.

AttachmentList
*/
type IssueListIssueCommentAttachmentsOK struct {
	Payload []*models.Attachment
}

// IsSuccess returns true when this issue list issue comment attachments o k response has a 2xx status code
func (o *IssueListIssueCommentAttachmentsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this issue list issue comment attachments o k response has a 3xx status code
func (o *IssueListIssueCommentAttachmentsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue list issue comment attachments o k response has a 4xx status code
func (o *IssueListIssueCommentAttachmentsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this issue list issue comment attachments o k response has a 5xx status code
func (o *IssueListIssueCommentAttachmentsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this issue list issue comment attachments o k response a status code equal to that given
func (o *IssueListIssueCommentAttachmentsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the issue list issue comment attachments o k response
func (o *IssueListIssueCommentAttachmentsOK) Code() int {
	return 200
}

func (o *IssueListIssueCommentAttachmentsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/issues/comments/{id}/assets][%d] issueListIssueCommentAttachmentsOK %s", 200, payload)
}

func (o *IssueListIssueCommentAttachmentsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/issues/comments/{id}/assets][%d] issueListIssueCommentAttachmentsOK %s", 200, payload)
}

func (o *IssueListIssueCommentAttachmentsOK) GetPayload() []*models.Attachment {
	return o.Payload
}

func (o *IssueListIssueCommentAttachmentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIssueListIssueCommentAttachmentsNotFound creates a IssueListIssueCommentAttachmentsNotFound with default headers values
func NewIssueListIssueCommentAttachmentsNotFound() *IssueListIssueCommentAttachmentsNotFound {
	return &IssueListIssueCommentAttachmentsNotFound{}
}

/*
IssueListIssueCommentAttachmentsNotFound describes a response with status code 404, with default header values.

APIError is error format response
*/
type IssueListIssueCommentAttachmentsNotFound struct {
	Message string
	URL     string
}

// IsSuccess returns true when this issue list issue comment attachments not found response has a 2xx status code
func (o *IssueListIssueCommentAttachmentsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this issue list issue comment attachments not found response has a 3xx status code
func (o *IssueListIssueCommentAttachmentsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue list issue comment attachments not found response has a 4xx status code
func (o *IssueListIssueCommentAttachmentsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this issue list issue comment attachments not found response has a 5xx status code
func (o *IssueListIssueCommentAttachmentsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this issue list issue comment attachments not found response a status code equal to that given
func (o *IssueListIssueCommentAttachmentsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the issue list issue comment attachments not found response
func (o *IssueListIssueCommentAttachmentsNotFound) Code() int {
	return 404
}

func (o *IssueListIssueCommentAttachmentsNotFound) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/issues/comments/{id}/assets][%d] issueListIssueCommentAttachmentsNotFound", 404)
}

func (o *IssueListIssueCommentAttachmentsNotFound) String() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/issues/comments/{id}/assets][%d] issueListIssueCommentAttachmentsNotFound", 404)
}

func (o *IssueListIssueCommentAttachmentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
