package handlers

import (
	"log"
	"net/http"

	"groupie-tracker/internal/utils"
)

// GeoHandler displays map page (currently disabled due to API issues)
func GeoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	pageData := utils.PageData{
		Title:           "Map",
		ActiveTab:       "map",
		ContentTemplate: "map",
	}

	// Map functionality is temporarily disabled
	if err := utils.RenderTemplate(w, "map.html", pageData); err != nil {
		log.Println("Error rendering template:", err)
		utils.ErrorHandler(w, http.StatusInternalServerError)
	}
}
