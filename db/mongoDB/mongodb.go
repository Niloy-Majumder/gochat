package mongoDB

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gochat/config/logger"
)

var (
	Client       = &MongoClient{}
	loggerStruct = logger.Logger{}
	Logger       = loggerStruct.New("Mongodb")
)

type MongoClient struct {
	Database *mongo.Database
	DbName   string
}

func (c *MongoClient) NewMongoClient(db string) *MongoClient {
	c.DbName = db
	return c
}

func (c *MongoClient) Connect(host string, port string) {
	connectionURI := fmt.Sprintf("mongodb://%s:%s", host, port)

	// Create the Connection pool - It runs in background doesn't throw error if client cannot connect
	mongoOptions := options.Client().ApplyURI(connectionURI)
	client, err := mongo.Connect(context.TODO(), mongoOptions)
	if err != nil {
		Logger.Fatal(err.Error())
	}

	// Check if the Connection is Successful
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		Logger.Fatal(err.Error())
	}

	Logger.Info("Connected to MongoDB on: ", host, ":", port)

	c.Database = client.Database(c.DbName)
}

func (c *MongoClient) SetIndexes(collection string, field ...string) error {
	var indexModel = make([]mongo.IndexModel, 0)

	for _, value := range field {
		indexModel = append(indexModel, mongo.IndexModel{
			Keys:    bson.D{{Key: value, Value: 1}},
			Options: options.Index().SetUnique(true),
		})
	}

	_, err := c.Database.Collection(collection).Indexes().CreateMany(context.Background(), indexModel)

	if err != nil {
		return err
	}
	return nil
}
