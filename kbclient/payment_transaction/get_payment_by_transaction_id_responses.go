// Code generated by go-swagger; DO NOT EDIT.

package payment_transaction

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

// GetPaymentByTransactionIDReader is a Reader for the GetPaymentByTransactionID structure.
type GetPaymentByTransactionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentByTransactionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentByTransactionIDOK()
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

// NewGetPaymentByTransactionIDOK creates a GetPaymentByTransactionIDOK with default headers values
func NewGetPaymentByTransactionIDOK() *GetPaymentByTransactionIDOK {
	return &GetPaymentByTransactionIDOK{}
}

/* GetPaymentByTransactionIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetPaymentByTransactionIDOK struct {
	Payload      *kbmodel.Payment
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get payment by transaction Id o k response
func (o *GetPaymentByTransactionIDOK) Code() int {
	return 200
}

// IsSuccess returns true when this get payment by transaction Id o k response has a 2xx status code
func (o *GetPaymentByTransactionIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get payment by transaction Id o k response has a 3xx status code
func (o *GetPaymentByTransactionIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get payment by transaction Id o k response has a 4xx status code
func (o *GetPaymentByTransactionIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get payment by transaction Id o k response has a 5xx status code
func (o *GetPaymentByTransactionIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get payment by transaction Id o k response a status code equal to that given
func (o *GetPaymentByTransactionIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetPaymentByTransactionIDOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/paymentTransactions/{transactionId}][%d] getPaymentByTransactionIdOK  %+v", 200, o.Payload)
}

func (o *GetPaymentByTransactionIDOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/paymentTransactions/{transactionId}][%d] getPaymentByTransactionIdOK  %+v", 200, o.Payload)
}

func (o *GetPaymentByTransactionIDOK) GetPayload() *kbmodel.Payment {
	return o.Payload
}

func (o *GetPaymentByTransactionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Payment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentByTransactionIDNotFound creates a GetPaymentByTransactionIDNotFound with default headers values
func NewGetPaymentByTransactionIDNotFound() *GetPaymentByTransactionIDNotFound {
	return &GetPaymentByTransactionIDNotFound{}
}

/* GetPaymentByTransactionIDNotFound describes a response with status code 404, with default header values.

Payment not found
*/
type GetPaymentByTransactionIDNotFound struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get payment by transaction Id not found response
func (o *GetPaymentByTransactionIDNotFound) Code() int {
	return 404
}

// IsSuccess returns true when this get payment by transaction Id not found response has a 2xx status code
func (o *GetPaymentByTransactionIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get payment by transaction Id not found response has a 3xx status code
func (o *GetPaymentByTransactionIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get payment by transaction Id not found response has a 4xx status code
func (o *GetPaymentByTransactionIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get payment by transaction Id not found response has a 5xx status code
func (o *GetPaymentByTransactionIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get payment by transaction Id not found response a status code equal to that given
func (o *GetPaymentByTransactionIDNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetPaymentByTransactionIDNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/paymentTransactions/{transactionId}][%d] getPaymentByTransactionIdNotFound ", 404)
}

func (o *GetPaymentByTransactionIDNotFound) String() string {
	return fmt.Sprintf("[GET /1.0/kb/paymentTransactions/{transactionId}][%d] getPaymentByTransactionIdNotFound ", 404)
}

func (o *GetPaymentByTransactionIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
