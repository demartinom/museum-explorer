package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/demartinom/museum-explorer/server/internal/daily"
	"github.com/demartinom/museum-explorer/server/internal/museums/metmuseum"
	"github.com/demartinom/museum-explorer/server/internal/routes"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	//Create Global Met client
	metClient := metmuseum.NewClient()
	// Create daily art struct and begin daily refresh of art
	dailyArt := daily.NewDailyArtworkManager()
	dailyArt.Start(metClient)

	// Creates server at port listed in .env file
	r := routes.RegisterRoutes(metClient, dailyArt)
	port := os.Getenv("PORT")

	fmt.Println("Server now running")
	http.ListenAndServe(port, r)
}
