package commands

import (
	"github.com/btylerh7/bootdev-pokedex/internal/api"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(cfg *Config, args ...string) error
}

type Config struct {
	Commands      map[string]CliCommand
	NextUrl       *string
	PrevUrl       *string
	CaughtPokemon []api.PokemonInfo
}
