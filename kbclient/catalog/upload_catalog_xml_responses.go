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

// UploadCatalogXMLReader is a Reader for the UploadCatalogXML structure.
type UploadCatalogXMLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadCatalogXMLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewUploadCatalogXMLCreated()
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

// NewUploadCatalogXMLCreated creates a UploadCatalogXMLCreated with default headers values
func NewUploadCatalogXMLCreated() *UploadCatalogXMLCreated {
	return &UploadCatalogXMLCreated{}
}

/* UploadCatalogXMLCreated describes a response with status code 201, with default header values.

Catalog XML created successfully
*/
type UploadCatalogXMLCreated struct {
	Payload      string
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the upload catalog Xml created response
func (o *UploadCatalogXMLCreated) Code() int {
	return 201
}

// IsSuccess returns true when this upload catalog Xml created response has a 2xx status code
func (o *UploadCatalogXMLCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this upload catalog Xml created response has a 3xx status code
func (o *UploadCatalogXMLCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload catalog Xml created response has a 4xx status code
func (o *UploadCatalogXMLCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this upload catalog Xml created response has a 5xx status code
func (o *UploadCatalogXMLCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this upload catalog Xml created response a status code equal to that given
func (o *UploadCatalogXMLCreated) IsCode(code int) bool {
	return code == 201
}

func (o *UploadCatalogXMLCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/catalog/xml][%d] uploadCatalogXmlCreated  %+v", 201, o.Payload)
}

func (o *UploadCatalogXMLCreated) String() string {
	return fmt.Sprintf("[POST /1.0/kb/catalog/xml][%d] uploadCatalogXmlCreated  %+v", 201, o.Payload)
}

func (o *UploadCatalogXMLCreated) GetPayload() string {
	return o.Payload
}

func (o *UploadCatalogXMLCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
