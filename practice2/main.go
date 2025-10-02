package main

import (
	"errors"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/user", withAuthAndLogging(http.HandlerFunc(userHandler)))

	srv := newServer(":8080", mux)

	log.Printf("Server listening on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Server error: %v", err)
	}
}
