package gbfs

type (
	ApiStationInformation struct {
		Metadata
		Data StationInformationData `json:"data"`
	}

	StationInformationData struct {
		Stations []StationInformation `json:"stations"`
	}

	StationInformation struct {
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
	}
)
