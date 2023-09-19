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

// CloseAccountReader is a Reader for the CloseAccount structure.
type CloseAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloseAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewCloseAccountNoContent()
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

// NewCloseAccountNoContent creates a CloseAccountNoContent with default headers values
func NewCloseAccountNoContent() *CloseAccountNoContent {
	return &CloseAccountNoContent{}
}

/* CloseAccountNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type CloseAccountNoContent struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the close account no content response
func (o *CloseAccountNoContent) Code() int {
	return 204
}

// IsSuccess returns true when this close account no content response has a 2xx status code
func (o *CloseAccountNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this close account no content response has a 3xx status code
func (o *CloseAccountNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this close account no content response has a 4xx status code
func (o *CloseAccountNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this close account no content response has a 5xx status code
func (o *CloseAccountNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this close account no content response a status code equal to that given
func (o *CloseAccountNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *CloseAccountNoContent) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/accounts/{accountId}][%d] closeAccountNoContent ", 204)
}

func (o *CloseAccountNoContent) String() string {
	return fmt.Sprintf("[DELETE /1.0/kb/accounts/{accountId}][%d] closeAccountNoContent ", 204)
}

func (o *CloseAccountNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCloseAccountBadRequest creates a CloseAccountBadRequest with default headers values
func NewCloseAccountBadRequest() *CloseAccountBadRequest {
	return &CloseAccountBadRequest{}
}

/* CloseAccountBadRequest describes a response with status code 400, with default header values.

Invalid account id supplied
*/
type CloseAccountBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the close account bad request response
func (o *CloseAccountBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this close account bad request response has a 2xx status code
func (o *CloseAccountBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this close account bad request response has a 3xx status code
func (o *CloseAccountBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this close account bad request response has a 4xx status code
func (o *CloseAccountBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this close account bad request response has a 5xx status code
func (o *CloseAccountBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this close account bad request response a status code equal to that given
func (o *CloseAccountBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CloseAccountBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/accounts/{accountId}][%d] closeAccountBadRequest ", 400)
}

func (o *CloseAccountBadRequest) String() string {
	return fmt.Sprintf("[DELETE /1.0/kb/accounts/{accountId}][%d] closeAccountBadRequest ", 400)
}

func (o *CloseAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
