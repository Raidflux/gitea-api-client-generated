// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// OrgGetRunnerRegistrationTokenReader is a Reader for the OrgGetRunnerRegistrationToken structure.
type OrgGetRunnerRegistrationTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrgGetRunnerRegistrationTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrgGetRunnerRegistrationTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /orgs/{org}/actions/runners/registration-token] orgGetRunnerRegistrationToken", response, response.Code())
	}
}

// NewOrgGetRunnerRegistrationTokenOK creates a OrgGetRunnerRegistrationTokenOK with default headers values
func NewOrgGetRunnerRegistrationTokenOK() *OrgGetRunnerRegistrationTokenOK {
	return &OrgGetRunnerRegistrationTokenOK{}
}

/*
OrgGetRunnerRegistrationTokenOK describes a response with status code 200, with default header values.

RegistrationToken is response related to registration token
*/
type OrgGetRunnerRegistrationTokenOK struct {
	Token string
}

// IsSuccess returns true when this org get runner registration token o k response has a 2xx status code
func (o *OrgGetRunnerRegistrationTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this org get runner registration token o k response has a 3xx status code
func (o *OrgGetRunnerRegistrationTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this org get runner registration token o k response has a 4xx status code
func (o *OrgGetRunnerRegistrationTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this org get runner registration token o k response has a 5xx status code
func (o *OrgGetRunnerRegistrationTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this org get runner registration token o k response a status code equal to that given
func (o *OrgGetRunnerRegistrationTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the org get runner registration token o k response
func (o *OrgGetRunnerRegistrationTokenOK) Code() int {
	return 200
}

func (o *OrgGetRunnerRegistrationTokenOK) Error() string {
	return fmt.Sprintf("[GET /orgs/{org}/actions/runners/registration-token][%d] orgGetRunnerRegistrationTokenOK", 200)
}

func (o *OrgGetRunnerRegistrationTokenOK) String() string {
	return fmt.Sprintf("[GET /orgs/{org}/actions/runners/registration-token][%d] orgGetRunnerRegistrationTokenOK", 200)
}

func (o *OrgGetRunnerRegistrationTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header token
	hdrToken := response.GetHeader("token")

	if hdrToken != "" {
		o.Token = hdrToken
	}

	return nil
}
