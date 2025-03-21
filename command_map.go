package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Config struct {
	Next     string
	Previous string
}

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

func getLocationInformation(url string, cfg *Config) (locationAreaResponse, error) {
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

func commandGetMap(cfg *Config) error {
	url := cfg.Next
	result, err := getLocationInformation(url, cfg)
	if err != nil {
		return fmt.Errorf("Location Information could not be retrieved: ", err)
	}

	if result.Next != nil {
		cfg.Next = *result.Next
	} else {
		cfg.Next = ""
	}

	if result.Previous != nil {
		cfg.Previous = *result.Previous
	} else {
		cfg.Previous = ""
	}

	for _, val := range result.Results {
		fmt.Println(val.Name)
	}
	return nil
}

func commandGetMapb(cfg *Config) error {
	if cfg.Previous == "" {
		fmt.Println("You are on the first Page")
		return nil
	}

	result, err := getLocationInformation(cfg.Previous, cfg)
	if err != nil {
		return fmt.Errorf("Location information could not be retrevied: ", err)
	}

	if result.Next != nil {
		cfg.Next = *result.Next
	} else {
		cfg.Next = ""
	}

	if result.Previous != nil {
		cfg.Previous = *result.Previous
	} else {
		cfg.Previous = ""
	}

	for _, val := range result.Results {
		fmt.Println(val.Name)
	}

	return nil
}
