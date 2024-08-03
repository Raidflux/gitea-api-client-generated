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

// RepoEditReader is a Reader for the RepoEdit structure.
type RepoEditReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoEditReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoEditOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewRepoEditForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRepoEditNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewRepoEditUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /repos/{owner}/{repo}] repoEdit", response, response.Code())
	}
}

// NewRepoEditOK creates a RepoEditOK with default headers values
func NewRepoEditOK() *RepoEditOK {
	return &RepoEditOK{}
}

/*
RepoEditOK describes a response with status code 200, with default header values.

Repository
*/
type RepoEditOK struct {
	Payload *models.Repository
}

// IsSuccess returns true when this repo edit o k response has a 2xx status code
func (o *RepoEditOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repo edit o k response has a 3xx status code
func (o *RepoEditOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo edit o k response has a 4xx status code
func (o *RepoEditOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this repo edit o k response has a 5xx status code
func (o *RepoEditOK) IsServerError() bool {
	return false
}

// IsCode returns true when this repo edit o k response a status code equal to that given
func (o *RepoEditOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the repo edit o k response
func (o *RepoEditOK) Code() int {
	return 200
}

func (o *RepoEditOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}][%d] repoEditOK %s", 200, payload)
}

func (o *RepoEditOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}][%d] repoEditOK %s", 200, payload)
}

func (o *RepoEditOK) GetPayload() *models.Repository {
	return o.Payload
}

func (o *RepoEditOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Repository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepoEditForbidden creates a RepoEditForbidden with default headers values
func NewRepoEditForbidden() *RepoEditForbidden {
	return &RepoEditForbidden{}
}

/*
RepoEditForbidden describes a response with status code 403, with default header values.

APIForbiddenError is a forbidden error response
*/
type RepoEditForbidden struct {
	Message string
	URL     string
}

// IsSuccess returns true when this repo edit forbidden response has a 2xx status code
func (o *RepoEditForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo edit forbidden response has a 3xx status code
func (o *RepoEditForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo edit forbidden response has a 4xx status code
func (o *RepoEditForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo edit forbidden response has a 5xx status code
func (o *RepoEditForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this repo edit forbidden response a status code equal to that given
func (o *RepoEditForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the repo edit forbidden response
func (o *RepoEditForbidden) Code() int {
	return 403
}

func (o *RepoEditForbidden) Error() string {
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}][%d] repoEditForbidden", 403)
}

func (o *RepoEditForbidden) String() string {
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}][%d] repoEditForbidden", 403)
}

func (o *RepoEditForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRepoEditNotFound creates a RepoEditNotFound with default headers values
func NewRepoEditNotFound() *RepoEditNotFound {
	return &RepoEditNotFound{}
}

/*
RepoEditNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type RepoEditNotFound struct {
}

// IsSuccess returns true when this repo edit not found response has a 2xx status code
func (o *RepoEditNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo edit not found response has a 3xx status code
func (o *RepoEditNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo edit not found response has a 4xx status code
func (o *RepoEditNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo edit not found response has a 5xx status code
func (o *RepoEditNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repo edit not found response a status code equal to that given
func (o *RepoEditNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the repo edit not found response
func (o *RepoEditNotFound) Code() int {
	return 404
}

func (o *RepoEditNotFound) Error() string {
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}][%d] repoEditNotFound", 404)
}

func (o *RepoEditNotFound) String() string {
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}][%d] repoEditNotFound", 404)
}

func (o *RepoEditNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRepoEditUnprocessableEntity creates a RepoEditUnprocessableEntity with default headers values
func NewRepoEditUnprocessableEntity() *RepoEditUnprocessableEntity {
	return &RepoEditUnprocessableEntity{}
}

/*
RepoEditUnprocessableEntity describes a response with status code 422, with default header values.

APIValidationError is error format response related to input validation
*/
type RepoEditUnprocessableEntity struct {
	Message string
	URL     string
}

// IsSuccess returns true when this repo edit unprocessable entity response has a 2xx status code
func (o *RepoEditUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo edit unprocessable entity response has a 3xx status code
func (o *RepoEditUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo edit unprocessable entity response has a 4xx status code
func (o *RepoEditUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo edit unprocessable entity response has a 5xx status code
func (o *RepoEditUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this repo edit unprocessable entity response a status code equal to that given
func (o *RepoEditUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the repo edit unprocessable entity response
func (o *RepoEditUnprocessableEntity) Code() int {
	return 422
}

func (o *RepoEditUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}][%d] repoEditUnprocessableEntity", 422)
}

func (o *RepoEditUnprocessableEntity) String() string {
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}][%d] repoEditUnprocessableEntity", 422)
}

func (o *RepoEditUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
