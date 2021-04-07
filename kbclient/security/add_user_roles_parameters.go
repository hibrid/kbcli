// Code generated by go-swagger; DO NOT EDIT.

package security

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

	kbmodel "github.com/killbill/kbcli/v2/kbmodel"
)

// NewAddUserRolesParams creates a new AddUserRolesParams object
// with the default values initialized.
func NewAddUserRolesParams() *AddUserRolesParams {
	var ()
	return &AddUserRolesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddUserRolesParamsWithTimeout creates a new AddUserRolesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddUserRolesParamsWithTimeout(timeout time.Duration) *AddUserRolesParams {
	var ()
	return &AddUserRolesParams{

		timeout: timeout,
	}
}

// NewAddUserRolesParamsWithContext creates a new AddUserRolesParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddUserRolesParamsWithContext(ctx context.Context) *AddUserRolesParams {
	var ()
	return &AddUserRolesParams{

		Context: ctx,
	}
}

// NewAddUserRolesParamsWithHTTPClient creates a new AddUserRolesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddUserRolesParamsWithHTTPClient(client *http.Client) *AddUserRolesParams {
	var ()
	return &AddUserRolesParams{
		HTTPClient: client,
	}
}

/*AddUserRolesParams contains all the parameters to send to the API endpoint
for the add user roles operation typically these are written to a http.Request
*/
type AddUserRolesParams struct {

	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*Body*/
	Body *kbmodel.UserRoles

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the add user roles params
func (o *AddUserRolesParams) WithTimeout(timeout time.Duration) *AddUserRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add user roles params
func (o *AddUserRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add user roles params
func (o *AddUserRolesParams) WithContext(ctx context.Context) *AddUserRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add user roles params
func (o *AddUserRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add user roles params
func (o *AddUserRolesParams) WithHTTPClient(client *http.Client) *AddUserRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add user roles params
func (o *AddUserRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the add user roles params
func (o *AddUserRolesParams) WithXKillbillComment(xKillbillComment *string) *AddUserRolesParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the add user roles params
func (o *AddUserRolesParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the add user roles params
func (o *AddUserRolesParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *AddUserRolesParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the add user roles params
func (o *AddUserRolesParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the add user roles params
func (o *AddUserRolesParams) WithXKillbillReason(xKillbillReason *string) *AddUserRolesParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the add user roles params
func (o *AddUserRolesParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the add user roles params
func (o *AddUserRolesParams) WithBody(body *kbmodel.UserRoles) *AddUserRolesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add user roles params
func (o *AddUserRolesParams) SetBody(body *kbmodel.UserRoles) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddUserRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
