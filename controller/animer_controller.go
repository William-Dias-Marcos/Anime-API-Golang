package controller

import (
	"net/http"

	"github.com/William-Dias-Marcos/Anime-API-Golang/model"
	"github.com/gin-gonic/gin"
)


type animeController struct {
	
}

func NewAnimeController() animeController {
	return animeController{}
}

func (p *animeController) GetAnimes(ctx *gin.Context) {
	animes := [] model.Anime{
		{
			GUID: "1",
			Title: "Naruto",
			Description: "A story about a young ninja.",
			Genre: "Action",
			Episodes: 220,
			ReleaseYear: 2002,
			CreatedAt: "2024-01-01T00:00:00Z",
			UpdatedAt: "2024-01-01T00:00:00Z",
		},
	}

	ctx.JSON(http.StatusOK, animes)
}
