package metmuseum

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

// Returns a set number of highlights from a specific department at the Met, chosen randomly
func (c *MetClient) cacheHighlights(dept int, numHighlights int) error {
	// TODO: No duplicates
	// TODO: no id = 0
	// TODO: Add temp cache
	//Reading lock for finding all highlight IDs in a specific department
	c.Mu.RLock()
	// Get all department IDs for the Met
	deptHighlightIDs := c.HighlightsIDCache[dept]
	c.Mu.RUnlock()

	if len(deptHighlightIDs) == 0 {
		return fmt.Errorf("no highlight IDs for dept %d", dept)
	}

	// Make highlight cache if doesn't alreay exist
	if c.CachedHighlights == nil {
		c.CachedHighlights = make(map[int][]MetSingleArtwork)
	}
	c.Mu.Lock()
	defer c.Mu.Unlock()

	// Make numHighlights # of calls to get random highglights and save them in cache.
	for i := 0; i < numHighlights; i++ {
		randomHighlight := deptHighlightIDs[rand.Intn(len(deptHighlightIDs))]
		result, err := c.GetSpecific(randomHighlight)
		if err != nil {
			return err
		}
		// Save highlights
		c.CachedHighlights[dept] = append(c.CachedHighlights[dept], *result)
		fmt.Printf("Cached artwork %d: %s\n", result.ObjectID, result.ObjectName)
	}

	return nil
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
	if c.HighlightsIDCache == nil {
		c.HighlightsIDCache = make(map[int][]int)
	}

	tempCache := make(map[int][]int)

	for _, highlight := range highlights {
		tempCache[highlight.DepartmentID] = append(tempCache[highlight.DepartmentID], highlight.ObjectID)
	}

	c.Mu.Lock()
	for deptID, ids := range tempCache {
		c.HighlightsIDCache[deptID] = append(c.HighlightsIDCache[deptID], ids...)
	}

	c.Mu.Unlock()
	return nil
}

// On startup, will run CacheHighlights for all deparments at the Met
// getting 5 highlights per department
func (c *MetClient) DeptHighlightsStartup() {
	var departments []int
	for deptID := range c.HighlightsIDCache {
		departments = append(departments, deptID)
	}

	for id := range departments {
		c.cacheHighlights(id, 5)
	}
}

// Selects a random highlight from the Met
func (c *MetClient) FetchRandom() (*MetSingleArtwork, error) {
	c.Mu.RLock()
	defer c.Mu.RUnlock()

	// Randomly select a department ID
	var departments []int
	for deptID := range c.HighlightsIDCache {
		departments = append(departments, deptID)
	}
	randomDepartment := departments[rand.Intn(len(departments))]

	// Randomly choose an object from the randomly selected department
	objects := c.HighlightsIDCache[randomDepartment]
	randomObject := objects[rand.Intn(len(objects))]

	// Gets information for randomly selected artwork
	result, err := c.GetSpecific(randomObject)
	if err != nil {
		return nil, err
	}

	return result, nil
}
