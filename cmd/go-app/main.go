package main

import (
	"go-app/internal/bootstrap"
	"log"
)

func main() {
	server, err := bootstrap.SetupAndGetServer()
	if err != nil {
		log.Fatalf("Failed to setup server: %v", err)
	}
	log.Printf("Server starting on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
