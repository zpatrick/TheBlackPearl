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

// The css/html: https://www.w3schools.com/w3css/tryw3css_templates_portfolio.htm
// https://www.w3schools.com/w3css/tryit.asp?filename=tryw3css_templates_portfolio&stacked=h
func (r *RootController) getRoot(c *fireball.Context) (fireball.Response, error) {
	return c.HTML(200, "root.html", nil)
}
