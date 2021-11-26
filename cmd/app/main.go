package main

import (
	"log"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	if err := router.Run(":3000"); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}