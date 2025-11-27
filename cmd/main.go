package main

import (
	"fmt"
	"log"
	"os"

	"github.com/William-Dias-Marcos/Anime-API-Golang/controller"
	"github.com/William-Dias-Marcos/Anime-API-Golang/db"
	"github.com/William-Dias-Marcos/Anime-API-Golang/repository"
	"github.com/William-Dias-Marcos/Anime-API-Golang/usecase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){

	err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env: %v", err)
    }
	
	server := gin.Default()
	ginMode := os.Getenv("GIN_MODE")
	gin.SetMode(ginMode)

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
	server.POST("/animes", AnimeController.CreateAnime)
	server.DELETE("/animes/:guid", AnimeController.DeleteAnime)

	port := os.Getenv("PORT")
	server.Run(fmt.Sprintf(":%s", port))
}