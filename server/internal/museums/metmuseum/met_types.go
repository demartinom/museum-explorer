package metmuseum

import "sync"

// Init new Client for Met API calls
type MetClient struct {
	BaseURL string
	// mu for reading and writing safely
	Mu sync.RWMutex
	// Map of departments with a slice of the highlights found in the department
	HighlightsIDCache map[string][]int
	CachedHighlights  map[string][]MetSingleArtwork
}

// Returns an array of the different departments within the Met
type MetDepartments struct {
	Departments []MetSingleDepartment `json:"departments"`
}

// Struct for returning list of departments at the Met
type DepartmentsResponse struct {
	Departments []MetSingleDepartment `json:"departments"`
}

// Container for a single department within the array of Met departments
type MetSingleDepartment struct {
	DepartmentID int    `json:"departmentId"`
	DisplayName  string `json:"displayName"`
}

type MetHighlightEntry struct {
	ObjectID   int    `json:"objectId"`
	Department string `json:"departmentName"`
}

// Struct for a single piece of art from the Met
type MetSingleArtwork struct {
	ObjectID          int    `json:"objectID"`
	IsHighlight       bool   `json:"isHighlight"`
	PrimaryImage      string `json:"primaryImage"`
	PrimaryImageSmall string `json:"primaryImageSmall"`
	Department        string `json:"department"`
	ObjectName        string `json:"objectName"`
	Title             string `json:"title"`
	ArtistDisplayName string `json:"artistDisplayName"`
	ObjectDate        string `json:"objectDate"`
}

var DepartmentIDToName = map[int]string{
	1:  "American Decorative Arts",
	3:  "Ancient West Asian Art",
	4:  "Arms and Armor",
	5:  "Arts of Africa, Oceania, and the Americas",
	6:  "Asian Art",
	7:  "The Cloisters",
	8:  "Costume Institute",
	9:  "Drawings and Prints",
	10: "Egyptian Art",
	11: "European Paintings",
	12: "European Sculpture and Decorative Arts",
	13: "Greek and Roman Art",
	14: "Islamic Art",
	15: "Robert Lehman Collection",
	16: "The Libraries",
	17: "Medieval Art",
	18: "Musical Instruments",
	19: "Photographs",
	21: "Modern Art",
}
