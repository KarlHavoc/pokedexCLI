package main

import (
	"errors"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the name of 20 location areas in the Pokemon world",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location-area>",
			description: "Get the pokemon that exist in a particular area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "Try to catch the specified pokemon",
			callback:    commandCatch,
		},
	}
}

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, c := range getCommands() {
		fmt.Printf("\t%s:  %s\n", c.name, c.description)
		fmt.Println()
	}
	return nil
}

func commandMapf(cfg *config, args ...string) error {
	locationsResp, err := cfg.pokeAPIClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.previousLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.previousLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationsResp, err := cfg.pokeAPIClient.ListLocations(cfg.previousLocationsURL)
	if err != nil {
		return err
	}
	cfg.nextLocationsURL = locationsResp.Next
	cfg.previousLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) > 1 {
		return errors.New("please search for a correct area")
	}
	area := args[0]
	fullURL := "https://pokeapi.co/api/v2/location-area/" + area + "/"
	areaResp, err := cfg.pokeAPIClient.GetLocation(&fullURL)
	if err != nil {
		return errors.New("please enter a valid area to explore")
	}
	fmt.Println()
	fmt.Printf("Exploring %s\n", area)
	fmt.Println()
	for _, pokemon := range areaResp.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}


