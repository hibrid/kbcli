// Code generated by go-swagger; DO NOT EDIT.

package payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCancelScheduledPaymentTransactionByIDParams creates a new CancelScheduledPaymentTransactionByIDParams object
// with the default values initialized.
func NewCancelScheduledPaymentTransactionByIDParams() *CancelScheduledPaymentTransactionByIDParams {
	var ()
	return &CancelScheduledPaymentTransactionByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCancelScheduledPaymentTransactionByIDParamsWithTimeout creates a new CancelScheduledPaymentTransactionByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCancelScheduledPaymentTransactionByIDParamsWithTimeout(timeout time.Duration) *CancelScheduledPaymentTransactionByIDParams {
	var ()
	return &CancelScheduledPaymentTransactionByIDParams{

		timeout: timeout,
	}
}

// NewCancelScheduledPaymentTransactionByIDParamsWithContext creates a new CancelScheduledPaymentTransactionByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewCancelScheduledPaymentTransactionByIDParamsWithContext(ctx context.Context) *CancelScheduledPaymentTransactionByIDParams {
	var ()
	return &CancelScheduledPaymentTransactionByIDParams{

		Context: ctx,
	}
}

// NewCancelScheduledPaymentTransactionByIDParamsWithHTTPClient creates a new CancelScheduledPaymentTransactionByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCancelScheduledPaymentTransactionByIDParamsWithHTTPClient(client *http.Client) *CancelScheduledPaymentTransactionByIDParams {
	var ()
	return &CancelScheduledPaymentTransactionByIDParams{
		HTTPClient: client,
	}
}

/*CancelScheduledPaymentTransactionByIDParams contains all the parameters to send to the API endpoint
for the cancel scheduled payment transaction by Id operation typically these are written to a http.Request
*/
type CancelScheduledPaymentTransactionByIDParams struct {

	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*PaymentTransactionID*/
	PaymentTransactionID strfmt.UUID

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the cancel scheduled payment transaction by Id params
func (o *CancelScheduledPaymentTransactionByIDParams) WithTimeout(timeout time.Duration) *CancelScheduledPaymentTransactionByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cancel scheduled payment transaction by Id params
func (o *CancelScheduledPaymentTransactionByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cancel scheduled payment transaction by Id params
func (o *CancelScheduledPaymentTransactionByIDParams) WithContext(ctx context.Context) *CancelScheduledPaymentTransactionByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cancel scheduled payment transaction by Id params
func (o *CancelScheduledPaymentTransactionByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cancel scheduled payment transaction by Id params
func (o *CancelScheduledPaymentTransactionByIDParams) WithHTTPClient(client *http.Client) *CancelScheduledPaymentTransactionByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cancel scheduled payment transaction by Id params
func (o *CancelScheduledPaymentTransactionByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the cancel scheduled payment transaction by Id params
func (o *CancelScheduledPaymentTransactionByIDParams) WithXKillbillComment(xKillbillComment *string) *CancelScheduledPaymentTransactionByIDParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the cancel scheduled payment transaction by Id params
func (o *CancelScheduledPaymentTransactionByIDParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the cancel scheduled payment transaction by Id params
func (o *CancelScheduledPaymentTransactionByIDParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *CancelScheduledPaymentTransactionByIDParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the cancel scheduled payment transaction by Id params
func (o *CancelScheduledPaymentTransactionByIDParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the cancel scheduled payment transaction by Id params
func (o *CancelScheduledPaymentTransactionByIDParams) WithXKillbillReason(xKillbillReason *string) *CancelScheduledPaymentTransactionByIDParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the cancel scheduled payment transaction by Id params
func (o *CancelScheduledPaymentTransactionByIDParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithPaymentTransactionID adds the paymentTransactionID to the cancel scheduled payment transaction by Id params
func (o *CancelScheduledPaymentTransactionByIDParams) WithPaymentTransactionID(paymentTransactionID strfmt.UUID) *CancelScheduledPaymentTransactionByIDParams {
	o.SetPaymentTransactionID(paymentTransactionID)
	return o
}

// SetPaymentTransactionID adds the paymentTransactionId to the cancel scheduled payment transaction by Id params
func (o *CancelScheduledPaymentTransactionByIDParams) SetPaymentTransactionID(paymentTransactionID strfmt.UUID) {
	o.PaymentTransactionID = paymentTransactionID
}

// WriteToRequest writes these params to a swagger request
func (o *CancelScheduledPaymentTransactionByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param paymentTransactionId
	if err := r.SetPathParam("paymentTransactionId", o.PaymentTransactionID.String()); err != nil {
		return err
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
