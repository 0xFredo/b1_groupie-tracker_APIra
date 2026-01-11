package test

import (
	"testing"

	"groupie-tracker/internal/services"
)

func TestApplyFilters(t *testing.T) {
	params := services.FilterParams{
		CreationDateMin: 1980,
		CreationDateMax: 2000,
		FirstAlbumMin:   0,
		FirstAlbumMax:   9999,
		MembersMin:      0,
		MembersMax:      100,
	}

	results, err := services.ApplyFilters(params)
	if err != nil {
		t.Fatalf("ApplyFilters failed: %v", err)
	}

	t.Logf("Found %d artists created between 1980-2000", len(results))

	for _, artist := range results {
		if artist.CreationDate < 1980 || artist.CreationDate > 2000 {
			t.Errorf("Artist %s creation date %d outside range 1980-2000", artist.Name, artist.CreationDate)
		}
	}
}

func TestMembersFilter(t *testing.T) {
	params := services.FilterParams{
		CreationDateMin: 0,
		CreationDateMax: 9999,
		FirstAlbumMin:   0,
		FirstAlbumMax:   9999,
		MembersMin:      4,
		MembersMax:      4,
	}

	results, err := services.ApplyFilters(params)
	if err != nil {
		t.Fatalf("ApplyFilters failed: %v", err)
	}

	t.Logf("Found %d artists with exactly 4 members", len(results))

	for _, artist := range results {
		if len(artist.Members) != 4 {
			t.Errorf("Artist %s has %d members, expected 4", artist.Name, len(artist.Members))
		}
	}
}
