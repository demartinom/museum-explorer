package methandlers

// func DepartmentHighlightsHandler(client *metmuseum.MetClient) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		idStr := chi.URLParam(r, "id")

// 		departmentId, err := strconv.Atoi(idStr)
// 		if err != nil {
// 			return
// 		}

// 		data, err := client.DepartmentHighlights(departmentId)
// 		if err != nil {
// 			return
// 		}

// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)

// 		if err := json.NewEncoder(w).Encode(data); err != nil {
// 			http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
// 		}
// 	}

// }
