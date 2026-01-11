package handlers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"groupie-tracker/internal/api"
	"groupie-tracker/internal/utils"
)

type ArtistData struct {
	Artist   *api.Artist
	Relation *api.Relation
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}

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

	artistData := ArtistData{
		Artist:   artist,
		Relation: relation,
	}

	pageData := utils.PageData{
		Title:           artist.Name,
		ActiveTab:       "",
		ContentTemplate: "artist",
		Data:            artistData,
	}

	if err := utils.RenderTemplate(w, "artist.html", pageData); err != nil {
		log.Println("Error rendering template:", err)
		utils.ErrorHandler(w, http.StatusInternalServerError)
	}
}
