// Code generated by go-swagger; DO NOT EDIT.

package bundle

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v3/kbcommon"
)

// PauseBundleReader is a Reader for the PauseBundle structure.
type PauseBundleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PauseBundleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPauseBundleNoContent()
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

// NewPauseBundleNoContent creates a PauseBundleNoContent with default headers values
func NewPauseBundleNoContent() *PauseBundleNoContent {
	return &PauseBundleNoContent{}
}

/* PauseBundleNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type PauseBundleNoContent struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the pause bundle no content response
func (o *PauseBundleNoContent) Code() int {
	return 204
}

// IsSuccess returns true when this pause bundle no content response has a 2xx status code
func (o *PauseBundleNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pause bundle no content response has a 3xx status code
func (o *PauseBundleNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pause bundle no content response has a 4xx status code
func (o *PauseBundleNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this pause bundle no content response has a 5xx status code
func (o *PauseBundleNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this pause bundle no content response a status code equal to that given
func (o *PauseBundleNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *PauseBundleNoContent) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/bundles/{bundleId}/pause][%d] pauseBundleNoContent ", 204)
}

func (o *PauseBundleNoContent) String() string {
	return fmt.Sprintf("[PUT /1.0/kb/bundles/{bundleId}/pause][%d] pauseBundleNoContent ", 204)
}

func (o *PauseBundleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPauseBundleBadRequest creates a PauseBundleBadRequest with default headers values
func NewPauseBundleBadRequest() *PauseBundleBadRequest {
	return &PauseBundleBadRequest{}
}

/* PauseBundleBadRequest describes a response with status code 400, with default header values.

Invalid bundle id supplied
*/
type PauseBundleBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the pause bundle bad request response
func (o *PauseBundleBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this pause bundle bad request response has a 2xx status code
func (o *PauseBundleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pause bundle bad request response has a 3xx status code
func (o *PauseBundleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pause bundle bad request response has a 4xx status code
func (o *PauseBundleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pause bundle bad request response has a 5xx status code
func (o *PauseBundleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pause bundle bad request response a status code equal to that given
func (o *PauseBundleBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PauseBundleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/bundles/{bundleId}/pause][%d] pauseBundleBadRequest ", 400)
}

func (o *PauseBundleBadRequest) String() string {
	return fmt.Sprintf("[PUT /1.0/kb/bundles/{bundleId}/pause][%d] pauseBundleBadRequest ", 400)
}

func (o *PauseBundleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPauseBundleNotFound creates a PauseBundleNotFound with default headers values
func NewPauseBundleNotFound() *PauseBundleNotFound {
	return &PauseBundleNotFound{}
}

/* PauseBundleNotFound describes a response with status code 404, with default header values.

Bundle not found
*/
type PauseBundleNotFound struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the pause bundle not found response
func (o *PauseBundleNotFound) Code() int {
	return 404
}

// IsSuccess returns true when this pause bundle not found response has a 2xx status code
func (o *PauseBundleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pause bundle not found response has a 3xx status code
func (o *PauseBundleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pause bundle not found response has a 4xx status code
func (o *PauseBundleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pause bundle not found response has a 5xx status code
func (o *PauseBundleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pause bundle not found response a status code equal to that given
func (o *PauseBundleNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PauseBundleNotFound) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/bundles/{bundleId}/pause][%d] pauseBundleNotFound ", 404)
}

func (o *PauseBundleNotFound) String() string {
	return fmt.Sprintf("[PUT /1.0/kb/bundles/{bundleId}/pause][%d] pauseBundleNotFound ", 404)
}

func (o *PauseBundleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
