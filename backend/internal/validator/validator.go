package validator

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func New() *CustomValidator {
	v := validator.New()
	return &CustomValidator{Validator: v}

}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		return err
	}
	return nil
}

func GetValidationErrors(err error) map[string]string {
	errors := make(map[string]string)
	if valErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range valErrors {
			errors[strings.ToLower(fieldErr.Field())] = getErrorText(fieldErr)
		}
	}

	return errors
}

func getErrorText(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "this field is required"
	case "min":
		return "min lenght is " + err.Param()
	case "email":
		return "must be a valid email address"
	default:
		return "invalid"
	}
}
