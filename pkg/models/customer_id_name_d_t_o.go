// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CustomerIDNameDTO customer Id name d t o
//
// swagger:model CustomerIdNameDTO
type CustomerIDNameDTO struct {

	// customer Id
	CustomerID string `json:"customerId,omitempty"`

	// exclude constellation
	ExcludeConstellation bool `json:"excludeConstellation,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this customer Id name d t o
func (m *CustomerIDNameDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this customer Id name d t o based on context it is used
func (m *CustomerIDNameDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustomerIDNameDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomerIDNameDTO) UnmarshalBinary(b []byte) error {
	var res CustomerIDNameDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
