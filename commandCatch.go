package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {

	if len(args) > 1 {
		return errors.New("you can catch only one")
	}
	pokeToCatch := args[0]

	pokeInfo, err := cfg.pokeAPIClient.GetPokemonInfo(pokeToCatch)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokeToCatch)
	pokeBaseEXP := pokeInfo.BaseExperience
	randNum := rand.Intn(pokeBaseEXP)
	if randNum > 40 {
		fmt.Printf("%s was caught!!\n", pokeToCatch)
		cfg.caughtPokemon[pokeToCatch] = pokeInfo
	} else {
		fmt.Printf("%s escaped!\n", pokeToCatch)
	}
	return nil
}
