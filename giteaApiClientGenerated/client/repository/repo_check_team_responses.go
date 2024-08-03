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

// RepoCheckTeamReader is a Reader for the RepoCheckTeam structure.
type RepoCheckTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoCheckTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoCheckTeamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRepoCheckTeamNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewRepoCheckTeamMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /repos/{owner}/{repo}/teams/{team}] repoCheckTeam", response, response.Code())
	}
}

// NewRepoCheckTeamOK creates a RepoCheckTeamOK with default headers values
func NewRepoCheckTeamOK() *RepoCheckTeamOK {
	return &RepoCheckTeamOK{}
}

/*
RepoCheckTeamOK describes a response with status code 200, with default header values.

Team
*/
type RepoCheckTeamOK struct {
	Payload *models.Team
}

// IsSuccess returns true when this repo check team o k response has a 2xx status code
func (o *RepoCheckTeamOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repo check team o k response has a 3xx status code
func (o *RepoCheckTeamOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo check team o k response has a 4xx status code
func (o *RepoCheckTeamOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this repo check team o k response has a 5xx status code
func (o *RepoCheckTeamOK) IsServerError() bool {
	return false
}

// IsCode returns true when this repo check team o k response a status code equal to that given
func (o *RepoCheckTeamOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the repo check team o k response
func (o *RepoCheckTeamOK) Code() int {
	return 200
}

func (o *RepoCheckTeamOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/teams/{team}][%d] repoCheckTeamOK %s", 200, payload)
}

func (o *RepoCheckTeamOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/teams/{team}][%d] repoCheckTeamOK %s", 200, payload)
}

func (o *RepoCheckTeamOK) GetPayload() *models.Team {
	return o.Payload
}

func (o *RepoCheckTeamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Team)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepoCheckTeamNotFound creates a RepoCheckTeamNotFound with default headers values
func NewRepoCheckTeamNotFound() *RepoCheckTeamNotFound {
	return &RepoCheckTeamNotFound{}
}

/*
RepoCheckTeamNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type RepoCheckTeamNotFound struct {
}

// IsSuccess returns true when this repo check team not found response has a 2xx status code
func (o *RepoCheckTeamNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo check team not found response has a 3xx status code
func (o *RepoCheckTeamNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo check team not found response has a 4xx status code
func (o *RepoCheckTeamNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo check team not found response has a 5xx status code
func (o *RepoCheckTeamNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repo check team not found response a status code equal to that given
func (o *RepoCheckTeamNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the repo check team not found response
func (o *RepoCheckTeamNotFound) Code() int {
	return 404
}

func (o *RepoCheckTeamNotFound) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/teams/{team}][%d] repoCheckTeamNotFound", 404)
}

func (o *RepoCheckTeamNotFound) String() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/teams/{team}][%d] repoCheckTeamNotFound", 404)
}

func (o *RepoCheckTeamNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRepoCheckTeamMethodNotAllowed creates a RepoCheckTeamMethodNotAllowed with default headers values
func NewRepoCheckTeamMethodNotAllowed() *RepoCheckTeamMethodNotAllowed {
	return &RepoCheckTeamMethodNotAllowed{}
}

/*
RepoCheckTeamMethodNotAllowed describes a response with status code 405, with default header values.

APIError is error format response
*/
type RepoCheckTeamMethodNotAllowed struct {
	Message string
	URL     string
}

// IsSuccess returns true when this repo check team method not allowed response has a 2xx status code
func (o *RepoCheckTeamMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo check team method not allowed response has a 3xx status code
func (o *RepoCheckTeamMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo check team method not allowed response has a 4xx status code
func (o *RepoCheckTeamMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo check team method not allowed response has a 5xx status code
func (o *RepoCheckTeamMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this repo check team method not allowed response a status code equal to that given
func (o *RepoCheckTeamMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

// Code gets the status code for the repo check team method not allowed response
func (o *RepoCheckTeamMethodNotAllowed) Code() int {
	return 405
}

func (o *RepoCheckTeamMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/teams/{team}][%d] repoCheckTeamMethodNotAllowed", 405)
}

func (o *RepoCheckTeamMethodNotAllowed) String() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/teams/{team}][%d] repoCheckTeamMethodNotAllowed", 405)
}

func (o *RepoCheckTeamMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
