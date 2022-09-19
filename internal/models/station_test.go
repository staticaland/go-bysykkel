package models

import (
	"testing"

	"github.com/staticaland/go-bysykkel/internal/assert"
	"github.com/staticaland/go-bysykkel/internal/client"
	"github.com/staticaland/go-bysykkel/internal/gbfs"
)

func TestJoinStationsByIdWithApiData(t *testing.T) {

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
	// 2351. This test is a bit silly, keeping it for entertainment purposes.
	expected := Station{
		StationID:         "2351",
		Name:              "Sogn Studentby",
		NumBikesAvailable: actual.NumBikesAvailable,
		NumDocksAvailable: actual.NumDocksAvailable,
	}

	assert.Equal(t, actual, expected)

}

func TestJoinStationsByIdWithMockData(t *testing.T) {

	stationInformation := gbfs.ApiStationInformation{
		Data: gbfs.StationInformationData{
			Stations: []gbfs.StationInformation{
				gbfs.StationInformation{
					StationID: "42",
					Name:      "Origo",
				},
			},
		},
	}

	stationStatus := gbfs.ApiStationStatus{
		Data: gbfs.StationStatusData{
			Stations: []gbfs.StationStatus{
				gbfs.StationStatus{
					StationID:         "42",
					NumBikesAvailable: 13,
					NumDocksAvailable: 37,
				},
			},
		},
	}

	stations, _ := joinStationsByID(&stationStatus, &stationInformation)

	actual := stations[0]

	expected := Station{
		StationID:         "42",
		Name:              "Origo",
		NumBikesAvailable: 13,
		NumDocksAvailable: 37,
	}

	assert.Equal(t, actual, expected)

}
