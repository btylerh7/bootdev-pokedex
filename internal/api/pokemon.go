package api

import (
	"fmt"
)

type getPokemonResponse struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
}

func GetPokemon(pokemon string) (getPokemonResponse, error) {
	apiUrl := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemon)
	return MakeAPICall[getPokemonResponse](apiUrl)
}
