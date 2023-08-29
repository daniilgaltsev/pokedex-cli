package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


type command struct {
	help string
	function func() error
}
var commands map[string]command

type cliConfig struct {
	prev string
	next string
	mapStart string
	currentOffset int
}
var config cliConfig // NOTE: This should be made a non-global variable


func exit() error {
	os.Exit(0)
	return nil
}

func help() error {
	fmt.Println("Commands:")
	for key, value := range commands {
		fmt.Printf("  %s: %s\n", key, value.help)
	}
	return nil
}


func clean_input(s string) string {
	s = strings.ToLower(s)
	return s
}

func initCli() {
	config = cliConfig{
		prev: "",
		next: "",
		mapStart: "https://pokeapi.co/api/v2/location-area/",
	}
	commands = map[string]command{
		"exit": command{
			help: "Exit the program",
			function: exit,
		},
		"help": command{
			help: "Show this help message",
			function: help,
		},
		"map": command{
			help: "Show the next 20 locations",
			function: pokemap,
		},
		"mapb": command{
			help: "Show the previous 20 locations",
			function: pokemapb,
		},
	}

}

func cli() {
	initCli()

	scanner := bufio.NewScanner(os.Stdin)
	for ;; {
		fmt.Print("pokedex > ")
		scanner.Scan()
		text := clean_input(scanner.Text())

		recognized := false
		for key, value := range commands {
			if text == key {
				err := value.function()
				if err != nil {
					fmt.Printf("Error when executing %s: %s\n", text, err)
				}

				recognized = true
			}
		}
		if !recognized {
			fmt.Printf("Command `%s` not recognized.\n", text)
		}
	}
}
