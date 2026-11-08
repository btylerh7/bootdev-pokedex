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
	callback    func() error
}

var Commands map[string]cliCommand

func initCommands() {
	Commands = map[string]cliCommand{
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
	}

}

func cleanInput(text string) []string {
	var result []string

	for word := range strings.SplitSeq(text, " ") {
		lower := strings.ToLower(word)
		result = append(result, strings.Trim(lower, " "))
	}

	return result

}
func startRepl() {
	initCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleaned := cleanInput(scanner.Text())
		command, ok := Commands[cleaned[0]]
		if ok {
			command.callback()
		}

	}
}
