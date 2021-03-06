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

// GetIdpUsingGET1Reader is a Reader for the GetIdpUsingGET1 structure.
type GetIdpUsingGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIdpUsingGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIdpUsingGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIdpUsingGET1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIdpUsingGET1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIdpUsingGET1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIdpUsingGET1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIdpUsingGET1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIdpUsingGET1OK creates a GetIdpUsingGET1OK with default headers values
func NewGetIdpUsingGET1OK() *GetIdpUsingGET1OK {
	return &GetIdpUsingGET1OK{}
}

/* GetIdpUsingGET1OK describes a response with status code 200, with default header values.

OK
*/
type GetIdpUsingGET1OK struct {
	Payload []*models.Idp
}

func (o *GetIdpUsingGET1OK) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/idp][%d] getIdpUsingGET1OK  %+v", 200, o.Payload)
}
func (o *GetIdpUsingGET1OK) GetPayload() []*models.Idp {
	return o.Payload
}

func (o *GetIdpUsingGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdpUsingGET1BadRequest creates a GetIdpUsingGET1BadRequest with default headers values
func NewGetIdpUsingGET1BadRequest() *GetIdpUsingGET1BadRequest {
	return &GetIdpUsingGET1BadRequest{}
}

/* GetIdpUsingGET1BadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type GetIdpUsingGET1BadRequest struct {
}

func (o *GetIdpUsingGET1BadRequest) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/idp][%d] getIdpUsingGET1BadRequest ", 400)
}

func (o *GetIdpUsingGET1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIdpUsingGET1Unauthorized creates a GetIdpUsingGET1Unauthorized with default headers values
func NewGetIdpUsingGET1Unauthorized() *GetIdpUsingGET1Unauthorized {
	return &GetIdpUsingGET1Unauthorized{}
}

/* GetIdpUsingGET1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetIdpUsingGET1Unauthorized struct {
}

func (o *GetIdpUsingGET1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/idp][%d] getIdpUsingGET1Unauthorized ", 401)
}

func (o *GetIdpUsingGET1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIdpUsingGET1Forbidden creates a GetIdpUsingGET1Forbidden with default headers values
func NewGetIdpUsingGET1Forbidden() *GetIdpUsingGET1Forbidden {
	return &GetIdpUsingGET1Forbidden{}
}

/* GetIdpUsingGET1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetIdpUsingGET1Forbidden struct {
}

func (o *GetIdpUsingGET1Forbidden) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/idp][%d] getIdpUsingGET1Forbidden ", 403)
}

func (o *GetIdpUsingGET1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIdpUsingGET1NotFound creates a GetIdpUsingGET1NotFound with default headers values
func NewGetIdpUsingGET1NotFound() *GetIdpUsingGET1NotFound {
	return &GetIdpUsingGET1NotFound{}
}

/* GetIdpUsingGET1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetIdpUsingGET1NotFound struct {
}

func (o *GetIdpUsingGET1NotFound) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/idp][%d] getIdpUsingGET1NotFound ", 404)
}

func (o *GetIdpUsingGET1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIdpUsingGET1TooManyRequests creates a GetIdpUsingGET1TooManyRequests with default headers values
func NewGetIdpUsingGET1TooManyRequests() *GetIdpUsingGET1TooManyRequests {
	return &GetIdpUsingGET1TooManyRequests{}
}

/* GetIdpUsingGET1TooManyRequests describes a response with status code 429, with default header values.

TooManyRequest
*/
type GetIdpUsingGET1TooManyRequests struct {
}

func (o *GetIdpUsingGET1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/idp][%d] getIdpUsingGET1TooManyRequests ", 429)
}

func (o *GetIdpUsingGET1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
