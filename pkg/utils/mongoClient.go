package utils

import (
	"context"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func GetDBCollection(col string) *mongo.Collection {
	return db.Collection(col)
}

func InitDB() error {
	print("dscvdsvcedvcedcved")
	uri := "mongodb://mongo:nU9LjIIaVuqib1sDpqCt@containers-us-west-164.railway.app:6266"
	if uri == "" {
		return errors.New("you must set your 'MONGODB_URI' environmental variable")
	}
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	db = client.Database(os.Getenv("dbName"))

	return nil
}

func CloseDB() error {
	return db.Client().Disconnect(context.Background())
}
