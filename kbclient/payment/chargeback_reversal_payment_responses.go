// Code generated by go-swagger; DO NOT EDIT.

package payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v3/kbcommon"

	"github.com/killbill/kbcli/v3/kbmodel"
)

// ChargebackReversalPaymentReader is a Reader for the ChargebackReversalPayment structure.
type ChargebackReversalPaymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChargebackReversalPaymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewChargebackReversalPaymentCreated()
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

// NewChargebackReversalPaymentCreated creates a ChargebackReversalPaymentCreated with default headers values
func NewChargebackReversalPaymentCreated() *ChargebackReversalPaymentCreated {
	return &ChargebackReversalPaymentCreated{}
}

/* ChargebackReversalPaymentCreated describes a response with status code 201, with default header values.

Payment transaction created successfully
*/
type ChargebackReversalPaymentCreated struct {
	Payload      *kbmodel.Payment
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the chargeback reversal payment created response
func (o *ChargebackReversalPaymentCreated) Code() int {
	return 201
}

// IsSuccess returns true when this chargeback reversal payment created response has a 2xx status code
func (o *ChargebackReversalPaymentCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this chargeback reversal payment created response has a 3xx status code
func (o *ChargebackReversalPaymentCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this chargeback reversal payment created response has a 4xx status code
func (o *ChargebackReversalPaymentCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this chargeback reversal payment created response has a 5xx status code
func (o *ChargebackReversalPaymentCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this chargeback reversal payment created response a status code equal to that given
func (o *ChargebackReversalPaymentCreated) IsCode(code int) bool {
	return code == 201
}

func (o *ChargebackReversalPaymentCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}/chargebackReversals][%d] chargebackReversalPaymentCreated  %+v", 201, o.Payload)
}

func (o *ChargebackReversalPaymentCreated) String() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}/chargebackReversals][%d] chargebackReversalPaymentCreated  %+v", 201, o.Payload)
}

func (o *ChargebackReversalPaymentCreated) GetPayload() *kbmodel.Payment {
	return o.Payload
}

func (o *ChargebackReversalPaymentCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Payment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChargebackReversalPaymentBadRequest creates a ChargebackReversalPaymentBadRequest with default headers values
func NewChargebackReversalPaymentBadRequest() *ChargebackReversalPaymentBadRequest {
	return &ChargebackReversalPaymentBadRequest{}
}

/* ChargebackReversalPaymentBadRequest describes a response with status code 400, with default header values.

Invalid paymentId supplied
*/
type ChargebackReversalPaymentBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the chargeback reversal payment bad request response
func (o *ChargebackReversalPaymentBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this chargeback reversal payment bad request response has a 2xx status code
func (o *ChargebackReversalPaymentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this chargeback reversal payment bad request response has a 3xx status code
func (o *ChargebackReversalPaymentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this chargeback reversal payment bad request response has a 4xx status code
func (o *ChargebackReversalPaymentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this chargeback reversal payment bad request response has a 5xx status code
func (o *ChargebackReversalPaymentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this chargeback reversal payment bad request response a status code equal to that given
func (o *ChargebackReversalPaymentBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ChargebackReversalPaymentBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}/chargebackReversals][%d] chargebackReversalPaymentBadRequest ", 400)
}

func (o *ChargebackReversalPaymentBadRequest) String() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}/chargebackReversals][%d] chargebackReversalPaymentBadRequest ", 400)
}

func (o *ChargebackReversalPaymentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChargebackReversalPaymentPaymentRequired creates a ChargebackReversalPaymentPaymentRequired with default headers values
func NewChargebackReversalPaymentPaymentRequired() *ChargebackReversalPaymentPaymentRequired {
	return &ChargebackReversalPaymentPaymentRequired{}
}

/* ChargebackReversalPaymentPaymentRequired describes a response with status code 402, with default header values.

Transaction declined by gateway
*/
type ChargebackReversalPaymentPaymentRequired struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the chargeback reversal payment payment required response
func (o *ChargebackReversalPaymentPaymentRequired) Code() int {
	return 402
}

// IsSuccess returns true when this chargeback reversal payment payment required response has a 2xx status code
func (o *ChargebackReversalPaymentPaymentRequired) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this chargeback reversal payment payment required response has a 3xx status code
func (o *ChargebackReversalPaymentPaymentRequired) IsRedirect() bool {
	return false
}

// IsClientError returns true when this chargeback reversal payment payment required response has a 4xx status code
func (o *ChargebackReversalPaymentPaymentRequired) IsClientError() bool {
	return true
}

// IsServerError returns true when this chargeback reversal payment payment required response has a 5xx status code
func (o *ChargebackReversalPaymentPaymentRequired) IsServerError() bool {
	return false
}

// IsCode returns true when this chargeback reversal payment payment required response a status code equal to that given
func (o *ChargebackReversalPaymentPaymentRequired) IsCode(code int) bool {
	return code == 402
}

func (o *ChargebackReversalPaymentPaymentRequired) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}/chargebackReversals][%d] chargebackReversalPaymentPaymentRequired ", 402)
}

