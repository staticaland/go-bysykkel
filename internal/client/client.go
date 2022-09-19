package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/staticaland/go-bysykkel/internal/gbfs"
)

const (
	BaseURL  = "https://gbfs.urbansharing.com/oslobysykkel.no"
	ClientID = "staticaland-go-bysykkel"
)

type Client struct {
	BaseURL    string
	ClientID   string
	HTTPClient *http.Client
}

type Option func(*Client) error

func SetBaseURL(baseURL string) Option {
	return func(c *Client) error {
		c.BaseURL = baseURL
		return nil
	}
}

func (c *Client) parseOptions(opts ...Option) error {
	for _, option := range opts {
		err := option(c)
		if err != nil {
			return err
		}
	}

	return nil
}

func CreateClient(opts ...Option) (*Client, error) {
	client := &Client{
		BaseURL: BaseURL,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}

	if err := client.parseOptions(opts...); err != nil {
		return nil, err
	}

	return client, nil
}

func (c *Client) GetStationInformation() (*gbfs.ApiStationInformation, error) {

	req, _ := http.NewRequest("GET", c.BaseURL+"/station_information.json", nil)

	res := gbfs.ApiStationInformation{}

	_, err := c.doRequest(req, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil

}

func (c *Client) GetStationStatus() (*gbfs.ApiStationStatus, error) {

	req, _ := http.NewRequest("GET", c.BaseURL+"/station_status.json", nil)

	res := gbfs.ApiStationStatus{}

	_, err := c.doRequest(req, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil

}

func (c *Client) doRequest(req *http.Request, v any) (any, error) {

	req.Header.Set("Client-Identifier", c.ClientID)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// I want to create custom success and error types, but this will do for now
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected response code: %v", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, v); err != nil {
		return nil, err
	}

	return v, nil
}
