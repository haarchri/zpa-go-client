// Code generated by go-swagger; DO NOT EDIT.

package trusted_network_controller

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

// NewGetTrustedNetworkUsingGET1Params creates a new GetTrustedNetworkUsingGET1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTrustedNetworkUsingGET1Params() *GetTrustedNetworkUsingGET1Params {
	return &GetTrustedNetworkUsingGET1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTrustedNetworkUsingGET1ParamsWithTimeout creates a new GetTrustedNetworkUsingGET1Params object
// with the ability to set a timeout on a request.
func NewGetTrustedNetworkUsingGET1ParamsWithTimeout(timeout time.Duration) *GetTrustedNetworkUsingGET1Params {
	return &GetTrustedNetworkUsingGET1Params{
		timeout: timeout,
	}
}

// NewGetTrustedNetworkUsingGET1ParamsWithContext creates a new GetTrustedNetworkUsingGET1Params object
// with the ability to set a context for a request.
func NewGetTrustedNetworkUsingGET1ParamsWithContext(ctx context.Context) *GetTrustedNetworkUsingGET1Params {
	return &GetTrustedNetworkUsingGET1Params{
		Context: ctx,
	}
}

// NewGetTrustedNetworkUsingGET1ParamsWithHTTPClient creates a new GetTrustedNetworkUsingGET1Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetTrustedNetworkUsingGET1ParamsWithHTTPClient(client *http.Client) *GetTrustedNetworkUsingGET1Params {
	return &GetTrustedNetworkUsingGET1Params{
		HTTPClient: client,
	}
}

/* GetTrustedNetworkUsingGET1Params contains all the parameters to send to the API endpoint
   for the get trusted network using g e t 1 operation.

   Typically these are written to a http.Request.
*/
type GetTrustedNetworkUsingGET1Params struct {

	/* CustomerID.

	   The unique identifier of the ZPA tenant.

	   Format: int64
	*/
	CustomerID int64

	/* ID.

	   The unique identifier of the Trusted Network.

	   Format: int64
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get trusted network using g e t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTrustedNetworkUsingGET1Params) WithDefaults() *GetTrustedNetworkUsingGET1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get trusted network using g e t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTrustedNetworkUsingGET1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get trusted network using g e t 1 params
func (o *GetTrustedNetworkUsingGET1Params) WithTimeout(timeout time.Duration) *GetTrustedNetworkUsingGET1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get trusted network using g e t 1 params
func (o *GetTrustedNetworkUsingGET1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get trusted network using g e t 1 params
func (o *GetTrustedNetworkUsingGET1Params) WithContext(ctx context.Context) *GetTrustedNetworkUsingGET1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get trusted network using g e t 1 params
func (o *GetTrustedNetworkUsingGET1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get trusted network using g e t 1 params
func (o *GetTrustedNetworkUsingGET1Params) WithHTTPClient(client *http.Client) *GetTrustedNetworkUsingGET1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get trusted network using g e t 1 params
func (o *GetTrustedNetworkUsingGET1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerID adds the customerID to the get trusted network using g e t 1 params
func (o *GetTrustedNetworkUsingGET1Params) WithCustomerID(customerID int64) *GetTrustedNetworkUsingGET1Params {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the get trusted network using g e t 1 params
func (o *GetTrustedNetworkUsingGET1Params) SetCustomerID(customerID int64) {
	o.CustomerID = customerID
}

// WithID adds the id to the get trusted network using g e t 1 params
func (o *GetTrustedNetworkUsingGET1Params) WithID(id int64) *GetTrustedNetworkUsingGET1Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the get trusted network using g e t 1 params
func (o *GetTrustedNetworkUsingGET1Params) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetTrustedNetworkUsingGET1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerId
	if err := r.SetPathParam("customerId", swag.FormatInt64(o.CustomerID)); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}