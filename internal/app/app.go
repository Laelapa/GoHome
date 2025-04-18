package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Laelapa/GoHome/internal/routes"
	"github.com/Laelapa/GoHome/internal/middleware"
	"go.uber.org/zap"
)

type serverOptions struct {
	shutdownTimeout time.Duration
}

type App struct {
	ctx           context.Context
	logger        *zap.SugaredLogger
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
	logger *zap.SugaredLogger,
	port string,
	staticDir string,
	shutdownTimeout time.Duration,
) *App {
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
func newMux(staticDir string, logger *zap.SugaredLogger) http.Handler {
	mux := routes.Setup(staticDir, logger)
	return attachBasicMiddleware(mux)
}

// attachBasicMiddleware wraps the provided handler with common middleware
// functions used across all routes.
func attachBasicMiddleware(handler http.Handler) http.Handler {

	handler = middleware.SecurityResponseHeaders(handler)
	handler = middleware.CacheControlHeader(handler)

	return handler
}

// SetServerShutdownTimeout configures the duration the server will wait
// during shutdown before forcefully terminating connections.
//
// Parameters:
//   - t: The duration to wait during shutdown in nanoseconds
func (app *App) SetServerShutdownTimeout(t time.Duration) {

	app.serverOptions.shutdownTimeout = t
	app.logger.Infof("Server shutdown timeout set to %vns\n", t)
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

		app.logger.Infof("Server running on %s", app.server.Addr)
		if err := app.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			app.logger.Errorf("Error by ListenAndServe(): %v\n", err)
			errChan <- err
		}
	}()

	select {
	case err := <-errChan:

		return fmt.Errorf("server failed to start: %v", err)

	case <-app.ctx.Done():

		app.logger.Infof("Shutting down server\n")
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
		app.logger.Errorf("Error during server shutdown: %v\n", err)
		app.logger.Infof("Closing server forcefully\n")
		app.server.Close()
	} else {
		app.logger.Infof("Server shut down successfully\n")
	}

}
