package repository

import (
	"database/sql"
	"fmt"

	"github.com/William-Dias-Marcos/Anime-API-Golang/model"
)

type AnimeRepository struct {
	connection *sql.DB
}

func NewAnimeRepository(connection *sql.DB) AnimeRepository {
	return AnimeRepository{
		connection: connection,
	}
}

func (ar *AnimeRepository) GetAnimes() ([]model.Anime, error) {
	query:= "SELECT guid, title, description, genre, episodes, release_year, created_at, updated_at FROM animes"
	rows, err := ar.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Anime{}, err
	}
	defer rows.Close()

	animes := []model.Anime{}
	for rows.Next() {
		var anime model.Anime
		err := rows.Scan(&anime.GUID, &anime.Title, &anime.Description, &anime.Genre, &anime.Episodes, &anime.ReleaseYear, &anime.CreatedAt, &anime.UpdatedAt)
		if err != nil {
			fmt.Println(err)
			return []model.Anime{}, err
		}
		animes = append(animes, anime)
	}
	rows.Close()
	
	return animes, nil
}