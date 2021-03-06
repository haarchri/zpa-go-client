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

// NewGetReauthPolicySetUsingGET1Params creates a new GetReauthPolicySetUsingGET1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetReauthPolicySetUsingGET1Params() *GetReauthPolicySetUsingGET1Params {
	return &GetReauthPolicySetUsingGET1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetReauthPolicySetUsingGET1ParamsWithTimeout creates a new GetReauthPolicySetUsingGET1Params object
// with the ability to set a timeout on a request.
func NewGetReauthPolicySetUsingGET1ParamsWithTimeout(timeout time.Duration) *GetReauthPolicySetUsingGET1Params {
	return &GetReauthPolicySetUsingGET1Params{
		timeout: timeout,
	}
}

// NewGetReauthPolicySetUsingGET1ParamsWithContext creates a new GetReauthPolicySetUsingGET1Params object
// with the ability to set a context for a request.
func NewGetReauthPolicySetUsingGET1ParamsWithContext(ctx context.Context) *GetReauthPolicySetUsingGET1Params {
	return &GetReauthPolicySetUsingGET1Params{
		Context: ctx,
	}
}

// NewGetReauthPolicySetUsingGET1ParamsWithHTTPClient creates a new GetReauthPolicySetUsingGET1Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetReauthPolicySetUsingGET1ParamsWithHTTPClient(client *http.Client) *GetReauthPolicySetUsingGET1Params {
	return &GetReauthPolicySetUsingGET1Params{
		HTTPClient: client,
	}
}

/* GetReauthPolicySetUsingGET1Params contains all the parameters to send to the API endpoint
   for the get reauth policy set using g e t 1 operation.

   Typically these are written to a http.Request.
*/
type GetReauthPolicySetUsingGET1Params struct {

	/* CustomerID.

	   The unique identifier of the ZPA tenant.
	*/
	CustomerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get reauth policy set using g e t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetReauthPolicySetUsingGET1Params) WithDefaults() *GetReauthPolicySetUsingGET1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get reauth policy set using g e t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetReauthPolicySetUsingGET1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get reauth policy set using g e t 1 params
func (o *GetReauthPolicySetUsingGET1Params) WithTimeout(timeout time.Duration) *GetReauthPolicySetUsingGET1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get reauth policy set using g e t 1 params
func (o *GetReauthPolicySetUsingGET1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get reauth policy set using g e t 1 params
func (o *GetReauthPolicySetUsingGET1Params) WithContext(ctx context.Context) *GetReauthPolicySetUsingGET1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get reauth policy set using g e t 1 params
func (o *GetReauthPolicySetUsingGET1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get reauth policy set using g e t 1 params
func (o *GetReauthPolicySetUsingGET1Params) WithHTTPClient(client *http.Client) *GetReauthPolicySetUsingGET1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get reauth policy set using g e t 1 params
func (o *GetReauthPolicySetUsingGET1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerID adds the customerID to the get reauth policy set using g e t 1 params
func (o *GetReauthPolicySetUsingGET1Params) WithCustomerID(customerID string) *GetReauthPolicySetUsingGET1Params {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the get reauth policy set using g e t 1 params
func (o *GetReauthPolicySetUsingGET1Params) SetCustomerID(customerID string) {
	o.CustomerID = customerID
}

// WriteToRequest writes these params to a swagger request
func (o *GetReauthPolicySetUsingGET1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
