package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

const (
	baseURL      = "https://groupietrackers.herokuapp.com/api"
	artistsURL   = baseURL + "/artists"
	locationsURL = baseURL + "/locations"
	datesURL     = baseURL + "/dates"
	relationsURL = baseURL + "/relation"
	cacheTimeout = 30 * time.Minute
)

var (
	cache      *APIData
	cacheMutex sync.RWMutex
	lastFetch  time.Time
)

// FetchAPI retrieves all API data with caching
func FetchAPI() (*APIData, error) {
	cacheMutex.RLock()
	if cache != nil && time.Since(lastFetch) < cacheTimeout {
		cacheMutex.RUnlock()
		return cache, nil
	}
	cacheMutex.RUnlock()

	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	// Double-check after acquiring write lock
	if cache != nil && time.Since(lastFetch) < cacheTimeout {
		return cache, nil
	}

	data := &APIData{}

	// Fetch all data in parallel
	var wg sync.WaitGroup
	errChan := make(chan error, 4)

	wg.Add(4)

	go func() {
		defer wg.Done()
		if err := fetchJSON(artistsURL, &data.Artists); err != nil {
			errChan <- fmt.Errorf("artists: %w", err)
		}
	}()

	go func() {
		defer wg.Done()
		if err := fetchJSON(locationsURL, &data.Locations); err != nil {
			errChan <- fmt.Errorf("locations: %w", err)
		}
	}()

	go func() {
		defer wg.Done()
		if err := fetchJSON(datesURL, &data.Dates); err != nil {
			errChan <- fmt.Errorf("dates: %w", err)
		}
	}()

	go func() {
		defer wg.Done()
		if err := fetchJSON(relationsURL, &data.Relations); err != nil {
			errChan <- fmt.Errorf("relations: %w", err)
		}
	}()

	wg.Wait()
	close(errChan)

	// Check for errors
	for err := range errChan {
		return nil, err
	}

	cache = data
	lastFetch = time.Now()
	return data, nil
}

// fetchJSON performs HTTP GET and decodes JSON
func fetchJSON(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("status %d: %s", resp.StatusCode, string(body))
	}

	return json.NewDecoder(resp.Body).Decode(target)
}

// GetArtistByID returns a specific artist by ID
func GetArtistByID(id int) (*Artist, error) {
	data, err := FetchAPI()
	if err != nil {
		return nil, err
	}

	for _, artist := range data.Artists {
		if artist.ID == id {
			return &artist, nil
		}
	}
	return nil, fmt.Errorf("artist not found")
}

// GetRelationByID returns relations for a specific artist
func GetRelationByID(id int) (*Relation, error) {
	data, err := FetchAPI()
	if err != nil {
		return nil, err
	}

	for _, rel := range data.Relations.Index {
		if rel.ID == id {
			return &rel, nil
		}
	}
	return nil, fmt.Errorf("relation not found")
}
