package routes

import (
	"github.com/demartinom/museum-explorer/server/internal/daily"
	"github.com/demartinom/museum-explorer/server/internal/handlers"
	"github.com/demartinom/museum-explorer/server/internal/handlers/dailyhandlers"
	"github.com/demartinom/museum-explorer/server/internal/handlers/methandlers"
	"github.com/demartinom/museum-explorer/server/internal/museums/metmuseum"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Creates routes for different API calls on server startup
// Takes in clients from main.go and passes them into respective handlers
func RegisterRoutes(metClient *metmuseum.MetClient, dailyArt *daily.DailyArtworkManager) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Registers routes that belong to the met
	// They will all fall under /api/met
	r.Route("/api/metmuseum", func(r chi.Router) {
		r.Get("/departments", methandlers.DepartmentsHandler(metClient))
	})
	// Register daily routes
	// They will all fall under /api/daily
	r.Route("/api/daily", func(r chi.Router) {
		r.Get("/dailyartwork", dailyhandlers.DailyArtworkHandler(dailyArt))
	})
	// Route for general museum endpoints
	r.Route("/api/museums", func(r chi.Router) {
		// Returns list of all museums who's APIs are used
		r.Get("/", handlers.MuseumListHandler())
	})

	return r
}
