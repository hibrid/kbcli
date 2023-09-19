// Code generated by go-swagger; DO NOT EDIT.

package catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v3/kbcommon"
)

// AddSimplePlanReader is a Reader for the AddSimplePlan structure.
type AddSimplePlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddSimplePlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewAddSimplePlanCreated()
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

// NewAddSimplePlanCreated creates a AddSimplePlanCreated with default headers values
func NewAddSimplePlanCreated() *AddSimplePlanCreated {
	return &AddSimplePlanCreated{}
}

/* AddSimplePlanCreated describes a response with status code 201, with default header values.

Created new plan successfully
*/
type AddSimplePlanCreated struct {
	Payload      string
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the add simple plan created response
func (o *AddSimplePlanCreated) Code() int {
	return 201
}

// IsSuccess returns true when this add simple plan created response has a 2xx status code
func (o *AddSimplePlanCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add simple plan created response has a 3xx status code
func (o *AddSimplePlanCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add simple plan created response has a 4xx status code
func (o *AddSimplePlanCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this add simple plan created response has a 5xx status code
func (o *AddSimplePlanCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this add simple plan created response a status code equal to that given
func (o *AddSimplePlanCreated) IsCode(code int) bool {
	return code == 201
}

func (o *AddSimplePlanCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/catalog/simplePlan][%d] addSimplePlanCreated  %+v", 201, o.Payload)
}

func (o *AddSimplePlanCreated) String() string {
	return fmt.Sprintf("[POST /1.0/kb/catalog/simplePlan][%d] addSimplePlanCreated  %+v", 201, o.Payload)
}

func (o *AddSimplePlanCreated) GetPayload() string {
	return o.Payload
}

func (o *AddSimplePlanCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
