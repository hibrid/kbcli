// Code generated by go-swagger; DO NOT EDIT.

package t_n_custom_field

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

// CreateTNCustomFieldReader is a Reader for the CreateTNCustomField structure.
type CreateTNCustomFieldReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTNCustomFieldReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewCreateTNCustomFieldCreated()
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

// NewCreateTNCustomFieldCreated creates a CreateTNCustomFieldCreated with default headers values
func NewCreateTNCustomFieldCreated() *CreateTNCustomFieldCreated {
	return &CreateTNCustomFieldCreated{}
}

/* CreateTNCustomFieldCreated describes a response with status code 201, with default header values.

successful operation
*/
type CreateTNCustomFieldCreated struct {
	Payload      *kbmodel.TNCustomField
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the create t n custom field created response
func (o *CreateTNCustomFieldCreated) Code() int {
	return 201
}

// IsSuccess returns true when this create t n custom field created response has a 2xx status code
func (o *CreateTNCustomFieldCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create t n custom field created response has a 3xx status code
func (o *CreateTNCustomFieldCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create t n custom field created response has a 4xx status code
func (o *CreateTNCustomFieldCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create t n custom field created response has a 5xx status code
func (o *CreateTNCustomFieldCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create t n custom field created response a status code equal to that given
func (o *CreateTNCustomFieldCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateTNCustomFieldCreated) Error() string {
	return fmt.Sprintf("[POST /plugins/textnow-custom-field-plugin/][%d] createTNCustomFieldCreated  %+v", 201, o.Payload)
}

func (o *CreateTNCustomFieldCreated) String() string {
	return fmt.Sprintf("[POST /plugins/textnow-custom-field-plugin/][%d] createTNCustomFieldCreated  %+v", 201, o.Payload)
}

func (o *CreateTNCustomFieldCreated) GetPayload() *kbmodel.TNCustomField {
	return o.Payload
}

func (o *CreateTNCustomFieldCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.TNCustomField)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTNCustomFieldBadRequest creates a CreateTNCustomFieldBadRequest with default headers values
func NewCreateTNCustomFieldBadRequest() *CreateTNCustomFieldBadRequest {
	return &CreateTNCustomFieldBadRequest{}
}

/* CreateTNCustomFieldBadRequest describes a response with status code 400, with default header values.

Invalid data supplied
*/
type CreateTNCustomFieldBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the create t n custom field bad request response
func (o *CreateTNCustomFieldBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this create t n custom field bad request response has a 2xx status code
func (o *CreateTNCustomFieldBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create t n custom field bad request response has a 3xx status code
func (o *CreateTNCustomFieldBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create t n custom field bad request response has a 4xx status code
func (o *CreateTNCustomFieldBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create t n custom field bad request response has a 5xx status code
func (o *CreateTNCustomFieldBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create t n custom field bad request response a status code equal to that given
func (o *CreateTNCustomFieldBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateTNCustomFieldBadRequest) Error() string {
	return fmt.Sprintf("[POST /plugins/textnow-custom-field-plugin/][%d] createTNCustomFieldBadRequest ", 400)
}

func (o *CreateTNCustomFieldBadRequest) String() string {
	return fmt.Sprintf("[POST /plugins/textnow-custom-field-plugin/][%d] createTNCustomFieldBadRequest ", 400)
}

func (o *CreateTNCustomFieldBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateTNCustomFieldInternalServerError creates a CreateTNCustomFieldInternalServerError with default headers values
func NewCreateTNCustomFieldInternalServerError() *CreateTNCustomFieldInternalServerError {
	return &CreateTNCustomFieldInternalServerError{}
}

/* CreateTNCustomFieldInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type CreateTNCustomFieldInternalServerError struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the create t n custom field internal server error response
func (o *CreateTNCustomFieldInternalServerError) Code() int {
	return 500
}

// IsSuccess returns true when this create t n custom field internal server error response has a 2xx status code
func (o *CreateTNCustomFieldInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create t n custom field internal server error response has a 3xx status code
func (o *CreateTNCustomFieldInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create t n custom field internal server error response has a 4xx status code
func (o *CreateTNCustomFieldInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create t n custom field internal server error response has a 5xx status code
func (o *CreateTNCustomFieldInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create t n custom field internal server error response a status code equal to that given
func (o *CreateTNCustomFieldInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CreateTNCustomFieldInternalServerError) Error() string {
	return fmt.Sprintf("[POST /plugins/textnow-custom-field-plugin/][%d] createTNCustomFieldInternalServerError ", 500)
}

func (o *CreateTNCustomFieldInternalServerError) String() string {
	return fmt.Sprintf("[POST /plugins/textnow-custom-field-plugin/][%d] createTNCustomFieldInternalServerError ", 500)
}

func (o *CreateTNCustomFieldInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
