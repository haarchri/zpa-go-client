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
	"github.com/go-openapi/swag"

	"github.com/haarchri/zpa-go-client/pkg/models"
)

// NewAddAppServerGroupUsingPOST1Params creates a new AddAppServerGroupUsingPOST1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddAppServerGroupUsingPOST1Params() *AddAppServerGroupUsingPOST1Params {
	return &AddAppServerGroupUsingPOST1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddAppServerGroupUsingPOST1ParamsWithTimeout creates a new AddAppServerGroupUsingPOST1Params object
// with the ability to set a timeout on a request.
func NewAddAppServerGroupUsingPOST1ParamsWithTimeout(timeout time.Duration) *AddAppServerGroupUsingPOST1Params {
	return &AddAppServerGroupUsingPOST1Params{
		timeout: timeout,
	}
}

// NewAddAppServerGroupUsingPOST1ParamsWithContext creates a new AddAppServerGroupUsingPOST1Params object
// with the ability to set a context for a request.
func NewAddAppServerGroupUsingPOST1ParamsWithContext(ctx context.Context) *AddAppServerGroupUsingPOST1Params {
	return &AddAppServerGroupUsingPOST1Params{
		Context: ctx,
	}
}

// NewAddAppServerGroupUsingPOST1ParamsWithHTTPClient creates a new AddAppServerGroupUsingPOST1Params object
// with the ability to set a custom HTTPClient for a request.
func NewAddAppServerGroupUsingPOST1ParamsWithHTTPClient(client *http.Client) *AddAppServerGroupUsingPOST1Params {
	return &AddAppServerGroupUsingPOST1Params{
		HTTPClient: client,
	}
}

/* AddAppServerGroupUsingPOST1Params contains all the parameters to send to the API endpoint
   for the add app server group using p o s t 1 operation.

   Typically these are written to a http.Request.
*/
type AddAppServerGroupUsingPOST1Params struct {

	/* CustomerID.

	   The unique identifier of the ZPA tenant.

	   Format: int64
	*/
	CustomerID int64

	/* Group.

	   The object of the Server Group.
	*/
	Group *models.ServerGroupDTO

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add app server group using p o s t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddAppServerGroupUsingPOST1Params) WithDefaults() *AddAppServerGroupUsingPOST1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add app server group using p o s t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddAppServerGroupUsingPOST1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add app server group using p o s t 1 params
func (o *AddAppServerGroupUsingPOST1Params) WithTimeout(timeout time.Duration) *AddAppServerGroupUsingPOST1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add app server group using p o s t 1 params
func (o *AddAppServerGroupUsingPOST1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add app server group using p o s t 1 params
func (o *AddAppServerGroupUsingPOST1Params) WithContext(ctx context.Context) *AddAppServerGroupUsingPOST1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add app server group using p o s t 1 params
func (o *AddAppServerGroupUsingPOST1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add app server group using p o s t 1 params
func (o *AddAppServerGroupUsingPOST1Params) WithHTTPClient(client *http.Client) *AddAppServerGroupUsingPOST1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add app server group using p o s t 1 params
func (o *AddAppServerGroupUsingPOST1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerID adds the customerID to the add app server group using p o s t 1 params
func (o *AddAppServerGroupUsingPOST1Params) WithCustomerID(customerID int64) *AddAppServerGroupUsingPOST1Params {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the add app server group using p o s t 1 params
func (o *AddAppServerGroupUsingPOST1Params) SetCustomerID(customerID int64) {
	o.CustomerID = customerID
}

// WithGroup adds the group to the add app server group using p o s t 1 params
func (o *AddAppServerGroupUsingPOST1Params) WithGroup(group *models.ServerGroupDTO) *AddAppServerGroupUsingPOST1Params {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the add app server group using p o s t 1 params
func (o *AddAppServerGroupUsingPOST1Params) SetGroup(group *models.ServerGroupDTO) {
	o.Group = group
}

// WriteToRequest writes these params to a swagger request
func (o *AddAppServerGroupUsingPOST1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerId
	if err := r.SetPathParam("customerId", swag.FormatInt64(o.CustomerID)); err != nil {
		return err
	}
	if o.Group != nil {
		if err := r.SetBodyParam(o.Group); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
