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
	"github.com/go-openapi/swag"

	"github.com/killbill/kbcli/v3/kbmodel"
)

// NewUpdateAccountParams creates a new UpdateAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAccountParams() *UpdateAccountParams {
	return &UpdateAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAccountParamsWithTimeout creates a new UpdateAccountParams object
// with the ability to set a timeout on a request.
func NewUpdateAccountParamsWithTimeout(timeout time.Duration) *UpdateAccountParams {
	return &UpdateAccountParams{
		timeout: timeout,
	}
}

// NewUpdateAccountParamsWithContext creates a new UpdateAccountParams object
// with the ability to set a context for a request.
func NewUpdateAccountParamsWithContext(ctx context.Context) *UpdateAccountParams {
	return &UpdateAccountParams{
		Context: ctx,
	}
}

// NewUpdateAccountParamsWithHTTPClient creates a new UpdateAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAccountParamsWithHTTPClient(client *http.Client) *UpdateAccountParams {
	return &UpdateAccountParams{
		HTTPClient: client,
	}
}

/* UpdateAccountParams contains all the parameters to send to the API endpoint
   for the update account operation.

   Typically these are written to a http.Request.
*/
type UpdateAccountParams struct {

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

	// Body.
	Body *kbmodel.Account

	// TreatNullAsReset.
	TreatNullAsReset *bool

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the update account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAccountParams) WithDefaults() *UpdateAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAccountParams) SetDefaults() {
	var (
		treatNullAsResetDefault = bool(false)
	)

	val := UpdateAccountParams{
		TreatNullAsReset: &treatNullAsResetDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update account params
func (o *UpdateAccountParams) WithTimeout(timeout time.Duration) *UpdateAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update account params
func (o *UpdateAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update account params
func (o *UpdateAccountParams) WithContext(ctx context.Context) *UpdateAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update account params
func (o *UpdateAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update account params
func (o *UpdateAccountParams) WithHTTPClient(client *http.Client) *UpdateAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update account params
func (o *UpdateAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the update account params
func (o *UpdateAccountParams) WithXKillbillComment(xKillbillComment *string) *UpdateAccountParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the update account params
func (o *UpdateAccountParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the update account params
func (o *UpdateAccountParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *UpdateAccountParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the update account params
func (o *UpdateAccountParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the update account params
func (o *UpdateAccountParams) WithXKillbillReason(xKillbillReason *string) *UpdateAccountParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the update account params
func (o *UpdateAccountParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithAccountID adds the accountID to the update account params
func (o *UpdateAccountParams) WithAccountID(accountID strfmt.UUID) *UpdateAccountParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the update account params
func (o *UpdateAccountParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WithBody adds the body to the update account params
func (o *UpdateAccountParams) WithBody(body *kbmodel.Account) *UpdateAccountParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update account params
func (o *UpdateAccountParams) SetBody(body *kbmodel.Account) {
	o.Body = body
}

// WithTreatNullAsReset adds the treatNullAsReset to the update account params
func (o *UpdateAccountParams) WithTreatNullAsReset(treatNullAsReset *bool) *UpdateAccountParams {
	o.SetTreatNullAsReset(treatNullAsReset)
	return o
}

// SetTreatNullAsReset adds the treatNullAsReset to the update account params
func (o *UpdateAccountParams) SetTreatNullAsReset(treatNullAsReset *bool) {
	o.TreatNullAsReset = treatNullAsReset
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.TreatNullAsReset != nil {

		// query param treatNullAsReset
		var qrTreatNullAsReset bool

		if o.TreatNullAsReset != nil {
			qrTreatNullAsReset = *o.TreatNullAsReset
		}
		qTreatNullAsReset := swag.FormatBool(qrTreatNullAsReset)
		if qTreatNullAsReset != "" {

			if err := r.SetQueryParam("treatNullAsReset", qTreatNullAsReset); err != nil {
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
