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

// RepoGetNoteReader is a Reader for the RepoGetNote structure.
type RepoGetNoteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoGetNoteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoGetNoteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRepoGetNoteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewRepoGetNoteUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /repos/{owner}/{repo}/git/notes/{sha}] repoGetNote", response, response.Code())
	}
}

// NewRepoGetNoteOK creates a RepoGetNoteOK with default headers values
func NewRepoGetNoteOK() *RepoGetNoteOK {
	return &RepoGetNoteOK{}
}

/*
RepoGetNoteOK describes a response with status code 200, with default header values.

Note
*/
type RepoGetNoteOK struct {
	Payload *models.Note
}

// IsSuccess returns true when this repo get note o k response has a 2xx status code
func (o *RepoGetNoteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repo get note o k response has a 3xx status code
func (o *RepoGetNoteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo get note o k response has a 4xx status code
func (o *RepoGetNoteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this repo get note o k response has a 5xx status code
func (o *RepoGetNoteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this repo get note o k response a status code equal to that given
func (o *RepoGetNoteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the repo get note o k response
func (o *RepoGetNoteOK) Code() int {
	return 200
}

func (o *RepoGetNoteOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/git/notes/{sha}][%d] repoGetNoteOK %s", 200, payload)
}

func (o *RepoGetNoteOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/git/notes/{sha}][%d] repoGetNoteOK %s", 200, payload)
}

func (o *RepoGetNoteOK) GetPayload() *models.Note {
	return o.Payload
}

func (o *RepoGetNoteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Note)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepoGetNoteNotFound creates a RepoGetNoteNotFound with default headers values
func NewRepoGetNoteNotFound() *RepoGetNoteNotFound {
	return &RepoGetNoteNotFound{}
}

/*
RepoGetNoteNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type RepoGetNoteNotFound struct {
}

// IsSuccess returns true when this repo get note not found response has a 2xx status code
func (o *RepoGetNoteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo get note not found response has a 3xx status code
func (o *RepoGetNoteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo get note not found response has a 4xx status code
func (o *RepoGetNoteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo get note not found response has a 5xx status code
func (o *RepoGetNoteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repo get note not found response a status code equal to that given
func (o *RepoGetNoteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the repo get note not found response
func (o *RepoGetNoteNotFound) Code() int {
	return 404
}

func (o *RepoGetNoteNotFound) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/git/notes/{sha}][%d] repoGetNoteNotFound", 404)
}

func (o *RepoGetNoteNotFound) String() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/git/notes/{sha}][%d] repoGetNoteNotFound", 404)
}

func (o *RepoGetNoteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRepoGetNoteUnprocessableEntity creates a RepoGetNoteUnprocessableEntity with default headers values
func NewRepoGetNoteUnprocessableEntity() *RepoGetNoteUnprocessableEntity {
	return &RepoGetNoteUnprocessableEntity{}
}

/*
RepoGetNoteUnprocessableEntity describes a response with status code 422, with default header values.

APIValidationError is error format response related to input validation
*/
type RepoGetNoteUnprocessableEntity struct {
	Message string
	URL     string
}

// IsSuccess returns true when this repo get note unprocessable entity response has a 2xx status code
func (o *RepoGetNoteUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo get note unprocessable entity response has a 3xx status code
func (o *RepoGetNoteUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo get note unprocessable entity response has a 4xx status code
func (o *RepoGetNoteUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo get note unprocessable entity response has a 5xx status code
func (o *RepoGetNoteUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this repo get note unprocessable entity response a status code equal to that given
func (o *RepoGetNoteUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the repo get note unprocessable entity response
func (o *RepoGetNoteUnprocessableEntity) Code() int {
	return 422
}

func (o *RepoGetNoteUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/git/notes/{sha}][%d] repoGetNoteUnprocessableEntity", 422)
}

func (o *RepoGetNoteUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/git/notes/{sha}][%d] repoGetNoteUnprocessableEntity", 422)
}

func (o *RepoGetNoteUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
