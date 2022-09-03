package gbfs

type (
	ApiStationStatus struct {
		Metadata
		Data StationStatusData `json:"data"`
	}

	StationStatusData struct {
		Stations []StationStatus `json:"stations"`
	}

	StationStatus struct {
		StationID         string `json:"station_id"`
		IsInstalled       int    `json:"is_installed"`
		IsRenting         int    `json:"is_renting"`
		IsReturning       int    `json:"is_returning"`
		LastReported      int    `json:"last_reported"`
		NumBikesAvailable int    `json:"num_bikes_available"`
		NumDocksAvailable int    `json:"num_docks_available"`
	}
)
