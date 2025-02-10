package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c Client) ListLocations(pageURL *string) (PokeAreas, error) {
	url := baseURL + "/location-area"
	locationsResp := PokeAreas{}
	if pageURL != nil {
		url = *pageURL
	}
	if val, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return locationsResp, err
		}
		fmt.Println()
		fmt.Println("Using the cache!!")
		fmt.Println()
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeAreas{}, err
	}

	resp, err := c.httpclient.Do(req)
	if err != nil {
		return PokeAreas{}, err
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokeAreas{}, err
	}
	c.cache.Add(url, dat)

	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return PokeAreas{}, err
	}
	return locationsResp, nil
}
