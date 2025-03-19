package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type locationAreaResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []locationArea `json:"results"`
}

type locationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func getLocationInformation(url string) (locationAreaResponse, error) {
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	}
	res, err := http.Get(url)
	if err != nil {
		return locationAreaResponse{}, fmt.Errorf("Could not get location data: %w", err)
	}
	defer res.Body.Close()

	var locationReponse locationAreaResponse
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locationReponse); err != nil {
		return locationAreaResponse{}, fmt.Errorf("Could not decode location information: %w", err)
	}
	return locationReponse, nil
}

func commandGet(cfg *Config) error {
	result, err := getLocationInformation()
	if err != nil {
		return fmt.Errorf("Location Information could not be retrieved: %w", err)
	}
	cfg.Next = *result.Next
	cfg.Previous = *result.Previous
	for _, val := range result.Results {
		fmt.Println(val.Name)
	}
	return nil
}

type Config struct {
	Next     string
	Previous string
}
