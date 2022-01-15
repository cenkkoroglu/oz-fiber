package util

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

var validate *validator.Validate

func InitValidator() {
	validate = validator.New()

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" || name == "" {
			return ""
		}

		return "{" + name + "}"
	})
}

func ValidateStruct(data interface{}) []string {
	errors := make([]string, 0)
	if err := validate.Struct(data); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, validation := range validationErrors {
			errors = append(errors, validation.Error())
		}
	}
	return errors
}
