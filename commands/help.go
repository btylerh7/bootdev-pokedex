package commands

import "fmt"

func Help(cfg *Config, _ ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, command := range cfg.Commands {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}
	return nil

}
