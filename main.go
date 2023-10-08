package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/hotels", func(c *gin.Context) {
		hotels, err := getAllHotels()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, hotels)
	})

	r.POST("/hotels", func(c *gin.Context) {
		var hotel Hotel
		err := c.BindJSON(&hotel)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		savedHotel, err := saveHotel(hotel)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, savedHotel)
	})

	r.Run()
}
