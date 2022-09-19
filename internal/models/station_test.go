package models

import (
	"testing"

	"github.com/staticaland/go-bysykkel/internal/assert"
	"github.com/staticaland/go-bysykkel/internal/client"
)

func TestJoinStationsByID(t *testing.T) {

	c := client.CreateClient()
	stationStatus, err := c.GetStationStatus()

	if err != nil {
		t.Fatal(err)
	}

	stationInformation, err := c.GetStationInformation()

	if err != nil {
		t.Fatal(err)
	}

	stations, _ := joinStationsByID(stationStatus, stationInformation)

	actual := stations[0]

	// Testing with the live API and assuming that the first element is always
	// 2351. This test is a bit silly, but I don't have time to refactor to
	// something better. I would probably create some mock gbfs structs by hand.
	expected := Station{
		StationID:         "2351",
		Name:              "Sogn Studentby",
		NumBikesAvailable: actual.NumBikesAvailable,
		NumDocksAvailable: actual.NumDocksAvailable,
	}

	assert.Equal(t, actual, expected)

}
