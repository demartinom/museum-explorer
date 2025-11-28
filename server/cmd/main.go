package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/demartinom/museum-explorer/server/internal/routes"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Creates server at port listed in .env file
	r := routes.RegisterRoutes()
	port := os.Getenv("PORT")

	fmt.Println("Server now running")
	http.ListenAndServe(port, r)
}
