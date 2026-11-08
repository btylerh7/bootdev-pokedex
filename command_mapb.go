package main

import (
	"fmt"
	"github.com/btylerh7/bootdev-pokedex/internal/api"
)

func commandMapb(cfg *config) error {
	jsonData, err := api.GetLocationAreas(cfg.PrevUrl)
	if err != nil {
		return err
	}
	cfg.NextUrl = jsonData.Next
	cfg.PrevUrl = jsonData.Previous
	for _, locationName := range jsonData.Results {
		fmt.Println(locationName.Name)
	}
	return nil
}
