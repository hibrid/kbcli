// Code generated by go-swagger; DO NOT EDIT.

package invoice_item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteInvoiceItemTagsParams creates a new DeleteInvoiceItemTagsParams object
// with the default values initialized.
func NewDeleteInvoiceItemTagsParams() *DeleteInvoiceItemTagsParams {
	var ()
	return &DeleteInvoiceItemTagsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteInvoiceItemTagsParamsWithTimeout creates a new DeleteInvoiceItemTagsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteInvoiceItemTagsParamsWithTimeout(timeout time.Duration) *DeleteInvoiceItemTagsParams {
	var ()
	return &DeleteInvoiceItemTagsParams{

		timeout: timeout,
	}
}

// NewDeleteInvoiceItemTagsParamsWithContext creates a new DeleteInvoiceItemTagsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteInvoiceItemTagsParamsWithContext(ctx context.Context) *DeleteInvoiceItemTagsParams {
	var ()
	return &DeleteInvoiceItemTagsParams{

		Context: ctx,
	}
}

// NewDeleteInvoiceItemTagsParamsWithHTTPClient creates a new DeleteInvoiceItemTagsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteInvoiceItemTagsParamsWithHTTPClient(client *http.Client) *DeleteInvoiceItemTagsParams {
	var ()
	return &DeleteInvoiceItemTagsParams{
		HTTPClient: client,
	}
}

/*DeleteInvoiceItemTagsParams contains all the parameters to send to the API endpoint
for the delete invoice item tags operation typically these are written to a http.Request
*/
type DeleteInvoiceItemTagsParams struct {

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
	/*InvoiceItemID*/
	InvoiceItemID strfmt.UUID
	/*TagDef*/
	TagDef []strfmt.UUID

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the delete invoice item tags params
func (o *DeleteInvoiceItemTagsParams) WithTimeout(timeout time.Duration) *DeleteInvoiceItemTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete invoice item tags params
func (o *DeleteInvoiceItemTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete invoice item tags params
func (o *DeleteInvoiceItemTagsParams) WithContext(ctx context.Context) *DeleteInvoiceItemTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete invoice item tags params
func (o *DeleteInvoiceItemTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete invoice item tags params
func (o *DeleteInvoiceItemTagsParams) WithHTTPClient(client *http.Client) *DeleteInvoiceItemTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete invoice item tags params
func (o *DeleteInvoiceItemTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the delete invoice item tags params
func (o *DeleteInvoiceItemTagsParams) WithXKillbillAPIKey(xKillbillAPIKey string) *DeleteInvoiceItemTagsParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the delete invoice item tags params
func (o *DeleteInvoiceItemTagsParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the delete invoice item tags params
func (o *DeleteInvoiceItemTagsParams) WithXKillbillAPISecret(xKillbillAPISecret string) *DeleteInvoiceItemTagsParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the delete invoice item tags params
func (o *DeleteInvoiceItemTagsParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithXKillbillComment adds the xKillbillComment to the delete invoice item tags params
func (o *DeleteInvoiceItemTagsParams) WithXKillbillComment(xKillbillComment *string) *DeleteInvoiceItemTagsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the delete invoice item tags params
func (o *DeleteInvoiceItemTagsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the delete invoice item tags params
func (o *DeleteInvoiceItemTagsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *DeleteInvoiceItemTagsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the delete invoice item tags params
func (o *DeleteInvoiceItemTagsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the delete invoice item tags params
func (o *DeleteInvoiceItemTagsParams) WithXKillbillReason(xKillbillReason *string) *DeleteInvoiceItemTagsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the delete invoice item tags params
func (o *DeleteInvoiceItemTagsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithInvoiceItemID adds the invoiceItemID to the delete invoice item tags params
func (o *DeleteInvoiceItemTagsParams) WithInvoiceItemID(invoiceItemID strfmt.UUID) *DeleteInvoiceItemTagsParams {
	o.SetInvoiceItemID(invoiceItemID)
	return o
}

// SetInvoiceItemID adds the invoiceItemId to the delete invoice item tags params
func (o *DeleteInvoiceItemTagsParams) SetInvoiceItemID(invoiceItemID strfmt.UUID) {
	o.InvoiceItemID = invoiceItemID
}

// WithTagDef adds the tagDef to the delete invoice item tags params
func (o *DeleteInvoiceItemTagsParams) WithTagDef(tagDef []strfmt.UUID) *DeleteInvoiceItemTagsParams {
	o.SetTagDef(tagDef)
	return o
}

// SetTagDef adds the tagDef to the delete invoice item tags params
func (o *DeleteInvoiceItemTagsParams) SetTagDef(tagDef []strfmt.UUID) {
	o.TagDef = tagDef
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteInvoiceItemTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param invoiceItemId
	if err := r.SetPathParam("invoiceItemId", o.InvoiceItemID.String()); err != nil {
		return err
	}

	var valuesTagDef []string
	for _, v := range o.TagDef {
		valuesTagDef = append(valuesTagDef, v.String())
	}

	joinedTagDef := swag.JoinByFormat(valuesTagDef, "multi")
	// query array param tagDef
	if err := r.SetQueryParam("tagDef", joinedTagDef...); err != nil {
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
