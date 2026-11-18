package main

import (
	"bufio"
	"fmt"
	"github.com/btylerh7/bootdev-pokedex/commands"
	"os"
	"strings"
)

var Commands map[string]commands.CliCommand

func initCommands() commands.Config {
	cfg := commands.Config{
		Commands: map[string]commands.CliCommand{
			"help": {
				Name:        "help",
				Description: "Displays a help message",
				Callback:    commands.Help,
			},
			"exit": {
				Name:        "exit",
				Description: "Exit the Pokedex",
				Callback:    commands.Exit,
			},
			"explore": {
				Name:        "explore",
				Description: "Explore Pokemon in an Area",
				Callback:    commands.Explore,
			},
			"map": {
				Name:        "map",
				Description: "Return a map of pokemon locations",
				Callback:    commands.Map,
			},
			"mapb": {
				Name:        "mapb",
				Description: "Return a map of pokemon locations from the previous page",
				Callback:    commands.Mapb,
			},
			"catch": {
				Name:        "catch",
				Description: "Take a pokemon name and see if you catch it!",
				Callback:    commands.Catch,
			},
			"inspect": {
				Name:        "inspect",
				Description: "View information on a pokemon if you have caught it",
				Callback:    commands.Inspect,
			},
			"pokedex": {
				Name:        "pokedex",
				Description: "View the names of all caught pokemon",
				Callback:    commands.Pokedex,
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
func startRepl(cfg *commands.Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleaned := cleanInput(scanner.Text())
		command, ok := cfg.Commands[cleaned[0]]
		if ok {
			err := command.Callback(cfg, cleaned[1:]...)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
}
