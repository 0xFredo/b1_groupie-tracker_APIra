package handlers

import (
	"log"
	"net/http"

	"groupie-tracker/internal/api"
	"groupie-tracker/internal/services"
	"groupie-tracker/internal/utils"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		utils.ErrorHandler(w, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		utils.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	hasFilters := r.URL.RawQuery != ""

	var artists []api.Artist
	var err error

	if hasFilters {
		filters := services.FilterParams{
			CreationDateMin: parseIntParam(r, "creation_min", 0),
			CreationDateMax: parseIntParam(r, "creation_max", 9999),
			FirstAlbumMin:   parseIntParam(r, "album_min", 0),
			FirstAlbumMax:   parseIntParam(r, "album_max", 9999),
			MembersMin:      parseIntParam(r, "members_min", 0),
			MembersMax:      parseIntParam(r, "members_max", 100),
			Locations:       parseArrayParam(r, "location"),
		}
		artists, err = services.ApplyFilters(filters)
	} else {
		var data *api.APIData
		data, err = api.FetchAPI()
		if err == nil {
			artists = data.Artists
		}
	}

	if err != nil {
		log.Println("Error fetching data:", err)
		utils.ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	pageData := utils.PageData{
		ActiveTab:       "home",
		ContentTemplate: "index",
		Data:            artists,
	}

	if err := utils.RenderTemplate(w, "index.html", pageData); err != nil {
		log.Println("Error rendering template:", err)
		utils.ErrorHandler(w, http.StatusInternalServerError)
	}
}
