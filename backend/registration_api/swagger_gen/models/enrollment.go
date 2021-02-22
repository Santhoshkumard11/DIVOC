// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Enrollment enrollment
//
// swagger:model enrollment
type Enrollment struct {

	// address
	Address *Address `json:"address,omitempty"`

	// appointment date
	// Format: date
	AppointmentDate strfmt.Date `json:"appointmentDate,omitempty"`

	// appointment slot
	AppointmentSlot string `json:"appointmentSlot,omitempty"`

	// beneficiary phone
	BeneficiaryPhone string `json:"beneficiaryPhone,omitempty"`

	// certified
	Certified *bool `json:"certified,omitempty"`

	// code
	Code string `json:"code,omitempty"`

	// dob
	// Required: true
	// Format: date
	Dob *strfmt.Date `json:"dob"`

	// email
	Email string `json:"email,omitempty"`

	// enrollment scope Id
	EnrollmentScopeID string `json:"enrollmentScopeId,omitempty"`

	// gender
	// Enum: [Male Female Other]
	Gender string `json:"gender,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// national Id
	// Required: true
	NationalID *string `json:"nationalId"`

	// phone
	Phone string `json:"phone,omitempty"`

	// program Id
	ProgramID string `json:"programId,omitempty"`

	// yob
	Yob int64 `json:"yob,omitempty"`
}

// Validate validates this enrollment
func (m *Enrollment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAppointmentDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDob(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGender(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNationalID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Enrollment) validateAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if m.Address != nil {
		if err := m.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("address")
			}
			return err
		}
	}

	return nil
}

func (m *Enrollment) validateAppointmentDate(formats strfmt.Registry) error {

	if swag.IsZero(m.AppointmentDate) { // not required
		return nil
	}

	if err := validate.FormatOf("appointmentDate", "body", "date", m.AppointmentDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Enrollment) validateDob(formats strfmt.Registry) error {

	if err := validate.Required("dob", "body", m.Dob); err != nil {
		return err
	}

	if err := validate.FormatOf("dob", "body", "date", m.Dob.String(), formats); err != nil {
		return err
	}

	return nil
}

var enrollmentTypeGenderPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Male","Female","Other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		enrollmentTypeGenderPropEnum = append(enrollmentTypeGenderPropEnum, v)
	}
}

const (

	// EnrollmentGenderMale captures enum value "Male"
	EnrollmentGenderMale string = "Male"

	// EnrollmentGenderFemale captures enum value "Female"
	EnrollmentGenderFemale string = "Female"

	// EnrollmentGenderOther captures enum value "Other"
	EnrollmentGenderOther string = "Other"
)

// prop value enum
func (m *Enrollment) validateGenderEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, enrollmentTypeGenderPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Enrollment) validateGender(formats strfmt.Registry) error {

	if swag.IsZero(m.Gender) { // not required
		return nil
	}

	// value enum
	if err := m.validateGenderEnum("gender", "body", m.Gender); err != nil {
		return err
	}

	return nil
}

func (m *Enrollment) validateNationalID(formats strfmt.Registry) error {

	if err := validate.Required("nationalId", "body", m.NationalID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Enrollment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Enrollment) UnmarshalBinary(b []byte) error {
	var res Enrollment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
