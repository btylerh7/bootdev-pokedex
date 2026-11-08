package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type location struct {
	name string
	url  string
}
type mapResponse struct {
	Results []struct {
		Name string `json:name`
		URL  string `json:url`
	} `json:results`
}

func commandMap(cfg *config) error {

	apiURL := "https://pokeapi.co/api/v2/location-area/"
	var jsonData mapResponse
	res, err := http.Get(apiURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&jsonData); err != nil {
		return err
	}
	for _, locationName := range jsonData.Results {
		fmt.Println(locationName.Name)
	}
	return nil
}
