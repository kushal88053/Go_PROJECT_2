package main

import (
	"fmt" // Import fmt for cleaner output
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kushal88053/Go_PROJECT_2/pkg/routes" // Ensure this path is correct
)

func main() {
	// Initialize the Gorilla Mux router
	router := mux.NewRouter()

	// Register all your application routes
	// Corrected typo: "registerRoutes" instead of "regirsterRoutes"
	routes.RegirsterRoutes(router)

	// Start the HTTP server
	port := ":8000"
	fmt.Printf("Server starting on port %s...\n", port) // Informative message

	err := http.ListenAndServe(port, router)

	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
