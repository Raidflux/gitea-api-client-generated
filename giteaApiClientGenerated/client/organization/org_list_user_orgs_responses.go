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

	"giteaApiClientGenerated/models"
)

// OrgListUserOrgsReader is a Reader for the OrgListUserOrgs structure.
type OrgListUserOrgsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrgListUserOrgsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrgListUserOrgsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewOrgListUserOrgsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /users/{username}/orgs] orgListUserOrgs", response, response.Code())
	}
}

// NewOrgListUserOrgsOK creates a OrgListUserOrgsOK with default headers values
func NewOrgListUserOrgsOK() *OrgListUserOrgsOK {
	return &OrgListUserOrgsOK{}
}

/*
OrgListUserOrgsOK describes a response with status code 200, with default header values.

OrganizationList
*/
type OrgListUserOrgsOK struct {
	Payload []*models.Organization
}

// IsSuccess returns true when this org list user orgs o k response has a 2xx status code
func (o *OrgListUserOrgsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this org list user orgs o k response has a 3xx status code
func (o *OrgListUserOrgsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this org list user orgs o k response has a 4xx status code
func (o *OrgListUserOrgsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this org list user orgs o k response has a 5xx status code
func (o *OrgListUserOrgsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this org list user orgs o k response a status code equal to that given
func (o *OrgListUserOrgsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the org list user orgs o k response
func (o *OrgListUserOrgsOK) Code() int {
	return 200
}

func (o *OrgListUserOrgsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/{username}/orgs][%d] orgListUserOrgsOK %s", 200, payload)
}

func (o *OrgListUserOrgsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/{username}/orgs][%d] orgListUserOrgsOK %s", 200, payload)
}

func (o *OrgListUserOrgsOK) GetPayload() []*models.Organization {
	return o.Payload
}

func (o *OrgListUserOrgsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrgListUserOrgsNotFound creates a OrgListUserOrgsNotFound with default headers values
func NewOrgListUserOrgsNotFound() *OrgListUserOrgsNotFound {
	return &OrgListUserOrgsNotFound{}
}

/*
OrgListUserOrgsNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type OrgListUserOrgsNotFound struct {
}

// IsSuccess returns true when this org list user orgs not found response has a 2xx status code
func (o *OrgListUserOrgsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this org list user orgs not found response has a 3xx status code
func (o *OrgListUserOrgsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this org list user orgs not found response has a 4xx status code
func (o *OrgListUserOrgsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this org list user orgs not found response has a 5xx status code
func (o *OrgListUserOrgsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this org list user orgs not found response a status code equal to that given
func (o *OrgListUserOrgsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the org list user orgs not found response
func (o *OrgListUserOrgsNotFound) Code() int {
	return 404
}

func (o *OrgListUserOrgsNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{username}/orgs][%d] orgListUserOrgsNotFound", 404)
}

func (o *OrgListUserOrgsNotFound) String() string {
	return fmt.Sprintf("[GET /users/{username}/orgs][%d] orgListUserOrgsNotFound", 404)
}

func (o *OrgListUserOrgsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
