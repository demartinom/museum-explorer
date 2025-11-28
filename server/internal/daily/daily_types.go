package daily

import (
	"sync"
	"time"

	"github.com/demartinom/museum-explorer/server/internal/core"
)

// Struct for holding the artwork of the day
type DailyArtworkManager struct {
	mu      sync.RWMutex
	Artwork *core.Artwork
}

// Allows functions to receive the selected artwork of they day
func (m *DailyArtworkManager) GetArt() *core.Artwork {
	m.mu.RLock()

	defer m.mu.RUnlock()

	return m.Artwork
}

// Takes the randomly selected artwork and sets it as the Artwork of the day
func (m *DailyArtworkManager) SetArtwork(a *core.Artwork) {
	m.mu.Lock()

	defer m.mu.Unlock()

	m.Artwork = a
}

// Spins up a new daily artwork manager
func (m *DailyArtworkManager) Start(client core.MuseumClient) {
	go func() {
		for {
			art, err := client.GetRandomArtwork()
			if err == nil {
				m.SetArtwork(art)
			}

			now := time.Now()
			next := now.Truncate(24 * time.Hour).Add(24 * time.Hour)
			sleepDuration := time.Until(next)

			time.Sleep(sleepDuration)
		}
	}()
}
