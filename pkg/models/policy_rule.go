// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PolicyRule policy rule
//
// swagger:model PolicyRule
type PolicyRule struct {

	// action
	// Enum: [ALLOW DENY LOG RE_AUTH NEVER BYPASS INTERCEPT NO_DOWNLOAD BYPASS_RE_AUTH INTERCEPT_ACCESSIBLE ISOLATE BYPASS_ISOLATE INSPECT BYPASS_INSPECT]
	Action string `json:"action,omitempty"`

	// action Id
	ActionID int64 `json:"actionId,omitempty"`

	// app connector groups
	AppConnectorGroups []*AppConnectorGroup `json:"appConnectorGroups"`

	// app server groups
	AppServerGroups []*AppServerGroup `json:"appServerGroups"`

	// bypass default rule
	BypassDefaultRule bool `json:"bypassDefaultRule,omitempty"`

	// conditions
	Conditions []*ConditionSet `json:"conditions"`

	// creation time
	CreationTime int32 `json:"creationTime,omitempty"`

	// custom msg
	CustomMsg string `json:"customMsg,omitempty"`

	// default rule
	DefaultRule bool `json:"defaultRule,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// isolation default rule
	IsolationDefaultRule bool `json:"isolationDefaultRule,omitempty"`

	// modified by
	ModifiedBy int64 `json:"modifiedBy,omitempty"`

	// modified time
	ModifiedTime int32 `json:"modifiedTime,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// operator
	// Enum: [AND OR]
	Operator string `json:"operator,omitempty"`

	// policy set Id
	PolicySetID int64 `json:"policySetId,omitempty"`

	// policy type
	PolicyType int32 `json:"policyType,omitempty"`

	// priority
	Priority int32 `json:"priority,omitempty"`

	// reauth default rule
	ReauthDefaultRule bool `json:"reauthDefaultRule,omitempty"`

	// reauth idle timeout
	ReauthIdleTimeout int32 `json:"reauthIdleTimeout,omitempty"`

	// reauth timeout
	ReauthTimeout int32 `json:"reauthTimeout,omitempty"`

	// rule order
	RuleOrder int32 `json:"ruleOrder,omitempty"`

	// siem default rule
	SiemDefaultRule bool `json:"siemDefaultRule,omitempty"`

	// zpn cbi profile Id
	ZpnCbiProfileID int64 `json:"zpnCbiProfileId,omitempty"`

	// zpn inspection profile Id
	ZpnInspectionProfileID int64 `json:"zpnInspectionProfileId,omitempty"`

	// zpn inspection profile name
	ZpnInspectionProfileName string `json:"zpnInspectionProfileName,omitempty"`
}

// Validate validates this policy rule
func (m *PolicyRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAppConnectorGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAppServerGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperator(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var policyRuleTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ALLOW","DENY","LOG","RE_AUTH","NEVER","BYPASS","INTERCEPT","NO_DOWNLOAD","BYPASS_RE_AUTH","INTERCEPT_ACCESSIBLE","ISOLATE","BYPASS_ISOLATE","INSPECT","BYPASS_INSPECT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		policyRuleTypeActionPropEnum = append(policyRuleTypeActionPropEnum, v)
	}
}

