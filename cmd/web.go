package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
	"votingMicroservicesApp/pkg/handlers"
)

// web is the entry point for the web mode of the application.
func web() {
	// Create a context that will listen for interrupt signals (SIGINT, SIGTERM).
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Create a new HTTP server with the router created by SetupRouter.
	srv := &http.Server{
		Addr:    ":8080",
		Handler: handlers.SetupRouter(),
	}

	// Start the HTTP server in a goroutine.
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start HTTP server: %s\n", err)
		}
	}()

	// Wait for the interrupt signal.
	<-ctx.Done()
	stop()
	log.Println("Shutting down gracefully. Press Ctrl+C again to force.")

	// Create a new context with a 5-second timeout to force close the server.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the server.
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Failed to shut down HTTP server gracefully: ", err)
	}

	log.Println("HTTP server exiting.")
}
