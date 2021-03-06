// Code generated by go-swagger; DO NOT EDIT.

package idp_controller_v_2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haarchri/zpa-go-client/pkg/models"
)

// GetAllIdpUsingGETReader is a Reader for the GetAllIdpUsingGET structure.
type GetAllIdpUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllIdpUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllIdpUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAllIdpUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAllIdpUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAllIdpUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAllIdpUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAllIdpUsingGETTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllIdpUsingGETOK creates a GetAllIdpUsingGETOK with default headers values
func NewGetAllIdpUsingGETOK() *GetAllIdpUsingGETOK {
	return &GetAllIdpUsingGETOK{}
}

/* GetAllIdpUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetAllIdpUsingGETOK struct {
	Payload *models.PageListOfIdp
}

func (o *GetAllIdpUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v2/admin/customers/{customerId}/idp][%d] getAllIdpUsingGETOK  %+v", 200, o.Payload)
}
func (o *GetAllIdpUsingGETOK) GetPayload() *models.PageListOfIdp {
	return o.Payload
}

func (o *GetAllIdpUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageListOfIdp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllIdpUsingGETBadRequest creates a GetAllIdpUsingGETBadRequest with default headers values
func NewGetAllIdpUsingGETBadRequest() *GetAllIdpUsingGETBadRequest {
	return &GetAllIdpUsingGETBadRequest{}
}

/* GetAllIdpUsingGETBadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type GetAllIdpUsingGETBadRequest struct {
}

func (o *GetAllIdpUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v2/admin/customers/{customerId}/idp][%d] getAllIdpUsingGETBadRequest ", 400)
}

func (o *GetAllIdpUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllIdpUsingGETUnauthorized creates a GetAllIdpUsingGETUnauthorized with default headers values
func NewGetAllIdpUsingGETUnauthorized() *GetAllIdpUsingGETUnauthorized {
	return &GetAllIdpUsingGETUnauthorized{}
}

/* GetAllIdpUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAllIdpUsingGETUnauthorized struct {
}

func (o *GetAllIdpUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v2/admin/customers/{customerId}/idp][%d] getAllIdpUsingGETUnauthorized ", 401)
}

func (o *GetAllIdpUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllIdpUsingGETForbidden creates a GetAllIdpUsingGETForbidden with default headers values
func NewGetAllIdpUsingGETForbidden() *GetAllIdpUsingGETForbidden {
	return &GetAllIdpUsingGETForbidden{}
}

/* GetAllIdpUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAllIdpUsingGETForbidden struct {
}

func (o *GetAllIdpUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v2/admin/customers/{customerId}/idp][%d] getAllIdpUsingGETForbidden ", 403)
}

func (o *GetAllIdpUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllIdpUsingGETNotFound creates a GetAllIdpUsingGETNotFound with default headers values
func NewGetAllIdpUsingGETNotFound() *GetAllIdpUsingGETNotFound {
	return &GetAllIdpUsingGETNotFound{}
}

/* GetAllIdpUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAllIdpUsingGETNotFound struct {
}

func (o *GetAllIdpUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v2/admin/customers/{customerId}/idp][%d] getAllIdpUsingGETNotFound ", 404)
}

func (o *GetAllIdpUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllIdpUsingGETTooManyRequests creates a GetAllIdpUsingGETTooManyRequests with default headers values
func NewGetAllIdpUsingGETTooManyRequests() *GetAllIdpUsingGETTooManyRequests {
	return &GetAllIdpUsingGETTooManyRequests{}
}

/* GetAllIdpUsingGETTooManyRequests describes a response with status code 429, with default header values.

TooManyRequest
*/
type GetAllIdpUsingGETTooManyRequests struct {
}

func (o *GetAllIdpUsingGETTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v2/admin/customers/{customerId}/idp][%d] getAllIdpUsingGETTooManyRequests ", 429)
}

func (o *GetAllIdpUsingGETTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
