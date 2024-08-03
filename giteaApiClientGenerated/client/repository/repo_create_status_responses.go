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

// RepoCreateStatusReader is a Reader for the RepoCreateStatus structure.
type RepoCreateStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoCreateStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewRepoCreateStatusCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRepoCreateStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRepoCreateStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /repos/{owner}/{repo}/statuses/{sha}] repoCreateStatus", response, response.Code())
	}
}

// NewRepoCreateStatusCreated creates a RepoCreateStatusCreated with default headers values
func NewRepoCreateStatusCreated() *RepoCreateStatusCreated {
	return &RepoCreateStatusCreated{}
}

/*
RepoCreateStatusCreated describes a response with status code 201, with default header values.

CommitStatus
*/
type RepoCreateStatusCreated struct {
	Payload *models.CommitStatus
}

// IsSuccess returns true when this repo create status created response has a 2xx status code
func (o *RepoCreateStatusCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repo create status created response has a 3xx status code
func (o *RepoCreateStatusCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo create status created response has a 4xx status code
func (o *RepoCreateStatusCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this repo create status created response has a 5xx status code
func (o *RepoCreateStatusCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this repo create status created response a status code equal to that given
func (o *RepoCreateStatusCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the repo create status created response
func (o *RepoCreateStatusCreated) Code() int {
	return 201
}

func (o *RepoCreateStatusCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/statuses/{sha}][%d] repoCreateStatusCreated %s", 201, payload)
}

func (o *RepoCreateStatusCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/statuses/{sha}][%d] repoCreateStatusCreated %s", 201, payload)
}

func (o *RepoCreateStatusCreated) GetPayload() *models.CommitStatus {
	return o.Payload
}

func (o *RepoCreateStatusCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommitStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepoCreateStatusBadRequest creates a RepoCreateStatusBadRequest with default headers values
func NewRepoCreateStatusBadRequest() *RepoCreateStatusBadRequest {
	return &RepoCreateStatusBadRequest{}
}

/*
RepoCreateStatusBadRequest describes a response with status code 400, with default header values.

APIError is error format response
*/
type RepoCreateStatusBadRequest struct {
	Message string
	URL     string
}

// IsSuccess returns true when this repo create status bad request response has a 2xx status code
func (o *RepoCreateStatusBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo create status bad request response has a 3xx status code
func (o *RepoCreateStatusBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo create status bad request response has a 4xx status code
func (o *RepoCreateStatusBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo create status bad request response has a 5xx status code
func (o *RepoCreateStatusBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this repo create status bad request response a status code equal to that given
func (o *RepoCreateStatusBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the repo create status bad request response
func (o *RepoCreateStatusBadRequest) Code() int {
	return 400
}

func (o *RepoCreateStatusBadRequest) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/statuses/{sha}][%d] repoCreateStatusBadRequest", 400)
}

func (o *RepoCreateStatusBadRequest) String() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/statuses/{sha}][%d] repoCreateStatusBadRequest", 400)
}

func (o *RepoCreateStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRepoCreateStatusNotFound creates a RepoCreateStatusNotFound with default headers values
func NewRepoCreateStatusNotFound() *RepoCreateStatusNotFound {
	return &RepoCreateStatusNotFound{}
}

/*
RepoCreateStatusNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type RepoCreateStatusNotFound struct {
}

// IsSuccess returns true when this repo create status not found response has a 2xx status code
func (o *RepoCreateStatusNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo create status not found response has a 3xx status code
func (o *RepoCreateStatusNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo create status not found response has a 4xx status code
func (o *RepoCreateStatusNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo create status not found response has a 5xx status code
func (o *RepoCreateStatusNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repo create status not found response a status code equal to that given
func (o *RepoCreateStatusNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the repo create status not found response
func (o *RepoCreateStatusNotFound) Code() int {
	return 404
}

func (o *RepoCreateStatusNotFound) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/statuses/{sha}][%d] repoCreateStatusNotFound", 404)
}

func (o *RepoCreateStatusNotFound) String() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/statuses/{sha}][%d] repoCreateStatusNotFound", 404)
}

func (o *RepoCreateStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
