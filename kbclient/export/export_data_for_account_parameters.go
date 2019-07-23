// Code generated by go-swagger; DO NOT EDIT.

package export

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

// NewExportDataForAccountParams creates a new ExportDataForAccountParams object
// with the default values initialized.
func NewExportDataForAccountParams() *ExportDataForAccountParams {
	var ()
	return &ExportDataForAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExportDataForAccountParamsWithTimeout creates a new ExportDataForAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExportDataForAccountParamsWithTimeout(timeout time.Duration) *ExportDataForAccountParams {
	var ()
	return &ExportDataForAccountParams{

		timeout: timeout,
	}
}

// NewExportDataForAccountParamsWithContext creates a new ExportDataForAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewExportDataForAccountParamsWithContext(ctx context.Context) *ExportDataForAccountParams {
	var ()
	return &ExportDataForAccountParams{

		Context: ctx,
	}
}

// NewExportDataForAccountParamsWithHTTPClient creates a new ExportDataForAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExportDataForAccountParamsWithHTTPClient(client *http.Client) *ExportDataForAccountParams {
	var ()
	return &ExportDataForAccountParams{
		HTTPClient: client,
	}
}

/*ExportDataForAccountParams contains all the parameters to send to the API endpoint
for the export data for account operation typically these are written to a http.Request
*/
type ExportDataForAccountParams struct {

	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*AccountID*/
	AccountID strfmt.UUID

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the export data for account params
func (o *ExportDataForAccountParams) WithTimeout(timeout time.Duration) *ExportDataForAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the export data for account params
func (o *ExportDataForAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the export data for account params
func (o *ExportDataForAccountParams) WithContext(ctx context.Context) *ExportDataForAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the export data for account params
func (o *ExportDataForAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the export data for account params
func (o *ExportDataForAccountParams) WithHTTPClient(client *http.Client) *ExportDataForAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the export data for account params
func (o *ExportDataForAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the export data for account params
func (o *ExportDataForAccountParams) WithXKillbillComment(xKillbillComment *string) *ExportDataForAccountParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the export data for account params
func (o *ExportDataForAccountParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the export data for account params
func (o *ExportDataForAccountParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *ExportDataForAccountParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the export data for account params
func (o *ExportDataForAccountParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the export data for account params
func (o *ExportDataForAccountParams) WithXKillbillReason(xKillbillReason *string) *ExportDataForAccountParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the export data for account params
func (o *ExportDataForAccountParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithAccountID adds the accountID to the export data for account params
func (o *ExportDataForAccountParams) WithAccountID(accountID strfmt.UUID) *ExportDataForAccountParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the export data for account params
func (o *ExportDataForAccountParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WriteToRequest writes these params to a swagger request
func (o *ExportDataForAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
