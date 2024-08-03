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

// RepoCreatePullRequestReader is a Reader for the RepoCreatePullRequest structure.
type RepoCreatePullRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoCreatePullRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewRepoCreatePullRequestCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewRepoCreatePullRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRepoCreatePullRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewRepoCreatePullRequestConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewRepoCreatePullRequestUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 423:
		result := NewRepoCreatePullRequestLocked()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /repos/{owner}/{repo}/pulls] repoCreatePullRequest", response, response.Code())
	}
}

// NewRepoCreatePullRequestCreated creates a RepoCreatePullRequestCreated with default headers values
func NewRepoCreatePullRequestCreated() *RepoCreatePullRequestCreated {
	return &RepoCreatePullRequestCreated{}
}

/*
RepoCreatePullRequestCreated describes a response with status code 201, with default header values.

PullRequest
*/
type RepoCreatePullRequestCreated struct {
	Payload *models.PullRequest
}

// IsSuccess returns true when this repo create pull request created response has a 2xx status code
func (o *RepoCreatePullRequestCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repo create pull request created response has a 3xx status code
func (o *RepoCreatePullRequestCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo create pull request created response has a 4xx status code
func (o *RepoCreatePullRequestCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this repo create pull request created response has a 5xx status code
func (o *RepoCreatePullRequestCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this repo create pull request created response a status code equal to that given
func (o *RepoCreatePullRequestCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the repo create pull request created response
func (o *RepoCreatePullRequestCreated) Code() int {
	return 201
}

func (o *RepoCreatePullRequestCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/pulls][%d] repoCreatePullRequestCreated %s", 201, payload)
}

func (o *RepoCreatePullRequestCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/pulls][%d] repoCreatePullRequestCreated %s", 201, payload)
}

func (o *RepoCreatePullRequestCreated) GetPayload() *models.PullRequest {
	return o.Payload
}

func (o *RepoCreatePullRequestCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PullRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepoCreatePullRequestForbidden creates a RepoCreatePullRequestForbidden with default headers values
func NewRepoCreatePullRequestForbidden() *RepoCreatePullRequestForbidden {
	return &RepoCreatePullRequestForbidden{}
}

/*
RepoCreatePullRequestForbidden describes a response with status code 403, with default header values.

APIForbiddenError is a forbidden error response
*/
type RepoCreatePullRequestForbidden struct {
	Message string
	URL     string
}

// IsSuccess returns true when this repo create pull request forbidden response has a 2xx status code
func (o *RepoCreatePullRequestForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo create pull request forbidden response has a 3xx status code
func (o *RepoCreatePullRequestForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo create pull request forbidden response has a 4xx status code
func (o *RepoCreatePullRequestForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo create pull request forbidden response has a 5xx status code
func (o *RepoCreatePullRequestForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this repo create pull request forbidden response a status code equal to that given
func (o *RepoCreatePullRequestForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the repo create pull request forbidden response
func (o *RepoCreatePullRequestForbidden) Code() int {
	return 403
}

func (o *RepoCreatePullRequestForbidden) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/pulls][%d] repoCreatePullRequestForbidden", 403)
}

func (o *RepoCreatePullRequestForbidden) String() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/pulls][%d] repoCreatePullRequestForbidden", 403)
}

func (o *RepoCreatePullRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRepoCreatePullRequestNotFound creates a RepoCreatePullRequestNotFound with default headers values
func NewRepoCreatePullRequestNotFound() *RepoCreatePullRequestNotFound {
	return &RepoCreatePullRequestNotFound{}
}

/*
RepoCreatePullRequestNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type RepoCreatePullRequestNotFound struct {
}

// IsSuccess returns true when this repo create pull request not found response has a 2xx status code
func (o *RepoCreatePullRequestNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo create pull request not found response has a 3xx status code
func (o *RepoCreatePullRequestNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo create pull request not found response has a 4xx status code
func (o *RepoCreatePullRequestNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo create pull request not found response has a 5xx status code
func (o *RepoCreatePullRequestNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repo create pull request not found response a status code equal to that given
func (o *RepoCreatePullRequestNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the repo create pull request not found response
func (o *RepoCreatePullRequestNotFound) Code() int {
	return 404
}

func (o *RepoCreatePullRequestNotFound) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/pulls][%d] repoCreatePullRequestNotFound", 404)
}

func (o *RepoCreatePullRequestNotFound) String() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/pulls][%d] repoCreatePullRequestNotFound", 404)
}

func (o *RepoCreatePullRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRepoCreatePullRequestConflict creates a RepoCreatePullRequestConflict with default headers values
func NewRepoCreatePullRequestConflict() *RepoCreatePullRequestConflict {
	return &RepoCreatePullRequestConflict{}
}

/*
RepoCreatePullRequestConflict describes a response with status code 409, with default header values.

APIError is error format response
*/
type RepoCreatePullRequestConflict struct {
	Message string
	URL     string
}

// IsSuccess returns true when this repo create pull request conflict response has a 2xx status code
func (o *RepoCreatePullRequestConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo create pull request conflict response has a 3xx status code
func (o *RepoCreatePullRequestConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo create pull request conflict response has a 4xx status code
func (o *RepoCreatePullRequestConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo create pull request conflict response has a 5xx status code
func (o *RepoCreatePullRequestConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this repo create pull request conflict response a status code equal to that given
func (o *RepoCreatePullRequestConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the repo create pull request conflict response
func (o *RepoCreatePullRequestConflict) Code() int {
	return 409
}

func (o *RepoCreatePullRequestConflict) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/pulls][%d] repoCreatePullRequestConflict", 409)
}

func (o *RepoCreatePullRequestConflict) String() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/pulls][%d] repoCreatePullRequestConflict", 409)
}

func (o *RepoCreatePullRequestConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRepoCreatePullRequestUnprocessableEntity creates a RepoCreatePullRequestUnprocessableEntity with default headers values
func NewRepoCreatePullRequestUnprocessableEntity() *RepoCreatePullRequestUnprocessableEntity {
	return &RepoCreatePullRequestUnprocessableEntity{}
}

/*
RepoCreatePullRequestUnprocessableEntity describes a response with status code 422, with default header values.

APIValidationError is error format response related to input validation
*/
type RepoCreatePullRequestUnprocessableEntity struct {
	Message string
	URL     string
}

// IsSuccess returns true when this repo create pull request unprocessable entity response has a 2xx status code
func (o *RepoCreatePullRequestUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo create pull request unprocessable entity response has a 3xx status code
func (o *RepoCreatePullRequestUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo create pull request unprocessable entity response has a 4xx status code
func (o *RepoCreatePullRequestUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo create pull request unprocessable entity response has a 5xx status code
func (o *RepoCreatePullRequestUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this repo create pull request unprocessable entity response a status code equal to that given
func (o *RepoCreatePullRequestUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the repo create pull request unprocessable entity response
func (o *RepoCreatePullRequestUnprocessableEntity) Code() int {
	return 422
}

func (o *RepoCreatePullRequestUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/pulls][%d] repoCreatePullRequestUnprocessableEntity", 422)
}

func (o *RepoCreatePullRequestUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/pulls][%d] repoCreatePullRequestUnprocessableEntity", 422)
}

func (o *RepoCreatePullRequestUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRepoCreatePullRequestLocked creates a RepoCreatePullRequestLocked with default headers values
func NewRepoCreatePullRequestLocked() *RepoCreatePullRequestLocked {
	return &RepoCreatePullRequestLocked{}
}

/*
RepoCreatePullRequestLocked describes a response with status code 423, with default header values.

APIRepoArchivedError is an error that is raised when an archived repo should be modified
*/
type RepoCreatePullRequestLocked struct {
	Message string
	URL     string
}

// IsSuccess returns true when this repo create pull request locked response has a 2xx status code
func (o *RepoCreatePullRequestLocked) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo create pull request locked response has a 3xx status code
func (o *RepoCreatePullRequestLocked) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo create pull request locked response has a 4xx status code
func (o *RepoCreatePullRequestLocked) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo create pull request locked response has a 5xx status code
func (o *RepoCreatePullRequestLocked) IsServerError() bool {
	return false
}

// IsCode returns true when this repo create pull request locked response a status code equal to that given
func (o *RepoCreatePullRequestLocked) IsCode(code int) bool {
	return code == 423
}

// Code gets the status code for the repo create pull request locked response
func (o *RepoCreatePullRequestLocked) Code() int {
	return 423
}

func (o *RepoCreatePullRequestLocked) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/pulls][%d] repoCreatePullRequestLocked", 423)
}

func (o *RepoCreatePullRequestLocked) String() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/pulls][%d] repoCreatePullRequestLocked", 423)
}

func (o *RepoCreatePullRequestLocked) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
