// Code generated by go-swagger; DO NOT EDIT.

package invoice

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetInvoiceTemplateParams creates a new GetInvoiceTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetInvoiceTemplateParams() *GetInvoiceTemplateParams {
	return &GetInvoiceTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetInvoiceTemplateParamsWithTimeout creates a new GetInvoiceTemplateParams object
// with the ability to set a timeout on a request.
func NewGetInvoiceTemplateParamsWithTimeout(timeout time.Duration) *GetInvoiceTemplateParams {
	return &GetInvoiceTemplateParams{
		timeout: timeout,
	}
}

// NewGetInvoiceTemplateParamsWithContext creates a new GetInvoiceTemplateParams object
// with the ability to set a context for a request.
func NewGetInvoiceTemplateParamsWithContext(ctx context.Context) *GetInvoiceTemplateParams {
	return &GetInvoiceTemplateParams{
		Context: ctx,
	}
}

// NewGetInvoiceTemplateParamsWithHTTPClient creates a new GetInvoiceTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetInvoiceTemplateParamsWithHTTPClient(client *http.Client) *GetInvoiceTemplateParams {
	return &GetInvoiceTemplateParams{
		HTTPClient: client,
	}
}

/* GetInvoiceTemplateParams contains all the parameters to send to the API endpoint
   for the get invoice template operation.

   Typically these are written to a http.Request.
*/
type GetInvoiceTemplateParams struct {
	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the get invoice template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInvoiceTemplateParams) WithDefaults() *GetInvoiceTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get invoice template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInvoiceTemplateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get invoice template params
func (o *GetInvoiceTemplateParams) WithTimeout(timeout time.Duration) *GetInvoiceTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get invoice template params
func (o *GetInvoiceTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get invoice template params
func (o *GetInvoiceTemplateParams) WithContext(ctx context.Context) *GetInvoiceTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get invoice template params
func (o *GetInvoiceTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get invoice template params
func (o *GetInvoiceTemplateParams) WithHTTPClient(client *http.Client) *GetInvoiceTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get invoice template params
func (o *GetInvoiceTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetInvoiceTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param WithProfilingInfo
	if o.WithProfilingInfo != nil && len(*o.WithProfilingInfo) > 0 {
		if err := r.SetHeaderParam("X-Killbill-Profiling-Req", *o.WithProfilingInfo); err != nil {
			return err
		}
	}

	// header param withStackTrace
	if o.WithStackTrace != nil && *o.WithStackTrace {
		if err := r.SetQueryParam("withStackTrace", "true"); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
