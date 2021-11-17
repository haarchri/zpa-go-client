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

// GetAppServerUsingGET1Reader is a Reader for the GetAppServerUsingGET1 structure.
type GetAppServerUsingGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAppServerUsingGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAppServerUsingGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAppServerUsingGET1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAppServerUsingGET1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAppServerUsingGET1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAppServerUsingGET1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAppServerUsingGET1OK creates a GetAppServerUsingGET1OK with default headers values
func NewGetAppServerUsingGET1OK() *GetAppServerUsingGET1OK {
	return &GetAppServerUsingGET1OK{}
}

/* GetAppServerUsingGET1OK describes a response with status code 200, with default header values.

OK
*/
type GetAppServerUsingGET1OK struct {
	Payload *models.ApplicationServer
}

func (o *GetAppServerUsingGET1OK) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/server/{serverId}][%d] getAppServerUsingGET1OK  %+v", 200, o.Payload)
}
func (o *GetAppServerUsingGET1OK) GetPayload() *models.ApplicationServer {
	return o.Payload
}

func (o *GetAppServerUsingGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApplicationServer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppServerUsingGET1BadRequest creates a GetAppServerUsingGET1BadRequest with default headers values
func NewGetAppServerUsingGET1BadRequest() *GetAppServerUsingGET1BadRequest {
	return &GetAppServerUsingGET1BadRequest{}
}

/* GetAppServerUsingGET1BadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type GetAppServerUsingGET1BadRequest struct {
}

func (o *GetAppServerUsingGET1BadRequest) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/server/{serverId}][%d] getAppServerUsingGET1BadRequest ", 400)
}

func (o *GetAppServerUsingGET1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAppServerUsingGET1Unauthorized creates a GetAppServerUsingGET1Unauthorized with default headers values
func NewGetAppServerUsingGET1Unauthorized() *GetAppServerUsingGET1Unauthorized {
	return &GetAppServerUsingGET1Unauthorized{}
}

/* GetAppServerUsingGET1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAppServerUsingGET1Unauthorized struct {
}

func (o *GetAppServerUsingGET1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/server/{serverId}][%d] getAppServerUsingGET1Unauthorized ", 401)
}

func (o *GetAppServerUsingGET1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAppServerUsingGET1Forbidden creates a GetAppServerUsingGET1Forbidden with default headers values
func NewGetAppServerUsingGET1Forbidden() *GetAppServerUsingGET1Forbidden {
	return &GetAppServerUsingGET1Forbidden{}
}

/* GetAppServerUsingGET1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAppServerUsingGET1Forbidden struct {
}

func (o *GetAppServerUsingGET1Forbidden) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/server/{serverId}][%d] getAppServerUsingGET1Forbidden ", 403)
}

func (o *GetAppServerUsingGET1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAppServerUsingGET1NotFound creates a GetAppServerUsingGET1NotFound with default headers values
func NewGetAppServerUsingGET1NotFound() *GetAppServerUsingGET1NotFound {
	return &GetAppServerUsingGET1NotFound{}
}

/* GetAppServerUsingGET1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAppServerUsingGET1NotFound struct {
}

func (o *GetAppServerUsingGET1NotFound) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/server/{serverId}][%d] getAppServerUsingGET1NotFound ", 404)
}

func (o *GetAppServerUsingGET1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
