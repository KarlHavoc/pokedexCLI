package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userText := scanner.Text()
		cleanText := cleanInput(userText)
		fmt.Printf("Your command was: %s\n", cleanText[0])
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
