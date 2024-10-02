package bootstrap

import (
	"fmt"
	"go-app/api/middlewares"
	"go-app/api/routers"
	"go-app/internal/config"
	"go-app/internal/database"
	_ "go-app/migrations"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func SetupAndGetServer() (*http.Server, error) {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Failed to load .env file\n")
	}
	cfg, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("Error loading config: %w", err)
	}

	db, err := database.SetupDB(cfg)
	if err != nil {
		return nil, fmt.Errorf("Error setting up database: %w", err)
	}

	chain := middlewares.CreateChain(
		middlewares.Logger,
		middlewares.APIKeyAuth,
	)

	mux := routers.SetupRoutes(db)

	serverAddr := fmt.Sprintf(":%d", cfg.ServerPort)
	server := &http.Server{
		Addr:    serverAddr,
		Handler: chain(mux),
	}
	return server, nil
}
