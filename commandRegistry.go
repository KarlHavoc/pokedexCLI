package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
	*config
}

type config struct {
	next     string
	previous string
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
			callback:    commandMap,
			config: &config{
				next: "",
				previous: "",
			},
		},
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, c := range getCommands() {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	return nil
}

func commandMap() error {
	areas, err := getAreas()
	
	if err != nil {
		fmt.Printf("error getting areas: %v", err)
	}

	for i := range areas.Results {
		fmt.Println(areas.Results[i].Name)
	}

	return nil
}
