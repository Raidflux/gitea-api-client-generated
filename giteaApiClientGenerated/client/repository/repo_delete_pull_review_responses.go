// Code generated by go-swagger; DO NOT EDIT.

package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RepoDeletePullReviewReader is a Reader for the RepoDeletePullReview structure.
type RepoDeletePullReviewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoDeletePullReviewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRepoDeletePullReviewNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewRepoDeletePullReviewForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRepoDeletePullReviewNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /repos/{owner}/{repo}/pulls/{index}/reviews/{id}] repoDeletePullReview", response, response.Code())
	}
}

// NewRepoDeletePullReviewNoContent creates a RepoDeletePullReviewNoContent with default headers values
func NewRepoDeletePullReviewNoContent() *RepoDeletePullReviewNoContent {
	return &RepoDeletePullReviewNoContent{}
}

/*
RepoDeletePullReviewNoContent describes a response with status code 204, with default header values.

APIEmpty is an empty response
*/
type RepoDeletePullReviewNoContent struct {
}

// IsSuccess returns true when this repo delete pull review no content response has a 2xx status code
func (o *RepoDeletePullReviewNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repo delete pull review no content response has a 3xx status code
func (o *RepoDeletePullReviewNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo delete pull review no content response has a 4xx status code
func (o *RepoDeletePullReviewNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this repo delete pull review no content response has a 5xx status code
func (o *RepoDeletePullReviewNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this repo delete pull review no content response a status code equal to that given
func (o *RepoDeletePullReviewNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the repo delete pull review no content response
func (o *RepoDeletePullReviewNoContent) Code() int {
	return 204
}

func (o *RepoDeletePullReviewNoContent) Error() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/pulls/{index}/reviews/{id}][%d] repoDeletePullReviewNoContent", 204)
}

func (o *RepoDeletePullReviewNoContent) String() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/pulls/{index}/reviews/{id}][%d] repoDeletePullReviewNoContent", 204)
}

func (o *RepoDeletePullReviewNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRepoDeletePullReviewForbidden creates a RepoDeletePullReviewForbidden with default headers values
func NewRepoDeletePullReviewForbidden() *RepoDeletePullReviewForbidden {
	return &RepoDeletePullReviewForbidden{}
}

/*
RepoDeletePullReviewForbidden describes a response with status code 403, with default header values.

APIForbiddenError is a forbidden error response
*/
type RepoDeletePullReviewForbidden struct {
	Message string
	URL     string
}

// IsSuccess returns true when this repo delete pull review forbidden response has a 2xx status code
func (o *RepoDeletePullReviewForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo delete pull review forbidden response has a 3xx status code
func (o *RepoDeletePullReviewForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo delete pull review forbidden response has a 4xx status code
func (o *RepoDeletePullReviewForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo delete pull review forbidden response has a 5xx status code
func (o *RepoDeletePullReviewForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this repo delete pull review forbidden response a status code equal to that given
func (o *RepoDeletePullReviewForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the repo delete pull review forbidden response
func (o *RepoDeletePullReviewForbidden) Code() int {
	return 403
}

func (o *RepoDeletePullReviewForbidden) Error() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/pulls/{index}/reviews/{id}][%d] repoDeletePullReviewForbidden", 403)
}

func (o *RepoDeletePullReviewForbidden) String() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/pulls/{index}/reviews/{id}][%d] repoDeletePullReviewForbidden", 403)
}

func (o *RepoDeletePullReviewForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRepoDeletePullReviewNotFound creates a RepoDeletePullReviewNotFound with default headers values
func NewRepoDeletePullReviewNotFound() *RepoDeletePullReviewNotFound {
	return &RepoDeletePullReviewNotFound{}
}

/*
RepoDeletePullReviewNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type RepoDeletePullReviewNotFound struct {
}

// IsSuccess returns true when this repo delete pull review not found response has a 2xx status code
func (o *RepoDeletePullReviewNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo delete pull review not found response has a 3xx status code
func (o *RepoDeletePullReviewNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo delete pull review not found response has a 4xx status code
func (o *RepoDeletePullReviewNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo delete pull review not found response has a 5xx status code
func (o *RepoDeletePullReviewNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repo delete pull review not found response a status code equal to that given
func (o *RepoDeletePullReviewNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the repo delete pull review not found response
func (o *RepoDeletePullReviewNotFound) Code() int {
	return 404
}

func (o *RepoDeletePullReviewNotFound) Error() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/pulls/{index}/reviews/{id}][%d] repoDeletePullReviewNotFound", 404)
}

func (o *RepoDeletePullReviewNotFound) String() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/pulls/{index}/reviews/{id}][%d] repoDeletePullReviewNotFound", 404)
}

func (o *RepoDeletePullReviewNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
