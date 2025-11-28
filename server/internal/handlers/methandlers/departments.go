package methandlers

import (
	"encoding/json"
	"net/http"

	"github.com/demartinom/museum-explorer/server/internal/museums/metmuseum"
)

// Handler that makes API call to department endpoint on Met API
// Uses handlerfunc to pass in met client from main.go
func DepartmentsHandler(client *metmuseum.MetClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := client.GetDepartments()
		if err != nil {
			http.Error(w, "Failed to fetch Met API", http.StatusBadGateway)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) // optional, default is 200

		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		}
	}
}
