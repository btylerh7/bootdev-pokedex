package api

import (
	"fmt"
)

type PokemonInfo struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	}
}

func GetPokemon(pokemon string) (PokemonInfo, error) {
	apiUrl := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemon)
	return MakeAPICall[PokemonInfo](apiUrl)
}
