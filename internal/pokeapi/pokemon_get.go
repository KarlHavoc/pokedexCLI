package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c Client) GetPokemonInfo(pokemon string) (Pokemon, error) {
	fullURL := baseURL + "/pokemon/" + pokemon + "/"
	pokeData := Pokemon{}

	if val, ok := c.cache.Get(pokemon); ok {
		err := json.Unmarshal(val, &pokeData)
		if err != nil {
			return Pokemon{}, err
		}
		return pokeData, nil
	}
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpclient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(pokemon, dat)

	err = json.Unmarshal(dat, &pokeData)
	if err != nil {
		return Pokemon{}, err
	}
	return pokeData, nil
}
