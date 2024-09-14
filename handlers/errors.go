package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func HandleValidationError(w http.ResponseWriter, err error) {
	errs := err.(validator.ValidationErrors)

	errorsResponse := make(map[string]string)
	for _, e := range errs {
		errorsResponse[e.Field()] = e.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "validation failed",
		"errors":  errorsResponse,
	})
}

func WriteErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func WriteJSONResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, "Failed to encode response")
	}
}
