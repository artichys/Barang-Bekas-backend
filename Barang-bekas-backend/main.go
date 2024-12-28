package main

import (
	"log"
	"net/http"
	"Barang-bekas-backend/database"
	"Barang-bekas-backend/routes"
)

func main() {
	// Connect to the database
	database.Connect()

	// Setup routes
	router := routes.SetupRoutes()

	// Start the server
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
