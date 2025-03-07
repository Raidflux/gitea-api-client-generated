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

// RepoDismissPullReviewReader is a Reader for the RepoDismissPullReview structure.
type RepoDismissPullReviewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoDismissPullReviewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoDismissPullReviewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewRepoDismissPullReviewForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRepoDismissPullReviewNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewRepoDismissPullReviewUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /repos/{owner}/{repo}/pulls/{index}/reviews/{id}/dismissals] repoDismissPullReview", response, response.Code())
	}
}

// NewRepoDismissPullReviewOK creates a RepoDismissPullReviewOK with default headers values
func NewRepoDismissPullReviewOK() *RepoDismissPullReviewOK {
	return &RepoDismissPullReviewOK{}
}

/*
RepoDismissPullReviewOK describes a response with status code 200, with default header values.

PullReview
*/
type RepoDismissPullReviewOK struct {
	Payload *models.PullReview
}

// IsSuccess returns true when this repo dismiss pull review o k response has a 2xx status code
func (o *RepoDismissPullReviewOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repo dismiss pull review o k response has a 3xx status code
func (o *RepoDismissPullReviewOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo dismiss pull review o k response has a 4xx status code
func (o *RepoDismissPullReviewOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this repo dismiss pull review o k response has a 5xx status code
func (o *RepoDismissPullReviewOK) IsServerError() bool {
	return false
}

// IsCode returns true when this repo dismiss pull review o k response a status code equal to that given
func (o *RepoDismissPullReviewOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the repo dismiss pull review o k response
func (o *RepoDismissPullReviewOK) Code() int {
	return 200
}

func (o *RepoDismissPullReviewOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/pulls/{index}/reviews/{id}/dismissals][%d] repoDismissPullReviewOK %s", 200, payload)
}

func (o *RepoDismissPullReviewOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/pulls/{index}/reviews/{id}/dismissals][%d] repoDismissPullReviewOK %s", 200, payload)
}

func (o *RepoDismissPullReviewOK) GetPayload() *models.PullReview {
	return o.Payload
}

func (o *RepoDismissPullReviewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PullReview)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepoDismissPullReviewForbidden creates a RepoDismissPullReviewForbidden with default headers values
func NewRepoDismissPullReviewForbidden() *RepoDismissPullReviewForbidden {
	return &RepoDismissPullReviewForbidden{}
}

/*
RepoDismissPullReviewForbidden describes a response with status code 403, with default header values.

APIForbiddenError is a forbidden error response
*/
type RepoDismissPullReviewForbidden struct {
	Message string
	URL     string
}

// IsSuccess returns true when this repo dismiss pull review forbidden response has a 2xx status code
func (o *RepoDismissPullReviewForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo dismiss pull review forbidden response has a 3xx status code
func (o *RepoDismissPullReviewForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo dismiss pull review forbidden response has a 4xx status code
func (o *RepoDismissPullReviewForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo dismiss pull review forbidden response has a 5xx status code
func (o *RepoDismissPullReviewForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this repo dismiss pull review forbidden response a status code equal to that given
func (o *RepoDismissPullReviewForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the repo dismiss pull review forbidden response
func (o *RepoDismissPullReviewForbidden) Code() int {
	return 403
}

func (o *RepoDismissPullReviewForbidden) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/pulls/{index}/reviews/{id}/dismissals][%d] repoDismissPullReviewForbidden", 403)
}

func (o *RepoDismissPullReviewForbidden) String() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/pulls/{index}/reviews/{id}/dismissals][%d] repoDismissPullReviewForbidden", 403)
}

func (o *RepoDismissPullReviewForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRepoDismissPullReviewNotFound creates a RepoDismissPullReviewNotFound with default headers values
func NewRepoDismissPullReviewNotFound() *RepoDismissPullReviewNotFound {
	return &RepoDismissPullReviewNotFound{}
}

/*
RepoDismissPullReviewNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type RepoDismissPullReviewNotFound struct {
}

// IsSuccess returns true when this repo dismiss pull review not found response has a 2xx status code
func (o *RepoDismissPullReviewNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo dismiss pull review not found response has a 3xx status code
func (o *RepoDismissPullReviewNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo dismiss pull review not found response has a 4xx status code
func (o *RepoDismissPullReviewNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo dismiss pull review not found response has a 5xx status code
func (o *RepoDismissPullReviewNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repo dismiss pull review not found response a status code equal to that given
func (o *RepoDismissPullReviewNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the repo dismiss pull review not found response
func (o *RepoDismissPullReviewNotFound) Code() int {
	return 404
}

func (o *RepoDismissPullReviewNotFound) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/pulls/{index}/reviews/{id}/dismissals][%d] repoDismissPullReviewNotFound", 404)
}

func (o *RepoDismissPullReviewNotFound) String() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/pulls/{index}/reviews/{id}/dismissals][%d] repoDismissPullReviewNotFound", 404)
}

func (o *RepoDismissPullReviewNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRepoDismissPullReviewUnprocessableEntity creates a RepoDismissPullReviewUnprocessableEntity with default headers values
func NewRepoDismissPullReviewUnprocessableEntity() *RepoDismissPullReviewUnprocessableEntity {
	return &RepoDismissPullReviewUnprocessableEntity{}
}

/*
RepoDismissPullReviewUnprocessableEntity describes a response with status code 422, with default header values.

APIValidationError is error format response related to input validation
*/
type RepoDismissPullReviewUnprocessableEntity struct {
	Message string
	URL     string
}

// IsSuccess returns true when this repo dismiss pull review unprocessable entity response has a 2xx status code
func (o *RepoDismissPullReviewUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo dismiss pull review unprocessable entity response has a 3xx status code
func (o *RepoDismissPullReviewUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo dismiss pull review unprocessable entity response has a 4xx status code
func (o *RepoDismissPullReviewUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo dismiss pull review unprocessable entity response has a 5xx status code
func (o *RepoDismissPullReviewUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this repo dismiss pull review unprocessable entity response a status code equal to that given
func (o *RepoDismissPullReviewUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the repo dismiss pull review unprocessable entity response
func (o *RepoDismissPullReviewUnprocessableEntity) Code() int {
	return 422
}

func (o *RepoDismissPullReviewUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/pulls/{index}/reviews/{id}/dismissals][%d] repoDismissPullReviewUnprocessableEntity", 422)
}

func (o *RepoDismissPullReviewUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/pulls/{index}/reviews/{id}/dismissals][%d] repoDismissPullReviewUnprocessableEntity", 422)
}

func (o *RepoDismissPullReviewUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
