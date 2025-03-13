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
		log.Fatal("FATAL: %v\n", err)
	}
}

// run initializes and orchestrates all components of the application:
//   - Sets up signal handling for graceful shutdown
//   - Loads environment variables
//   - Initializes the logger
//   - Launches the HTTP server
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
	
	app := app.New(ctx, logger)
	if err = app.LaunchServer(); err != nil {
		return fmt.Errorf("error launching server: %w", err) // FIXME: add error handling
	}

	return nil
}
