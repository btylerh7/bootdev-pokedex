package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, command := range cfg.Commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil

}
