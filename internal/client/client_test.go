package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/staticaland/go-bysykkel/internal/assert"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
	client *Client
)

func fixture(path string) string {
	b, err := ioutil.ReadFile("testdata/fixtures/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func setup() func() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client = CreateTestClient()

	return func() {
		server.Close()
	}
}

func TestGetStationInformation(t *testing.T) {

	teardown := setup()
	defer teardown()

	mux.HandleFunc("/station_information.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("station_information.json"))
	})

	stationInformation, err := client.GetStationInformation()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, stationInformation.Data.Stations[0].Name, "Sogn Studentby")

}

func TestGetStationStatus(t *testing.T) {

	teardown := setup()
	defer teardown()

	mux.HandleFunc("/station_status.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("station_status.json"))
	})

	stationStatus, err := client.GetStationStatus()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, stationStatus.Data.Stations[0].StationID, "2351")

}
