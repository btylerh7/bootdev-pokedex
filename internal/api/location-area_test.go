package api

import (
	"fmt"
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
			fmt.Errorf("Did not get expected value of %s, got %s", location.err, err)
		}
		resString := make([]string, 0, len(result.Pokemon))
		for _, pokemon := range result.Pokemon {
			resString = append(resString, pokemon.Name)
		}
		if !slices.Equal(resString, location.output) {
			fmt.Errorf("Did not get expected value of %t")
		}
	}
}
