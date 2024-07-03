package actions

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	if userActions.exists(bson.M{"email": email}) {
		return errors.New("email already exists")
	}

	user := models.User{
		Name:     name,
		Email:    email,
		Password: password,
		Contacts: make([]models.Contact, 0),
	}

	if err := userActions.insert(user); err != nil {
		return err
	}

	return nil
}

func (userActions *UserActions) GetUserWithEmail(email string) (*models.User, error) {

	result, err := userActions.find(bson.D{{Key: "email", Value: email}}, 1, 0, nil)

	if err != nil || len(result) < 1 {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Email not found")
	}

	user := models.User{}
	bytes, err := bson.Marshal(result[0])
	if err != nil {
		return nil, err
	}
	err = bson.Unmarshal(bytes, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (userActions *UserActions) GetUserById(id string) (*models.User, error) {
	user := models.User{}
	err := userActions.findById(id, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (userActions *UserActions) UpdateContactsByUserDocument(user *models.User, contactName string, contactEmail string) (*models.User, error) {
	contactUser, err := userActions.GetUserWithEmail(contactEmail)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Contact not found")
	}

	contact := models.Contact{
		Name:   contactName,
		UserId: contactUser.Id,
		Email:  contactEmail,
	}
	user.Contacts = append(user.Contacts, contact)
	userId, _ := primitive.ObjectIDFromHex(user.Id)
	user.Id = ""
	if err := userActions.replace(bson.M{"_id": userId}, user); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return user, nil
}
