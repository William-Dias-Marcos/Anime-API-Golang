package controller

import (
	"net/http"

	"github.com/William-Dias-Marcos/Anime-API-Golang/model"
	"github.com/William-Dias-Marcos/Anime-API-Golang/usecase"
	"github.com/gin-gonic/gin"
)


type animeController struct {
	animeUseCase usecase.AnimeUsecase
}

func NewAnimeController(usecase usecase.AnimeUsecase) animeController {
	return animeController{
		animeUseCase: usecase,
	}
}

func (a *animeController) GetAnimes(ctx *gin.Context) {
	animes, err := a.animeUseCase.GetAnimes()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch animes"})
		return
	}

	ctx.JSON(http.StatusOK, animes)
}

func (a *animeController) CreateAnime(ctx *gin.Context) {
	var anime model.Anime
	err := ctx.ShouldBindJSON(&anime)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	guidCreatedAnime, err := a.animeUseCase.CretateAnime(anime)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create anime"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"guid": guidCreatedAnime})
	
}

func (a *animeController) DeleteAnime(ctx *gin.Context) {
	guid := ctx.Param("guid")

	err := a.animeUseCase.DeleteAnime(guid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete anime"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Anime deleted successfully"})
}