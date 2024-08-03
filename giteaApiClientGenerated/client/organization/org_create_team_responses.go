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

// OrgCreateTeamReader is a Reader for the OrgCreateTeam structure.
type OrgCreateTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrgCreateTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewOrgCreateTeamCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewOrgCreateTeamNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewOrgCreateTeamUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /orgs/{org}/teams] orgCreateTeam", response, response.Code())
	}
}

// NewOrgCreateTeamCreated creates a OrgCreateTeamCreated with default headers values
func NewOrgCreateTeamCreated() *OrgCreateTeamCreated {
	return &OrgCreateTeamCreated{}
}

/*
OrgCreateTeamCreated describes a response with status code 201, with default header values.

Team
*/
type OrgCreateTeamCreated struct {
	Payload *models.Team
}

// IsSuccess returns true when this org create team created response has a 2xx status code
func (o *OrgCreateTeamCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this org create team created response has a 3xx status code
func (o *OrgCreateTeamCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this org create team created response has a 4xx status code
func (o *OrgCreateTeamCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this org create team created response has a 5xx status code
func (o *OrgCreateTeamCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this org create team created response a status code equal to that given
func (o *OrgCreateTeamCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the org create team created response
func (o *OrgCreateTeamCreated) Code() int {
	return 201
}

func (o *OrgCreateTeamCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /orgs/{org}/teams][%d] orgCreateTeamCreated %s", 201, payload)
}

func (o *OrgCreateTeamCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /orgs/{org}/teams][%d] orgCreateTeamCreated %s", 201, payload)
}

func (o *OrgCreateTeamCreated) GetPayload() *models.Team {
	return o.Payload
}

func (o *OrgCreateTeamCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Team)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrgCreateTeamNotFound creates a OrgCreateTeamNotFound with default headers values
func NewOrgCreateTeamNotFound() *OrgCreateTeamNotFound {
	return &OrgCreateTeamNotFound{}
}

/*
OrgCreateTeamNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type OrgCreateTeamNotFound struct {
}

// IsSuccess returns true when this org create team not found response has a 2xx status code
func (o *OrgCreateTeamNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this org create team not found response has a 3xx status code
func (o *OrgCreateTeamNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this org create team not found response has a 4xx status code
func (o *OrgCreateTeamNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this org create team not found response has a 5xx status code
func (o *OrgCreateTeamNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this org create team not found response a status code equal to that given
func (o *OrgCreateTeamNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the org create team not found response
func (o *OrgCreateTeamNotFound) Code() int {
	return 404
}

func (o *OrgCreateTeamNotFound) Error() string {
	return fmt.Sprintf("[POST /orgs/{org}/teams][%d] orgCreateTeamNotFound", 404)
}

func (o *OrgCreateTeamNotFound) String() string {
	return fmt.Sprintf("[POST /orgs/{org}/teams][%d] orgCreateTeamNotFound", 404)
}

func (o *OrgCreateTeamNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOrgCreateTeamUnprocessableEntity creates a OrgCreateTeamUnprocessableEntity with default headers values
func NewOrgCreateTeamUnprocessableEntity() *OrgCreateTeamUnprocessableEntity {
	return &OrgCreateTeamUnprocessableEntity{}
}

/*
OrgCreateTeamUnprocessableEntity describes a response with status code 422, with default header values.

APIValidationError is error format response related to input validation
*/
type OrgCreateTeamUnprocessableEntity struct {
	Message string
	URL     string
}

// IsSuccess returns true when this org create team unprocessable entity response has a 2xx status code
func (o *OrgCreateTeamUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this org create team unprocessable entity response has a 3xx status code
func (o *OrgCreateTeamUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this org create team unprocessable entity response has a 4xx status code
func (o *OrgCreateTeamUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this org create team unprocessable entity response has a 5xx status code
func (o *OrgCreateTeamUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this org create team unprocessable entity response a status code equal to that given
func (o *OrgCreateTeamUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the org create team unprocessable entity response
func (o *OrgCreateTeamUnprocessableEntity) Code() int {
	return 422
}

func (o *OrgCreateTeamUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /orgs/{org}/teams][%d] orgCreateTeamUnprocessableEntity", 422)
}

func (o *OrgCreateTeamUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /orgs/{org}/teams][%d] orgCreateTeamUnprocessableEntity", 422)
}

func (o *OrgCreateTeamUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header message
	hdrMessage := response.GetHeader("message")

	if hdrMessage != "" {
		o.Message = hdrMessage
	}

	// hydrates response header url
	hdrURL := response.GetHeader("url")

	if hdrURL != "" {
		o.URL = hdrURL
	}

	return nil
}
