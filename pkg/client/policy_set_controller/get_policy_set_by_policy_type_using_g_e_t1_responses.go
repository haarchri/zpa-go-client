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

// GetPolicySetByPolicyTypeUsingGET1Reader is a Reader for the GetPolicySetByPolicyTypeUsingGET1 structure.
type GetPolicySetByPolicyTypeUsingGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPolicySetByPolicyTypeUsingGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPolicySetByPolicyTypeUsingGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPolicySetByPolicyTypeUsingGET1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetPolicySetByPolicyTypeUsingGET1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPolicySetByPolicyTypeUsingGET1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPolicySetByPolicyTypeUsingGET1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetPolicySetByPolicyTypeUsingGET1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPolicySetByPolicyTypeUsingGET1OK creates a GetPolicySetByPolicyTypeUsingGET1OK with default headers values
func NewGetPolicySetByPolicyTypeUsingGET1OK() *GetPolicySetByPolicyTypeUsingGET1OK {
	return &GetPolicySetByPolicyTypeUsingGET1OK{}
}

/* GetPolicySetByPolicyTypeUsingGET1OK describes a response with status code 200, with default header values.

OK
*/
type GetPolicySetByPolicyTypeUsingGET1OK struct {
	Payload *models.PolicySet
}

func (o *GetPolicySetByPolicyTypeUsingGET1OK) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/policySet/policyType/{policyType}][%d] getPolicySetByPolicyTypeUsingGET1OK  %+v", 200, o.Payload)
}
func (o *GetPolicySetByPolicyTypeUsingGET1OK) GetPayload() *models.PolicySet {
	return o.Payload
}

func (o *GetPolicySetByPolicyTypeUsingGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PolicySet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPolicySetByPolicyTypeUsingGET1BadRequest creates a GetPolicySetByPolicyTypeUsingGET1BadRequest with default headers values
func NewGetPolicySetByPolicyTypeUsingGET1BadRequest() *GetPolicySetByPolicyTypeUsingGET1BadRequest {
	return &GetPolicySetByPolicyTypeUsingGET1BadRequest{}
}

/* GetPolicySetByPolicyTypeUsingGET1BadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type GetPolicySetByPolicyTypeUsingGET1BadRequest struct {
}

func (o *GetPolicySetByPolicyTypeUsingGET1BadRequest) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/policySet/policyType/{policyType}][%d] getPolicySetByPolicyTypeUsingGET1BadRequest ", 400)
}

func (o *GetPolicySetByPolicyTypeUsingGET1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPolicySetByPolicyTypeUsingGET1Unauthorized creates a GetPolicySetByPolicyTypeUsingGET1Unauthorized with default headers values
func NewGetPolicySetByPolicyTypeUsingGET1Unauthorized() *GetPolicySetByPolicyTypeUsingGET1Unauthorized {
	return &GetPolicySetByPolicyTypeUsingGET1Unauthorized{}
}

/* GetPolicySetByPolicyTypeUsingGET1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetPolicySetByPolicyTypeUsingGET1Unauthorized struct {
}

func (o *GetPolicySetByPolicyTypeUsingGET1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/policySet/policyType/{policyType}][%d] getPolicySetByPolicyTypeUsingGET1Unauthorized ", 401)
}

func (o *GetPolicySetByPolicyTypeUsingGET1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPolicySetByPolicyTypeUsingGET1Forbidden creates a GetPolicySetByPolicyTypeUsingGET1Forbidden with default headers values
func NewGetPolicySetByPolicyTypeUsingGET1Forbidden() *GetPolicySetByPolicyTypeUsingGET1Forbidden {
	return &GetPolicySetByPolicyTypeUsingGET1Forbidden{}
}

/* GetPolicySetByPolicyTypeUsingGET1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetPolicySetByPolicyTypeUsingGET1Forbidden struct {
}

func (o *GetPolicySetByPolicyTypeUsingGET1Forbidden) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/policySet/policyType/{policyType}][%d] getPolicySetByPolicyTypeUsingGET1Forbidden ", 403)
}

func (o *GetPolicySetByPolicyTypeUsingGET1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPolicySetByPolicyTypeUsingGET1NotFound creates a GetPolicySetByPolicyTypeUsingGET1NotFound with default headers values
func NewGetPolicySetByPolicyTypeUsingGET1NotFound() *GetPolicySetByPolicyTypeUsingGET1NotFound {
	return &GetPolicySetByPolicyTypeUsingGET1NotFound{}
}

/* GetPolicySetByPolicyTypeUsingGET1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetPolicySetByPolicyTypeUsingGET1NotFound struct {
}

func (o *GetPolicySetByPolicyTypeUsingGET1NotFound) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/policySet/policyType/{policyType}][%d] getPolicySetByPolicyTypeUsingGET1NotFound ", 404)
}

func (o *GetPolicySetByPolicyTypeUsingGET1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPolicySetByPolicyTypeUsingGET1TooManyRequests creates a GetPolicySetByPolicyTypeUsingGET1TooManyRequests with default headers values
func NewGetPolicySetByPolicyTypeUsingGET1TooManyRequests() *GetPolicySetByPolicyTypeUsingGET1TooManyRequests {
	return &GetPolicySetByPolicyTypeUsingGET1TooManyRequests{}
}

/* GetPolicySetByPolicyTypeUsingGET1TooManyRequests describes a response with status code 429, with default header values.

TooManyRequest
*/
type GetPolicySetByPolicyTypeUsingGET1TooManyRequests struct {
}

func (o *GetPolicySetByPolicyTypeUsingGET1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/policySet/policyType/{policyType}][%d] getPolicySetByPolicyTypeUsingGET1TooManyRequests ", 429)
}

func (o *GetPolicySetByPolicyTypeUsingGET1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}