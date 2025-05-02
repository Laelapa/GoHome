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

	"go.uber.org/zap"

	"github.com/Laelapa/GoHome/internal/app"
	"github.com/Laelapa/GoHome/logging"
)

const (
	DefaultShutdownTimeout = 5 * time.Second // Time until forceful shutdown
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

	// FIXME: add error handling, automate loading of .env file or pulling from secrets manager
	//
	// Disabled for production, use your service's secrets instead
	// Refer to the dotenv example file for the required environment variables
	// Uncomment the following lines to load environment variables from a .env file in a local development environment
	//
	// err := godotenv.Load()
	// if err != nil {
	// 	return fmt.Errorf("error loading .env file: %w", err)
	// }

	logger, err := logging.NewLogger(os.Getenv("ENVIRONMENT"))
	if err != nil {
		return fmt.Errorf("error creating logger: %w", err) // FIXME: add error handling
	}

	defer func() {
		if syncErr := logger.Sync(); syncErr != nil { // FIXME: handle case of writing to unbuffered output that doesnt support sync
			// Print the error without crashing the program
			fmt.Fprintf(os.Stderr, "Failed to sync logger: %v\n", syncErr)
		}
	}()

	// Parse the server shutdown timeout from the environment
	shutdownTimeout, err := time.ParseDuration(os.Getenv("SERVER_SHUTDOWN_TIMEOUT") + "s")
	if err != nil {
		shutdownTimeout = DefaultShutdownTimeout // fallback default
		logger.LogAppWarn(
			"Failed to parse SERVER_SHUTDOWN_TIMEOUT, falling back to default",
			zap.Duration("shutdown timeout", shutdownTimeout),
		)
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
