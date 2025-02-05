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

// NewDeleteCBAParams creates a new DeleteCBAParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteCBAParams() *DeleteCBAParams {
	return &DeleteCBAParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCBAParamsWithTimeout creates a new DeleteCBAParams object
// with the ability to set a timeout on a request.
func NewDeleteCBAParamsWithTimeout(timeout time.Duration) *DeleteCBAParams {
	return &DeleteCBAParams{
		timeout: timeout,
	}
}

// NewDeleteCBAParamsWithContext creates a new DeleteCBAParams object
// with the ability to set a context for a request.
func NewDeleteCBAParamsWithContext(ctx context.Context) *DeleteCBAParams {
	return &DeleteCBAParams{
		Context: ctx,
	}
}

// NewDeleteCBAParamsWithHTTPClient creates a new DeleteCBAParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteCBAParamsWithHTTPClient(client *http.Client) *DeleteCBAParams {
	return &DeleteCBAParams{
		HTTPClient: client,
	}
}

/* DeleteCBAParams contains all the parameters to send to the API endpoint
   for the delete c b a operation.

   Typically these are written to a http.Request.
*/
type DeleteCBAParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// AccountID.
	//
	// Format: uuid
	AccountID strfmt.UUID

	// InvoiceID.
	//
	// Format: uuid
	InvoiceID strfmt.UUID

	// InvoiceItemID.
	//
	// Format: uuid
	InvoiceItemID strfmt.UUID

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the delete c b a params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCBAParams) WithDefaults() *DeleteCBAParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete c b a params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCBAParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete c b a params
func (o *DeleteCBAParams) WithTimeout(timeout time.Duration) *DeleteCBAParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete c b a params
func (o *DeleteCBAParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete c b a params
func (o *DeleteCBAParams) WithContext(ctx context.Context) *DeleteCBAParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete c b a params
func (o *DeleteCBAParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete c b a params
func (o *DeleteCBAParams) WithHTTPClient(client *http.Client) *DeleteCBAParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete c b a params
func (o *DeleteCBAParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the delete c b a params
func (o *DeleteCBAParams) WithXKillbillComment(xKillbillComment *string) *DeleteCBAParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the delete c b a params
func (o *DeleteCBAParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the delete c b a params
func (o *DeleteCBAParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *DeleteCBAParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the delete c b a params
func (o *DeleteCBAParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the delete c b a params
func (o *DeleteCBAParams) WithXKillbillReason(xKillbillReason *string) *DeleteCBAParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the delete c b a params
func (o *DeleteCBAParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithAccountID adds the accountID to the delete c b a params
func (o *DeleteCBAParams) WithAccountID(accountID strfmt.UUID) *DeleteCBAParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the delete c b a params
func (o *DeleteCBAParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WithInvoiceID adds the invoiceID to the delete c b a params
func (o *DeleteCBAParams) WithInvoiceID(invoiceID strfmt.UUID) *DeleteCBAParams {
	o.SetInvoiceID(invoiceID)
	return o
}

// SetInvoiceID adds the invoiceId to the delete c b a params
func (o *DeleteCBAParams) SetInvoiceID(invoiceID strfmt.UUID) {
	o.InvoiceID = invoiceID
}

// WithInvoiceItemID adds the invoiceItemID to the delete c b a params
func (o *DeleteCBAParams) WithInvoiceItemID(invoiceItemID strfmt.UUID) *DeleteCBAParams {
	o.SetInvoiceItemID(invoiceItemID)
	return o
}

// SetInvoiceItemID adds the invoiceItemId to the delete c b a params
func (o *DeleteCBAParams) SetInvoiceItemID(invoiceItemID strfmt.UUID) {
	o.InvoiceItemID = invoiceItemID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCBAParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param accountId
	qrAccountID := o.AccountID
	qAccountID := qrAccountID.String()
	if qAccountID != "" {

		if err := r.SetQueryParam("accountId", qAccountID); err != nil {
			return err
		}
	}

	// path param invoiceId
	if err := r.SetPathParam("invoiceId", o.InvoiceID.String()); err != nil {
		return err
	}

	// path param invoiceItemId
	if err := r.SetPathParam("invoiceItemId", o.InvoiceItemID.String()); err != nil {
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
