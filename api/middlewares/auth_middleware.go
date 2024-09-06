package middlewares

import (
	"net/http"
	"os"
)

func APIKeyAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")
		if apiKey == "" {
			http.Error(w, "Missing API key", http.StatusUnauthorized)
			return
		}
		validApiKey := os.Getenv("API_KEY")
		if validApiKey == "" {
			http.Error(w, "Server error: API key not configured", http.StatusInternalServerError)
			return
		}

		if apiKey != validApiKey {
			http.Error(w, "Invalid API key", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
