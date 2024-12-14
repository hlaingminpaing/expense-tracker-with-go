package main

import (
	"expense-tracker/config"
	"expense-tracker/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize database connection
	config.ConnectDB()

	// Create a new router
	r := mux.NewRouter()

	// Register routes
	routes.RegisterExpenseRoutes(r)

	// Serve static files for frontend
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("../frontend/"))))

	// Start the server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
