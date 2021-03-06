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

// BACertificate b a certificate
//
// swagger:model BACertificate
type BACertificate struct {

	// c name
	CName string `json:"cName,omitempty"`

	// cert chain
	CertChain string `json:"certChain,omitempty"`

	// The certificate text is in PEM format.
	// Required: true
	Certificate *string `json:"certificate"`

	// creation time
	CreationTime string `json:"creationTime,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// issued by
	IssuedBy string `json:"issuedBy,omitempty"`

	// issued to
	IssuedTo string `json:"issuedTo,omitempty"`

	// modified by
	ModifiedBy string `json:"modifiedBy,omitempty"`

	// modified time
	ModifiedTime string `json:"modifiedTime,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// public key
	PublicKey string `json:"publicKey,omitempty"`

	// san
	San []string `json:"san"`

	// serial no
	SerialNo string `json:"serialNo,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// valid from in epoch sec
	ValidFromInEpochSec string `json:"validFromInEpochSec,omitempty"`

	// valid to in epoch sec
	ValidToInEpochSec string `json:"validToInEpochSec,omitempty"`
}

// Validate validates this b a certificate
func (m *BACertificate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCertificate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BACertificate) validateCertificate(formats strfmt.Registry) error {

	if err := validate.Required("certificate", "body", m.Certificate); err != nil {
		return err
	}

	return nil
}

func (m *BACertificate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this b a certificate based on context it is used
func (m *BACertificate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BACertificate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BACertificate) UnmarshalBinary(b []byte) error {
	var res BACertificate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
