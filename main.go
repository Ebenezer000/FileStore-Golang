package main

import (
	"filestore-golang/api/routes"
	"filestore-golang/config"
	"log"
	"net/http"
)

func main() {
	// Initialize the database connection
	config.ConnectDB()

	// Initialize the routes
	router := routes.InitRoutes()

	// Start the server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
