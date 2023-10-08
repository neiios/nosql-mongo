package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Hotel struct {
	Name    string   `json:"name"`
	Address string   `json:"address"`
	Country string   `json:"country"`
	Rooms   []Room   `json:"rooms"`
	Workers []Worker `json:"workers"`
}

type Worker struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Position string `json:"position"`
}

type Room struct {
	Price  int  `json:"price"`
	Booked bool `json:"booked"`
}

func GetHotelsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, hotels)
}

func AddHotelHandler(c *gin.Context) {
	var hotel Hotel
	err := c.BindJSON(&hotel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hotels = append(hotels, hotel)
	c.JSON(http.StatusOK, hotel)
}

func main() {
	r := gin.Default()

	r.GET("/hotels", GetHotelsHandler)
	r.POST("/hotels", AddHotelHandler)

	r.Run()
}
