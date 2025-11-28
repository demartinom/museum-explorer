package dailyhandlers

import (
	"encoding/json"
	"net/http"

	"github.com/demartinom/museum-explorer/server/internal/daily"
)

// Handler that makes API call return random artwork from the Met's highlights for the Artwork of the day
// Uses handlerfunc to pass in met client from main.go
func DailyArtworkHandler(manager *daily.DailyArtworkManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		art := manager.GetArt()
		if art == nil {
			http.Error(w, "Artwork not ready", http.StatusServiceUnavailable)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(art); err != nil {
			http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		}
	}
}
