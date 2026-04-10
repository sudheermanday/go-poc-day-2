package main

import (
	"encoding/json"
	"net/http"

	"go-poc-day-2/model"
)

// getUsersHandler returns a list of users as JSON
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow GET requests
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Sample user data
	users := []model.User{
		{ID: 1, Username: "john_doe", Email: "john@example.com"},
		{ID: 2, Username: "jane_smith", Email: "jane@example.com"},
		{ID: 3, Username: "bob_wilson", Email: "bob@example.com"},
	}

	// Set response header to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode users to JSON and write response
	json.NewEncoder(w).Encode(users)
}

func main() {
	http.HandleFunc("/users", getUsersHandler)

	httpServer := http.Server{
		Addr: ":8080",
	}

	httpServer.ListenAndServe()
}
