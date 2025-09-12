package middleware

import (
	"net/http"
	"slices"

	"github.com/gregriff/weather-api-go/internal/config"
)

const (
	allowedMethods = "GET, POST, PUT, DELETE, OPTIONS"
	allowedHeaders = "Content-Type, Authorization"
)

func NewCORSHandler(next http.Handler) http.Handler {
	cfg := config.Get()
	allowedOrigins := cfg.Api.AllowedOrigins
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If no origin (same-origin request), allow it
		if origin := r.Header.Get("Origin"); origin == "" {
			w.Header().Set("Access-Control-Allow-Methods", allowedMethods)
			w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
		} else {
			if slices.Contains(allowedOrigins, origin) {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Methods", allowedMethods)
				w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
			}

		}

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
