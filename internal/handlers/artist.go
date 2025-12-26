package handlers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"groupie-tracker/internal/api"
	"groupie-tracker/internal/utils"
)

// ArtistData combines artist info with relations
type ArtistData struct {
	Artist   *api.Artist
	Relation *api.Relation
}

// ArtistHandler displays detailed artist page with concerts
func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from path /artist/{id}
	path := strings.TrimPrefix(r.URL.Path, "/artist/")
	id, err := strconv.Atoi(path)
	if err != nil || id < 1 {
		utils.ErrorHandler(w, http.StatusBadRequest)
		return
	}

	artist, err := api.GetArtistByID(id)
	if err != nil {
		log.Println("Error fetching artist:", err)
		utils.ErrorHandler(w, http.StatusNotFound)
		return
	}

	relation, err := api.GetRelationByID(id)
	if err != nil {
		log.Println("Error fetching relation:", err)
		utils.ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	data := ArtistData{
		Artist:   artist,
		Relation: relation,
	}

	if err := utils.RenderTemplate(w, "artist.html", data); err != nil {
		log.Println("Error rendering template:", err)
		utils.ErrorHandler(w, http.StatusInternalServerError)
	}
}
