package api

// Artist represents a band/artist from the API
type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`    // URL to locations
	ConcertDates string   `json:"concertDates"` // URL to dates
	Relations    string   `json:"relations"`    // URL to relations
}

// Location represents concert locations
type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

// LocationIndex is the root locations object
type LocationIndex struct {
	Index []Location `json:"index"`
}

// Date represents concert dates
type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

// DateIndex is the root dates object
type DateIndex struct {
	Index []Date `json:"index"`
}

// Relation represents the relationship between locations and dates
type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// RelationIndex is the root relation object
type RelationIndex struct {
	Index []Relation `json:"index"`
}

// APIData holds all the combined data
type APIData struct {
	Artists   []Artist
	Locations LocationIndex
	Dates     DateIndex
	Relations RelationIndex
}
