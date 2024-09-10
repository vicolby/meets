package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/vicolby/events/database"
	"github.com/vicolby/events/types"
	"go.uber.org/zap"
)

type UserHandler struct {
	Logger    *zap.SugaredLogger
	Validator *validator.Validate
}

func NewUserHandler(logger *zap.SugaredLogger, validator *validator.Validate) *UserHandler {
	return &UserHandler{
		Logger:    logger,
		Validator: validator,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user types.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.Validator.Struct(user); err != nil {
		HandleValidationError(w, err)
		return
	}

	if err := database.Insert(&user); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var user types.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.Validator.Struct(user); err != nil {
		HandleValidationError(w, err)
		return
	}

	if err := database.Delete(&user); err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User deleted successfully",
	})
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := database.GetUsers()
	if err != nil {
		http.Error(w, "Failed to get users", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
