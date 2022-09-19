package main

import (
	"net/http"
)

func (app *application) showStationsHandler(w http.ResponseWriter, r *http.Request) {

	stations, _ := app.stations.All()

	err := app.writeJSON(w, http.StatusOK, envelope{"stations": stations}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	app.logger.Info().
		Int("count", len(stations)).
		Msg("Served stations")
}
