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

// RepoGetLatestReleaseReader is a Reader for the RepoGetLatestRelease structure.
type RepoGetLatestReleaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoGetLatestReleaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoGetLatestReleaseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRepoGetLatestReleaseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /repos/{owner}/{repo}/releases/latest] repoGetLatestRelease", response, response.Code())
	}
}

// NewRepoGetLatestReleaseOK creates a RepoGetLatestReleaseOK with default headers values
func NewRepoGetLatestReleaseOK() *RepoGetLatestReleaseOK {
	return &RepoGetLatestReleaseOK{}
}

/*
RepoGetLatestReleaseOK describes a response with status code 200, with default header values.

Release
*/
type RepoGetLatestReleaseOK struct {
	Payload *models.Release
}

// IsSuccess returns true when this repo get latest release o k response has a 2xx status code
func (o *RepoGetLatestReleaseOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repo get latest release o k response has a 3xx status code
func (o *RepoGetLatestReleaseOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo get latest release o k response has a 4xx status code
func (o *RepoGetLatestReleaseOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this repo get latest release o k response has a 5xx status code
func (o *RepoGetLatestReleaseOK) IsServerError() bool {
	return false
}

// IsCode returns true when this repo get latest release o k response a status code equal to that given
func (o *RepoGetLatestReleaseOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the repo get latest release o k response
func (o *RepoGetLatestReleaseOK) Code() int {
	return 200
}

func (o *RepoGetLatestReleaseOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/releases/latest][%d] repoGetLatestReleaseOK %s", 200, payload)
}

func (o *RepoGetLatestReleaseOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/releases/latest][%d] repoGetLatestReleaseOK %s", 200, payload)
}

func (o *RepoGetLatestReleaseOK) GetPayload() *models.Release {
	return o.Payload
}

func (o *RepoGetLatestReleaseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Release)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepoGetLatestReleaseNotFound creates a RepoGetLatestReleaseNotFound with default headers values
func NewRepoGetLatestReleaseNotFound() *RepoGetLatestReleaseNotFound {
	return &RepoGetLatestReleaseNotFound{}
}

/*
RepoGetLatestReleaseNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type RepoGetLatestReleaseNotFound struct {
}

// IsSuccess returns true when this repo get latest release not found response has a 2xx status code
func (o *RepoGetLatestReleaseNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo get latest release not found response has a 3xx status code
func (o *RepoGetLatestReleaseNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo get latest release not found response has a 4xx status code
func (o *RepoGetLatestReleaseNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo get latest release not found response has a 5xx status code
func (o *RepoGetLatestReleaseNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repo get latest release not found response a status code equal to that given
func (o *RepoGetLatestReleaseNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the repo get latest release not found response
func (o *RepoGetLatestReleaseNotFound) Code() int {
	return 404
}

func (o *RepoGetLatestReleaseNotFound) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/releases/latest][%d] repoGetLatestReleaseNotFound", 404)
}

func (o *RepoGetLatestReleaseNotFound) String() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/releases/latest][%d] repoGetLatestReleaseNotFound", 404)
}

func (o *RepoGetLatestReleaseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
