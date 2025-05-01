package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Laelapa/GoHome/internal/middleware"
	"github.com/Laelapa/GoHome/internal/routes"
	"github.com/Laelapa/GoHome/logging"
	"go.uber.org/zap"
)

type serverOptions struct {
	shutdownTimeout time.Duration
}

type App struct {
	ctx           context.Context
	logger        *logging.Logger
	server        *http.Server
	serverOptions *serverOptions
}

// New creates and returns a new App instance with the provided dependencies.
// It initializes the HTTP server with default configuration and prepares it
// for handling requests.
//
// Parameters:
//   - ctx: The context for the application lifecycle
//   - logger: A configured zap logger for application logging
//   - port: The port on which the server will listen for incoming requests
//   - staticDir: The directory containing static files to serve
//   - shutdownTimeout: The duration to wait during shutdown before forcefully terminating
func New(
	ctx context.Context,
	logger *logging.Logger,
	port string,
	staticDir string,
	shutdownTimeout time.Duration,
) *App {

	if port == "" {
		logger.LogAppWarn("Port not specified, using default port 8080")
		port = "8080"
	}
	if staticDir == "" {
		logger.LogAppWarn("Static directory not specified, using default directory 'static'")
		staticDir = "static"
	}

	return &App{
		ctx:    ctx,
		logger: logger,
		server: &http.Server{
			Addr:    fmt.Sprintf(":%s", port),
			Handler: newMux(staticDir, logger),
		},
		serverOptions: &serverOptions{
			shutdownTimeout: shutdownTimeout,
		},
	}
}

// newMux creates and configures the HTTP request multiplexer with all routes
// and middleware attached.
func newMux(staticDir string, logger *logging.Logger) http.Handler {
	mux := routes.Setup(staticDir, logger)
	return attachBasicMiddleware(mux, logger)
}

// attachBasicMiddleware wraps the provided handler with common middleware
// functions used across all routes.
func attachBasicMiddleware(handler http.Handler, logger *logging.Logger) http.Handler {

	handler = middleware.SecurityResponseHeaders(handler)
	handler = middleware.CacheControlHeader(handler)
	handler = middleware.RequestLogger(handler, logger)

	return handler
}

// SetServerShutdownTimeout configures the duration the server will wait
// during shutdown before forcefully terminating connections.
//
// Parameters:
//   - t: The duration to wait during shutdown in nanoseconds
func (app *App) SetServerShutdownTimeout(t time.Duration) {

	app.serverOptions.shutdownTimeout = t
	app.logger.LogAppInfo(
		"Server shutdown timeout set",
		zap.Duration(logging.FieldDuration, t),
	)
}

// LaunchServer starts the HTTP server and manages its lifecycle. It will run
// until either a server error occurs or the application context is cancelled.
// When the context is cancelled, it triggers a graceful shutdown.
//
// Returns an error if the server fails to start or encounters an error while running.
func (app *App) LaunchServer() error {

	errChan := make(chan error, 1)
	defer close(errChan)

	go func() {

		app.logger.LogAppInfo(
			"Server running",
			zap.String(logging.FieldServerAddr, app.server.Addr),
		)
		if err := app.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			app.logger.LogAppError("Error thrown by ListenAndServe", err)
			errChan <- err
		}
	}()

	select {
	case err := <-errChan:

		return fmt.Errorf("server failed to start: %v", err)

	case <-app.ctx.Done():

		app.logger.LogAppInfo("Shutting down server")
		app.ShutdownServer()
		return nil
	}
}

// ShutdownServer attempts to gracefully shut down the HTTP server within the
// configured shutdown timeout duration. If graceful shutdown fails, it forces
// the server to close. The shutdown status is logged through the application logger.
func (app *App) ShutdownServer() {

	ctxServerShutdown, cancel := context.WithTimeout(context.Background(), app.serverOptions.shutdownTimeout)
	defer cancel()

	if err := app.server.Shutdown(ctxServerShutdown); err != nil && err != http.ErrServerClosed {
		app.logger.LogAppError("Error during server shutdown", err)
		app.logger.LogAppWarn("Closing server forcefully")
		app.server.Close()
	} else {
		app.logger.LogAppInfo("Server shut down successfully")
	}

}
