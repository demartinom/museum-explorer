package museums

import (
	"encoding/json"
	"os"
)

// Function to read museum list json and put it in slice of Museum structs
func LoadMuseums(path string) ([]Museum, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var museums []Museum
	if err := json.Unmarshal(data, &museums); err != nil {
		return nil, err
	}
	return museums, nil
}
