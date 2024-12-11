package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	
	// Check if the response is in the cache
	cachedResp, ok := c.cache.Get(url)
	if ok {
		var resp RespShallowLocations
		err := json.Unmarshal(cachedResp, &resp)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return resp, nil
	}
	
	// if not in cache, make the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	// Add the response to the cache
	c.cache.Add(url, data)
	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
	

}