// Code generated by go-swagger; DO NOT EDIT.

package app_server_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteAppServerUsingDELETE1Reader is a Reader for the DeleteAppServerUsingDELETE1 structure.
type DeleteAppServerUsingDELETE1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAppServerUsingDELETE1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteAppServerUsingDELETE1NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteAppServerUsingDELETE1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteAppServerUsingDELETE1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAppServerUsingDELETE1NoContent creates a DeleteAppServerUsingDELETE1NoContent with default headers values
func NewDeleteAppServerUsingDELETE1NoContent() *DeleteAppServerUsingDELETE1NoContent {
	return &DeleteAppServerUsingDELETE1NoContent{}
}

/* DeleteAppServerUsingDELETE1NoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteAppServerUsingDELETE1NoContent struct {
}

func (o *DeleteAppServerUsingDELETE1NoContent) Error() string {
	return fmt.Sprintf("[DELETE /mgmtconfig/v1/admin/customers/{customerId}/server/{serverId}][%d] deleteAppServerUsingDELETE1NoContent ", 204)
}

func (o *DeleteAppServerUsingDELETE1NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAppServerUsingDELETE1Unauthorized creates a DeleteAppServerUsingDELETE1Unauthorized with default headers values
func NewDeleteAppServerUsingDELETE1Unauthorized() *DeleteAppServerUsingDELETE1Unauthorized {
	return &DeleteAppServerUsingDELETE1Unauthorized{}
}

/* DeleteAppServerUsingDELETE1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteAppServerUsingDELETE1Unauthorized struct {
}

func (o *DeleteAppServerUsingDELETE1Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /mgmtconfig/v1/admin/customers/{customerId}/server/{serverId}][%d] deleteAppServerUsingDELETE1Unauthorized ", 401)
}

func (o *DeleteAppServerUsingDELETE1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAppServerUsingDELETE1Forbidden creates a DeleteAppServerUsingDELETE1Forbidden with default headers values
func NewDeleteAppServerUsingDELETE1Forbidden() *DeleteAppServerUsingDELETE1Forbidden {
	return &DeleteAppServerUsingDELETE1Forbidden{}
}

/* DeleteAppServerUsingDELETE1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteAppServerUsingDELETE1Forbidden struct {
}

func (o *DeleteAppServerUsingDELETE1Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /mgmtconfig/v1/admin/customers/{customerId}/server/{serverId}][%d] deleteAppServerUsingDELETE1Forbidden ", 403)
}

func (o *DeleteAppServerUsingDELETE1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
