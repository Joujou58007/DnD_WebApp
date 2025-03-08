package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	// Define the handler
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Go backend!")
	})

	// Set up CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5000"},  // Frontend origin
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"}, // Allowed HTTP methods
		AllowedHeaders:   []string{"*"},                      // Allow all headers (adjust as needed)
		AllowCredentials: true,                               // Optional: if you need cookies or auth
	})

	// Wrap the default handler with CORS middleware
	handler := c.Handler(http.DefaultServeMux)

	log.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
