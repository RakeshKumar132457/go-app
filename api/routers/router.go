package routers

import (
	"database/sql"
	"net/http"
)

func SetupRoutes(db *sql.DB) *http.ServeMux {
	mainMux := http.NewServeMux()

	v1Mux := http.NewServeMux()
	v1Mux.Handle("/user/", http.StripPrefix("/user", SetupUserRoutes(db)))

	mainMux.Handle("/api/v1/", http.StripPrefix("/api/v1", v1Mux))

	return mainMux
}

