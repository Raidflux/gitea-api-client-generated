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

// RepoCreateTagReader is a Reader for the RepoCreateTag structure.
type RepoCreateTagReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoCreateTagReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoCreateTagOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRepoCreateTagNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewRepoCreateTagMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewRepoCreateTagConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewRepoCreateTagUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 423:
		result := NewRepoCreateTagLocked()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /repos/{owner}/{repo}/tags] repoCreateTag", response, response.Code())
	}
}

// NewRepoCreateTagOK creates a RepoCreateTagOK with default headers values
func NewRepoCreateTagOK() *RepoCreateTagOK {
	return &RepoCreateTagOK{}
}

/*
RepoCreateTagOK describes a response with status code 200, with default header values.

Tag
*/
type RepoCreateTagOK struct {
	Payload *models.Tag
}

// IsSuccess returns true when this repo create tag o k response has a 2xx status code
func (o *RepoCreateTagOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repo create tag o k response has a 3xx status code
func (o *RepoCreateTagOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo create tag o k response has a 4xx status code
func (o *RepoCreateTagOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this repo create tag o k response has a 5xx status code
func (o *RepoCreateTagOK) IsServerError() bool {
	return false
}

// IsCode returns true when this repo create tag o k response a status code equal to that given
func (o *RepoCreateTagOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the repo create tag o k response
func (o *RepoCreateTagOK) Code() int {
	return 200
}

func (o *RepoCreateTagOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/tags][%d] repoCreateTagOK %s", 200, payload)
}

func (o *RepoCreateTagOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/tags][%d] repoCreateTagOK %s", 200, payload)
}

func (o *RepoCreateTagOK) GetPayload() *models.Tag {
	return o.Payload
}

func (o *RepoCreateTagOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepoCreateTagNotFound creates a RepoCreateTagNotFound with default headers values
func NewRepoCreateTagNotFound() *RepoCreateTagNotFound {
	return &RepoCreateTagNotFound{}
}

/*
RepoCreateTagNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type RepoCreateTagNotFound struct {
}

// IsSuccess returns true when this repo create tag not found response has a 2xx status code
func (o *RepoCreateTagNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo create tag not found response has a 3xx status code
func (o *RepoCreateTagNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo create tag not found response has a 4xx status code
func (o *RepoCreateTagNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo create tag not found response has a 5xx status code
func (o *RepoCreateTagNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repo create tag not found response a status code equal to that given
func (o *RepoCreateTagNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the repo create tag not found response
func (o *RepoCreateTagNotFound) Code() int {
	return 404
}

func (o *RepoCreateTagNotFound) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/tags][%d] repoCreateTagNotFound", 404)
}

func (o *RepoCreateTagNotFound) String() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/tags][%d] repoCreateTagNotFound", 404)
}

func (o *RepoCreateTagNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRepoCreateTagMethodNotAllowed creates a RepoCreateTagMethodNotAllowed with default headers values
func NewRepoCreateTagMethodNotAllowed() *RepoCreateTagMethodNotAllowed {
	return &RepoCreateTagMethodNotAllowed{}
}

/*
RepoCreateTagMethodNotAllowed describes a response with status code 405, with default header values.

APIEmpty is an empty response
*/
type RepoCreateTagMethodNotAllowed struct {
}

// IsSuccess returns true when this repo create tag method not allowed response has a 2xx status code
func (o *RepoCreateTagMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo create tag method not allowed response has a 3xx status code
func (o *RepoCreateTagMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo create tag method not allowed response has a 4xx status code
func (o *RepoCreateTagMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo create tag method not allowed response has a 5xx status code
func (o *RepoCreateTagMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this repo create tag method not allowed response a status code equal to that given
func (o *RepoCreateTagMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

// Code gets the status code for the repo create tag method not allowed response
func (o *RepoCreateTagMethodNotAllowed) Code() int {
	return 405
}

func (o *RepoCreateTagMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/tags][%d] repoCreateTagMethodNotAllowed", 405)
}

func (o *RepoCreateTagMethodNotAllowed) String() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/tags][%d] repoCreateTagMethodNotAllowed", 405)
}

func (o *RepoCreateTagMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRepoCreateTagConflict creates a RepoCreateTagConflict with default headers values
func NewRepoCreateTagConflict() *RepoCreateTagConflict {
	return &RepoCreateTagConflict{}
}

/*
RepoCreateTagConflict describes a response with status code 409, with default header values.

APIConflict is a conflict empty response
*/
type RepoCreateTagConflict struct {
}

// IsSuccess returns true when this repo create tag conflict response has a 2xx status code
func (o *RepoCreateTagConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo create tag conflict response has a 3xx status code
func (o *RepoCreateTagConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo create tag conflict response has a 4xx status code
func (o *RepoCreateTagConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo create tag conflict response has a 5xx status code
func (o *RepoCreateTagConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this repo create tag conflict response a status code equal to that given
func (o *RepoCreateTagConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the repo create tag conflict response
func (o *RepoCreateTagConflict) Code() int {
	return 409
}

func (o *RepoCreateTagConflict) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/tags][%d] repoCreateTagConflict", 409)
}

func (o *RepoCreateTagConflict) String() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/tags][%d] repoCreateTagConflict", 409)
}

func (o *RepoCreateTagConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRepoCreateTagUnprocessableEntity creates a RepoCreateTagUnprocessableEntity with default headers values
func NewRepoCreateTagUnprocessableEntity() *RepoCreateTagUnprocessableEntity {
	return &RepoCreateTagUnprocessableEntity{}
}

/*
RepoCreateTagUnprocessableEntity describes a response with status code 422, with default header values.

APIValidationError is error format response related to input validation
*/
type RepoCreateTagUnprocessableEntity struct {
	Message string
	URL     string
}

// IsSuccess returns true when this repo create tag unprocessable entity response has a 2xx status code
func (o *RepoCreateTagUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo create tag unprocessable entity response has a 3xx status code
func (o *RepoCreateTagUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo create tag unprocessable entity response has a 4xx status code
func (o *RepoCreateTagUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo create tag unprocessable entity response has a 5xx status code
func (o *RepoCreateTagUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this repo create tag unprocessable entity response a status code equal to that given
func (o *RepoCreateTagUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the repo create tag unprocessable entity response
func (o *RepoCreateTagUnprocessableEntity) Code() int {
	return 422
}

func (o *RepoCreateTagUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/tags][%d] repoCreateTagUnprocessableEntity", 422)
}

func (o *RepoCreateTagUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/tags][%d] repoCreateTagUnprocessableEntity", 422)
}

func (o *RepoCreateTagUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRepoCreateTagLocked creates a RepoCreateTagLocked with default headers values
func NewRepoCreateTagLocked() *RepoCreateTagLocked {
	return &RepoCreateTagLocked{}
}

/*
RepoCreateTagLocked describes a response with status code 423, with default header values.

APIRepoArchivedError is an error that is raised when an archived repo should be modified
*/
type RepoCreateTagLocked struct {
	Message string
	URL     string
}

// IsSuccess returns true when this repo create tag locked response has a 2xx status code
func (o *RepoCreateTagLocked) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo create tag locked response has a 3xx status code
func (o *RepoCreateTagLocked) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo create tag locked response has a 4xx status code
func (o *RepoCreateTagLocked) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo create tag locked response has a 5xx status code
func (o *RepoCreateTagLocked) IsServerError() bool {
	return false
}

// IsCode returns true when this repo create tag locked response a status code equal to that given
func (o *RepoCreateTagLocked) IsCode(code int) bool {
	return code == 423
}

// Code gets the status code for the repo create tag locked response
func (o *RepoCreateTagLocked) Code() int {
	return 423
}

func (o *RepoCreateTagLocked) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/tags][%d] repoCreateTagLocked", 423)
}

func (o *RepoCreateTagLocked) String() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/tags][%d] repoCreateTagLocked", 423)
}

func (o *RepoCreateTagLocked) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
