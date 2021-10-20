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

// Znf znf
//
// swagger:model Znf
type Znf struct {

	// creation time
	CreationTime int32 `json:"creationTime,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// fingerprint
	Fingerprint string `json:"fingerprint,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// ip Acl
	IPACL []string `json:"ipAcl"`

	// issued cert Id
	IssuedCertID int64 `json:"issuedCertId,omitempty"`

	// modified by
	ModifiedBy int64 `json:"modifiedBy,omitempty"`

	// modified time
	ModifiedTime int32 `json:"modifiedTime,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// signing cert
	SigningCert map[string]string `json:"signingCert,omitempty"`
}

// Validate validates this znf
func (m *Znf) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Znf) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this znf based on context it is used
func (m *Znf) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Znf) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Znf) UnmarshalBinary(b []byte) error {
	var res Znf
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
