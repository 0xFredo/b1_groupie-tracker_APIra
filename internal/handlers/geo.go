package handlers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"groupie-tracker/internal/api"
	"groupie-tracker/internal/services"
	"groupie-tracker/internal/utils"
)

// GeoData combines artist and geocoded locations
type GeoData struct {
	Artist    *api.Artist
	Locations []services.GeoLocation
}

// GeoHandler displays map with concert locations
func GeoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from path /map/{id}
	path := strings.TrimPrefix(r.URL.Path, "/map/")
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

	locations, err := services.GeocodeLocations(relation)
	if err != nil {
		log.Println("Error geocoding locations:", err)
		utils.ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	data := GeoData{
		Artist:    artist,
		Locations: locations,
	}

	if err := utils.RenderTemplate(w, "map.html", data); err != nil {
		log.Println("Error rendering template:", err)
		utils.ErrorHandler(w, http.StatusInternalServerError)
	}
}

// Page/section géolocalisation si tu fais l'option