func (o *ChargebackReversalPaymentPaymentRequired) String() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}/chargebackReversals][%d] chargebackReversalPaymentPaymentRequired ", 402)
}

func (o *ChargebackReversalPaymentPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChargebackReversalPaymentNotFound creates a ChargebackReversalPaymentNotFound with default headers values
func NewChargebackReversalPaymentNotFound() *ChargebackReversalPaymentNotFound {
	return &ChargebackReversalPaymentNotFound{}
}

/* ChargebackReversalPaymentNotFound describes a response with status code 404, with default header values.

Account or payment not found
*/
type ChargebackReversalPaymentNotFound struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the chargeback reversal payment not found response
func (o *ChargebackReversalPaymentNotFound) Code() int {
	return 404
}

// IsSuccess returns true when this chargeback reversal payment not found response has a 2xx status code
func (o *ChargebackReversalPaymentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this chargeback reversal payment not found response has a 3xx status code
func (o *ChargebackReversalPaymentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this chargeback reversal payment not found response has a 4xx status code
func (o *ChargebackReversalPaymentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this chargeback reversal payment not found response has a 5xx status code
func (o *ChargebackReversalPaymentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this chargeback reversal payment not found response a status code equal to that given
func (o *ChargebackReversalPaymentNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ChargebackReversalPaymentNotFound) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}/chargebackReversals][%d] chargebackReversalPaymentNotFound ", 404)
}

func (o *ChargebackReversalPaymentNotFound) String() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}/chargebackReversals][%d] chargebackReversalPaymentNotFound ", 404)
}

func (o *ChargebackReversalPaymentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChargebackReversalPaymentUnprocessableEntity creates a ChargebackReversalPaymentUnprocessableEntity with default headers values
func NewChargebackReversalPaymentUnprocessableEntity() *ChargebackReversalPaymentUnprocessableEntity {
	return &ChargebackReversalPaymentUnprocessableEntity{}
}

/* ChargebackReversalPaymentUnprocessableEntity describes a response with status code 422, with default header values.

Payment is aborted by a control plugin
*/
type ChargebackReversalPaymentUnprocessableEntity struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the chargeback reversal payment unprocessable entity response
func (o *ChargebackReversalPaymentUnprocessableEntity) Code() int {
	return 422
}

// IsSuccess returns true when this chargeback reversal payment unprocessable entity response has a 2xx status code
func (o *ChargebackReversalPaymentUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this chargeback reversal payment unprocessable entity response has a 3xx status code
func (o *ChargebackReversalPaymentUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this chargeback reversal payment unprocessable entity response has a 4xx status code
func (o *ChargebackReversalPaymentUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this chargeback reversal payment unprocessable entity response has a 5xx status code
func (o *ChargebackReversalPaymentUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this chargeback reversal payment unprocessable entity response a status code equal to that given
func (o *ChargebackReversalPaymentUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *ChargebackReversalPaymentUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}/chargebackReversals][%d] chargebackReversalPaymentUnprocessableEntity ", 422)
}

func (o *ChargebackReversalPaymentUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}/chargebackReversals][%d] chargebackReversalPaymentUnprocessableEntity ", 422)
}

func (o *ChargebackReversalPaymentUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChargebackReversalPaymentBadGateway creates a ChargebackReversalPaymentBadGateway with default headers values
func NewChargebackReversalPaymentBadGateway() *ChargebackReversalPaymentBadGateway {
	return &ChargebackReversalPaymentBadGateway{}
}

/* ChargebackReversalPaymentBadGateway describes a response with status code 502, with default header values.

Failed to submit payment transaction
*/
type ChargebackReversalPaymentBadGateway struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the chargeback reversal payment bad gateway response
func (o *ChargebackReversalPaymentBadGateway) Code() int {
	return 502
}

// IsSuccess returns true when this chargeback reversal payment bad gateway response has a 2xx status code
func (o *ChargebackReversalPaymentBadGateway) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this chargeback reversal payment bad gateway response has a 3xx status code
func (o *ChargebackReversalPaymentBadGateway) IsRedirect() bool {
	return false
}

// IsClientError returns true when this chargeback reversal payment bad gateway response has a 4xx status code
func (o *ChargebackReversalPaymentBadGateway) IsClientError() bool {
	return false
}

// IsServerError returns true when this chargeback reversal payment bad gateway response has a 5xx status code
func (o *ChargebackReversalPaymentBadGateway) IsServerError() bool {
	return true
}

// IsCode returns true when this chargeback reversal payment bad gateway response a status code equal to that given
func (o *ChargebackReversalPaymentBadGateway) IsCode(code int) bool {
	return code == 502
}

func (o *ChargebackReversalPaymentBadGateway) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}/chargebackReversals][%d] chargebackReversalPaymentBadGateway ", 502)
}

