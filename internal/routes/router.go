package routes

import (
	"net/http"

	"github.com/Laelapa/GoHome/internal/routes/handlers"
)

// Setup initializes and returns a configured router with all application routes.
func Setup() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.HandleGetHome)
	mux.HandleFunc("GET /health", handlers.HandleGetHealth)

	return mux
}
