package commands

import (
	"fmt"
	"math/rand"

	"github.com/btylerh7/bootdev-pokedex/internal/api"
)

func Catch(config *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Must supply a location name")
	}
	pokemon := args[0]
	result, err := api.GetPokemon(pokemon)
	if err != nil {
		return fmt.Errorf("Unable to find pokemon: %s", err.Error())
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	chance := rand.Intn(50 + result.BaseExperience)
	if result.BaseExperience < chance {
		fmt.Printf("%s was caught!\n", pokemon)
		config.CaughtPokemon = append(config.CaughtPokemon, result)
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
	}

	return nil
}
