package handlers

import (
	"log"
	"net/http"

	"groupie-tracker/internal/api"
	"groupie-tracker/internal/utils"
)

// HomeHandler displays the main page with artist list
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		utils.ErrorHandler(w, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		utils.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	data, err := api.FetchAPI()
	if err != nil {
		log.Println("Error fetching API:", err)
		utils.ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	if err := utils.RenderTemplate(w, "index.html", data.Artists); err != nil {
		log.Println("Error rendering template:", err)
		utils.ErrorHandler(w, http.StatusInternalServerError)
	}
}
