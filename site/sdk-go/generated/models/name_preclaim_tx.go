// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"

	utils "github.com/aeternity/aepp-sdk-go/utils"
)

// NamePreclaimTx name preclaim tx
// swagger:model NamePreclaimTx
type NamePreclaimTx struct {

	// account id
	// Required: true
	AccountID EncodedHash `json:"account_id"`

	// commitment id
	// Required: true
	CommitmentID EncodedHash `json:"commitment_id"`

	// fee
	// Required: true
	Fee utils.BigInt `json:"fee"`

	// nonce
	Nonce uint64 `json:"nonce,omitempty"`

	// ttl
	TTL uint64 `json:"ttl,omitempty"`
}

// Validate validates this name preclaim tx
func (m *NamePreclaimTx) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommitmentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFee(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NamePreclaimTx) validateAccountID(formats strfmt.Registry) error {

	if err := m.AccountID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("account_id")
		}
		return err
	}

	return nil
}

func (m *NamePreclaimTx) validateCommitmentID(formats strfmt.Registry) error {

	if err := m.CommitmentID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("commitment_id")
		}
		return err
	}

	return nil
}

func (m *NamePreclaimTx) validateFee(formats strfmt.Registry) error {

	if err := m.Fee.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("fee")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NamePreclaimTx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NamePreclaimTx) UnmarshalBinary(b []byte) error {
	var res NamePreclaimTx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
