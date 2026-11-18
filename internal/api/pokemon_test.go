package api

import (
	"strings"
	"testing"
)

type getPokemonTestCase struct {
	input      string
	output     getPokemonResponse
	errMessage string
}

func TestGetPokemon(t *testing.T) {
	cases := []getPokemonTestCase{
		{
			input:      "pikachu",
			output:     getPokemonResponse{Name: "pikachu", BaseExperience: 112},
			errMessage: "",
		},
		{
			input:      "notreal",
			output:     getPokemonResponse{},
			errMessage: "Unsuccessful HTTP response",
		},
	}
	for _, testCase := range cases {
		result, err := GetPokemon(testCase.input)
		if err != nil {
			if testCase.errMessage == "" {
				t.Errorf("Expected result and received error: %v", err)
			}
			if !strings.Contains(err.Error(), testCase.errMessage) {
				t.Errorf("Expected error message %s and received error: %v", testCase.errMessage, err)
			}
			continue
		}
		if (result == testCase.output) == false {
			t.Errorf("Expected %v and received %v", testCase.output, result)
		}
	}
}
