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

// RepoListStatusesReader is a Reader for the RepoListStatuses structure.
type RepoListStatusesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoListStatusesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoListStatusesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRepoListStatusesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRepoListStatusesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /repos/{owner}/{repo}/statuses/{sha}] repoListStatuses", response, response.Code())
	}
}

// NewRepoListStatusesOK creates a RepoListStatusesOK with default headers values
func NewRepoListStatusesOK() *RepoListStatusesOK {
	return &RepoListStatusesOK{}
}

/*
RepoListStatusesOK describes a response with status code 200, with default header values.

CommitStatusList
*/
type RepoListStatusesOK struct {
	Payload []*models.CommitStatus
}

// IsSuccess returns true when this repo list statuses o k response has a 2xx status code
func (o *RepoListStatusesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repo list statuses o k response has a 3xx status code
func (o *RepoListStatusesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo list statuses o k response has a 4xx status code
func (o *RepoListStatusesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this repo list statuses o k response has a 5xx status code
func (o *RepoListStatusesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this repo list statuses o k response a status code equal to that given
func (o *RepoListStatusesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the repo list statuses o k response
func (o *RepoListStatusesOK) Code() int {
	return 200
}

func (o *RepoListStatusesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/statuses/{sha}][%d] repoListStatusesOK %s", 200, payload)
}

func (o *RepoListStatusesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/statuses/{sha}][%d] repoListStatusesOK %s", 200, payload)
}

func (o *RepoListStatusesOK) GetPayload() []*models.CommitStatus {
	return o.Payload
}

func (o *RepoListStatusesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepoListStatusesBadRequest creates a RepoListStatusesBadRequest with default headers values
func NewRepoListStatusesBadRequest() *RepoListStatusesBadRequest {
	return &RepoListStatusesBadRequest{}
}

/*
RepoListStatusesBadRequest describes a response with status code 400, with default header values.

APIError is error format response
*/
type RepoListStatusesBadRequest struct {
	Message string
	URL     string
}

// IsSuccess returns true when this repo list statuses bad request response has a 2xx status code
func (o *RepoListStatusesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo list statuses bad request response has a 3xx status code
func (o *RepoListStatusesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo list statuses bad request response has a 4xx status code
func (o *RepoListStatusesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo list statuses bad request response has a 5xx status code
func (o *RepoListStatusesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this repo list statuses bad request response a status code equal to that given
func (o *RepoListStatusesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the repo list statuses bad request response
func (o *RepoListStatusesBadRequest) Code() int {
	return 400
}

func (o *RepoListStatusesBadRequest) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/statuses/{sha}][%d] repoListStatusesBadRequest", 400)
}

func (o *RepoListStatusesBadRequest) String() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/statuses/{sha}][%d] repoListStatusesBadRequest", 400)
}

func (o *RepoListStatusesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRepoListStatusesNotFound creates a RepoListStatusesNotFound with default headers values
func NewRepoListStatusesNotFound() *RepoListStatusesNotFound {
	return &RepoListStatusesNotFound{}
}

/*
RepoListStatusesNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type RepoListStatusesNotFound struct {
}

// IsSuccess returns true when this repo list statuses not found response has a 2xx status code
func (o *RepoListStatusesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo list statuses not found response has a 3xx status code
func (o *RepoListStatusesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo list statuses not found response has a 4xx status code
func (o *RepoListStatusesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo list statuses not found response has a 5xx status code
func (o *RepoListStatusesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repo list statuses not found response a status code equal to that given
func (o *RepoListStatusesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the repo list statuses not found response
func (o *RepoListStatusesNotFound) Code() int {
	return 404
}

func (o *RepoListStatusesNotFound) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/statuses/{sha}][%d] repoListStatusesNotFound", 404)
}

func (o *RepoListStatusesNotFound) String() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/statuses/{sha}][%d] repoListStatusesNotFound", 404)
}

func (o *RepoListStatusesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
