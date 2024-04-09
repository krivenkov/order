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

// Pagination Pagination
//
// swagger:model Pagination
type Pagination struct {

	// limit of orders
	// Required: true
	// Minimum: 1
	Limit *float64 `json:"limit"`

	// offset
	// Required: true
	// Minimum: 1
	Offset *float64 `json:"offset"`

	// Total items
	// Required: true
	// Minimum: 0
	Total *float64 `json:"total"`
}

// Validate validates this pagination
func (m *Pagination) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOffset(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Pagination) validateLimit(formats strfmt.Registry) error {

	if err := validate.Required("limit", "body", m.Limit); err != nil {
		return err
	}

	if err := validate.Minimum("limit", "body", *m.Limit, 1, false); err != nil {
		return err
	}

	return nil
}

func (m *Pagination) validateOffset(formats strfmt.Registry) error {

	if err := validate.Required("offset", "body", m.Offset); err != nil {
		return err
	}

	if err := validate.Minimum("offset", "body", *m.Offset, 1, false); err != nil {
		return err
	}

	return nil
}

func (m *Pagination) validateTotal(formats strfmt.Registry) error {

	if err := validate.Required("total", "body", m.Total); err != nil {
		return err
	}

	if err := validate.Minimum("total", "body", *m.Total, 0, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this pagination based on context it is used
func (m *Pagination) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Pagination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Pagination) UnmarshalBinary(b []byte) error {
	var res Pagination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}