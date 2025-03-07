// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UserGetRunnerRegistrationTokenReader is a Reader for the UserGetRunnerRegistrationToken structure.
type UserGetRunnerRegistrationTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserGetRunnerRegistrationTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserGetRunnerRegistrationTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /user/actions/runners/registration-token] userGetRunnerRegistrationToken", response, response.Code())
	}
}

// NewUserGetRunnerRegistrationTokenOK creates a UserGetRunnerRegistrationTokenOK with default headers values
func NewUserGetRunnerRegistrationTokenOK() *UserGetRunnerRegistrationTokenOK {
	return &UserGetRunnerRegistrationTokenOK{}
}

/*
UserGetRunnerRegistrationTokenOK describes a response with status code 200, with default header values.

RegistrationToken is response related to registration token
*/
type UserGetRunnerRegistrationTokenOK struct {
	Token string
}

// IsSuccess returns true when this user get runner registration token o k response has a 2xx status code
func (o *UserGetRunnerRegistrationTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user get runner registration token o k response has a 3xx status code
func (o *UserGetRunnerRegistrationTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user get runner registration token o k response has a 4xx status code
func (o *UserGetRunnerRegistrationTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user get runner registration token o k response has a 5xx status code
func (o *UserGetRunnerRegistrationTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this user get runner registration token o k response a status code equal to that given
func (o *UserGetRunnerRegistrationTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the user get runner registration token o k response
func (o *UserGetRunnerRegistrationTokenOK) Code() int {
	return 200
}

func (o *UserGetRunnerRegistrationTokenOK) Error() string {
	return fmt.Sprintf("[GET /user/actions/runners/registration-token][%d] userGetRunnerRegistrationTokenOK", 200)
}

func (o *UserGetRunnerRegistrationTokenOK) String() string {
	return fmt.Sprintf("[GET /user/actions/runners/registration-token][%d] userGetRunnerRegistrationTokenOK", 200)
}

func (o *UserGetRunnerRegistrationTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header token
	hdrToken := response.GetHeader("token")

	if hdrToken != "" {
		o.Token = hdrToken
	}

	return nil
}
