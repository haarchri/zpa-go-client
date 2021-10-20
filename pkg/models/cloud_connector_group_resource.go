// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CloudConnectorGroupResource cloud connector group resource
//
// swagger:model CloudConnectorGroupResource
type CloudConnectorGroupResource struct {

	// cloud connectors
	CloudConnectors []*Znf `json:"cloudConnectors"`

	// Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.
	CreationTime int32 `json:"creationTime,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// geo location Id
	GeoLocationID int64 `json:"geoLocationId,omitempty"`

	// Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.
	ID string `json:"id,omitempty"`

	// Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.
	ModifiedBy string `json:"modifiedBy,omitempty"`

	// Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.
	ModifiedTime int32 `json:"modifiedTime,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// zia cloud
	ZiaCloud string `json:"ziaCloud,omitempty"`

	// zia org Id
	ZiaOrgID int64 `json:"ziaOrgId,omitempty"`
}

// Validate validates this cloud connector group resource
func (m *CloudConnectorGroupResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudConnectors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudConnectorGroupResource) validateCloudConnectors(formats strfmt.Registry) error {
	if swag.IsZero(m.CloudConnectors) { // not required
		return nil
	}

	for i := 0; i < len(m.CloudConnectors); i++ {
		if swag.IsZero(m.CloudConnectors[i]) { // not required
			continue
		}

		if m.CloudConnectors[i] != nil {
			if err := m.CloudConnectors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cloudConnectors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cloudConnectors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cloud connector group resource based on the context it is used
func (m *CloudConnectorGroupResource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCloudConnectors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudConnectorGroupResource) contextValidateCloudConnectors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CloudConnectors); i++ {

		if m.CloudConnectors[i] != nil {
			if err := m.CloudConnectors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cloudConnectors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cloudConnectors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudConnectorGroupResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudConnectorGroupResource) UnmarshalBinary(b []byte) error {
	var res CloudConnectorGroupResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}