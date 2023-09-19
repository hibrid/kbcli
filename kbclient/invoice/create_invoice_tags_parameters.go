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

// NewCreateInvoiceTagsParams creates a new CreateInvoiceTagsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateInvoiceTagsParams() *CreateInvoiceTagsParams {
	return &CreateInvoiceTagsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateInvoiceTagsParamsWithTimeout creates a new CreateInvoiceTagsParams object
// with the ability to set a timeout on a request.
func NewCreateInvoiceTagsParamsWithTimeout(timeout time.Duration) *CreateInvoiceTagsParams {
	return &CreateInvoiceTagsParams{
		timeout: timeout,
	}
}

// NewCreateInvoiceTagsParamsWithContext creates a new CreateInvoiceTagsParams object
// with the ability to set a context for a request.
func NewCreateInvoiceTagsParamsWithContext(ctx context.Context) *CreateInvoiceTagsParams {
	return &CreateInvoiceTagsParams{
		Context: ctx,
	}
}

// NewCreateInvoiceTagsParamsWithHTTPClient creates a new CreateInvoiceTagsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateInvoiceTagsParamsWithHTTPClient(client *http.Client) *CreateInvoiceTagsParams {
	return &CreateInvoiceTagsParams{
		HTTPClient: client,
	}
}

/* CreateInvoiceTagsParams contains all the parameters to send to the API endpoint
   for the create invoice tags operation.

   Typically these are written to a http.Request.
*/
type CreateInvoiceTagsParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// Body.
	Body []strfmt.UUID

	// InvoiceID.
	//
	// Format: uuid
	InvoiceID strfmt.UUID

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the create invoice tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateInvoiceTagsParams) WithDefaults() *CreateInvoiceTagsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create invoice tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateInvoiceTagsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create invoice tags params
func (o *CreateInvoiceTagsParams) WithTimeout(timeout time.Duration) *CreateInvoiceTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create invoice tags params
func (o *CreateInvoiceTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create invoice tags params
func (o *CreateInvoiceTagsParams) WithContext(ctx context.Context) *CreateInvoiceTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create invoice tags params
func (o *CreateInvoiceTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create invoice tags params
func (o *CreateInvoiceTagsParams) WithHTTPClient(client *http.Client) *CreateInvoiceTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create invoice tags params
func (o *CreateInvoiceTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the create invoice tags params
func (o *CreateInvoiceTagsParams) WithXKillbillComment(xKillbillComment *string) *CreateInvoiceTagsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the create invoice tags params
func (o *CreateInvoiceTagsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the create invoice tags params
func (o *CreateInvoiceTagsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *CreateInvoiceTagsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the create invoice tags params
func (o *CreateInvoiceTagsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the create invoice tags params
func (o *CreateInvoiceTagsParams) WithXKillbillReason(xKillbillReason *string) *CreateInvoiceTagsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the create invoice tags params
func (o *CreateInvoiceTagsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the create invoice tags params
func (o *CreateInvoiceTagsParams) WithBody(body []strfmt.UUID) *CreateInvoiceTagsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create invoice tags params
func (o *CreateInvoiceTagsParams) SetBody(body []strfmt.UUID) {
	o.Body = body
}

// WithInvoiceID adds the invoiceID to the create invoice tags params
func (o *CreateInvoiceTagsParams) WithInvoiceID(invoiceID strfmt.UUID) *CreateInvoiceTagsParams {
	o.SetInvoiceID(invoiceID)
	return o
}

// SetInvoiceID adds the invoiceId to the create invoice tags params
func (o *CreateInvoiceTagsParams) SetInvoiceID(invoiceID strfmt.UUID) {
	o.InvoiceID = invoiceID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateInvoiceTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XKillbillComment != nil {

		// header param X-Killbill-Comment
		if err := r.SetHeaderParam("X-Killbill-Comment", *o.XKillbillComment); err != nil {
			return err
		}
	}

	// header param X-Killbill-CreatedBy
	if err := r.SetHeaderParam("X-Killbill-CreatedBy", o.XKillbillCreatedBy); err != nil {
		return err
	}

	if o.XKillbillReason != nil {

		// header param X-Killbill-Reason
		if err := r.SetHeaderParam("X-Killbill-Reason", *o.XKillbillReason); err != nil {
			return err
		}
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param invoiceId
	if err := r.SetPathParam("invoiceId", o.InvoiceID.String()); err != nil {
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
