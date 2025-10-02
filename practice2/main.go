package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/user", withAuthAndLogging(http.HandlerFunc(userHandler)))

	srv := newServer(":8080", mux)

	// Graceful shutdown
	idleConnsClosed := make(chan struct{})
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

		<-sigCh
		log.Println("Shutdown signal received, shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("Server shutdown error: %v", err)
		}
		close(idleConnsClosed)
	}()

	log.Printf("Server listening on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Server error: %v", err)
	}

	<-idleConnsClosed
	log.Println("Server stopped cleanly")
}
