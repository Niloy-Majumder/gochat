package request

import (
	"github.com/go-playground/validator/v10"
	"unicode"
)

type CreateUserRequest struct {
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,password"`
}

type LoginRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,password"`
}

func _passwordValidator(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	// Check if the password contains at least one lowercase letter
	containsLower := false
	for _, char := range password {
		if unicode.IsLower(char) {
			containsLower = true
			break
		}
	}

	// Check if the password contains at least one uppercase letter
	containsUpper := false
	for _, char := range password {
		if unicode.IsUpper(char) {
			containsUpper = true
			break
		}
	}

	// Check if the password contains at least one digit
	containsDigit := false
	for _, char := range password {
		if unicode.IsDigit(char) {
			containsDigit = true
			break
		}
	}

	// Ensure password length is at least 8 characters
	isValidLength := len(password) >= 8

	// Return true if all conditions are met
	return containsLower && containsUpper && containsDigit && isValidLength
}

func (req *CreateUserRequest) Validate() error {

	customValidator := CustomValidator{tag: "password", fn: _passwordValidator}

	err := _validate(req, customValidator)
	if err != nil {
		return err
	}
	return nil
}

func (req *LoginRequest) Validate() error {
	customValidator := CustomValidator{tag: "password", fn: _passwordValidator}
	err := _validate(req, customValidator)
	if err != nil {
		return err
	}
	return nil
}
