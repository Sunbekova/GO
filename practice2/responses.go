package main

import (
	"encoding/json"
	"net/http"
)

// safe JSON error payload
type ErrorResponse struct {
	Error string `json:"error"`
}

// Success responses
type UserResponse struct {
	UserID int `json:"user_id,omitempty"` //omit tag - do not show zero values from response
}

type CreateUserResponse struct {
	Created string `json:"created,omitempty"`
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func decodeJSON(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
