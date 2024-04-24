package services

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"gochat/config/logger"
	"gochat/db/mongoDB/actions"
	"gochat/types/dto/request"
	"regexp"
)

var loggerStruct = logger.Logger{}
var Logger = loggerStruct.New("UserService")

type UserService struct {
	userActions actions.Actions
}

func _passwordValidator(fl validator.FieldLevel) bool {
	var passwordRegex = regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d]{8,}$`)
	return passwordRegex.MatchString(fl.Field().String())
}

func (s *UserService) CreateUser(body []byte) error {

	createUserRequest := &request.CreateUserRequest{}
	_ = json.Unmarshal(body, createUserRequest)

	err := createUserRequest.Validate()
	if err != nil {
		return err
	}

	userActions := actions.UserActions{}
	userActions.Init()

	err = userActions.Create(createUserRequest.Name, createUserRequest.Email, createUserRequest.Password)
	if err != nil {
		return err
	}
	return nil
}
