package metmuseum

import (
	"encoding/json"
	"fmt"
	"net/http"

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
