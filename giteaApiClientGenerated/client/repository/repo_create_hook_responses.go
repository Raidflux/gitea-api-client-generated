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

// RepoCreateHookReader is a Reader for the RepoCreateHook structure.
type RepoCreateHookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoCreateHookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewRepoCreateHookCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRepoCreateHookNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /repos/{owner}/{repo}/hooks] repoCreateHook", response, response.Code())
	}
}

// NewRepoCreateHookCreated creates a RepoCreateHookCreated with default headers values
func NewRepoCreateHookCreated() *RepoCreateHookCreated {
	return &RepoCreateHookCreated{}
}

/*
RepoCreateHookCreated describes a response with status code 201, with default header values.

Hook
*/
type RepoCreateHookCreated struct {
	Payload *models.Hook
}

// IsSuccess returns true when this repo create hook created response has a 2xx status code
func (o *RepoCreateHookCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repo create hook created response has a 3xx status code
func (o *RepoCreateHookCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo create hook created response has a 4xx status code
func (o *RepoCreateHookCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this repo create hook created response has a 5xx status code
func (o *RepoCreateHookCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this repo create hook created response a status code equal to that given
func (o *RepoCreateHookCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the repo create hook created response
func (o *RepoCreateHookCreated) Code() int {
	return 201
}

func (o *RepoCreateHookCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/hooks][%d] repoCreateHookCreated %s", 201, payload)
}

func (o *RepoCreateHookCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/hooks][%d] repoCreateHookCreated %s", 201, payload)
}

func (o *RepoCreateHookCreated) GetPayload() *models.Hook {
	return o.Payload
}

func (o *RepoCreateHookCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Hook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepoCreateHookNotFound creates a RepoCreateHookNotFound with default headers values
func NewRepoCreateHookNotFound() *RepoCreateHookNotFound {
	return &RepoCreateHookNotFound{}
}

/*
RepoCreateHookNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type RepoCreateHookNotFound struct {
}

// IsSuccess returns true when this repo create hook not found response has a 2xx status code
func (o *RepoCreateHookNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo create hook not found response has a 3xx status code
func (o *RepoCreateHookNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo create hook not found response has a 4xx status code
func (o *RepoCreateHookNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo create hook not found response has a 5xx status code
func (o *RepoCreateHookNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repo create hook not found response a status code equal to that given
func (o *RepoCreateHookNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the repo create hook not found response
func (o *RepoCreateHookNotFound) Code() int {
	return 404
}

func (o *RepoCreateHookNotFound) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/hooks][%d] repoCreateHookNotFound", 404)
}

func (o *RepoCreateHookNotFound) String() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/hooks][%d] repoCreateHookNotFound", 404)
}

func (o *RepoCreateHookNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
