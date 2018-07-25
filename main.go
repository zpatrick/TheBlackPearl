package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/quintilesims/d.ims.io/logging"
	"github.com/urfave/cli"
	"github.com/zpatrick/TheBlackPearl/controllers"
	"github.com/zpatrick/TheBlackPearl/video"
	"github.com/zpatrick/fireball"
)

/*
Design:
  Routes:
    /           -> Main screen that has big search bar, maybe some 'most-recent'
    /videos     -> Shows all videos in a table, alphabetical order
      ?search=string  -> Only return results that match the search
      ?limit=int      -> Only returns N results
      ?start=int      -> Only start after N results
    /videos/:videoID    -> Shows specific video, description, etc.. Has button for Watch

Bucket Layout:
  - Use a single folder (should probably pass in a param from cli, with 'videos/' as default)
  - The file name is the movie name, e.g. 'Inception.mp4'
  - For series, the file name is always <Series> - S<season #>E<episode #>, e.g. 'The Office - S1E1.mp4'
  - All other metadata, series, season, picture, etc., is recorded via tags

Cache:
  - Use a wrapper store that caches results
  - Cache expires every 30 minutes
  - Include a button in the website's header to refresh
*/

const (
	FlagPort         = "port"
	FlagDebug        = "debug"
	FlagUsername     = "username"
	FlagPassword     = "password"
	FlagAWSAccessKey = "aws-access-key"
	FlagAWSSecretKey = "aws-secret-key"
	FlagAWSRegion    = "aws-region"
	FlagS3Bucket     = "s3-bucket"
)

const (
	EnvVarPort         = "TBP_PORT"
	EnvVarDebug        = "TBP_DEBUG"
	EnvVarUsername     = "TBP_USERNAME"
	EnvVarPassword     = "TBP_PASSWORD"
	EnvVarAWSAccessKey = "TBP_AWS_ACCESS_KEY"
	EnvVarAWSSecretKey = "TBP_AWS_SECRET_KEY"
	EnvVarAWSRegion    = "TBP_AWS_REGION"
	EnvVarS3Bucket     = "TBP_S3_BUCKET"
)

func main() {
	app := cli.NewApp()
	app.Name = "The Black Pearl"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:   FlagPort,
			EnvVar: EnvVarPort,
			Value:  9090,
		},
		cli.BoolFlag{
			Name:   FlagDebug,
			EnvVar: EnvVarDebug,
		},
		cli.StringFlag{
			Name:   FlagUsername,
			EnvVar: EnvVarUsername,
		},
		cli.StringFlag{
			Name:   FlagPassword,
			EnvVar: EnvVarPassword,
		},
		cli.StringFlag{
			Name:   FlagAWSAccessKey,
			EnvVar: EnvVarAWSAccessKey,
		},
		cli.StringFlag{
			Name:   FlagAWSSecretKey,
			EnvVar: EnvVarAWSSecretKey,
		},
		cli.StringFlag{
			Name:   FlagAWSRegion,
			EnvVar: EnvVarAWSRegion,
			Value:  "us-west-2",
		},
		cli.StringFlag{
			Name:   FlagS3Bucket,
			EnvVar: EnvVarS3Bucket,
		},
	}

	app.Before = func(c *cli.Context) error {
		requiredFlags := []string{
			FlagUsername,
			FlagPassword,
			FlagAWSAccessKey,
			FlagAWSSecretKey,
			FlagS3Bucket,
		}

		for _, flag := range requiredFlags {
			if !c.IsSet(flag) {
				return fmt.Errorf("Required flag '%s' is not set", flag)
			}
		}

		debug := c.Bool(FlagDebug)
		log.SetOutput(logging.NewLogWriter(debug))

		return nil
	}

	app.Action = func(c *cli.Context) error {
		credentials := credentials.NewStaticCredentials(
			c.String(FlagAWSAccessKey),
			c.String(FlagAWSSecretKey),
			"")
		config := defaults.Get().Config.
			WithCredentials(credentials).
			WithRegion(c.String(FlagAWSRegion))
		session := session.New(config)

		store := video.NewS3Store(c.String(FlagS3Bucket), s3.New(session))
		rootController := controllers.NewRootController(store)
		videoController := controllers.NewVideoController(store)

		// todo: make custom error route handler that returns html w the error
		routes := rootController.Routes()
		routes = append(routes, videoController.Routes()...)
		routes = fireball.Decorate(routes,
			fireball.LogDecorator(),
			fireball.BasicAuthDecorator(c.String(FlagUsername), c.String(FlagPassword)))
		routes = fireball.EnableCORS(routes)

		fb := fireball.NewApp(routes)
		http.Handle("/", fb)

		fs := http.FileServer(http.Dir("static"))
		http.Handle("/static/", http.StripPrefix("/static", fs))

		port := fmt.Sprintf("0.0.0.0:%d", c.Int(FlagPort))
		log.Printf("[INFO] Running on port %s\n", port)
		return http.ListenAndServe(port, nil)
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
