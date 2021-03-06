// Code generated by go-swagger; DO NOT EDIT.

package machine_group_controller

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

// NewGetMachineGroupUsingGETParams creates a new GetMachineGroupUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMachineGroupUsingGETParams() *GetMachineGroupUsingGETParams {
	return &GetMachineGroupUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMachineGroupUsingGETParamsWithTimeout creates a new GetMachineGroupUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetMachineGroupUsingGETParamsWithTimeout(timeout time.Duration) *GetMachineGroupUsingGETParams {
	return &GetMachineGroupUsingGETParams{
		timeout: timeout,
	}
}

// NewGetMachineGroupUsingGETParamsWithContext creates a new GetMachineGroupUsingGETParams object
// with the ability to set a context for a request.
func NewGetMachineGroupUsingGETParamsWithContext(ctx context.Context) *GetMachineGroupUsingGETParams {
	return &GetMachineGroupUsingGETParams{
		Context: ctx,
	}
}

// NewGetMachineGroupUsingGETParamsWithHTTPClient creates a new GetMachineGroupUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMachineGroupUsingGETParamsWithHTTPClient(client *http.Client) *GetMachineGroupUsingGETParams {
	return &GetMachineGroupUsingGETParams{
		HTTPClient: client,
	}
}

/* GetMachineGroupUsingGETParams contains all the parameters to send to the API endpoint
   for the get machine group using g e t operation.

   Typically these are written to a http.Request.
*/
type GetMachineGroupUsingGETParams struct {

	/* ID.

	   The unique identifier of the Machine Group.
	*/
	ID string

	/* CustomerID.

	   The unique identifier of the ZPA tenant.
	*/
	CustomerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get machine group using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMachineGroupUsingGETParams) WithDefaults() *GetMachineGroupUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get machine group using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMachineGroupUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get machine group using g e t params
func (o *GetMachineGroupUsingGETParams) WithTimeout(timeout time.Duration) *GetMachineGroupUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get machine group using g e t params
func (o *GetMachineGroupUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get machine group using g e t params
func (o *GetMachineGroupUsingGETParams) WithContext(ctx context.Context) *GetMachineGroupUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get machine group using g e t params
func (o *GetMachineGroupUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get machine group using g e t params
func (o *GetMachineGroupUsingGETParams) WithHTTPClient(client *http.Client) *GetMachineGroupUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get machine group using g e t params
func (o *GetMachineGroupUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get machine group using g e t params
func (o *GetMachineGroupUsingGETParams) WithID(id string) *GetMachineGroupUsingGETParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get machine group using g e t params
func (o *GetMachineGroupUsingGETParams) SetID(id string) {
	o.ID = id
}

// WithCustomerID adds the customerID to the get machine group using g e t params
func (o *GetMachineGroupUsingGETParams) WithCustomerID(customerID string) *GetMachineGroupUsingGETParams {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the get machine group using g e t params
func (o *GetMachineGroupUsingGETParams) SetCustomerID(customerID string) {
	o.CustomerID = customerID
}

// WriteToRequest writes these params to a swagger request
func (o *GetMachineGroupUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Id
	if err := r.SetPathParam("Id", o.ID); err != nil {
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
