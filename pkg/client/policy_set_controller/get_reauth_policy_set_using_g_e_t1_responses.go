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

// GetReauthPolicySetUsingGET1Reader is a Reader for the GetReauthPolicySetUsingGET1 structure.
type GetReauthPolicySetUsingGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReauthPolicySetUsingGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReauthPolicySetUsingGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetReauthPolicySetUsingGET1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetReauthPolicySetUsingGET1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetReauthPolicySetUsingGET1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetReauthPolicySetUsingGET1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetReauthPolicySetUsingGET1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetReauthPolicySetUsingGET1OK creates a GetReauthPolicySetUsingGET1OK with default headers values
func NewGetReauthPolicySetUsingGET1OK() *GetReauthPolicySetUsingGET1OK {
	return &GetReauthPolicySetUsingGET1OK{}
}

/* GetReauthPolicySetUsingGET1OK describes a response with status code 200, with default header values.

OK
*/
type GetReauthPolicySetUsingGET1OK struct {
	Payload *models.PolicySet
}

func (o *GetReauthPolicySetUsingGET1OK) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/policySet/reauth][%d] getReauthPolicySetUsingGET1OK  %+v", 200, o.Payload)
}
func (o *GetReauthPolicySetUsingGET1OK) GetPayload() *models.PolicySet {
	return o.Payload
}

func (o *GetReauthPolicySetUsingGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PolicySet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReauthPolicySetUsingGET1BadRequest creates a GetReauthPolicySetUsingGET1BadRequest with default headers values
func NewGetReauthPolicySetUsingGET1BadRequest() *GetReauthPolicySetUsingGET1BadRequest {
	return &GetReauthPolicySetUsingGET1BadRequest{}
}

/* GetReauthPolicySetUsingGET1BadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type GetReauthPolicySetUsingGET1BadRequest struct {
}

func (o *GetReauthPolicySetUsingGET1BadRequest) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/policySet/reauth][%d] getReauthPolicySetUsingGET1BadRequest ", 400)
}

func (o *GetReauthPolicySetUsingGET1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetReauthPolicySetUsingGET1Unauthorized creates a GetReauthPolicySetUsingGET1Unauthorized with default headers values
func NewGetReauthPolicySetUsingGET1Unauthorized() *GetReauthPolicySetUsingGET1Unauthorized {
	return &GetReauthPolicySetUsingGET1Unauthorized{}
}

/* GetReauthPolicySetUsingGET1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetReauthPolicySetUsingGET1Unauthorized struct {
}

func (o *GetReauthPolicySetUsingGET1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/policySet/reauth][%d] getReauthPolicySetUsingGET1Unauthorized ", 401)
}

func (o *GetReauthPolicySetUsingGET1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetReauthPolicySetUsingGET1Forbidden creates a GetReauthPolicySetUsingGET1Forbidden with default headers values
func NewGetReauthPolicySetUsingGET1Forbidden() *GetReauthPolicySetUsingGET1Forbidden {
	return &GetReauthPolicySetUsingGET1Forbidden{}
}

/* GetReauthPolicySetUsingGET1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetReauthPolicySetUsingGET1Forbidden struct {
}

func (o *GetReauthPolicySetUsingGET1Forbidden) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/policySet/reauth][%d] getReauthPolicySetUsingGET1Forbidden ", 403)
}

func (o *GetReauthPolicySetUsingGET1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetReauthPolicySetUsingGET1NotFound creates a GetReauthPolicySetUsingGET1NotFound with default headers values
func NewGetReauthPolicySetUsingGET1NotFound() *GetReauthPolicySetUsingGET1NotFound {
	return &GetReauthPolicySetUsingGET1NotFound{}
}

/* GetReauthPolicySetUsingGET1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetReauthPolicySetUsingGET1NotFound struct {
}

func (o *GetReauthPolicySetUsingGET1NotFound) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/policySet/reauth][%d] getReauthPolicySetUsingGET1NotFound ", 404)
}

func (o *GetReauthPolicySetUsingGET1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetReauthPolicySetUsingGET1TooManyRequests creates a GetReauthPolicySetUsingGET1TooManyRequests with default headers values
func NewGetReauthPolicySetUsingGET1TooManyRequests() *GetReauthPolicySetUsingGET1TooManyRequests {
	return &GetReauthPolicySetUsingGET1TooManyRequests{}
}

/* GetReauthPolicySetUsingGET1TooManyRequests describes a response with status code 429, with default header values.

TooManyRequest
*/
type GetReauthPolicySetUsingGET1TooManyRequests struct {
}

func (o *GetReauthPolicySetUsingGET1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/policySet/reauth][%d] getReauthPolicySetUsingGET1TooManyRequests ", 429)
}

func (o *GetReauthPolicySetUsingGET1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
