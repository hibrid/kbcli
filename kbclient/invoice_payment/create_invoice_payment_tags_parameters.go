// Code generated by go-swagger; DO NOT EDIT.

package invoice_payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCreateInvoicePaymentTagsParams creates a new CreateInvoicePaymentTagsParams object
// with the default values initialized.
func NewCreateInvoicePaymentTagsParams() *CreateInvoicePaymentTagsParams {
	var ()
	return &CreateInvoicePaymentTagsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateInvoicePaymentTagsParamsWithTimeout creates a new CreateInvoicePaymentTagsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateInvoicePaymentTagsParamsWithTimeout(timeout time.Duration) *CreateInvoicePaymentTagsParams {
	var ()
	return &CreateInvoicePaymentTagsParams{

		timeout: timeout,
	}
}

// NewCreateInvoicePaymentTagsParamsWithContext creates a new CreateInvoicePaymentTagsParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateInvoicePaymentTagsParamsWithContext(ctx context.Context) *CreateInvoicePaymentTagsParams {
	var ()
	return &CreateInvoicePaymentTagsParams{

		Context: ctx,
	}
}

// NewCreateInvoicePaymentTagsParamsWithHTTPClient creates a new CreateInvoicePaymentTagsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateInvoicePaymentTagsParamsWithHTTPClient(client *http.Client) *CreateInvoicePaymentTagsParams {
	var ()
	return &CreateInvoicePaymentTagsParams{
		HTTPClient: client,
	}
}

/*CreateInvoicePaymentTagsParams contains all the parameters to send to the API endpoint
for the create invoice payment tags operation typically these are written to a http.Request
*/
type CreateInvoicePaymentTagsParams struct {

	/*XKillbillAPIKey*/
	XKillbillAPIKey string
	/*XKillbillAPISecret*/
	XKillbillAPISecret string
	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*Body*/
	Body []strfmt.UUID
	/*PaymentID*/
	PaymentID strfmt.UUID

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the create invoice payment tags params
func (o *CreateInvoicePaymentTagsParams) WithTimeout(timeout time.Duration) *CreateInvoicePaymentTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create invoice payment tags params
func (o *CreateInvoicePaymentTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create invoice payment tags params
func (o *CreateInvoicePaymentTagsParams) WithContext(ctx context.Context) *CreateInvoicePaymentTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create invoice payment tags params
func (o *CreateInvoicePaymentTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create invoice payment tags params
func (o *CreateInvoicePaymentTagsParams) WithHTTPClient(client *http.Client) *CreateInvoicePaymentTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create invoice payment tags params
func (o *CreateInvoicePaymentTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the create invoice payment tags params
func (o *CreateInvoicePaymentTagsParams) WithXKillbillAPIKey(xKillbillAPIKey string) *CreateInvoicePaymentTagsParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the create invoice payment tags params
func (o *CreateInvoicePaymentTagsParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the create invoice payment tags params
func (o *CreateInvoicePaymentTagsParams) WithXKillbillAPISecret(xKillbillAPISecret string) *CreateInvoicePaymentTagsParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the create invoice payment tags params
func (o *CreateInvoicePaymentTagsParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithXKillbillComment adds the xKillbillComment to the create invoice payment tags params
func (o *CreateInvoicePaymentTagsParams) WithXKillbillComment(xKillbillComment *string) *CreateInvoicePaymentTagsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the create invoice payment tags params
func (o *CreateInvoicePaymentTagsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the create invoice payment tags params
func (o *CreateInvoicePaymentTagsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *CreateInvoicePaymentTagsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the create invoice payment tags params
func (o *CreateInvoicePaymentTagsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the create invoice payment tags params
func (o *CreateInvoicePaymentTagsParams) WithXKillbillReason(xKillbillReason *string) *CreateInvoicePaymentTagsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the create invoice payment tags params
func (o *CreateInvoicePaymentTagsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the create invoice payment tags params
func (o *CreateInvoicePaymentTagsParams) WithBody(body []strfmt.UUID) *CreateInvoicePaymentTagsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create invoice payment tags params
func (o *CreateInvoicePaymentTagsParams) SetBody(body []strfmt.UUID) {
	o.Body = body
}

// WithPaymentID adds the paymentID to the create invoice payment tags params
func (o *CreateInvoicePaymentTagsParams) WithPaymentID(paymentID strfmt.UUID) *CreateInvoicePaymentTagsParams {
	o.SetPaymentID(paymentID)
	return o
}

// SetPaymentID adds the paymentId to the create invoice payment tags params
func (o *CreateInvoicePaymentTagsParams) SetPaymentID(paymentID strfmt.UUID) {
	o.PaymentID = paymentID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateInvoicePaymentTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Killbill-ApiKey
	if err := r.SetHeaderParam("X-Killbill-ApiKey", o.XKillbillAPIKey); err != nil {
		return err
	}

	// header param X-Killbill-ApiSecret
	if err := r.SetHeaderParam("X-Killbill-ApiSecret", o.XKillbillAPISecret); err != nil {
		return err
	}

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

	// path param paymentId
	if err := r.SetPathParam("paymentId", o.PaymentID.String()); err != nil {
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
