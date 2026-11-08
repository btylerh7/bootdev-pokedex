package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	var result []string

	for word := range strings.SplitSeq(text, " ") {
		result = append(result, strings.Trim(word, " "))
	}

	return result

}
func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleaned := cleanInput(scanner.Text())
		if len(cleaned) > 0 {
			fmt.Printf("Your command was: %s\n", cleaned[0])
		}

	}
}
