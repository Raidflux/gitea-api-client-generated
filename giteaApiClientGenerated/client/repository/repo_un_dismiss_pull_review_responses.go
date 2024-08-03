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

// RepoUnDismissPullReviewReader is a Reader for the RepoUnDismissPullReview structure.
type RepoUnDismissPullReviewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoUnDismissPullReviewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoUnDismissPullReviewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewRepoUnDismissPullReviewForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRepoUnDismissPullReviewNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewRepoUnDismissPullReviewUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /repos/{owner}/{repo}/pulls/{index}/reviews/{id}/undismissals] repoUnDismissPullReview", response, response.Code())
	}
}

// NewRepoUnDismissPullReviewOK creates a RepoUnDismissPullReviewOK with default headers values
func NewRepoUnDismissPullReviewOK() *RepoUnDismissPullReviewOK {
	return &RepoUnDismissPullReviewOK{}
}

/*
RepoUnDismissPullReviewOK describes a response with status code 200, with default header values.

PullReview
*/
type RepoUnDismissPullReviewOK struct {
	Payload *models.PullReview
}

// IsSuccess returns true when this repo un dismiss pull review o k response has a 2xx status code
func (o *RepoUnDismissPullReviewOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repo un dismiss pull review o k response has a 3xx status code
func (o *RepoUnDismissPullReviewOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo un dismiss pull review o k response has a 4xx status code
func (o *RepoUnDismissPullReviewOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this repo un dismiss pull review o k response has a 5xx status code
func (o *RepoUnDismissPullReviewOK) IsServerError() bool {
	return false
}

// IsCode returns true when this repo un dismiss pull review o k response a status code equal to that given
func (o *RepoUnDismissPullReviewOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the repo un dismiss pull review o k response
func (o *RepoUnDismissPullReviewOK) Code() int {
	return 200
}

func (o *RepoUnDismissPullReviewOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/pulls/{index}/reviews/{id}/undismissals][%d] repoUnDismissPullReviewOK %s", 200, payload)
}

func (o *RepoUnDismissPullReviewOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/pulls/{index}/reviews/{id}/undismissals][%d] repoUnDismissPullReviewOK %s", 200, payload)
}

func (o *RepoUnDismissPullReviewOK) GetPayload() *models.PullReview {
	return o.Payload
}

func (o *RepoUnDismissPullReviewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PullReview)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepoUnDismissPullReviewForbidden creates a RepoUnDismissPullReviewForbidden with default headers values
func NewRepoUnDismissPullReviewForbidden() *RepoUnDismissPullReviewForbidden {
	return &RepoUnDismissPullReviewForbidden{}
}

/*
RepoUnDismissPullReviewForbidden describes a response with status code 403, with default header values.

APIForbiddenError is a forbidden error response
*/
type RepoUnDismissPullReviewForbidden struct {
	Message string
	URL     string
}

// IsSuccess returns true when this repo un dismiss pull review forbidden response has a 2xx status code
func (o *RepoUnDismissPullReviewForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo un dismiss pull review forbidden response has a 3xx status code
func (o *RepoUnDismissPullReviewForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo un dismiss pull review forbidden response has a 4xx status code
func (o *RepoUnDismissPullReviewForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo un dismiss pull review forbidden response has a 5xx status code
func (o *RepoUnDismissPullReviewForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this repo un dismiss pull review forbidden response a status code equal to that given
func (o *RepoUnDismissPullReviewForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the repo un dismiss pull review forbidden response
func (o *RepoUnDismissPullReviewForbidden) Code() int {
	return 403
}

func (o *RepoUnDismissPullReviewForbidden) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/pulls/{index}/reviews/{id}/undismissals][%d] repoUnDismissPullReviewForbidden", 403)
}

func (o *RepoUnDismissPullReviewForbidden) String() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/pulls/{index}/reviews/{id}/undismissals][%d] repoUnDismissPullReviewForbidden", 403)
}

func (o *RepoUnDismissPullReviewForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRepoUnDismissPullReviewNotFound creates a RepoUnDismissPullReviewNotFound with default headers values
func NewRepoUnDismissPullReviewNotFound() *RepoUnDismissPullReviewNotFound {
	return &RepoUnDismissPullReviewNotFound{}
}

/*
RepoUnDismissPullReviewNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type RepoUnDismissPullReviewNotFound struct {
}

// IsSuccess returns true when this repo un dismiss pull review not found response has a 2xx status code
func (o *RepoUnDismissPullReviewNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo un dismiss pull review not found response has a 3xx status code
func (o *RepoUnDismissPullReviewNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo un dismiss pull review not found response has a 4xx status code
func (o *RepoUnDismissPullReviewNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo un dismiss pull review not found response has a 5xx status code
func (o *RepoUnDismissPullReviewNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repo un dismiss pull review not found response a status code equal to that given
func (o *RepoUnDismissPullReviewNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the repo un dismiss pull review not found response
func (o *RepoUnDismissPullReviewNotFound) Code() int {
	return 404
}

func (o *RepoUnDismissPullReviewNotFound) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/pulls/{index}/reviews/{id}/undismissals][%d] repoUnDismissPullReviewNotFound", 404)
}

func (o *RepoUnDismissPullReviewNotFound) String() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/pulls/{index}/reviews/{id}/undismissals][%d] repoUnDismissPullReviewNotFound", 404)
}

func (o *RepoUnDismissPullReviewNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRepoUnDismissPullReviewUnprocessableEntity creates a RepoUnDismissPullReviewUnprocessableEntity with default headers values
func NewRepoUnDismissPullReviewUnprocessableEntity() *RepoUnDismissPullReviewUnprocessableEntity {
	return &RepoUnDismissPullReviewUnprocessableEntity{}
}

/*
RepoUnDismissPullReviewUnprocessableEntity describes a response with status code 422, with default header values.

APIValidationError is error format response related to input validation
*/
type RepoUnDismissPullReviewUnprocessableEntity struct {
	Message string
	URL     string
}

// IsSuccess returns true when this repo un dismiss pull review unprocessable entity response has a 2xx status code
func (o *RepoUnDismissPullReviewUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo un dismiss pull review unprocessable entity response has a 3xx status code
func (o *RepoUnDismissPullReviewUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo un dismiss pull review unprocessable entity response has a 4xx status code
func (o *RepoUnDismissPullReviewUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo un dismiss pull review unprocessable entity response has a 5xx status code
func (o *RepoUnDismissPullReviewUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this repo un dismiss pull review unprocessable entity response a status code equal to that given
func (o *RepoUnDismissPullReviewUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the repo un dismiss pull review unprocessable entity response
func (o *RepoUnDismissPullReviewUnprocessableEntity) Code() int {
	return 422
}

func (o *RepoUnDismissPullReviewUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/pulls/{index}/reviews/{id}/undismissals][%d] repoUnDismissPullReviewUnprocessableEntity", 422)
}

func (o *RepoUnDismissPullReviewUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/pulls/{index}/reviews/{id}/undismissals][%d] repoUnDismissPullReviewUnprocessableEntity", 422)
}

func (o *RepoUnDismissPullReviewUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
