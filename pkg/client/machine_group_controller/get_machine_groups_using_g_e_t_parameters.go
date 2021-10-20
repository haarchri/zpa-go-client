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
	"github.com/go-openapi/swag"
)

// NewGetMachineGroupsUsingGETParams creates a new GetMachineGroupsUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMachineGroupsUsingGETParams() *GetMachineGroupsUsingGETParams {
	return &GetMachineGroupsUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMachineGroupsUsingGETParamsWithTimeout creates a new GetMachineGroupsUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetMachineGroupsUsingGETParamsWithTimeout(timeout time.Duration) *GetMachineGroupsUsingGETParams {
	return &GetMachineGroupsUsingGETParams{
		timeout: timeout,
	}
}

// NewGetMachineGroupsUsingGETParamsWithContext creates a new GetMachineGroupsUsingGETParams object
// with the ability to set a context for a request.
func NewGetMachineGroupsUsingGETParamsWithContext(ctx context.Context) *GetMachineGroupsUsingGETParams {
	return &GetMachineGroupsUsingGETParams{
		Context: ctx,
	}
}

// NewGetMachineGroupsUsingGETParamsWithHTTPClient creates a new GetMachineGroupsUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMachineGroupsUsingGETParamsWithHTTPClient(client *http.Client) *GetMachineGroupsUsingGETParams {
	return &GetMachineGroupsUsingGETParams{
		HTTPClient: client,
	}
}

/* GetMachineGroupsUsingGETParams contains all the parameters to send to the API endpoint
   for the get machine groups using g e t operation.

   Typically these are written to a http.Request.
*/
type GetMachineGroupsUsingGETParams struct {

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

// WithDefaults hydrates default values in the get machine groups using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMachineGroupsUsingGETParams) WithDefaults() *GetMachineGroupsUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get machine groups using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMachineGroupsUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get machine groups using g e t params
func (o *GetMachineGroupsUsingGETParams) WithTimeout(timeout time.Duration) *GetMachineGroupsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get machine groups using g e t params
func (o *GetMachineGroupsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get machine groups using g e t params
func (o *GetMachineGroupsUsingGETParams) WithContext(ctx context.Context) *GetMachineGroupsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get machine groups using g e t params
func (o *GetMachineGroupsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get machine groups using g e t params
func (o *GetMachineGroupsUsingGETParams) WithHTTPClient(client *http.Client) *GetMachineGroupsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get machine groups using g e t params
func (o *GetMachineGroupsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerID adds the customerID to the get machine groups using g e t params
func (o *GetMachineGroupsUsingGETParams) WithCustomerID(customerID int64) *GetMachineGroupsUsingGETParams {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the get machine groups using g e t params
func (o *GetMachineGroupsUsingGETParams) SetCustomerID(customerID int64) {
	o.CustomerID = customerID
}

// WithPage adds the page to the get machine groups using g e t params
func (o *GetMachineGroupsUsingGETParams) WithPage(page int32) *GetMachineGroupsUsingGETParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get machine groups using g e t params
func (o *GetMachineGroupsUsingGETParams) SetPage(page int32) {
	o.Page = page
}

// WithPagesize adds the pagesize to the get machine groups using g e t params
func (o *GetMachineGroupsUsingGETParams) WithPagesize(pagesize int32) *GetMachineGroupsUsingGETParams {
	o.SetPagesize(pagesize)
	return o
}

// SetPagesize adds the pagesize to the get machine groups using g e t params
func (o *GetMachineGroupsUsingGETParams) SetPagesize(pagesize int32) {
	o.Pagesize = pagesize
}

// WithSearch adds the search to the get machine groups using g e t params
func (o *GetMachineGroupsUsingGETParams) WithSearch(search string) *GetMachineGroupsUsingGETParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the get machine groups using g e t params
func (o *GetMachineGroupsUsingGETParams) SetSearch(search string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *GetMachineGroupsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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