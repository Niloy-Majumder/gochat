package actions

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gochat/db/mongoDB"
)

type Actions interface {
	insert(entity interface{}) error
	findById(id string, result interface{}) error
	find(query interface{}, limit int, skip int, sort interface{}, result []interface{}) error
	findOneAndUpdate(query interface{}, entity interface{}) error
	delete(id string) error
}

type BaseActions struct {
	collectionName string
}

func (b *BaseActions) insert(entity interface{}) error {
	collection := mongoDB.Client.Database.Collection(b.collectionName)
	_, err := collection.InsertOne(context.TODO(), entity)
	if err != nil {
		return err
	}
	return nil
}

func (b *BaseActions) findById(id string, result interface{}) error {
	collection := mongoDB.Client.Database.Collection(b.collectionName)
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	if err := collection.FindOne(context.TODO(), _id).Decode(&result); err != nil {
		return err
	}
	return nil
}

func (b *BaseActions) find(query interface{}, limit int, skip int, sort interface{}, result []interface{}) (e error) {
	return nil
}
func (b *BaseActions) findOneAndUpdate(query interface{}, entity interface{}) error {
	return nil
}

func (b *BaseActions) delete(id string) error {
	return nil
}
