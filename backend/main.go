package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rs/cors"
)

type Message struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

func main() {
	initDB()

	mux := http.NewServeMux()

	// Serve static frontend
	mux.Handle("/", http.FileServer(http.Dir("../frontend")))

	// API endpoint
	mux.HandleFunc("/api/messages", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var msg Message
			err := json.NewDecoder(r.Body).Decode(&msg)
			if err != nil {
				http.Error(w, "Invalid request", http.StatusBadRequest)
				return
			}
			if msg.Name == "" || msg.Content == "" {
				http.Error(w, "Missing fields", http.StatusBadRequest)
				return
			}
			err = insertMessage(msg.Name, msg.Content)
			if err != nil {
				http.Error(w, "Could not save message", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
			return
		}

		if r.Method == http.MethodGet {
			messages, err := getAllMessages()
			if err != nil {
				http.Error(w, "Error fetching messages", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(messages)
			return
		}

		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	handler := cors.Default().Handler(mux)

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
