package main

import (
	"github.com/zpatrick/TheBlackPearl/api/controllers"
	"github.com/zpatrick/TheBlackPearl/api/models"
	"github.com/zpatrick/TheBlackPearl/api/stores"
	"github.com/zpatrick/fireball"
	"github.com/zpatrick/go-config"
	"github.com/zpatrick/go-sdata/container"
	"log"
	"net/http"
)

func main() {
	c := initConfig()
	movieStore := getMovieStore()
	movieController := controllers.NewMovieController(movieStore)

	movie := &models.Movie{
		ID:          "1",
		Title:       "Gladiator",
		Description: "string",
		PosterURL:   "http://google.com",
	}
	movieStore.Insert(movie).Execute()

	routes := fireball.Decorate(
		movieController.Routes(),
		fireball.LogDecorator(),
	)

	app := fireball.NewApp(routes)
	app.Before = enableCORS

	log.Println("Running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", app))

	log.Println(c.StringOr("aws.access_key", "none"))
}

func enableCORS(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token, Authorization")
}

func initConfig() *config.Config {
	ini := config.NewINIFile("config.ini")
	env := config.NewEnvironment(map[string]string{
		"AWS_ACCESS_KEY_ID":     "aws.access_key",
		"AWS_SECRET_ACCESS_KEY": "aws.secret_key",
		"AWS_BUCKET":            "aws.bucket",
	})

	c := config.NewConfig([]config.Provider{ini, env})
	if err := c.Load(); err != nil {
		log.Fatal(err)
	}

	return c
}

func getMovieStore() *stores.MovieStore {
	fileContainer := container.NewStringFileContainer("movies.json", nil)

	movieStore := stores.NewMovieStore(fileContainer)
	if err := movieStore.Init(); err != nil {
		log.Fatal(err)
	}

	return movieStore
}
