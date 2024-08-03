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

// OrgEditReader is a Reader for the OrgEdit structure.
type OrgEditReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrgEditReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrgEditOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewOrgEditNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /orgs/{org}] orgEdit", response, response.Code())
	}
}

// NewOrgEditOK creates a OrgEditOK with default headers values
func NewOrgEditOK() *OrgEditOK {
	return &OrgEditOK{}
}

/*
OrgEditOK describes a response with status code 200, with default header values.

Organization
*/
type OrgEditOK struct {
	Payload *models.Organization
}

// IsSuccess returns true when this org edit o k response has a 2xx status code
func (o *OrgEditOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this org edit o k response has a 3xx status code
func (o *OrgEditOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this org edit o k response has a 4xx status code
func (o *OrgEditOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this org edit o k response has a 5xx status code
func (o *OrgEditOK) IsServerError() bool {
	return false
}

// IsCode returns true when this org edit o k response a status code equal to that given
func (o *OrgEditOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the org edit o k response
func (o *OrgEditOK) Code() int {
	return 200
}

func (o *OrgEditOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /orgs/{org}][%d] orgEditOK %s", 200, payload)
}

func (o *OrgEditOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /orgs/{org}][%d] orgEditOK %s", 200, payload)
}

func (o *OrgEditOK) GetPayload() *models.Organization {
	return o.Payload
}

func (o *OrgEditOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Organization)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrgEditNotFound creates a OrgEditNotFound with default headers values
func NewOrgEditNotFound() *OrgEditNotFound {
	return &OrgEditNotFound{}
}

/*
OrgEditNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type OrgEditNotFound struct {
}

// IsSuccess returns true when this org edit not found response has a 2xx status code
func (o *OrgEditNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this org edit not found response has a 3xx status code
func (o *OrgEditNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this org edit not found response has a 4xx status code
func (o *OrgEditNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this org edit not found response has a 5xx status code
func (o *OrgEditNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this org edit not found response a status code equal to that given
func (o *OrgEditNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the org edit not found response
func (o *OrgEditNotFound) Code() int {
	return 404
}

func (o *OrgEditNotFound) Error() string {
	return fmt.Sprintf("[PATCH /orgs/{org}][%d] orgEditNotFound", 404)
}

func (o *OrgEditNotFound) String() string {
	return fmt.Sprintf("[PATCH /orgs/{org}][%d] orgEditNotFound", 404)
}

func (o *OrgEditNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
