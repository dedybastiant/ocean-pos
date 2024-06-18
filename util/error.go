package util

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) string {
	var message string
	for _, err := range err.(validator.ValidationErrors) {
		switch err.Tag() {
		case "required":
			message = fmt.Sprintf("%s is required", err.Field())
		case "email":
			message = fmt.Sprintf("%s must be a valid email address", err.Field())
		case "min":
			message = fmt.Sprintf("%s must be at least %s characters long", err.Field(), err.Param())
		default:
			message = fmt.Sprintf("%s is invalid", err.Field())
		}
	}
	return message
}
