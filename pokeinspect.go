package main

import (
	"fmt"
)

func inspect(args []string) error {
	name := args[0]
	fmt.Printf("Inspecting %s...\n", name)

	pokemon, ok := config.pokemons[name]
	if !ok {
		fmt.Printf("%s not caught.\n", name)
		return nil
	}

	// print name, height, weight, stats, and types
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, type_ := range pokemon.Types {
		fmt.Printf(" - %s\n", type_.Type.Name)
	}

	return nil
}