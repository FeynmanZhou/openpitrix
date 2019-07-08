// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixDescribeAppVersionAuditsResponse openpitrix describe app version audits response
// swagger:model openpitrixDescribeAppVersionAuditsResponse
type OpenpitrixDescribeAppVersionAuditsResponse struct {

	// app version audit set
	AppVersionAuditSet OpenpitrixDescribeAppVersionAuditsResponseAppVersionAuditSet `json:"app_version_audit_set"`

	// total count of audits of app with specific version
	TotalCount int64 `json:"total_count,omitempty"`
}

// Validate validates this openpitrix describe app version audits response
func (m *OpenpitrixDescribeAppVersionAuditsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixDescribeAppVersionAuditsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixDescribeAppVersionAuditsResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixDescribeAppVersionAuditsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}