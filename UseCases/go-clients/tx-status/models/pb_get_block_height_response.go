// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PbGetBlockHeightResponse pb get block height response
//
// swagger:model pbGetBlockHeightResponse
type PbGetBlockHeightResponse struct {

	// block height
	BlockHeight *PbProcessedTick `json:"blockHeight,omitempty"`
}

// Validate validates this pb get block height response
func (m *PbGetBlockHeightResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlockHeight(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PbGetBlockHeightResponse) validateBlockHeight(formats strfmt.Registry) error {
	if swag.IsZero(m.BlockHeight) { // not required
		return nil
	}

	if m.BlockHeight != nil {
		if err := m.BlockHeight.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("blockHeight")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("blockHeight")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this pb get block height response based on the context it is used
func (m *PbGetBlockHeightResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBlockHeight(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PbGetBlockHeightResponse) contextValidateBlockHeight(ctx context.Context, formats strfmt.Registry) error {

	if m.BlockHeight != nil {

		if swag.IsZero(m.BlockHeight) { // not required
			return nil
		}

		if err := m.BlockHeight.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("blockHeight")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("blockHeight")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PbGetBlockHeightResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PbGetBlockHeightResponse) UnmarshalBinary(b []byte) error {
	var res PbGetBlockHeightResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
