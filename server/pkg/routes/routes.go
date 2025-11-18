package routes

import (
	"github.com/demartinom/museum-explorer/server/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RegisterRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("api/met/departments", handlers.MetDepartments)

	return r
}
