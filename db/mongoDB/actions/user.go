package actions

import (
	"gochat/db/mongoDB/models"
)

type UserActions struct {
	*BaseActions
}

func (userActions *UserActions) Init() {
	userActions.BaseActions = &BaseActions{}
	userActions.collectionName = "users"
}

func (userActions *UserActions) Create(name string, email string, password string) error {
	user := models.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	if err := userActions.insert(user); err != nil {
		return err
	}
	return nil
}
