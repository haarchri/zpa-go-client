// Code generated by go-swagger; DO NOT EDIT.

package service_edge_group_controller

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

// NewGetServiceEdgeGroupUsingGETParams creates a new GetServiceEdgeGroupUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetServiceEdgeGroupUsingGETParams() *GetServiceEdgeGroupUsingGETParams {
	return &GetServiceEdgeGroupUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetServiceEdgeGroupUsingGETParamsWithTimeout creates a new GetServiceEdgeGroupUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetServiceEdgeGroupUsingGETParamsWithTimeout(timeout time.Duration) *GetServiceEdgeGroupUsingGETParams {
	return &GetServiceEdgeGroupUsingGETParams{
		timeout: timeout,
	}
}

// NewGetServiceEdgeGroupUsingGETParamsWithContext creates a new GetServiceEdgeGroupUsingGETParams object
// with the ability to set a context for a request.
func NewGetServiceEdgeGroupUsingGETParamsWithContext(ctx context.Context) *GetServiceEdgeGroupUsingGETParams {
	return &GetServiceEdgeGroupUsingGETParams{
		Context: ctx,
	}
}

// NewGetServiceEdgeGroupUsingGETParamsWithHTTPClient creates a new GetServiceEdgeGroupUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetServiceEdgeGroupUsingGETParamsWithHTTPClient(client *http.Client) *GetServiceEdgeGroupUsingGETParams {
	return &GetServiceEdgeGroupUsingGETParams{
		HTTPClient: client,
	}
}

/* GetServiceEdgeGroupUsingGETParams contains all the parameters to send to the API endpoint
   for the get service edge group using g e t operation.

   Typically these are written to a http.Request.
*/
type GetServiceEdgeGroupUsingGETParams struct {

	/* CustomerID.

	   The unique identifier of the ZPA tenant.
	*/
	CustomerID string

	/* ServiceEdgeGroupID.

	   The unique identifier of the Service Edge Group.
	*/
	ServiceEdgeGroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get service edge group using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetServiceEdgeGroupUsingGETParams) WithDefaults() *GetServiceEdgeGroupUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get service edge group using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetServiceEdgeGroupUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get service edge group using g e t params
func (o *GetServiceEdgeGroupUsingGETParams) WithTimeout(timeout time.Duration) *GetServiceEdgeGroupUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get service edge group using g e t params
func (o *GetServiceEdgeGroupUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get service edge group using g e t params
func (o *GetServiceEdgeGroupUsingGETParams) WithContext(ctx context.Context) *GetServiceEdgeGroupUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get service edge group using g e t params
func (o *GetServiceEdgeGroupUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get service edge group using g e t params
func (o *GetServiceEdgeGroupUsingGETParams) WithHTTPClient(client *http.Client) *GetServiceEdgeGroupUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get service edge group using g e t params
func (o *GetServiceEdgeGroupUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerID adds the customerID to the get service edge group using g e t params
func (o *GetServiceEdgeGroupUsingGETParams) WithCustomerID(customerID string) *GetServiceEdgeGroupUsingGETParams {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the get service edge group using g e t params
func (o *GetServiceEdgeGroupUsingGETParams) SetCustomerID(customerID string) {
	o.CustomerID = customerID
}

// WithServiceEdgeGroupID adds the serviceEdgeGroupID to the get service edge group using g e t params
func (o *GetServiceEdgeGroupUsingGETParams) WithServiceEdgeGroupID(serviceEdgeGroupID string) *GetServiceEdgeGroupUsingGETParams {
	o.SetServiceEdgeGroupID(serviceEdgeGroupID)
	return o
}

// SetServiceEdgeGroupID adds the serviceEdgeGroupId to the get service edge group using g e t params
func (o *GetServiceEdgeGroupUsingGETParams) SetServiceEdgeGroupID(serviceEdgeGroupID string) {
	o.ServiceEdgeGroupID = serviceEdgeGroupID
}

// WriteToRequest writes these params to a swagger request
func (o *GetServiceEdgeGroupUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerId
	if err := r.SetPathParam("customerId", o.CustomerID); err != nil {
		return err
	}

	// path param serviceEdgeGroupId
	if err := r.SetPathParam("serviceEdgeGroupId", o.ServiceEdgeGroupID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}