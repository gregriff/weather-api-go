package middleware

import (
	"log"
	"net/http"

	"github.com/gregriff/weather-api-go/internal/config"
)

// NewCSRFHandler configures and returns a handler that provides CSRF protection
func NewCSRFHandler(next http.Handler) http.Handler {
	cfg := config.Get()
	allowedOrigins := cfg.Api.AllowedOrigins
	if len(allowedOrigins) > 0 {
		log.Printf("Allowing trusted requests from %+v", allowedOrigins)
	}

	csrf := http.NewCrossOriginProtection()
	for _, origin := range allowedOrigins {
		csrf.AddTrustedOrigin(origin)
	}
	return csrf.Handler(next)
}
