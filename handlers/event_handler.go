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
		h.Logger.Errorf("Ivalid request payload: %v", err)
		WriteErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := h.Validator.Struct(event); err != nil {
		h.Logger.Errorf("Validation error: %v", err)
		HandleValidationError(w, err)
		return
	}

	ctx := r.Context()

	if err := database.Insert(ctx, &event); err != nil {
		h.Logger.Errorf("Failed to insert event: %v", err)
		WriteErrorResponse(w, http.StatusInternalServerError, "Failed to create event")
		return
	}

	WriteJSONResponse(w, http.StatusCreated, event)

}

func (h *EventHandler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	var event types.Event

	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		h.Logger.Errorf("Failed to decode request body: %v", err)
		WriteErrorResponse(w, http.StatusBadRequest, "Ivalid request payload")
		return
	}

	ctx := r.Context()

	if err := database.Delete(ctx, &event); err != nil {
		h.Logger.Errorf("Failed to delete event: %v", err)
		WriteErrorResponse(w, http.StatusInternalServerError, "Failed to delete event")
		return
	}

	WriteJSONResponse(w, http.StatusOK, map[string]string{
		"message": "Event deleted successfully",
	})
}

func (h *EventHandler) GetEvents(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	events, err := database.GetEvents(ctx)
	if err != nil {
		h.Logger.Errorf("Failed to get events: %v", err)
		WriteErrorResponse(w, http.StatusInternalServerError, "Failed to get events")
		return
	}

	WriteJSONResponse(w, http.StatusOK, events)
}

func (h *EventHandler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	var event types.Event

	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		h.Logger.Errorf("Failed to decode request body: %v", err)
		WriteErrorResponse(w, http.StatusBadRequest, "Ivalid request payload")
		return
	}

	ctx := r.Context()

	if err := database.UpdateEvent(ctx, &event); err != nil {
		h.Logger.Errorf("Failed to update event: %v", err)
		WriteErrorResponse(w, http.StatusInternalServerError, "Failed to update event")
		return
	}

	WriteJSONResponse(w, http.StatusOK, map[string]string{
		"message": "Event updated successfully",
	})

}
