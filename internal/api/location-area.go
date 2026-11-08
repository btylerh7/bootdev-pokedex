package api

import (
	"encoding/json"
	"net/http"
)

type mapResponse struct {
	Count    int
	Next     *string
	Previous *string
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(url *string) (mapResponse, error) {

	apiURL := "https://pokeapi.co/api/v2/location-area/"
	if url != nil {
		apiURL = *url
	}
	var jsonData mapResponse
	res, err := http.Get(apiURL)
	if err != nil {
		return jsonData, err
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&jsonData); err != nil {
		return jsonData, err
	}
	return jsonData, nil
}
