package main

import (
	"fmt"
	"log"
	"net/http"

	"groupie-tracker/internal/handlers"
	"groupie-tracker/internal/utils"
)

func main() {
	// Initialize templates
	if err := utils.InitTemplates(); err != nil {
		log.Fatal("Failed to load templates:", err)
	}

	// Setup routes
	mux := http.NewServeMux()

	// Static files
	fs := http.FileServer(http.Dir("web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Pages
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/artist/", handlers.ArtistHandler)
	mux.HandleFunc("/search", handlers.SearchHandler)
	mux.HandleFunc("/api/suggestions", handlers.SuggestionsHandler)
	mux.HandleFunc("/filters", handlers.FiltersHandler)
	mux.HandleFunc("/map/", handlers.GeoHandler)

	port := ":8080"
	fmt.Println("Server running on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, mux))
}
