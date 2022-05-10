// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Matcher Matcher models the matching of a label.
//
// swagger:model Matcher
type Matcher struct {

	// name
	Name string `json:"Name,omitempty"`

	// type
	Type MatchType `json:"Type,omitempty"`

	// value
	Value string `json:"Value,omitempty"`
}

// Validate validates this matcher
func (m *Matcher) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Matcher) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Type")
		}
		return err
	}

	return nil
}

// ContextValidate validate this matcher based on the context it is used
func (m *Matcher) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Matcher) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Type.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Matcher) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Matcher) UnmarshalBinary(b []byte) error {
	var res Matcher
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
