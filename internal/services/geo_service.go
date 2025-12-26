package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"groupie-tracker/internal/api"
)

// GeoLocation represents a geocoded location with coordinates
type GeoLocation struct {
	Name      string
	Latitude  float64
	Longitude float64
	Dates     []string
}

// NominatimResponse represents the response from Nominatim API
type NominatimResponse []struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

// GeocodeLocations converts relation locations to geocoded coordinates
func GeocodeLocations(relation *api.Relation) ([]GeoLocation, error) {
	var geoLocations []GeoLocation

	for location, dates := range relation.DatesLocations {
		// Format location for geocoding (e.g., "seattle-washington-usa" -> "Seattle, Washington, USA")
		formattedLocation := formatLocation(location)

		lat, lon, err := geocode(formattedLocation)
		if err != nil {
			// If geocoding fails, skip this location
			continue
		}

		geoLocations = append(geoLocations, GeoLocation{
			Name:      formattedLocation,
			Latitude:  lat,
			Longitude: lon,
			Dates:     dates,
		})
	}

	return geoLocations, nil
}

// formatLocation converts API location format to human-readable
func formatLocation(location string) string {
	// Replace hyphens and underscores with spaces, capitalize words
	parts := strings.Split(location, "-")
	for i, part := range parts {
		part = strings.ReplaceAll(part, "_", " ")
		parts[i] = strings.Title(strings.ToLower(part))
	}
	return strings.Join(parts, ", ")
}

// geocode converts address to coordinates using Nominatim (OpenStreetMap)
func geocode(address string) (float64, float64, error) {
	// Using Nominatim API (free, no API key required)
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

	// Nominatim requires a User-Agent header
	req.Header.Set("User-Agent", "Groupie-Tracker/1.0")

	resp, err := http.DefaultClient.Do(req)
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

	return lat, lon, nil
}

// Conversion adresses → coordonnées
