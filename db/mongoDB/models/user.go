package models

type User struct {
	Id       string `json:"_id,omitempty" bson:"-"`
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}
