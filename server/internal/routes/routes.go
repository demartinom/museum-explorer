package routes

import (
	"github.com/demartinom/museum-explorer/server/internal/handlers/methandlers"
	"github.com/demartinom/museum-explorer/server/internal/museums/metmuseum"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Creates routes for different API calls on server startup
// Takes in metClient from main.go and passes it into handlers
func RegisterRoutes(metClient *metmuseum.MetClient) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Registers routes that belong to the met
	// They will all fall under /api/met
	r.Route("/api/met", func(r chi.Router) {
		r.Get("/departments", methandlers.DepartmentsHandler(metClient))
		r.Get("/random", methandlers.RandomHandler(metClient))
	})

	return r
}
