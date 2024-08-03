// Code generated by go-swagger; DO NOT EDIT.

package repository

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

// RepoCreateReleaseAttachmentReader is a Reader for the RepoCreateReleaseAttachment structure.
type RepoCreateReleaseAttachmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoCreateReleaseAttachmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewRepoCreateReleaseAttachmentCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRepoCreateReleaseAttachmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRepoCreateReleaseAttachmentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /repos/{owner}/{repo}/releases/{id}/assets] repoCreateReleaseAttachment", response, response.Code())
	}
}

// NewRepoCreateReleaseAttachmentCreated creates a RepoCreateReleaseAttachmentCreated with default headers values
func NewRepoCreateReleaseAttachmentCreated() *RepoCreateReleaseAttachmentCreated {
	return &RepoCreateReleaseAttachmentCreated{}
}

/*
RepoCreateReleaseAttachmentCreated describes a response with status code 201, with default header values.

Attachment
*/
type RepoCreateReleaseAttachmentCreated struct {
	Payload *models.Attachment
}

// IsSuccess returns true when this repo create release attachment created response has a 2xx status code
func (o *RepoCreateReleaseAttachmentCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repo create release attachment created response has a 3xx status code
func (o *RepoCreateReleaseAttachmentCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo create release attachment created response has a 4xx status code
func (o *RepoCreateReleaseAttachmentCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this repo create release attachment created response has a 5xx status code
func (o *RepoCreateReleaseAttachmentCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this repo create release attachment created response a status code equal to that given
func (o *RepoCreateReleaseAttachmentCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the repo create release attachment created response
func (o *RepoCreateReleaseAttachmentCreated) Code() int {
	return 201
}

func (o *RepoCreateReleaseAttachmentCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/releases/{id}/assets][%d] repoCreateReleaseAttachmentCreated %s", 201, payload)
}

func (o *RepoCreateReleaseAttachmentCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/releases/{id}/assets][%d] repoCreateReleaseAttachmentCreated %s", 201, payload)
}

func (o *RepoCreateReleaseAttachmentCreated) GetPayload() *models.Attachment {
	return o.Payload
}

func (o *RepoCreateReleaseAttachmentCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Attachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepoCreateReleaseAttachmentBadRequest creates a RepoCreateReleaseAttachmentBadRequest with default headers values
func NewRepoCreateReleaseAttachmentBadRequest() *RepoCreateReleaseAttachmentBadRequest {
	return &RepoCreateReleaseAttachmentBadRequest{}
}

/*
RepoCreateReleaseAttachmentBadRequest describes a response with status code 400, with default header values.

APIError is error format response
*/
type RepoCreateReleaseAttachmentBadRequest struct {
	Message string
	URL     string
}

// IsSuccess returns true when this repo create release attachment bad request response has a 2xx status code
func (o *RepoCreateReleaseAttachmentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo create release attachment bad request response has a 3xx status code
func (o *RepoCreateReleaseAttachmentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo create release attachment bad request response has a 4xx status code
func (o *RepoCreateReleaseAttachmentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo create release attachment bad request response has a 5xx status code
func (o *RepoCreateReleaseAttachmentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this repo create release attachment bad request response a status code equal to that given
func (o *RepoCreateReleaseAttachmentBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the repo create release attachment bad request response
func (o *RepoCreateReleaseAttachmentBadRequest) Code() int {
	return 400
}

func (o *RepoCreateReleaseAttachmentBadRequest) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/releases/{id}/assets][%d] repoCreateReleaseAttachmentBadRequest", 400)
}

func (o *RepoCreateReleaseAttachmentBadRequest) String() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/releases/{id}/assets][%d] repoCreateReleaseAttachmentBadRequest", 400)
}

func (o *RepoCreateReleaseAttachmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRepoCreateReleaseAttachmentNotFound creates a RepoCreateReleaseAttachmentNotFound with default headers values
func NewRepoCreateReleaseAttachmentNotFound() *RepoCreateReleaseAttachmentNotFound {
	return &RepoCreateReleaseAttachmentNotFound{}
}

/*
RepoCreateReleaseAttachmentNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type RepoCreateReleaseAttachmentNotFound struct {
}

// IsSuccess returns true when this repo create release attachment not found response has a 2xx status code
func (o *RepoCreateReleaseAttachmentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo create release attachment not found response has a 3xx status code
func (o *RepoCreateReleaseAttachmentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo create release attachment not found response has a 4xx status code
func (o *RepoCreateReleaseAttachmentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo create release attachment not found response has a 5xx status code
func (o *RepoCreateReleaseAttachmentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repo create release attachment not found response a status code equal to that given
func (o *RepoCreateReleaseAttachmentNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the repo create release attachment not found response
func (o *RepoCreateReleaseAttachmentNotFound) Code() int {
	return 404
}

func (o *RepoCreateReleaseAttachmentNotFound) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/releases/{id}/assets][%d] repoCreateReleaseAttachmentNotFound", 404)
}

func (o *RepoCreateReleaseAttachmentNotFound) String() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/releases/{id}/assets][%d] repoCreateReleaseAttachmentNotFound", 404)
}

func (o *RepoCreateReleaseAttachmentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
