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

// RepoCreatePullReviewRequestsReader is a Reader for the RepoCreatePullReviewRequests structure.
type RepoCreatePullReviewRequestsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoCreatePullReviewRequestsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewRepoCreatePullReviewRequestsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRepoCreatePullReviewRequestsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewRepoCreatePullReviewRequestsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /repos/{owner}/{repo}/pulls/{index}/requested_reviewers] repoCreatePullReviewRequests", response, response.Code())
	}
}

// NewRepoCreatePullReviewRequestsCreated creates a RepoCreatePullReviewRequestsCreated with default headers values
func NewRepoCreatePullReviewRequestsCreated() *RepoCreatePullReviewRequestsCreated {
	return &RepoCreatePullReviewRequestsCreated{}
}

/*
RepoCreatePullReviewRequestsCreated describes a response with status code 201, with default header values.

PullReviewList
*/
type RepoCreatePullReviewRequestsCreated struct {
	Payload []*models.PullReview
}

// IsSuccess returns true when this repo create pull review requests created response has a 2xx status code
func (o *RepoCreatePullReviewRequestsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repo create pull review requests created response has a 3xx status code
func (o *RepoCreatePullReviewRequestsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo create pull review requests created response has a 4xx status code
func (o *RepoCreatePullReviewRequestsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this repo create pull review requests created response has a 5xx status code
func (o *RepoCreatePullReviewRequestsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this repo create pull review requests created response a status code equal to that given
func (o *RepoCreatePullReviewRequestsCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the repo create pull review requests created response
func (o *RepoCreatePullReviewRequestsCreated) Code() int {
	return 201
}

func (o *RepoCreatePullReviewRequestsCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/pulls/{index}/requested_reviewers][%d] repoCreatePullReviewRequestsCreated %s", 201, payload)
}

func (o *RepoCreatePullReviewRequestsCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/pulls/{index}/requested_reviewers][%d] repoCreatePullReviewRequestsCreated %s", 201, payload)
}

func (o *RepoCreatePullReviewRequestsCreated) GetPayload() []*models.PullReview {
	return o.Payload
}

func (o *RepoCreatePullReviewRequestsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepoCreatePullReviewRequestsNotFound creates a RepoCreatePullReviewRequestsNotFound with default headers values
func NewRepoCreatePullReviewRequestsNotFound() *RepoCreatePullReviewRequestsNotFound {
	return &RepoCreatePullReviewRequestsNotFound{}
}

/*
RepoCreatePullReviewRequestsNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type RepoCreatePullReviewRequestsNotFound struct {
}

// IsSuccess returns true when this repo create pull review requests not found response has a 2xx status code
func (o *RepoCreatePullReviewRequestsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo create pull review requests not found response has a 3xx status code
func (o *RepoCreatePullReviewRequestsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo create pull review requests not found response has a 4xx status code
func (o *RepoCreatePullReviewRequestsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo create pull review requests not found response has a 5xx status code
func (o *RepoCreatePullReviewRequestsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repo create pull review requests not found response a status code equal to that given
func (o *RepoCreatePullReviewRequestsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the repo create pull review requests not found response
func (o *RepoCreatePullReviewRequestsNotFound) Code() int {
	return 404
}

func (o *RepoCreatePullReviewRequestsNotFound) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/pulls/{index}/requested_reviewers][%d] repoCreatePullReviewRequestsNotFound", 404)
}

func (o *RepoCreatePullReviewRequestsNotFound) String() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/pulls/{index}/requested_reviewers][%d] repoCreatePullReviewRequestsNotFound", 404)
}

func (o *RepoCreatePullReviewRequestsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRepoCreatePullReviewRequestsUnprocessableEntity creates a RepoCreatePullReviewRequestsUnprocessableEntity with default headers values
func NewRepoCreatePullReviewRequestsUnprocessableEntity() *RepoCreatePullReviewRequestsUnprocessableEntity {
	return &RepoCreatePullReviewRequestsUnprocessableEntity{}
}

/*
RepoCreatePullReviewRequestsUnprocessableEntity describes a response with status code 422, with default header values.

APIValidationError is error format response related to input validation
*/
type RepoCreatePullReviewRequestsUnprocessableEntity struct {
	Message string
	URL     string
}

// IsSuccess returns true when this repo create pull review requests unprocessable entity response has a 2xx status code
func (o *RepoCreatePullReviewRequestsUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo create pull review requests unprocessable entity response has a 3xx status code
func (o *RepoCreatePullReviewRequestsUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo create pull review requests unprocessable entity response has a 4xx status code
func (o *RepoCreatePullReviewRequestsUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo create pull review requests unprocessable entity response has a 5xx status code
func (o *RepoCreatePullReviewRequestsUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this repo create pull review requests unprocessable entity response a status code equal to that given
func (o *RepoCreatePullReviewRequestsUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the repo create pull review requests unprocessable entity response
func (o *RepoCreatePullReviewRequestsUnprocessableEntity) Code() int {
	return 422
}

func (o *RepoCreatePullReviewRequestsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/pulls/{index}/requested_reviewers][%d] repoCreatePullReviewRequestsUnprocessableEntity", 422)
}

func (o *RepoCreatePullReviewRequestsUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/pulls/{index}/requested_reviewers][%d] repoCreatePullReviewRequestsUnprocessableEntity", 422)
}

func (o *RepoCreatePullReviewRequestsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
