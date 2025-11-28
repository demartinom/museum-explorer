package daily

import (
	"sync"
	"time"

	"github.com/demartinom/museum-explorer/server/internal/core"
)

type DailyArtworkManager struct {
	mu      sync.RWMutex
	Artwork *core.Artwork
}

func (m *DailyArtworkManager) GetArt() *core.Artwork {
	m.mu.RLock()

	defer m.mu.RUnlock()

	return m.Artwork
}

func (m *DailyArtworkManager) SetArtwork(a *core.Artwork) {
	m.mu.Lock()

	defer m.mu.Lock()

	m.Artwork = a
}

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
