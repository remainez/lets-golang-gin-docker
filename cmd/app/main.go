package main

import (
	"log"

	"github.com/nagaihiroya/lets-golang-gin-docker/pkg/infrastructure"
	"github.com/nagaihiroya/lets-golang-gin-docker/pkg/usecase"
	"github.com/nagaihiroya/lets-golang-gin-docker/pkg/interfaces/handler"

	"github.com/gin-gonic/gin"
)

func main() {
    playlistRepository := infrastructure.NewPlaylistRepository()
    playlistUsecase := usecase.NewPlaylistUsecase(playlistRepository)
    playlistHandler := handler.NewPlaylistHandler(playlistUsecase)

	router := gin.Default()
	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "success",
		})
	})
	playlist := router.Group("/playlist")
	{
		playlist.GET("/", playlistHandler.Get)
	}

	if err := router.Run(":3000"); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}