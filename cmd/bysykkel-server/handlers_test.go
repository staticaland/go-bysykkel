package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rs/zerolog"

	"github.com/staticaland/go-bysykkel/internal/assert"
	"github.com/staticaland/go-bysykkel/internal/client"
	"github.com/staticaland/go-bysykkel/internal/models"
)

func TestHealthCheck(t *testing.T) {

	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	app := application{}

	app.healthCheckHandler(rr, r)

	rs := rr.Result()

	assert.Equal(t, rs.StatusCode, http.StatusOK)

}

func TestShowStations(t *testing.T) {

	logger := zerolog.New(io.Discard)

	c := client.CreateClient()

	app := &application{
		logger:   logger,
		stations: &models.StationModel{Client: c},
	}

	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	app.healthCheckHandler(rr, r)

	rs := rr.Result()

	assert.Equal(t, rs.StatusCode, http.StatusOK)

}
