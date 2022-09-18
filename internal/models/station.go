package models

import (
	"github.com/staticaland/go-bysykkel/internal/client"
	"github.com/staticaland/go-bysykkel/internal/gbfs"
)

type Station struct {
	StationID         string `json:"station_id"`
	Name              string `json:"name"`
	NumBikesAvailable int    `json:"num_bikes_available"`
	NumDocksAvailable int    `json:"num_docks_available"`
}

type StationModel struct {
	Client *client.Client
}

func (s *StationModel) All() ([]Station, error) {

	stationInformation, err := s.Client.GetStationInformation()

	if err != nil {
		return nil, err
	}

	stationStatus, err := s.Client.GetStationStatus()

	if err != nil {
		return nil, err
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

	return stations, nil

}
