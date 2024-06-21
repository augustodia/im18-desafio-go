package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/augustodia/im18-desafio-go/handlers"
	"github.com/augustodia/im18-desafio-go/models"
	"github.com/gorilla/mux"
)

func main() {
	var data models.Data
	file, err := os.Open("data.json")
	if err != nil {
		log.Fatalf("failed to open data file: %v", err)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&data); err != nil {
		log.Fatalf("failed to decode data file: %v", err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/events", handlers.GetEvents(data.Events)).Methods("GET")
	r.HandleFunc("/events/{eventId}", handlers.GetEvent(data.Events)).Methods("GET")
	r.HandleFunc("/events/{eventId}/spots", handlers.GetEventSpots(data.Spots)).Methods("GET")
	r.HandleFunc("/event/{eventId}/reserve", handlers.ReserveSpot(&data)).Methods("POST")

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
