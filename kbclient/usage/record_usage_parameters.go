// Code generated by go-swagger; DO NOT EDIT.

package usage

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

	"github.com/killbill/kbcli/v3/kbmodel"
)

// NewRecordUsageParams creates a new RecordUsageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRecordUsageParams() *RecordUsageParams {
	return &RecordUsageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRecordUsageParamsWithTimeout creates a new RecordUsageParams object
// with the ability to set a timeout on a request.
func NewRecordUsageParamsWithTimeout(timeout time.Duration) *RecordUsageParams {
	return &RecordUsageParams{
		timeout: timeout,
	}
}

// NewRecordUsageParamsWithContext creates a new RecordUsageParams object
// with the ability to set a context for a request.
func NewRecordUsageParamsWithContext(ctx context.Context) *RecordUsageParams {
	return &RecordUsageParams{
		Context: ctx,
	}
}

// NewRecordUsageParamsWithHTTPClient creates a new RecordUsageParams object
// with the ability to set a custom HTTPClient for a request.
func NewRecordUsageParamsWithHTTPClient(client *http.Client) *RecordUsageParams {
	return &RecordUsageParams{
		HTTPClient: client,
	}
}

/* RecordUsageParams contains all the parameters to send to the API endpoint
   for the record usage operation.

   Typically these are written to a http.Request.
*/
type RecordUsageParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// Body.
	Body *kbmodel.SubscriptionUsageRecord

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the record usage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RecordUsageParams) WithDefaults() *RecordUsageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the record usage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RecordUsageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the record usage params
func (o *RecordUsageParams) WithTimeout(timeout time.Duration) *RecordUsageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the record usage params
func (o *RecordUsageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the record usage params
func (o *RecordUsageParams) WithContext(ctx context.Context) *RecordUsageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the record usage params
func (o *RecordUsageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the record usage params
func (o *RecordUsageParams) WithHTTPClient(client *http.Client) *RecordUsageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the record usage params
func (o *RecordUsageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the record usage params
func (o *RecordUsageParams) WithXKillbillComment(xKillbillComment *string) *RecordUsageParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the record usage params
func (o *RecordUsageParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the record usage params
func (o *RecordUsageParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *RecordUsageParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the record usage params
func (o *RecordUsageParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the record usage params
func (o *RecordUsageParams) WithXKillbillReason(xKillbillReason *string) *RecordUsageParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the record usage params
func (o *RecordUsageParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the record usage params
func (o *RecordUsageParams) WithBody(body *kbmodel.SubscriptionUsageRecord) *RecordUsageParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the record usage params
func (o *RecordUsageParams) SetBody(body *kbmodel.SubscriptionUsageRecord) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RecordUsageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
