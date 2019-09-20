// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FingerPath FingerPath
// swagger:model FingerPath
type FingerPath struct {

	// blur
	// Required: true
	Blur *bool `json:"blur"`

	// board color
	// Required: true
	BoardColor *int32 `json:"boardColor"`

	// clear
	// Required: true
	Clear *bool `json:"clear"`

	// dash
	// Required: true
	Dash *bool `json:"dash"`

	// finger points
	// Required: true
	FingerPoints []*FingerPoint `json:"fingerPoints"`

	// path color
	// Required: true
	PathColor *int32 `json:"pathColor"`

	// path Id
	PathID int32 `json:"pathId,omitempty"`

	// stroke width
	// Required: true
	StrokeWidth *int32 `json:"strokeWidth"`

	// user Id
	UserID string `json:"userId,omitempty"`
}

// Validate validates this finger path
func (m *FingerPath) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlur(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBoardColor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClear(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFingerPoints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePathColor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStrokeWidth(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FingerPath) validateBlur(formats strfmt.Registry) error {

	if err := validate.Required("blur", "body", m.Blur); err != nil {
		return err
	}

	return nil
}

func (m *FingerPath) validateBoardColor(formats strfmt.Registry) error {

	if err := validate.Required("boardColor", "body", m.BoardColor); err != nil {
		return err
	}

	return nil
}

func (m *FingerPath) validateClear(formats strfmt.Registry) error {

	if err := validate.Required("clear", "body", m.Clear); err != nil {
		return err
	}

	return nil
}

func (m *FingerPath) validateDash(formats strfmt.Registry) error {

	if err := validate.Required("dash", "body", m.Dash); err != nil {
		return err
	}

	return nil
}

func (m *FingerPath) validateFingerPoints(formats strfmt.Registry) error {

	if err := validate.Required("fingerPoints", "body", m.FingerPoints); err != nil {
		return err
	}

	for i := 0; i < len(m.FingerPoints); i++ {
		if swag.IsZero(m.FingerPoints[i]) { // not required
			continue
		}

		if m.FingerPoints[i] != nil {
			if err := m.FingerPoints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fingerPoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FingerPath) validatePathColor(formats strfmt.Registry) error {

	if err := validate.Required("pathColor", "body", m.PathColor); err != nil {
		return err
	}

	return nil
}

func (m *FingerPath) validateStrokeWidth(formats strfmt.Registry) error {

	if err := validate.Required("strokeWidth", "body", m.StrokeWidth); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FingerPath) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FingerPath) UnmarshalBinary(b []byte) error {
	var res FingerPath
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
