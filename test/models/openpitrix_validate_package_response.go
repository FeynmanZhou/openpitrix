// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixValidatePackageResponse openpitrix validate package response
// swagger:model openpitrixValidatePackageResponse
type OpenpitrixValidatePackageResponse struct {

	// error eg.[json error]
	Error string `json:"error,omitempty"`

	// filename map to detail
	ErrorDetails map[string]string `json:"error_details,omitempty"`

	// app name eg.[wordpress|mysql|...]
	Name string `json:"name,omitempty"`

	// app version name.eg.[0.1.0]
	VersionName string `json:"version_name,omitempty"`
}

// Validate validates this openpitrix validate package response
func (m *OpenpitrixValidatePackageResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixValidatePackageResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixValidatePackageResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixValidatePackageResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}