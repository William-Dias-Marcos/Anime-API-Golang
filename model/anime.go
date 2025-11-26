package model

import "time"

type Anime struct {
	GUID        string    `json:"guid" validate:"required"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Genre       string    `json:"genre" validate:"required"`
	Episodes    int       `json:"episodes" validate:"required"`
	ReleaseYear int       `json:"release_year" validate:"required"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
