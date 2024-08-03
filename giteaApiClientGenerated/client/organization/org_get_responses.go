// Code generated by go-swagger; DO NOT EDIT.

package organization

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

// OrgGetReader is a Reader for the OrgGet structure.
type OrgGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrgGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrgGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewOrgGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /orgs/{org}] orgGet", response, response.Code())
	}
}

// NewOrgGetOK creates a OrgGetOK with default headers values
func NewOrgGetOK() *OrgGetOK {
	return &OrgGetOK{}
}

/*
OrgGetOK describes a response with status code 200, with default header values.

Organization
*/
type OrgGetOK struct {
	Payload *models.Organization
}

// IsSuccess returns true when this org get o k response has a 2xx status code
func (o *OrgGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this org get o k response has a 3xx status code
func (o *OrgGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this org get o k response has a 4xx status code
func (o *OrgGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this org get o k response has a 5xx status code
func (o *OrgGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this org get o k response a status code equal to that given
func (o *OrgGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the org get o k response
func (o *OrgGetOK) Code() int {
	return 200
}

func (o *OrgGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{org}][%d] orgGetOK %s", 200, payload)
}

func (o *OrgGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{org}][%d] orgGetOK %s", 200, payload)
}

func (o *OrgGetOK) GetPayload() *models.Organization {
	return o.Payload
}

func (o *OrgGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Organization)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrgGetNotFound creates a OrgGetNotFound with default headers values
func NewOrgGetNotFound() *OrgGetNotFound {
	return &OrgGetNotFound{}
}

/*
OrgGetNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type OrgGetNotFound struct {
}

// IsSuccess returns true when this org get not found response has a 2xx status code
func (o *OrgGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this org get not found response has a 3xx status code
func (o *OrgGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this org get not found response has a 4xx status code
func (o *OrgGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this org get not found response has a 5xx status code
func (o *OrgGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this org get not found response a status code equal to that given
func (o *OrgGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the org get not found response
func (o *OrgGetNotFound) Code() int {
	return 404
}

func (o *OrgGetNotFound) Error() string {
	return fmt.Sprintf("[GET /orgs/{org}][%d] orgGetNotFound", 404)
}

func (o *OrgGetNotFound) String() string {
	return fmt.Sprintf("[GET /orgs/{org}][%d] orgGetNotFound", 404)
}

func (o *OrgGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
