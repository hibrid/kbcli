// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAccountParams creates a new GetAccountParams object
// with the default values initialized.
func NewGetAccountParams() *GetAccountParams {
	var (
		accountWithBalanceDefault       = bool(false)
		accountWithBalanceAndCBADefault = bool(false)
		auditDefault                    = string("NONE")
	)
	return &GetAccountParams{
		AccountWithBalance:       &accountWithBalanceDefault,
		AccountWithBalanceAndCBA: &accountWithBalanceAndCBADefault,
		Audit: &auditDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountParamsWithTimeout creates a new GetAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAccountParamsWithTimeout(timeout time.Duration) *GetAccountParams {
	var (
		accountWithBalanceDefault       = bool(false)
		accountWithBalanceAndCBADefault = bool(false)
		auditDefault                    = string("NONE")
	)
	return &GetAccountParams{
		AccountWithBalance:       &accountWithBalanceDefault,
		AccountWithBalanceAndCBA: &accountWithBalanceAndCBADefault,
		Audit: &auditDefault,

		timeout: timeout,
	}
}

// NewGetAccountParamsWithContext creates a new GetAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAccountParamsWithContext(ctx context.Context) *GetAccountParams {
	var (
		accountWithBalanceDefault       = bool(false)
		accountWithBalanceAndCBADefault = bool(false)
		auditDefault                    = string("NONE")
	)
	return &GetAccountParams{
		AccountWithBalance:       &accountWithBalanceDefault,
		AccountWithBalanceAndCBA: &accountWithBalanceAndCBADefault,
		Audit: &auditDefault,

		Context: ctx,
	}
}

// NewGetAccountParamsWithHTTPClient creates a new GetAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAccountParamsWithHTTPClient(client *http.Client) *GetAccountParams {
	var (
		accountWithBalanceDefault       = bool(false)
		accountWithBalanceAndCBADefault = bool(false)
		auditDefault                    = string("NONE")
	)
	return &GetAccountParams{
		AccountWithBalance:       &accountWithBalanceDefault,
		AccountWithBalanceAndCBA: &accountWithBalanceAndCBADefault,
		Audit:      &auditDefault,
		HTTPClient: client,
	}
}

/*GetAccountParams contains all the parameters to send to the API endpoint
for the get account operation typically these are written to a http.Request
*/
type GetAccountParams struct {

	/*XKillbillAPIKey*/
	XKillbillAPIKey string
	/*XKillbillAPISecret*/
	XKillbillAPISecret string
	/*AccountID*/
	AccountID strfmt.UUID
	/*AccountWithBalance*/
	AccountWithBalance *bool
	/*AccountWithBalanceAndCBA*/
	AccountWithBalanceAndCBA *bool
	/*Audit*/
	Audit *string

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get account params
func (o *GetAccountParams) WithTimeout(timeout time.Duration) *GetAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get account params
func (o *GetAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get account params
func (o *GetAccountParams) WithContext(ctx context.Context) *GetAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get account params
func (o *GetAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get account params
func (o *GetAccountParams) WithHTTPClient(client *http.Client) *GetAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get account params
func (o *GetAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the get account params
func (o *GetAccountParams) WithXKillbillAPIKey(xKillbillAPIKey string) *GetAccountParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the get account params
func (o *GetAccountParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the get account params
func (o *GetAccountParams) WithXKillbillAPISecret(xKillbillAPISecret string) *GetAccountParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the get account params
func (o *GetAccountParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithAccountID adds the accountID to the get account params
func (o *GetAccountParams) WithAccountID(accountID strfmt.UUID) *GetAccountParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get account params
func (o *GetAccountParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WithAccountWithBalance adds the accountWithBalance to the get account params
func (o *GetAccountParams) WithAccountWithBalance(accountWithBalance *bool) *GetAccountParams {
	o.SetAccountWithBalance(accountWithBalance)
	return o
}

// SetAccountWithBalance adds the accountWithBalance to the get account params
func (o *GetAccountParams) SetAccountWithBalance(accountWithBalance *bool) {
	o.AccountWithBalance = accountWithBalance
}

// WithAccountWithBalanceAndCBA adds the accountWithBalanceAndCBA to the get account params
func (o *GetAccountParams) WithAccountWithBalanceAndCBA(accountWithBalanceAndCBA *bool) *GetAccountParams {
	o.SetAccountWithBalanceAndCBA(accountWithBalanceAndCBA)
	return o
}

// SetAccountWithBalanceAndCBA adds the accountWithBalanceAndCBA to the get account params
func (o *GetAccountParams) SetAccountWithBalanceAndCBA(accountWithBalanceAndCBA *bool) {
	o.AccountWithBalanceAndCBA = accountWithBalanceAndCBA
}

// WithAudit adds the audit to the get account params
func (o *GetAccountParams) WithAudit(audit *string) *GetAccountParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get account params
func (o *GetAccountParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID.String()); err != nil {
		return err
	}

	if o.AccountWithBalance != nil {

		// query param accountWithBalance
		var qrAccountWithBalance bool
		if o.AccountWithBalance != nil {
			qrAccountWithBalance = *o.AccountWithBalance
		}
		qAccountWithBalance := swag.FormatBool(qrAccountWithBalance)
		if qAccountWithBalance != "" {
			if err := r.SetQueryParam("accountWithBalance", qAccountWithBalance); err != nil {
				return err
			}
		}

	}

	if o.AccountWithBalanceAndCBA != nil {

		// query param accountWithBalanceAndCBA
		var qrAccountWithBalanceAndCBA bool
		if o.AccountWithBalanceAndCBA != nil {
			qrAccountWithBalanceAndCBA = *o.AccountWithBalanceAndCBA
		}
		qAccountWithBalanceAndCBA := swag.FormatBool(qrAccountWithBalanceAndCBA)
		if qAccountWithBalanceAndCBA != "" {
			if err := r.SetQueryParam("accountWithBalanceAndCBA", qAccountWithBalanceAndCBA); err != nil {
				return err
			}
		}

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
