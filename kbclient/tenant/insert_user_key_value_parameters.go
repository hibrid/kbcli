// Code generated by go-swagger; DO NOT EDIT.

package tenant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewInsertUserKeyValueParams creates a new InsertUserKeyValueParams object
// with the default values initialized.
func NewInsertUserKeyValueParams() *InsertUserKeyValueParams {
	var ()
	return &InsertUserKeyValueParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewInsertUserKeyValueParamsWithTimeout creates a new InsertUserKeyValueParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInsertUserKeyValueParamsWithTimeout(timeout time.Duration) *InsertUserKeyValueParams {
	var ()
	return &InsertUserKeyValueParams{

		timeout: timeout,
	}
}

// NewInsertUserKeyValueParamsWithContext creates a new InsertUserKeyValueParams object
// with the default values initialized, and the ability to set a context for a request
func NewInsertUserKeyValueParamsWithContext(ctx context.Context) *InsertUserKeyValueParams {
	var ()
	return &InsertUserKeyValueParams{

		Context: ctx,
	}
}

// NewInsertUserKeyValueParamsWithHTTPClient creates a new InsertUserKeyValueParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInsertUserKeyValueParamsWithHTTPClient(client *http.Client) *InsertUserKeyValueParams {
	var ()
	return &InsertUserKeyValueParams{
		HTTPClient: client,
	}
}

/*InsertUserKeyValueParams contains all the parameters to send to the API endpoint
for the insert user key value operation typically these are written to a http.Request
*/
type InsertUserKeyValueParams struct {

	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*Body*/
	Body string
	/*KeyName*/
	KeyName string

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the insert user key value params
func (o *InsertUserKeyValueParams) WithTimeout(timeout time.Duration) *InsertUserKeyValueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the insert user key value params
func (o *InsertUserKeyValueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the insert user key value params
func (o *InsertUserKeyValueParams) WithContext(ctx context.Context) *InsertUserKeyValueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the insert user key value params
func (o *InsertUserKeyValueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the insert user key value params
func (o *InsertUserKeyValueParams) WithHTTPClient(client *http.Client) *InsertUserKeyValueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the insert user key value params
func (o *InsertUserKeyValueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the insert user key value params
func (o *InsertUserKeyValueParams) WithXKillbillComment(xKillbillComment *string) *InsertUserKeyValueParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the insert user key value params
func (o *InsertUserKeyValueParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the insert user key value params
func (o *InsertUserKeyValueParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *InsertUserKeyValueParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the insert user key value params
func (o *InsertUserKeyValueParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the insert user key value params
func (o *InsertUserKeyValueParams) WithXKillbillReason(xKillbillReason *string) *InsertUserKeyValueParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the insert user key value params
func (o *InsertUserKeyValueParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the insert user key value params
func (o *InsertUserKeyValueParams) WithBody(body string) *InsertUserKeyValueParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the insert user key value params
func (o *InsertUserKeyValueParams) SetBody(body string) {
	o.Body = body
}

// WithKeyName adds the keyName to the insert user key value params
func (o *InsertUserKeyValueParams) WithKeyName(keyName string) *InsertUserKeyValueParams {
	o.SetKeyName(keyName)
	return o
}

// SetKeyName adds the keyName to the insert user key value params
func (o *InsertUserKeyValueParams) SetKeyName(keyName string) {
	o.KeyName = keyName
}

// WriteToRequest writes these params to a swagger request
func (o *InsertUserKeyValueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param keyName
	if err := r.SetPathParam("keyName", o.KeyName); err != nil {
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
