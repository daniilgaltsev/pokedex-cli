package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


type command struct {
	// name string
	help string
	function func() error
}
var commands map[string]command


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

func cli() {
	commands = map[string]command{
		"exit": command{
			// name: "exit",
			help: "Exit the program",
			function: exit,
		},
		"help": command{
			// name: "help",
			help: "Show this help message",
			function: help,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for ;; {
		fmt.Print("pokedex > ")
		scanner.Scan()
		text := clean_input(scanner.Text())

		recognized := false
		for key, value := range commands {
			if text == key {
				value.function()
				recognized = true
			}
		}
		if !recognized {
			fmt.Printf("Command `%s` not recognized.\n", text)
		}
	}
}
