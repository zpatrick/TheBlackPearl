package models

import (
	"sort"
)

type Movie struct {
	ID          string `data:"primary_key"`
	Title       string
	Description string
	PosterURL   string
}

type byTitle []*Movie

func (t byTitle) Len() int           { return len(t) }
func (t byTitle) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t byTitle) Less(i, j int) bool { return t[i].Title < t[j].Title }

func SortMovies(movies []*Movie) []*Movie {
	m := movies[:]
	sort.Sort(byTitle(m))
	return m
}
