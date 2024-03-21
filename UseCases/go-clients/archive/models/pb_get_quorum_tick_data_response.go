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

// PbGetQuorumTickDataResponse pb get quorum tick data response
//
// swagger:model pbGetQuorumTickDataResponse
type PbGetQuorumTickDataResponse struct {

	// quorum tick data
	QuorumTickData *PbQuorumTickData `json:"quorumTickData,omitempty"`
}

// Validate validates this pb get quorum tick data response
func (m *PbGetQuorumTickDataResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQuorumTickData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PbGetQuorumTickDataResponse) validateQuorumTickData(formats strfmt.Registry) error {
	if swag.IsZero(m.QuorumTickData) { // not required
		return nil
	}

	if m.QuorumTickData != nil {
		if err := m.QuorumTickData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quorumTickData")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("quorumTickData")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this pb get quorum tick data response based on the context it is used
func (m *PbGetQuorumTickDataResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateQuorumTickData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PbGetQuorumTickDataResponse) contextValidateQuorumTickData(ctx context.Context, formats strfmt.Registry) error {

	if m.QuorumTickData != nil {

		if swag.IsZero(m.QuorumTickData) { // not required
			return nil
		}

		if err := m.QuorumTickData.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quorumTickData")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("quorumTickData")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PbGetQuorumTickDataResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PbGetQuorumTickDataResponse) UnmarshalBinary(b []byte) error {
	var res PbGetQuorumTickDataResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
