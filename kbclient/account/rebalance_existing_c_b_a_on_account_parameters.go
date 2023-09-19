// Code generated by go-swagger; DO NOT EDIT.

package account

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

// NewRebalanceExistingCBAOnAccountParams creates a new RebalanceExistingCBAOnAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRebalanceExistingCBAOnAccountParams() *RebalanceExistingCBAOnAccountParams {
	return &RebalanceExistingCBAOnAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRebalanceExistingCBAOnAccountParamsWithTimeout creates a new RebalanceExistingCBAOnAccountParams object
// with the ability to set a timeout on a request.
func NewRebalanceExistingCBAOnAccountParamsWithTimeout(timeout time.Duration) *RebalanceExistingCBAOnAccountParams {
	return &RebalanceExistingCBAOnAccountParams{
		timeout: timeout,
	}
}

// NewRebalanceExistingCBAOnAccountParamsWithContext creates a new RebalanceExistingCBAOnAccountParams object
// with the ability to set a context for a request.
func NewRebalanceExistingCBAOnAccountParamsWithContext(ctx context.Context) *RebalanceExistingCBAOnAccountParams {
	return &RebalanceExistingCBAOnAccountParams{
		Context: ctx,
	}
}

// NewRebalanceExistingCBAOnAccountParamsWithHTTPClient creates a new RebalanceExistingCBAOnAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewRebalanceExistingCBAOnAccountParamsWithHTTPClient(client *http.Client) *RebalanceExistingCBAOnAccountParams {
	return &RebalanceExistingCBAOnAccountParams{
		HTTPClient: client,
	}
}

/* RebalanceExistingCBAOnAccountParams contains all the parameters to send to the API endpoint
   for the rebalance existing c b a on account operation.

   Typically these are written to a http.Request.
*/
type RebalanceExistingCBAOnAccountParams struct {

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

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the rebalance existing c b a on account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RebalanceExistingCBAOnAccountParams) WithDefaults() *RebalanceExistingCBAOnAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the rebalance existing c b a on account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RebalanceExistingCBAOnAccountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the rebalance existing c b a on account params
func (o *RebalanceExistingCBAOnAccountParams) WithTimeout(timeout time.Duration) *RebalanceExistingCBAOnAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rebalance existing c b a on account params
func (o *RebalanceExistingCBAOnAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rebalance existing c b a on account params
func (o *RebalanceExistingCBAOnAccountParams) WithContext(ctx context.Context) *RebalanceExistingCBAOnAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rebalance existing c b a on account params
func (o *RebalanceExistingCBAOnAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rebalance existing c b a on account params
func (o *RebalanceExistingCBAOnAccountParams) WithHTTPClient(client *http.Client) *RebalanceExistingCBAOnAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rebalance existing c b a on account params
func (o *RebalanceExistingCBAOnAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the rebalance existing c b a on account params
func (o *RebalanceExistingCBAOnAccountParams) WithXKillbillComment(xKillbillComment *string) *RebalanceExistingCBAOnAccountParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the rebalance existing c b a on account params
func (o *RebalanceExistingCBAOnAccountParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the rebalance existing c b a on account params
func (o *RebalanceExistingCBAOnAccountParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *RebalanceExistingCBAOnAccountParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the rebalance existing c b a on account params
func (o *RebalanceExistingCBAOnAccountParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the rebalance existing c b a on account params
func (o *RebalanceExistingCBAOnAccountParams) WithXKillbillReason(xKillbillReason *string) *RebalanceExistingCBAOnAccountParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the rebalance existing c b a on account params
func (o *RebalanceExistingCBAOnAccountParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithAccountID adds the accountID to the rebalance existing c b a on account params
func (o *RebalanceExistingCBAOnAccountParams) WithAccountID(accountID strfmt.UUID) *RebalanceExistingCBAOnAccountParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the rebalance existing c b a on account params
func (o *RebalanceExistingCBAOnAccountParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WriteToRequest writes these params to a swagger request
func (o *RebalanceExistingCBAOnAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID.String()); err != nil {
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
