// Code generated by go-swagger; DO NOT EDIT.

package payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/kbcommon"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/kbmodel"
)

// ChargebackPaymentReader is a Reader for the ChargebackPayment structure.
type ChargebackPaymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChargebackPaymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewChargebackPaymentCreated()
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

// NewChargebackPaymentCreated creates a ChargebackPaymentCreated with default headers values
func NewChargebackPaymentCreated() *ChargebackPaymentCreated {
	return &ChargebackPaymentCreated{}
}

/*ChargebackPaymentCreated handles this case with default header values.

Payment transaction created successfully
*/
type ChargebackPaymentCreated struct {
	Payload *kbmodel.Payment

	HttpResponse runtime.ClientResponse
}

func (o *ChargebackPaymentCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}/chargebacks][%d] chargebackPaymentCreated  %+v", 201, o.Payload)
}

func (o *ChargebackPaymentCreated) GetPayload() *kbmodel.Payment {
	return o.Payload
}

func (o *ChargebackPaymentCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Payment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChargebackPaymentBadRequest creates a ChargebackPaymentBadRequest with default headers values
func NewChargebackPaymentBadRequest() *ChargebackPaymentBadRequest {
	return &ChargebackPaymentBadRequest{}
}

/*ChargebackPaymentBadRequest handles this case with default header values.

Invalid paymentId supplied
*/
type ChargebackPaymentBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *ChargebackPaymentBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}/chargebacks][%d] chargebackPaymentBadRequest ", 400)
}

func (o *ChargebackPaymentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChargebackPaymentPaymentRequired creates a ChargebackPaymentPaymentRequired with default headers values
func NewChargebackPaymentPaymentRequired() *ChargebackPaymentPaymentRequired {
	return &ChargebackPaymentPaymentRequired{}
}

/*ChargebackPaymentPaymentRequired handles this case with default header values.

Transaction declined by gateway
*/
type ChargebackPaymentPaymentRequired struct {
	HttpResponse runtime.ClientResponse
}

func (o *ChargebackPaymentPaymentRequired) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}/chargebacks][%d] chargebackPaymentPaymentRequired ", 402)
}

func (o *ChargebackPaymentPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChargebackPaymentNotFound creates a ChargebackPaymentNotFound with default headers values
func NewChargebackPaymentNotFound() *ChargebackPaymentNotFound {
	return &ChargebackPaymentNotFound{}
}

/*ChargebackPaymentNotFound handles this case with default header values.

Account or payment not found
*/
type ChargebackPaymentNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *ChargebackPaymentNotFound) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}/chargebacks][%d] chargebackPaymentNotFound ", 404)
}

func (o *ChargebackPaymentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChargebackPaymentUnprocessableEntity creates a ChargebackPaymentUnprocessableEntity with default headers values
func NewChargebackPaymentUnprocessableEntity() *ChargebackPaymentUnprocessableEntity {
	return &ChargebackPaymentUnprocessableEntity{}
}

/*ChargebackPaymentUnprocessableEntity handles this case with default header values.

Payment is aborted by a control plugin
*/
type ChargebackPaymentUnprocessableEntity struct {
	HttpResponse runtime.ClientResponse
}

func (o *ChargebackPaymentUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}/chargebacks][%d] chargebackPaymentUnprocessableEntity ", 422)
}

func (o *ChargebackPaymentUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChargebackPaymentBadGateway creates a ChargebackPaymentBadGateway with default headers values
func NewChargebackPaymentBadGateway() *ChargebackPaymentBadGateway {
	return &ChargebackPaymentBadGateway{}
}

/*ChargebackPaymentBadGateway handles this case with default header values.

Failed to submit payment transaction
*/
type ChargebackPaymentBadGateway struct {
	HttpResponse runtime.ClientResponse
}

func (o *ChargebackPaymentBadGateway) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}/chargebacks][%d] chargebackPaymentBadGateway ", 502)
}

func (o *ChargebackPaymentBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChargebackPaymentServiceUnavailable creates a ChargebackPaymentServiceUnavailable with default headers values
func NewChargebackPaymentServiceUnavailable() *ChargebackPaymentServiceUnavailable {
	return &ChargebackPaymentServiceUnavailable{}
}

/*ChargebackPaymentServiceUnavailable handles this case with default header values.

Payment in unknown status, failed to receive gateway response
*/
type ChargebackPaymentServiceUnavailable struct {
	HttpResponse runtime.ClientResponse
}

func (o *ChargebackPaymentServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}/chargebacks][%d] chargebackPaymentServiceUnavailable ", 503)
}

func (o *ChargebackPaymentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChargebackPaymentGatewayTimeout creates a ChargebackPaymentGatewayTimeout with default headers values
func NewChargebackPaymentGatewayTimeout() *ChargebackPaymentGatewayTimeout {
	return &ChargebackPaymentGatewayTimeout{}
}

/*ChargebackPaymentGatewayTimeout handles this case with default header values.

Payment operation timeout
*/
type ChargebackPaymentGatewayTimeout struct {
	HttpResponse runtime.ClientResponse
}

func (o *ChargebackPaymentGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}/chargebacks][%d] chargebackPaymentGatewayTimeout ", 504)
}

func (o *ChargebackPaymentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
