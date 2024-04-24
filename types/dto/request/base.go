package request

import (
	"github.com/go-playground/validator/v10"
	validator2 "gochat/api/v1/services/validator"
)

type RequestDTO interface {
	Validate() error
}

type CustomValidator struct {
	tag string
	fn  validator.Func
}

func _validate(entity interface{}, customValidators ...CustomValidator) error {
	validate := validator.New()

	for _, validators := range customValidators {
		_ = validate.RegisterValidation(validators.tag, validators.fn)

	}

	err := validator2.Validator(validate, entity)
	if err != nil {
		return err
	}
	return nil
}
