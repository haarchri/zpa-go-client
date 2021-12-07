// Code generated by go-swagger; DO NOT EDIT.

package lss_config_controller_v_2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteLssUsingDELETE3Reader is a Reader for the DeleteLssUsingDELETE3 structure.
type DeleteLssUsingDELETE3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLssUsingDELETE3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteLssUsingDELETE3NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteLssUsingDELETE3BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteLssUsingDELETE3Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteLssUsingDELETE3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteLssUsingDELETE3TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteLssUsingDELETE3NoContent creates a DeleteLssUsingDELETE3NoContent with default headers values
func NewDeleteLssUsingDELETE3NoContent() *DeleteLssUsingDELETE3NoContent {
	return &DeleteLssUsingDELETE3NoContent{}
}

/* DeleteLssUsingDELETE3NoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteLssUsingDELETE3NoContent struct {
}

func (o *DeleteLssUsingDELETE3NoContent) Error() string {
	return fmt.Sprintf("[DELETE /mgmtconfig/v2/admin/customers/{customerId}/lssConfig/{lssId}][%d] deleteLssUsingDELETE3NoContent ", 204)
}

func (o *DeleteLssUsingDELETE3NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLssUsingDELETE3BadRequest creates a DeleteLssUsingDELETE3BadRequest with default headers values
func NewDeleteLssUsingDELETE3BadRequest() *DeleteLssUsingDELETE3BadRequest {
	return &DeleteLssUsingDELETE3BadRequest{}
}

/* DeleteLssUsingDELETE3BadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type DeleteLssUsingDELETE3BadRequest struct {
}

func (o *DeleteLssUsingDELETE3BadRequest) Error() string {
	return fmt.Sprintf("[DELETE /mgmtconfig/v2/admin/customers/{customerId}/lssConfig/{lssId}][%d] deleteLssUsingDELETE3BadRequest ", 400)
}

func (o *DeleteLssUsingDELETE3BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLssUsingDELETE3Unauthorized creates a DeleteLssUsingDELETE3Unauthorized with default headers values
func NewDeleteLssUsingDELETE3Unauthorized() *DeleteLssUsingDELETE3Unauthorized {
	return &DeleteLssUsingDELETE3Unauthorized{}
}

/* DeleteLssUsingDELETE3Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteLssUsingDELETE3Unauthorized struct {
}

func (o *DeleteLssUsingDELETE3Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /mgmtconfig/v2/admin/customers/{customerId}/lssConfig/{lssId}][%d] deleteLssUsingDELETE3Unauthorized ", 401)
}

func (o *DeleteLssUsingDELETE3Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLssUsingDELETE3Forbidden creates a DeleteLssUsingDELETE3Forbidden with default headers values
func NewDeleteLssUsingDELETE3Forbidden() *DeleteLssUsingDELETE3Forbidden {
	return &DeleteLssUsingDELETE3Forbidden{}
}

/* DeleteLssUsingDELETE3Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteLssUsingDELETE3Forbidden struct {
}

func (o *DeleteLssUsingDELETE3Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /mgmtconfig/v2/admin/customers/{customerId}/lssConfig/{lssId}][%d] deleteLssUsingDELETE3Forbidden ", 403)
}

func (o *DeleteLssUsingDELETE3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLssUsingDELETE3TooManyRequests creates a DeleteLssUsingDELETE3TooManyRequests with default headers values
func NewDeleteLssUsingDELETE3TooManyRequests() *DeleteLssUsingDELETE3TooManyRequests {
	return &DeleteLssUsingDELETE3TooManyRequests{}
}

/* DeleteLssUsingDELETE3TooManyRequests describes a response with status code 429, with default header values.

TooManyRequest
*/
type DeleteLssUsingDELETE3TooManyRequests struct {
}

func (o *DeleteLssUsingDELETE3TooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /mgmtconfig/v2/admin/customers/{customerId}/lssConfig/{lssId}][%d] deleteLssUsingDELETE3TooManyRequests ", 429)
}

func (o *DeleteLssUsingDELETE3TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}