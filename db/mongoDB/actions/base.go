package actions

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gochat/db/mongoDB"
)

type Actions interface {
	insert(entity interface{}) error
	findById(id string, result interface{}) error
	find(filter bson.D, limit int64, skip int64, sort bson.D) (result []interface{}, e error)
	findOneAndUpdate(query interface{}, entity interface{}) (result interface{}, e error)
	delete(id string) error
	exists(filter bson.M) bool
	replace(filter bson.M, entity interface{}) error
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

	if err := collection.FindOne(context.TODO(), bson.M{"_id": _id}).Decode(result); err != nil {
		return err
	}
	return nil
}

func (b *BaseActions) find(filter bson.D, limit int64, skip int64, sort bson.D) (result []interface{}, e error) {
	collection := mongoDB.Client.Database.Collection(b.collectionName)

	findOptions := &options.FindOptions{Limit: &limit, Skip: &skip}
	if sort != nil {
		findOptions.SetSort(sort)
	}

	cursor, err := collection.Find(context.Background(), filter, findOptions)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var res []interface{} = make([]interface{}, 1)

	err = cursor.All(context.TODO(), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (b *BaseActions) findOneAndUpdate(query interface{}, entity interface{}) (result interface{}, e error) {
	collection := mongoDB.Client.Database.Collection(b.collectionName)
	findOptions := options.FindOneAndUpdateOptions{}
	findOptions.SetReturnDocument(options.After)
	res := collection.FindOneAndUpdate(context.TODO(), query, entity, &findOptions)
	if res.Err() != nil {
		return nil, res.Err()
	}
	raw, _ := res.Raw()
	return raw, nil

}

func (b *BaseActions) delete(id string) error {
	collection := mongoDB.Client.Database.Collection(b.collectionName)
	deleteQuery := bson.D{{"_id", id}}
	_, err := collection.DeleteOne(context.TODO(), deleteQuery)
	if err != nil {
		return err
	}
	return nil
}

func (b *BaseActions) exists(filter bson.M) bool {
	collection := mongoDB.Client.Database.Collection(b.collectionName)
	result := collection.FindOne(context.Background(), filter)
	if result.Err() != nil {
		return false
	}

	return true
}

func (b *BaseActions) replace(filter bson.M, entity interface{}) error {
	collection := mongoDB.Client.Database.Collection(b.collectionName)
	update, err := collection.ReplaceOne(context.TODO(), filter, entity)
	if err != nil {
		return err
	}
	if update.ModifiedCount == 0 {
		return errors.New("document not updated")
	}
	return nil
}
