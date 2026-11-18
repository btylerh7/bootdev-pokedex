package main

import (
	"fmt"
	"github.com/btylerh7/bootdev-pokedex/internal/api"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Must supply a location name")
	}
	location := args[0]
	jsonData, err := api.GetPokemonForLocation(location)
	if err != nil {
		return err
	}
	for _, pokemon := range jsonData.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}
