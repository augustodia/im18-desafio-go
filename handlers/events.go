package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/augustodia/im18-desafio-go/models"
	"github.com/gorilla/mux"
)

func GetEvent(events []models.Event) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		for _, event := range events {
			if strconv.Itoa(event.ID) == params["eventId"] {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(event)
				return
			}
		}
		http.Error(w, "Event not found", http.StatusNotFound)
	}
}

func GetEvents(events []models.Event) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(events)
	}
}

func ReserveSpot(data *models.Data) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		eventID, err := strconv.Atoi(params["eventId"])
		if err != nil {
			http.Error(w, "Invalid event ID", http.StatusBadRequest)
			return
		}

		var reservation models.Reservation
		if err := json.NewDecoder(r.Body).Decode(&reservation); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		for _, spotName := range reservation.Spots {
			reserved := false
			for i, spot := range data.Spots {
				print(spot.EventID, eventID, spot.Name, spotName)
				if spot.EventID == eventID && spot.Name == spotName {
					if spot.Status != "available" {
						http.Error(w, "Spot "+spotName+" already reserved", http.StatusBadRequest)
						return
					}
					data.Spots[i].Status = "reserved"
					reserved = true
					break
				}
			}
			if !reserved {
				http.Error(w, "Spot "+spotName+" not found for event", http.StatusNotFound)
				return
			}
		}

		w.WriteHeader(http.StatusCreated)
	}
}
