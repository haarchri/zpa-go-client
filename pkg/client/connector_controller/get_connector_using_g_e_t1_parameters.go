// Code generated by go-swagger; DO NOT EDIT.

package connector_controller

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

// NewGetConnectorUsingGET1Params creates a new GetConnectorUsingGET1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetConnectorUsingGET1Params() *GetConnectorUsingGET1Params {
	return &GetConnectorUsingGET1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetConnectorUsingGET1ParamsWithTimeout creates a new GetConnectorUsingGET1Params object
// with the ability to set a timeout on a request.
func NewGetConnectorUsingGET1ParamsWithTimeout(timeout time.Duration) *GetConnectorUsingGET1Params {
	return &GetConnectorUsingGET1Params{
		timeout: timeout,
	}
}

// NewGetConnectorUsingGET1ParamsWithContext creates a new GetConnectorUsingGET1Params object
// with the ability to set a context for a request.
func NewGetConnectorUsingGET1ParamsWithContext(ctx context.Context) *GetConnectorUsingGET1Params {
	return &GetConnectorUsingGET1Params{
		Context: ctx,
	}
}

// NewGetConnectorUsingGET1ParamsWithHTTPClient creates a new GetConnectorUsingGET1Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetConnectorUsingGET1ParamsWithHTTPClient(client *http.Client) *GetConnectorUsingGET1Params {
	return &GetConnectorUsingGET1Params{
		HTTPClient: client,
	}
}

/* GetConnectorUsingGET1Params contains all the parameters to send to the API endpoint
   for the get connector using g e t 1 operation.

   Typically these are written to a http.Request.
*/
type GetConnectorUsingGET1Params struct {

	/* ConnectorID.

	   The unique identifier of the Connector.
	*/
	ConnectorID string

	/* CustomerID.

	   The unique identifier of the ZPA tenant.
	*/
	CustomerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get connector using g e t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConnectorUsingGET1Params) WithDefaults() *GetConnectorUsingGET1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get connector using g e t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConnectorUsingGET1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get connector using g e t 1 params
func (o *GetConnectorUsingGET1Params) WithTimeout(timeout time.Duration) *GetConnectorUsingGET1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get connector using g e t 1 params
func (o *GetConnectorUsingGET1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get connector using g e t 1 params
func (o *GetConnectorUsingGET1Params) WithContext(ctx context.Context) *GetConnectorUsingGET1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get connector using g e t 1 params
func (o *GetConnectorUsingGET1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get connector using g e t 1 params
func (o *GetConnectorUsingGET1Params) WithHTTPClient(client *http.Client) *GetConnectorUsingGET1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get connector using g e t 1 params
func (o *GetConnectorUsingGET1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnectorID adds the connectorID to the get connector using g e t 1 params
func (o *GetConnectorUsingGET1Params) WithConnectorID(connectorID string) *GetConnectorUsingGET1Params {
	o.SetConnectorID(connectorID)
	return o
}

// SetConnectorID adds the connectorId to the get connector using g e t 1 params
func (o *GetConnectorUsingGET1Params) SetConnectorID(connectorID string) {
	o.ConnectorID = connectorID
}

// WithCustomerID adds the customerID to the get connector using g e t 1 params
func (o *GetConnectorUsingGET1Params) WithCustomerID(customerID string) *GetConnectorUsingGET1Params {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the get connector using g e t 1 params
func (o *GetConnectorUsingGET1Params) SetCustomerID(customerID string) {
	o.CustomerID = customerID
}

// WriteToRequest writes these params to a swagger request
func (o *GetConnectorUsingGET1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param connectorId
	if err := r.SetPathParam("connectorId", o.ConnectorID); err != nil {
		return err
	}

	// path param customerId
	if err := r.SetPathParam("customerId", o.CustomerID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
