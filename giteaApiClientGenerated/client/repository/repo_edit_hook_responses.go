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

// RepoEditHookReader is a Reader for the RepoEditHook structure.
type RepoEditHookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoEditHookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoEditHookOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRepoEditHookNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /repos/{owner}/{repo}/hooks/{id}] repoEditHook", response, response.Code())
	}
}

// NewRepoEditHookOK creates a RepoEditHookOK with default headers values
func NewRepoEditHookOK() *RepoEditHookOK {
	return &RepoEditHookOK{}
}

/*
RepoEditHookOK describes a response with status code 200, with default header values.

Hook
*/
type RepoEditHookOK struct {
	Payload *models.Hook
}

// IsSuccess returns true when this repo edit hook o k response has a 2xx status code
func (o *RepoEditHookOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repo edit hook o k response has a 3xx status code
func (o *RepoEditHookOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo edit hook o k response has a 4xx status code
func (o *RepoEditHookOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this repo edit hook o k response has a 5xx status code
func (o *RepoEditHookOK) IsServerError() bool {
	return false
}

// IsCode returns true when this repo edit hook o k response a status code equal to that given
func (o *RepoEditHookOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the repo edit hook o k response
func (o *RepoEditHookOK) Code() int {
	return 200
}

func (o *RepoEditHookOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}/hooks/{id}][%d] repoEditHookOK %s", 200, payload)
}

func (o *RepoEditHookOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}/hooks/{id}][%d] repoEditHookOK %s", 200, payload)
}

func (o *RepoEditHookOK) GetPayload() *models.Hook {
	return o.Payload
}

func (o *RepoEditHookOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Hook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepoEditHookNotFound creates a RepoEditHookNotFound with default headers values
func NewRepoEditHookNotFound() *RepoEditHookNotFound {
	return &RepoEditHookNotFound{}
}

/*
RepoEditHookNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type RepoEditHookNotFound struct {
}

// IsSuccess returns true when this repo edit hook not found response has a 2xx status code
func (o *RepoEditHookNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo edit hook not found response has a 3xx status code
func (o *RepoEditHookNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo edit hook not found response has a 4xx status code
func (o *RepoEditHookNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo edit hook not found response has a 5xx status code
func (o *RepoEditHookNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repo edit hook not found response a status code equal to that given
func (o *RepoEditHookNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the repo edit hook not found response
func (o *RepoEditHookNotFound) Code() int {
	return 404
}

func (o *RepoEditHookNotFound) Error() string {
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}/hooks/{id}][%d] repoEditHookNotFound", 404)
}

func (o *RepoEditHookNotFound) String() string {
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}/hooks/{id}][%d] repoEditHookNotFound", 404)
}

func (o *RepoEditHookNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
