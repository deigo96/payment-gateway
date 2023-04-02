package middleware

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func ValidateStruct(i interface{}) *string {
	validate = validator.New()

	err := validate.Struct(i)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil
		}
		errMessage := err.(validator.ValidationErrors)[0].StructField()
		return &errMessage
	}

	return nil
}
