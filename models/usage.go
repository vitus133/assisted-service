// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Usage usage
//
// swagger:model usage
type Usage struct {

	// additional properties of the feature
	Data map[string]interface{} `json:"data,omitempty"`

	// Unique idenftifier of the feature
	ID string `json:"id,omitempty"`

	// name of the feature to track
	Name string `json:"name,omitempty"`
}

// Validate validates this usage
func (m *Usage) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Usage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Usage) UnmarshalBinary(b []byte) error {
	var res Usage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
