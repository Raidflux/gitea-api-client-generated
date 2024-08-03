// Code generated by go-swagger; DO NOT EDIT.

package user

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

// UserListSubscriptionsReader is a Reader for the UserListSubscriptions structure.
type UserListSubscriptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserListSubscriptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserListSubscriptionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUserListSubscriptionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /users/{username}/subscriptions] userListSubscriptions", response, response.Code())
	}
}

// NewUserListSubscriptionsOK creates a UserListSubscriptionsOK with default headers values
func NewUserListSubscriptionsOK() *UserListSubscriptionsOK {
	return &UserListSubscriptionsOK{}
}

/*
UserListSubscriptionsOK describes a response with status code 200, with default header values.

RepositoryList
*/
type UserListSubscriptionsOK struct {
	Payload []*models.Repository
}

// IsSuccess returns true when this user list subscriptions o k response has a 2xx status code
func (o *UserListSubscriptionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user list subscriptions o k response has a 3xx status code
func (o *UserListSubscriptionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user list subscriptions o k response has a 4xx status code
func (o *UserListSubscriptionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user list subscriptions o k response has a 5xx status code
func (o *UserListSubscriptionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this user list subscriptions o k response a status code equal to that given
func (o *UserListSubscriptionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the user list subscriptions o k response
func (o *UserListSubscriptionsOK) Code() int {
	return 200
}

func (o *UserListSubscriptionsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/{username}/subscriptions][%d] userListSubscriptionsOK %s", 200, payload)
}

func (o *UserListSubscriptionsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/{username}/subscriptions][%d] userListSubscriptionsOK %s", 200, payload)
}

func (o *UserListSubscriptionsOK) GetPayload() []*models.Repository {
	return o.Payload
}

func (o *UserListSubscriptionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserListSubscriptionsNotFound creates a UserListSubscriptionsNotFound with default headers values
func NewUserListSubscriptionsNotFound() *UserListSubscriptionsNotFound {
	return &UserListSubscriptionsNotFound{}
}

/*
UserListSubscriptionsNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type UserListSubscriptionsNotFound struct {
}

// IsSuccess returns true when this user list subscriptions not found response has a 2xx status code
func (o *UserListSubscriptionsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user list subscriptions not found response has a 3xx status code
func (o *UserListSubscriptionsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user list subscriptions not found response has a 4xx status code
func (o *UserListSubscriptionsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this user list subscriptions not found response has a 5xx status code
func (o *UserListSubscriptionsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this user list subscriptions not found response a status code equal to that given
func (o *UserListSubscriptionsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the user list subscriptions not found response
func (o *UserListSubscriptionsNotFound) Code() int {
	return 404
}

func (o *UserListSubscriptionsNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{username}/subscriptions][%d] userListSubscriptionsNotFound", 404)
}

func (o *UserListSubscriptionsNotFound) String() string {
	return fmt.Sprintf("[GET /users/{username}/subscriptions][%d] userListSubscriptionsNotFound", 404)
}

func (o *UserListSubscriptionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
