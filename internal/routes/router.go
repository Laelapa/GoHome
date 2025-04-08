package routes

import (
	"net/http"

	"github.com/Laelapa/GoHome/internal/routes/handlers"
	"go.uber.org/zap"
)

// Setup initializes and returns a configured router with all application routes
// as well as static file serving.
//
// Parameters:
//   - fsDir: The directory containing static files to serve
func Setup(staticDir string, logger *zap.SugaredLogger) *http.ServeMux {

	mux := http.NewServeMux()
	fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir)))
	mux.Handle("GET /static/", fileServer)
	mux.HandleFunc("GET /", handlers.HandleGetHome)
	mux.HandleFunc("GET /health", handlers.HandleGetHealth)

	return mux
}
