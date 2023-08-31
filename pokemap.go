package main

import (
	"errors"
	"fmt"
	"encoding/json"

	"github.com/daniilgaltsev/pokedex-cli/internal/pokeapi"
)

type mapResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func pokemapGetAndParse(url string) (mapResponse, error) {

	body, err := pokeapi.Request(url, mapCache)
	if err != nil {
		return mapResponse{}, err
	}

	locations := mapResponse{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return mapResponse{}, err
	}

	return locations, nil
}

func pokemap(args []string) error {
	var url string
	if config.next == "" {
		url = config.mapStart
	} else {
		url = config.next
	}
	
	locations, err := pokemapGetAndParse(url)
	if err != nil {
		return err
	}
	printLocations(locations, true)

	config.prev = url
	config.next = locations.Next

	return nil
}

func pokemapb(args []string) error {
	var url string
	if config.prev == "" {
		return errors.New("No previous locations")
	} else {
		url = config.prev
	}
	
	locations, err := pokemapGetAndParse(url)
	if err != nil {
		return err
	}
	printLocations(locations, false)

	config.prev = locations.Previous
	config.next = url

	return nil
}

func printLocations(locations mapResponse, isNext bool) {
	if isNext {
		for _, location := range locations.Results {
			config.currentOffset += 1
			fmt.Printf("%d: %s\n", config.currentOffset, location.Name)
		}
	} else {
		for i := len(locations.Results) - 1; i >= 0; i-- {
			fmt.Printf("%d: %s\n", config.currentOffset, locations.Results[i].Name)
			config.currentOffset -= 1
		}
	}
}
