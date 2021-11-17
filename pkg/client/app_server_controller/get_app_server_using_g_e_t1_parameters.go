// Code generated by go-swagger; DO NOT EDIT.

package app_server_controller

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

// NewGetAppServerUsingGET1Params creates a new GetAppServerUsingGET1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAppServerUsingGET1Params() *GetAppServerUsingGET1Params {
	return &GetAppServerUsingGET1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAppServerUsingGET1ParamsWithTimeout creates a new GetAppServerUsingGET1Params object
// with the ability to set a timeout on a request.
func NewGetAppServerUsingGET1ParamsWithTimeout(timeout time.Duration) *GetAppServerUsingGET1Params {
	return &GetAppServerUsingGET1Params{
		timeout: timeout,
	}
}

// NewGetAppServerUsingGET1ParamsWithContext creates a new GetAppServerUsingGET1Params object
// with the ability to set a context for a request.
func NewGetAppServerUsingGET1ParamsWithContext(ctx context.Context) *GetAppServerUsingGET1Params {
	return &GetAppServerUsingGET1Params{
		Context: ctx,
	}
}

// NewGetAppServerUsingGET1ParamsWithHTTPClient creates a new GetAppServerUsingGET1Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetAppServerUsingGET1ParamsWithHTTPClient(client *http.Client) *GetAppServerUsingGET1Params {
	return &GetAppServerUsingGET1Params{
		HTTPClient: client,
	}
}

/* GetAppServerUsingGET1Params contains all the parameters to send to the API endpoint
   for the get app server using g e t 1 operation.

   Typically these are written to a http.Request.
*/
type GetAppServerUsingGET1Params struct {

	/* CustomerID.

	   The unique identifier of the ZPA tenant.
	*/
	CustomerID string

	/* ServerID.

	   The unique identifier of the Server.
	*/
	ServerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get app server using g e t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAppServerUsingGET1Params) WithDefaults() *GetAppServerUsingGET1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get app server using g e t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAppServerUsingGET1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get app server using g e t 1 params
func (o *GetAppServerUsingGET1Params) WithTimeout(timeout time.Duration) *GetAppServerUsingGET1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get app server using g e t 1 params
func (o *GetAppServerUsingGET1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get app server using g e t 1 params
func (o *GetAppServerUsingGET1Params) WithContext(ctx context.Context) *GetAppServerUsingGET1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get app server using g e t 1 params
func (o *GetAppServerUsingGET1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get app server using g e t 1 params
func (o *GetAppServerUsingGET1Params) WithHTTPClient(client *http.Client) *GetAppServerUsingGET1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get app server using g e t 1 params
func (o *GetAppServerUsingGET1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerID adds the customerID to the get app server using g e t 1 params
func (o *GetAppServerUsingGET1Params) WithCustomerID(customerID string) *GetAppServerUsingGET1Params {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the get app server using g e t 1 params
func (o *GetAppServerUsingGET1Params) SetCustomerID(customerID string) {
	o.CustomerID = customerID
}

// WithServerID adds the serverID to the get app server using g e t 1 params
func (o *GetAppServerUsingGET1Params) WithServerID(serverID string) *GetAppServerUsingGET1Params {
	o.SetServerID(serverID)
	return o
}

// SetServerID adds the serverId to the get app server using g e t 1 params
func (o *GetAppServerUsingGET1Params) SetServerID(serverID string) {
	o.ServerID = serverID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAppServerUsingGET1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerId
	if err := r.SetPathParam("customerId", o.CustomerID); err != nil {
		return err
	}

	// path param serverId
	if err := r.SetPathParam("serverId", o.ServerID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
