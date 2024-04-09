package utils

import (
	"github.com/go-playground/validator/v10"
)

// Use a single instance of Validate, it caches struct info
var valid *validator.Validate

func init() {
	valid = validator.New()
}

// Validate struct fields
func ValidateStruct(s interface{}) error {
	return valid.Struct(s)
}
