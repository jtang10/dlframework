// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DlframeworkPredictorOpenRequest dlframework predictor open request
// swagger:model dlframeworkPredictorOpenRequest
type DlframeworkPredictorOpenRequest struct {

	// framework name
	FrameworkName string `json:"framework_name,omitempty"`

	// framework version
	FrameworkVersion string `json:"framework_version,omitempty"`

	// model name
	ModelName string `json:"model_name,omitempty"`

	// model version
	ModelVersion string `json:"model_version,omitempty"`

	// options
	Options *DlframeworkPredictionOptions `json:"options,omitempty"`

	// persist
	Persist bool `json:"persist,omitempty"`
}

// Validate validates this dlframework predictor open request
func (m *DlframeworkPredictorOpenRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DlframeworkPredictorOpenRequest) validateOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.Options) { // not required
		return nil
	}

	if m.Options != nil {
		if err := m.Options.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("options")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DlframeworkPredictorOpenRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DlframeworkPredictorOpenRequest) UnmarshalBinary(b []byte) error {
	var res DlframeworkPredictorOpenRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
