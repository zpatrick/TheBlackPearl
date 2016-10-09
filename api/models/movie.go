package models

type Movie struct {
	ID          string `data:"primary_key"`
	Title       string
	Description string
	Year        int
	PosterURL   string
}
