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

// OrgGetAllReader is a Reader for the OrgGetAll structure.
type OrgGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrgGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrgGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /orgs] orgGetAll", response, response.Code())
	}
}

// NewOrgGetAllOK creates a OrgGetAllOK with default headers values
func NewOrgGetAllOK() *OrgGetAllOK {
	return &OrgGetAllOK{}
}

/*
OrgGetAllOK describes a response with status code 200, with default header values.

OrganizationList
*/
type OrgGetAllOK struct {
	Payload []*models.Organization
}

// IsSuccess returns true when this org get all o k response has a 2xx status code
func (o *OrgGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this org get all o k response has a 3xx status code
func (o *OrgGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this org get all o k response has a 4xx status code
func (o *OrgGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this org get all o k response has a 5xx status code
func (o *OrgGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this org get all o k response a status code equal to that given
func (o *OrgGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the org get all o k response
func (o *OrgGetAllOK) Code() int {
	return 200
}

func (o *OrgGetAllOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs][%d] orgGetAllOK %s", 200, payload)
}

func (o *OrgGetAllOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs][%d] orgGetAllOK %s", 200, payload)
}

func (o *OrgGetAllOK) GetPayload() []*models.Organization {
	return o.Payload
}

func (o *OrgGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
