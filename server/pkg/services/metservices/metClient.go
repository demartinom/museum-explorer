package metservices

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

// Initialize a new client for the Met API
func NewClient() *MetClient {
	return &MetClient{BaseURL: "https://collectionapi.metmuseum.org/public/collection/v1"}
}

// Returns a list of all departments within the Met
func (c *MetClient) GetDepartments() (*DepartmentsResponse, error) {
	// Takes base Met url from client and adds departments to the end
	url := fmt.Sprintf("%s/departments", c.BaseURL)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Instance of DepartmentsResponse struct for data to be put into
	var result DepartmentsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// Gets a list of Art IDs for the highlights of the Met and then returns a random artwork from the highlights
func (c *MetClient) GetRandom() (*MetSingleArtwork, error) {
	url := fmt.Sprintf("%s/search?q=isHightlight=true", c.BaseURL)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var randomResult GetHighlightIDs
	if err := json.NewDecoder(resp.Body).Decode(&randomResult); err != nil {
		return nil, err
	}

	// TODO: Ensure it doesn't return something blank or with no picture
	// Selects a random ID
	randArtworkID := rand.Intn(len(randomResult.ObjectIds))

	// Gets information for randomly selected artwork
	result, err := c.GetSpecific(randomResult.ObjectIds[randArtworkID])
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Searches for specific work in collection using it's unique object ID
func (c *MetClient) GetSpecific(objectID int) (*MetSingleArtwork, error) {
	url := fmt.Sprintf("%s/objects/%d", c.BaseURL, objectID)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result MetSingleArtwork
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
