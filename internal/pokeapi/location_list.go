package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"fmt"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseUrl + "/locatoin-area"

	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return RespShallowLocations{}, fmt.Errorf("bad status code:%v", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}
	
	return locationsResp, nil
	
}


