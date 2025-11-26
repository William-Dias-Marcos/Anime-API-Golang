package usecase

import (
	"github.com/William-Dias-Marcos/Anime-API-Golang/model"
	"github.com/William-Dias-Marcos/Anime-API-Golang/repository"
)

type AnimeUsecase struct {
	repository repository.AnimeRepository
}

func NewAnimeUsecase(repo repository.AnimeRepository) AnimeUsecase {
	return AnimeUsecase{
		repository: repo,
	}
}

func (au * AnimeUsecase) GetAnimes() ([]model.Anime, error) {
	return au.repository.GetAnimes()
}