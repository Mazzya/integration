// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PbGetTransferTransactionsPerTickResponse pb get transfer transactions per tick response
//
// swagger:model pbGetTransferTransactionsPerTickResponse
type PbGetTransferTransactionsPerTickResponse struct {

	// transfer transactions per tick
	TransferTransactionsPerTick []*PbTransferTransactionsPerTick `json:"transferTransactionsPerTick"`
}

// Validate validates this pb get transfer transactions per tick response
func (m *PbGetTransferTransactionsPerTickResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTransferTransactionsPerTick(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PbGetTransferTransactionsPerTickResponse) validateTransferTransactionsPerTick(formats strfmt.Registry) error {
	if swag.IsZero(m.TransferTransactionsPerTick) { // not required
		return nil
	}

	for i := 0; i < len(m.TransferTransactionsPerTick); i++ {
		if swag.IsZero(m.TransferTransactionsPerTick[i]) { // not required
			continue
		}

		if m.TransferTransactionsPerTick[i] != nil {
			if err := m.TransferTransactionsPerTick[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("transferTransactionsPerTick" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("transferTransactionsPerTick" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this pb get transfer transactions per tick response based on the context it is used
func (m *PbGetTransferTransactionsPerTickResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTransferTransactionsPerTick(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PbGetTransferTransactionsPerTickResponse) contextValidateTransferTransactionsPerTick(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TransferTransactionsPerTick); i++ {

		if m.TransferTransactionsPerTick[i] != nil {

			if swag.IsZero(m.TransferTransactionsPerTick[i]) { // not required
				return nil
			}

			if err := m.TransferTransactionsPerTick[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("transferTransactionsPerTick" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("transferTransactionsPerTick" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PbGetTransferTransactionsPerTickResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PbGetTransferTransactionsPerTickResponse) UnmarshalBinary(b []byte) error {
	var res PbGetTransferTransactionsPerTickResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
