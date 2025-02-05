// Code generated by go-swagger; DO NOT EDIT.

package payment_method

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

// GetPaymentMethodReader is a Reader for the GetPaymentMethod structure.
type GetPaymentMethodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentMethodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentMethodOK()
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

// NewGetPaymentMethodOK creates a GetPaymentMethodOK with default headers values
func NewGetPaymentMethodOK() *GetPaymentMethodOK {
	return &GetPaymentMethodOK{}
}

/* GetPaymentMethodOK describes a response with status code 200, with default header values.

successful operation
*/
type GetPaymentMethodOK struct {
	Payload      *kbmodel.PaymentMethod
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get payment method o k response
func (o *GetPaymentMethodOK) Code() int {
	return 200
}

// IsSuccess returns true when this get payment method o k response has a 2xx status code
func (o *GetPaymentMethodOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get payment method o k response has a 3xx status code
func (o *GetPaymentMethodOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get payment method o k response has a 4xx status code
func (o *GetPaymentMethodOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get payment method o k response has a 5xx status code
func (o *GetPaymentMethodOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get payment method o k response a status code equal to that given
func (o *GetPaymentMethodOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetPaymentMethodOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/paymentMethods/{paymentMethodId}][%d] getPaymentMethodOK  %+v", 200, o.Payload)
}

func (o *GetPaymentMethodOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/paymentMethods/{paymentMethodId}][%d] getPaymentMethodOK  %+v", 200, o.Payload)
}

func (o *GetPaymentMethodOK) GetPayload() *kbmodel.PaymentMethod {
	return o.Payload
}

func (o *GetPaymentMethodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.PaymentMethod)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentMethodBadRequest creates a GetPaymentMethodBadRequest with default headers values
func NewGetPaymentMethodBadRequest() *GetPaymentMethodBadRequest {
	return &GetPaymentMethodBadRequest{}
}

/* GetPaymentMethodBadRequest describes a response with status code 400, with default header values.

Invalid paymentMethodId supplied
*/
type GetPaymentMethodBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get payment method bad request response
func (o *GetPaymentMethodBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this get payment method bad request response has a 2xx status code
func (o *GetPaymentMethodBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get payment method bad request response has a 3xx status code
func (o *GetPaymentMethodBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get payment method bad request response has a 4xx status code
func (o *GetPaymentMethodBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get payment method bad request response has a 5xx status code
func (o *GetPaymentMethodBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get payment method bad request response a status code equal to that given
func (o *GetPaymentMethodBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetPaymentMethodBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/paymentMethods/{paymentMethodId}][%d] getPaymentMethodBadRequest ", 400)
}

func (o *GetPaymentMethodBadRequest) String() string {
	return fmt.Sprintf("[GET /1.0/kb/paymentMethods/{paymentMethodId}][%d] getPaymentMethodBadRequest ", 400)
}

func (o *GetPaymentMethodBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPaymentMethodNotFound creates a GetPaymentMethodNotFound with default headers values
func NewGetPaymentMethodNotFound() *GetPaymentMethodNotFound {
	return &GetPaymentMethodNotFound{}
}

/* GetPaymentMethodNotFound describes a response with status code 404, with default header values.

Account or payment method not found
*/
type GetPaymentMethodNotFound struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get payment method not found response
func (o *GetPaymentMethodNotFound) Code() int {
	return 404
}

// IsSuccess returns true when this get payment method not found response has a 2xx status code
func (o *GetPaymentMethodNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get payment method not found response has a 3xx status code
func (o *GetPaymentMethodNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get payment method not found response has a 4xx status code
func (o *GetPaymentMethodNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get payment method not found response has a 5xx status code
func (o *GetPaymentMethodNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get payment method not found response a status code equal to that given
func (o *GetPaymentMethodNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetPaymentMethodNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/paymentMethods/{paymentMethodId}][%d] getPaymentMethodNotFound ", 404)
}

func (o *GetPaymentMethodNotFound) String() string {
	return fmt.Sprintf("[GET /1.0/kb/paymentMethods/{paymentMethodId}][%d] getPaymentMethodNotFound ", 404)
}

func (o *GetPaymentMethodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
