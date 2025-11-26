package controller

import (
	"net/http"

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
