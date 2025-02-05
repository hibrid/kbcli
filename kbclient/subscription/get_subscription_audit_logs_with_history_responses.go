// Code generated by go-swagger; DO NOT EDIT.

package subscription

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

// GetSubscriptionAuditLogsWithHistoryReader is a Reader for the GetSubscriptionAuditLogsWithHistory structure.
type GetSubscriptionAuditLogsWithHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSubscriptionAuditLogsWithHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSubscriptionAuditLogsWithHistoryOK()
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

// NewGetSubscriptionAuditLogsWithHistoryOK creates a GetSubscriptionAuditLogsWithHistoryOK with default headers values
func NewGetSubscriptionAuditLogsWithHistoryOK() *GetSubscriptionAuditLogsWithHistoryOK {
	return &GetSubscriptionAuditLogsWithHistoryOK{}
}

/* GetSubscriptionAuditLogsWithHistoryOK describes a response with status code 200, with default header values.

successful operation
*/
type GetSubscriptionAuditLogsWithHistoryOK struct {
	Payload      []*kbmodel.AuditLog
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get subscription audit logs with history o k response
func (o *GetSubscriptionAuditLogsWithHistoryOK) Code() int {
	return 200
}

// IsSuccess returns true when this get subscription audit logs with history o k response has a 2xx status code
func (o *GetSubscriptionAuditLogsWithHistoryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get subscription audit logs with history o k response has a 3xx status code
func (o *GetSubscriptionAuditLogsWithHistoryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get subscription audit logs with history o k response has a 4xx status code
func (o *GetSubscriptionAuditLogsWithHistoryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get subscription audit logs with history o k response has a 5xx status code
func (o *GetSubscriptionAuditLogsWithHistoryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get subscription audit logs with history o k response a status code equal to that given
func (o *GetSubscriptionAuditLogsWithHistoryOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetSubscriptionAuditLogsWithHistoryOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/subscriptions/{subscriptionId}/auditLogsWithHistory][%d] getSubscriptionAuditLogsWithHistoryOK  %+v", 200, o.Payload)
}

func (o *GetSubscriptionAuditLogsWithHistoryOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/subscriptions/{subscriptionId}/auditLogsWithHistory][%d] getSubscriptionAuditLogsWithHistoryOK  %+v", 200, o.Payload)
}

func (o *GetSubscriptionAuditLogsWithHistoryOK) GetPayload() []*kbmodel.AuditLog {
	return o.Payload
}

func (o *GetSubscriptionAuditLogsWithHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubscriptionAuditLogsWithHistoryNotFound creates a GetSubscriptionAuditLogsWithHistoryNotFound with default headers values
func NewGetSubscriptionAuditLogsWithHistoryNotFound() *GetSubscriptionAuditLogsWithHistoryNotFound {
	return &GetSubscriptionAuditLogsWithHistoryNotFound{}
}

/* GetSubscriptionAuditLogsWithHistoryNotFound describes a response with status code 404, with default header values.

Subscription not found
*/
type GetSubscriptionAuditLogsWithHistoryNotFound struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get subscription audit logs with history not found response
func (o *GetSubscriptionAuditLogsWithHistoryNotFound) Code() int {
	return 404
}

// IsSuccess returns true when this get subscription audit logs with history not found response has a 2xx status code
func (o *GetSubscriptionAuditLogsWithHistoryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get subscription audit logs with history not found response has a 3xx status code
func (o *GetSubscriptionAuditLogsWithHistoryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get subscription audit logs with history not found response has a 4xx status code
func (o *GetSubscriptionAuditLogsWithHistoryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get subscription audit logs with history not found response has a 5xx status code
func (o *GetSubscriptionAuditLogsWithHistoryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get subscription audit logs with history not found response a status code equal to that given
func (o *GetSubscriptionAuditLogsWithHistoryNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetSubscriptionAuditLogsWithHistoryNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/subscriptions/{subscriptionId}/auditLogsWithHistory][%d] getSubscriptionAuditLogsWithHistoryNotFound ", 404)
}

func (o *GetSubscriptionAuditLogsWithHistoryNotFound) String() string {
	return fmt.Sprintf("[GET /1.0/kb/subscriptions/{subscriptionId}/auditLogsWithHistory][%d] getSubscriptionAuditLogsWithHistoryNotFound ", 404)
}

func (o *GetSubscriptionAuditLogsWithHistoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
