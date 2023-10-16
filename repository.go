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

func getChain(id int) (Chain, error) {
	var chain Chain
	filter := bson.D{{Key: "_id", Value: id}}
	err := chainColl.FindOne(ctx, filter).Decode(&chain)
	if err != nil {
		return Chain{}, err
	}
	return chain, nil
}

func saveChain(chain Chain) (Chain, error) {
	_, err := chainColl.InsertOne(ctx, chain)
	if err != nil {
		return Chain{}, err
	}
	return getChain(chain.ID)
}

// get nested documents

func getRooms(id int) ([]Room, error) {
	var hotel Hotel
	filter := bson.D{{Key: "_id", Value: id}}
	err := hotelColl.FindOne(ctx, filter).Decode(&hotel)
	if err != nil {
		return nil, err
	}
	return hotel.Rooms, nil
}

func getWorkers(id int) ([]Worker, error) {
	var hotel Hotel
	filter := bson.D{{Key: "_id", Value: id}}
	err := hotelColl.FindOne(ctx, filter).Decode(&hotel)
	if err != nil {
		return nil, err
	}
	return hotel.Workers, nil
}

// aggregation pipelines

func getWorkersByChain(chainID int) ([]Worker, error) {
	workers := make([]Worker, 0)

	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.D{{Key: "chainid", Value: chainID}}}},
		{{Key: "$unwind", Value: "$workers"}},
		{{Key: "$replaceRoot", Value: bson.D{{Key: "newRoot", Value: "$workers"}}}},
	}

	cursor, err := hotelColl.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var worker Worker
		err := cursor.Decode(&worker)
		if err != nil {
			return nil, err
		}
		workers = append(workers, worker)
	}

	return workers, nil
}

func countWorkersByPosition(hotelId int) (map[string]int32, error) {
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.D{{Key: "_id", Value: hotelId}}}},
		{{Key: "$unwind", Value: "$workers"}},
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$workers.position"},
			{Key: "count", Value: bson.D{{Key: "$sum", Value: 1}}},
		}}},
	}

	cursor, err := hotelColl.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	var positions []bson.M
	if err = cursor.All(ctx, &positions); err != nil {
		return nil, err
	}

	result := make(map[string]int32)
	for _, position := range positions {
		name := position["_id"].(string)
		count := position["count"].(int32)
		result[name] = count
	}

	return result, nil
}
