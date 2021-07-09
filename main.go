package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	// cards := addDeck()
	// cards.shuffle()
	// cards.print()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Conenction to server success!!",
		})
	})
	r.Run()
}
