package routes

import (
	metHandlers "github.com/demartinom/museum-explorer/server/pkg/handlers/met"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RegisterRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Registers routes that belong to the met
	// They will all fall under /api/met
	r.Route("/api/met", func(r chi.Router) { r.Get("api/met/departments", metHandlers.MetDepartments) })

	return r
}
