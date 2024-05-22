package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ExploreArea
func (c *Client) ExploreArea(areaName string) (RespLocationArea, error) {
	url := baseURL + "/location-area/" + areaName

	if val, ok := c.cache.Get(url); ok {
		locationsResp := RespLocationArea{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespLocationArea{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationArea{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationArea{}, err
	}

	locationsResp := RespLocationArea{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespLocationArea{}, err
	}

	c.cache.Add(url, dat)

	return locationsResp, nil
}
