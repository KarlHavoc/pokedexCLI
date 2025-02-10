package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) > 1 {
		return errors.New("please enter the name of one pokemon")
	}
	name := args[0]
	val, exists := cfg.caughtPokemon[name]
	if !exists {
		return errors.New("you haven't caught that pokemon yet")
	}
	fmt.Printf("Name: %s\n", val.Name)
	fmt.Printf("Height: %v\n", val.Height)
	fmt.Printf("Weight: %v\n", val.Weight)
	fmt.Println("Stats:")
	for _, stat := range val.Stats {
		fmt.Printf("\t-%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range val.Types {
		fmt.Printf("\t-%v\n", t.Type.Name)
	}
	return nil
}
