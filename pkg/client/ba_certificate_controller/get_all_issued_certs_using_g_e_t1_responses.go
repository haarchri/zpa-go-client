// Code generated by go-swagger; DO NOT EDIT.

package ba_certificate_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haarchri/zpa-go-client/pkg/models"
)

// GetAllIssuedCertsUsingGET1Reader is a Reader for the GetAllIssuedCertsUsingGET1 structure.
type GetAllIssuedCertsUsingGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllIssuedCertsUsingGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllIssuedCertsUsingGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAllIssuedCertsUsingGET1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAllIssuedCertsUsingGET1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAllIssuedCertsUsingGET1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllIssuedCertsUsingGET1OK creates a GetAllIssuedCertsUsingGET1OK with default headers values
func NewGetAllIssuedCertsUsingGET1OK() *GetAllIssuedCertsUsingGET1OK {
	return &GetAllIssuedCertsUsingGET1OK{}
}

/* GetAllIssuedCertsUsingGET1OK describes a response with status code 200, with default header values.

OK
*/
type GetAllIssuedCertsUsingGET1OK struct {
	Payload []*models.BACertificate
}

func (o *GetAllIssuedCertsUsingGET1OK) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/clientlessCertificate/issued][%d] getAllIssuedCertsUsingGET1OK  %+v", 200, o.Payload)
}
func (o *GetAllIssuedCertsUsingGET1OK) GetPayload() []*models.BACertificate {
	return o.Payload
}

func (o *GetAllIssuedCertsUsingGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllIssuedCertsUsingGET1Unauthorized creates a GetAllIssuedCertsUsingGET1Unauthorized with default headers values
func NewGetAllIssuedCertsUsingGET1Unauthorized() *GetAllIssuedCertsUsingGET1Unauthorized {
	return &GetAllIssuedCertsUsingGET1Unauthorized{}
}

/* GetAllIssuedCertsUsingGET1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAllIssuedCertsUsingGET1Unauthorized struct {
}

func (o *GetAllIssuedCertsUsingGET1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/clientlessCertificate/issued][%d] getAllIssuedCertsUsingGET1Unauthorized ", 401)
}

func (o *GetAllIssuedCertsUsingGET1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllIssuedCertsUsingGET1Forbidden creates a GetAllIssuedCertsUsingGET1Forbidden with default headers values
func NewGetAllIssuedCertsUsingGET1Forbidden() *GetAllIssuedCertsUsingGET1Forbidden {
	return &GetAllIssuedCertsUsingGET1Forbidden{}
}

/* GetAllIssuedCertsUsingGET1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAllIssuedCertsUsingGET1Forbidden struct {
}

func (o *GetAllIssuedCertsUsingGET1Forbidden) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/clientlessCertificate/issued][%d] getAllIssuedCertsUsingGET1Forbidden ", 403)
}

func (o *GetAllIssuedCertsUsingGET1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllIssuedCertsUsingGET1NotFound creates a GetAllIssuedCertsUsingGET1NotFound with default headers values
func NewGetAllIssuedCertsUsingGET1NotFound() *GetAllIssuedCertsUsingGET1NotFound {
	return &GetAllIssuedCertsUsingGET1NotFound{}
}

/* GetAllIssuedCertsUsingGET1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAllIssuedCertsUsingGET1NotFound struct {
}

func (o *GetAllIssuedCertsUsingGET1NotFound) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/clientlessCertificate/issued][%d] getAllIssuedCertsUsingGET1NotFound ", 404)
}

func (o *GetAllIssuedCertsUsingGET1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}