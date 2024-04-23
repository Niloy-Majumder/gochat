package services

import "gochat/db/mongoDB/actions"

type UserService struct {
	userActions actions.Actions
}

func (s *UserService) CreateUser(name string, email string, password string) error {
	userActions := actions.UserActions{}
	userActions.Init()

	err := userActions.Create(name, email, password)
	if err != nil {
		return err
	}
	return nil
}
