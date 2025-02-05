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
)

// DeletePushNotificationCallbacksReader is a Reader for the DeletePushNotificationCallbacks structure.
type DeletePushNotificationCallbacksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePushNotificationCallbacksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeletePushNotificationCallbacksNoContent()
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

// NewDeletePushNotificationCallbacksNoContent creates a DeletePushNotificationCallbacksNoContent with default headers values
func NewDeletePushNotificationCallbacksNoContent() *DeletePushNotificationCallbacksNoContent {
	return &DeletePushNotificationCallbacksNoContent{}
}

/* DeletePushNotificationCallbacksNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type DeletePushNotificationCallbacksNoContent struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the delete push notification callbacks no content response
func (o *DeletePushNotificationCallbacksNoContent) Code() int {
	return 204
}

// IsSuccess returns true when this delete push notification callbacks no content response has a 2xx status code
func (o *DeletePushNotificationCallbacksNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete push notification callbacks no content response has a 3xx status code
func (o *DeletePushNotificationCallbacksNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete push notification callbacks no content response has a 4xx status code
func (o *DeletePushNotificationCallbacksNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete push notification callbacks no content response has a 5xx status code
func (o *DeletePushNotificationCallbacksNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete push notification callbacks no content response a status code equal to that given
func (o *DeletePushNotificationCallbacksNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeletePushNotificationCallbacksNoContent) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/tenants/registerNotificationCallback][%d] deletePushNotificationCallbacksNoContent ", 204)
}

func (o *DeletePushNotificationCallbacksNoContent) String() string {
	return fmt.Sprintf("[DELETE /1.0/kb/tenants/registerNotificationCallback][%d] deletePushNotificationCallbacksNoContent ", 204)
}

func (o *DeletePushNotificationCallbacksNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePushNotificationCallbacksBadRequest creates a DeletePushNotificationCallbacksBadRequest with default headers values
func NewDeletePushNotificationCallbacksBadRequest() *DeletePushNotificationCallbacksBadRequest {
	return &DeletePushNotificationCallbacksBadRequest{}
}

/* DeletePushNotificationCallbacksBadRequest describes a response with status code 400, with default header values.

Invalid tenantId supplied
*/
type DeletePushNotificationCallbacksBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the delete push notification callbacks bad request response
func (o *DeletePushNotificationCallbacksBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this delete push notification callbacks bad request response has a 2xx status code
func (o *DeletePushNotificationCallbacksBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete push notification callbacks bad request response has a 3xx status code
func (o *DeletePushNotificationCallbacksBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete push notification callbacks bad request response has a 4xx status code
func (o *DeletePushNotificationCallbacksBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete push notification callbacks bad request response has a 5xx status code
func (o *DeletePushNotificationCallbacksBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete push notification callbacks bad request response a status code equal to that given
func (o *DeletePushNotificationCallbacksBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeletePushNotificationCallbacksBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/tenants/registerNotificationCallback][%d] deletePushNotificationCallbacksBadRequest ", 400)
}

func (o *DeletePushNotificationCallbacksBadRequest) String() string {
	return fmt.Sprintf("[DELETE /1.0/kb/tenants/registerNotificationCallback][%d] deletePushNotificationCallbacksBadRequest ", 400)
}

func (o *DeletePushNotificationCallbacksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
