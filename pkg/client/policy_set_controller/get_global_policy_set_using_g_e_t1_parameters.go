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
)

// NewGetGlobalPolicySetUsingGET1Params creates a new GetGlobalPolicySetUsingGET1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetGlobalPolicySetUsingGET1Params() *GetGlobalPolicySetUsingGET1Params {
	return &GetGlobalPolicySetUsingGET1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetGlobalPolicySetUsingGET1ParamsWithTimeout creates a new GetGlobalPolicySetUsingGET1Params object
// with the ability to set a timeout on a request.
func NewGetGlobalPolicySetUsingGET1ParamsWithTimeout(timeout time.Duration) *GetGlobalPolicySetUsingGET1Params {
	return &GetGlobalPolicySetUsingGET1Params{
		timeout: timeout,
	}
}

// NewGetGlobalPolicySetUsingGET1ParamsWithContext creates a new GetGlobalPolicySetUsingGET1Params object
// with the ability to set a context for a request.
func NewGetGlobalPolicySetUsingGET1ParamsWithContext(ctx context.Context) *GetGlobalPolicySetUsingGET1Params {
	return &GetGlobalPolicySetUsingGET1Params{
		Context: ctx,
	}
}

// NewGetGlobalPolicySetUsingGET1ParamsWithHTTPClient creates a new GetGlobalPolicySetUsingGET1Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetGlobalPolicySetUsingGET1ParamsWithHTTPClient(client *http.Client) *GetGlobalPolicySetUsingGET1Params {
	return &GetGlobalPolicySetUsingGET1Params{
		HTTPClient: client,
	}
}

/* GetGlobalPolicySetUsingGET1Params contains all the parameters to send to the API endpoint
   for the get global policy set using g e t 1 operation.

   Typically these are written to a http.Request.
*/
type GetGlobalPolicySetUsingGET1Params struct {

	/* CustomerID.

	   The unique identifier of the ZPA tenant.
	*/
	CustomerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get global policy set using g e t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGlobalPolicySetUsingGET1Params) WithDefaults() *GetGlobalPolicySetUsingGET1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get global policy set using g e t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGlobalPolicySetUsingGET1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get global policy set using g e t 1 params
func (o *GetGlobalPolicySetUsingGET1Params) WithTimeout(timeout time.Duration) *GetGlobalPolicySetUsingGET1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get global policy set using g e t 1 params
func (o *GetGlobalPolicySetUsingGET1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get global policy set using g e t 1 params
func (o *GetGlobalPolicySetUsingGET1Params) WithContext(ctx context.Context) *GetGlobalPolicySetUsingGET1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get global policy set using g e t 1 params
func (o *GetGlobalPolicySetUsingGET1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get global policy set using g e t 1 params
func (o *GetGlobalPolicySetUsingGET1Params) WithHTTPClient(client *http.Client) *GetGlobalPolicySetUsingGET1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get global policy set using g e t 1 params
func (o *GetGlobalPolicySetUsingGET1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerID adds the customerID to the get global policy set using g e t 1 params
func (o *GetGlobalPolicySetUsingGET1Params) WithCustomerID(customerID string) *GetGlobalPolicySetUsingGET1Params {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the get global policy set using g e t 1 params
func (o *GetGlobalPolicySetUsingGET1Params) SetCustomerID(customerID string) {
	o.CustomerID = customerID
}

// WriteToRequest writes these params to a swagger request
func (o *GetGlobalPolicySetUsingGET1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerId
	if err := r.SetPathParam("customerId", o.CustomerID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
