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
		h.Logger.Errorf("Failed to decode request body: %v", err)
		WriteErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := h.Validator.Struct(user); err != nil {
		HandleValidationError(w, err)
		return
	}

	ctx := r.Context()

	if err := database.Insert(ctx, &user); err != nil {
		h.Logger.Errorf("Failed to insert user: %v", err)
		WriteErrorResponse(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	WriteJSONResponse(w, http.StatusCreated, user)

}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var user types.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.Logger.Errorf("Failed to decode request body: %v", err)
		WriteErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := h.Validator.Struct(user); err != nil {
		HandleValidationError(w, err)
		return
	}

	ctx := r.Context()

	if err := database.Delete(ctx, &user); err != nil {
		h.Logger.Errorf("Failed to delete user: %v", err)
		WriteErrorResponse(w, http.StatusInternalServerError, "Failed to delete user")
		return
	}

	WriteJSONResponse(w, http.StatusOK, map[string]string{
		"message": "User deleted successfully",
	})

}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := database.GetUsers(ctx)
	if err != nil {
		h.Logger.Errorf("Failed to get users: %v", err)
		WriteErrorResponse(w, http.StatusInternalServerError, "Failed to get users")
	}

	WriteJSONResponse(w, http.StatusOK, users)
}
