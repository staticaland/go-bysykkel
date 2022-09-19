package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"runtime/debug"
	"time"

	"github.com/rs/zerolog"
	"github.com/staticaland/go-bysykkel/internal/client"
	"github.com/staticaland/go-bysykkel/internal/models"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config   config
	stations *models.StationModel
	logger   zerolog.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	var logger zerolog.Logger

	if cfg.env == "development" {
		buildInfo, _ := debug.ReadBuildInfo()
		logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
			Level(zerolog.TraceLevel).
			With().
			Timestamp().
			Caller().
			Int("pid", os.Getpid()).
			Str("go_version", buildInfo.GoVersion).
			Logger()
	} else {
		logger = zerolog.New(os.Stdout).Level(zerolog.InfoLevel)
	}

	c := client.CreateClient()

	app := &application{
		config:   cfg,
		logger:   logger,
		stations: &models.StationModel{Client: c},
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	app.logger.Info().
		Str("env", cfg.env).
		Str("address", srv.Addr).
		Int("port", cfg.port).
		Msg("Starting server")

	err := srv.ListenAndServe()
	app.logger.Fatal().Err(err)
}
