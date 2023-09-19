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

	"github.com/killbill/kbcli/v3/kbmodel"
)

// NewModifyAccountCustomFieldsParams creates a new ModifyAccountCustomFieldsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewModifyAccountCustomFieldsParams() *ModifyAccountCustomFieldsParams {
	return &ModifyAccountCustomFieldsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewModifyAccountCustomFieldsParamsWithTimeout creates a new ModifyAccountCustomFieldsParams object
// with the ability to set a timeout on a request.
func NewModifyAccountCustomFieldsParamsWithTimeout(timeout time.Duration) *ModifyAccountCustomFieldsParams {
	return &ModifyAccountCustomFieldsParams{
		timeout: timeout,
	}
}

// NewModifyAccountCustomFieldsParamsWithContext creates a new ModifyAccountCustomFieldsParams object
// with the ability to set a context for a request.
func NewModifyAccountCustomFieldsParamsWithContext(ctx context.Context) *ModifyAccountCustomFieldsParams {
	return &ModifyAccountCustomFieldsParams{
		Context: ctx,
	}
}

// NewModifyAccountCustomFieldsParamsWithHTTPClient creates a new ModifyAccountCustomFieldsParams object
// with the ability to set a custom HTTPClient for a request.
func NewModifyAccountCustomFieldsParamsWithHTTPClient(client *http.Client) *ModifyAccountCustomFieldsParams {
	return &ModifyAccountCustomFieldsParams{
		HTTPClient: client,
	}
}

/* ModifyAccountCustomFieldsParams contains all the parameters to send to the API endpoint
   for the modify account custom fields operation.

   Typically these are written to a http.Request.
*/
type ModifyAccountCustomFieldsParams struct {

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
	Body []*kbmodel.CustomField

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the modify account custom fields params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ModifyAccountCustomFieldsParams) WithDefaults() *ModifyAccountCustomFieldsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the modify account custom fields params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ModifyAccountCustomFieldsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the modify account custom fields params
func (o *ModifyAccountCustomFieldsParams) WithTimeout(timeout time.Duration) *ModifyAccountCustomFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the modify account custom fields params
func (o *ModifyAccountCustomFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the modify account custom fields params
func (o *ModifyAccountCustomFieldsParams) WithContext(ctx context.Context) *ModifyAccountCustomFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the modify account custom fields params
func (o *ModifyAccountCustomFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the modify account custom fields params
func (o *ModifyAccountCustomFieldsParams) WithHTTPClient(client *http.Client) *ModifyAccountCustomFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the modify account custom fields params
func (o *ModifyAccountCustomFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the modify account custom fields params
func (o *ModifyAccountCustomFieldsParams) WithXKillbillComment(xKillbillComment *string) *ModifyAccountCustomFieldsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the modify account custom fields params
func (o *ModifyAccountCustomFieldsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the modify account custom fields params
func (o *ModifyAccountCustomFieldsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *ModifyAccountCustomFieldsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the modify account custom fields params
func (o *ModifyAccountCustomFieldsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the modify account custom fields params
func (o *ModifyAccountCustomFieldsParams) WithXKillbillReason(xKillbillReason *string) *ModifyAccountCustomFieldsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the modify account custom fields params
func (o *ModifyAccountCustomFieldsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithAccountID adds the accountID to the modify account custom fields params
func (o *ModifyAccountCustomFieldsParams) WithAccountID(accountID strfmt.UUID) *ModifyAccountCustomFieldsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the modify account custom fields params
func (o *ModifyAccountCustomFieldsParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WithBody adds the body to the modify account custom fields params
func (o *ModifyAccountCustomFieldsParams) WithBody(body []*kbmodel.CustomField) *ModifyAccountCustomFieldsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the modify account custom fields params
func (o *ModifyAccountCustomFieldsParams) SetBody(body []*kbmodel.CustomField) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ModifyAccountCustomFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
