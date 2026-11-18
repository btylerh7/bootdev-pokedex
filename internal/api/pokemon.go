package api

import (
	"fmt"
)

type PokemonInfo struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
}

func GetPokemon(pokemon string) (PokemonInfo, error) {
	apiUrl := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemon)
	return MakeAPICall[PokemonInfo](apiUrl)
}
