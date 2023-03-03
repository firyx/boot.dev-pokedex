package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasList, error) {
	var fullURL string
	if pageURL != nil {
		fullURL = *pageURL
	} else {
		endpoint := "/location-area"
		fullURL = baseURL + endpoint
	}

	data, ok := c.cache.Get(fullURL)
	if ok {
		//cache hit
		locationAreasResp := LocationAreasList{}
		err := json.Unmarshal(data, &locationAreasResp)
		if err != nil {
			return LocationAreasList{}, err
		}
		return locationAreasResp, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasList{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasList{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasList{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasList{}, err
	}

	locationAreasResp := LocationAreasList{}
	err = json.Unmarshal(data, &locationAreasResp)
	if err != nil {
		return LocationAreasList{}, err
	}

	c.cache.Add(fullURL, data)

	return locationAreasResp, nil
}

func (c *Client) GetLocationArea(loactionAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + loactionAreaName
	fullURL := baseURL + endpoint

	data, ok := c.cache.Get(fullURL)
	if ok {
		//cache hit
		locationAreaResp := LocationArea{}
		err := json.Unmarshal(data, &locationAreaResp)
		if err != nil {
			return LocationArea{}, err
		}
		return locationAreaResp, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationAreaResp := LocationArea{}
	err = json.Unmarshal(data, &locationAreaResp)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(fullURL, data)

	return locationAreaResp, nil
}
