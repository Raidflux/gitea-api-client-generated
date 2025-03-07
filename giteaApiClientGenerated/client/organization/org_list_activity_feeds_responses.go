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

// OrgListActivityFeedsReader is a Reader for the OrgListActivityFeeds structure.
type OrgListActivityFeedsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrgListActivityFeedsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrgListActivityFeedsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewOrgListActivityFeedsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /orgs/{org}/activities/feeds] orgListActivityFeeds", response, response.Code())
	}
}

// NewOrgListActivityFeedsOK creates a OrgListActivityFeedsOK with default headers values
func NewOrgListActivityFeedsOK() *OrgListActivityFeedsOK {
	return &OrgListActivityFeedsOK{}
}

/*
OrgListActivityFeedsOK describes a response with status code 200, with default header values.

ActivityFeedsList
*/
type OrgListActivityFeedsOK struct {
	Payload []*models.Activity
}

// IsSuccess returns true when this org list activity feeds o k response has a 2xx status code
func (o *OrgListActivityFeedsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this org list activity feeds o k response has a 3xx status code
func (o *OrgListActivityFeedsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this org list activity feeds o k response has a 4xx status code
func (o *OrgListActivityFeedsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this org list activity feeds o k response has a 5xx status code
func (o *OrgListActivityFeedsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this org list activity feeds o k response a status code equal to that given
func (o *OrgListActivityFeedsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the org list activity feeds o k response
func (o *OrgListActivityFeedsOK) Code() int {
	return 200
}

func (o *OrgListActivityFeedsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{org}/activities/feeds][%d] orgListActivityFeedsOK %s", 200, payload)
}

func (o *OrgListActivityFeedsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{org}/activities/feeds][%d] orgListActivityFeedsOK %s", 200, payload)
}

func (o *OrgListActivityFeedsOK) GetPayload() []*models.Activity {
	return o.Payload
}

func (o *OrgListActivityFeedsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrgListActivityFeedsNotFound creates a OrgListActivityFeedsNotFound with default headers values
func NewOrgListActivityFeedsNotFound() *OrgListActivityFeedsNotFound {
	return &OrgListActivityFeedsNotFound{}
}

/*
OrgListActivityFeedsNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type OrgListActivityFeedsNotFound struct {
}

// IsSuccess returns true when this org list activity feeds not found response has a 2xx status code
func (o *OrgListActivityFeedsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this org list activity feeds not found response has a 3xx status code
func (o *OrgListActivityFeedsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this org list activity feeds not found response has a 4xx status code
func (o *OrgListActivityFeedsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this org list activity feeds not found response has a 5xx status code
func (o *OrgListActivityFeedsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this org list activity feeds not found response a status code equal to that given
func (o *OrgListActivityFeedsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the org list activity feeds not found response
func (o *OrgListActivityFeedsNotFound) Code() int {
	return 404
}

func (o *OrgListActivityFeedsNotFound) Error() string {
	return fmt.Sprintf("[GET /orgs/{org}/activities/feeds][%d] orgListActivityFeedsNotFound", 404)
}

func (o *OrgListActivityFeedsNotFound) String() string {
	return fmt.Sprintf("[GET /orgs/{org}/activities/feeds][%d] orgListActivityFeedsNotFound", 404)
}

func (o *OrgListActivityFeedsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
