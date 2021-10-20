// Code generated by go-swagger; DO NOT EDIT.

package idp_controller

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

// NewGetIdpByIDUsingGET1Params creates a new GetIdpByIDUsingGET1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetIdpByIDUsingGET1Params() *GetIdpByIDUsingGET1Params {
	return &GetIdpByIDUsingGET1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetIdpByIDUsingGET1ParamsWithTimeout creates a new GetIdpByIDUsingGET1Params object
// with the ability to set a timeout on a request.
func NewGetIdpByIDUsingGET1ParamsWithTimeout(timeout time.Duration) *GetIdpByIDUsingGET1Params {
	return &GetIdpByIDUsingGET1Params{
		timeout: timeout,
	}
}

// NewGetIdpByIDUsingGET1ParamsWithContext creates a new GetIdpByIDUsingGET1Params object
// with the ability to set a context for a request.
func NewGetIdpByIDUsingGET1ParamsWithContext(ctx context.Context) *GetIdpByIDUsingGET1Params {
	return &GetIdpByIDUsingGET1Params{
		Context: ctx,
	}
}

// NewGetIdpByIDUsingGET1ParamsWithHTTPClient creates a new GetIdpByIDUsingGET1Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetIdpByIDUsingGET1ParamsWithHTTPClient(client *http.Client) *GetIdpByIDUsingGET1Params {
	return &GetIdpByIDUsingGET1Params{
		HTTPClient: client,
	}
}

/* GetIdpByIDUsingGET1Params contains all the parameters to send to the API endpoint
   for the get idp by Id using g e t 1 operation.

   Typically these are written to a http.Request.
*/
type GetIdpByIDUsingGET1Params struct {

	/* CustomerID.

	   The unique identifier of the ZPA tenant.

	   Format: int64
	*/
	CustomerID int64

	/* IdpID.

	   The unique identifier of the IdP.

	   Format: int64
	*/
	IdpID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get idp by Id using g e t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIdpByIDUsingGET1Params) WithDefaults() *GetIdpByIDUsingGET1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get idp by Id using g e t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIdpByIDUsingGET1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get idp by Id using g e t 1 params
func (o *GetIdpByIDUsingGET1Params) WithTimeout(timeout time.Duration) *GetIdpByIDUsingGET1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get idp by Id using g e t 1 params
func (o *GetIdpByIDUsingGET1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get idp by Id using g e t 1 params
func (o *GetIdpByIDUsingGET1Params) WithContext(ctx context.Context) *GetIdpByIDUsingGET1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get idp by Id using g e t 1 params
func (o *GetIdpByIDUsingGET1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get idp by Id using g e t 1 params
func (o *GetIdpByIDUsingGET1Params) WithHTTPClient(client *http.Client) *GetIdpByIDUsingGET1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get idp by Id using g e t 1 params
func (o *GetIdpByIDUsingGET1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerID adds the customerID to the get idp by Id using g e t 1 params
func (o *GetIdpByIDUsingGET1Params) WithCustomerID(customerID int64) *GetIdpByIDUsingGET1Params {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the get idp by Id using g e t 1 params
func (o *GetIdpByIDUsingGET1Params) SetCustomerID(customerID int64) {
	o.CustomerID = customerID
}

// WithIdpID adds the idpID to the get idp by Id using g e t 1 params
func (o *GetIdpByIDUsingGET1Params) WithIdpID(idpID int64) *GetIdpByIDUsingGET1Params {
	o.SetIdpID(idpID)
	return o
}

// SetIdpID adds the idpId to the get idp by Id using g e t 1 params
func (o *GetIdpByIDUsingGET1Params) SetIdpID(idpID int64) {
	o.IdpID = idpID
}

// WriteToRequest writes these params to a swagger request
func (o *GetIdpByIDUsingGET1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerId
	if err := r.SetPathParam("customerId", swag.FormatInt64(o.CustomerID)); err != nil {
		return err
	}

	// path param idpId
	if err := r.SetPathParam("idpId", swag.FormatInt64(o.IdpID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
