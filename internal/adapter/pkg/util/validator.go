package util

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func GetErrorValidationMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", fe.Field())
	case "email":
		return fmt.Sprintf("%s is not a valid email", fe.Field())
	case "max":
		return fmt.Sprintf("%s must be less than %s", fe.Field(), fe.Param())
	case "min":
		return fmt.Sprintf("%s must be more than %s", fe.Field(), fe.Param())
	case "number":
		return fmt.Sprintf("%s must be a number", fe.Field())
	case "e164":
		return fmt.Sprintf("%s must be a valid phone number", fe.Field())
	case "phone":
		return fmt.Sprintf("%s must be a valid phone number", fe.Field())
	default:
		return fmt.Sprintf("%s is not valid", fe.Field())
	}
}
