package api

import (
	"slices"
	"testing"
)

type testCase struct {
	input  string
	output []string
	err    error
}

func TestGetPokemonForLocation(t *testing.T) {
	locations := []testCase{
		{
			input:  "sinnoh-pokemon-league-area",
			output: []string{"tentacruel", "magikarp", "gyarados", "remoraid", "octillery", "wingull", "pelipper", "luvdisc"},
			err:    nil,
		},
	}
	for _, location := range locations {
		result, err := GetPokemonForLocation(location.input)
		if err != location.err {
			t.Errorf("Did not get expected value of %s, got %s", location.err, err)
		}
		resString := make([]string, 0, len(result.PokemonEncounters))
		for _, pokemon := range result.PokemonEncounters {
			resString = append(resString, pokemon.Pokemon.Name)
		}
		if !slices.Equal(resString, location.output) {
			t.Errorf("Did not get expected value of %v", location.output)
		}
	}
}
