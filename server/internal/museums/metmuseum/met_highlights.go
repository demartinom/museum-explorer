package metmuseum

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

// Returns a set number of highlights from a specific department at the Met, chosen randomly
func (c *MetClient) cacheHighlights(dept string, numHighlights int) error {
	// TODO: No duplicates
	// TODO: no id = 0
	// TODO: Add temp cache
	//Reading lock for finding all highlight IDs in a specific department
	c.Mu.RLock()
	// Get all department IDs for the Met
	deptHighlightIDs := c.HighlightsIDCache[dept]
	c.Mu.RUnlock()

	if len(deptHighlightIDs) == 0 {
		return fmt.Errorf("no highlight IDs for dept %s", dept)
	}

	// Make highlight cache if doesn't alreay exist
	if c.CachedHighlights == nil {
		c.CachedHighlights = make(map[string][]MetSingleArtwork)
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
		c.HighlightsIDCache = make(map[string][]int)
	}

	tempCache := make(map[string][]int)

	for _, highlight := range highlights {
		tempCache[highlight.Department] = append(tempCache[highlight.Department], highlight.ObjectID)
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
	var departments []string

	if os.Getenv("DEV_CACHE") == "1" {
		file, err := os.ReadFile("./data/cachedhighlights.json")
		if err != nil {
			return
		}
		var cachedHighlights []MetSingleArtwork
		json.Unmarshal(file, &cachedHighlights)

		// Make highlight cache if doesn't alreay exist
		if c.CachedHighlights == nil {
			c.CachedHighlights = make(map[string][]MetSingleArtwork)
		}
		for _, item := range cachedHighlights {
			c.CachedHighlights[item.Department] = append(c.CachedHighlights[item.Department], item)
		}
		return
	}

	for deptID := range c.HighlightsIDCache {
		departments = append(departments, deptID)
	}

	for _, deptID := range departments {
		c.cacheHighlights(deptID, 5)
	}
}

// Selects a random highlight from the Met
func (c *MetClient) FetchRandom() (*MetSingleArtwork, error) {
	c.Mu.RLock()
	defer c.Mu.RUnlock()

	// Randomly select a department ID
	var departments []string
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
