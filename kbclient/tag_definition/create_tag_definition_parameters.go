// Code generated by go-swagger; DO NOT EDIT.

package tag_definition

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

// NewCreateTagDefinitionParams creates a new CreateTagDefinitionParams object
// with the default values initialized.
func NewCreateTagDefinitionParams() *CreateTagDefinitionParams {
	var ()
	return &CreateTagDefinitionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTagDefinitionParamsWithTimeout creates a new CreateTagDefinitionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateTagDefinitionParamsWithTimeout(timeout time.Duration) *CreateTagDefinitionParams {
	var ()
	return &CreateTagDefinitionParams{

		timeout: timeout,
	}
}

// NewCreateTagDefinitionParamsWithContext creates a new CreateTagDefinitionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateTagDefinitionParamsWithContext(ctx context.Context) *CreateTagDefinitionParams {
	var ()
	return &CreateTagDefinitionParams{

		Context: ctx,
	}
}

// NewCreateTagDefinitionParamsWithHTTPClient creates a new CreateTagDefinitionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateTagDefinitionParamsWithHTTPClient(client *http.Client) *CreateTagDefinitionParams {
	var ()
	return &CreateTagDefinitionParams{
		HTTPClient: client,
	}
}

/*CreateTagDefinitionParams contains all the parameters to send to the API endpoint
for the create tag definition operation typically these are written to a http.Request
*/
type CreateTagDefinitionParams struct {

	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*Body*/
	Body *kbmodel.TagDefinition

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the create tag definition params
func (o *CreateTagDefinitionParams) WithTimeout(timeout time.Duration) *CreateTagDefinitionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create tag definition params
func (o *CreateTagDefinitionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create tag definition params
func (o *CreateTagDefinitionParams) WithContext(ctx context.Context) *CreateTagDefinitionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create tag definition params
func (o *CreateTagDefinitionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create tag definition params
func (o *CreateTagDefinitionParams) WithHTTPClient(client *http.Client) *CreateTagDefinitionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create tag definition params
func (o *CreateTagDefinitionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the create tag definition params
func (o *CreateTagDefinitionParams) WithXKillbillComment(xKillbillComment *string) *CreateTagDefinitionParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the create tag definition params
func (o *CreateTagDefinitionParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the create tag definition params
func (o *CreateTagDefinitionParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *CreateTagDefinitionParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the create tag definition params
func (o *CreateTagDefinitionParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the create tag definition params
func (o *CreateTagDefinitionParams) WithXKillbillReason(xKillbillReason *string) *CreateTagDefinitionParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the create tag definition params
func (o *CreateTagDefinitionParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the create tag definition params
func (o *CreateTagDefinitionParams) WithBody(body *kbmodel.TagDefinition) *CreateTagDefinitionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create tag definition params
func (o *CreateTagDefinitionParams) SetBody(body *kbmodel.TagDefinition) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTagDefinitionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
