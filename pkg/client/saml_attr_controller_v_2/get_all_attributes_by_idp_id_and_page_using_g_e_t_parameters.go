// Code generated by go-swagger; DO NOT EDIT.

package saml_attr_controller_v_2

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

// NewGetAllAttributesByIdpIDAndPageUsingGETParams creates a new GetAllAttributesByIdpIDAndPageUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllAttributesByIdpIDAndPageUsingGETParams() *GetAllAttributesByIdpIDAndPageUsingGETParams {
	return &GetAllAttributesByIdpIDAndPageUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllAttributesByIdpIDAndPageUsingGETParamsWithTimeout creates a new GetAllAttributesByIdpIDAndPageUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetAllAttributesByIdpIDAndPageUsingGETParamsWithTimeout(timeout time.Duration) *GetAllAttributesByIdpIDAndPageUsingGETParams {
	return &GetAllAttributesByIdpIDAndPageUsingGETParams{
		timeout: timeout,
	}
}

// NewGetAllAttributesByIdpIDAndPageUsingGETParamsWithContext creates a new GetAllAttributesByIdpIDAndPageUsingGETParams object
// with the ability to set a context for a request.
func NewGetAllAttributesByIdpIDAndPageUsingGETParamsWithContext(ctx context.Context) *GetAllAttributesByIdpIDAndPageUsingGETParams {
	return &GetAllAttributesByIdpIDAndPageUsingGETParams{
		Context: ctx,
	}
}

// NewGetAllAttributesByIdpIDAndPageUsingGETParamsWithHTTPClient creates a new GetAllAttributesByIdpIDAndPageUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllAttributesByIdpIDAndPageUsingGETParamsWithHTTPClient(client *http.Client) *GetAllAttributesByIdpIDAndPageUsingGETParams {
	return &GetAllAttributesByIdpIDAndPageUsingGETParams{
		HTTPClient: client,
	}
}

/* GetAllAttributesByIdpIDAndPageUsingGETParams contains all the parameters to send to the API endpoint
   for the get all attributes by idp Id and page using g e t operation.

   Typically these are written to a http.Request.
*/
type GetAllAttributesByIdpIDAndPageUsingGETParams struct {

	/* CustomerID.

	   The unique identifier of the ZPA tenant.
	*/
	CustomerID string

	/* IdpID.

	   The unique identifier of the IdP.
	*/
	IdpID string

	/* Page.

	   Specifies the page number.
	*/
	Page string

	/* Pagesize.

	   Specifies the page size. If not provided, the default page size is 20. The max page size is 500.
	*/
	Pagesize string

	/* Search.

	   The search string used to support search by features and fields for the API.
	*/
	Search string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get all attributes by idp Id and page using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllAttributesByIdpIDAndPageUsingGETParams) WithDefaults() *GetAllAttributesByIdpIDAndPageUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all attributes by idp Id and page using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllAttributesByIdpIDAndPageUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get all attributes by idp Id and page using g e t params
func (o *GetAllAttributesByIdpIDAndPageUsingGETParams) WithTimeout(timeout time.Duration) *GetAllAttributesByIdpIDAndPageUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all attributes by idp Id and page using g e t params
func (o *GetAllAttributesByIdpIDAndPageUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all attributes by idp Id and page using g e t params
func (o *GetAllAttributesByIdpIDAndPageUsingGETParams) WithContext(ctx context.Context) *GetAllAttributesByIdpIDAndPageUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all attributes by idp Id and page using g e t params
func (o *GetAllAttributesByIdpIDAndPageUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all attributes by idp Id and page using g e t params
func (o *GetAllAttributesByIdpIDAndPageUsingGETParams) WithHTTPClient(client *http.Client) *GetAllAttributesByIdpIDAndPageUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all attributes by idp Id and page using g e t params
func (o *GetAllAttributesByIdpIDAndPageUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerID adds the customerID to the get all attributes by idp Id and page using g e t params
func (o *GetAllAttributesByIdpIDAndPageUsingGETParams) WithCustomerID(customerID string) *GetAllAttributesByIdpIDAndPageUsingGETParams {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the get all attributes by idp Id and page using g e t params
func (o *GetAllAttributesByIdpIDAndPageUsingGETParams) SetCustomerID(customerID string) {
	o.CustomerID = customerID
}

// WithIdpID adds the idpID to the get all attributes by idp Id and page using g e t params
func (o *GetAllAttributesByIdpIDAndPageUsingGETParams) WithIdpID(idpID string) *GetAllAttributesByIdpIDAndPageUsingGETParams {
	o.SetIdpID(idpID)
	return o
}

// SetIdpID adds the idpId to the get all attributes by idp Id and page using g e t params
func (o *GetAllAttributesByIdpIDAndPageUsingGETParams) SetIdpID(idpID string) {
	o.IdpID = idpID
}

// WithPage adds the page to the get all attributes by idp Id and page using g e t params
func (o *GetAllAttributesByIdpIDAndPageUsingGETParams) WithPage(page string) *GetAllAttributesByIdpIDAndPageUsingGETParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get all attributes by idp Id and page using g e t params
func (o *GetAllAttributesByIdpIDAndPageUsingGETParams) SetPage(page string) {
	o.Page = page
}

// WithPagesize adds the pagesize to the get all attributes by idp Id and page using g e t params
func (o *GetAllAttributesByIdpIDAndPageUsingGETParams) WithPagesize(pagesize string) *GetAllAttributesByIdpIDAndPageUsingGETParams {
	o.SetPagesize(pagesize)
	return o
}

// SetPagesize adds the pagesize to the get all attributes by idp Id and page using g e t params
func (o *GetAllAttributesByIdpIDAndPageUsingGETParams) SetPagesize(pagesize string) {
	o.Pagesize = pagesize
}

// WithSearch adds the search to the get all attributes by idp Id and page using g e t params
func (o *GetAllAttributesByIdpIDAndPageUsingGETParams) WithSearch(search string) *GetAllAttributesByIdpIDAndPageUsingGETParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the get all attributes by idp Id and page using g e t params
func (o *GetAllAttributesByIdpIDAndPageUsingGETParams) SetSearch(search string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllAttributesByIdpIDAndPageUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerId
	if err := r.SetPathParam("customerId", o.CustomerID); err != nil {
		return err
	}

	// path param idpId
	if err := r.SetPathParam("idpId", o.IdpID); err != nil {
		return err
	}

	// query param page
	qrPage := o.Page
	qPage := qrPage

	if err := r.SetQueryParam("page", qPage); err != nil {
		return err
	}

	// query param pagesize
	qrPagesize := o.Pagesize
	qPagesize := qrPagesize

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
