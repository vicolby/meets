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
		h.Logger.Errorf("Failed to decode request body: %v", err)
		WriteErrorResponse(w, http.StatusBadRequest, "Ivalid request payload")
		return
	}

	if err := h.Validator.Struct(location); err != nil {
		h.Logger.Errorf("Validation error: %v", err)
		HandleValidationError(w, err)
		return
	}

	ctx := r.Context()

	if err := database.Insert(ctx, &location); err != nil {
		h.Logger.Errorf("Failed to insert location: %v", err)
		WriteErrorResponse(w, http.StatusInternalServerError, "Failed to create location")
		return
	}

	WriteJSONResponse(w, http.StatusCreated, location)
}

func (h *LocationHandler) DeleteLocation(w http.ResponseWriter, r *http.Request) {
	var location types.Location

	if err := json.NewDecoder(r.Body).Decode(&location); err != nil {
		h.Logger.Errorf("Failed to decode request body: %v", err)
		WriteErrorResponse(w, http.StatusBadRequest, "Ivalid request payload")
		return
	}

	if err := h.Validator.Struct(location); err != nil {
		h.Logger.Errorf("Validation error: %v", err)
		HandleValidationError(w, err)
		return
	}

	ctx := r.Context()

	if err := database.Delete(ctx, &location); err != nil {
		h.Logger.Errorf("Failed to delete location: %v", err)
		WriteErrorResponse(w, http.StatusInternalServerError, "Failed to delete location")
		return
	}

	WriteJSONResponse(w, http.StatusOK, map[string]string{
		"message": "Location deleted successfully",
	})
}

func (h *LocationHandler) GetLocations(w http.ResponseWriter, r *http.Request) {
	locations, err := database.GetLocations()
	if err != nil {
		h.Logger.Errorf("Failed to get locations: %v", err)
		WriteErrorResponse(w, http.StatusInternalServerError, "Failed to get locations")
		return
	}
	WriteJSONResponse(w, http.StatusOK, locations)
}
