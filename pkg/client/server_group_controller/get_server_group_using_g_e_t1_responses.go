// Code generated by go-swagger; DO NOT EDIT.

package server_group_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haarchri/zpa-go-client/pkg/models"
)

// GetServerGroupUsingGET1Reader is a Reader for the GetServerGroupUsingGET1 structure.
type GetServerGroupUsingGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServerGroupUsingGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServerGroupUsingGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetServerGroupUsingGET1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetServerGroupUsingGET1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetServerGroupUsingGET1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetServerGroupUsingGET1OK creates a GetServerGroupUsingGET1OK with default headers values
func NewGetServerGroupUsingGET1OK() *GetServerGroupUsingGET1OK {
	return &GetServerGroupUsingGET1OK{}
}

/* GetServerGroupUsingGET1OK describes a response with status code 200, with default header values.

OK
*/
type GetServerGroupUsingGET1OK struct {
	Payload *models.ServerGroupDTO
}

func (o *GetServerGroupUsingGET1OK) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/serverGroup/{groupId}][%d] getServerGroupUsingGET1OK  %+v", 200, o.Payload)
}
func (o *GetServerGroupUsingGET1OK) GetPayload() *models.ServerGroupDTO {
	return o.Payload
}

func (o *GetServerGroupUsingGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerGroupDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServerGroupUsingGET1Unauthorized creates a GetServerGroupUsingGET1Unauthorized with default headers values
func NewGetServerGroupUsingGET1Unauthorized() *GetServerGroupUsingGET1Unauthorized {
	return &GetServerGroupUsingGET1Unauthorized{}
}

/* GetServerGroupUsingGET1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetServerGroupUsingGET1Unauthorized struct {
}

func (o *GetServerGroupUsingGET1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/serverGroup/{groupId}][%d] getServerGroupUsingGET1Unauthorized ", 401)
}

func (o *GetServerGroupUsingGET1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetServerGroupUsingGET1Forbidden creates a GetServerGroupUsingGET1Forbidden with default headers values
func NewGetServerGroupUsingGET1Forbidden() *GetServerGroupUsingGET1Forbidden {
	return &GetServerGroupUsingGET1Forbidden{}
}

/* GetServerGroupUsingGET1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetServerGroupUsingGET1Forbidden struct {
}

func (o *GetServerGroupUsingGET1Forbidden) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/serverGroup/{groupId}][%d] getServerGroupUsingGET1Forbidden ", 403)
}

func (o *GetServerGroupUsingGET1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetServerGroupUsingGET1NotFound creates a GetServerGroupUsingGET1NotFound with default headers values
func NewGetServerGroupUsingGET1NotFound() *GetServerGroupUsingGET1NotFound {
	return &GetServerGroupUsingGET1NotFound{}
}

/* GetServerGroupUsingGET1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetServerGroupUsingGET1NotFound struct {
}

func (o *GetServerGroupUsingGET1NotFound) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/serverGroup/{groupId}][%d] getServerGroupUsingGET1NotFound ", 404)
}

func (o *GetServerGroupUsingGET1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}