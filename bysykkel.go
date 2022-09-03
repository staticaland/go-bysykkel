package main

import (
	"log"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"

	"github.com/staticaland/go-bysykkel/client"
	"github.com/staticaland/go-bysykkel/gbfs"
)

type Station struct {
	StationID         string `json:"station_id"`
	Name              string `json:"name"`
	NumBikesAvailable int    `json:"num_bikes_available"`
	NumDocksAvailable int    `json:"num_docks_available"`
}

func main() {

	// Trigger CI
	c := client.CreateClient()

	stationInformation, err := c.GetStationInformation()

	if err != nil {
		log.Fatal(err)
	}

	stationStatus, err := c.GetStationStatus()

	if err != nil {
		log.Fatal(err)
	}

	stationStatusByID := make(map[string]gbfs.StationStatus)

	for _, s := range stationStatus.Data.Stations {
		stationStatusByID[s.StationID] = s
	}

	var stations []Station

	for _, s := range stationInformation.Data.Stations {
		stations = append(stations, Station{
			StationID:         s.StationID,
			Name:              s.Name,
			NumBikesAvailable: stationStatusByID[s.StationID].NumBikesAvailable,
			NumDocksAvailable: stationStatusByID[s.StationID].NumDocksAvailable,
		})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Bikes", "Locks"})
	table.SetAutoFormatHeaders(false)

	for _, s := range stations {
		table.Append([]string{
			s.Name,
			strconv.Itoa(s.NumBikesAvailable),
			strconv.Itoa(s.NumDocksAvailable),
		})
	}

	table.Render()

}
