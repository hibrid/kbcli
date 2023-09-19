// Code generated by go-swagger; DO NOT EDIT.

package credit

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

// GetCreditReader is a Reader for the GetCredit structure.
type GetCreditReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCreditReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCreditOK()
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

// NewGetCreditOK creates a GetCreditOK with default headers values
func NewGetCreditOK() *GetCreditOK {
	return &GetCreditOK{}
}

/* GetCreditOK describes a response with status code 200, with default header values.

successful operation
*/
type GetCreditOK struct {
	Payload      *kbmodel.InvoiceItem
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get credit o k response
func (o *GetCreditOK) Code() int {
	return 200
}

// IsSuccess returns true when this get credit o k response has a 2xx status code
func (o *GetCreditOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get credit o k response has a 3xx status code
func (o *GetCreditOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get credit o k response has a 4xx status code
func (o *GetCreditOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get credit o k response has a 5xx status code
func (o *GetCreditOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get credit o k response a status code equal to that given
func (o *GetCreditOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetCreditOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/credits/{creditId}][%d] getCreditOK  %+v", 200, o.Payload)
}

func (o *GetCreditOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/credits/{creditId}][%d] getCreditOK  %+v", 200, o.Payload)
}

func (o *GetCreditOK) GetPayload() *kbmodel.InvoiceItem {
	return o.Payload
}

func (o *GetCreditOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.InvoiceItem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCreditBadRequest creates a GetCreditBadRequest with default headers values
func NewGetCreditBadRequest() *GetCreditBadRequest {
	return &GetCreditBadRequest{}
}

/* GetCreditBadRequest describes a response with status code 400, with default header values.

Invalid credit id supplied
*/
type GetCreditBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get credit bad request response
func (o *GetCreditBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this get credit bad request response has a 2xx status code
func (o *GetCreditBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get credit bad request response has a 3xx status code
func (o *GetCreditBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get credit bad request response has a 4xx status code
func (o *GetCreditBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get credit bad request response has a 5xx status code
func (o *GetCreditBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get credit bad request response a status code equal to that given
func (o *GetCreditBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetCreditBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/credits/{creditId}][%d] getCreditBadRequest ", 400)
}

func (o *GetCreditBadRequest) String() string {
	return fmt.Sprintf("[GET /1.0/kb/credits/{creditId}][%d] getCreditBadRequest ", 400)
}

func (o *GetCreditBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCreditNotFound creates a GetCreditNotFound with default headers values
func NewGetCreditNotFound() *GetCreditNotFound {
	return &GetCreditNotFound{}
}

/* GetCreditNotFound describes a response with status code 404, with default header values.

Credit not found
*/
type GetCreditNotFound struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get credit not found response
func (o *GetCreditNotFound) Code() int {
	return 404
}

// IsSuccess returns true when this get credit not found response has a 2xx status code
func (o *GetCreditNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get credit not found response has a 3xx status code
func (o *GetCreditNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get credit not found response has a 4xx status code
func (o *GetCreditNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get credit not found response has a 5xx status code
func (o *GetCreditNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get credit not found response a status code equal to that given
func (o *GetCreditNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetCreditNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/credits/{creditId}][%d] getCreditNotFound ", 404)
}

func (o *GetCreditNotFound) String() string {
	return fmt.Sprintf("[GET /1.0/kb/credits/{creditId}][%d] getCreditNotFound ", 404)
}

func (o *GetCreditNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
