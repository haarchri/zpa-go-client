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

// AddAppServerUsingPOST1Reader is a Reader for the AddAppServerUsingPOST1 structure.
type AddAppServerUsingPOST1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddAppServerUsingPOST1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddAppServerUsingPOST1Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAddAppServerUsingPOST1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddAppServerUsingPOST1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddAppServerUsingPOST1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddAppServerUsingPOST1Created creates a AddAppServerUsingPOST1Created with default headers values
func NewAddAppServerUsingPOST1Created() *AddAppServerUsingPOST1Created {
	return &AddAppServerUsingPOST1Created{}
}

/* AddAppServerUsingPOST1Created describes a response with status code 201, with default header values.

Created
*/
type AddAppServerUsingPOST1Created struct {
	Payload *models.ApplicationServer
}

func (o *AddAppServerUsingPOST1Created) Error() string {
	return fmt.Sprintf("[POST /mgmtconfig/v1/admin/customers/{customerId}/server][%d] addAppServerUsingPOST1Created  %+v", 201, o.Payload)
}
func (o *AddAppServerUsingPOST1Created) GetPayload() *models.ApplicationServer {
	return o.Payload
}

func (o *AddAppServerUsingPOST1Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApplicationServer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddAppServerUsingPOST1Unauthorized creates a AddAppServerUsingPOST1Unauthorized with default headers values
func NewAddAppServerUsingPOST1Unauthorized() *AddAppServerUsingPOST1Unauthorized {
	return &AddAppServerUsingPOST1Unauthorized{}
}

/* AddAppServerUsingPOST1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AddAppServerUsingPOST1Unauthorized struct {
}

func (o *AddAppServerUsingPOST1Unauthorized) Error() string {
	return fmt.Sprintf("[POST /mgmtconfig/v1/admin/customers/{customerId}/server][%d] addAppServerUsingPOST1Unauthorized ", 401)
}

func (o *AddAppServerUsingPOST1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddAppServerUsingPOST1Forbidden creates a AddAppServerUsingPOST1Forbidden with default headers values
func NewAddAppServerUsingPOST1Forbidden() *AddAppServerUsingPOST1Forbidden {
	return &AddAppServerUsingPOST1Forbidden{}
}

/* AddAppServerUsingPOST1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AddAppServerUsingPOST1Forbidden struct {
}

func (o *AddAppServerUsingPOST1Forbidden) Error() string {
	return fmt.Sprintf("[POST /mgmtconfig/v1/admin/customers/{customerId}/server][%d] addAppServerUsingPOST1Forbidden ", 403)
}

func (o *AddAppServerUsingPOST1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddAppServerUsingPOST1NotFound creates a AddAppServerUsingPOST1NotFound with default headers values
func NewAddAppServerUsingPOST1NotFound() *AddAppServerUsingPOST1NotFound {
	return &AddAppServerUsingPOST1NotFound{}
}

/* AddAppServerUsingPOST1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type AddAppServerUsingPOST1NotFound struct {
}

func (o *AddAppServerUsingPOST1NotFound) Error() string {
	return fmt.Sprintf("[POST /mgmtconfig/v1/admin/customers/{customerId}/server][%d] addAppServerUsingPOST1NotFound ", 404)
}

func (o *AddAppServerUsingPOST1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
