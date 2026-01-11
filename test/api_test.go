package test

import (
	"testing"

	"groupie-tracker/internal/api"
)

func TestFetchAPI(t *testing.T) {
	data, err := api.FetchAPI()
	if err != nil {
		t.Fatalf("FetchAPI failed: %v", err)
	}

	if data == nil {
		t.Fatal("Expected data, got nil")
	}

	if len(data.Artists) == 0 {
		t.Error("Expected artists, got none")
	}

	t.Logf("Fetched %d artists", len(data.Artists))
}

func TestGetArtistByID(t *testing.T) {
	artist, err := api.GetArtistByID(1)
	if err != nil {
		t.Fatalf("GetArtistByID failed: %v", err)
	}

	if artist == nil {
		t.Fatal("Expected artist, got nil")
	}

	if artist.ID != 1 {
		t.Errorf("Expected ID 1, got %d", artist.ID)
	}

	t.Logf("Fetched artist: %s", artist.Name)
}

func TestGetRelationByID(t *testing.T) {
	relation, err := api.GetRelationByID(1)
	if err != nil {
		t.Fatalf("GetRelationByID failed: %v", err)
	}

	if relation == nil {
		t.Fatal("Expected relation, got nil")
	}

	if relation.ID != 1 {
		t.Errorf("Expected ID 1, got %d", relation.ID)
	}

	t.Logf("Fetched relation with %d locations", len(relation.DatesLocations))
}
