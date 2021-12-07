// Code generated by go-swagger; DO NOT EDIT.

package scim_attribute_header_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haarchri/zpa-go-client/pkg/models"
)

// GetAllSCIMAttributesUsingGET1Reader is a Reader for the GetAllSCIMAttributesUsingGET1 structure.
type GetAllSCIMAttributesUsingGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllSCIMAttributesUsingGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllSCIMAttributesUsingGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAllSCIMAttributesUsingGET1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAllSCIMAttributesUsingGET1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAllSCIMAttributesUsingGET1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAllSCIMAttributesUsingGET1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAllSCIMAttributesUsingGET1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllSCIMAttributesUsingGET1OK creates a GetAllSCIMAttributesUsingGET1OK with default headers values
func NewGetAllSCIMAttributesUsingGET1OK() *GetAllSCIMAttributesUsingGET1OK {
	return &GetAllSCIMAttributesUsingGET1OK{}
}

/* GetAllSCIMAttributesUsingGET1OK describes a response with status code 200, with default header values.

OK
*/
type GetAllSCIMAttributesUsingGET1OK struct {
	Payload *models.PageListOfSCIMAttributeHeader
}

func (o *GetAllSCIMAttributesUsingGET1OK) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/idp/{idpId}/scimattribute][%d] getAllSCIMAttributesUsingGET1OK  %+v", 200, o.Payload)
}
func (o *GetAllSCIMAttributesUsingGET1OK) GetPayload() *models.PageListOfSCIMAttributeHeader {
	return o.Payload
}

func (o *GetAllSCIMAttributesUsingGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageListOfSCIMAttributeHeader)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllSCIMAttributesUsingGET1BadRequest creates a GetAllSCIMAttributesUsingGET1BadRequest with default headers values
func NewGetAllSCIMAttributesUsingGET1BadRequest() *GetAllSCIMAttributesUsingGET1BadRequest {
	return &GetAllSCIMAttributesUsingGET1BadRequest{}
}

/* GetAllSCIMAttributesUsingGET1BadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type GetAllSCIMAttributesUsingGET1BadRequest struct {
}

func (o *GetAllSCIMAttributesUsingGET1BadRequest) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/idp/{idpId}/scimattribute][%d] getAllSCIMAttributesUsingGET1BadRequest ", 400)
}

func (o *GetAllSCIMAttributesUsingGET1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllSCIMAttributesUsingGET1Unauthorized creates a GetAllSCIMAttributesUsingGET1Unauthorized with default headers values
func NewGetAllSCIMAttributesUsingGET1Unauthorized() *GetAllSCIMAttributesUsingGET1Unauthorized {
	return &GetAllSCIMAttributesUsingGET1Unauthorized{}
}

/* GetAllSCIMAttributesUsingGET1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAllSCIMAttributesUsingGET1Unauthorized struct {
}

func (o *GetAllSCIMAttributesUsingGET1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/idp/{idpId}/scimattribute][%d] getAllSCIMAttributesUsingGET1Unauthorized ", 401)
}

func (o *GetAllSCIMAttributesUsingGET1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllSCIMAttributesUsingGET1Forbidden creates a GetAllSCIMAttributesUsingGET1Forbidden with default headers values
func NewGetAllSCIMAttributesUsingGET1Forbidden() *GetAllSCIMAttributesUsingGET1Forbidden {
	return &GetAllSCIMAttributesUsingGET1Forbidden{}
}

/* GetAllSCIMAttributesUsingGET1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAllSCIMAttributesUsingGET1Forbidden struct {
}

func (o *GetAllSCIMAttributesUsingGET1Forbidden) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/idp/{idpId}/scimattribute][%d] getAllSCIMAttributesUsingGET1Forbidden ", 403)
}

func (o *GetAllSCIMAttributesUsingGET1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllSCIMAttributesUsingGET1NotFound creates a GetAllSCIMAttributesUsingGET1NotFound with default headers values
func NewGetAllSCIMAttributesUsingGET1NotFound() *GetAllSCIMAttributesUsingGET1NotFound {
	return &GetAllSCIMAttributesUsingGET1NotFound{}
}

/* GetAllSCIMAttributesUsingGET1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAllSCIMAttributesUsingGET1NotFound struct {
}

func (o *GetAllSCIMAttributesUsingGET1NotFound) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/idp/{idpId}/scimattribute][%d] getAllSCIMAttributesUsingGET1NotFound ", 404)
}

func (o *GetAllSCIMAttributesUsingGET1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllSCIMAttributesUsingGET1TooManyRequests creates a GetAllSCIMAttributesUsingGET1TooManyRequests with default headers values
func NewGetAllSCIMAttributesUsingGET1TooManyRequests() *GetAllSCIMAttributesUsingGET1TooManyRequests {
	return &GetAllSCIMAttributesUsingGET1TooManyRequests{}
}

/* GetAllSCIMAttributesUsingGET1TooManyRequests describes a response with status code 429, with default header values.

TooManyRequest
*/
type GetAllSCIMAttributesUsingGET1TooManyRequests struct {
}

func (o *GetAllSCIMAttributesUsingGET1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/idp/{idpId}/scimattribute][%d] getAllSCIMAttributesUsingGET1TooManyRequests ", 429)
}

func (o *GetAllSCIMAttributesUsingGET1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
