package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/vicolby/events/database"
	"github.com/vicolby/events/types"
	"go.uber.org/zap"
)

type LocationHandler struct {
	Logger    *zap.SugaredLogger
	Validator *validator.Validate
}

func NewLocationHandler(logger *zap.SugaredLogger, validator *validator.Validate) *LocationHandler {
	return &LocationHandler{
		Logger:    logger,
		Validator: validator,
	}
}

func (h *LocationHandler) CreateLocation(w http.ResponseWriter, r *http.Request) {
	var location types.Location

	if err := json.NewDecoder(r.Body).Decode(&location); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.Validator.Struct(location); err != nil {
		HandleValidationError(w, err)
		return
	}

	if err := database.Insert(&location); err != nil {
		http.Error(w, "Failed to create location", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(location)
}

func (h *LocationHandler) DeleteLocation(w http.ResponseWriter, r *http.Request) {
	var location types.Location

	if err := json.NewDecoder(r.Body).Decode(&location); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.Validator.Struct(location); err != nil {
		HandleValidationError(w, err)
		return
	}

	if err := database.Delete(&location); err != nil {
		http.Error(w, "Failed to delete location", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Location deleted successfully",
	})

}

func (h *LocationHandler) GetLocations(w http.ResponseWriter, r *http.Request) {
	locations, err := database.GetLocations()
	if err != nil {
		http.Error(w, "Failed to get locations", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(locations); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
