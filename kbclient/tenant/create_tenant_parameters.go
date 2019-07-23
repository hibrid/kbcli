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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/kbmodel"
)

// NewCreateTenantParams creates a new CreateTenantParams object
// with the default values initialized.
func NewCreateTenantParams() *CreateTenantParams {
	var (
		useGlobalDefaultDefault = bool(false)
	)
	return &CreateTenantParams{
		UseGlobalDefault: &useGlobalDefaultDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTenantParamsWithTimeout creates a new CreateTenantParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateTenantParamsWithTimeout(timeout time.Duration) *CreateTenantParams {
	var (
		useGlobalDefaultDefault = bool(false)
	)
	return &CreateTenantParams{
		UseGlobalDefault: &useGlobalDefaultDefault,

		timeout: timeout,
	}
}

// NewCreateTenantParamsWithContext creates a new CreateTenantParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateTenantParamsWithContext(ctx context.Context) *CreateTenantParams {
	var (
		useGlobalDefaultDefault = bool(false)
	)
	return &CreateTenantParams{
		UseGlobalDefault: &useGlobalDefaultDefault,

		Context: ctx,
	}
}

// NewCreateTenantParamsWithHTTPClient creates a new CreateTenantParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateTenantParamsWithHTTPClient(client *http.Client) *CreateTenantParams {
	var (
		useGlobalDefaultDefault = bool(false)
	)
	return &CreateTenantParams{
		UseGlobalDefault: &useGlobalDefaultDefault,
		HTTPClient:       client,
	}
}

/*CreateTenantParams contains all the parameters to send to the API endpoint
for the create tenant operation typically these are written to a http.Request
*/
type CreateTenantParams struct {

	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*Body*/
	Body *kbmodel.Tenant
	/*UseGlobalDefault*/
	UseGlobalDefault *bool

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the create tenant params
func (o *CreateTenantParams) WithTimeout(timeout time.Duration) *CreateTenantParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create tenant params
func (o *CreateTenantParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create tenant params
func (o *CreateTenantParams) WithContext(ctx context.Context) *CreateTenantParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create tenant params
func (o *CreateTenantParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create tenant params
func (o *CreateTenantParams) WithHTTPClient(client *http.Client) *CreateTenantParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create tenant params
func (o *CreateTenantParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the create tenant params
func (o *CreateTenantParams) WithXKillbillComment(xKillbillComment *string) *CreateTenantParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the create tenant params
func (o *CreateTenantParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the create tenant params
func (o *CreateTenantParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *CreateTenantParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the create tenant params
func (o *CreateTenantParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the create tenant params
func (o *CreateTenantParams) WithXKillbillReason(xKillbillReason *string) *CreateTenantParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the create tenant params
func (o *CreateTenantParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the create tenant params
func (o *CreateTenantParams) WithBody(body *kbmodel.Tenant) *CreateTenantParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create tenant params
func (o *CreateTenantParams) SetBody(body *kbmodel.Tenant) {
	o.Body = body
}

// WithUseGlobalDefault adds the useGlobalDefault to the create tenant params
func (o *CreateTenantParams) WithUseGlobalDefault(useGlobalDefault *bool) *CreateTenantParams {
	o.SetUseGlobalDefault(useGlobalDefault)
	return o
}

// SetUseGlobalDefault adds the useGlobalDefault to the create tenant params
func (o *CreateTenantParams) SetUseGlobalDefault(useGlobalDefault *bool) {
	o.UseGlobalDefault = useGlobalDefault
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTenantParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.UseGlobalDefault != nil {

		// query param useGlobalDefault
		var qrUseGlobalDefault bool
		if o.UseGlobalDefault != nil {
			qrUseGlobalDefault = *o.UseGlobalDefault
		}
		qUseGlobalDefault := swag.FormatBool(qrUseGlobalDefault)
		if qUseGlobalDefault != "" {
			if err := r.SetQueryParam("useGlobalDefault", qUseGlobalDefault); err != nil {
				return err
			}
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
