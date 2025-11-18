package handlers

import (
	"io"
	"net/http"
)

func MetDepartments(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://collectionapi.metmuseum.org/public/collection/v1/departments")
	if err != nil {
		http.Error(w, "failed to fetch Met API", http.StatusBadGateway)
		return
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
