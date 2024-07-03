package models

type Contact struct {
	Name   string `json:"name" bson:"name"`
	UserId string `json:"userId" bson:"userId"`
	Email  string `json:"email" bson:"email"`
}

type User struct {
	Id       string    `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string    `json:"name" bson:"name"`
	Email    string    `json:"email" bson:"email"`
	Password string    `json:"password" bson:"password"`
	Contacts []Contact `json:"contacts" bson:"contacts"`
}
