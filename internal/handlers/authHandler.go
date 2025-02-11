package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/timut2/avito_test_task/internal/models"
	"github.com/timut2/avito_test_task/internal/service"
	"github.com/timut2/avito_test_task/internal/utils"
)

type AuthHandler struct {
	AuthService *service.AuthService
}

func NewAuthHandler(authSerice *service.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authSerice}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req models.AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := utils.ValidatePassword(req.Password); err != nil {
		http.Error(w, "Wrong format for password", http.StatusBadRequest)
		return
	}

	user, err := h.AuthService.Authentication(req.Username, req.Password)
	if err != nil {
		if errors.Is(err, service.ErrInvalidPassword) {
			http.Error(w, "Invalid password", http.StatusUnauthorized)
		}
		return
	}

	token, err := utils.GenerateJwt(uint(user.Id))
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.AuthResponse{Token: token})

}
