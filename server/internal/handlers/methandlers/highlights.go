package methandlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/demartinom/museum-explorer/server/internal/museums/metmuseum"
	"github.com/go-chi/chi/v5"
)

func DepartmentHighlightsHandler(c *metmuseum.MetClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")

		departmentID, err := strconv.Atoi(idStr)
		if err != nil {
			return
		}

		c.Mu.RLock()
		data := c.CachedHighlights[metmuseum.DepartmentIDToName[departmentID]]
		c.Mu.RUnlock()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, "failed to encode JSON", http.StatusInternalServerError)
		}

	}
}
