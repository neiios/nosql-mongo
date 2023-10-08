package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getCleanMongoConnection() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://root:pass@localhost:27017")
	connection, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}
	// remove later
	connection.Database("nosql-mongo-task").Drop(ctx)
	return connection
}

var ctx = context.TODO()
var connection = getCleanMongoConnection()
var hotelColl = connection.Database("nosql-mongo-task").Collection("hotels")

func saveHotel(hotel Hotel) (*mongo.InsertOneResult, error) {
	return hotelColl.InsertOne(ctx, hotel)
}

func getAllHotels() ([]Hotel, error) {
	var hotels []Hotel
	cursor, err := hotelColl.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &hotels); err != nil {
		return nil, err
	}
	return hotels, nil
}
