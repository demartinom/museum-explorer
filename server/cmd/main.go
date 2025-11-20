package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/demartinom/museum-explorer/server/pkg/routes"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Creates server at port 3000
	r := routes.RegisterRoutes()
	port := os.Getenv("PORT")
	fmt.Println("Server now running")
	http.ListenAndServe(port, r)
}
