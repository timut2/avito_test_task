package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/timut2/avito_test_task/internal/middleware"
	"github.com/timut2/avito_test_task/internal/service"
	"github.com/timut2/avito_test_task/internal/utils"
)

type infoHandler struct {
	infoService *service.InfoService
}

func NewInfoHandler(infoService *service.InfoService) *infoHandler {
	return &infoHandler{infoService: infoService}
}

func (h *infoHandler) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.SendErrorResponse(w, http.StatusBadRequest, "недопустимый метод запроса")
		return
	}
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		utils.SendErrorResponse(w, http.StatusUnauthorized, "идентификатор пользователя не найден в контексте")
		return
	}

	log.Println("Идентификатор пользователя:", userID)

	info, err := h.infoService.GetInformation(userID)
	if err != nil {
		if errors.Is(err, service.ErrInvalidPassword) {
			utils.SendErrorResponse(w, http.StatusUnauthorized, "неверный пароль")
		} else {
			utils.SendErrorResponse(w, http.StatusInternalServerError, "не удалось получить информацию")
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(info)
	if err != nil {
		log.Printf("Error with decoding JSON forresponse: %v", err)
	}
}
