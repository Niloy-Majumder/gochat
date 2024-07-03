package services

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"gochat/config/logger"
	"gochat/db/mongoDB/actions"
	"gochat/db/mongoDB/models"
	"gochat/types/dto/request"
	"gochat/utils"
	"golang.org/x/crypto/bcrypt"
	"sync"
)

var loggerStruct = logger.Logger{}
var Logger = loggerStruct.New("UserService")

type UserService struct {
	mutex       sync.Mutex
	userActions actions.UserActions
}

func (s *UserService) CreateUser(body []byte) error {

	createUserRequest := &request.CreateUserRequest{}
	_ = json.Unmarshal(body, createUserRequest)

	err := createUserRequest.Validate()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(createUserRequest.Password), bcrypt.DefaultCost)

	createUserRequest.Password = string(password)

	s.userActions = actions.UserActions{}
	s.userActions.Init()

	s.mutex.Lock()
	err = s.userActions.Create(createUserRequest.Name, createUserRequest.Email, createUserRequest.Password)
	s.mutex.Unlock()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return nil
}

func (s *UserService) LoginUser(body []byte) (string, error) {
	loginRequest := &request.LoginRequest{}
	_ = json.Unmarshal(body, loginRequest)
	err := loginRequest.Validate()
	if err != nil {
		return "", fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	s.userActions = actions.UserActions{}
	s.userActions.Init()

	user, err := s.userActions.GetUserWithEmail(loginRequest.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		return "", fiber.NewError(fiber.StatusBadRequest, "invalid password")
	}

	token, err := utils.CreateJWTToken(user.Id)
	if err != nil {
		return "", fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return token, nil
}

func (s *UserService) GetUserById(id string) (*models.User, error) {
	s.userActions = actions.UserActions{}
	s.userActions.Init()
	user, err := s.userActions.GetUserById(id)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return user, nil
}

func (s *UserService) UpdateContacts(id string, contactName string, contactEmail string) (*models.User, error) {
	s.userActions = actions.UserActions{}
	s.userActions.Init()
	user, err := s.userActions.GetUserById(id)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	for _, contact := range user.Contacts {
		if contact.Email == contactEmail {
			return nil, fiber.NewError(fiber.StatusBadRequest, "Email Already Saved")
		} else if contact.Name == contactName {
			return nil, fiber.NewError(fiber.StatusBadRequest, "Contact Name Already Exists")
		}
	}
	updatedUser, err := s.userActions.UpdateContactsByUserDocument(user, contactName, contactEmail)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return updatedUser, nil
}