const (

	// PolicyRuleActionALLOW captures enum value "ALLOW"
	PolicyRuleActionALLOW string = "ALLOW"

	// PolicyRuleActionDENY captures enum value "DENY"
	PolicyRuleActionDENY string = "DENY"

	// PolicyRuleActionLOG captures enum value "LOG"
	PolicyRuleActionLOG string = "LOG"

	// PolicyRuleActionREAUTH captures enum value "RE_AUTH"
	PolicyRuleActionREAUTH string = "RE_AUTH"

	// PolicyRuleActionNEVER captures enum value "NEVER"
	PolicyRuleActionNEVER string = "NEVER"

	// PolicyRuleActionBYPASS captures enum value "BYPASS"
	PolicyRuleActionBYPASS string = "BYPASS"

	// PolicyRuleActionINTERCEPT captures enum value "INTERCEPT"
	PolicyRuleActionINTERCEPT string = "INTERCEPT"

	// PolicyRuleActionNODOWNLOAD captures enum value "NO_DOWNLOAD"
	PolicyRuleActionNODOWNLOAD string = "NO_DOWNLOAD"

	// PolicyRuleActionBYPASSREAUTH captures enum value "BYPASS_RE_AUTH"
	PolicyRuleActionBYPASSREAUTH string = "BYPASS_RE_AUTH"

	// PolicyRuleActionINTERCEPTACCESSIBLE captures enum value "INTERCEPT_ACCESSIBLE"
	PolicyRuleActionINTERCEPTACCESSIBLE string = "INTERCEPT_ACCESSIBLE"

	// PolicyRuleActionISOLATE captures enum value "ISOLATE"
	PolicyRuleActionISOLATE string = "ISOLATE"

	// PolicyRuleActionBYPASSISOLATE captures enum value "BYPASS_ISOLATE"
	PolicyRuleActionBYPASSISOLATE string = "BYPASS_ISOLATE"

	// PolicyRuleActionINSPECT captures enum value "INSPECT"
	PolicyRuleActionINSPECT string = "INSPECT"

	// PolicyRuleActionBYPASSINSPECT captures enum value "BYPASS_INSPECT"
	PolicyRuleActionBYPASSINSPECT string = "BYPASS_INSPECT"
)

// prop value enum
func (m *PolicyRule) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, policyRuleTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PolicyRule) validateAction(formats strfmt.Registry) error {
	if swag.IsZero(m.Action) { // not required
		return nil
	}

	// value enum
	if err := m.validateActionEnum("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

func (m *PolicyRule) validateAppConnectorGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.AppConnectorGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.AppConnectorGroups); i++ {
		if swag.IsZero(m.AppConnectorGroups[i]) { // not required
			continue
		}

		if m.AppConnectorGroups[i] != nil {
			if err := m.AppConnectorGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("appConnectorGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("appConnectorGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PolicyRule) validateAppServerGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.AppServerGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.AppServerGroups); i++ {
		if swag.IsZero(m.AppServerGroups[i]) { // not required
			continue
		}

		if m.AppServerGroups[i] != nil {
			if err := m.AppServerGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("appServerGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("appServerGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PolicyRule) validateConditions(formats strfmt.Registry) error {
	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	for i := 0; i < len(m.Conditions); i++ {
		if swag.IsZero(m.Conditions[i]) { // not required
			continue
		}

		if m.Conditions[i] != nil {
			if err := m.Conditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PolicyRule) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var policyRuleTypeOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AND","OR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		policyRuleTypeOperatorPropEnum = append(policyRuleTypeOperatorPropEnum, v)
	}
}

const (

	// PolicyRuleOperatorAND captures enum value "AND"
	PolicyRuleOperatorAND string = "AND"

	// PolicyRuleOperatorOR captures enum value "OR"
	PolicyRuleOperatorOR string = "OR"
)

// prop value enum
func (m *PolicyRule) validateOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, policyRuleTypeOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PolicyRule) validateOperator(formats strfmt.Registry) error {
	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperatorEnum("operator", "body", m.Operator); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this policy rule based on the context it is used
func (m *PolicyRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppConnectorGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAppServerGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PolicyRule) contextValidateAppConnectorGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AppConnectorGroups); i++ {

		if m.AppConnectorGroups[i] != nil {
			if err := m.AppConnectorGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("appConnectorGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("appConnectorGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PolicyRule) contextValidateAppServerGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AppServerGroups); i++ {

		if m.AppServerGroups[i] != nil {
			if err := m.AppServerGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("appServerGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("appServerGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PolicyRule) contextValidateConditions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Conditions); i++ {

		if m.Conditions[i] != nil {
			if err := m.Conditions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PolicyRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyRule) UnmarshalBinary(b []byte) error {
	var res PolicyRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}