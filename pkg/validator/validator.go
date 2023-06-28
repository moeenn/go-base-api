package validator

import (
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	FailedField string `json:"failed_field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

var validate = validator.New()

func ValidateStruct(data any) []*ErrorResponse {
	var errors []*ErrorResponse

	if err := validate.Struct(data); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}

	return errors
}
