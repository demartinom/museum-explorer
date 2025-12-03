package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/demartinom/museum-explorer/server/internal/museums"
)

// Takes json list of museums and allows it to be sent to the frontend
func MuseumListHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		museums, err := museums.LoadMuseums("internal/museums/metmuseum/museumlist.json")
		if err != nil {
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(museums)
	}
}
