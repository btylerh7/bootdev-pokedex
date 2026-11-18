package commands

import "fmt"

func Pokedex(cfg *Config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.CaughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
