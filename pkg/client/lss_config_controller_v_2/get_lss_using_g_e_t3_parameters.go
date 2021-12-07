// Code generated by go-swagger; DO NOT EDIT.

package lss_config_controller_v_2

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

// NewGetLssUsingGET3Params creates a new GetLssUsingGET3Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLssUsingGET3Params() *GetLssUsingGET3Params {
	return &GetLssUsingGET3Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLssUsingGET3ParamsWithTimeout creates a new GetLssUsingGET3Params object
// with the ability to set a timeout on a request.
func NewGetLssUsingGET3ParamsWithTimeout(timeout time.Duration) *GetLssUsingGET3Params {
	return &GetLssUsingGET3Params{
		timeout: timeout,
	}
}

// NewGetLssUsingGET3ParamsWithContext creates a new GetLssUsingGET3Params object
// with the ability to set a context for a request.
func NewGetLssUsingGET3ParamsWithContext(ctx context.Context) *GetLssUsingGET3Params {
	return &GetLssUsingGET3Params{
		Context: ctx,
	}
}

// NewGetLssUsingGET3ParamsWithHTTPClient creates a new GetLssUsingGET3Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetLssUsingGET3ParamsWithHTTPClient(client *http.Client) *GetLssUsingGET3Params {
	return &GetLssUsingGET3Params{
		HTTPClient: client,
	}
}

/* GetLssUsingGET3Params contains all the parameters to send to the API endpoint
   for the get lss using g e t 3 operation.

   Typically these are written to a http.Request.
*/
type GetLssUsingGET3Params struct {

	/* CustomerID.

	   The unique identifier of the ZPA tenant.
	*/
	CustomerID string

	/* LssID.

	   The unique identifier of the LSS resource.
	*/
	LssID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get lss using g e t 3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLssUsingGET3Params) WithDefaults() *GetLssUsingGET3Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get lss using g e t 3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLssUsingGET3Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get lss using g e t 3 params
func (o *GetLssUsingGET3Params) WithTimeout(timeout time.Duration) *GetLssUsingGET3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lss using g e t 3 params
func (o *GetLssUsingGET3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lss using g e t 3 params
func (o *GetLssUsingGET3Params) WithContext(ctx context.Context) *GetLssUsingGET3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lss using g e t 3 params
func (o *GetLssUsingGET3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lss using g e t 3 params
func (o *GetLssUsingGET3Params) WithHTTPClient(client *http.Client) *GetLssUsingGET3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lss using g e t 3 params
func (o *GetLssUsingGET3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerID adds the customerID to the get lss using g e t 3 params
func (o *GetLssUsingGET3Params) WithCustomerID(customerID string) *GetLssUsingGET3Params {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the get lss using g e t 3 params
func (o *GetLssUsingGET3Params) SetCustomerID(customerID string) {
	o.CustomerID = customerID
}

// WithLssID adds the lssID to the get lss using g e t 3 params
func (o *GetLssUsingGET3Params) WithLssID(lssID string) *GetLssUsingGET3Params {
	o.SetLssID(lssID)
	return o
}

// SetLssID adds the lssId to the get lss using g e t 3 params
func (o *GetLssUsingGET3Params) SetLssID(lssID string) {
	o.LssID = lssID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLssUsingGET3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerId
	if err := r.SetPathParam("customerId", o.CustomerID); err != nil {
		return err
	}

	// path param lssId
	if err := r.SetPathParam("lssId", o.LssID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
