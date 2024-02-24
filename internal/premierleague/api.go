package premierleague

import (
	"encoding/json"
	"io"
	"net/http"
)

type APIClient struct {
	Client *http.Client
}

func NewClient() APIClient {
	client := &http.Client{}
	return APIClient{Client: client}
}

func (c *APIClient) GetFixtures() (Fixtures, error) {
	resp, err := c.Client.Get("https://fantasy.premierleague.com/api/fixtures/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var fixtures Fixtures
	err = json.Unmarshal(data, &fixtures)
	if err != nil {
		return nil, err
	}

	return fixtures, nil
}

func (c *APIClient) GetBootstrapData() (BootstrapData, error) {
	resp, err := c.Client.Get("https://fantasy.premierleague.com/api/bootstrap-static/")
	if err != nil {
		return BootstrapData{}, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return BootstrapData{}, err
	}

	var bootstrap BootstrapData
	err = json.Unmarshal(data, &bootstrap)
	if err != nil {
		return BootstrapData{}, err
	}
	return bootstrap, nil
}
