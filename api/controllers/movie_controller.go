package controllers

import (
	"fmt"
	"github.com/zpatrick/TheBlackPearl/api/models"
	"github.com/zpatrick/TheBlackPearl/api/stores"
	"github.com/zpatrick/fireball"
)

type MovieController struct {
	Store *stores.MovieStore
}

func NewMovieController(store *stores.MovieStore) *MovieController {
	return &MovieController{
		Store: store,
	}
}

func (m *MovieController) Routes() []*fireball.Route {
	return []*fireball.Route{
		{
			Path: "/movies",
			Handlers: fireball.Handlers{
				"GET": m.ListMovies,
			},
		},
		{
			Path: "/movies/:id",
			Handlers: fireball.Handlers{
				"GET":    m.GetMovie,
				"DELETE": m.DeleteMovie,
			},
		},
	}
}

func (m *MovieController) ListMovies(c *fireball.Context) (fireball.Response, error) {
	movies, err := m.Store.SelectAll().Execute()
	if err != nil {
		return nil, err
	}

	return fireball.NewJSONResponse(200, movies)
}

func (m *MovieController) GetMovie(c *fireball.Context) (fireball.Response, error) {
	id := c.PathVariables["id"]

	movieIDMatch := func(m *models.Movie) bool {
		return m.ID == id
	}

	movie, err := m.Store.SelectAll().Where(movieIDMatch).FirstOrNil().Execute()
	if err != nil {
		return nil, err
	}

	if movie == nil {
		return nil, fmt.Errorf("Movie with id '%s' does not exist", id)
	}

	return fireball.NewJSONResponse(200, movie)
}

func (m *MovieController) DeleteMovie(c *fireball.Context) (fireball.Response, error) {
	id := c.PathVariables["id"]

	existed, err := m.Store.Delete(id).Execute()
	if err != nil {
		return nil, err
	}

	if !existed {
		return nil, fmt.Errorf("Movie with id '%s' does not exist", id)
	}

	return fireball.NewJSONResponse(200, nil)
}
