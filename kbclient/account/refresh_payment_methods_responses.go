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

// RefreshPaymentMethodsReader is a Reader for the RefreshPaymentMethods structure.
type RefreshPaymentMethodsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RefreshPaymentMethodsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewRefreshPaymentMethodsNoContent()
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

// NewRefreshPaymentMethodsNoContent creates a RefreshPaymentMethodsNoContent with default headers values
func NewRefreshPaymentMethodsNoContent() *RefreshPaymentMethodsNoContent {
	return &RefreshPaymentMethodsNoContent{}
}

/* RefreshPaymentMethodsNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type RefreshPaymentMethodsNoContent struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the refresh payment methods no content response
func (o *RefreshPaymentMethodsNoContent) Code() int {
	return 204
}

// IsSuccess returns true when this refresh payment methods no content response has a 2xx status code
func (o *RefreshPaymentMethodsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this refresh payment methods no content response has a 3xx status code
func (o *RefreshPaymentMethodsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this refresh payment methods no content response has a 4xx status code
func (o *RefreshPaymentMethodsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this refresh payment methods no content response has a 5xx status code
func (o *RefreshPaymentMethodsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this refresh payment methods no content response a status code equal to that given
func (o *RefreshPaymentMethodsNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *RefreshPaymentMethodsNoContent) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/accounts/{accountId}/paymentMethods/refresh][%d] refreshPaymentMethodsNoContent ", 204)
}

func (o *RefreshPaymentMethodsNoContent) String() string {
	return fmt.Sprintf("[PUT /1.0/kb/accounts/{accountId}/paymentMethods/refresh][%d] refreshPaymentMethodsNoContent ", 204)
}

func (o *RefreshPaymentMethodsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRefreshPaymentMethodsBadRequest creates a RefreshPaymentMethodsBadRequest with default headers values
func NewRefreshPaymentMethodsBadRequest() *RefreshPaymentMethodsBadRequest {
	return &RefreshPaymentMethodsBadRequest{}
}

/* RefreshPaymentMethodsBadRequest describes a response with status code 400, with default header values.

Invalid account id supplied
*/
type RefreshPaymentMethodsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the refresh payment methods bad request response
func (o *RefreshPaymentMethodsBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this refresh payment methods bad request response has a 2xx status code
func (o *RefreshPaymentMethodsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this refresh payment methods bad request response has a 3xx status code
func (o *RefreshPaymentMethodsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this refresh payment methods bad request response has a 4xx status code
func (o *RefreshPaymentMethodsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this refresh payment methods bad request response has a 5xx status code
func (o *RefreshPaymentMethodsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this refresh payment methods bad request response a status code equal to that given
func (o *RefreshPaymentMethodsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *RefreshPaymentMethodsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/accounts/{accountId}/paymentMethods/refresh][%d] refreshPaymentMethodsBadRequest ", 400)
}

func (o *RefreshPaymentMethodsBadRequest) String() string {
	return fmt.Sprintf("[PUT /1.0/kb/accounts/{accountId}/paymentMethods/refresh][%d] refreshPaymentMethodsBadRequest ", 400)
}

func (o *RefreshPaymentMethodsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRefreshPaymentMethodsNotFound creates a RefreshPaymentMethodsNotFound with default headers values
func NewRefreshPaymentMethodsNotFound() *RefreshPaymentMethodsNotFound {
	return &RefreshPaymentMethodsNotFound{}
}

/* RefreshPaymentMethodsNotFound describes a response with status code 404, with default header values.

Account not found
*/
type RefreshPaymentMethodsNotFound struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the refresh payment methods not found response
func (o *RefreshPaymentMethodsNotFound) Code() int {
	return 404
}

// IsSuccess returns true when this refresh payment methods not found response has a 2xx status code
func (o *RefreshPaymentMethodsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this refresh payment methods not found response has a 3xx status code
func (o *RefreshPaymentMethodsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this refresh payment methods not found response has a 4xx status code
func (o *RefreshPaymentMethodsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this refresh payment methods not found response has a 5xx status code
func (o *RefreshPaymentMethodsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this refresh payment methods not found response a status code equal to that given
func (o *RefreshPaymentMethodsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *RefreshPaymentMethodsNotFound) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/accounts/{accountId}/paymentMethods/refresh][%d] refreshPaymentMethodsNotFound ", 404)
}

func (o *RefreshPaymentMethodsNotFound) String() string {
	return fmt.Sprintf("[PUT /1.0/kb/accounts/{accountId}/paymentMethods/refresh][%d] refreshPaymentMethodsNotFound ", 404)
}

func (o *RefreshPaymentMethodsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
