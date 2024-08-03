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

// RepoListKeysReader is a Reader for the RepoListKeys structure.
type RepoListKeysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoListKeysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoListKeysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRepoListKeysNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /repos/{owner}/{repo}/keys] repoListKeys", response, response.Code())
	}
}

// NewRepoListKeysOK creates a RepoListKeysOK with default headers values
func NewRepoListKeysOK() *RepoListKeysOK {
	return &RepoListKeysOK{}
}

/*
RepoListKeysOK describes a response with status code 200, with default header values.

DeployKeyList
*/
type RepoListKeysOK struct {
	Payload []*models.DeployKey
}

// IsSuccess returns true when this repo list keys o k response has a 2xx status code
func (o *RepoListKeysOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repo list keys o k response has a 3xx status code
func (o *RepoListKeysOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo list keys o k response has a 4xx status code
func (o *RepoListKeysOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this repo list keys o k response has a 5xx status code
func (o *RepoListKeysOK) IsServerError() bool {
	return false
}

// IsCode returns true when this repo list keys o k response a status code equal to that given
func (o *RepoListKeysOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the repo list keys o k response
func (o *RepoListKeysOK) Code() int {
	return 200
}

func (o *RepoListKeysOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/keys][%d] repoListKeysOK %s", 200, payload)
}

func (o *RepoListKeysOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/keys][%d] repoListKeysOK %s", 200, payload)
}

func (o *RepoListKeysOK) GetPayload() []*models.DeployKey {
	return o.Payload
}

func (o *RepoListKeysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepoListKeysNotFound creates a RepoListKeysNotFound with default headers values
func NewRepoListKeysNotFound() *RepoListKeysNotFound {
	return &RepoListKeysNotFound{}
}

/*
RepoListKeysNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type RepoListKeysNotFound struct {
}

// IsSuccess returns true when this repo list keys not found response has a 2xx status code
func (o *RepoListKeysNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo list keys not found response has a 3xx status code
func (o *RepoListKeysNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo list keys not found response has a 4xx status code
func (o *RepoListKeysNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo list keys not found response has a 5xx status code
func (o *RepoListKeysNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repo list keys not found response a status code equal to that given
func (o *RepoListKeysNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the repo list keys not found response
func (o *RepoListKeysNotFound) Code() int {
	return 404
}

func (o *RepoListKeysNotFound) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/keys][%d] repoListKeysNotFound", 404)
}

func (o *RepoListKeysNotFound) String() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/keys][%d] repoListKeysNotFound", 404)
}

func (o *RepoListKeysNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
