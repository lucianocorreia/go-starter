package services

import "github.com/go-playground/validator/v10"

// Validator provides validation mainly validating structs within the web context
type Validator struct {
	validator *validator.Validate
}

// NewValidator creates a new validator
func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

// Validate validates a struct
func (v *Validator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		return err
	}
	return nil
}
