package utils

import (
	"github.com/go-playground/validator/v10"
)

// CustomValidator struct that embeds the validator library
type CustomValidator struct {
	Validator *validator.Validate
}

// Validate function implements the echo.Validator interface
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
