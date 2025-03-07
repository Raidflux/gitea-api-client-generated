// Code generated by go-swagger; DO NOT EDIT.

package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RepoDeleteReleaseReader is a Reader for the RepoDeleteRelease structure.
type RepoDeleteReleaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoDeleteReleaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRepoDeleteReleaseNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRepoDeleteReleaseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewRepoDeleteReleaseUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /repos/{owner}/{repo}/releases/{id}] repoDeleteRelease", response, response.Code())
	}
}

// NewRepoDeleteReleaseNoContent creates a RepoDeleteReleaseNoContent with default headers values
func NewRepoDeleteReleaseNoContent() *RepoDeleteReleaseNoContent {
	return &RepoDeleteReleaseNoContent{}
}

/*
RepoDeleteReleaseNoContent describes a response with status code 204, with default header values.

APIEmpty is an empty response
*/
type RepoDeleteReleaseNoContent struct {
}

// IsSuccess returns true when this repo delete release no content response has a 2xx status code
func (o *RepoDeleteReleaseNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repo delete release no content response has a 3xx status code
func (o *RepoDeleteReleaseNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo delete release no content response has a 4xx status code
func (o *RepoDeleteReleaseNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this repo delete release no content response has a 5xx status code
func (o *RepoDeleteReleaseNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this repo delete release no content response a status code equal to that given
func (o *RepoDeleteReleaseNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the repo delete release no content response
func (o *RepoDeleteReleaseNoContent) Code() int {
	return 204
}

func (o *RepoDeleteReleaseNoContent) Error() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/releases/{id}][%d] repoDeleteReleaseNoContent", 204)
}

func (o *RepoDeleteReleaseNoContent) String() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/releases/{id}][%d] repoDeleteReleaseNoContent", 204)
}

func (o *RepoDeleteReleaseNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRepoDeleteReleaseNotFound creates a RepoDeleteReleaseNotFound with default headers values
func NewRepoDeleteReleaseNotFound() *RepoDeleteReleaseNotFound {
	return &RepoDeleteReleaseNotFound{}
}

/*
RepoDeleteReleaseNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type RepoDeleteReleaseNotFound struct {
}

// IsSuccess returns true when this repo delete release not found response has a 2xx status code
func (o *RepoDeleteReleaseNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo delete release not found response has a 3xx status code
func (o *RepoDeleteReleaseNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo delete release not found response has a 4xx status code
func (o *RepoDeleteReleaseNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo delete release not found response has a 5xx status code
func (o *RepoDeleteReleaseNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repo delete release not found response a status code equal to that given
func (o *RepoDeleteReleaseNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the repo delete release not found response
func (o *RepoDeleteReleaseNotFound) Code() int {
	return 404
}

func (o *RepoDeleteReleaseNotFound) Error() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/releases/{id}][%d] repoDeleteReleaseNotFound", 404)
}

func (o *RepoDeleteReleaseNotFound) String() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/releases/{id}][%d] repoDeleteReleaseNotFound", 404)
}

func (o *RepoDeleteReleaseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRepoDeleteReleaseUnprocessableEntity creates a RepoDeleteReleaseUnprocessableEntity with default headers values
func NewRepoDeleteReleaseUnprocessableEntity() *RepoDeleteReleaseUnprocessableEntity {
	return &RepoDeleteReleaseUnprocessableEntity{}
}

/*
RepoDeleteReleaseUnprocessableEntity describes a response with status code 422, with default header values.

APIValidationError is error format response related to input validation
*/
type RepoDeleteReleaseUnprocessableEntity struct {
	Message string
	URL     string
}

// IsSuccess returns true when this repo delete release unprocessable entity response has a 2xx status code
func (o *RepoDeleteReleaseUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo delete release unprocessable entity response has a 3xx status code
func (o *RepoDeleteReleaseUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo delete release unprocessable entity response has a 4xx status code
func (o *RepoDeleteReleaseUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo delete release unprocessable entity response has a 5xx status code
func (o *RepoDeleteReleaseUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this repo delete release unprocessable entity response a status code equal to that given
func (o *RepoDeleteReleaseUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the repo delete release unprocessable entity response
func (o *RepoDeleteReleaseUnprocessableEntity) Code() int {
	return 422
}

func (o *RepoDeleteReleaseUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/releases/{id}][%d] repoDeleteReleaseUnprocessableEntity", 422)
}

func (o *RepoDeleteReleaseUnprocessableEntity) String() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/releases/{id}][%d] repoDeleteReleaseUnprocessableEntity", 422)
}

func (o *RepoDeleteReleaseUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
