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

	"github.com/Raidflux/gitea-api-client-generated/giteaApiClientGenerated/models"
)

// RepoEditReleaseAttachmentReader is a Reader for the RepoEditReleaseAttachment structure.
type RepoEditReleaseAttachmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoEditReleaseAttachmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewRepoEditReleaseAttachmentCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRepoEditReleaseAttachmentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /repos/{owner}/{repo}/releases/{id}/assets/{attachment_id}] repoEditReleaseAttachment", response, response.Code())
	}
}

// NewRepoEditReleaseAttachmentCreated creates a RepoEditReleaseAttachmentCreated with default headers values
func NewRepoEditReleaseAttachmentCreated() *RepoEditReleaseAttachmentCreated {
	return &RepoEditReleaseAttachmentCreated{}
}

/*
RepoEditReleaseAttachmentCreated describes a response with status code 201, with default header values.

Attachment
*/
type RepoEditReleaseAttachmentCreated struct {
	Payload *models.Attachment
}

// IsSuccess returns true when this repo edit release attachment created response has a 2xx status code
func (o *RepoEditReleaseAttachmentCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repo edit release attachment created response has a 3xx status code
func (o *RepoEditReleaseAttachmentCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo edit release attachment created response has a 4xx status code
func (o *RepoEditReleaseAttachmentCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this repo edit release attachment created response has a 5xx status code
func (o *RepoEditReleaseAttachmentCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this repo edit release attachment created response a status code equal to that given
func (o *RepoEditReleaseAttachmentCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the repo edit release attachment created response
func (o *RepoEditReleaseAttachmentCreated) Code() int {
	return 201
}

func (o *RepoEditReleaseAttachmentCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}/releases/{id}/assets/{attachment_id}][%d] repoEditReleaseAttachmentCreated %s", 201, payload)
}

func (o *RepoEditReleaseAttachmentCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}/releases/{id}/assets/{attachment_id}][%d] repoEditReleaseAttachmentCreated %s", 201, payload)
}

func (o *RepoEditReleaseAttachmentCreated) GetPayload() *models.Attachment {
	return o.Payload
}

func (o *RepoEditReleaseAttachmentCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Attachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepoEditReleaseAttachmentNotFound creates a RepoEditReleaseAttachmentNotFound with default headers values
func NewRepoEditReleaseAttachmentNotFound() *RepoEditReleaseAttachmentNotFound {
	return &RepoEditReleaseAttachmentNotFound{}
}

/*
RepoEditReleaseAttachmentNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type RepoEditReleaseAttachmentNotFound struct {
}

// IsSuccess returns true when this repo edit release attachment not found response has a 2xx status code
func (o *RepoEditReleaseAttachmentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo edit release attachment not found response has a 3xx status code
func (o *RepoEditReleaseAttachmentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo edit release attachment not found response has a 4xx status code
func (o *RepoEditReleaseAttachmentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo edit release attachment not found response has a 5xx status code
func (o *RepoEditReleaseAttachmentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repo edit release attachment not found response a status code equal to that given
func (o *RepoEditReleaseAttachmentNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the repo edit release attachment not found response
func (o *RepoEditReleaseAttachmentNotFound) Code() int {
	return 404
}

func (o *RepoEditReleaseAttachmentNotFound) Error() string {
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}/releases/{id}/assets/{attachment_id}][%d] repoEditReleaseAttachmentNotFound", 404)
}

func (o *RepoEditReleaseAttachmentNotFound) String() string {
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}/releases/{id}/assets/{attachment_id}][%d] repoEditReleaseAttachmentNotFound", 404)
}

func (o *RepoEditReleaseAttachmentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
