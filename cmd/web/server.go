package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func newServer() *http.Server {
	// Create a new logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		logger.Error("unable to read env file", "err", err)
		os.Exit(1)
	}

	// Convert the PORT environment variable to an integer
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		logger.Error("unable to convert port to int", "err", err)
		os.Exit(1)
	}

	// Setup the routes
	router := setupRoutes()

	// Create a new http.Server instance
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Info(fmt.Sprintf("server is running at http://localhost:%d", port))

	return server
}
