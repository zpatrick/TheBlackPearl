package main

import (
	"github.com/zpatrick/TheBlackPearl/api/controllers"
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

	routes := fireball.Decorate(
		movieController.Routes(),
		fireball.LogDecorator(),
	)

	app := fireball.NewApp(routes)
	log.Println("Running on port 80")
	log.Fatal(http.ListenAndServe(":80", app))

	log.Println(c.StringOr("aws.access_key", "none"))
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
