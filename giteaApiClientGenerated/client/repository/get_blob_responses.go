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

// GetBlobReader is a Reader for the GetBlob structure.
type GetBlobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBlobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBlobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetBlobBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetBlobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /repos/{owner}/{repo}/git/blobs/{sha}] GetBlob", response, response.Code())
	}
}

// NewGetBlobOK creates a GetBlobOK with default headers values
func NewGetBlobOK() *GetBlobOK {
	return &GetBlobOK{}
}

/*
GetBlobOK describes a response with status code 200, with default header values.

GitBlobResponse
*/
type GetBlobOK struct {
	Payload *models.GitBlobResponse
}

// IsSuccess returns true when this get blob o k response has a 2xx status code
func (o *GetBlobOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get blob o k response has a 3xx status code
func (o *GetBlobOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get blob o k response has a 4xx status code
func (o *GetBlobOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get blob o k response has a 5xx status code
func (o *GetBlobOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get blob o k response a status code equal to that given
func (o *GetBlobOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get blob o k response
func (o *GetBlobOK) Code() int {
	return 200
}

func (o *GetBlobOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/git/blobs/{sha}][%d] getBlobOK %s", 200, payload)
}

func (o *GetBlobOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/git/blobs/{sha}][%d] getBlobOK %s", 200, payload)
}

func (o *GetBlobOK) GetPayload() *models.GitBlobResponse {
	return o.Payload
}

func (o *GetBlobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GitBlobResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBlobBadRequest creates a GetBlobBadRequest with default headers values
func NewGetBlobBadRequest() *GetBlobBadRequest {
	return &GetBlobBadRequest{}
}

/*
GetBlobBadRequest describes a response with status code 400, with default header values.

APIError is error format response
*/
type GetBlobBadRequest struct {
	Message string
	URL     string
}

// IsSuccess returns true when this get blob bad request response has a 2xx status code
func (o *GetBlobBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get blob bad request response has a 3xx status code
func (o *GetBlobBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get blob bad request response has a 4xx status code
func (o *GetBlobBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get blob bad request response has a 5xx status code
func (o *GetBlobBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get blob bad request response a status code equal to that given
func (o *GetBlobBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get blob bad request response
func (o *GetBlobBadRequest) Code() int {
	return 400
}

func (o *GetBlobBadRequest) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/git/blobs/{sha}][%d] getBlobBadRequest", 400)
}

func (o *GetBlobBadRequest) String() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/git/blobs/{sha}][%d] getBlobBadRequest", 400)
}

func (o *GetBlobBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetBlobNotFound creates a GetBlobNotFound with default headers values
func NewGetBlobNotFound() *GetBlobNotFound {
	return &GetBlobNotFound{}
}

/*
GetBlobNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type GetBlobNotFound struct {
}

// IsSuccess returns true when this get blob not found response has a 2xx status code
func (o *GetBlobNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get blob not found response has a 3xx status code
func (o *GetBlobNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get blob not found response has a 4xx status code
func (o *GetBlobNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get blob not found response has a 5xx status code
func (o *GetBlobNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get blob not found response a status code equal to that given
func (o *GetBlobNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get blob not found response
func (o *GetBlobNotFound) Code() int {
	return 404
}

func (o *GetBlobNotFound) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/git/blobs/{sha}][%d] getBlobNotFound", 404)
}

func (o *GetBlobNotFound) String() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/git/blobs/{sha}][%d] getBlobNotFound", 404)
}

func (o *GetBlobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
