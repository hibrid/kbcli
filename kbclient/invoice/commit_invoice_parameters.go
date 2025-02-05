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

// NewCommitInvoiceParams creates a new CommitInvoiceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCommitInvoiceParams() *CommitInvoiceParams {
	return &CommitInvoiceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCommitInvoiceParamsWithTimeout creates a new CommitInvoiceParams object
// with the ability to set a timeout on a request.
func NewCommitInvoiceParamsWithTimeout(timeout time.Duration) *CommitInvoiceParams {
	return &CommitInvoiceParams{
		timeout: timeout,
	}
}

// NewCommitInvoiceParamsWithContext creates a new CommitInvoiceParams object
// with the ability to set a context for a request.
func NewCommitInvoiceParamsWithContext(ctx context.Context) *CommitInvoiceParams {
	return &CommitInvoiceParams{
		Context: ctx,
	}
}

// NewCommitInvoiceParamsWithHTTPClient creates a new CommitInvoiceParams object
// with the ability to set a custom HTTPClient for a request.
func NewCommitInvoiceParamsWithHTTPClient(client *http.Client) *CommitInvoiceParams {
	return &CommitInvoiceParams{
		HTTPClient: client,
	}
}

/* CommitInvoiceParams contains all the parameters to send to the API endpoint
   for the commit invoice operation.

   Typically these are written to a http.Request.
*/
type CommitInvoiceParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

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

// WithDefaults hydrates default values in the commit invoice params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CommitInvoiceParams) WithDefaults() *CommitInvoiceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the commit invoice params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CommitInvoiceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the commit invoice params
func (o *CommitInvoiceParams) WithTimeout(timeout time.Duration) *CommitInvoiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the commit invoice params
func (o *CommitInvoiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the commit invoice params
func (o *CommitInvoiceParams) WithContext(ctx context.Context) *CommitInvoiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the commit invoice params
func (o *CommitInvoiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the commit invoice params
func (o *CommitInvoiceParams) WithHTTPClient(client *http.Client) *CommitInvoiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the commit invoice params
func (o *CommitInvoiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the commit invoice params
func (o *CommitInvoiceParams) WithXKillbillComment(xKillbillComment *string) *CommitInvoiceParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the commit invoice params
func (o *CommitInvoiceParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the commit invoice params
func (o *CommitInvoiceParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *CommitInvoiceParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the commit invoice params
func (o *CommitInvoiceParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the commit invoice params
func (o *CommitInvoiceParams) WithXKillbillReason(xKillbillReason *string) *CommitInvoiceParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the commit invoice params
func (o *CommitInvoiceParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithInvoiceID adds the invoiceID to the commit invoice params
func (o *CommitInvoiceParams) WithInvoiceID(invoiceID strfmt.UUID) *CommitInvoiceParams {
	o.SetInvoiceID(invoiceID)
	return o
}

// SetInvoiceID adds the invoiceId to the commit invoice params
func (o *CommitInvoiceParams) SetInvoiceID(invoiceID strfmt.UUID) {
	o.InvoiceID = invoiceID
}

// WriteToRequest writes these params to a swagger request
func (o *CommitInvoiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
