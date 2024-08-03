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

// OrgListTeamMemberReader is a Reader for the OrgListTeamMember structure.
type OrgListTeamMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrgListTeamMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrgListTeamMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewOrgListTeamMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /teams/{id}/members/{username}] orgListTeamMember", response, response.Code())
	}
}

// NewOrgListTeamMemberOK creates a OrgListTeamMemberOK with default headers values
func NewOrgListTeamMemberOK() *OrgListTeamMemberOK {
	return &OrgListTeamMemberOK{}
}

/*
OrgListTeamMemberOK describes a response with status code 200, with default header values.

User
*/
type OrgListTeamMemberOK struct {
	Payload *models.User
}

// IsSuccess returns true when this org list team member o k response has a 2xx status code
func (o *OrgListTeamMemberOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this org list team member o k response has a 3xx status code
func (o *OrgListTeamMemberOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this org list team member o k response has a 4xx status code
func (o *OrgListTeamMemberOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this org list team member o k response has a 5xx status code
func (o *OrgListTeamMemberOK) IsServerError() bool {
	return false
}

// IsCode returns true when this org list team member o k response a status code equal to that given
func (o *OrgListTeamMemberOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the org list team member o k response
func (o *OrgListTeamMemberOK) Code() int {
	return 200
}

func (o *OrgListTeamMemberOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /teams/{id}/members/{username}][%d] orgListTeamMemberOK %s", 200, payload)
}

func (o *OrgListTeamMemberOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /teams/{id}/members/{username}][%d] orgListTeamMemberOK %s", 200, payload)
}

func (o *OrgListTeamMemberOK) GetPayload() *models.User {
	return o.Payload
}

func (o *OrgListTeamMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrgListTeamMemberNotFound creates a OrgListTeamMemberNotFound with default headers values
func NewOrgListTeamMemberNotFound() *OrgListTeamMemberNotFound {
	return &OrgListTeamMemberNotFound{}
}

/*
OrgListTeamMemberNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type OrgListTeamMemberNotFound struct {
}

// IsSuccess returns true when this org list team member not found response has a 2xx status code
func (o *OrgListTeamMemberNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this org list team member not found response has a 3xx status code
func (o *OrgListTeamMemberNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this org list team member not found response has a 4xx status code
func (o *OrgListTeamMemberNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this org list team member not found response has a 5xx status code
func (o *OrgListTeamMemberNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this org list team member not found response a status code equal to that given
func (o *OrgListTeamMemberNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the org list team member not found response
func (o *OrgListTeamMemberNotFound) Code() int {
	return 404
}

func (o *OrgListTeamMemberNotFound) Error() string {
	return fmt.Sprintf("[GET /teams/{id}/members/{username}][%d] orgListTeamMemberNotFound", 404)
}

func (o *OrgListTeamMemberNotFound) String() string {
	return fmt.Sprintf("[GET /teams/{id}/members/{username}][%d] orgListTeamMemberNotFound", 404)
}

func (o *OrgListTeamMemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
