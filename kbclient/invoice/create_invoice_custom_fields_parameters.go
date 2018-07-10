// Code generated by go-swagger; DO NOT EDIT.

package invoice

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

	kbmodel "github.com/killbill/kbcli/kbmodel"
)

// NewCreateInvoiceCustomFieldsParams creates a new CreateInvoiceCustomFieldsParams object
// with the default values initialized.
func NewCreateInvoiceCustomFieldsParams() *CreateInvoiceCustomFieldsParams {
	var ()
	return &CreateInvoiceCustomFieldsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateInvoiceCustomFieldsParamsWithTimeout creates a new CreateInvoiceCustomFieldsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateInvoiceCustomFieldsParamsWithTimeout(timeout time.Duration) *CreateInvoiceCustomFieldsParams {
	var ()
	return &CreateInvoiceCustomFieldsParams{

		timeout: timeout,
	}
}

// NewCreateInvoiceCustomFieldsParamsWithContext creates a new CreateInvoiceCustomFieldsParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateInvoiceCustomFieldsParamsWithContext(ctx context.Context) *CreateInvoiceCustomFieldsParams {
	var ()
	return &CreateInvoiceCustomFieldsParams{

		Context: ctx,
	}
}

// NewCreateInvoiceCustomFieldsParamsWithHTTPClient creates a new CreateInvoiceCustomFieldsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateInvoiceCustomFieldsParamsWithHTTPClient(client *http.Client) *CreateInvoiceCustomFieldsParams {
	var ()
	return &CreateInvoiceCustomFieldsParams{
		HTTPClient: client,
	}
}

/*CreateInvoiceCustomFieldsParams contains all the parameters to send to the API endpoint
for the create invoice custom fields operation typically these are written to a http.Request
*/
type CreateInvoiceCustomFieldsParams struct {

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
	Body []*kbmodel.CustomField
	/*InvoiceID*/
	InvoiceID strfmt.UUID

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the create invoice custom fields params
func (o *CreateInvoiceCustomFieldsParams) WithTimeout(timeout time.Duration) *CreateInvoiceCustomFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create invoice custom fields params
func (o *CreateInvoiceCustomFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create invoice custom fields params
func (o *CreateInvoiceCustomFieldsParams) WithContext(ctx context.Context) *CreateInvoiceCustomFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create invoice custom fields params
func (o *CreateInvoiceCustomFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create invoice custom fields params
func (o *CreateInvoiceCustomFieldsParams) WithHTTPClient(client *http.Client) *CreateInvoiceCustomFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create invoice custom fields params
func (o *CreateInvoiceCustomFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the create invoice custom fields params
func (o *CreateInvoiceCustomFieldsParams) WithXKillbillAPIKey(xKillbillAPIKey string) *CreateInvoiceCustomFieldsParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the create invoice custom fields params
func (o *CreateInvoiceCustomFieldsParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the create invoice custom fields params
func (o *CreateInvoiceCustomFieldsParams) WithXKillbillAPISecret(xKillbillAPISecret string) *CreateInvoiceCustomFieldsParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the create invoice custom fields params
func (o *CreateInvoiceCustomFieldsParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithXKillbillComment adds the xKillbillComment to the create invoice custom fields params
func (o *CreateInvoiceCustomFieldsParams) WithXKillbillComment(xKillbillComment *string) *CreateInvoiceCustomFieldsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the create invoice custom fields params
func (o *CreateInvoiceCustomFieldsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the create invoice custom fields params
func (o *CreateInvoiceCustomFieldsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *CreateInvoiceCustomFieldsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the create invoice custom fields params
func (o *CreateInvoiceCustomFieldsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the create invoice custom fields params
func (o *CreateInvoiceCustomFieldsParams) WithXKillbillReason(xKillbillReason *string) *CreateInvoiceCustomFieldsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the create invoice custom fields params
func (o *CreateInvoiceCustomFieldsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the create invoice custom fields params
func (o *CreateInvoiceCustomFieldsParams) WithBody(body []*kbmodel.CustomField) *CreateInvoiceCustomFieldsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create invoice custom fields params
func (o *CreateInvoiceCustomFieldsParams) SetBody(body []*kbmodel.CustomField) {
	o.Body = body
}

// WithInvoiceID adds the invoiceID to the create invoice custom fields params
func (o *CreateInvoiceCustomFieldsParams) WithInvoiceID(invoiceID strfmt.UUID) *CreateInvoiceCustomFieldsParams {
	o.SetInvoiceID(invoiceID)
	return o
}

// SetInvoiceID adds the invoiceId to the create invoice custom fields params
func (o *CreateInvoiceCustomFieldsParams) SetInvoiceID(invoiceID strfmt.UUID) {
	o.InvoiceID = invoiceID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateInvoiceCustomFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param invoiceId
	if err := r.SetPathParam("invoiceId", o.InvoiceID.String()); err != nil {
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
