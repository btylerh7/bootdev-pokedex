package commands

import (
	"fmt"
	"github.com/btylerh7/bootdev-pokedex/internal/api"
)

func Map(cfg *Config, _ ...string) error {
	jsonData, err := api.GetLocationAreas(cfg.NextUrl)
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
