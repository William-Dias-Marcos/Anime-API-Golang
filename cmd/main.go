package main

import (
	"github.com/William-Dias-Marcos/Anime-API-Golang/controller"
	"github.com/gin-gonic/gin"
)

func main(){
	server := gin.Default()

	animeController := controller.NewAnimeController()

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/animes", animeController.GetAnimes)

	server.Run(":8000")
}