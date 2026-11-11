package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}
type config struct {
	Commands map[string]cliCommand
	NextUrl  *string
	PrevUrl  *string
}

var Commands map[string]cliCommand

func initCommands() config {
	cfg := config{
		Commands: map[string]cliCommand{
			"help": {
				name:        "help",
				description: "Displays a help message",
				callback:    commandHelp,
			},
			"exit": {
				name:        "exit",
				description: "Exit the Pokedex",
				callback:    commandExit,
			},
			"map": {
				name:        "map",
				description: "Return a map of pokemon locations",
				callback:    commandMap,
			},
			"mapb": {
				name:        "mapb",
				description: "Return a map of pokemon locations from the previous page",
				callback:    commandMapb,
			},
		},
	}
	return cfg

}

func cleanInput(text string) []string {
	var result []string

	for word := range strings.SplitSeq(text, " ") {
		lower := strings.ToLower(word)
		if len(word) == 0 {
			continue
		}
		result = append(result, strings.Trim(lower, " "))
	}

	return result

}
func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleaned := cleanInput(scanner.Text())
		command, ok := cfg.Commands[cleaned[0]]
		if ok {
			command.callback(cfg)
		}

	}
}
