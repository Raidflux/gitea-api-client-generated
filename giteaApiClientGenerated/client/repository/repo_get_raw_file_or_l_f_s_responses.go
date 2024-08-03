// Code generated by go-swagger; DO NOT EDIT.

package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RepoGetRawFileOrLFSReader is a Reader for the RepoGetRawFileOrLFS structure.
type RepoGetRawFileOrLFSReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *RepoGetRawFileOrLFSReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoGetRawFileOrLFSOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRepoGetRawFileOrLFSNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /repos/{owner}/{repo}/media/{filepath}] repoGetRawFileOrLFS", response, response.Code())
	}
}

// NewRepoGetRawFileOrLFSOK creates a RepoGetRawFileOrLFSOK with default headers values
func NewRepoGetRawFileOrLFSOK(writer io.Writer) *RepoGetRawFileOrLFSOK {
	return &RepoGetRawFileOrLFSOK{

		Payload: writer,
	}
}

/*
RepoGetRawFileOrLFSOK describes a response with status code 200, with default header values.

Returns raw file content.
*/
type RepoGetRawFileOrLFSOK struct {
	Payload io.Writer
}

// IsSuccess returns true when this repo get raw file or l f s o k response has a 2xx status code
func (o *RepoGetRawFileOrLFSOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repo get raw file or l f s o k response has a 3xx status code
func (o *RepoGetRawFileOrLFSOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo get raw file or l f s o k response has a 4xx status code
func (o *RepoGetRawFileOrLFSOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this repo get raw file or l f s o k response has a 5xx status code
func (o *RepoGetRawFileOrLFSOK) IsServerError() bool {
	return false
}

// IsCode returns true when this repo get raw file or l f s o k response a status code equal to that given
func (o *RepoGetRawFileOrLFSOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the repo get raw file or l f s o k response
func (o *RepoGetRawFileOrLFSOK) Code() int {
	return 200
}

func (o *RepoGetRawFileOrLFSOK) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/media/{filepath}][%d] repoGetRawFileOrLFSOK", 200)
}

func (o *RepoGetRawFileOrLFSOK) String() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/media/{filepath}][%d] repoGetRawFileOrLFSOK", 200)
}

func (o *RepoGetRawFileOrLFSOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *RepoGetRawFileOrLFSOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepoGetRawFileOrLFSNotFound creates a RepoGetRawFileOrLFSNotFound with default headers values
func NewRepoGetRawFileOrLFSNotFound() *RepoGetRawFileOrLFSNotFound {
	return &RepoGetRawFileOrLFSNotFound{}
}

/*
RepoGetRawFileOrLFSNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type RepoGetRawFileOrLFSNotFound struct {
}

// IsSuccess returns true when this repo get raw file or l f s not found response has a 2xx status code
func (o *RepoGetRawFileOrLFSNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo get raw file or l f s not found response has a 3xx status code
func (o *RepoGetRawFileOrLFSNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo get raw file or l f s not found response has a 4xx status code
func (o *RepoGetRawFileOrLFSNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo get raw file or l f s not found response has a 5xx status code
func (o *RepoGetRawFileOrLFSNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repo get raw file or l f s not found response a status code equal to that given
func (o *RepoGetRawFileOrLFSNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the repo get raw file or l f s not found response
func (o *RepoGetRawFileOrLFSNotFound) Code() int {
	return 404
}

func (o *RepoGetRawFileOrLFSNotFound) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/media/{filepath}][%d] repoGetRawFileOrLFSNotFound", 404)
}

func (o *RepoGetRawFileOrLFSNotFound) String() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/media/{filepath}][%d] repoGetRawFileOrLFSNotFound", 404)
}

func (o *RepoGetRawFileOrLFSNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
