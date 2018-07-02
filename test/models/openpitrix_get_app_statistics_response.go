// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixGetAppStatisticsResponse openpitrix get app statistics response
// swagger:model openpitrixGetAppStatisticsResponse
type OpenpitrixGetAppStatisticsResponse struct {

	// app count
	AppCount int64 `json:"app_count,omitempty"`

	// last two week created
	LastTwoWeekCreated map[string]int64 `json:"last_two_week_created,omitempty"`

	// repo count
	RepoCount int64 `json:"repo_count,omitempty"`

	// top ten repos
	TopTenRepos map[string]int64 `json:"top_ten_repos,omitempty"`
}

// Validate validates this openpitrix get app statistics response
func (m *OpenpitrixGetAppStatisticsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixGetAppStatisticsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixGetAppStatisticsResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixGetAppStatisticsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
