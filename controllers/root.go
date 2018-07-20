package controllers

import (
	"github.com/zpatrick/TheBlackPearl/video"
	"github.com/zpatrick/fireball"
)

type RootController struct {
	store video.Store
}

func NewRootController(s video.Store) *RootController {
	return &RootController{
		store: s,
	}
}

func (r *RootController) Routes() []*fireball.Route {
	routes := []*fireball.Route{
		{
			Path: "/",
			Handlers: fireball.Handlers{
				"GET": r.getRoot,
			},
		},
	}

	return routes
}

func (r *RootController) getRoot(c *fireball.Context) (fireball.Response, error) {
	return c.HTML(200, "root.html", nil)
}
