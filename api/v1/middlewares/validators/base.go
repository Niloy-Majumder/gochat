package validators

import (
	"github.com/go-playground/validator/v10"
	"gochat/types/constants"
)

type XValidator struct {
	validator *validator.Validate
}

var validate = validator.New()

func (v XValidator) Validate(data interface{}) []constants.ErrorResponse {
	var validationErrors []constants.ErrorResponse

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem constants.ErrorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}
