package routes

import (
	"github.com/demartinom/museum-explorer/server/pkg/handlers/methandlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Creates routes for different API calls on server startup
func RegisterRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Registers routes that belong to the met
	// They will all fall under /api/met
	r.Route("/api/met", func(r chi.Router) {
		r.Get("/departments", methandlers.DepartmentsHandler)
		r.Get("/random", methandlers.RandomHandler)
	})

	return r
}
