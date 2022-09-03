package main

import (
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"

	"github.com/staticaland/go-bysykkel/api"
	"github.com/staticaland/go-bysykkel/bysykkel"
)

func main() {

	c := bysykkel.CreateClient()

	stationInformation, _ := c.GetStationInformation()

	stationStatus, _ := c.GetStationStatus()

    stationStatusByID := make(map[string]api.StatusStation)

    for _, s := range stationStatus.Data.Stations {
        stationStatusByID[s.StationID] = s
    }

	var stations []api.Station

    for _, s := range stationInformation.Data.Stations {
		stations = append(stations, api.Station{
			StationID: s.StationID,
			Name: s.Name,
			NumBikesAvailable: stationStatusByID[s.StationID].NumBikesAvailable,
			NumDocksAvailable: stationStatusByID[s.StationID].NumDocksAvailable,
		})
    }

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Bikes", "Locks"})

	for _, s := range stations {
		table.Append([]string{
			s.Name,
			strconv.Itoa(s.NumBikesAvailable),
			strconv.Itoa(s.NumDocksAvailable),
		})
	}

	table.Render()

}
