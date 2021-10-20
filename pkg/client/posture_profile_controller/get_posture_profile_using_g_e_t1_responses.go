// Code generated by go-swagger; DO NOT EDIT.

package posture_profile_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haarchri/zpa-go-client/pkg/models"
)

// GetPostureProfileUsingGET1Reader is a Reader for the GetPostureProfileUsingGET1 structure.
type GetPostureProfileUsingGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPostureProfileUsingGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPostureProfileUsingGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetPostureProfileUsingGET1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPostureProfileUsingGET1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPostureProfileUsingGET1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPostureProfileUsingGET1OK creates a GetPostureProfileUsingGET1OK with default headers values
func NewGetPostureProfileUsingGET1OK() *GetPostureProfileUsingGET1OK {
	return &GetPostureProfileUsingGET1OK{}
}

/* GetPostureProfileUsingGET1OK describes a response with status code 200, with default header values.

OK
*/
type GetPostureProfileUsingGET1OK struct {
	Payload *models.PostureProfile
}

func (o *GetPostureProfileUsingGET1OK) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/posture/{id}][%d] getPostureProfileUsingGET1OK  %+v", 200, o.Payload)
}
func (o *GetPostureProfileUsingGET1OK) GetPayload() *models.PostureProfile {
	return o.Payload
}

func (o *GetPostureProfileUsingGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostureProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPostureProfileUsingGET1Unauthorized creates a GetPostureProfileUsingGET1Unauthorized with default headers values
func NewGetPostureProfileUsingGET1Unauthorized() *GetPostureProfileUsingGET1Unauthorized {
	return &GetPostureProfileUsingGET1Unauthorized{}
}

/* GetPostureProfileUsingGET1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetPostureProfileUsingGET1Unauthorized struct {
}

func (o *GetPostureProfileUsingGET1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/posture/{id}][%d] getPostureProfileUsingGET1Unauthorized ", 401)
}

func (o *GetPostureProfileUsingGET1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPostureProfileUsingGET1Forbidden creates a GetPostureProfileUsingGET1Forbidden with default headers values
func NewGetPostureProfileUsingGET1Forbidden() *GetPostureProfileUsingGET1Forbidden {
	return &GetPostureProfileUsingGET1Forbidden{}
}

/* GetPostureProfileUsingGET1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetPostureProfileUsingGET1Forbidden struct {
}

func (o *GetPostureProfileUsingGET1Forbidden) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/posture/{id}][%d] getPostureProfileUsingGET1Forbidden ", 403)
}

func (o *GetPostureProfileUsingGET1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPostureProfileUsingGET1NotFound creates a GetPostureProfileUsingGET1NotFound with default headers values
func NewGetPostureProfileUsingGET1NotFound() *GetPostureProfileUsingGET1NotFound {
	return &GetPostureProfileUsingGET1NotFound{}
}

/* GetPostureProfileUsingGET1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetPostureProfileUsingGET1NotFound struct {
}

func (o *GetPostureProfileUsingGET1NotFound) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/posture/{id}][%d] getPostureProfileUsingGET1NotFound ", 404)
}

func (o *GetPostureProfileUsingGET1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
