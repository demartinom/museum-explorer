package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Creates server at port 3000
	r := chi.NewRouter()
	port := os.Getenv("PORT")

	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Museum Explorer Backend"))
	})

	http.ListenAndServe(port, r)
}
