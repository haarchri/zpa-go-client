// Code generated by go-swagger; DO NOT EDIT.

package scim_attribute_header_controller

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

// NewGetAllSCIMAttributesUsingGET1Params creates a new GetAllSCIMAttributesUsingGET1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllSCIMAttributesUsingGET1Params() *GetAllSCIMAttributesUsingGET1Params {
	return &GetAllSCIMAttributesUsingGET1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllSCIMAttributesUsingGET1ParamsWithTimeout creates a new GetAllSCIMAttributesUsingGET1Params object
// with the ability to set a timeout on a request.
func NewGetAllSCIMAttributesUsingGET1ParamsWithTimeout(timeout time.Duration) *GetAllSCIMAttributesUsingGET1Params {
	return &GetAllSCIMAttributesUsingGET1Params{
		timeout: timeout,
	}
}

// NewGetAllSCIMAttributesUsingGET1ParamsWithContext creates a new GetAllSCIMAttributesUsingGET1Params object
// with the ability to set a context for a request.
func NewGetAllSCIMAttributesUsingGET1ParamsWithContext(ctx context.Context) *GetAllSCIMAttributesUsingGET1Params {
	return &GetAllSCIMAttributesUsingGET1Params{
		Context: ctx,
	}
}

// NewGetAllSCIMAttributesUsingGET1ParamsWithHTTPClient creates a new GetAllSCIMAttributesUsingGET1Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllSCIMAttributesUsingGET1ParamsWithHTTPClient(client *http.Client) *GetAllSCIMAttributesUsingGET1Params {
	return &GetAllSCIMAttributesUsingGET1Params{
		HTTPClient: client,
	}
}

/* GetAllSCIMAttributesUsingGET1Params contains all the parameters to send to the API endpoint
   for the get all s c i m attributes using g e t 1 operation.

   Typically these are written to a http.Request.
*/
type GetAllSCIMAttributesUsingGET1Params struct {

	/* CustomerID.

	   The unique identifier of the ZPA tenant.

	   Format: int64
	*/
	CustomerID int64

	/* IdpID.

	   The unique identifier of the IdP.

	   Format: int64
	*/
	IdpID int64

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

// WithDefaults hydrates default values in the get all s c i m attributes using g e t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllSCIMAttributesUsingGET1Params) WithDefaults() *GetAllSCIMAttributesUsingGET1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all s c i m attributes using g e t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllSCIMAttributesUsingGET1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get all s c i m attributes using g e t 1 params
func (o *GetAllSCIMAttributesUsingGET1Params) WithTimeout(timeout time.Duration) *GetAllSCIMAttributesUsingGET1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all s c i m attributes using g e t 1 params
func (o *GetAllSCIMAttributesUsingGET1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all s c i m attributes using g e t 1 params
func (o *GetAllSCIMAttributesUsingGET1Params) WithContext(ctx context.Context) *GetAllSCIMAttributesUsingGET1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all s c i m attributes using g e t 1 params
func (o *GetAllSCIMAttributesUsingGET1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all s c i m attributes using g e t 1 params
func (o *GetAllSCIMAttributesUsingGET1Params) WithHTTPClient(client *http.Client) *GetAllSCIMAttributesUsingGET1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all s c i m attributes using g e t 1 params
func (o *GetAllSCIMAttributesUsingGET1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerID adds the customerID to the get all s c i m attributes using g e t 1 params
func (o *GetAllSCIMAttributesUsingGET1Params) WithCustomerID(customerID int64) *GetAllSCIMAttributesUsingGET1Params {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the get all s c i m attributes using g e t 1 params
func (o *GetAllSCIMAttributesUsingGET1Params) SetCustomerID(customerID int64) {
	o.CustomerID = customerID
}

// WithIdpID adds the idpID to the get all s c i m attributes using g e t 1 params
func (o *GetAllSCIMAttributesUsingGET1Params) WithIdpID(idpID int64) *GetAllSCIMAttributesUsingGET1Params {
	o.SetIdpID(idpID)
	return o
}

// SetIdpID adds the idpId to the get all s c i m attributes using g e t 1 params
func (o *GetAllSCIMAttributesUsingGET1Params) SetIdpID(idpID int64) {
	o.IdpID = idpID
}

// WithPage adds the page to the get all s c i m attributes using g e t 1 params
func (o *GetAllSCIMAttributesUsingGET1Params) WithPage(page int32) *GetAllSCIMAttributesUsingGET1Params {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get all s c i m attributes using g e t 1 params
func (o *GetAllSCIMAttributesUsingGET1Params) SetPage(page int32) {
	o.Page = page
}

// WithPagesize adds the pagesize to the get all s c i m attributes using g e t 1 params
func (o *GetAllSCIMAttributesUsingGET1Params) WithPagesize(pagesize int32) *GetAllSCIMAttributesUsingGET1Params {
	o.SetPagesize(pagesize)
	return o
}

// SetPagesize adds the pagesize to the get all s c i m attributes using g e t 1 params
func (o *GetAllSCIMAttributesUsingGET1Params) SetPagesize(pagesize int32) {
	o.Pagesize = pagesize
}

// WithSearch adds the search to the get all s c i m attributes using g e t 1 params
func (o *GetAllSCIMAttributesUsingGET1Params) WithSearch(search string) *GetAllSCIMAttributesUsingGET1Params {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the get all s c i m attributes using g e t 1 params
func (o *GetAllSCIMAttributesUsingGET1Params) SetSearch(search string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllSCIMAttributesUsingGET1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerId
	if err := r.SetPathParam("customerId", swag.FormatInt64(o.CustomerID)); err != nil {
		return err
	}

	// path param idpId
	if err := r.SetPathParam("idpId", swag.FormatInt64(o.IdpID)); err != nil {
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