func (o *ChargebackReversalPaymentBadGateway) String() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}/chargebackReversals][%d] chargebackReversalPaymentBadGateway ", 502)
}

func (o *ChargebackReversalPaymentBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChargebackReversalPaymentServiceUnavailable creates a ChargebackReversalPaymentServiceUnavailable with default headers values
func NewChargebackReversalPaymentServiceUnavailable() *ChargebackReversalPaymentServiceUnavailable {
	return &ChargebackReversalPaymentServiceUnavailable{}
}

/* ChargebackReversalPaymentServiceUnavailable describes a response with status code 503, with default header values.

Payment in unknown status, failed to receive gateway response
*/
type ChargebackReversalPaymentServiceUnavailable struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the chargeback reversal payment service unavailable response
func (o *ChargebackReversalPaymentServiceUnavailable) Code() int {
	return 503
}

// IsSuccess returns true when this chargeback reversal payment service unavailable response has a 2xx status code
func (o *ChargebackReversalPaymentServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this chargeback reversal payment service unavailable response has a 3xx status code
func (o *ChargebackReversalPaymentServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this chargeback reversal payment service unavailable response has a 4xx status code
func (o *ChargebackReversalPaymentServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this chargeback reversal payment service unavailable response has a 5xx status code
func (o *ChargebackReversalPaymentServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this chargeback reversal payment service unavailable response a status code equal to that given
func (o *ChargebackReversalPaymentServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *ChargebackReversalPaymentServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}/chargebackReversals][%d] chargebackReversalPaymentServiceUnavailable ", 503)
}

func (o *ChargebackReversalPaymentServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}/chargebackReversals][%d] chargebackReversalPaymentServiceUnavailable ", 503)
}

func (o *ChargebackReversalPaymentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChargebackReversalPaymentGatewayTimeout creates a ChargebackReversalPaymentGatewayTimeout with default headers values
func NewChargebackReversalPaymentGatewayTimeout() *ChargebackReversalPaymentGatewayTimeout {
	return &ChargebackReversalPaymentGatewayTimeout{}
}

/* ChargebackReversalPaymentGatewayTimeout describes a response with status code 504, with default header values.

Payment operation timeout
*/
type ChargebackReversalPaymentGatewayTimeout struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the chargeback reversal payment gateway timeout response
func (o *ChargebackReversalPaymentGatewayTimeout) Code() int {
	return 504
}

// IsSuccess returns true when this chargeback reversal payment gateway timeout response has a 2xx status code
func (o *ChargebackReversalPaymentGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this chargeback reversal payment gateway timeout response has a 3xx status code
func (o *ChargebackReversalPaymentGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this chargeback reversal payment gateway timeout response has a 4xx status code
func (o *ChargebackReversalPaymentGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this chargeback reversal payment gateway timeout response has a 5xx status code
func (o *ChargebackReversalPaymentGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this chargeback reversal payment gateway timeout response a status code equal to that given
func (o *ChargebackReversalPaymentGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *ChargebackReversalPaymentGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}/chargebackReversals][%d] chargebackReversalPaymentGatewayTimeout ", 504)
}

func (o *ChargebackReversalPaymentGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}/chargebackReversals][%d] chargebackReversalPaymentGatewayTimeout ", 504)
}

func (o *ChargebackReversalPaymentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
