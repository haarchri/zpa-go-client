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

// SCIMAttributeHeader s c i m attribute header
//
// swagger:model SCIMAttributeHeader
type SCIMAttributeHeader struct {

	// canonical values
	CanonicalValues []string `json:"canonicalValues"`

	// case sensitive
	CaseSensitive bool `json:"caseSensitive,omitempty"`

	// creation time
	CreationTime string `json:"creationTime,omitempty"`

	// data type
	DataType string `json:"dataType,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// idp Id
	IdpID string `json:"idpId,omitempty"`

	// modified by
	ModifiedBy string `json:"modifiedBy,omitempty"`

	// modified time
	ModifiedTime string `json:"modifiedTime,omitempty"`

	// multivalued
	Multivalued bool `json:"multivalued,omitempty"`

	// mutability
	Mutability string `json:"mutability,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// required
	Required bool `json:"required,omitempty"`

	// returned
	Returned string `json:"returned,omitempty"`

	// schema URI
	SchemaURI string `json:"schemaURI,omitempty"`

	// uniqueness
	Uniqueness bool `json:"uniqueness,omitempty"`
}

// Validate validates this s c i m attribute header
func (m *SCIMAttributeHeader) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SCIMAttributeHeader) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this s c i m attribute header based on context it is used
func (m *SCIMAttributeHeader) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SCIMAttributeHeader) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SCIMAttributeHeader) UnmarshalBinary(b []byte) error {
	var res SCIMAttributeHeader
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
