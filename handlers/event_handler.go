package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/vicolby/events/database"
	"github.com/vicolby/events/types"
	"go.uber.org/zap"
)

type EventHandler struct {
	Logger    *zap.SugaredLogger
	Validator *validator.Validate
}

func NewEventHandler(logger *zap.SugaredLogger, validator *validator.Validate) *EventHandler {
	return &EventHandler{
		Logger:    logger,
		Validator: validator,
	}
}

func (h *EventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event types.Event

	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.Validator.Struct(event); err != nil {
		HandleValidationError(w, err)
		return
	}

	if err := database.Insert(&event); err != nil {
		http.Error(w, "Failed to create event", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(event)
}

func (h *EventHandler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	var event types.Event

	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := database.Delete(&event); err != nil {
		http.Error(w, "Failed to delete event", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Event deleted successfully",
	})
}

func (h *EventHandler) GetEvents(w http.ResponseWriter, r *http.Request) {
	events, err := database.GetEvents()
	if err != nil {
		http.Error(w, "Failed to retrieve events", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(events); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *EventHandler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	var event types.Event

	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := database.UpdateEvent(&event); err != nil {
		http.Error(w, "Failed to update event", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Event updated successfully",
	})
}
