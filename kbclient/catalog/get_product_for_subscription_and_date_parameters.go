// Code generated by go-swagger; DO NOT EDIT.

package catalog

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

// NewGetProductForSubscriptionAndDateParams creates a new GetProductForSubscriptionAndDateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProductForSubscriptionAndDateParams() *GetProductForSubscriptionAndDateParams {
	return &GetProductForSubscriptionAndDateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProductForSubscriptionAndDateParamsWithTimeout creates a new GetProductForSubscriptionAndDateParams object
// with the ability to set a timeout on a request.
func NewGetProductForSubscriptionAndDateParamsWithTimeout(timeout time.Duration) *GetProductForSubscriptionAndDateParams {
	return &GetProductForSubscriptionAndDateParams{
		timeout: timeout,
	}
}

// NewGetProductForSubscriptionAndDateParamsWithContext creates a new GetProductForSubscriptionAndDateParams object
// with the ability to set a context for a request.
func NewGetProductForSubscriptionAndDateParamsWithContext(ctx context.Context) *GetProductForSubscriptionAndDateParams {
	return &GetProductForSubscriptionAndDateParams{
		Context: ctx,
	}
}

// NewGetProductForSubscriptionAndDateParamsWithHTTPClient creates a new GetProductForSubscriptionAndDateParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProductForSubscriptionAndDateParamsWithHTTPClient(client *http.Client) *GetProductForSubscriptionAndDateParams {
	return &GetProductForSubscriptionAndDateParams{
		HTTPClient: client,
	}
}

/* GetProductForSubscriptionAndDateParams contains all the parameters to send to the API endpoint
   for the get product for subscription and date operation.

   Typically these are written to a http.Request.
*/
type GetProductForSubscriptionAndDateParams struct {

	// RequestedDate.
	//
	// Format: date
	RequestedDate *strfmt.Date

	// SubscriptionID.
	//
	// Format: uuid
	SubscriptionID *strfmt.UUID

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the get product for subscription and date params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProductForSubscriptionAndDateParams) WithDefaults() *GetProductForSubscriptionAndDateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get product for subscription and date params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProductForSubscriptionAndDateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get product for subscription and date params
func (o *GetProductForSubscriptionAndDateParams) WithTimeout(timeout time.Duration) *GetProductForSubscriptionAndDateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get product for subscription and date params
func (o *GetProductForSubscriptionAndDateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get product for subscription and date params
func (o *GetProductForSubscriptionAndDateParams) WithContext(ctx context.Context) *GetProductForSubscriptionAndDateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get product for subscription and date params
func (o *GetProductForSubscriptionAndDateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get product for subscription and date params
func (o *GetProductForSubscriptionAndDateParams) WithHTTPClient(client *http.Client) *GetProductForSubscriptionAndDateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get product for subscription and date params
func (o *GetProductForSubscriptionAndDateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestedDate adds the requestedDate to the get product for subscription and date params
func (o *GetProductForSubscriptionAndDateParams) WithRequestedDate(requestedDate *strfmt.Date) *GetProductForSubscriptionAndDateParams {
	o.SetRequestedDate(requestedDate)
	return o
}

// SetRequestedDate adds the requestedDate to the get product for subscription and date params
func (o *GetProductForSubscriptionAndDateParams) SetRequestedDate(requestedDate *strfmt.Date) {
	o.RequestedDate = requestedDate
}

// WithSubscriptionID adds the subscriptionID to the get product for subscription and date params
func (o *GetProductForSubscriptionAndDateParams) WithSubscriptionID(subscriptionID *strfmt.UUID) *GetProductForSubscriptionAndDateParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the get product for subscription and date params
func (o *GetProductForSubscriptionAndDateParams) SetSubscriptionID(subscriptionID *strfmt.UUID) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetProductForSubscriptionAndDateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.RequestedDate != nil {

		// query param requestedDate
		var qrRequestedDate strfmt.Date

		if o.RequestedDate != nil {
			qrRequestedDate = *o.RequestedDate
		}
		qRequestedDate := qrRequestedDate.String()
		if qRequestedDate != "" {

			if err := r.SetQueryParam("requestedDate", qRequestedDate); err != nil {
				return err
			}
		}
	}

	if o.SubscriptionID != nil {

		// query param subscriptionId
		var qrSubscriptionID strfmt.UUID

		if o.SubscriptionID != nil {
			qrSubscriptionID = *o.SubscriptionID
		}
		qSubscriptionID := qrSubscriptionID.String()
		if qSubscriptionID != "" {

			if err := r.SetQueryParam("subscriptionId", qSubscriptionID); err != nil {
				return err
			}
		}
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
