// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixCreateClusterRequest openpitrix create cluster request
// swagger:model openpitrixCreateClusterRequest
type OpenpitrixCreateClusterRequest struct {

	// advanced param
	AdvancedParam []string `json:"advanced_param"`

	// app id
	AppID string `json:"app_id,omitempty"`

	// conf
	Conf string `json:"conf,omitempty"`

	// runtime env id
	RuntimeEnvID string `json:"runtime_env_id,omitempty"`

	// version id
	VersionID string `json:"version_id,omitempty"`
}

// Validate validates this openpitrix create cluster request
func (m *OpenpitrixCreateClusterRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdvancedParam(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixCreateClusterRequest) validateAdvancedParam(formats strfmt.Registry) error {

	if swag.IsZero(m.AdvancedParam) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixCreateClusterRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixCreateClusterRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixCreateClusterRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
