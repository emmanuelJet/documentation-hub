// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Calldata calldata
// swagger:model Calldata
type Calldata struct {

	// calldata
	// Required: true
	Calldata EncodedByteArray `json:"calldata"`
}

// Validate validates this calldata
func (m *Calldata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCalldata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Calldata) validateCalldata(formats strfmt.Registry) error {

	if err := m.Calldata.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("calldata")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Calldata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Calldata) UnmarshalBinary(b []byte) error {
	var res Calldata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}