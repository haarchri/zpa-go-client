// Code generated by go-swagger; DO NOT EDIT.

package policy_set_controller

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

// NewGetBypassPolicySetUsingGET1Params creates a new GetBypassPolicySetUsingGET1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBypassPolicySetUsingGET1Params() *GetBypassPolicySetUsingGET1Params {
	return &GetBypassPolicySetUsingGET1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBypassPolicySetUsingGET1ParamsWithTimeout creates a new GetBypassPolicySetUsingGET1Params object
// with the ability to set a timeout on a request.
func NewGetBypassPolicySetUsingGET1ParamsWithTimeout(timeout time.Duration) *GetBypassPolicySetUsingGET1Params {
	return &GetBypassPolicySetUsingGET1Params{
		timeout: timeout,
	}
}

// NewGetBypassPolicySetUsingGET1ParamsWithContext creates a new GetBypassPolicySetUsingGET1Params object
// with the ability to set a context for a request.
func NewGetBypassPolicySetUsingGET1ParamsWithContext(ctx context.Context) *GetBypassPolicySetUsingGET1Params {
	return &GetBypassPolicySetUsingGET1Params{
		Context: ctx,
	}
}

// NewGetBypassPolicySetUsingGET1ParamsWithHTTPClient creates a new GetBypassPolicySetUsingGET1Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetBypassPolicySetUsingGET1ParamsWithHTTPClient(client *http.Client) *GetBypassPolicySetUsingGET1Params {
	return &GetBypassPolicySetUsingGET1Params{
		HTTPClient: client,
	}
}

/* GetBypassPolicySetUsingGET1Params contains all the parameters to send to the API endpoint
   for the get bypass policy set using g e t 1 operation.

   Typically these are written to a http.Request.
*/
type GetBypassPolicySetUsingGET1Params struct {

	/* CustomerID.

	   The unique identifier of the ZPA tenant.

	   Format: int64
	*/
	CustomerID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get bypass policy set using g e t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBypassPolicySetUsingGET1Params) WithDefaults() *GetBypassPolicySetUsingGET1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get bypass policy set using g e t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBypassPolicySetUsingGET1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get bypass policy set using g e t 1 params
func (o *GetBypassPolicySetUsingGET1Params) WithTimeout(timeout time.Duration) *GetBypassPolicySetUsingGET1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get bypass policy set using g e t 1 params
func (o *GetBypassPolicySetUsingGET1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get bypass policy set using g e t 1 params
func (o *GetBypassPolicySetUsingGET1Params) WithContext(ctx context.Context) *GetBypassPolicySetUsingGET1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get bypass policy set using g e t 1 params
func (o *GetBypassPolicySetUsingGET1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get bypass policy set using g e t 1 params
func (o *GetBypassPolicySetUsingGET1Params) WithHTTPClient(client *http.Client) *GetBypassPolicySetUsingGET1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get bypass policy set using g e t 1 params
func (o *GetBypassPolicySetUsingGET1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerID adds the customerID to the get bypass policy set using g e t 1 params
func (o *GetBypassPolicySetUsingGET1Params) WithCustomerID(customerID int64) *GetBypassPolicySetUsingGET1Params {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the get bypass policy set using g e t 1 params
func (o *GetBypassPolicySetUsingGET1Params) SetCustomerID(customerID int64) {
	o.CustomerID = customerID
}

// WriteToRequest writes these params to a swagger request
func (o *GetBypassPolicySetUsingGET1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerId
	if err := r.SetPathParam("customerId", swag.FormatInt64(o.CustomerID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
