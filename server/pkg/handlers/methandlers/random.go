package methandlers

import (
	"encoding/json"
	"net/http"
)

// Handler that makes API call return random artwork from the Met's highlights
func RandomHandler(w http.ResponseWriter, r *http.Request) {
	data, err := client.GetRandom()
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
