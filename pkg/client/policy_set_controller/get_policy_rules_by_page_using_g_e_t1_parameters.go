// Code generated by go-swagger; DO NOT EDIT.

package policy_set_controller

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

// NewGetPolicyRulesByPageUsingGET1Params creates a new GetPolicyRulesByPageUsingGET1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPolicyRulesByPageUsingGET1Params() *GetPolicyRulesByPageUsingGET1Params {
	return &GetPolicyRulesByPageUsingGET1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPolicyRulesByPageUsingGET1ParamsWithTimeout creates a new GetPolicyRulesByPageUsingGET1Params object
// with the ability to set a timeout on a request.
func NewGetPolicyRulesByPageUsingGET1ParamsWithTimeout(timeout time.Duration) *GetPolicyRulesByPageUsingGET1Params {
	return &GetPolicyRulesByPageUsingGET1Params{
		timeout: timeout,
	}
}

// NewGetPolicyRulesByPageUsingGET1ParamsWithContext creates a new GetPolicyRulesByPageUsingGET1Params object
// with the ability to set a context for a request.
func NewGetPolicyRulesByPageUsingGET1ParamsWithContext(ctx context.Context) *GetPolicyRulesByPageUsingGET1Params {
	return &GetPolicyRulesByPageUsingGET1Params{
		Context: ctx,
	}
}

// NewGetPolicyRulesByPageUsingGET1ParamsWithHTTPClient creates a new GetPolicyRulesByPageUsingGET1Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetPolicyRulesByPageUsingGET1ParamsWithHTTPClient(client *http.Client) *GetPolicyRulesByPageUsingGET1Params {
	return &GetPolicyRulesByPageUsingGET1Params{
		HTTPClient: client,
	}
}

/* GetPolicyRulesByPageUsingGET1Params contains all the parameters to send to the API endpoint
   for the get policy rules by page using g e t 1 operation.

   Typically these are written to a http.Request.
*/
type GetPolicyRulesByPageUsingGET1Params struct {

	/* CustomerID.

	   The unique identifier of the ZPA tenant.
	*/
	CustomerID string

	/* Page.

	   Specifies the page number.
	*/
	Page string

	/* Pagesize.

	   Specifies the page size. If not provided, the default page size is 20. The max page size is 500.
	*/
	Pagesize string

	/* PolicyType.

	   Type for differentiating policy types. Supported values : ACCESS_POLICY/GLOBAL_POLICY, TIMEOUT_POLICY/REAUTH_POLICY, SIEM_POLICY, CLIENT_FORWARDING_POLICY/BYPASS_POLICY
	*/
	PolicyType string

	/* Search.

	   The search string used to support search by features and fields for the API.
	*/
	Search string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get policy rules by page using g e t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPolicyRulesByPageUsingGET1Params) WithDefaults() *GetPolicyRulesByPageUsingGET1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get policy rules by page using g e t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPolicyRulesByPageUsingGET1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get policy rules by page using g e t 1 params
func (o *GetPolicyRulesByPageUsingGET1Params) WithTimeout(timeout time.Duration) *GetPolicyRulesByPageUsingGET1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get policy rules by page using g e t 1 params
func (o *GetPolicyRulesByPageUsingGET1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get policy rules by page using g e t 1 params
func (o *GetPolicyRulesByPageUsingGET1Params) WithContext(ctx context.Context) *GetPolicyRulesByPageUsingGET1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get policy rules by page using g e t 1 params
func (o *GetPolicyRulesByPageUsingGET1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get policy rules by page using g e t 1 params
func (o *GetPolicyRulesByPageUsingGET1Params) WithHTTPClient(client *http.Client) *GetPolicyRulesByPageUsingGET1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get policy rules by page using g e t 1 params
func (o *GetPolicyRulesByPageUsingGET1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerID adds the customerID to the get policy rules by page using g e t 1 params
func (o *GetPolicyRulesByPageUsingGET1Params) WithCustomerID(customerID string) *GetPolicyRulesByPageUsingGET1Params {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the get policy rules by page using g e t 1 params
func (o *GetPolicyRulesByPageUsingGET1Params) SetCustomerID(customerID string) {
	o.CustomerID = customerID
}

// WithPage adds the page to the get policy rules by page using g e t 1 params
func (o *GetPolicyRulesByPageUsingGET1Params) WithPage(page string) *GetPolicyRulesByPageUsingGET1Params {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get policy rules by page using g e t 1 params
func (o *GetPolicyRulesByPageUsingGET1Params) SetPage(page string) {
	o.Page = page
}

// WithPagesize adds the pagesize to the get policy rules by page using g e t 1 params
func (o *GetPolicyRulesByPageUsingGET1Params) WithPagesize(pagesize string) *GetPolicyRulesByPageUsingGET1Params {
	o.SetPagesize(pagesize)
	return o
}

// SetPagesize adds the pagesize to the get policy rules by page using g e t 1 params
func (o *GetPolicyRulesByPageUsingGET1Params) SetPagesize(pagesize string) {
	o.Pagesize = pagesize
}

// WithPolicyType adds the policyType to the get policy rules by page using g e t 1 params
func (o *GetPolicyRulesByPageUsingGET1Params) WithPolicyType(policyType string) *GetPolicyRulesByPageUsingGET1Params {
	o.SetPolicyType(policyType)
	return o
}

// SetPolicyType adds the policyType to the get policy rules by page using g e t 1 params
func (o *GetPolicyRulesByPageUsingGET1Params) SetPolicyType(policyType string) {
	o.PolicyType = policyType
}

// WithSearch adds the search to the get policy rules by page using g e t 1 params
func (o *GetPolicyRulesByPageUsingGET1Params) WithSearch(search string) *GetPolicyRulesByPageUsingGET1Params {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the get policy rules by page using g e t 1 params
func (o *GetPolicyRulesByPageUsingGET1Params) SetSearch(search string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *GetPolicyRulesByPageUsingGET1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerId
	if err := r.SetPathParam("customerId", o.CustomerID); err != nil {
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

	// path param policyType
	if err := r.SetPathParam("policyType", o.PolicyType); err != nil {
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
