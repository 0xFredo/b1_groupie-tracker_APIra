package test

import (
	"testing"

	"groupie-tracker/internal/services"
)

// TestSearchArtists tests the search functionality
func TestSearchArtists(t *testing.T) {
	tests := []struct {
		query    string
		expected bool
	}{
		{"queen", true},
		{"freddie", true},
		{"1985", true},
		{"nonexistent", false},
	}

	for _, tt := range tests {
		t.Run(tt.query, func(t *testing.T) {
			results, err := services.SearchArtists(tt.query)
			if err != nil {
				t.Fatalf("SearchArtists failed: %v", err)
			}

			hasResults := len(results) > 0
			if hasResults != tt.expected {
				t.Errorf("For query %q, expected hasResults=%v, got %v", tt.query, tt.expected, hasResults)
			}

			t.Logf("Query %q returned %d results", tt.query, len(results))
		})
	}
}

// TestGetSuggestions tests search suggestions
func TestGetSuggestions(t *testing.T) {
	suggestions, err := services.GetSuggestions("qu")
	if err != nil {
		t.Fatalf("GetSuggestions failed: %v", err)
	}

	t.Logf("Got %d suggestions for 'qu'", len(suggestions))

	// Test empty query
	emptySuggestions, err := services.GetSuggestions("")
	if err != nil {
		t.Fatalf("GetSuggestions with empty query failed: %v", err)
	}

	if len(emptySuggestions) > 0 {
		t.Error("Expected no suggestions for empty query")
	}
}
