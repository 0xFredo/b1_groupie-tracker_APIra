package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"groupie-tracker/internal/api"
	"groupie-tracker/internal/services"
	"groupie-tracker/internal/utils"
)

// SearchData holds search page data
type SearchData struct {
	Query   string
	Results []api.Artist
}

// SearchHandler handles search form submission
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	query := strings.TrimSpace(r.URL.Query().Get("q"))

	data := SearchData{
		Query: query,
	}

	if query != "" {
		results, err := services.SearchArtists(query)
		if err != nil {
			log.Println("Error searching:", err)
			utils.ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		data.Results = results
	}

	if err := utils.RenderTemplate(w, "search.html", data); err != nil {
		log.Println("Error rendering template:", err)
		utils.ErrorHandler(w, http.StatusInternalServerError)
	}
}

// SuggestionsHandler returns JSON suggestions for search bar
func SuggestionsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	query := strings.TrimSpace(r.URL.Query().Get("q"))
	if query == "" {
		json.NewEncoder(w).Encode([]services.Suggestion{})
		return
	}

	suggestions, err := services.GetSuggestions(query)
	if err != nil {
		log.Println("Error getting suggestions:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(suggestions)
}
