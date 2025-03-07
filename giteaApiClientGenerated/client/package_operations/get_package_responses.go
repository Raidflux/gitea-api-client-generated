// Code generated by go-swagger; DO NOT EDIT.

package package_operations

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

// GetPackageReader is a Reader for the GetPackage structure.
type GetPackageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPackageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPackageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetPackageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /packages/{owner}/{type}/{name}/{version}] getPackage", response, response.Code())
	}
}

// NewGetPackageOK creates a GetPackageOK with default headers values
func NewGetPackageOK() *GetPackageOK {
	return &GetPackageOK{}
}

/*
GetPackageOK describes a response with status code 200, with default header values.

Package
*/
type GetPackageOK struct {
	Payload *models.Package
}

// IsSuccess returns true when this get package o k response has a 2xx status code
func (o *GetPackageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get package o k response has a 3xx status code
func (o *GetPackageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get package o k response has a 4xx status code
func (o *GetPackageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get package o k response has a 5xx status code
func (o *GetPackageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get package o k response a status code equal to that given
func (o *GetPackageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get package o k response
func (o *GetPackageOK) Code() int {
	return 200
}

func (o *GetPackageOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /packages/{owner}/{type}/{name}/{version}][%d] getPackageOK %s", 200, payload)
}

func (o *GetPackageOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /packages/{owner}/{type}/{name}/{version}][%d] getPackageOK %s", 200, payload)
}

func (o *GetPackageOK) GetPayload() *models.Package {
	return o.Payload
}

func (o *GetPackageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Package)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPackageNotFound creates a GetPackageNotFound with default headers values
func NewGetPackageNotFound() *GetPackageNotFound {
	return &GetPackageNotFound{}
}

/*
GetPackageNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type GetPackageNotFound struct {
}

// IsSuccess returns true when this get package not found response has a 2xx status code
func (o *GetPackageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get package not found response has a 3xx status code
func (o *GetPackageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get package not found response has a 4xx status code
func (o *GetPackageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get package not found response has a 5xx status code
func (o *GetPackageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get package not found response a status code equal to that given
func (o *GetPackageNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get package not found response
func (o *GetPackageNotFound) Code() int {
	return 404
}

func (o *GetPackageNotFound) Error() string {
	return fmt.Sprintf("[GET /packages/{owner}/{type}/{name}/{version}][%d] getPackageNotFound", 404)
}

func (o *GetPackageNotFound) String() string {
	return fmt.Sprintf("[GET /packages/{owner}/{type}/{name}/{version}][%d] getPackageNotFound", 404)
}

func (o *GetPackageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
