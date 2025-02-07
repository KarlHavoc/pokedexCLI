package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/KarlHavoc/pokedexCLI/internal/pokeapi"
)

type config struct {
	pokeAPIClient        pokeapi.Client
	nextLocationsURL     *string
	previousLocationsURL *string
}

func startREPL(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	areaName := ""
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		if len(words) == 2 {
			areaName = words[1]
		}
		commandName := words[0]
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, areaName)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	splitStrings := strings.Split(text, " ")
	lowered := []string{}
	for _, word := range splitStrings {
		if word == "" {
			continue
		}
		l := strings.ToLower(word)
		lowered = append(lowered, l)
	}
	return lowered

}
