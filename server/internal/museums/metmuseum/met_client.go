package metmuseum

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"

	"github.com/demartinom/museum-explorer/server/internal/core"
)

// Initialize a new client for the Met API
func NewClient() *MetClient {
	return &MetClient{BaseURL: "https://collectionapi.metmuseum.org/public/collection/v1"}
}

// Reusable function for making API calls
// Uses type parameter of T for inputting structs into function
func fetchJSON[T any](url string) (*T, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result T
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// Returns a list of all departments within the Met
func (c *MetClient) GetDepartments() (*DepartmentsResponse, error) {
	// Takes base Met url from client and adds departments to the end
	url := fmt.Sprintf("%s/departments", c.BaseURL)

	return fetchJSON[DepartmentsResponse](url)
}

// Selects a random highlight from the Met
func (c *MetClient) FetchRandom() (*MetSingleArtwork, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	// Randomly select a department ID
	var departments []int
	for deptID := range c.HighlightsCache {
		departments = append(departments, deptID)
	}
	randomDepartment := departments[rand.Intn(len(departments))]

	// Randomly choose an object from the randomly selected department
	objects := c.HighlightsCache[randomDepartment]
	randomObject := objects[rand.Intn(len(objects))]

	// Gets information for randomly selected artwork
	result, err := c.GetSpecific(randomObject)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Searches for specific work in collection using it's unique object ID
func (c *MetClient) GetSpecific(objectID int) (*MetSingleArtwork, error) {
	url := fmt.Sprintf("%s/objects/%d", c.BaseURL, objectID)

	return fetchJSON[MetSingleArtwork](url)
}

// Calls function to get random artwork and put it in Core artwork struct
func (c *MetClient) GetRandomArtwork() (*core.Artwork, error) {
	selectedArt, err := c.FetchRandom()
	if err != nil {
		return nil, err
	}

	// Shared type for all different museums
	return &core.Artwork{
		Title:    selectedArt.Title,
		Artist:   selectedArt.ArtistDisplayName,
		ImageURL: selectedArt.PrimaryImage,
		Museum:   "Metropolitan Museum of Art",
	}, nil
}

// Reads JSON with list of Met highlights
// Creates a new cache if one doesn't exist
// Saves highlight ids in map organized by department
func (c *MetClient) DepartmentHighlights(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	var highlights []MetHighlightEntry

	if err := json.NewDecoder(file).Decode(&highlights); err != nil {
		return err
	}
	if c.HighlightsCache == nil {
		c.HighlightsCache = make(map[int][]int)
	}

	tempCache := make(map[int][]int)

	for _, highlight := range highlights {
		tempCache[highlight.DepartmentID] = append(tempCache[highlight.DepartmentID], highlight.ObjectID)
	}

	c.mu.Lock()
	for deptID, ids := range tempCache {
		c.HighlightsCache[deptID] = append(c.HighlightsCache[deptID], ids...)
	}

	c.mu.Unlock()
	return nil
}
