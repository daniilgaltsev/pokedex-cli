package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/daniilgaltsev/pokedex-cli/internal/cache"
)

type command struct {
	help string
	function func(args []string) error
	nargs int
}
var commands map[string]command

type cliConfig struct {
	prev string
	next string
	mapStart string
	currentOffset int
}
var config cliConfig // NOTE: This should be made a non-global variable

var mapCache = cache.NewCache(20 * time.Second) // NOTE: This should be made a non-global variable
var exploreCache = cache.NewCache(20 * time.Second) // NOTE: This should _also_ be made a non-global variable


func exit(args []string) error {
	os.Exit(0)
	return nil
}

func help(args []string) error {
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
			nargs: 0,
		},
		"help": command{
			help: "Show this help message",
			function: help,
			nargs: 0,
		},
		"map": command{
			help: "Show the next 20 locations",
			function: pokemap,
			nargs: 0,
		},
		"mapb": command{
			help: "Show the previous 20 locations",
			function: pokemapb,
			nargs: 0,
		},
		"explore": command{
			help: "Explore a location",
			function: explore,
			nargs: 1,
		},
	}

}

func cli() {
	initCli()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for ;; {
		fmt.Print("pokedex > ")
		scanner.Scan()
		text := clean_input(scanner.Text())

		recognized := false
		for key, value := range commands {
			if text == key {
				args := make([]string, value.nargs)
				for i := 0; i < value.nargs; i++ {
					scanner.Scan()
					args[i] = scanner.Text()
				}

				err := value.function(args)
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
