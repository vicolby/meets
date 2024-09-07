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
