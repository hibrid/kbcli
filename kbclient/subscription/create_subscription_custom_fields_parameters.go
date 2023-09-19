// Code generated by go-swagger; DO NOT EDIT.

package subscription

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

// NewCreateSubscriptionCustomFieldsParams creates a new CreateSubscriptionCustomFieldsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateSubscriptionCustomFieldsParams() *CreateSubscriptionCustomFieldsParams {
	return &CreateSubscriptionCustomFieldsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSubscriptionCustomFieldsParamsWithTimeout creates a new CreateSubscriptionCustomFieldsParams object
// with the ability to set a timeout on a request.
func NewCreateSubscriptionCustomFieldsParamsWithTimeout(timeout time.Duration) *CreateSubscriptionCustomFieldsParams {
	return &CreateSubscriptionCustomFieldsParams{
		timeout: timeout,
	}
}

// NewCreateSubscriptionCustomFieldsParamsWithContext creates a new CreateSubscriptionCustomFieldsParams object
// with the ability to set a context for a request.
func NewCreateSubscriptionCustomFieldsParamsWithContext(ctx context.Context) *CreateSubscriptionCustomFieldsParams {
	return &CreateSubscriptionCustomFieldsParams{
		Context: ctx,
	}
}

// NewCreateSubscriptionCustomFieldsParamsWithHTTPClient creates a new CreateSubscriptionCustomFieldsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateSubscriptionCustomFieldsParamsWithHTTPClient(client *http.Client) *CreateSubscriptionCustomFieldsParams {
	return &CreateSubscriptionCustomFieldsParams{
		HTTPClient: client,
	}
}

/* CreateSubscriptionCustomFieldsParams contains all the parameters to send to the API endpoint
   for the create subscription custom fields operation.

   Typically these are written to a http.Request.
*/
type CreateSubscriptionCustomFieldsParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// Body.
	Body []*kbmodel.CustomField

	// SubscriptionID.
	//
	// Format: uuid
	SubscriptionID strfmt.UUID

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the create subscription custom fields params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSubscriptionCustomFieldsParams) WithDefaults() *CreateSubscriptionCustomFieldsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create subscription custom fields params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSubscriptionCustomFieldsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create subscription custom fields params
func (o *CreateSubscriptionCustomFieldsParams) WithTimeout(timeout time.Duration) *CreateSubscriptionCustomFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create subscription custom fields params
func (o *CreateSubscriptionCustomFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create subscription custom fields params
func (o *CreateSubscriptionCustomFieldsParams) WithContext(ctx context.Context) *CreateSubscriptionCustomFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create subscription custom fields params
func (o *CreateSubscriptionCustomFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create subscription custom fields params
func (o *CreateSubscriptionCustomFieldsParams) WithHTTPClient(client *http.Client) *CreateSubscriptionCustomFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create subscription custom fields params
func (o *CreateSubscriptionCustomFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the create subscription custom fields params
func (o *CreateSubscriptionCustomFieldsParams) WithXKillbillComment(xKillbillComment *string) *CreateSubscriptionCustomFieldsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the create subscription custom fields params
func (o *CreateSubscriptionCustomFieldsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the create subscription custom fields params
func (o *CreateSubscriptionCustomFieldsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *CreateSubscriptionCustomFieldsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the create subscription custom fields params
func (o *CreateSubscriptionCustomFieldsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the create subscription custom fields params
func (o *CreateSubscriptionCustomFieldsParams) WithXKillbillReason(xKillbillReason *string) *CreateSubscriptionCustomFieldsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the create subscription custom fields params
func (o *CreateSubscriptionCustomFieldsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the create subscription custom fields params
func (o *CreateSubscriptionCustomFieldsParams) WithBody(body []*kbmodel.CustomField) *CreateSubscriptionCustomFieldsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create subscription custom fields params
func (o *CreateSubscriptionCustomFieldsParams) SetBody(body []*kbmodel.CustomField) {
	o.Body = body
}

// WithSubscriptionID adds the subscriptionID to the create subscription custom fields params
func (o *CreateSubscriptionCustomFieldsParams) WithSubscriptionID(subscriptionID strfmt.UUID) *CreateSubscriptionCustomFieldsParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the create subscription custom fields params
func (o *CreateSubscriptionCustomFieldsParams) SetSubscriptionID(subscriptionID strfmt.UUID) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSubscriptionCustomFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param subscriptionId
	if err := r.SetPathParam("subscriptionId", o.SubscriptionID.String()); err != nil {
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
