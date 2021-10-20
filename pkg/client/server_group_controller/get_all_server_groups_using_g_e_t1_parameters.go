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
)

// NewGetAllServerGroupsUsingGET1Params creates a new GetAllServerGroupsUsingGET1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllServerGroupsUsingGET1Params() *GetAllServerGroupsUsingGET1Params {
	return &GetAllServerGroupsUsingGET1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllServerGroupsUsingGET1ParamsWithTimeout creates a new GetAllServerGroupsUsingGET1Params object
// with the ability to set a timeout on a request.
func NewGetAllServerGroupsUsingGET1ParamsWithTimeout(timeout time.Duration) *GetAllServerGroupsUsingGET1Params {
	return &GetAllServerGroupsUsingGET1Params{
		timeout: timeout,
	}
}

// NewGetAllServerGroupsUsingGET1ParamsWithContext creates a new GetAllServerGroupsUsingGET1Params object
// with the ability to set a context for a request.
func NewGetAllServerGroupsUsingGET1ParamsWithContext(ctx context.Context) *GetAllServerGroupsUsingGET1Params {
	return &GetAllServerGroupsUsingGET1Params{
		Context: ctx,
	}
}

// NewGetAllServerGroupsUsingGET1ParamsWithHTTPClient creates a new GetAllServerGroupsUsingGET1Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllServerGroupsUsingGET1ParamsWithHTTPClient(client *http.Client) *GetAllServerGroupsUsingGET1Params {
	return &GetAllServerGroupsUsingGET1Params{
		HTTPClient: client,
	}
}

/* GetAllServerGroupsUsingGET1Params contains all the parameters to send to the API endpoint
   for the get all server groups using g e t 1 operation.

   Typically these are written to a http.Request.
*/
type GetAllServerGroupsUsingGET1Params struct {

	/* CustomerID.

	   The unique identifier of the ZPA tenant.

	   Format: int64
	*/
	CustomerID int64

	/* Page.

	   Specifies the page number.

	   Format: int32
	*/
	Page int32

	/* Pagesize.

	   Specifies the page size. If not provided, the default page size is 20. The max page size is 500.

	   Format: int32
	*/
	Pagesize int32

	/* Search.

	   The search string used to support search by features and fields for the API.
	*/
	Search string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get all server groups using g e t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllServerGroupsUsingGET1Params) WithDefaults() *GetAllServerGroupsUsingGET1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all server groups using g e t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllServerGroupsUsingGET1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get all server groups using g e t 1 params
func (o *GetAllServerGroupsUsingGET1Params) WithTimeout(timeout time.Duration) *GetAllServerGroupsUsingGET1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all server groups using g e t 1 params
func (o *GetAllServerGroupsUsingGET1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all server groups using g e t 1 params
func (o *GetAllServerGroupsUsingGET1Params) WithContext(ctx context.Context) *GetAllServerGroupsUsingGET1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all server groups using g e t 1 params
func (o *GetAllServerGroupsUsingGET1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all server groups using g e t 1 params
func (o *GetAllServerGroupsUsingGET1Params) WithHTTPClient(client *http.Client) *GetAllServerGroupsUsingGET1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all server groups using g e t 1 params
func (o *GetAllServerGroupsUsingGET1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerID adds the customerID to the get all server groups using g e t 1 params
func (o *GetAllServerGroupsUsingGET1Params) WithCustomerID(customerID int64) *GetAllServerGroupsUsingGET1Params {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the get all server groups using g e t 1 params
func (o *GetAllServerGroupsUsingGET1Params) SetCustomerID(customerID int64) {
	o.CustomerID = customerID
}

// WithPage adds the page to the get all server groups using g e t 1 params
func (o *GetAllServerGroupsUsingGET1Params) WithPage(page int32) *GetAllServerGroupsUsingGET1Params {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get all server groups using g e t 1 params
func (o *GetAllServerGroupsUsingGET1Params) SetPage(page int32) {
	o.Page = page
}

// WithPagesize adds the pagesize to the get all server groups using g e t 1 params
func (o *GetAllServerGroupsUsingGET1Params) WithPagesize(pagesize int32) *GetAllServerGroupsUsingGET1Params {
	o.SetPagesize(pagesize)
	return o
}

// SetPagesize adds the pagesize to the get all server groups using g e t 1 params
func (o *GetAllServerGroupsUsingGET1Params) SetPagesize(pagesize int32) {
	o.Pagesize = pagesize
}

// WithSearch adds the search to the get all server groups using g e t 1 params
func (o *GetAllServerGroupsUsingGET1Params) WithSearch(search string) *GetAllServerGroupsUsingGET1Params {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the get all server groups using g e t 1 params
func (o *GetAllServerGroupsUsingGET1Params) SetSearch(search string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllServerGroupsUsingGET1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerId
	if err := r.SetPathParam("customerId", swag.FormatInt64(o.CustomerID)); err != nil {
		return err
	}

	// query param page
	qrPage := o.Page
	qPage := swag.FormatInt32(qrPage)

	if err := r.SetQueryParam("page", qPage); err != nil {
		return err
	}

	// query param pagesize
	qrPagesize := o.Pagesize
	qPagesize := swag.FormatInt32(qrPagesize)

	if err := r.SetQueryParam("pagesize", qPagesize); err != nil {
		return err
	}

	// query param search
	qrSearch := o.Search
	qSearch := qrSearch

	if err := r.SetQueryParam("search", qSearch); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}