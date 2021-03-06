// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AppConnectorGroupResource app connector group resource
//
// swagger:model AppConnectorGroupResource
type AppConnectorGroupResource struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this app connector group resource
func (m *AppConnectorGroupResource) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this app connector group resource based on context it is used
func (m *AppConnectorGroupResource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AppConnectorGroupResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppConnectorGroupResource) UnmarshalBinary(b []byte) error {
	var res AppConnectorGroupResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
