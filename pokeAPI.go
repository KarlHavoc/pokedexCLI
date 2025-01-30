package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type PokeAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func getAreas() (PokeAreas, error) {
	areaURL := "https://pokeapi.co/api/v2/location-area/"
	
	res, err := http.Get(areaURL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var areas PokeAreas

	if err := json.Unmarshal(body, &areas); err != nil {
		log.Fatal(err)
	}
	return areas, nil
}
