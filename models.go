package main

type Chain struct {
	ID   int    `json:"id" binding:"required" bson:"_id"`
	Name string `json:"name" binding:"required"`
}

type Hotel struct {
	ID      int      `json:"id" binding:"required" bson:"_id"`
	ChainID int      `json:"chain_id" binding:"required"`
	Name    string   `json:"name" binding:"required"`
	Address string   `json:"address" binding:"required"`
	Country string   `json:"country" binding:"required"`
	Rooms   []Room   `json:"rooms" binding:"required,dive"`
	Workers []Worker `json:"workers" binding:"required,dive"`
}

type Worker struct {
	ID       int    `json:"id" binding:"required" bson:"_id"`
	Name     string `json:"name" binding:"required"`
	Age      int    `json:"age" binding:"required"`
	Position string `json:"position" binding:"required"`
}

type Room struct {
	ID     int  `json:"id" binding:"required" bson:"_id"`
	Number int  `json:"number" binding:"required"`
	Price  int  `json:"price" binding:"required"`
	Booked bool `json:"booked"`
}
