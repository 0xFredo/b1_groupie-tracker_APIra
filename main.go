package main

import (
	"flag"
	"fmt"
	"groupie-tracker/internal/handlers"
	"groupie-tracker/internal/utils"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func openBrowser(url string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	default:
		log.Printf("Could not determine how to open browser on %s", runtime.GOOS)
		return
	}
	if err := cmd.Start(); err != nil {
		log.Printf("Failed to open browser: %v", err)
	}
}

func main() {
	defaultPort := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		defaultPort = envPort
	}
	addr := flag.String("addr", ":"+defaultPort, "HTTP network address")
	flag.Parse()

	if err := utils.InitTemplates(); err != nil {
		log.Fatal("Failed to load templates:", err)
	}

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/artist/", handlers.ArtistHandler)
	mux.HandleFunc("/search", handlers.SearchHandler)
	mux.HandleFunc("/api/suggestions", handlers.SuggestionsHandler)
	mux.HandleFunc("/map/", handlers.GeoHandler)

	server := &http.Server{
		Addr:    *addr,
		Handler: mux,
	}

	url := fmt.Sprintf("http://localhost%s", *addr)
	log.Printf("Starting server on %s", url)
	go func() {
		time.Sleep(500 * time.Millisecond)
		openBrowser(url)
	}()

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
