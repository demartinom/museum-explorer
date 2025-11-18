package main

import (
	"net/http"
	"os"

	"github.com/demartinom/museum-explorer/server/pkg/routes"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	godotenv.Load()

	// Creates server at port 3000
	r := routes.RegisterRoutes()
	port := os.Getenv("PORT")
	http.ListenAndServe(port, r)
}
