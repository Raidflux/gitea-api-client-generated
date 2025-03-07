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

// RepoGetReleaseReader is a Reader for the RepoGetRelease structure.
type RepoGetReleaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoGetReleaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoGetReleaseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRepoGetReleaseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /repos/{owner}/{repo}/releases/{id}] repoGetRelease", response, response.Code())
	}
}

// NewRepoGetReleaseOK creates a RepoGetReleaseOK with default headers values
func NewRepoGetReleaseOK() *RepoGetReleaseOK {
	return &RepoGetReleaseOK{}
}

/*
RepoGetReleaseOK describes a response with status code 200, with default header values.

Release
*/
type RepoGetReleaseOK struct {
	Payload *models.Release
}

// IsSuccess returns true when this repo get release o k response has a 2xx status code
func (o *RepoGetReleaseOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repo get release o k response has a 3xx status code
func (o *RepoGetReleaseOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo get release o k response has a 4xx status code
func (o *RepoGetReleaseOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this repo get release o k response has a 5xx status code
func (o *RepoGetReleaseOK) IsServerError() bool {
	return false
}

// IsCode returns true when this repo get release o k response a status code equal to that given
func (o *RepoGetReleaseOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the repo get release o k response
func (o *RepoGetReleaseOK) Code() int {
	return 200
}

func (o *RepoGetReleaseOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/releases/{id}][%d] repoGetReleaseOK %s", 200, payload)
}

func (o *RepoGetReleaseOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/releases/{id}][%d] repoGetReleaseOK %s", 200, payload)
}

func (o *RepoGetReleaseOK) GetPayload() *models.Release {
	return o.Payload
}

func (o *RepoGetReleaseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Release)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepoGetReleaseNotFound creates a RepoGetReleaseNotFound with default headers values
func NewRepoGetReleaseNotFound() *RepoGetReleaseNotFound {
	return &RepoGetReleaseNotFound{}
}

/*
RepoGetReleaseNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type RepoGetReleaseNotFound struct {
}

// IsSuccess returns true when this repo get release not found response has a 2xx status code
func (o *RepoGetReleaseNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo get release not found response has a 3xx status code
func (o *RepoGetReleaseNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo get release not found response has a 4xx status code
func (o *RepoGetReleaseNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo get release not found response has a 5xx status code
func (o *RepoGetReleaseNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repo get release not found response a status code equal to that given
func (o *RepoGetReleaseNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the repo get release not found response
func (o *RepoGetReleaseNotFound) Code() int {
	return 404
}

func (o *RepoGetReleaseNotFound) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/releases/{id}][%d] repoGetReleaseNotFound", 404)
}

func (o *RepoGetReleaseNotFound) String() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/releases/{id}][%d] repoGetReleaseNotFound", 404)
}

func (o *RepoGetReleaseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
