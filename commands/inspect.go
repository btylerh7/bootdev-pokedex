package commands

import (
	"fmt"
	"github.com/btylerh7/bootdev-pokedex/internal/api"
	"slices"
)

func Inspect(cfg *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Must supply a pokemon name")
	}
	inspectPokemon := args[0]
	foundIndex := slices.IndexFunc(cfg.CaughtPokemon, func(pokemon api.PokemonInfo) bool { return pokemon.Name == inspectPokemon })
	if foundIndex == -1 {
		fmt.Println("you have not caught that pokemon")
	}
	foundPokemon := cfg.CaughtPokemon[foundIndex]
	stats := ""
	for _, stat := range foundPokemon.Stats {
		stats = stats + fmt.Sprintf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	types := ""
	for _, tp := range foundPokemon.Types {
		types = types + fmt.Sprintf("  -  %s\n", tp.Type.Name)
	}
	fmt.Printf(`
Name: %s
Height: %s
Weight: %s
Stats:
%s
Types:
%s`, foundPokemon.Name, foundPokemon.Height, foundPokemon.Weight, stats, types)
	return nil

}
