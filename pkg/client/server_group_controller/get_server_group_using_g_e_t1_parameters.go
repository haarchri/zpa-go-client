// Code generated by go-swagger; DO NOT EDIT.

package server_group_controller

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

// NewGetServerGroupUsingGET1Params creates a new GetServerGroupUsingGET1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetServerGroupUsingGET1Params() *GetServerGroupUsingGET1Params {
	return &GetServerGroupUsingGET1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetServerGroupUsingGET1ParamsWithTimeout creates a new GetServerGroupUsingGET1Params object
// with the ability to set a timeout on a request.
func NewGetServerGroupUsingGET1ParamsWithTimeout(timeout time.Duration) *GetServerGroupUsingGET1Params {
	return &GetServerGroupUsingGET1Params{
		timeout: timeout,
	}
}

// NewGetServerGroupUsingGET1ParamsWithContext creates a new GetServerGroupUsingGET1Params object
// with the ability to set a context for a request.
func NewGetServerGroupUsingGET1ParamsWithContext(ctx context.Context) *GetServerGroupUsingGET1Params {
	return &GetServerGroupUsingGET1Params{
		Context: ctx,
	}
}

// NewGetServerGroupUsingGET1ParamsWithHTTPClient creates a new GetServerGroupUsingGET1Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetServerGroupUsingGET1ParamsWithHTTPClient(client *http.Client) *GetServerGroupUsingGET1Params {
	return &GetServerGroupUsingGET1Params{
		HTTPClient: client,
	}
}

/* GetServerGroupUsingGET1Params contains all the parameters to send to the API endpoint
   for the get server group using g e t 1 operation.

   Typically these are written to a http.Request.
*/
type GetServerGroupUsingGET1Params struct {

	/* CustomerID.

	   The unique identifier of the ZPA tenant.
	*/
	CustomerID string

	/* GroupID.

	   The unique identifier of the Segment Group.
	*/
	GroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get server group using g e t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetServerGroupUsingGET1Params) WithDefaults() *GetServerGroupUsingGET1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get server group using g e t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetServerGroupUsingGET1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get server group using g e t 1 params
func (o *GetServerGroupUsingGET1Params) WithTimeout(timeout time.Duration) *GetServerGroupUsingGET1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get server group using g e t 1 params
func (o *GetServerGroupUsingGET1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get server group using g e t 1 params
func (o *GetServerGroupUsingGET1Params) WithContext(ctx context.Context) *GetServerGroupUsingGET1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get server group using g e t 1 params
func (o *GetServerGroupUsingGET1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get server group using g e t 1 params
func (o *GetServerGroupUsingGET1Params) WithHTTPClient(client *http.Client) *GetServerGroupUsingGET1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get server group using g e t 1 params
func (o *GetServerGroupUsingGET1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerID adds the customerID to the get server group using g e t 1 params
func (o *GetServerGroupUsingGET1Params) WithCustomerID(customerID string) *GetServerGroupUsingGET1Params {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the get server group using g e t 1 params
func (o *GetServerGroupUsingGET1Params) SetCustomerID(customerID string) {
	o.CustomerID = customerID
}

// WithGroupID adds the groupID to the get server group using g e t 1 params
func (o *GetServerGroupUsingGET1Params) WithGroupID(groupID string) *GetServerGroupUsingGET1Params {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the get server group using g e t 1 params
func (o *GetServerGroupUsingGET1Params) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WriteToRequest writes these params to a swagger request
func (o *GetServerGroupUsingGET1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerId
	if err := r.SetPathParam("customerId", o.CustomerID); err != nil {
		return err
	}

	// path param groupId
	if err := r.SetPathParam("groupId", o.GroupID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
