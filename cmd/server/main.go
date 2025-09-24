package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gregriff/weather-api-go/internal/config"
	"github.com/gregriff/weather-api-go/internal/middleware"
	"github.com/gregriff/weather-api-go/internal/v1/routes"
)

func Run() {
	cfg := config.Get()
	// db := database.Connect(cfg.DatabaseURL)
	// defer db.Close()

	// Initialize handlers with dependencies
	h := routes.NewRouteHandler()

	mux := http.NewServeMux()
	createRoutes(mux, h)

	// apply middlewares
	var handler http.Handler
	if cfg.Debug {
		handler = middleware.DebugLogging(mux)
	} else {
		handler = mux
	}
	handler = middleware.NewCORSHandler(handler)
	handler = middleware.NewCSRFHandler(handler)

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.Api.Host, cfg.Api.Port),
		Handler: handler,
	}

	// Graceful shutdown
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		server.Shutdown(ctx)
	}()

	log.Printf("Starting server on %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}

// createRoutes creates the routing rules for the webserver
func createRoutes(mux *http.ServeMux, d *routes.RouteHandler) {
	// mapbox endpoints
	mux.HandleFunc("POST /v1/geocode/place", d.GeocodePlace)

	// nws endpoints
	mux.HandleFunc("POST /v1/weather/forecast", d.GetForecast)
	mux.HandleFunc("POST /v1/weather/forecast/hourly", d.GetHourlyForecast)
}
