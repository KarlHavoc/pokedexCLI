package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c Client) GetLocation(pageURL *string) (PokemonInArea, error) {
	pokeList := PokemonInArea{}

	if val, ok := c.cache.Get(*pageURL); ok {
		err := json.Unmarshal(val, &pokeList)
		if err != nil {
			return PokemonInArea{}, err
		}
		fmt.Println()
		fmt.Println("Using the cache!")
		fmt.Println()
	}

	req, err := http.NewRequest("GET", *pageURL, nil)
	if err != nil {
		return PokemonInArea{}, err
	}

	resp, err := c.httpclient.Do(req)
	if err != nil {
		return PokemonInArea{}, err
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonInArea{}, err
	}
	c.cache.Add(*pageURL, dat)

	err = json.Unmarshal(dat, &pokeList)
	if err != nil {
		return PokemonInArea{}, err
	}
	return pokeList, nil
}
