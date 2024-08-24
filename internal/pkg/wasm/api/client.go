package api

import (
	"encoding/json"
	"net/http"

	"homelabscm.com/scm/pkg/api_model"
)

type Client struct {
	apiEndpoint string
	httpClient  *http.Client
}

func NewAPIClient(apiEndpoint string) *Client {
	return &Client{
		apiEndpoint: apiEndpoint,
		httpClient:  &http.Client{},
	}
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.httpClient.Do(req)
}

func (c *Client) Status() (*api_model.StatusResponse, error) {
	req, err := http.NewRequest("GET", c.apiEndpoint+"/status", nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	
	status := &api_model.StatusResponse{}
	err = json.NewDecoder(resp.Body).Decode(status)
	if err != nil {
		return nil, err
	}
	return status, nil
}