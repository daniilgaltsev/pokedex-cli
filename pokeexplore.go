package main

import (
	"encoding/json"
	"fmt"

	"github.com/daniilgaltsev/pokedex-cli/internal/pokeapi"
)

type ExploreResponse struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}


func explore(args []string) error {
	name := args[0]
	fmt.Printf("Exploring %s...\n", name)

	const baseUrl = "https://pokeapi.co/api/v2/location-area/"
	url := baseUrl + name
	body, err := pokeapi.Request(url, exploreCache)
	if err != nil {
		return err
	}

	exploreResponse := ExploreResponse{}
	err = json.Unmarshal(body, &exploreResponse)
	if err != nil {
		return err
	}

	fancyName := exploreResponse.Names[0].Name
	pokemons := exploreResponse.PokemonEncounters
	fmt.Printf("Found %d pokemons in %s\n", len(pokemons), fancyName)
	for _, pokemon := range pokemons {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
