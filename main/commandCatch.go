package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/KarlHavoc/pokedexCLI/internal/pokeapi"
)

type Pokedex struct {
	pokemon map[string]pokeapi.PokemonData
}

func commandCatch(cfg *config, args ...string) error {
	pokeDex := Pokedex{}
	if len(args) > 1 {
		return errors.New("you can catch only one")
	}
	pokeToCatch := args[0]

	pokeInfo, err := cfg.pokeAPIClient.GetPokemonInfo(pokeToCatch)
	if err != nil {
		return err
	}
	// pokeLevel := pokeInfo.BaseExperience
	// catchPoke := pokeLevel%10
	fmt.Printf("Trying to catch %s\n", pokeToCatch)
	randNum := rand.Intn(3)
	if randNum == 3 {
		fmt.Printf("%s was caught!!\n", pokeToCatch)
		pokeDex.pokemon[pokeToCatch] = pokeInfo
	} else {
		fmt.Printf("%s escaped!\n", pokeToCatch)
	}
	return nil
}
