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
	"github.com/gregriff/weather-api-go/internal/v1/routes"
)

func Run() {
	cfg := config.Get()
	// db := database.Connect(cfg.DatabaseURL)
	// defer db.Close()

	// Initialize handlers with dependencies
	h := routes.New()

	mux := http.NewServeMux()
	setupRoutes(mux, h)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: mux,
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

	log.Printf("Server starting on %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}

func setupRoutes(mux *http.ServeMux, h *routes.Handler) {
	// mapbox endpoints
	mux.HandleFunc("POST /geocode/place", h.GeocodePlace)

	// nws endpoints
	mux.HandleFunc("POST /weather", h.TestForecast)
	mux.HandleFunc("POST /weather/gridpoints", h.TestGridpoints)
	mux.HandleFunc("POST /weather/forecast", h.GetForecast)
	mux.HandleFunc("POST /weather/forecast/hourly", h.GetHourlyForecast)
}
