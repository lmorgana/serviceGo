package main

import (
	"encoding/json"
	"net/http"
)

func sendErrorJSON(w http.ResponseWriter, statusCode int, message string, description string) error {
	data := struct {
		Error       bool
		Status      int
		Message     string
		Description string
	}{true, statusCode, message, description}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	return err
}

func checkSliceForInterval(max, min int, arr ...int) bool {
	for _, val := range arr {
		if val >= max || val < min {
			return false
		}
	}
	return true
}
