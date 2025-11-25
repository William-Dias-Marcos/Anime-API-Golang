package model

type Anime struct {
	GUID 	     string  `json:"guid"`
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	Genre        string  `json:"genre"`
	Episodes     int     `json:"episodes"`
	ReleaseYear  int     `json:"release_year"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}
	  