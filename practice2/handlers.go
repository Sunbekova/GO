package main

import (
	"net/http"
	"strconv"
)

func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUserHandler(w, r)
	case http.MethodPost:
		createUserHandler(w, r)
	default:
		writeJSON(w, http.StatusMethodNotAllowed, ErrorResponse{Error: "method not allowed"})
	}
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || idStr == "" {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "invalid id"})
		return
	}
	writeJSON(w, http.StatusOK, UserResponse{UserID: id})
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name string `json:"name"`
	}

	if err := decodeJSON(r, &body); err != nil || body.Name == "" {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "invalid name"})
		return
	}

	writeJSON(w, http.StatusCreated, CreateUserResponse{Created: body.Name})
}
