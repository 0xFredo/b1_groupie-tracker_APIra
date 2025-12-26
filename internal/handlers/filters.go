package handlers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"groupie-tracker/internal/services"
	"groupie-tracker/internal/utils"
)

// FiltersHandler applies filters to artist list
func FiltersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	// Parse filter parameters
	filters := services.FilterParams{
		CreationDateMin: parseIntParam(r, "creation_min", 0),
		CreationDateMax: parseIntParam(r, "creation_max", 9999),
		FirstAlbumMin:   parseIntParam(r, "album_min", 0),
		FirstAlbumMax:   parseIntParam(r, "album_max", 9999),
		MembersMin:      parseIntParam(r, "members_min", 0),
		MembersMax:      parseIntParam(r, "members_max", 100),
		Locations:       parseArrayParam(r, "location"),
	}

	results, err := services.ApplyFilters(filters)
	if err != nil {
		log.Println("Error applying filters:", err)
		utils.ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	if err := utils.RenderTemplate(w, "index.html", results); err != nil {
		log.Println("Error rendering template:", err)
		utils.ErrorHandler(w, http.StatusInternalServerError)
	}
}

func parseIntParam(r *http.Request, key string, defaultVal int) int {
	val := r.URL.Query().Get(key)
	if val == "" {
		return defaultVal
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		return defaultVal
	}
	return i
}

func parseArrayParam(r *http.Request, key string) []string {
	val := r.URL.Query().Get(key)
	if val == "" {
		return []string{}
	}
	return strings.Split(val, ",")
}

// Traitement des filtres (query params)
