package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"regexp"
	"unicode/utf8"

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

	if r.Method != http.MethodPost {
		utils.SendErrorResponse(w, http.StatusBadRequest, "неверный метод запроса")
		return
	}

	var req models.AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "неверный формат запроса")
		return
	}

	if len([]rune(req.Username)) < 5 {
		utils.SendErrorResponse(w, http.StatusBadRequest, "неверный формат имени пользователя")
		return
	}

	if err := ValidatePassword(req.Password); err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.AuthService.Authentication(req.Username, req.Password)
	if err != nil {
		if errors.Is(err, service.ErrInvalidPassword) {
			utils.SendErrorResponse(w, http.StatusUnauthorized, "неверный пароль")
		} else {
			utils.SendErrorResponse(w, http.StatusInternalServerError, "произошла ошибка")
		}
		return
	}

	token, err := utils.GenerateJWT(user.Id)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "ошибка при генерации токена")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(models.AuthResponse{Token: token})
	if err != nil {
		log.Printf("Error with decoding JSON forresponse: %v", err)
	}
}

var (
	ErrPasswordTooShort      = errors.New("пароль должен содержать не менее 8 символов")
	ErrPasswordNoDigit       = errors.New("пароль должен содержать хотя бы одну цифру")
	ErrPasswordNoSpecialChar = errors.New("пароль должен содержать хотя бы один специальный символ (!@#$%^&*)")
)

func ValidatePassword(password string) error {
	if utf8.RuneCountInString(password) < 8 {
		return ErrPasswordTooShort
	}

	if !regexp.MustCompile(`[0-9]`).MatchString(password) {
		return ErrPasswordNoDigit
	}

	if !regexp.MustCompile(`[!@#$%^&*]`).MatchString(password) {
		return ErrPasswordNoSpecialChar
	}

	return nil
}
