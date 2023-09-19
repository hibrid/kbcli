// Code generated by go-swagger; DO NOT EDIT.

package tenant

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

// InsertUserKeyValueReader is a Reader for the InsertUserKeyValue structure.
type InsertUserKeyValueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InsertUserKeyValueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewInsertUserKeyValueCreated()
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

// NewInsertUserKeyValueCreated creates a InsertUserKeyValueCreated with default headers values
func NewInsertUserKeyValueCreated() *InsertUserKeyValueCreated {
	return &InsertUserKeyValueCreated{}
}

/* InsertUserKeyValueCreated describes a response with status code 201, with default header values.

Per tenant config uploaded successfully
*/
type InsertUserKeyValueCreated struct {
	Payload      *kbmodel.TenantKeyValue
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the insert user key value created response
func (o *InsertUserKeyValueCreated) Code() int {
	return 201
}

// IsSuccess returns true when this insert user key value created response has a 2xx status code
func (o *InsertUserKeyValueCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this insert user key value created response has a 3xx status code
func (o *InsertUserKeyValueCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this insert user key value created response has a 4xx status code
func (o *InsertUserKeyValueCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this insert user key value created response has a 5xx status code
func (o *InsertUserKeyValueCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this insert user key value created response a status code equal to that given
func (o *InsertUserKeyValueCreated) IsCode(code int) bool {
	return code == 201
}

func (o *InsertUserKeyValueCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/tenants/userKeyValue/{keyName}][%d] insertUserKeyValueCreated  %+v", 201, o.Payload)
}

func (o *InsertUserKeyValueCreated) String() string {
	return fmt.Sprintf("[POST /1.0/kb/tenants/userKeyValue/{keyName}][%d] insertUserKeyValueCreated  %+v", 201, o.Payload)
}

func (o *InsertUserKeyValueCreated) GetPayload() *kbmodel.TenantKeyValue {
	return o.Payload
}

func (o *InsertUserKeyValueCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.TenantKeyValue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInsertUserKeyValueBadRequest creates a InsertUserKeyValueBadRequest with default headers values
func NewInsertUserKeyValueBadRequest() *InsertUserKeyValueBadRequest {
	return &InsertUserKeyValueBadRequest{}
}

/* InsertUserKeyValueBadRequest describes a response with status code 400, with default header values.

Invalid tenantId supplied
*/
type InsertUserKeyValueBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the insert user key value bad request response
func (o *InsertUserKeyValueBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this insert user key value bad request response has a 2xx status code
func (o *InsertUserKeyValueBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this insert user key value bad request response has a 3xx status code
func (o *InsertUserKeyValueBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this insert user key value bad request response has a 4xx status code
func (o *InsertUserKeyValueBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this insert user key value bad request response has a 5xx status code
func (o *InsertUserKeyValueBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this insert user key value bad request response a status code equal to that given
func (o *InsertUserKeyValueBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *InsertUserKeyValueBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/tenants/userKeyValue/{keyName}][%d] insertUserKeyValueBadRequest ", 400)
}

func (o *InsertUserKeyValueBadRequest) String() string {
	return fmt.Sprintf("[POST /1.0/kb/tenants/userKeyValue/{keyName}][%d] insertUserKeyValueBadRequest ", 400)
}

func (o *InsertUserKeyValueBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
