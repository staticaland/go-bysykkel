package client

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/staticaland/go-bysykkel/gbfs"
)

const (
	BaseURL               = "https://gbfs.urbansharing.com/oslobysykkel.no"
	ClientID              = "staticaland-go-bysykkel"
	StationInformationUrl = BaseURL + "/station_information.json"
	StationStatusUrl      = BaseURL + "/station_status.json"
)

type Client struct {
	BaseURL    string
	ClientID   string
	HTTPClient *http.Client
}

func CreateClient() *Client {
	return &Client{
		BaseURL: BaseURL,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) GetStationInformation() (*gbfs.ApiStationInformation, error) {

	req, _ := http.NewRequest("GET", StationInformationUrl, nil)

	req.Header.Set("Client-Identifier", c.ClientID)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("unexpected response code", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var stationInformation gbfs.ApiStationInformation

	if err := json.Unmarshal(body, &stationInformation); err != nil {
		log.Fatalln(err)
	}

	return &stationInformation, nil

}

func (c *Client) GetStationStatus() (*gbfs.ApiStationStatus, error) {

	req, _ := http.NewRequest("GET", StationStatusUrl, nil)

	req.Header.Set("Client-Identifier", c.ClientID)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("unexpected response code", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var stationStatus gbfs.ApiStationStatus

	if err := json.Unmarshal(body, &stationStatus); err != nil {
		log.Fatalln(err)
	}

	return &stationStatus, nil

}
