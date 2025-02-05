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

// InvalidatesCacheReader is a Reader for the InvalidatesCache structure.
type InvalidatesCacheReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InvalidatesCacheReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewInvalidatesCacheNoContent()
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

// NewInvalidatesCacheNoContent creates a InvalidatesCacheNoContent with default headers values
func NewInvalidatesCacheNoContent() *InvalidatesCacheNoContent {
	return &InvalidatesCacheNoContent{}
}

/* InvalidatesCacheNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type InvalidatesCacheNoContent struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the invalidates cache no content response
func (o *InvalidatesCacheNoContent) Code() int {
	return 204
}

// IsSuccess returns true when this invalidates cache no content response has a 2xx status code
func (o *InvalidatesCacheNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this invalidates cache no content response has a 3xx status code
func (o *InvalidatesCacheNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this invalidates cache no content response has a 4xx status code
func (o *InvalidatesCacheNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this invalidates cache no content response has a 5xx status code
func (o *InvalidatesCacheNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this invalidates cache no content response a status code equal to that given
func (o *InvalidatesCacheNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *InvalidatesCacheNoContent) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/admin/cache][%d] invalidatesCacheNoContent ", 204)
}

func (o *InvalidatesCacheNoContent) String() string {
	return fmt.Sprintf("[DELETE /1.0/kb/admin/cache][%d] invalidatesCacheNoContent ", 204)
}

func (o *InvalidatesCacheNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInvalidatesCacheBadRequest creates a InvalidatesCacheBadRequest with default headers values
func NewInvalidatesCacheBadRequest() *InvalidatesCacheBadRequest {
	return &InvalidatesCacheBadRequest{}
}

/* InvalidatesCacheBadRequest describes a response with status code 400, with default header values.

Cache name does not exist or is not alive
*/
type InvalidatesCacheBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the invalidates cache bad request response
func (o *InvalidatesCacheBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this invalidates cache bad request response has a 2xx status code
func (o *InvalidatesCacheBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this invalidates cache bad request response has a 3xx status code
func (o *InvalidatesCacheBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this invalidates cache bad request response has a 4xx status code
func (o *InvalidatesCacheBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this invalidates cache bad request response has a 5xx status code
func (o *InvalidatesCacheBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this invalidates cache bad request response a status code equal to that given
func (o *InvalidatesCacheBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *InvalidatesCacheBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/admin/cache][%d] invalidatesCacheBadRequest ", 400)
}

func (o *InvalidatesCacheBadRequest) String() string {
	return fmt.Sprintf("[DELETE /1.0/kb/admin/cache][%d] invalidatesCacheBadRequest ", 400)
}

func (o *InvalidatesCacheBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
