package main

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getHotel(id int) (Hotel, error) {
	var hotel Hotel
	filter := bson.D{{Key: "_id", Value: id}}
	err := hotelColl.FindOne(ctx, filter).Decode(&hotel)
	if err != nil {
		return Hotel{}, err
	}
	return hotel, nil
}

func getHotels() ([]Hotel, error) {
	hotels := make([]Hotel, 0)
	cursor, err := hotelColl.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &hotels); err != nil {
		return nil, err
	}
	return hotels, nil
}

func saveHotel(hotel Hotel) (Hotel, error) {
	_, err := hotelColl.InsertOne(ctx, hotel)
	if err != nil {
		return Hotel{}, err
	}
	return getHotel(hotel.ID)
}

func replaceHotel(id int, hotel Hotel) (Hotel, error) {
	var replacedHotel Hotel
	if id != hotel.ID {
		return Hotel{}, errors.New("ID in URL and ID in body are different")
	}
	filter := bson.D{{Key: "_id", Value: id}}
	opts := options.FindOneAndReplace().SetReturnDocument(options.After)
	err := hotelColl.FindOneAndReplace(ctx, filter, hotel, opts).Decode(&replacedHotel)
	if err != nil {
		return Hotel{}, err
	}
	return replacedHotel, nil
}

func deleteHotel(id int) error {
	filter := bson.D{{Key: "_id", Value: id}}
	_, err := hotelColl.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
