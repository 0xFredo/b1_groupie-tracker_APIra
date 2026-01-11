package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"groupie-tracker/internal/api"
)

type GeoLocation struct {
	Name      string
	Latitude  float64
	Longitude float64
	Dates     []string
}

type NominatimResponse []struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

var (
	geoCache      = make(map[string][2]float64)
	geoCacheMutex sync.RWMutex
)

func GeocodeLocations(relation *api.Relation) ([]GeoLocation, error) {
	var geoLocations []GeoLocation
	count := 0
	maxLocations := 5

	for location, dates := range relation.DatesLocations {
		if count >= maxLocations {
			break
		}

		formattedLocation := formatLocation(location)

		lat, lon, err := geocode(formattedLocation)
		if err != nil {
			continue
		}

		geoLocations = append(geoLocations, GeoLocation{
			Name:      formattedLocation,
			Latitude:  lat,
			Longitude: lon,
			Dates:     dates,
		})
		count++
	}

	return geoLocations, nil
}

func formatLocation(location string) string {
	parts := strings.Split(location, "-")
	for i, part := range parts {
		part = strings.ReplaceAll(part, "_", " ")
		parts[i] = strings.Title(strings.ToLower(part))
	}
	return strings.Join(parts, ", ")
}

func geocode(address string) (float64, float64, error) {
	geoCacheMutex.RLock()
	if coords, ok := geoCache[address]; ok {
		geoCacheMutex.RUnlock()
		return coords[0], coords[1], nil
	}
	geoCacheMutex.RUnlock()

	baseURL := "https://nominatim.openstreetmap.org/search"
	params := url.Values{}
	params.Add("q", address)
	params.Add("format", "json")
	params.Add("limit", "1")

	reqURL := baseURL + "?" + params.Encode()

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return 0, 0, err
	}

	req.Header.Set("User-Agent", "Groupie-Tracker/1.0")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, 0, fmt.Errorf("geocoding failed: status %d", resp.StatusCode)
	}

	var result NominatimResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, 0, err
	}

	if len(result) == 0 {
		return 0, 0, fmt.Errorf("location not found")
	}

	var lat, lon float64
	fmt.Sscanf(result[0].Lat, "%f", &lat)
	fmt.Sscanf(result[0].Lon, "%f", &lon)

	geoCacheMutex.Lock()
	geoCache[address] = [2]float64{lat, lon}
	geoCacheMutex.Unlock()

	return lat, lon, nil
}
