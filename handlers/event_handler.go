package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/vicolby/events/db"
	"github.com/vicolby/events/types"
)

func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	var event types.Event

	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.InsertEvent(&event); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(event)
}

func DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	var event types.Event

	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.DeleteEvent(&event); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}

func GetEventsHandler(w http.ResponseWriter, r *http.Request) {
	events, err := db.GetEvents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(events)
}
