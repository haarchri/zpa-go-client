// Code generated by go-swagger; DO NOT EDIT.

package service_edge_controller

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

	"github.com/haarchri/zpa-go-client/pkg/models"
)

// NewUpdateServiceEdgeUsingPUTParams creates a new UpdateServiceEdgeUsingPUTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateServiceEdgeUsingPUTParams() *UpdateServiceEdgeUsingPUTParams {
	return &UpdateServiceEdgeUsingPUTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateServiceEdgeUsingPUTParamsWithTimeout creates a new UpdateServiceEdgeUsingPUTParams object
// with the ability to set a timeout on a request.
func NewUpdateServiceEdgeUsingPUTParamsWithTimeout(timeout time.Duration) *UpdateServiceEdgeUsingPUTParams {
	return &UpdateServiceEdgeUsingPUTParams{
		timeout: timeout,
	}
}

// NewUpdateServiceEdgeUsingPUTParamsWithContext creates a new UpdateServiceEdgeUsingPUTParams object
// with the ability to set a context for a request.
func NewUpdateServiceEdgeUsingPUTParamsWithContext(ctx context.Context) *UpdateServiceEdgeUsingPUTParams {
	return &UpdateServiceEdgeUsingPUTParams{
		Context: ctx,
	}
}

// NewUpdateServiceEdgeUsingPUTParamsWithHTTPClient creates a new UpdateServiceEdgeUsingPUTParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateServiceEdgeUsingPUTParamsWithHTTPClient(client *http.Client) *UpdateServiceEdgeUsingPUTParams {
	return &UpdateServiceEdgeUsingPUTParams{
		HTTPClient: client,
	}
}

/* UpdateServiceEdgeUsingPUTParams contains all the parameters to send to the API endpoint
   for the update service edge using p u t operation.

   Typically these are written to a http.Request.
*/
type UpdateServiceEdgeUsingPUTParams struct {

	/* CustomerID.

	   The unique identifier of the ZPA tenant.
	*/
	CustomerID string

	/* ServiceEdge.

	   The object of the Service Edge.
	*/
	ServiceEdge *models.ServiceEdge

	/* ServiceEdgeID.

	   The unique identifier of the Service Edge.
	*/
	ServiceEdgeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update service edge using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateServiceEdgeUsingPUTParams) WithDefaults() *UpdateServiceEdgeUsingPUTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update service edge using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateServiceEdgeUsingPUTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update service edge using p u t params
func (o *UpdateServiceEdgeUsingPUTParams) WithTimeout(timeout time.Duration) *UpdateServiceEdgeUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update service edge using p u t params
func (o *UpdateServiceEdgeUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update service edge using p u t params
func (o *UpdateServiceEdgeUsingPUTParams) WithContext(ctx context.Context) *UpdateServiceEdgeUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update service edge using p u t params
func (o *UpdateServiceEdgeUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update service edge using p u t params
func (o *UpdateServiceEdgeUsingPUTParams) WithHTTPClient(client *http.Client) *UpdateServiceEdgeUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update service edge using p u t params
func (o *UpdateServiceEdgeUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerID adds the customerID to the update service edge using p u t params
func (o *UpdateServiceEdgeUsingPUTParams) WithCustomerID(customerID string) *UpdateServiceEdgeUsingPUTParams {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the update service edge using p u t params
func (o *UpdateServiceEdgeUsingPUTParams) SetCustomerID(customerID string) {
	o.CustomerID = customerID
}

// WithServiceEdge adds the serviceEdge to the update service edge using p u t params
func (o *UpdateServiceEdgeUsingPUTParams) WithServiceEdge(serviceEdge *models.ServiceEdge) *UpdateServiceEdgeUsingPUTParams {
	o.SetServiceEdge(serviceEdge)
	return o
}

// SetServiceEdge adds the serviceEdge to the update service edge using p u t params
func (o *UpdateServiceEdgeUsingPUTParams) SetServiceEdge(serviceEdge *models.ServiceEdge) {
	o.ServiceEdge = serviceEdge
}

// WithServiceEdgeID adds the serviceEdgeID to the update service edge using p u t params
func (o *UpdateServiceEdgeUsingPUTParams) WithServiceEdgeID(serviceEdgeID string) *UpdateServiceEdgeUsingPUTParams {
	o.SetServiceEdgeID(serviceEdgeID)
	return o
}

// SetServiceEdgeID adds the serviceEdgeId to the update service edge using p u t params
func (o *UpdateServiceEdgeUsingPUTParams) SetServiceEdgeID(serviceEdgeID string) {
	o.ServiceEdgeID = serviceEdgeID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateServiceEdgeUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerId
	if err := r.SetPathParam("customerId", o.CustomerID); err != nil {
		return err
	}
	if o.ServiceEdge != nil {
		if err := r.SetBodyParam(o.ServiceEdge); err != nil {
			return err
		}
	}

	// path param serviceEdgeId
	if err := r.SetPathParam("serviceEdgeId", o.ServiceEdgeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}