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

// GetTransactionTagsReader is a Reader for the GetTransactionTags structure.
type GetTransactionTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTransactionTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTransactionTagsOK()
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

// NewGetTransactionTagsOK creates a GetTransactionTagsOK with default headers values
func NewGetTransactionTagsOK() *GetTransactionTagsOK {
	return &GetTransactionTagsOK{}
}

/* GetTransactionTagsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetTransactionTagsOK struct {
	Payload      []*kbmodel.Tag
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get transaction tags o k response
func (o *GetTransactionTagsOK) Code() int {
	return 200
}

// IsSuccess returns true when this get transaction tags o k response has a 2xx status code
func (o *GetTransactionTagsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get transaction tags o k response has a 3xx status code
func (o *GetTransactionTagsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get transaction tags o k response has a 4xx status code
func (o *GetTransactionTagsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get transaction tags o k response has a 5xx status code
func (o *GetTransactionTagsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get transaction tags o k response a status code equal to that given
func (o *GetTransactionTagsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetTransactionTagsOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/paymentTransactions/{transactionId}/tags][%d] getTransactionTagsOK  %+v", 200, o.Payload)
}

func (o *GetTransactionTagsOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/paymentTransactions/{transactionId}/tags][%d] getTransactionTagsOK  %+v", 200, o.Payload)
}

func (o *GetTransactionTagsOK) GetPayload() []*kbmodel.Tag {
	return o.Payload
}

func (o *GetTransactionTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionTagsBadRequest creates a GetTransactionTagsBadRequest with default headers values
func NewGetTransactionTagsBadRequest() *GetTransactionTagsBadRequest {
	return &GetTransactionTagsBadRequest{}
}

/* GetTransactionTagsBadRequest describes a response with status code 400, with default header values.

Invalid transaction id supplied
*/
type GetTransactionTagsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get transaction tags bad request response
func (o *GetTransactionTagsBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this get transaction tags bad request response has a 2xx status code
func (o *GetTransactionTagsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get transaction tags bad request response has a 3xx status code
func (o *GetTransactionTagsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get transaction tags bad request response has a 4xx status code
func (o *GetTransactionTagsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get transaction tags bad request response has a 5xx status code
func (o *GetTransactionTagsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get transaction tags bad request response a status code equal to that given
func (o *GetTransactionTagsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetTransactionTagsBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/paymentTransactions/{transactionId}/tags][%d] getTransactionTagsBadRequest ", 400)
}

func (o *GetTransactionTagsBadRequest) String() string {
	return fmt.Sprintf("[GET /1.0/kb/paymentTransactions/{transactionId}/tags][%d] getTransactionTagsBadRequest ", 400)
}

func (o *GetTransactionTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTransactionTagsNotFound creates a GetTransactionTagsNotFound with default headers values
func NewGetTransactionTagsNotFound() *GetTransactionTagsNotFound {
	return &GetTransactionTagsNotFound{}
}

/* GetTransactionTagsNotFound describes a response with status code 404, with default header values.

Invoice not found
*/
type GetTransactionTagsNotFound struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get transaction tags not found response
func (o *GetTransactionTagsNotFound) Code() int {
	return 404
}

// IsSuccess returns true when this get transaction tags not found response has a 2xx status code
func (o *GetTransactionTagsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get transaction tags not found response has a 3xx status code
func (o *GetTransactionTagsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get transaction tags not found response has a 4xx status code
func (o *GetTransactionTagsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get transaction tags not found response has a 5xx status code
func (o *GetTransactionTagsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get transaction tags not found response a status code equal to that given
func (o *GetTransactionTagsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetTransactionTagsNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/paymentTransactions/{transactionId}/tags][%d] getTransactionTagsNotFound ", 404)
}

func (o *GetTransactionTagsNotFound) String() string {
	return fmt.Sprintf("[GET /1.0/kb/paymentTransactions/{transactionId}/tags][%d] getTransactionTagsNotFound ", 404)
}

func (o *GetTransactionTagsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
