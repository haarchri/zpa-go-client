// Code generated by go-swagger; DO NOT EDIT.

package saml_attr_controller

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

// NewGetAllAttributesUsingGET3Params creates a new GetAllAttributesUsingGET3Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllAttributesUsingGET3Params() *GetAllAttributesUsingGET3Params {
	return &GetAllAttributesUsingGET3Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllAttributesUsingGET3ParamsWithTimeout creates a new GetAllAttributesUsingGET3Params object
// with the ability to set a timeout on a request.
func NewGetAllAttributesUsingGET3ParamsWithTimeout(timeout time.Duration) *GetAllAttributesUsingGET3Params {
	return &GetAllAttributesUsingGET3Params{
		timeout: timeout,
	}
}

// NewGetAllAttributesUsingGET3ParamsWithContext creates a new GetAllAttributesUsingGET3Params object
// with the ability to set a context for a request.
func NewGetAllAttributesUsingGET3ParamsWithContext(ctx context.Context) *GetAllAttributesUsingGET3Params {
	return &GetAllAttributesUsingGET3Params{
		Context: ctx,
	}
}

// NewGetAllAttributesUsingGET3ParamsWithHTTPClient creates a new GetAllAttributesUsingGET3Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllAttributesUsingGET3ParamsWithHTTPClient(client *http.Client) *GetAllAttributesUsingGET3Params {
	return &GetAllAttributesUsingGET3Params{
		HTTPClient: client,
	}
}

/* GetAllAttributesUsingGET3Params contains all the parameters to send to the API endpoint
   for the get all attributes using g e t 3 operation.

   Typically these are written to a http.Request.
*/
type GetAllAttributesUsingGET3Params struct {

	/* CustomerID.

	   The unique identifier of the ZPA tenant.

	   Format: int64
	*/
	CustomerID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get all attributes using g e t 3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllAttributesUsingGET3Params) WithDefaults() *GetAllAttributesUsingGET3Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all attributes using g e t 3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllAttributesUsingGET3Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get all attributes using g e t 3 params
func (o *GetAllAttributesUsingGET3Params) WithTimeout(timeout time.Duration) *GetAllAttributesUsingGET3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all attributes using g e t 3 params
func (o *GetAllAttributesUsingGET3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all attributes using g e t 3 params
func (o *GetAllAttributesUsingGET3Params) WithContext(ctx context.Context) *GetAllAttributesUsingGET3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all attributes using g e t 3 params
func (o *GetAllAttributesUsingGET3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all attributes using g e t 3 params
func (o *GetAllAttributesUsingGET3Params) WithHTTPClient(client *http.Client) *GetAllAttributesUsingGET3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all attributes using g e t 3 params
func (o *GetAllAttributesUsingGET3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerID adds the customerID to the get all attributes using g e t 3 params
func (o *GetAllAttributesUsingGET3Params) WithCustomerID(customerID int64) *GetAllAttributesUsingGET3Params {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the get all attributes using g e t 3 params
func (o *GetAllAttributesUsingGET3Params) SetCustomerID(customerID int64) {
	o.CustomerID = customerID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllAttributesUsingGET3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
