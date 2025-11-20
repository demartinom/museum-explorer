package methandlers

import (
	"encoding/json"
	"net/http"

	"github.com/demartinom/museum-explorer/server/pkg/services/metservices"
)

var client = metservices.NewClient()

func DepartmentsHandler(w http.ResponseWriter, r *http.Request) {
	data, err := client.GetDepartments()
	if err != nil {
		http.Error(w, "Failed to fetch Met API", http.StatusBadGateway)
		return
	}

	// Tell the browser we're sending JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // optional, default is 200

	// Pretty-print the JSON for browser display
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
