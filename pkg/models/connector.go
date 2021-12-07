// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Connector connector
//
// swagger:model Connector
type Connector struct {

	// app connector group Id
	AppConnectorGroupID string `json:"appConnectorGroupId,omitempty"`

	// app connector group name
	AppConnectorGroupName string `json:"appConnectorGroupName,omitempty"`

	// application start time
	ApplicationStartTime string `json:"applicationStartTime,omitempty"`

	// Read only. Ignored in PUT/POST calls. Expected values: UNKNOWN/ZPN_STATUS_AUTHENTICATED(1)/ZPN_STATUS_DISCONNECTED
	// Enum: [UNKNOWN ZPN_STATUS_AUTHENTICATED ZPN_STATUS_DISCONNECTED]
	ControlChannelStatus string `json:"controlChannelStatus,omitempty"`

	// creation time
	CreationTime string `json:"creationTime,omitempty"`

	// Read only. Ignored in PUT/POST calls.
	CtrlBrokerName string `json:"ctrlBrokerName,omitempty"`

	// Read only. Ignored in PUT/POST calls.
	CurrentVersion string `json:"currentVersion,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// enrollment cert
	EnrollmentCert map[string]string `json:"enrollmentCert,omitempty"`

	// Read only. Ignored in PUT/POST calls.
	ExpectedUpgradeTime string `json:"expectedUpgradeTime,omitempty"`

	// Read only. Ignored in PUT/POST calls.
	ExpectedVersion string `json:"expectedVersion,omitempty"`

	// fingerprint
	Fingerprint string `json:"fingerprint,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// ip Acl
	IPACL []string `json:"ipAcl"`

	// issued cert Id
	IssuedCertID string `json:"issuedCertId,omitempty"`

	// Read only. Ignored in PUT/POST calls.
	LastBrokerConnectTime string `json:"lastBrokerConnectTime,omitempty"`

	// Read only. Ignored in PUT/POST calls.
	LastBrokerConnectTimeDuration string `json:"lastBrokerConnectTimeDuration,omitempty"`

	// Read only. Ignored in PUT/POST calls.
	LastBrokerDisconnectTime string `json:"lastBrokerDisconnectTime,omitempty"`

	// Read only. Ignored in PUT/POST calls.
	LastBrokerDisconnectTimeDuration string `json:"lastBrokerDisconnectTimeDuration,omitempty"`

	// Read only. Ignored in PUT/POST calls.
	LastUpgradeTime string `json:"lastUpgradeTime,omitempty"`

	// latitude
	Latitude float64 `json:"latitude,omitempty"`

	// location
	Location string `json:"location,omitempty"`

	// longitude
	Longitude float64 `json:"longitude,omitempty"`

	// modified by
	ModifiedBy string `json:"modifiedBy,omitempty"`

	// modified time
	ModifiedTime string `json:"modifiedTime,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// platform
	Platform string `json:"platform,omitempty"`

	// Read only. Ignored in PUT/POST calls.
	PreviousVersion string `json:"previousVersion,omitempty"`

	// private Ip
	PrivateIP string `json:"privateIp,omitempty"`

	// provisioning key Id
	ProvisioningKeyID string `json:"provisioningKeyId,omitempty"`

	// provisioning key name
	ProvisioningKeyName string `json:"provisioningKeyName,omitempty"`

	// public Ip
	PublicIP string `json:"publicIp,omitempty"`

	// sarge version
	SargeVersion string `json:"sargeVersion,omitempty"`

	// Read only. Ignored in PUT/POST calls.
	UpgradeAttempt string `json:"upgradeAttempt,omitempty"`

	// Read only. Ignored in PUT/POST calls.
	// Enum: [COMPLETE IN_PROGRESS FAILED UNKNOWN RESTARTING]
	UpgradeStatus string `json:"upgradeStatus,omitempty"`
}

// Validate validates this connector
func (m *Connector) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateControlChannelStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpgradeStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var connectorTypeControlChannelStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UNKNOWN","ZPN_STATUS_AUTHENTICATED","ZPN_STATUS_DISCONNECTED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		connectorTypeControlChannelStatusPropEnum = append(connectorTypeControlChannelStatusPropEnum, v)
	}
}

const (

	// ConnectorControlChannelStatusUNKNOWN captures enum value "UNKNOWN"
	ConnectorControlChannelStatusUNKNOWN string = "UNKNOWN"

	// ConnectorControlChannelStatusZPNSTATUSAUTHENTICATED captures enum value "ZPN_STATUS_AUTHENTICATED"
	ConnectorControlChannelStatusZPNSTATUSAUTHENTICATED string = "ZPN_STATUS_AUTHENTICATED"

	// ConnectorControlChannelStatusZPNSTATUSDISCONNECTED captures enum value "ZPN_STATUS_DISCONNECTED"
	ConnectorControlChannelStatusZPNSTATUSDISCONNECTED string = "ZPN_STATUS_DISCONNECTED"
)

// prop value enum
func (m *Connector) validateControlChannelStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, connectorTypeControlChannelStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Connector) validateControlChannelStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.ControlChannelStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateControlChannelStatusEnum("controlChannelStatus", "body", m.ControlChannelStatus); err != nil {
		return err
	}

	return nil
}

func (m *Connector) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var connectorTypeUpgradeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["COMPLETE","IN_PROGRESS","FAILED","UNKNOWN","RESTARTING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		connectorTypeUpgradeStatusPropEnum = append(connectorTypeUpgradeStatusPropEnum, v)
	}
}

const (

	// ConnectorUpgradeStatusCOMPLETE captures enum value "COMPLETE"
	ConnectorUpgradeStatusCOMPLETE string = "COMPLETE"

	// ConnectorUpgradeStatusINPROGRESS captures enum value "IN_PROGRESS"
	ConnectorUpgradeStatusINPROGRESS string = "IN_PROGRESS"

	// ConnectorUpgradeStatusFAILED captures enum value "FAILED"
	ConnectorUpgradeStatusFAILED string = "FAILED"

	// ConnectorUpgradeStatusUNKNOWN captures enum value "UNKNOWN"
	ConnectorUpgradeStatusUNKNOWN string = "UNKNOWN"

	// ConnectorUpgradeStatusRESTARTING captures enum value "RESTARTING"
	ConnectorUpgradeStatusRESTARTING string = "RESTARTING"
)

// prop value enum
func (m *Connector) validateUpgradeStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, connectorTypeUpgradeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Connector) validateUpgradeStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.UpgradeStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateUpgradeStatusEnum("upgradeStatus", "body", m.UpgradeStatus); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this connector based on context it is used
func (m *Connector) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Connector) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Connector) UnmarshalBinary(b []byte) error {
	var res Connector
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
