package metservices

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/demartinom/museum-explorer/server/internal/services/core"
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

// Gets a list of Art IDs for the highlights of the Met and then returns a random artwork from the highlights
func (c *MetClient) GetRandom() (*MetSingleArtwork, error) {
	url := fmt.Sprintf("%s/search?q=&isHighlight=true", c.BaseURL)

	HighlightIDs, err := fetchJSON[GetHighlightIDs](url)
	if err != nil {
		return nil, err
	}

	if len(HighlightIDs.ObjectIds) == 0 {
		return nil, fmt.Errorf("no highlight artworks returned from API")
	}
	// Selects a random ID
	randArtworkID := rand.Intn(len(HighlightIDs.ObjectIds))

	// Gets information for randomly selected artwork
	result, err := c.GetSpecific(HighlightIDs.ObjectIds[randArtworkID])
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
func (c *MetClient) FetchRandomDaily() (*core.Artwork, error) {
	selectedArt, _ := c.GetRandom()

	// Shared type for all different museums
	return &core.Artwork{
		Title:    selectedArt.Title,
		Artist:   selectedArt.ArtistDisplayName,
		ImageURL: selectedArt.PrimaryImage,
		Museum:   "Metropolitan Museum of Art",
	}, nil
}
