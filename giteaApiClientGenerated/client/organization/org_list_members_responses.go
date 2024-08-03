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

// OrgListMembersReader is a Reader for the OrgListMembers structure.
type OrgListMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrgListMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrgListMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewOrgListMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /orgs/{org}/members] orgListMembers", response, response.Code())
	}
}

// NewOrgListMembersOK creates a OrgListMembersOK with default headers values
func NewOrgListMembersOK() *OrgListMembersOK {
	return &OrgListMembersOK{}
}

/*
OrgListMembersOK describes a response with status code 200, with default header values.

UserList
*/
type OrgListMembersOK struct {
	Payload []*models.User
}

// IsSuccess returns true when this org list members o k response has a 2xx status code
func (o *OrgListMembersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this org list members o k response has a 3xx status code
func (o *OrgListMembersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this org list members o k response has a 4xx status code
func (o *OrgListMembersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this org list members o k response has a 5xx status code
func (o *OrgListMembersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this org list members o k response a status code equal to that given
func (o *OrgListMembersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the org list members o k response
func (o *OrgListMembersOK) Code() int {
	return 200
}

func (o *OrgListMembersOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{org}/members][%d] orgListMembersOK %s", 200, payload)
}

func (o *OrgListMembersOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{org}/members][%d] orgListMembersOK %s", 200, payload)
}

func (o *OrgListMembersOK) GetPayload() []*models.User {
	return o.Payload
}

func (o *OrgListMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrgListMembersNotFound creates a OrgListMembersNotFound with default headers values
func NewOrgListMembersNotFound() *OrgListMembersNotFound {
	return &OrgListMembersNotFound{}
}

/*
OrgListMembersNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type OrgListMembersNotFound struct {
}

// IsSuccess returns true when this org list members not found response has a 2xx status code
func (o *OrgListMembersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this org list members not found response has a 3xx status code
func (o *OrgListMembersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this org list members not found response has a 4xx status code
func (o *OrgListMembersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this org list members not found response has a 5xx status code
func (o *OrgListMembersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this org list members not found response a status code equal to that given
func (o *OrgListMembersNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the org list members not found response
func (o *OrgListMembersNotFound) Code() int {
	return 404
}

func (o *OrgListMembersNotFound) Error() string {
	return fmt.Sprintf("[GET /orgs/{org}/members][%d] orgListMembersNotFound", 404)
}

func (o *OrgListMembersNotFound) String() string {
	return fmt.Sprintf("[GET /orgs/{org}/members][%d] orgListMembersNotFound", 404)
}

func (o *OrgListMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
