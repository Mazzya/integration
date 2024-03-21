// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PbTickData pb tick data
//
// swagger:model pbTickData
type PbTickData struct {

	// computor index
	ComputorIndex int64 `json:"computorIndex,omitempty"`

	// contract fees
	ContractFees []string `json:"contractFees"`

	// epoch
	Epoch int64 `json:"epoch,omitempty"`

	// signature hex
	SignatureHex string `json:"signatureHex,omitempty"`

	// tick number
	TickNumber int64 `json:"tickNumber,omitempty"`

	// time lock
	// Format: byte
	TimeLock strfmt.Base64 `json:"timeLock,omitempty"`

	// timestamp
	Timestamp string `json:"timestamp,omitempty"`

	// transaction ids
	TransactionIds []string `json:"transactionIds"`

	// var struct
	// Format: byte
	VarStruct strfmt.Base64 `json:"varStruct,omitempty"`
}

// Validate validates this pb tick data
func (m *PbTickData) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pb tick data based on context it is used
func (m *PbTickData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PbTickData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PbTickData) UnmarshalBinary(b []byte) error {
	var res PbTickData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
