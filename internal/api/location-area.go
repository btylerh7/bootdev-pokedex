package api

import (
	"fmt"
)

type mapResponse struct {
	Count    int
	Next     *string
	Previous *string
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type locationResponse struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func GetLocationAreas(url *string) (mapResponse, error) {
	apiURL := "https://pokeapi.co/api/v2/location-area/"
	if url != nil {
		apiURL = *url
	}
	return MakeAPICall[mapResponse](apiURL)
}
func GetPokemonForLocation(location string) (locationResponse, error) {
	apiURL := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", location)
	return MakeAPICall[locationResponse](apiURL)
}
