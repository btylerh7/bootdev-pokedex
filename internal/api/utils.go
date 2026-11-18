package api

import (
	"encoding/json"
	"fmt"
	"github.com/btylerh7/bootdev-pokedex/internal/pokecache"
	"net/http"
	"time"
)

var cache = pokecache.NewCache(time.Minute * 5)

func MakeAPICall[T any](url string) (T, error) {
	var jsonData T
	cacheData, ok := cache.Get(url)
	if ok {
		if err := json.Unmarshal(cacheData, &jsonData); err != nil {
			return jsonData, err
		}
		return jsonData, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return jsonData, fmt.Errorf("Unsuccessful HTTP response: %v", err)
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&jsonData); err != nil {
		return jsonData, fmt.Errorf("Unable to decode data: %v", err)
	}

	bytes, err := json.Marshal(jsonData)
	if err != nil {
		var blankRes T
		return blankRes, fmt.Errorf("Unable to encode cache data: %v", err)
	}
	cache.Add(url, bytes)
	return jsonData, nil
}
