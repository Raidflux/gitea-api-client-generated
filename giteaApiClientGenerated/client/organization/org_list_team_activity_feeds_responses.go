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

// OrgListTeamActivityFeedsReader is a Reader for the OrgListTeamActivityFeeds structure.
type OrgListTeamActivityFeedsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrgListTeamActivityFeedsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrgListTeamActivityFeedsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewOrgListTeamActivityFeedsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /teams/{id}/activities/feeds] orgListTeamActivityFeeds", response, response.Code())
	}
}

// NewOrgListTeamActivityFeedsOK creates a OrgListTeamActivityFeedsOK with default headers values
func NewOrgListTeamActivityFeedsOK() *OrgListTeamActivityFeedsOK {
	return &OrgListTeamActivityFeedsOK{}
}

/*
OrgListTeamActivityFeedsOK describes a response with status code 200, with default header values.

ActivityFeedsList
*/
type OrgListTeamActivityFeedsOK struct {
	Payload []*models.Activity
}

// IsSuccess returns true when this org list team activity feeds o k response has a 2xx status code
func (o *OrgListTeamActivityFeedsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this org list team activity feeds o k response has a 3xx status code
func (o *OrgListTeamActivityFeedsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this org list team activity feeds o k response has a 4xx status code
func (o *OrgListTeamActivityFeedsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this org list team activity feeds o k response has a 5xx status code
func (o *OrgListTeamActivityFeedsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this org list team activity feeds o k response a status code equal to that given
func (o *OrgListTeamActivityFeedsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the org list team activity feeds o k response
func (o *OrgListTeamActivityFeedsOK) Code() int {
	return 200
}

func (o *OrgListTeamActivityFeedsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /teams/{id}/activities/feeds][%d] orgListTeamActivityFeedsOK %s", 200, payload)
}

func (o *OrgListTeamActivityFeedsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /teams/{id}/activities/feeds][%d] orgListTeamActivityFeedsOK %s", 200, payload)
}

func (o *OrgListTeamActivityFeedsOK) GetPayload() []*models.Activity {
	return o.Payload
}

func (o *OrgListTeamActivityFeedsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrgListTeamActivityFeedsNotFound creates a OrgListTeamActivityFeedsNotFound with default headers values
func NewOrgListTeamActivityFeedsNotFound() *OrgListTeamActivityFeedsNotFound {
	return &OrgListTeamActivityFeedsNotFound{}
}

/*
OrgListTeamActivityFeedsNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type OrgListTeamActivityFeedsNotFound struct {
}

// IsSuccess returns true when this org list team activity feeds not found response has a 2xx status code
func (o *OrgListTeamActivityFeedsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this org list team activity feeds not found response has a 3xx status code
func (o *OrgListTeamActivityFeedsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this org list team activity feeds not found response has a 4xx status code
func (o *OrgListTeamActivityFeedsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this org list team activity feeds not found response has a 5xx status code
func (o *OrgListTeamActivityFeedsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this org list team activity feeds not found response a status code equal to that given
func (o *OrgListTeamActivityFeedsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the org list team activity feeds not found response
func (o *OrgListTeamActivityFeedsNotFound) Code() int {
	return 404
}

func (o *OrgListTeamActivityFeedsNotFound) Error() string {
	return fmt.Sprintf("[GET /teams/{id}/activities/feeds][%d] orgListTeamActivityFeedsNotFound", 404)
}

func (o *OrgListTeamActivityFeedsNotFound) String() string {
	return fmt.Sprintf("[GET /teams/{id}/activities/feeds][%d] orgListTeamActivityFeedsNotFound", 404)
}

func (o *OrgListTeamActivityFeedsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
