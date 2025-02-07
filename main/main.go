package main

import (
	"time"

	"github.com/KarlHavoc/pokedexCLI/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeAPIClient: pokeClient,
		caughtPokemon: map[string]pokeapi.Pokemon{},
	}
	startREPL(cfg)
}
