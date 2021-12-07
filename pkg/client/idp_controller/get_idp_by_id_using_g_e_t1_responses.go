// Code generated by go-swagger; DO NOT EDIT.

package idp_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haarchri/zpa-go-client/pkg/models"
)

// GetIdpByIDUsingGET1Reader is a Reader for the GetIdpByIDUsingGET1 structure.
type GetIdpByIDUsingGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIdpByIDUsingGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIdpByIDUsingGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIdpByIDUsingGET1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIdpByIDUsingGET1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIdpByIDUsingGET1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIdpByIDUsingGET1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIdpByIDUsingGET1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIdpByIDUsingGET1OK creates a GetIdpByIDUsingGET1OK with default headers values
func NewGetIdpByIDUsingGET1OK() *GetIdpByIDUsingGET1OK {
	return &GetIdpByIDUsingGET1OK{}
}

/* GetIdpByIDUsingGET1OK describes a response with status code 200, with default header values.

OK
*/
type GetIdpByIDUsingGET1OK struct {
	Payload *models.Idp
}

func (o *GetIdpByIDUsingGET1OK) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/idp/{idpId}][%d] getIdpByIdUsingGET1OK  %+v", 200, o.Payload)
}
func (o *GetIdpByIDUsingGET1OK) GetPayload() *models.Idp {
	return o.Payload
}

func (o *GetIdpByIDUsingGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Idp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdpByIDUsingGET1BadRequest creates a GetIdpByIDUsingGET1BadRequest with default headers values
func NewGetIdpByIDUsingGET1BadRequest() *GetIdpByIDUsingGET1BadRequest {
	return &GetIdpByIDUsingGET1BadRequest{}
}

/* GetIdpByIDUsingGET1BadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type GetIdpByIDUsingGET1BadRequest struct {
}

func (o *GetIdpByIDUsingGET1BadRequest) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/idp/{idpId}][%d] getIdpByIdUsingGET1BadRequest ", 400)
}

func (o *GetIdpByIDUsingGET1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIdpByIDUsingGET1Unauthorized creates a GetIdpByIDUsingGET1Unauthorized with default headers values
func NewGetIdpByIDUsingGET1Unauthorized() *GetIdpByIDUsingGET1Unauthorized {
	return &GetIdpByIDUsingGET1Unauthorized{}
}

/* GetIdpByIDUsingGET1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetIdpByIDUsingGET1Unauthorized struct {
}

func (o *GetIdpByIDUsingGET1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/idp/{idpId}][%d] getIdpByIdUsingGET1Unauthorized ", 401)
}

func (o *GetIdpByIDUsingGET1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIdpByIDUsingGET1Forbidden creates a GetIdpByIDUsingGET1Forbidden with default headers values
func NewGetIdpByIDUsingGET1Forbidden() *GetIdpByIDUsingGET1Forbidden {
	return &GetIdpByIDUsingGET1Forbidden{}
}

/* GetIdpByIDUsingGET1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetIdpByIDUsingGET1Forbidden struct {
}

func (o *GetIdpByIDUsingGET1Forbidden) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/idp/{idpId}][%d] getIdpByIdUsingGET1Forbidden ", 403)
}

func (o *GetIdpByIDUsingGET1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIdpByIDUsingGET1NotFound creates a GetIdpByIDUsingGET1NotFound with default headers values
func NewGetIdpByIDUsingGET1NotFound() *GetIdpByIDUsingGET1NotFound {
	return &GetIdpByIDUsingGET1NotFound{}
}

/* GetIdpByIDUsingGET1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetIdpByIDUsingGET1NotFound struct {
}

func (o *GetIdpByIDUsingGET1NotFound) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/idp/{idpId}][%d] getIdpByIdUsingGET1NotFound ", 404)
}

func (o *GetIdpByIDUsingGET1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIdpByIDUsingGET1TooManyRequests creates a GetIdpByIDUsingGET1TooManyRequests with default headers values
func NewGetIdpByIDUsingGET1TooManyRequests() *GetIdpByIDUsingGET1TooManyRequests {
	return &GetIdpByIDUsingGET1TooManyRequests{}
}

/* GetIdpByIDUsingGET1TooManyRequests describes a response with status code 429, with default header values.

TooManyRequest
*/
type GetIdpByIDUsingGET1TooManyRequests struct {
}

func (o *GetIdpByIDUsingGET1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/idp/{idpId}][%d] getIdpByIdUsingGET1TooManyRequests ", 429)
}

func (o *GetIdpByIDUsingGET1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
