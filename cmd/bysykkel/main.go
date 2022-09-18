package main

import (
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"

	"github.com/staticaland/go-bysykkel/internal/client"
	"github.com/staticaland/go-bysykkel/internal/models"
)

func main() {

	c := client.CreateClient()

	stationModel := models.StationModel{
		Client: c,
	}

	stations, _ := stationModel.All()

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
