// Code generated by go-swagger; DO NOT EDIT.

package tag_definition

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetTagDefinitionParams creates a new GetTagDefinitionParams object
// with the default values initialized.
func NewGetTagDefinitionParams() *GetTagDefinitionParams {
	var (
		auditDefault = string("NONE")
	)
	return &GetTagDefinitionParams{
		Audit: &auditDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTagDefinitionParamsWithTimeout creates a new GetTagDefinitionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTagDefinitionParamsWithTimeout(timeout time.Duration) *GetTagDefinitionParams {
	var (
		auditDefault = string("NONE")
	)
	return &GetTagDefinitionParams{
		Audit: &auditDefault,

		timeout: timeout,
	}
}

// NewGetTagDefinitionParamsWithContext creates a new GetTagDefinitionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTagDefinitionParamsWithContext(ctx context.Context) *GetTagDefinitionParams {
	var (
		auditDefault = string("NONE")
	)
	return &GetTagDefinitionParams{
		Audit: &auditDefault,

		Context: ctx,
	}
}

// NewGetTagDefinitionParamsWithHTTPClient creates a new GetTagDefinitionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTagDefinitionParamsWithHTTPClient(client *http.Client) *GetTagDefinitionParams {
	var (
		auditDefault = string("NONE")
	)
	return &GetTagDefinitionParams{
		Audit:      &auditDefault,
		HTTPClient: client,
	}
}

/*GetTagDefinitionParams contains all the parameters to send to the API endpoint
for the get tag definition operation typically these are written to a http.Request
*/
type GetTagDefinitionParams struct {

	/*XKillbillAPIKey*/
	XKillbillAPIKey string
	/*XKillbillAPISecret*/
	XKillbillAPISecret string
	/*Audit*/
	Audit *string
	/*TagDefinitionID*/
	TagDefinitionID strfmt.UUID

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get tag definition params
func (o *GetTagDefinitionParams) WithTimeout(timeout time.Duration) *GetTagDefinitionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tag definition params
func (o *GetTagDefinitionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tag definition params
func (o *GetTagDefinitionParams) WithContext(ctx context.Context) *GetTagDefinitionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tag definition params
func (o *GetTagDefinitionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tag definition params
func (o *GetTagDefinitionParams) WithHTTPClient(client *http.Client) *GetTagDefinitionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tag definition params
func (o *GetTagDefinitionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the get tag definition params
func (o *GetTagDefinitionParams) WithXKillbillAPIKey(xKillbillAPIKey string) *GetTagDefinitionParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the get tag definition params
func (o *GetTagDefinitionParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the get tag definition params
func (o *GetTagDefinitionParams) WithXKillbillAPISecret(xKillbillAPISecret string) *GetTagDefinitionParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the get tag definition params
func (o *GetTagDefinitionParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithAudit adds the audit to the get tag definition params
func (o *GetTagDefinitionParams) WithAudit(audit *string) *GetTagDefinitionParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get tag definition params
func (o *GetTagDefinitionParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithTagDefinitionID adds the tagDefinitionID to the get tag definition params
func (o *GetTagDefinitionParams) WithTagDefinitionID(tagDefinitionID strfmt.UUID) *GetTagDefinitionParams {
	o.SetTagDefinitionID(tagDefinitionID)
	return o
}

// SetTagDefinitionID adds the tagDefinitionId to the get tag definition params
func (o *GetTagDefinitionParams) SetTagDefinitionID(tagDefinitionID strfmt.UUID) {
	o.TagDefinitionID = tagDefinitionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTagDefinitionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Killbill-ApiKey
	if err := r.SetHeaderParam("X-Killbill-ApiKey", o.XKillbillAPIKey); err != nil {
		return err
	}

	// header param X-Killbill-ApiSecret
	if err := r.SetHeaderParam("X-Killbill-ApiSecret", o.XKillbillAPISecret); err != nil {
		return err
	}

	if o.Audit != nil {

		// query param audit
		var qrAudit string
		if o.Audit != nil {
			qrAudit = *o.Audit
		}
		qAudit := qrAudit
		if qAudit != "" {
			if err := r.SetQueryParam("audit", qAudit); err != nil {
				return err
			}
		}

	}

	// path param tagDefinitionId
	if err := r.SetPathParam("tagDefinitionId", o.TagDefinitionID.String()); err != nil {
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
