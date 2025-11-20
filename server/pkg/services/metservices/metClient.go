package metservices

import (
	"encoding/json"
	"fmt"
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
