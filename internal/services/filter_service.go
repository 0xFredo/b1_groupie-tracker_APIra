package services

import (
	"strconv"
	"strings"

	"groupie-tracker/internal/api"
)

// FilterParams holds all filter criteria
type FilterParams struct {
	CreationDateMin int
	CreationDateMax int
	FirstAlbumMin   int
	FirstAlbumMax   int
	MembersMin      int
	MembersMax      int
	Locations       []string
}

// ApplyFilters filters artists based on criteria
func ApplyFilters(params FilterParams) ([]api.Artist, error) {
	data, err := api.FetchAPI()
	if err != nil {
		return nil, err
	}

	var results []api.Artist

	for _, artist := range data.Artists {
		if matchesFilters(artist, params, data) {
			results = append(results, artist)
		}
	}

	return results, nil
}

// matchesFilters checks if artist matches all filter criteria
func matchesFilters(artist api.Artist, params FilterParams, data *api.APIData) bool {
	// Creation date filter
	if artist.CreationDate < params.CreationDateMin || artist.CreationDate > params.CreationDateMax {
		return false
	}

	// First album filter (extract year from string like "14-12-1995")
	albumYear := extractYear(artist.FirstAlbum)
	if albumYear < params.FirstAlbumMin || albumYear > params.FirstAlbumMax {
		return false
	}

	// Members count filter
	memberCount := len(artist.Members)
	if memberCount < params.MembersMin || memberCount > params.MembersMax {
		return false
	}

	// Location filter (check if artist has concerts in selected locations)
	if len(params.Locations) > 0 {
		if !hasLocationMatch(artist.ID, params.Locations, data) {
			return false
		}
	}

	return true
}

// extractYear extracts year from date string (e.g., "14-12-1995" -> 1995)
func extractYear(dateStr string) int {
	parts := strings.Split(dateStr, "-")
	if len(parts) == 3 {
		year, err := strconv.Atoi(parts[2])
		if err == nil {
			return year
		}
	}
	return 0
}

// hasLocationMatch checks if artist has concerts in specified locations
func hasLocationMatch(artistID int, locations []string, data *api.APIData) bool {
	for _, relation := range data.Relations.Index {
		if relation.ID == artistID {
			for location := range relation.DatesLocations {
				locationLower := strings.ToLower(location)
				for _, filterLoc := range locations {
					filterLower := strings.ToLower(filterLoc)
					// Check if location contains filter (e.g., "seattle-washington-usa" contains "washington")
					if strings.Contains(locationLower, filterLower) {
						return true
					}
				}
			}
			break
		}
	}
	return false
}

// Logique de filtrage (dates, membres, lieux)
