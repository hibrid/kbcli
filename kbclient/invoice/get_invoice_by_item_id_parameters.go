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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetInvoiceByItemIDParams creates a new GetInvoiceByItemIDParams object
// with the default values initialized.
func NewGetInvoiceByItemIDParams() *GetInvoiceByItemIDParams {
	var (
		auditDefault             = string("NONE")
		withChildrenItemsDefault = bool(false)
		withItemsDefault         = bool(false)
	)
	return &GetInvoiceByItemIDParams{
		Audit:             &auditDefault,
		WithChildrenItems: &withChildrenItemsDefault,
		WithItems:         &withItemsDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInvoiceByItemIDParamsWithTimeout creates a new GetInvoiceByItemIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInvoiceByItemIDParamsWithTimeout(timeout time.Duration) *GetInvoiceByItemIDParams {
	var (
		auditDefault             = string("NONE")
		withChildrenItemsDefault = bool(false)
		withItemsDefault         = bool(false)
	)
	return &GetInvoiceByItemIDParams{
		Audit:             &auditDefault,
		WithChildrenItems: &withChildrenItemsDefault,
		WithItems:         &withItemsDefault,

		timeout: timeout,
	}
}

// NewGetInvoiceByItemIDParamsWithContext creates a new GetInvoiceByItemIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInvoiceByItemIDParamsWithContext(ctx context.Context) *GetInvoiceByItemIDParams {
	var (
		auditDefault             = string("NONE")
		withChildrenItemsDefault = bool(false)
		withItemsDefault         = bool(false)
	)
	return &GetInvoiceByItemIDParams{
		Audit:             &auditDefault,
		WithChildrenItems: &withChildrenItemsDefault,
		WithItems:         &withItemsDefault,

		Context: ctx,
	}
}

// NewGetInvoiceByItemIDParamsWithHTTPClient creates a new GetInvoiceByItemIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInvoiceByItemIDParamsWithHTTPClient(client *http.Client) *GetInvoiceByItemIDParams {
	var (
		auditDefault             = string("NONE")
		withChildrenItemsDefault = bool(false)
		withItemsDefault         = bool(false)
	)
	return &GetInvoiceByItemIDParams{
		Audit:             &auditDefault,
		WithChildrenItems: &withChildrenItemsDefault,
		WithItems:         &withItemsDefault,
		HTTPClient:        client,
	}
}

/*GetInvoiceByItemIDParams contains all the parameters to send to the API endpoint
for the get invoice by item Id operation typically these are written to a http.Request
*/
type GetInvoiceByItemIDParams struct {

	/*XKillbillAPIKey*/
	XKillbillAPIKey string
	/*XKillbillAPISecret*/
	XKillbillAPISecret string
	/*Audit*/
	Audit *string
	/*ItemID*/
	ItemID strfmt.UUID
	/*WithChildrenItems*/
	WithChildrenItems *bool
	/*WithItems*/
	WithItems *bool

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get invoice by item Id params
func (o *GetInvoiceByItemIDParams) WithTimeout(timeout time.Duration) *GetInvoiceByItemIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get invoice by item Id params
func (o *GetInvoiceByItemIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get invoice by item Id params
func (o *GetInvoiceByItemIDParams) WithContext(ctx context.Context) *GetInvoiceByItemIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get invoice by item Id params
func (o *GetInvoiceByItemIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get invoice by item Id params
func (o *GetInvoiceByItemIDParams) WithHTTPClient(client *http.Client) *GetInvoiceByItemIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get invoice by item Id params
func (o *GetInvoiceByItemIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the get invoice by item Id params
func (o *GetInvoiceByItemIDParams) WithXKillbillAPIKey(xKillbillAPIKey string) *GetInvoiceByItemIDParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the get invoice by item Id params
func (o *GetInvoiceByItemIDParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the get invoice by item Id params
func (o *GetInvoiceByItemIDParams) WithXKillbillAPISecret(xKillbillAPISecret string) *GetInvoiceByItemIDParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the get invoice by item Id params
func (o *GetInvoiceByItemIDParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithAudit adds the audit to the get invoice by item Id params
func (o *GetInvoiceByItemIDParams) WithAudit(audit *string) *GetInvoiceByItemIDParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get invoice by item Id params
func (o *GetInvoiceByItemIDParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithItemID adds the itemID to the get invoice by item Id params
func (o *GetInvoiceByItemIDParams) WithItemID(itemID strfmt.UUID) *GetInvoiceByItemIDParams {
	o.SetItemID(itemID)
	return o
}

// SetItemID adds the itemId to the get invoice by item Id params
func (o *GetInvoiceByItemIDParams) SetItemID(itemID strfmt.UUID) {
	o.ItemID = itemID
}

// WithWithChildrenItems adds the withChildrenItems to the get invoice by item Id params
func (o *GetInvoiceByItemIDParams) WithWithChildrenItems(withChildrenItems *bool) *GetInvoiceByItemIDParams {
	o.SetWithChildrenItems(withChildrenItems)
	return o
}

// SetWithChildrenItems adds the withChildrenItems to the get invoice by item Id params
func (o *GetInvoiceByItemIDParams) SetWithChildrenItems(withChildrenItems *bool) {
	o.WithChildrenItems = withChildrenItems
}

// WithWithItems adds the withItems to the get invoice by item Id params
func (o *GetInvoiceByItemIDParams) WithWithItems(withItems *bool) *GetInvoiceByItemIDParams {
	o.SetWithItems(withItems)
	return o
}

// SetWithItems adds the withItems to the get invoice by item Id params
func (o *GetInvoiceByItemIDParams) SetWithItems(withItems *bool) {
	o.WithItems = withItems
}

// WriteToRequest writes these params to a swagger request
func (o *GetInvoiceByItemIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Audit != nil {

		// query param audit
		var qrAudit string
		if o.Audit != nil {
			qrAudit = *o.Audit
		}
		qAudit := qrAudit
		if qAudit != "" {
			if err := r.SetQueryParam("audit", qAudit); err != nil {
				return err
			}
		}

	}

	// path param itemId
	if err := r.SetPathParam("itemId", o.ItemID.String()); err != nil {
		return err
	}

	if o.WithChildrenItems != nil {

		// query param withChildrenItems
		var qrWithChildrenItems bool
		if o.WithChildrenItems != nil {
			qrWithChildrenItems = *o.WithChildrenItems
		}
		qWithChildrenItems := swag.FormatBool(qrWithChildrenItems)
		if qWithChildrenItems != "" {
			if err := r.SetQueryParam("withChildrenItems", qWithChildrenItems); err != nil {
				return err
			}
		}

	}

	if o.WithItems != nil {

		// query param withItems
		var qrWithItems bool
		if o.WithItems != nil {
			qrWithItems = *o.WithItems
		}
		qWithItems := swag.FormatBool(qrWithItems)
		if qWithItems != "" {
			if err := r.SetQueryParam("withItems", qWithItems); err != nil {
				return err
			}
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
