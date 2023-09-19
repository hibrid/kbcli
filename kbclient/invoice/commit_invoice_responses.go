// Code generated by go-swagger; DO NOT EDIT.

package invoice

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v3/kbcommon"
)

// CommitInvoiceReader is a Reader for the CommitInvoice structure.
type CommitInvoiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CommitInvoiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewCommitInvoiceNoContent()
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

// NewCommitInvoiceNoContent creates a CommitInvoiceNoContent with default headers values
func NewCommitInvoiceNoContent() *CommitInvoiceNoContent {
	return &CommitInvoiceNoContent{}
}

/* CommitInvoiceNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type CommitInvoiceNoContent struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the commit invoice no content response
func (o *CommitInvoiceNoContent) Code() int {
	return 204
}

// IsSuccess returns true when this commit invoice no content response has a 2xx status code
func (o *CommitInvoiceNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this commit invoice no content response has a 3xx status code
func (o *CommitInvoiceNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this commit invoice no content response has a 4xx status code
func (o *CommitInvoiceNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this commit invoice no content response has a 5xx status code
func (o *CommitInvoiceNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this commit invoice no content response a status code equal to that given
func (o *CommitInvoiceNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *CommitInvoiceNoContent) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/invoices/{invoiceId}/commitInvoice][%d] commitInvoiceNoContent ", 204)
}

func (o *CommitInvoiceNoContent) String() string {
	return fmt.Sprintf("[PUT /1.0/kb/invoices/{invoiceId}/commitInvoice][%d] commitInvoiceNoContent ", 204)
}

func (o *CommitInvoiceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCommitInvoiceNotFound creates a CommitInvoiceNotFound with default headers values
func NewCommitInvoiceNotFound() *CommitInvoiceNotFound {
	return &CommitInvoiceNotFound{}
}

/* CommitInvoiceNotFound describes a response with status code 404, with default header values.

Invoice not found
*/
type CommitInvoiceNotFound struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the commit invoice not found response
func (o *CommitInvoiceNotFound) Code() int {
	return 404
}

// IsSuccess returns true when this commit invoice not found response has a 2xx status code
func (o *CommitInvoiceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this commit invoice not found response has a 3xx status code
func (o *CommitInvoiceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this commit invoice not found response has a 4xx status code
func (o *CommitInvoiceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this commit invoice not found response has a 5xx status code
func (o *CommitInvoiceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this commit invoice not found response a status code equal to that given
func (o *CommitInvoiceNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CommitInvoiceNotFound) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/invoices/{invoiceId}/commitInvoice][%d] commitInvoiceNotFound ", 404)
}

func (o *CommitInvoiceNotFound) String() string {
	return fmt.Sprintf("[PUT /1.0/kb/invoices/{invoiceId}/commitInvoice][%d] commitInvoiceNotFound ", 404)
}

func (o *CommitInvoiceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
