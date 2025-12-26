package services

import (
	"strconv"
	"strings"

	"groupie-tracker/internal/api"
)

// Suggestion represents a search suggestion with type
type Suggestion struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

// SearchArtists searches for artists by multiple criteria (case-insensitive)
func SearchArtists(query string) ([]api.Artist, error) {
	data, err := api.FetchAPI()
	if err != nil {
		return nil, err
	}

	query = strings.ToLower(query)
	var results []api.Artist

	for _, artist := range data.Artists {
		if matchesArtist(artist, query) {
			results = append(results, artist)
		}
	}

	return results, nil
}

// matchesArtist checks if artist matches search query
func matchesArtist(artist api.Artist, query string) bool {
	// Artist/band name
	if strings.Contains(strings.ToLower(artist.Name), query) {
		return true
	}

	// Members
	for _, member := range artist.Members {
		if strings.Contains(strings.ToLower(member), query) {
			return true
		}
	}

	// Creation date
	if strings.Contains(strconv.Itoa(artist.CreationDate), query) {
		return true
	}

	// First album
	if strings.Contains(strings.ToLower(artist.FirstAlbum), query) {
		return true
	}

	// TODO: Add location matching (requires relation data)

	return false
}

// GetSuggestions returns typed suggestions for search bar
func GetSuggestions(query string) ([]Suggestion, error) {
	if query == "" {
		return []Suggestion{}, nil
	}

	data, err := api.FetchAPI()
	if err != nil {
		return nil, err
	}

	query = strings.ToLower(query)
	suggestions := []Suggestion{}
	seen := make(map[string]bool)

	for _, artist := range data.Artists {
		// Artist name
		if strings.Contains(strings.ToLower(artist.Name), query) {
			key := artist.Name + "-artist"
			if !seen[key] {
				suggestions = append(suggestions, Suggestion{
					Text: artist.Name,
					Type: "artist/band",
				})
				seen[key] = true
			}
		}

		// Members
		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), query) {
				key := member + "-member"
				if !seen[key] {
					suggestions = append(suggestions, Suggestion{
						Text: member,
						Type: "member",
					})
					seen[key] = true
				}
			}
		}

		// Creation date
		creationStr := strconv.Itoa(artist.CreationDate)
		if strings.Contains(creationStr, query) {
			key := creationStr + "-creation"
			if !seen[key] {
				suggestions = append(suggestions, Suggestion{
					Text: creationStr,
					Type: "creation date",
				})
				seen[key] = true
			}
		}

		// First album
		if strings.Contains(strings.ToLower(artist.FirstAlbum), query) {
			key := artist.FirstAlbum + "-album"
			if !seen[key] {
				suggestions = append(suggestions, Suggestion{
					Text: artist.FirstAlbum,
					Type: "first album date",
				})
				seen[key] = true
			}
		}
	}

	// Limit suggestions
	if len(suggestions) > 10 {
		suggestions = suggestions[:10]
	}

	return suggestions, nil
}

// Logique de recherche (case-insensitive)
