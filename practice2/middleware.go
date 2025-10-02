package main

import (
	"log"
	"net/http"
)

const apiKey = "secret123"

func withAuthAndLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)

		if r.Header.Get("X-API-Key") != apiKey {
			writeJSON(w, http.StatusUnauthorized, ErrorResponse{Error: "unauthorized"})
			return
		}
		next.ServeHTTP(w, r)
	})
}
