package mongoDB

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gochat/config/logger"
)

var (
	Client       = &mongoClient{}
	loggerStruct = logger.Logger{}
	Logger       = loggerStruct.New("Mongodb")
)

type mongoClient struct {
	Database *mongo.Database
	DbName   string
}

func (c *mongoClient) NewMongoClient(db string) *mongoClient {
	c.DbName = db
	return c
}

func (c *mongoClient) Connect(host string, port string) {
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
