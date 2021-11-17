// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TrustedNetwork trusted network
//
// swagger:model TrustedNetwork
type TrustedNetwork struct {

	// creation time
	CreationTime string `json:"creationTime,omitempty"`

	// domain
	Domain string `json:"domain,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// master customer Id
	MasterCustomerID string `json:"masterCustomerId,omitempty"`

	// modified by
	ModifiedBy string `json:"modifiedBy,omitempty"`

	// modified time
	ModifiedTime string `json:"modifiedTime,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// network Id
	NetworkID string `json:"networkId,omitempty"`

	// zscaler cloud
	ZscalerCloud string `json:"zscalerCloud,omitempty"`
}

// Validate validates this trusted network
func (m *TrustedNetwork) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TrustedNetwork) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this trusted network based on context it is used
func (m *TrustedNetwork) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TrustedNetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TrustedNetwork) UnmarshalBinary(b []byte) error {
	var res TrustedNetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
