package api


// These are receiver functions that can access whatever is in the Client struct

// Could use receiver function to update fields
// Ineffective since the API returns everything
type Station struct {
	StationID         string `json:"station_id"`
	Name              string `json:"name"`
	NumBikesAvailable int    `json:"num_bikes_available"`
	NumDocksAvailable int    `json:"num_docks_available"`
}

type StationInformation struct {
	LastUpdated int    `json:"last_updated"`
	TTL         int    `json:"ttl"`
	Version     string `json:"version"`
	Data        struct {
		Stations []struct {
			StationID  string `json:"station_id"`
			Name       string `json:"name"`
			Address    string `json:"address"`
			RentalUris struct {
				Android string `json:"android"`
				Ios     string `json:"ios"`
			} `json:"rental_uris"`
			Lat      float64 `json:"lat"`
			Lon      float64 `json:"lon"`
			Capacity int     `json:"capacity"`
		} `json:"stations"`
	} `json:"data"`
}

type StationStatus struct {
	LastUpdated int    `json:"last_updated"`
	TTL         int    `json:"ttl"`
	Version     string `json:"version"`
	Data        struct {
		Stations []StatusStation `json:"stations"`
	} `json:"data"`
}

type StatusStation struct {
	StationID         string `json:"station_id"`
	IsInstalled       int    `json:"is_installed"`
	IsRenting         int    `json:"is_renting"`
	IsReturning       int    `json:"is_returning"`
	LastReported      int    `json:"last_reported"`
	NumBikesAvailable int    `json:"num_bikes_available"`
	NumDocksAvailable int    `json:"num_docks_available"`
}
