// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v3/kbcommon"
)

// GetQueueEntriesReader is a Reader for the GetQueueEntries structure.
type GetQueueEntriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetQueueEntriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetQueueEntriesOK()
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

// NewGetQueueEntriesOK creates a GetQueueEntriesOK with default headers values
func NewGetQueueEntriesOK() *GetQueueEntriesOK {
	return &GetQueueEntriesOK{}
}

/* GetQueueEntriesOK describes a response with status code 200, with default header values.

Success
*/
type GetQueueEntriesOK struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get queue entries o k response
func (o *GetQueueEntriesOK) Code() int {
	return 200
}

// IsSuccess returns true when this get queue entries o k response has a 2xx status code
func (o *GetQueueEntriesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get queue entries o k response has a 3xx status code
func (o *GetQueueEntriesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get queue entries o k response has a 4xx status code
func (o *GetQueueEntriesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get queue entries o k response has a 5xx status code
func (o *GetQueueEntriesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get queue entries o k response a status code equal to that given
func (o *GetQueueEntriesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetQueueEntriesOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/admin/queues][%d] getQueueEntriesOK ", 200)
}

func (o *GetQueueEntriesOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/admin/queues][%d] getQueueEntriesOK ", 200)
}

func (o *GetQueueEntriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetQueueEntriesBadRequest creates a GetQueueEntriesBadRequest with default headers values
func NewGetQueueEntriesBadRequest() *GetQueueEntriesBadRequest {
	return &GetQueueEntriesBadRequest{}
}

/* GetQueueEntriesBadRequest describes a response with status code 400, with default header values.

Invalid account id supplied
*/
type GetQueueEntriesBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get queue entries bad request response
func (o *GetQueueEntriesBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this get queue entries bad request response has a 2xx status code
func (o *GetQueueEntriesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get queue entries bad request response has a 3xx status code
func (o *GetQueueEntriesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get queue entries bad request response has a 4xx status code
func (o *GetQueueEntriesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get queue entries bad request response has a 5xx status code
func (o *GetQueueEntriesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get queue entries bad request response a status code equal to that given
func (o *GetQueueEntriesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetQueueEntriesBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/admin/queues][%d] getQueueEntriesBadRequest ", 400)
}

func (o *GetQueueEntriesBadRequest) String() string {
	return fmt.Sprintf("[GET /1.0/kb/admin/queues][%d] getQueueEntriesBadRequest ", 400)
}

func (o *GetQueueEntriesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetQueueEntriesNotFound creates a GetQueueEntriesNotFound with default headers values
func NewGetQueueEntriesNotFound() *GetQueueEntriesNotFound {
	return &GetQueueEntriesNotFound{}
}

/* GetQueueEntriesNotFound describes a response with status code 404, with default header values.

Account not found
*/
type GetQueueEntriesNotFound struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get queue entries not found response
func (o *GetQueueEntriesNotFound) Code() int {
	return 404
}

// IsSuccess returns true when this get queue entries not found response has a 2xx status code
func (o *GetQueueEntriesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get queue entries not found response has a 3xx status code
func (o *GetQueueEntriesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get queue entries not found response has a 4xx status code
func (o *GetQueueEntriesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get queue entries not found response has a 5xx status code
func (o *GetQueueEntriesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get queue entries not found response a status code equal to that given
func (o *GetQueueEntriesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetQueueEntriesNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/admin/queues][%d] getQueueEntriesNotFound ", 404)
}

func (o *GetQueueEntriesNotFound) String() string {
	return fmt.Sprintf("[GET /1.0/kb/admin/queues][%d] getQueueEntriesNotFound ", 404)
}

func (o *GetQueueEntriesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
