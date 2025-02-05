// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v3/kbcommon"
)

// UpdateAccountReader is a Reader for the UpdateAccount structure.
type UpdateAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateAccountNoContent()
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

// NewUpdateAccountNoContent creates a UpdateAccountNoContent with default headers values
func NewUpdateAccountNoContent() *UpdateAccountNoContent {
	return &UpdateAccountNoContent{}
}

/* UpdateAccountNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type UpdateAccountNoContent struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the update account no content response
func (o *UpdateAccountNoContent) Code() int {
	return 204
}

// IsSuccess returns true when this update account no content response has a 2xx status code
func (o *UpdateAccountNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update account no content response has a 3xx status code
func (o *UpdateAccountNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update account no content response has a 4xx status code
func (o *UpdateAccountNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update account no content response has a 5xx status code
func (o *UpdateAccountNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update account no content response a status code equal to that given
func (o *UpdateAccountNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *UpdateAccountNoContent) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/accounts/{accountId}][%d] updateAccountNoContent ", 204)
}

func (o *UpdateAccountNoContent) String() string {
	return fmt.Sprintf("[PUT /1.0/kb/accounts/{accountId}][%d] updateAccountNoContent ", 204)
}

func (o *UpdateAccountNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAccountBadRequest creates a UpdateAccountBadRequest with default headers values
func NewUpdateAccountBadRequest() *UpdateAccountBadRequest {
	return &UpdateAccountBadRequest{}
}

/* UpdateAccountBadRequest describes a response with status code 400, with default header values.

Invalid account data supplied
*/
type UpdateAccountBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the update account bad request response
func (o *UpdateAccountBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this update account bad request response has a 2xx status code
func (o *UpdateAccountBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update account bad request response has a 3xx status code
func (o *UpdateAccountBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update account bad request response has a 4xx status code
func (o *UpdateAccountBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update account bad request response has a 5xx status code
func (o *UpdateAccountBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update account bad request response a status code equal to that given
func (o *UpdateAccountBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UpdateAccountBadRequest) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/accounts/{accountId}][%d] updateAccountBadRequest ", 400)
}

func (o *UpdateAccountBadRequest) String() string {
	return fmt.Sprintf("[PUT /1.0/kb/accounts/{accountId}][%d] updateAccountBadRequest ", 400)
}

func (o *UpdateAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
