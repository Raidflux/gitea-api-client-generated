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

// OrgListReposReader is a Reader for the OrgListRepos structure.
type OrgListReposReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrgListReposReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrgListReposOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewOrgListReposNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /orgs/{org}/repos] orgListRepos", response, response.Code())
	}
}

// NewOrgListReposOK creates a OrgListReposOK with default headers values
func NewOrgListReposOK() *OrgListReposOK {
	return &OrgListReposOK{}
}

/*
OrgListReposOK describes a response with status code 200, with default header values.

RepositoryList
*/
type OrgListReposOK struct {
	Payload []*models.Repository
}

// IsSuccess returns true when this org list repos o k response has a 2xx status code
func (o *OrgListReposOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this org list repos o k response has a 3xx status code
func (o *OrgListReposOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this org list repos o k response has a 4xx status code
func (o *OrgListReposOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this org list repos o k response has a 5xx status code
func (o *OrgListReposOK) IsServerError() bool {
	return false
}

// IsCode returns true when this org list repos o k response a status code equal to that given
func (o *OrgListReposOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the org list repos o k response
func (o *OrgListReposOK) Code() int {
	return 200
}

func (o *OrgListReposOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{org}/repos][%d] orgListReposOK %s", 200, payload)
}

func (o *OrgListReposOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{org}/repos][%d] orgListReposOK %s", 200, payload)
}

func (o *OrgListReposOK) GetPayload() []*models.Repository {
	return o.Payload
}

func (o *OrgListReposOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrgListReposNotFound creates a OrgListReposNotFound with default headers values
func NewOrgListReposNotFound() *OrgListReposNotFound {
	return &OrgListReposNotFound{}
}

/*
OrgListReposNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type OrgListReposNotFound struct {
}

// IsSuccess returns true when this org list repos not found response has a 2xx status code
func (o *OrgListReposNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this org list repos not found response has a 3xx status code
func (o *OrgListReposNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this org list repos not found response has a 4xx status code
func (o *OrgListReposNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this org list repos not found response has a 5xx status code
func (o *OrgListReposNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this org list repos not found response a status code equal to that given
func (o *OrgListReposNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the org list repos not found response
func (o *OrgListReposNotFound) Code() int {
	return 404
}

func (o *OrgListReposNotFound) Error() string {
	return fmt.Sprintf("[GET /orgs/{org}/repos][%d] orgListReposNotFound", 404)
}

func (o *OrgListReposNotFound) String() string {
	return fmt.Sprintf("[GET /orgs/{org}/repos][%d] orgListReposNotFound", 404)
}

func (o *OrgListReposNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
