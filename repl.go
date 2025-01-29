package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var comMap = map[string]cliCommand{
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
}

func startREPL() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		if len(userInput) == 0 {
			continue
		}
		cleanCommand := cleanInput(userInput)
		command, ok := comMap[cleanCommand[0]]
		if ok {
			err := command.callback()
			if cleanCommand[0] == "help" {
				for c := range comMap {
					fmt.Printf("%s: %s\n", c, comMap[c].description)
				}
			}
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
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

func commandExit() error {
	fmt.Println("Closing the Pokedex.....Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	return nil
}
