// Package main implements the entry point for the GoHome server.
// It handles the initialization of core components and the HTTP server.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

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

	app := app.New(ctx, logger, "static") // TODO: add static dir configuration through env var
	if err = app.LaunchServer(); err != nil {
		return fmt.Errorf("error launching server: %w", err) // FIXME: add error handling
	}

	return nil
}
