package test

import (
	"testing"

	"groupie-tracker/internal/api"
	"groupie-tracker/internal/services"
)

func TestGeocodeLocations(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping geocoding test in short mode")
	}

	relation, err := api.GetRelationByID(1)
	if err != nil {
		t.Fatalf("Failed to get relation: %v", err)
	}

	locations, err := services.GeocodeLocations(relation)
	if err != nil {
		t.Fatalf("GeocodeLocations failed: %v", err)
	}

	t.Logf("Geocoded %d locations", len(locations))

	if len(locations) == 0 {
		t.Log("Warning: No locations were successfully geocoded")
	}

	for _, loc := range locations {
		if loc.Latitude == 0 && loc.Longitude == 0 {
			t.Errorf("Location %s has invalid coordinates (0,0)", loc.Name)
		}
		t.Logf("Location: %s (%.4f, %.4f)", loc.Name, loc.Latitude, loc.Longitude)
	}
}
