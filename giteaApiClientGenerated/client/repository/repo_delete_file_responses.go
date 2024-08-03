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

// RepoDeleteFileReader is a Reader for the RepoDeleteFile structure.
type RepoDeleteFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoDeleteFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoDeleteFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRepoDeleteFileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRepoDeleteFileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRepoDeleteFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 423:
		result := NewRepoDeleteFileLocked()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /repos/{owner}/{repo}/contents/{filepath}] repoDeleteFile", response, response.Code())
	}
}

// NewRepoDeleteFileOK creates a RepoDeleteFileOK with default headers values
func NewRepoDeleteFileOK() *RepoDeleteFileOK {
	return &RepoDeleteFileOK{}
}

/*
RepoDeleteFileOK describes a response with status code 200, with default header values.

FileDeleteResponse
*/
type RepoDeleteFileOK struct {
	Payload *models.FileDeleteResponse
}

// IsSuccess returns true when this repo delete file o k response has a 2xx status code
func (o *RepoDeleteFileOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repo delete file o k response has a 3xx status code
func (o *RepoDeleteFileOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo delete file o k response has a 4xx status code
func (o *RepoDeleteFileOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this repo delete file o k response has a 5xx status code
func (o *RepoDeleteFileOK) IsServerError() bool {
	return false
}

// IsCode returns true when this repo delete file o k response a status code equal to that given
func (o *RepoDeleteFileOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the repo delete file o k response
func (o *RepoDeleteFileOK) Code() int {
	return 200
}

func (o *RepoDeleteFileOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/contents/{filepath}][%d] repoDeleteFileOK %s", 200, payload)
}

func (o *RepoDeleteFileOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/contents/{filepath}][%d] repoDeleteFileOK %s", 200, payload)
}

func (o *RepoDeleteFileOK) GetPayload() *models.FileDeleteResponse {
	return o.Payload
}

func (o *RepoDeleteFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FileDeleteResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepoDeleteFileBadRequest creates a RepoDeleteFileBadRequest with default headers values
func NewRepoDeleteFileBadRequest() *RepoDeleteFileBadRequest {
	return &RepoDeleteFileBadRequest{}
}

/*
RepoDeleteFileBadRequest describes a response with status code 400, with default header values.

APIError is error format response
*/
type RepoDeleteFileBadRequest struct {
	Message string
	URL     string
}

// IsSuccess returns true when this repo delete file bad request response has a 2xx status code
func (o *RepoDeleteFileBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo delete file bad request response has a 3xx status code
func (o *RepoDeleteFileBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo delete file bad request response has a 4xx status code
func (o *RepoDeleteFileBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo delete file bad request response has a 5xx status code
func (o *RepoDeleteFileBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this repo delete file bad request response a status code equal to that given
func (o *RepoDeleteFileBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the repo delete file bad request response
func (o *RepoDeleteFileBadRequest) Code() int {
	return 400
}

func (o *RepoDeleteFileBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/contents/{filepath}][%d] repoDeleteFileBadRequest", 400)
}

func (o *RepoDeleteFileBadRequest) String() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/contents/{filepath}][%d] repoDeleteFileBadRequest", 400)
}

func (o *RepoDeleteFileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRepoDeleteFileForbidden creates a RepoDeleteFileForbidden with default headers values
func NewRepoDeleteFileForbidden() *RepoDeleteFileForbidden {
	return &RepoDeleteFileForbidden{}
}

/*
RepoDeleteFileForbidden describes a response with status code 403, with default header values.

APIError is error format response
*/
type RepoDeleteFileForbidden struct {
	Message string
	URL     string
}

// IsSuccess returns true when this repo delete file forbidden response has a 2xx status code
func (o *RepoDeleteFileForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo delete file forbidden response has a 3xx status code
func (o *RepoDeleteFileForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo delete file forbidden response has a 4xx status code
func (o *RepoDeleteFileForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo delete file forbidden response has a 5xx status code
func (o *RepoDeleteFileForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this repo delete file forbidden response a status code equal to that given
func (o *RepoDeleteFileForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the repo delete file forbidden response
func (o *RepoDeleteFileForbidden) Code() int {
	return 403
}

func (o *RepoDeleteFileForbidden) Error() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/contents/{filepath}][%d] repoDeleteFileForbidden", 403)
}

func (o *RepoDeleteFileForbidden) String() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/contents/{filepath}][%d] repoDeleteFileForbidden", 403)
}

func (o *RepoDeleteFileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRepoDeleteFileNotFound creates a RepoDeleteFileNotFound with default headers values
func NewRepoDeleteFileNotFound() *RepoDeleteFileNotFound {
	return &RepoDeleteFileNotFound{}
}

/*
RepoDeleteFileNotFound describes a response with status code 404, with default header values.

APIError is error format response
*/
type RepoDeleteFileNotFound struct {
	Message string
	URL     string
}

// IsSuccess returns true when this repo delete file not found response has a 2xx status code
func (o *RepoDeleteFileNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo delete file not found response has a 3xx status code
func (o *RepoDeleteFileNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo delete file not found response has a 4xx status code
func (o *RepoDeleteFileNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo delete file not found response has a 5xx status code
func (o *RepoDeleteFileNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repo delete file not found response a status code equal to that given
func (o *RepoDeleteFileNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the repo delete file not found response
func (o *RepoDeleteFileNotFound) Code() int {
	return 404
}

func (o *RepoDeleteFileNotFound) Error() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/contents/{filepath}][%d] repoDeleteFileNotFound", 404)
}

func (o *RepoDeleteFileNotFound) String() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/contents/{filepath}][%d] repoDeleteFileNotFound", 404)
}

func (o *RepoDeleteFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRepoDeleteFileLocked creates a RepoDeleteFileLocked with default headers values
func NewRepoDeleteFileLocked() *RepoDeleteFileLocked {
	return &RepoDeleteFileLocked{}
}

/*
RepoDeleteFileLocked describes a response with status code 423, with default header values.

APIRepoArchivedError is an error that is raised when an archived repo should be modified
*/
type RepoDeleteFileLocked struct {
	Message string
	URL     string
}

// IsSuccess returns true when this repo delete file locked response has a 2xx status code
func (o *RepoDeleteFileLocked) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo delete file locked response has a 3xx status code
func (o *RepoDeleteFileLocked) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo delete file locked response has a 4xx status code
func (o *RepoDeleteFileLocked) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo delete file locked response has a 5xx status code
func (o *RepoDeleteFileLocked) IsServerError() bool {
	return false
}

// IsCode returns true when this repo delete file locked response a status code equal to that given
func (o *RepoDeleteFileLocked) IsCode(code int) bool {
	return code == 423
}

// Code gets the status code for the repo delete file locked response
func (o *RepoDeleteFileLocked) Code() int {
	return 423
}

func (o *RepoDeleteFileLocked) Error() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/contents/{filepath}][%d] repoDeleteFileLocked", 423)
}

func (o *RepoDeleteFileLocked) String() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/contents/{filepath}][%d] repoDeleteFileLocked", 423)
}

func (o *RepoDeleteFileLocked) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
