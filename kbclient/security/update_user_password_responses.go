// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v3/kbcommon"
)

// UpdateUserPasswordReader is a Reader for the UpdateUserPassword structure.
type UpdateUserPasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserPasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateUserPasswordNoContent()
		result.HttpResponse = response
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		errorResult := kbcommon.NewKillbillError(response.Code())
		if err := consumer.Consume(response.Body(), &errorResult); err != nil && err != io.EOF {
			return nil, err
		}
		return nil, errorResult
	}
}

// NewUpdateUserPasswordNoContent creates a UpdateUserPasswordNoContent with default headers values
func NewUpdateUserPasswordNoContent() *UpdateUserPasswordNoContent {
	return &UpdateUserPasswordNoContent{}
}

/* UpdateUserPasswordNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type UpdateUserPasswordNoContent struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the update user password no content response
func (o *UpdateUserPasswordNoContent) Code() int {
	return 204
}

// IsSuccess returns true when this update user password no content response has a 2xx status code
func (o *UpdateUserPasswordNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update user password no content response has a 3xx status code
func (o *UpdateUserPasswordNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user password no content response has a 4xx status code
func (o *UpdateUserPasswordNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update user password no content response has a 5xx status code
func (o *UpdateUserPasswordNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update user password no content response a status code equal to that given
func (o *UpdateUserPasswordNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *UpdateUserPasswordNoContent) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/security/users/{username}/password][%d] updateUserPasswordNoContent ", 204)
}

func (o *UpdateUserPasswordNoContent) String() string {
	return fmt.Sprintf("[PUT /1.0/kb/security/users/{username}/password][%d] updateUserPasswordNoContent ", 204)
}

func (o *UpdateUserPasswordNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
