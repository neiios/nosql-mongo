package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hotel struct {
	Id      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name    string             `json:"name"`
	Address string             `json:"address"`
	Country string             `json:"country"`
	Rooms   []Room             `json:"rooms"`
	Workers []Worker           `json:"workers"`
}

type Worker struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Position string `json:"position"`
}

type Room struct {
	Number int  `json:"number"`
	Price  int  `json:"price"`
	Booked bool `json:"booked"`
}
