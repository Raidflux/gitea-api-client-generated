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

// OrgListCurrentUserOrgsReader is a Reader for the OrgListCurrentUserOrgs structure.
type OrgListCurrentUserOrgsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrgListCurrentUserOrgsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrgListCurrentUserOrgsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewOrgListCurrentUserOrgsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /user/orgs] orgListCurrentUserOrgs", response, response.Code())
	}
}

// NewOrgListCurrentUserOrgsOK creates a OrgListCurrentUserOrgsOK with default headers values
func NewOrgListCurrentUserOrgsOK() *OrgListCurrentUserOrgsOK {
	return &OrgListCurrentUserOrgsOK{}
}

/*
OrgListCurrentUserOrgsOK describes a response with status code 200, with default header values.

OrganizationList
*/
type OrgListCurrentUserOrgsOK struct {
	Payload []*models.Organization
}

// IsSuccess returns true when this org list current user orgs o k response has a 2xx status code
func (o *OrgListCurrentUserOrgsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this org list current user orgs o k response has a 3xx status code
func (o *OrgListCurrentUserOrgsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this org list current user orgs o k response has a 4xx status code
func (o *OrgListCurrentUserOrgsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this org list current user orgs o k response has a 5xx status code
func (o *OrgListCurrentUserOrgsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this org list current user orgs o k response a status code equal to that given
func (o *OrgListCurrentUserOrgsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the org list current user orgs o k response
func (o *OrgListCurrentUserOrgsOK) Code() int {
	return 200
}

func (o *OrgListCurrentUserOrgsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /user/orgs][%d] orgListCurrentUserOrgsOK %s", 200, payload)
}

func (o *OrgListCurrentUserOrgsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /user/orgs][%d] orgListCurrentUserOrgsOK %s", 200, payload)
}

func (o *OrgListCurrentUserOrgsOK) GetPayload() []*models.Organization {
	return o.Payload
}

func (o *OrgListCurrentUserOrgsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrgListCurrentUserOrgsNotFound creates a OrgListCurrentUserOrgsNotFound with default headers values
func NewOrgListCurrentUserOrgsNotFound() *OrgListCurrentUserOrgsNotFound {
	return &OrgListCurrentUserOrgsNotFound{}
}

/*
OrgListCurrentUserOrgsNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type OrgListCurrentUserOrgsNotFound struct {
}

// IsSuccess returns true when this org list current user orgs not found response has a 2xx status code
func (o *OrgListCurrentUserOrgsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this org list current user orgs not found response has a 3xx status code
func (o *OrgListCurrentUserOrgsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this org list current user orgs not found response has a 4xx status code
func (o *OrgListCurrentUserOrgsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this org list current user orgs not found response has a 5xx status code
func (o *OrgListCurrentUserOrgsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this org list current user orgs not found response a status code equal to that given
func (o *OrgListCurrentUserOrgsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the org list current user orgs not found response
func (o *OrgListCurrentUserOrgsNotFound) Code() int {
	return 404
}

func (o *OrgListCurrentUserOrgsNotFound) Error() string {
	return fmt.Sprintf("[GET /user/orgs][%d] orgListCurrentUserOrgsNotFound", 404)
}

func (o *OrgListCurrentUserOrgsNotFound) String() string {
	return fmt.Sprintf("[GET /user/orgs][%d] orgListCurrentUserOrgsNotFound", 404)
}

func (o *OrgListCurrentUserOrgsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
