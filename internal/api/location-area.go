package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/btylerh7/bootdev-pokedex/internal/pokecache"
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

type locationResponse struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

var cache = pokecache.NewCache(time.Minute * 5)

func GetLocationAreas(url *string) (mapResponse, error) {

	apiURL := "https://pokeapi.co/api/v2/location-area/"
	if url != nil {
		apiURL = *url
	}
	var jsonData mapResponse
	cacheData, ok := cache.Get(apiURL)
	if ok {
		if err := json.Unmarshal(cacheData, &jsonData); err != nil {
			return jsonData, err
		}
		return jsonData, nil
	}

	res, err := http.Get(apiURL)
	if err != nil {
		return jsonData, fmt.Errorf("Unsuccessful HTTP response: %v", err)
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&jsonData); err != nil {
		return jsonData, fmt.Errorf("Unable to decode data: %v", err)
	}

	bytes, err := json.Marshal(jsonData)
	if err != nil {
		var mapres mapResponse
		return mapres, fmt.Errorf("Unable to encode cache data: %v", err)
	}
	cache.Add(apiURL, bytes)
	return jsonData, nil

}
func GetPokemonForLocation(location string) (locationResponse, error) {
	apiURL := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", location)
	var jsonData locationResponse
	cacheData, ok := cache.Get(apiURL)
	if ok {
		if err := json.Unmarshal(cacheData, &jsonData); err != nil {
			return jsonData, err
		}
		return jsonData, nil

	}
	res, err := http.Get(apiURL)
	if err != nil {
		return jsonData, fmt.Errorf("Unsuccessful HTTP response: %v", err)
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&jsonData); err != nil {
		return jsonData, fmt.Errorf("Unable to decode data: %v", err)
	}
	bytes, err := json.Marshal(jsonData)
	if err != nil {
		var locres locationResponse
		return locres, fmt.Errorf("Unable to encode cache data: %v", err)
	}
	cache.Add(apiURL, bytes)
	return jsonData, nil

}
