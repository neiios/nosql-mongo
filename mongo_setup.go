package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getCleanMongoConnection() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://root:pass@localhost:27017")
	connection, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}
	// connection.Database("nosql-mongo-task").Drop(ctx)
	return connection
}

var ctx = context.TODO()
var connection = getCleanMongoConnection()
var hotelColl = connection.Database("nosql-mongo-task").Collection("hotels")
var chainColl = connection.Database("nosql-mongo-task").Collection("chains")
