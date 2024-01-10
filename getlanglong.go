package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type GeoResponse struct {
	Results []LatLong `json:"results"`
}

type LatLong struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func getLatLong(city string) (*LatLong, error) {
	endpoint := fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1&language=en&format=json", url.QueryEscape(city))
	result, err := http.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("error getting location")
	}

	defer result.Body.Close()

	var response GeoResponse
	if err := json.NewDecoder(result.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("error decoding json")
	}

	if len(response.Results) < 1 {
		return nil, errors.New("no results found")
	}

	return &response.Results[0], nil
}

func getWeather(latlong LatLong) (string, error) {
	endpoint := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%.6f&longitude=%.6f&hourly=temperature_2m", latlong.Latitude, latlong.Longitude)
	result, err := http.Get(endpoint)
	if err != nil {
		return "", fmt.Errorf("error getting weather")
	}

	defer result.Body.Close()

	body, err := io.ReadAll(result.Body)
	if err != nil {
		return "", fmt.Errorf("error reading data")
	}

	sudoku := map[string]bool{
		"1": false,
	}

	fmt.Println(sudoku)

	return string(body), nil
}
