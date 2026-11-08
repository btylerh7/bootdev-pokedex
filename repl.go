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
		},
	}
	return cfg

}

func cleanInput(text string) []string {
	var result []string

	for word := range strings.SplitSeq(text, " ") {
		lower := strings.ToLower(word)
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
