package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/augustodia/im18-desafio-go/models"
	"github.com/gorilla/mux"
)

func GetEventSpots(spots []models.Spot) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var eventSpots []models.Spot
		for _, spot := range spots {
			eventID, err := strconv.Atoi(params["eventId"])
			if err != nil {
				http.Error(w, "Invalid event ID", http.StatusBadRequest)
				return
			}
			if spot.EventID == eventID {
				eventSpots = append(eventSpots, spot)
			}
		}
		if len(eventSpots) == 0 {
			http.Error(w, "No spots found for this event", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(eventSpots)
	}
}
