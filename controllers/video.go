package controllers

import (
	"fmt"
	"strconv"

	"github.com/zpatrick/TheBlackPearl/video"
	"github.com/zpatrick/fireball"
)

type VideoController struct {
	store video.Store
}

func NewVideoController(s video.Store) *VideoController {
	return &VideoController{
		store: s,
	}
}

func (v *VideoController) Routes() []*fireball.Route {
	routes := []*fireball.Route{
		{
			Path: "/videos",
			Handlers: fireball.Handlers{
				"GET": v.listVideos,
			},
		},
		{
			Path: "/videos/:videoID",
			Handlers: fireball.Handlers{
				"GET": v.getVideo,
			},
		},
	}

	return routes
}

func (v *VideoController) listVideos(c *fireball.Context) (fireball.Response, error) {
	query := c.Request.URL.Query()

	reducers := []video.Reducer{}
	if search := query.Get("search"); search != "" {
		reducers = append(reducers, video.NewSearchReducer(search))
	}

	if start := query.Get("start"); start != "" {
		s, err := strconv.Atoi(start)
		if err != nil {
			return nil, fireball.NewError(400, fmt.Errorf("'start' must be an integer"), nil)
		}

		reducers = append(reducers, video.NewStartReducer(s))
	}

	if limit := query.Get("limit"); limit != "" {
		l, err := strconv.Atoi(limit)
		if err != nil {
			return nil, fireball.NewError(400, fmt.Errorf("'limit' must be an integer"), nil)
		}

		reducers = append(reducers, video.NewLimitReducer(l))
	}

	videos, err := v.store.ListVideos()
	if err != nil {
		return nil, err
	}

	for _, reduce := range reducers {
		videos = reduce(videos)
	}

	return c.HTML(200, "videos/list_videos.html", videos)
}

func (v *VideoController) getVideo(c *fireball.Context) (fireball.Response, error) {
	videoID := c.PathVariables["videoID"]
	video, err := v.store.GetVideo(videoID)
	if err != nil {
		// todo: capture this case in error decorator
		return nil, err
	}

	return c.HTML(200, "videos/get_video.html", video)
}
