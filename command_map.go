package main

import (
	"fmt"
	"github.com/btylerh7/bootdev-pokedex/internal/api"
)

func commandMap(cfg *config) error {
	jsonData, err := api.GetLocationAreas(nil)
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
