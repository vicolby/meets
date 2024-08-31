package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/vicolby/events/db"
	"github.com/vicolby/events/types"
)

func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	var event types.Event

	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.Insert(&event); err != nil {
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

	if err := db.Delete(&event); err != nil {
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

func UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	var event *types.Event

	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.UpdateEvent(event); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func AddEventParticipant(w http.ResponseWriter, r *http.Request) {
	var addReq types.AddParticipantReq
	eventParam := chi.URLParam(r, "eventID")
	eventParamInt, _ := strconv.Atoi(eventParam)

	if err := json.NewDecoder(r.Body).Decode(&addReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.AddEventParticipant(eventParamInt, addReq.UsersID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func DeleteEventParticipant(w http.ResponseWriter, r *http.Request) {
	var delReq types.DeleteParticipantReq
	eventParam := chi.URLParam(r, "eventID")
	eventParamInt, _ := strconv.Atoi(eventParam)

	if err := json.NewDecoder(r.Body).Decode(&delReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.DeleteEventParticipant(eventParamInt, delReq.UserID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
