// Code generated by go-swagger; DO NOT EDIT.

package policy_set_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haarchri/zpa-go-client/pkg/models"
)

// GetPolicyRulesByPageUsingGET1Reader is a Reader for the GetPolicyRulesByPageUsingGET1 structure.
type GetPolicyRulesByPageUsingGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPolicyRulesByPageUsingGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPolicyRulesByPageUsingGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPolicyRulesByPageUsingGET1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetPolicyRulesByPageUsingGET1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPolicyRulesByPageUsingGET1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPolicyRulesByPageUsingGET1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetPolicyRulesByPageUsingGET1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPolicyRulesByPageUsingGET1OK creates a GetPolicyRulesByPageUsingGET1OK with default headers values
func NewGetPolicyRulesByPageUsingGET1OK() *GetPolicyRulesByPageUsingGET1OK {
	return &GetPolicyRulesByPageUsingGET1OK{}
}

/* GetPolicyRulesByPageUsingGET1OK describes a response with status code 200, with default header values.

OK
*/
type GetPolicyRulesByPageUsingGET1OK struct {
	Payload *models.PageListOfPolicyRule
}

func (o *GetPolicyRulesByPageUsingGET1OK) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/policySet/rules/policyType/{policyType}][%d] getPolicyRulesByPageUsingGET1OK  %+v", 200, o.Payload)
}
func (o *GetPolicyRulesByPageUsingGET1OK) GetPayload() *models.PageListOfPolicyRule {
	return o.Payload
}

func (o *GetPolicyRulesByPageUsingGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageListOfPolicyRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPolicyRulesByPageUsingGET1BadRequest creates a GetPolicyRulesByPageUsingGET1BadRequest with default headers values
func NewGetPolicyRulesByPageUsingGET1BadRequest() *GetPolicyRulesByPageUsingGET1BadRequest {
	return &GetPolicyRulesByPageUsingGET1BadRequest{}
}

/* GetPolicyRulesByPageUsingGET1BadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type GetPolicyRulesByPageUsingGET1BadRequest struct {
}

func (o *GetPolicyRulesByPageUsingGET1BadRequest) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/policySet/rules/policyType/{policyType}][%d] getPolicyRulesByPageUsingGET1BadRequest ", 400)
}

func (o *GetPolicyRulesByPageUsingGET1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPolicyRulesByPageUsingGET1Unauthorized creates a GetPolicyRulesByPageUsingGET1Unauthorized with default headers values
func NewGetPolicyRulesByPageUsingGET1Unauthorized() *GetPolicyRulesByPageUsingGET1Unauthorized {
	return &GetPolicyRulesByPageUsingGET1Unauthorized{}
}

/* GetPolicyRulesByPageUsingGET1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetPolicyRulesByPageUsingGET1Unauthorized struct {
}

func (o *GetPolicyRulesByPageUsingGET1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/policySet/rules/policyType/{policyType}][%d] getPolicyRulesByPageUsingGET1Unauthorized ", 401)
}

func (o *GetPolicyRulesByPageUsingGET1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPolicyRulesByPageUsingGET1Forbidden creates a GetPolicyRulesByPageUsingGET1Forbidden with default headers values
func NewGetPolicyRulesByPageUsingGET1Forbidden() *GetPolicyRulesByPageUsingGET1Forbidden {
	return &GetPolicyRulesByPageUsingGET1Forbidden{}
}

/* GetPolicyRulesByPageUsingGET1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetPolicyRulesByPageUsingGET1Forbidden struct {
}

func (o *GetPolicyRulesByPageUsingGET1Forbidden) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/policySet/rules/policyType/{policyType}][%d] getPolicyRulesByPageUsingGET1Forbidden ", 403)
}

func (o *GetPolicyRulesByPageUsingGET1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPolicyRulesByPageUsingGET1NotFound creates a GetPolicyRulesByPageUsingGET1NotFound with default headers values
func NewGetPolicyRulesByPageUsingGET1NotFound() *GetPolicyRulesByPageUsingGET1NotFound {
	return &GetPolicyRulesByPageUsingGET1NotFound{}
}

/* GetPolicyRulesByPageUsingGET1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetPolicyRulesByPageUsingGET1NotFound struct {
}

func (o *GetPolicyRulesByPageUsingGET1NotFound) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/policySet/rules/policyType/{policyType}][%d] getPolicyRulesByPageUsingGET1NotFound ", 404)
}

func (o *GetPolicyRulesByPageUsingGET1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPolicyRulesByPageUsingGET1TooManyRequests creates a GetPolicyRulesByPageUsingGET1TooManyRequests with default headers values
func NewGetPolicyRulesByPageUsingGET1TooManyRequests() *GetPolicyRulesByPageUsingGET1TooManyRequests {
	return &GetPolicyRulesByPageUsingGET1TooManyRequests{}
}

/* GetPolicyRulesByPageUsingGET1TooManyRequests describes a response with status code 429, with default header values.

TooManyRequest
*/
type GetPolicyRulesByPageUsingGET1TooManyRequests struct {
}

func (o *GetPolicyRulesByPageUsingGET1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/policySet/rules/policyType/{policyType}][%d] getPolicyRulesByPageUsingGET1TooManyRequests ", 429)
}

func (o *GetPolicyRulesByPageUsingGET1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
