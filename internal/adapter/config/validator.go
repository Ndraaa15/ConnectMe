package config

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func NewValidator() *validator.Validate {
	v := validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	if err := v.RegisterValidation("phone", phoneValidation); err != nil {
		log.Fatal().Err(err).Msg("Failed to register phone validation")
	}

	return v
}

func phoneValidation(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	if len(phone) < 10 || len(phone) > 15 {
		return false
	}

	if phone[:3] != "+62" {
		return false
	}

	return true
}
