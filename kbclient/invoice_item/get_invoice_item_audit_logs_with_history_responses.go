// Code generated by go-swagger; DO NOT EDIT.

package invoice_item

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

// GetInvoiceItemAuditLogsWithHistoryReader is a Reader for the GetInvoiceItemAuditLogsWithHistory structure.
type GetInvoiceItemAuditLogsWithHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInvoiceItemAuditLogsWithHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInvoiceItemAuditLogsWithHistoryOK()
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

// NewGetInvoiceItemAuditLogsWithHistoryOK creates a GetInvoiceItemAuditLogsWithHistoryOK with default headers values
func NewGetInvoiceItemAuditLogsWithHistoryOK() *GetInvoiceItemAuditLogsWithHistoryOK {
	return &GetInvoiceItemAuditLogsWithHistoryOK{}
}

/* GetInvoiceItemAuditLogsWithHistoryOK describes a response with status code 200, with default header values.

successful operation
*/
type GetInvoiceItemAuditLogsWithHistoryOK struct {
	Payload      []*kbmodel.AuditLog
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get invoice item audit logs with history o k response
func (o *GetInvoiceItemAuditLogsWithHistoryOK) Code() int {
	return 200
}

// IsSuccess returns true when this get invoice item audit logs with history o k response has a 2xx status code
func (o *GetInvoiceItemAuditLogsWithHistoryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get invoice item audit logs with history o k response has a 3xx status code
func (o *GetInvoiceItemAuditLogsWithHistoryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoice item audit logs with history o k response has a 4xx status code
func (o *GetInvoiceItemAuditLogsWithHistoryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get invoice item audit logs with history o k response has a 5xx status code
func (o *GetInvoiceItemAuditLogsWithHistoryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoice item audit logs with history o k response a status code equal to that given
func (o *GetInvoiceItemAuditLogsWithHistoryOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetInvoiceItemAuditLogsWithHistoryOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoiceItems/{invoiceItemId}/auditLogsWithHistory][%d] getInvoiceItemAuditLogsWithHistoryOK  %+v", 200, o.Payload)
}

func (o *GetInvoiceItemAuditLogsWithHistoryOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/invoiceItems/{invoiceItemId}/auditLogsWithHistory][%d] getInvoiceItemAuditLogsWithHistoryOK  %+v", 200, o.Payload)
}

func (o *GetInvoiceItemAuditLogsWithHistoryOK) GetPayload() []*kbmodel.AuditLog {
	return o.Payload
}

func (o *GetInvoiceItemAuditLogsWithHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoiceItemAuditLogsWithHistoryNotFound creates a GetInvoiceItemAuditLogsWithHistoryNotFound with default headers values
func NewGetInvoiceItemAuditLogsWithHistoryNotFound() *GetInvoiceItemAuditLogsWithHistoryNotFound {
	return &GetInvoiceItemAuditLogsWithHistoryNotFound{}
}

/* GetInvoiceItemAuditLogsWithHistoryNotFound describes a response with status code 404, with default header values.

Invoice item not found
*/
type GetInvoiceItemAuditLogsWithHistoryNotFound struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get invoice item audit logs with history not found response
func (o *GetInvoiceItemAuditLogsWithHistoryNotFound) Code() int {
	return 404
}

// IsSuccess returns true when this get invoice item audit logs with history not found response has a 2xx status code
func (o *GetInvoiceItemAuditLogsWithHistoryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get invoice item audit logs with history not found response has a 3xx status code
func (o *GetInvoiceItemAuditLogsWithHistoryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoice item audit logs with history not found response has a 4xx status code
func (o *GetInvoiceItemAuditLogsWithHistoryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get invoice item audit logs with history not found response has a 5xx status code
func (o *GetInvoiceItemAuditLogsWithHistoryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoice item audit logs with history not found response a status code equal to that given
func (o *GetInvoiceItemAuditLogsWithHistoryNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetInvoiceItemAuditLogsWithHistoryNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoiceItems/{invoiceItemId}/auditLogsWithHistory][%d] getInvoiceItemAuditLogsWithHistoryNotFound ", 404)
}

func (o *GetInvoiceItemAuditLogsWithHistoryNotFound) String() string {
	return fmt.Sprintf("[GET /1.0/kb/invoiceItems/{invoiceItemId}/auditLogsWithHistory][%d] getInvoiceItemAuditLogsWithHistoryNotFound ", 404)
}

func (o *GetInvoiceItemAuditLogsWithHistoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
