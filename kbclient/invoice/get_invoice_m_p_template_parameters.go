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

// NewGetInvoiceMPTemplateParams creates a new GetInvoiceMPTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetInvoiceMPTemplateParams() *GetInvoiceMPTemplateParams {
	return &GetInvoiceMPTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetInvoiceMPTemplateParamsWithTimeout creates a new GetInvoiceMPTemplateParams object
// with the ability to set a timeout on a request.
func NewGetInvoiceMPTemplateParamsWithTimeout(timeout time.Duration) *GetInvoiceMPTemplateParams {
	return &GetInvoiceMPTemplateParams{
		timeout: timeout,
	}
}

// NewGetInvoiceMPTemplateParamsWithContext creates a new GetInvoiceMPTemplateParams object
// with the ability to set a context for a request.
func NewGetInvoiceMPTemplateParamsWithContext(ctx context.Context) *GetInvoiceMPTemplateParams {
	return &GetInvoiceMPTemplateParams{
		Context: ctx,
	}
}

// NewGetInvoiceMPTemplateParamsWithHTTPClient creates a new GetInvoiceMPTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetInvoiceMPTemplateParamsWithHTTPClient(client *http.Client) *GetInvoiceMPTemplateParams {
	return &GetInvoiceMPTemplateParams{
		HTTPClient: client,
	}
}

/* GetInvoiceMPTemplateParams contains all the parameters to send to the API endpoint
   for the get invoice m p template operation.

   Typically these are written to a http.Request.
*/
type GetInvoiceMPTemplateParams struct {

	// Locale.
	Locale string

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the get invoice m p template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInvoiceMPTemplateParams) WithDefaults() *GetInvoiceMPTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get invoice m p template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInvoiceMPTemplateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get invoice m p template params
func (o *GetInvoiceMPTemplateParams) WithTimeout(timeout time.Duration) *GetInvoiceMPTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get invoice m p template params
func (o *GetInvoiceMPTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get invoice m p template params
func (o *GetInvoiceMPTemplateParams) WithContext(ctx context.Context) *GetInvoiceMPTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get invoice m p template params
func (o *GetInvoiceMPTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get invoice m p template params
func (o *GetInvoiceMPTemplateParams) WithHTTPClient(client *http.Client) *GetInvoiceMPTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get invoice m p template params
func (o *GetInvoiceMPTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLocale adds the locale to the get invoice m p template params
func (o *GetInvoiceMPTemplateParams) WithLocale(locale string) *GetInvoiceMPTemplateParams {
	o.SetLocale(locale)
	return o
}

// SetLocale adds the locale to the get invoice m p template params
func (o *GetInvoiceMPTemplateParams) SetLocale(locale string) {
	o.Locale = locale
}

// WriteToRequest writes these params to a swagger request
func (o *GetInvoiceMPTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param locale
	if err := r.SetPathParam("locale", o.Locale); err != nil {
		return err
	}

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
