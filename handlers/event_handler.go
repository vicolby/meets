package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/vicolby/events/database"
	"github.com/vicolby/events/types"
)

func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	var event types.Event

	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.Insert(&event); err != nil {
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

	if err := database.Delete(&event); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}

func GetEventsHandler(w http.ResponseWriter, r *http.Request) {
	events, err := database.GetEvents()
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

	if err := database.UpdateEvent(event); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// func AddEventParticipant(w http.ResponseWriter, r *http.Request) {
// 	var addReq types.AddParticipantReq
// 	eventParam := chi.URLParam(r, "eventID")
// 	eventParamInt, _ := strconv.Atoi(eventParam)

// 	if err := json.NewDecoder(r.Body).Decode(&addReq); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	if err := db.AddEventParticipant(eventParamInt, addReq.UsersID); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// }

// func DeleteEventParticipant(w http.ResponseWriter, r *http.Request) {
// 	var delReq types.DeleteParticipantReq
// 	eventParam := chi.URLParam(r, "eventID")
// 	eventParamInt, _ := strconv.Atoi(eventParam)

// 	if err := json.NewDecoder(r.Body).Decode(&delReq); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	if err := db.DeleteEventParticipant(eventParamInt, delReq.UserID); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// }
