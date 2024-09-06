package routers

import (
	"database/sql"
	"go-app/api/handlers"
	"go-app/api/repositories"
	"go-app/api/services"
	"net/http"
)

func SetupUserRoutes(db *sql.DB) *http.ServeMux {
	userMux := http.NewServeMux()

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	userMux.HandleFunc("GET /{id}", userHandler.GetUser)

	return userMux
}
