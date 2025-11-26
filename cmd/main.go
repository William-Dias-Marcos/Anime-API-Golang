package main

import (
	"github.com/William-Dias-Marcos/Anime-API-Golang/controller"
	"github.com/William-Dias-Marcos/Anime-API-Golang/db"
	"github.com/William-Dias-Marcos/Anime-API-Golang/repository"
	"github.com/William-Dias-Marcos/Anime-API-Golang/usecase"
	"github.com/gin-gonic/gin"
)

func main(){
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic("Failed to connect to database, " + err.Error())
	}

	// repository
	AnimeRepository := repository.NewAnimeRepository(dbConnection)

	// usecase
	AnimeUseCase := usecase.NewAnimeUsecase(AnimeRepository)

	// controller
	AnimeController := controller.NewAnimeController(AnimeUseCase)

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/animes", AnimeController.GetAnimes)

	server.Run(":8000")
}