// Code generated by go-swagger; DO NOT EDIT.

package bundle

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
)

// NewDeleteBundleCustomFieldsParams creates a new DeleteBundleCustomFieldsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteBundleCustomFieldsParams() *DeleteBundleCustomFieldsParams {
	return &DeleteBundleCustomFieldsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBundleCustomFieldsParamsWithTimeout creates a new DeleteBundleCustomFieldsParams object
// with the ability to set a timeout on a request.
func NewDeleteBundleCustomFieldsParamsWithTimeout(timeout time.Duration) *DeleteBundleCustomFieldsParams {
	return &DeleteBundleCustomFieldsParams{
		timeout: timeout,
	}
}

// NewDeleteBundleCustomFieldsParamsWithContext creates a new DeleteBundleCustomFieldsParams object
// with the ability to set a context for a request.
func NewDeleteBundleCustomFieldsParamsWithContext(ctx context.Context) *DeleteBundleCustomFieldsParams {
	return &DeleteBundleCustomFieldsParams{
		Context: ctx,
	}
}

// NewDeleteBundleCustomFieldsParamsWithHTTPClient creates a new DeleteBundleCustomFieldsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteBundleCustomFieldsParamsWithHTTPClient(client *http.Client) *DeleteBundleCustomFieldsParams {
	return &DeleteBundleCustomFieldsParams{
		HTTPClient: client,
	}
}

/* DeleteBundleCustomFieldsParams contains all the parameters to send to the API endpoint
   for the delete bundle custom fields operation.

   Typically these are written to a http.Request.
*/
type DeleteBundleCustomFieldsParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// BundleID.
	//
	// Format: uuid
	BundleID strfmt.UUID

	// CustomField.
	CustomField []strfmt.UUID

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the delete bundle custom fields params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBundleCustomFieldsParams) WithDefaults() *DeleteBundleCustomFieldsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete bundle custom fields params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBundleCustomFieldsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete bundle custom fields params
func (o *DeleteBundleCustomFieldsParams) WithTimeout(timeout time.Duration) *DeleteBundleCustomFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete bundle custom fields params
func (o *DeleteBundleCustomFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete bundle custom fields params
func (o *DeleteBundleCustomFieldsParams) WithContext(ctx context.Context) *DeleteBundleCustomFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete bundle custom fields params
func (o *DeleteBundleCustomFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete bundle custom fields params
func (o *DeleteBundleCustomFieldsParams) WithHTTPClient(client *http.Client) *DeleteBundleCustomFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete bundle custom fields params
func (o *DeleteBundleCustomFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the delete bundle custom fields params
func (o *DeleteBundleCustomFieldsParams) WithXKillbillComment(xKillbillComment *string) *DeleteBundleCustomFieldsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the delete bundle custom fields params
func (o *DeleteBundleCustomFieldsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the delete bundle custom fields params
func (o *DeleteBundleCustomFieldsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *DeleteBundleCustomFieldsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the delete bundle custom fields params
func (o *DeleteBundleCustomFieldsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the delete bundle custom fields params
func (o *DeleteBundleCustomFieldsParams) WithXKillbillReason(xKillbillReason *string) *DeleteBundleCustomFieldsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the delete bundle custom fields params
func (o *DeleteBundleCustomFieldsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBundleID adds the bundleID to the delete bundle custom fields params
func (o *DeleteBundleCustomFieldsParams) WithBundleID(bundleID strfmt.UUID) *DeleteBundleCustomFieldsParams {
	o.SetBundleID(bundleID)
	return o
}

// SetBundleID adds the bundleId to the delete bundle custom fields params
func (o *DeleteBundleCustomFieldsParams) SetBundleID(bundleID strfmt.UUID) {
	o.BundleID = bundleID
}

// WithCustomField adds the customField to the delete bundle custom fields params
func (o *DeleteBundleCustomFieldsParams) WithCustomField(customField []strfmt.UUID) *DeleteBundleCustomFieldsParams {
	o.SetCustomField(customField)
	return o
}

// SetCustomField adds the customField to the delete bundle custom fields params
func (o *DeleteBundleCustomFieldsParams) SetCustomField(customField []strfmt.UUID) {
	o.CustomField = customField
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBundleCustomFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param bundleId
	if err := r.SetPathParam("bundleId", o.BundleID.String()); err != nil {
		return err
	}

	if o.CustomField != nil {

		// binding items for customField
		joinedCustomField := o.bindParamCustomField(reg)

		// query array param customField
		if err := r.SetQueryParam("customField", joinedCustomField...); err != nil {
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

// bindParamDeleteBundleCustomFields binds the parameter customField
func (o *DeleteBundleCustomFieldsParams) bindParamCustomField(formats strfmt.Registry) []string {
	customFieldIR := o.CustomField

	var customFieldIC []string
	for _, customFieldIIR := range customFieldIR { // explode []strfmt.UUID

		customFieldIIV := customFieldIIR.String() // strfmt.UUID as string
		customFieldIC = append(customFieldIC, customFieldIIV)
	}

	// items.CollectionFormat: "multi"
	customFieldIS := swag.JoinByFormat(customFieldIC, "multi")

	return customFieldIS
}
