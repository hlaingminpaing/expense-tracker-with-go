package main

import (
	"expense-tracker/config"
	"expense-tracker/routes"
	"log"
	"net/http"

	"github.com/rs/cors" //CORS middleware
)

func main() {
	// Connect to the database
	config.ConnectDB()

	// Set up routes
	router := routes.SetupRoutes()

	//Add CORS support
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost", "*"}, // Add your frontend's domain or * for all
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler(router)

	// Start the server
	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}
