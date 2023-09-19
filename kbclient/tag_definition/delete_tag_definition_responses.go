// Code generated by go-swagger; DO NOT EDIT.

package tag_definition

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v3/kbcommon"
)

// DeleteTagDefinitionReader is a Reader for the DeleteTagDefinition structure.
type DeleteTagDefinitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTagDefinitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteTagDefinitionNoContent()
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

// NewDeleteTagDefinitionNoContent creates a DeleteTagDefinitionNoContent with default headers values
func NewDeleteTagDefinitionNoContent() *DeleteTagDefinitionNoContent {
	return &DeleteTagDefinitionNoContent{}
}

/* DeleteTagDefinitionNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type DeleteTagDefinitionNoContent struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the delete tag definition no content response
func (o *DeleteTagDefinitionNoContent) Code() int {
	return 204
}

// IsSuccess returns true when this delete tag definition no content response has a 2xx status code
func (o *DeleteTagDefinitionNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete tag definition no content response has a 3xx status code
func (o *DeleteTagDefinitionNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete tag definition no content response has a 4xx status code
func (o *DeleteTagDefinitionNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete tag definition no content response has a 5xx status code
func (o *DeleteTagDefinitionNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete tag definition no content response a status code equal to that given
func (o *DeleteTagDefinitionNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteTagDefinitionNoContent) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/tagDefinitions/{tagDefinitionId}][%d] deleteTagDefinitionNoContent ", 204)
}

func (o *DeleteTagDefinitionNoContent) String() string {
	return fmt.Sprintf("[DELETE /1.0/kb/tagDefinitions/{tagDefinitionId}][%d] deleteTagDefinitionNoContent ", 204)
}

func (o *DeleteTagDefinitionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTagDefinitionBadRequest creates a DeleteTagDefinitionBadRequest with default headers values
func NewDeleteTagDefinitionBadRequest() *DeleteTagDefinitionBadRequest {
	return &DeleteTagDefinitionBadRequest{}
}

/* DeleteTagDefinitionBadRequest describes a response with status code 400, with default header values.

Invalid tagDefinitionId supplied
*/
type DeleteTagDefinitionBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the delete tag definition bad request response
func (o *DeleteTagDefinitionBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this delete tag definition bad request response has a 2xx status code
func (o *DeleteTagDefinitionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete tag definition bad request response has a 3xx status code
func (o *DeleteTagDefinitionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete tag definition bad request response has a 4xx status code
func (o *DeleteTagDefinitionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete tag definition bad request response has a 5xx status code
func (o *DeleteTagDefinitionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete tag definition bad request response a status code equal to that given
func (o *DeleteTagDefinitionBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteTagDefinitionBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/tagDefinitions/{tagDefinitionId}][%d] deleteTagDefinitionBadRequest ", 400)
}

func (o *DeleteTagDefinitionBadRequest) String() string {
	return fmt.Sprintf("[DELETE /1.0/kb/tagDefinitions/{tagDefinitionId}][%d] deleteTagDefinitionBadRequest ", 400)
}

func (o *DeleteTagDefinitionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
