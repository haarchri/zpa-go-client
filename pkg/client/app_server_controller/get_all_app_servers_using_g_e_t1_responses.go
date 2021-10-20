// Code generated by go-swagger; DO NOT EDIT.

package app_server_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haarchri/zpa-go-client/pkg/models"
)

// GetAllAppServersUsingGET1Reader is a Reader for the GetAllAppServersUsingGET1 structure.
type GetAllAppServersUsingGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllAppServersUsingGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllAppServersUsingGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAllAppServersUsingGET1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAllAppServersUsingGET1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAllAppServersUsingGET1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllAppServersUsingGET1OK creates a GetAllAppServersUsingGET1OK with default headers values
func NewGetAllAppServersUsingGET1OK() *GetAllAppServersUsingGET1OK {
	return &GetAllAppServersUsingGET1OK{}
}

/* GetAllAppServersUsingGET1OK describes a response with status code 200, with default header values.

OK
*/
type GetAllAppServersUsingGET1OK struct {
	Payload *models.PageListOfApplicationServer
}

func (o *GetAllAppServersUsingGET1OK) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/server][%d] getAllAppServersUsingGET1OK  %+v", 200, o.Payload)
}
func (o *GetAllAppServersUsingGET1OK) GetPayload() *models.PageListOfApplicationServer {
	return o.Payload
}

func (o *GetAllAppServersUsingGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageListOfApplicationServer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllAppServersUsingGET1Unauthorized creates a GetAllAppServersUsingGET1Unauthorized with default headers values
func NewGetAllAppServersUsingGET1Unauthorized() *GetAllAppServersUsingGET1Unauthorized {
	return &GetAllAppServersUsingGET1Unauthorized{}
}

/* GetAllAppServersUsingGET1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAllAppServersUsingGET1Unauthorized struct {
}

func (o *GetAllAppServersUsingGET1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/server][%d] getAllAppServersUsingGET1Unauthorized ", 401)
}

func (o *GetAllAppServersUsingGET1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllAppServersUsingGET1Forbidden creates a GetAllAppServersUsingGET1Forbidden with default headers values
func NewGetAllAppServersUsingGET1Forbidden() *GetAllAppServersUsingGET1Forbidden {
	return &GetAllAppServersUsingGET1Forbidden{}
}

/* GetAllAppServersUsingGET1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAllAppServersUsingGET1Forbidden struct {
}

func (o *GetAllAppServersUsingGET1Forbidden) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/server][%d] getAllAppServersUsingGET1Forbidden ", 403)
}

func (o *GetAllAppServersUsingGET1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllAppServersUsingGET1NotFound creates a GetAllAppServersUsingGET1NotFound with default headers values
func NewGetAllAppServersUsingGET1NotFound() *GetAllAppServersUsingGET1NotFound {
	return &GetAllAppServersUsingGET1NotFound{}
}

/* GetAllAppServersUsingGET1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAllAppServersUsingGET1NotFound struct {
}

func (o *GetAllAppServersUsingGET1NotFound) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/server][%d] getAllAppServersUsingGET1NotFound ", 404)
}

func (o *GetAllAppServersUsingGET1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}