package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreaResp, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	val, ok := c.cache.Get(url)
	if ok {
		locationAreaResp := LocationAreaResp{}
		err := json.Unmarshal(val, &locationAreaResp)
		if err != nil {
			return LocationAreaResp{}, err
		}
		return locationAreaResp, nil
	}


	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResp{}, err 
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResp{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}

	locationResp := LocationAreaResp{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return LocationAreaResp{}, err
	}
	c.cache.Add(url, dat)
	return locationResp, nil
}
