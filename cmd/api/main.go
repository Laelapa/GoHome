// Package main implements the entry point for the GoHome server.
// It handles the initialization of core components and the HTTP server.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/Laelapa/GoHome/internal/app"
	"github.com/Laelapa/GoHome/internal/logging"

	"github.com/joho/godotenv"
)

// main serves as the entry point for the application and acts as a thin wrapper
// around the run function. It will terminate the application with a fatal log
// if run encounters an error.
func main() {
	if err := run(); err != nil {
		log.Fatalf("FATAL: %v\n", err)
	}
}

// run initializes and orchestrates all components of the application:
//   - Sets up signal handling for graceful shutdown
//   - Loads environment variables
//   - Initializes the zap logger
//   - Launches the HTTP server and the static file server
//
// Returns an error if any initialization step fails.
func run() error {

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err) // FIXME: add error handling
	}

	logger, err := logging.NewLogger(os.Getenv("ENVIRONMENT"))
	if err != nil {
		return fmt.Errorf("error creating logger: %w", err) // FIXME: add error handling
	}

	defer func() {
		if syncErr := logger.Sync(); syncErr != nil {
			// Print the error without crashing the program
			fmt.Fprintf(os.Stderr, "Failed to sync logger: %v\n", syncErr)
		}
	}()

	// Parse the server shutdown timeout from the environment
	shutdownTimeout, err := time.ParseDuration(os.Getenv("SERVER_SHUTDOWN_TIMEOUT") + "s")
	if err != nil {
		shutdownTimeout = 5 * time.Second // fallback default
		logger.Warnf("Failed to parse SERVER_SHUTDOWN_TIMEOUT, using default: %v\n", shutdownTimeout)
	}

	app := app.New(
		ctx,
		logger,
		os.Getenv("SERVER_PORT"),
		os.Getenv("STATIC_DIR"), // FIXME: check if this is a valid path
		shutdownTimeout,
	)
	if err = app.LaunchServer(); err != nil {
		return fmt.Errorf("error launching server: %w", err) // FIXME: add error handling
	}

	return nil
}
